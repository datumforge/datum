package graphapi_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Yamashou/gqlgenc/clientv2"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/testutils"
)

var (
	EntClient   *ent.Client
	DBContainer *testutils.TC
)

const (
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
	ctx := context.Background()

	// don't setup the datastore if we already have one
	if EntClient != nil {
		return
	}

	logger := zap.NewNop().Sugar()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")

	ctr := testutils.GetTestURI(ctx, testDBURI)
	DBContainer = ctr

	dbconf := entdb.Config{
		Debug:           true,
		DriverName:      ctr.Dialect,
		PrimaryDBSource: ctr.URI,
		CacheTTL:        -1 * time.Second, // do not cache results in tests
	}

	entConfig := entdb.NewDBConfig(dbconf, logger)

	opts := []ent.Option{ent.Logger(*logger)}

	c, err := entConfig.NewMultiDriverDBClient(ctx, opts)
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

	if DBContainer == nil {
		t.Fatalf("DBContainer is nil")
	}

	dbconf := entdb.Config{
		Debug:           true,
		DriverName:      DBContainer.Dialect,
		PrimaryDBSource: DBContainer.URI,
		CacheTTL:        -1 * time.Second, // do not cache results in tests
	}

	entConfig := entdb.NewDBConfig(dbconf, logger)

	ctx := context.Background()

	opts := []ent.Option{ent.Logger(*logger), ent.Authz(*fc)}

	c, err := entConfig.NewMultiDriverDBClient(ctx, opts)
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

func graphTestClient(t *testing.T, c *ent.Client) datumclient.DatumClient {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	srv := handler.NewDefaultServer(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(c, true).WithLogger(logger)},
		))

	graphapi.WithTransactions(srv, c)

	g := &graphClient{
		srvURL:     "query",
		httpClient: &http.Client{Transport: localRoundTripper{handler: srv}},
	}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := datumclient.WithAccessToken(rawToken)

	return datumclient.NewClient(g.httpClient, g.srvURL, opt, i)
}

func graphTestClientNoAuth(t *testing.T, c *ent.Client) datumclient.DatumClient {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	srv := handler.NewDefaultServer(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(c, false).WithLogger(logger)},
		))

	graphapi.WithTransactions(srv, c)

	g := &graphClient{
		srvURL:     "query",
		httpClient: &http.Client{Transport: localRoundTripper{handler: srv}},
	}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := datumclient.WithEmptyInterceptor()

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
			Writes: []ofgaclient.ClientWriteRequestWriteResponse{
				{
					Status: ofgaclient.SUCCESS,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, nil)
	} else {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteRequestWriteResponse{
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

func mockReadAny(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context) {
	mockExecute := mock_client.NewMockSdkClientReadRequestInterface(mockCtrl)

	resp := ofgaclient.ClientReadResponse{}

	mockExecute.EXPECT().Execute().Return(&resp, nil)

	mockRequest := mock_client.NewMockSdkClientReadRequestInterface(mockCtrl)

	mockRequest.EXPECT().Options(gomock.Any()).Return(mockExecute)

	c.EXPECT().Read(gomock.Any()).Return(mockRequest)
}
