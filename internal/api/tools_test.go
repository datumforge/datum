package api_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Yamashou/gqlgenc/clientv2"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/api"
	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
)

var (
	defaultDBURI = "file:ent?mode=memory&cache=shared&_fk=1"
	EntClient    *ent.Client
)

const (
	subClaim = "1234567890"
	rawToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.oGFhqfFFDi9sJMJ1U2dWJZNYEiUQBEtZRVuwKE7Uiak" //nolint:gosec
)

func TestMain(m *testing.M) {
	// setup the database if needed
	setupDB()
	// run the tests
	code := m.Run()
	// teardown the database
	teardownDB()
	// return the test response code
	os.Exit(code)
}

func setupDB() {
	// don't setup the datastore if we already have one
	if EntClient != nil {
		return
	}

	logger := zap.NewNop().Sugar()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")
	if testDBURI == "" {
		testDBURI = defaultDBURI
	}

	entConfig := entdb.EntClientConfig{
		Debug:           true,
		DriverName:      dialect.SQLite,
		Logger:          *logger,
		PrimaryDBSource: testDBURI,
	}

	ctx := context.Background()

	opts := []ent.Option{ent.Logger(*logger)}

	c, err := entConfig.NewEntDBDriver(ctx, opts)
	if err != nil {
		errPanic("failed opening connection to database:", err)
	}

	errPanic("failed creating db schema", c.Schema.Create(ctx))

	EntClient = c
}

func setupAuthEntDB(t *testing.T, mockCtrl *gomock.Controller, mc *mock_client.MockSdkClient) *ent.Client {
	fc, err := fga.NewTestFGAClient(t, mockCtrl, mc)
	if err != nil {
		t.Fatalf("enable to create test FGA client")
	}

	logger := zap.NewNop().Sugar()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")
	if testDBURI == "" {
		testDBURI = defaultDBURI
	}

	entConfig := entdb.EntClientConfig{
		Debug:           true,
		DriverName:      dialect.SQLite,
		Logger:          *logger,
		PrimaryDBSource: testDBURI,
	}

	ctx := context.Background()

	opts := []ent.Option{ent.Logger(*logger), ent.Authz(*fc)}

	c, err := entConfig.NewEntDBDriver(ctx, opts)
	if err != nil {
		errPanic("failed opening connection to database:", err)
	}

	errPanic("failed creating db schema", c.Schema.Create(ctx))

	return c
}

func teardownDB() {
	if EntClient != nil {
		errPanic("teardown failed to close database connection", EntClient.Close())
	}
}

func errPanic(msg string, err error) {
	if err != nil {
		log.Panicf("%s err: %s", msg, err.Error())
	}
}

type graphClient struct {
	srvURL     string
	httpClient *http.Client
}

func graphTestClient(c *ent.Client) datumclient.DatumClient {
	g := &graphClient{
		srvURL: "query",
		httpClient: &http.Client{Transport: localRoundTripper{handler: handler.NewDefaultServer(
			api.NewExecutableSchema(
				api.Config{Resolvers: api.NewResolver(c).WithLogger(zap.NewNop().Sugar())},
			))}},
	}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := datumclient.WithAccessToken(rawToken)

	return datumclient.NewClient(g.httpClient, g.srvURL, opt, i)
}

func graphTestClientNoAuth(c *ent.Client) datumclient.DatumClient {
	g := &graphClient{
		srvURL: "query",
		httpClient: &http.Client{Transport: localRoundTripper{handler: handler.NewDefaultServer(
			api.NewExecutableSchema(
				api.Config{Resolvers: api.NewResolver(c).WithLogger(zap.NewNop().Sugar()).WithAuthDisabled(true)},
			))}},
	}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		return next(ctx, req, gqlInfo, res)
	}

	return datumclient.NewClient(g.httpClient, g.srvURL, opt, i)
}

// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)

	return w.Result(), nil
}

// mockWriteTuples creates mock responses based on the mock FGA client
func mockWriteTuplesAny(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, errMsg error) {
	mockExecute := mock_client.NewMockSdkClientWriteTuplesRequestInterface(mockCtrl)

	if errMsg == nil {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteSingleResponse{
				{
					Status: ofgaclient.SUCCESS,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, nil)
	} else {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteSingleResponse{
				{
					Status: ofgaclient.FAILURE,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, errMsg)
	}

	mockRequest := mock_client.NewMockSdkClientWriteTuplesRequestInterface(mockCtrl)

	mockRequest.EXPECT().Options(gomock.Any()).Return(mockExecute)

	mockBody := mock_client.NewMockSdkClientWriteTuplesRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(gomock.Any()).Return(mockRequest)

	c.EXPECT().WriteTuples(gomock.Any()).Return(mockBody)
}

func mockCheckAny(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, allowed bool) {
	mockExecute := mock_client.NewMockSdkClientCheckRequestInterface(mockCtrl)

	resp := ofgaclient.ClientCheckResponse{
		CheckResponse: openfga.CheckResponse{
			Allowed: openfga.PtrBool(allowed),
		},
	}

	mockExecute.EXPECT().Execute().Return(&resp, nil)

	mockBody := mock_client.NewMockSdkClientCheckRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(gomock.Any()).Return(mockExecute)

	c.EXPECT().Check(gomock.Any()).Return(mockBody)
}

func mockListAny(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, allowedObjects []string) {
	mockExecute := mock_client.NewMockSdkClientListObjectsRequestInterface(mockCtrl)

	resp := ofgaclient.ClientListObjectsResponse{}
	resp.SetObjects(allowedObjects)

	mockExecute.EXPECT().Execute().Return(&resp, nil)

	mockBody := mock_client.NewMockSdkClientListObjectsRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(gomock.Any()).Return(mockExecute)

	c.EXPECT().ListObjects(gomock.Any()).Return(mockBody)
}
