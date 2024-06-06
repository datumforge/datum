package graphapi_test

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/datumforge/fgax"
	mock_fga "github.com/datumforge/fgax/mockery"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/httpsling"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/testutils"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/marionette"
	"github.com/datumforge/datum/pkg/utils/totp"
	"github.com/datumforge/datum/pkg/utils/ulids"

	"github.com/datumforge/datum/internal/graphapi"
)

var (
	testUser          *ent.User
	testPersonalOrgID string
)

// TestGraphTestSuite runs all the tests in the GraphTestSuite
func TestGraphTestSuite(t *testing.T) {
	suite.Run(t, new(GraphTestSuite))
}

// GraphTestSuite handles the setup and teardown between tests
type GraphTestSuite struct {
	suite.Suite
	client *client
	tf     *testutils.TestFixture
}

// client contains all the clients the test need to interact with
type client struct {
	db    *ent.Client
	datum datumclient.DatumGraphClient
	fga   *mock_fga.MockSdkClient
}

const (
	rawToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.oGFhqfFFDi9sJMJ1U2dWJZNYEiUQBEtZRVuwKE7Uiak"                                                                                  //nolint:gosec
	session  = "MTcwNTY0MjU5NnxkR1FweHFEX0RONDVzVDg0LTVuT3hLQmQ5THNicGJuZDk2dm8wbm5RMjZSdGFpY0xtcVBFdE1SR1IxT19IcTZhMzd1SWJBYldQWncwWlVmWGd6a0FzTDFMYlNjWkVJb3BRX1htM05qVjdOYS1hYy11SGo2aWRRcnFZYXRuWWJKXy1HNlF8AXSjkzY_IpNBe7u1T5YfHMcsKCwzdKKW2yeNbnmm_Z0=" // # spellcheck:off
)

type graphClient struct {
	srvURL     string
	httpClient *http.Client
}

func (suite *GraphTestSuite) SetupSuite() {
	suite.tf = entdb.NewTestFixture()
}

func (suite *GraphTestSuite) SetupTest() {
	t := suite.T()

	ctx := context.Background()

	// setup fga mock
	c := &client{
		fga: mock_fga.NewMockSdkClient(t),
	}

	// create mock FGA client
	fc := fgax.NewMockFGAClient(t, c.fga)

	// setup logger
	logger := zap.NewNop().Sugar()

	// setup email manager
	emConfig := emails.Config{
		Testing:   true,
		Archive:   filepath.Join("fixtures", "emails"),
		FromEmail: "mitb@datum.net",
	}

	em, err := emails.New(emConfig)
	if err != nil {
		t.Fatal("error creating email manager")
	}

	// setup task manager
	tmConfig := marionette.Config{
		Logger: zap.NewNop().Sugar(),
	}

	taskMan := marionette.New(tmConfig)

	taskMan.Start()

	// setup otp manager
	otpOpts := []totp.ConfigOption{
		totp.WithCodeLength(6),
		totp.WithIssuer("datum"),
		totp.WithSecret(totp.Secret{
			Version: 0,
			Key:     ulids.New().String(),
		}),
	}

	otpMan := totp.NewOTP(otpOpts...)

	opts := []ent.Option{
		ent.Logger(*logger),
		ent.Authz(*fc),
		ent.Emails(em),
		ent.Marionette(taskMan),
		ent.Analytics(&analytics.EventManager{Enabled: false}),
		ent.TOTP(&totp.Manager{
			TOTPManager: otpMan,
		}),
	}

	// create database connection
	db, err := entdb.NewTestClient(ctx, suite.tf, opts)
	require.NoError(t, err, "failed opening connection to database")

	// assign values
	c.db = db
	c.datum = graphTestClient(t, c.db)

	// create test user
	ctx = echocontext.NewTestContext()
	testUser = (&UserBuilder{client: c}).MustNew(ctx, t)
	testPersonalOrg, err := testUser.Edges.Setting.DefaultOrg(ctx)
	require.NoError(t, err)

	testPersonalOrgID = testPersonalOrg.ID

	suite.client = c
}

func (suite *GraphTestSuite) TearDownTest() {
	// clear all fga mocks
	mock_fga.ClearMocks(suite.client.fga)

	if suite.client.db != nil {
		if err := suite.client.db.Close(); err != nil {
			log.Fatalf("failed to close database: %s", err)
		}
	}
}

func (suite *GraphTestSuite) TearDownSuite() {
	testutils.TeardownFixture(suite.tf)
}

func graphTestClient(t *testing.T, c *ent.Client) datumclient.DatumGraphClient {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	srv := handler.NewDefaultServer(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(c).WithLogger(logger)},
		))

	graphapi.WithTransactions(srv, c)

	httpClient := &httpsling.Client{
		HTTPClient: &http.Client{Transport: localRoundTripper{handler: srv}},
	}

	// setup interceptors
	opts := []datumclient.ClientOption{
		datumclient.WithCredentials(
			datumclient.Authorization{
				BearerToken: rawToken,
				Session:     session,
			}),
		datumclient.WithClient(httpClient),
	}

	config := datumclient.NewDefaultConfig()

	client, err := datumclient.New(config, opts...)
	if err != nil {
		t.Fatalf("failed to create client: %s", err)

		return nil
	}

	return client
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

// userContext creates a new user in the database and returns a context with
// the user claims attached
func userContext() (context.Context, error) {
	return auth.NewTestContextWithOrgID(testUser.ID, testPersonalOrgID)
}

// userContextWithID creates a new user context with the provided user ID
func userContextWithID(userID string) (context.Context, error) {
	// Use that user to create the organization
	ec, err := auth.NewTestEchoContextWithValidUser(userID)
	if err != nil {
		return nil, err
	}

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	return reqCtx, nil
}

func (suite *GraphTestSuite) captureOutput(funcToRun func()) string {
	var buffer bytes.Buffer

	oldLogger := suite.client.db.Logger
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buffer)

	logger := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel)).
		Sugar()

	suite.client.db.Logger = *logger

	funcToRun()
	writer.Flush()

	suite.client.db.Logger = oldLogger

	return buffer.String()
}
