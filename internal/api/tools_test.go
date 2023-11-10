package api_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Yamashou/gqlgenc/clientv2"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/api"
	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
)

var (
	defaultDBURI = "file:ent?mode=memory&cache=shared&_fk=1"
	EntClient    *ent.Client
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

	logger := zap.NewNop()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")
	if testDBURI == "" {
		testDBURI = defaultDBURI
	}

	entConfig := entdb.EntClientConfig{
		Debug:           true,
		DriverName:      dialect.SQLite,
		Logger:          *logger.Sugar(),
		PrimaryDBSource: testDBURI,
	}

	ctx := context.Background()

	c, err := entConfig.NewEntDBDriver(ctx)
	if err != nil {
		errPanic("failed opening connection to database:", err)
	}

	errPanic("failed creating db scema", c.Schema.Create(ctx))
	EntClient = c
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

func graphTestClient() datumclient.DatumClient {
	g := &graphClient{
		srvURL: "query",
		httpClient: &http.Client{Transport: localRoundTripper{handler: handler.NewDefaultServer(
			api.NewExecutableSchema(
				api.Config{Resolvers: api.NewResolver(EntClient, zap.NewNop().Sugar())},
			))}},
	}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		// TODO: Add Auth Headers
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
