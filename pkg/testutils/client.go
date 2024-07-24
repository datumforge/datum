package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/middleware/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	echo "github.com/datumforge/echox"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	server *echo.Echo
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.server.ServeHTTP(w, req)

	return w.Result(), nil
}

// DatumTestClient creates a new DatumClient for testing
func DatumTestClient(t *testing.T, c *generated.Client, opts ...datumclient.ClientOption) (*datumclient.DatumClient, error) {
	e := testEchoServer(t, c, false)

	// setup interceptors
	if opts == nil {
		opts = []datumclient.ClientOption{}
	}

	opts = append(opts, datumclient.WithTransport(localRoundTripper{server: e}))

	config := datumclient.NewDefaultConfig()

	return datumclient.New(config, opts...)
}

// DatumTestClientWithAuth creates a new DatumClient for testing that includes the auth middleware
func DatumTestClientWithAuth(t *testing.T, c *generated.Client, opts ...datumclient.ClientOption) (*datumclient.DatumClient, error) {
	e := testEchoServer(t, c, true)

	// setup interceptors
	if opts == nil {
		opts = []datumclient.ClientOption{}
	}

	opts = append(opts, datumclient.WithTransport(localRoundTripper{server: e}))

	config := datumclient.NewDefaultConfig()

	return datumclient.New(config, opts...)
}

// testEchoServer creates a new echo server for testing the graph api
// and optionally includes the middleware for authentication testing
func testEchoServer(t *testing.T, c *generated.Client, includeMiddleware bool) *echo.Echo {
	srv := testGraphServer(t, c)

	e := echo.New()

	if includeMiddleware {
		e.Use(echocontext.EchoContextToContextMiddleware())
		e.Use(auth.Authenticate(createAuthConfig(c)))
	}

	e.POST("/query", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		srv.ServeHTTP(res, req)

		return nil
	})

	return e
}

// createAuthConfig creates a new auth config for testing with the provided client
// and local validator
func createAuthConfig(c *generated.Client) *auth.AuthOptions {
	// setup auth middleware
	opts := []auth.AuthOption{
		auth.WithDBClient(c),
	}

	authConfig := auth.NewAuthOptions(opts...)

	authConfig.WithLocalValidator()

	return &authConfig
}

// testGraphServer creates a new graphql server for testing the graph api
func testGraphServer(t *testing.T, c *generated.Client) *handler.Server {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	srv := handler.NewDefaultServer(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(c).WithLogger(logger)},
		))

	// lower the cache size for testing
	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100), //nolint:mnd
	})

	// add all extension to the server
	graphapi.AddAllExtensions(srv)

	graphapi.WithTransactions(srv, c)

	// if you do not want sleeps (the writer prefers naps anyways), skip cache
	graphapi.WithSkipCache(srv)

	return srv
}
