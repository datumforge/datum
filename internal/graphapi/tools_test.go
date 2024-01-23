package graphapi_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Yamashou/gqlgenc/clientv2"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"

	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/testutils"
)

var (
	dbContainer *testutils.TC
	testUser    *ent.User
)

// client contains all the clients the test need to intrract with
type client struct {
	db    *ent.Client
	datum datumclient.DatumClient
	fga   *mock_fga.MockSdkClient
}

const (
	rawToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.oGFhqfFFDi9sJMJ1U2dWJZNYEiUQBEtZRVuwKE7Uiak" //nolint:gosec
	session  = "MTcwNTY0MjU5NnxkR1FweHFEX0RONDVzVDg0LTVuT3hLQmQ5THNicGJuZDk2dm8wbm5RMjZSdGFpY0xtcVBFdE1SR1IxT19IcTZhMzd1SWJBYldQWncwWlVmWGd6a0FzTDFMYlNjWkVJb3BRX1htM05qVjdOYS1hYy11SGo2aWRRcnFZYXRuWWJKXy1HNlF8AXSjkzY_IpNBe7u1T5YfHMcsKCwzdKKW2yeNbnmm_Z0="
)

type graphClient struct {
	srvURL     string
	httpClient *http.Client
}

func setupTest(t *testing.T) *client {
	ctx := context.Background()

	// setup fga mock
	c := &client{
		fga: mock_fga.NewMockSdkClient(t),
	}

	// create mock FGA client
	fc := fga.NewMockFGAClient(t, c.fga)

	// setup logger
	logger := zap.NewNop().Sugar()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")

	ctr := testutils.GetTestURI(ctx, testDBURI)
	dbContainer = ctr

	dbconf := entdb.Config{
		Debug:           true,
		DriverName:      dbContainer.Dialect,
		PrimaryDBSource: dbContainer.URI,
		CacheTTL:        -1 * time.Second, // do not cache results in tests
	}

	entConfig := entdb.NewDBConfig(dbconf, logger)

	opts := []ent.Option{ent.Logger(*logger), ent.Authz(*fc)}

	db, err := entConfig.NewMultiDriverDBClient(ctx, opts)
	if err != nil {
		require.NoError(t, err, "failed opening connection to database")
	}

	if err := db.Schema.Create(ctx); err != nil {
		require.NoError(t, err, "failed creating database schema")
	}

	c.db = db
	c.datum = graphTestClient(t, c.db)

	// create test user
	testUser = (&UserBuilder{client: c}).MustNew(context.Background(), t)

	return c
}

func graphTestClient(t *testing.T, c *ent.Client) datumclient.DatumClient {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	srv := handler.NewDefaultServer(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(c).WithLogger(logger)},
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
	i := datumclient.WithAuthorization(rawToken, session)

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

// userContext creates a new user in the database and returns a context with
// the user claims attached
func userContext() (context.Context, error) {
	// Use that user to create the organization
	ec, err := auth.NewTestContextWithValidUser(testUser.ID)
	if err != nil {
		return nil, err
	}

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	return reqCtx, nil
}
