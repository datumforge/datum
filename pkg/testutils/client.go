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
	"github.com/datumforge/datum/pkg/httpsling"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

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

// DatumTestClient creates a new DatumClient for testing
func DatumTestClient(t *testing.T, c *generated.Client) (*datumclient.DatumClient, error) {
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

	graphapi.WithTransactions(srv, c)

	// if you do not want sleeps (the writer prefers naps anyways), skip cache
	graphapi.WithSkipCache(srv)

	httpClient := &httpsling.Client{
		HTTPClient: &http.Client{Transport: localRoundTripper{handler: srv}},
	}

	// setup interceptors
	opts := []datumclient.ClientOption{
		datumclient.WithClient(httpClient),
	}

	config := datumclient.NewDefaultConfig()

	return datumclient.New(config, opts...)
}
