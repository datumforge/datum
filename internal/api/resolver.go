// Package api is the graphapi package
package api

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/wundergraph/graphql-go-tools/pkg/playground"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var (
	graphPath      = "query"
	playgroundPath = "playground"

	graphFullPath = fmt.Sprintf("/%s", graphPath)
)

// Resolver provides a graph response resolver
type Resolver struct {
	client       *ent.Client
	logger       *zap.SugaredLogger
	authDisabled bool
}

// NewResolver returns a resolver configured with the given ent client
func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{
		client: client,
		// do not disabled auth by default
		authDisabled: false,
	}
}

func (r Resolver) WithLogger(l *zap.SugaredLogger) *Resolver {
	r.logger = l

	return &r
}

func (r Resolver) WithAuthDisabled(authDisabled bool) *Resolver {
	r.authDisabled = authDisabled

	return &r
}

// Handler is an http handler wrapping a Resolver
type Handler struct {
	r              *Resolver
	graphqlHandler http.Handler
	playground     *playground.Playground
	middleware     []echo.MiddlewareFunc
}

// Handler returns an http handler for a graph resolver
func (r *Resolver) Handler(withPlayground bool, middleware ...echo.MiddlewareFunc) *Handler {
	srv := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{
				Resolvers: r,
			},
		),
	)

	h := &Handler{
		r:              r,
		middleware:     middleware,
		graphqlHandler: srv,
	}

	if withPlayground {
		h.playground = playground.New(playground.Config{
			PathPrefix:          "/",
			PlaygroundPath:      playgroundPath,
			GraphqlEndpointPath: graphFullPath,
		})
	}

	return h
}

// Handler returns the http.HandlerFunc for the GraphAPI
func (h *Handler) Handler() http.HandlerFunc {
	return h.graphqlHandler.ServeHTTP
}

// Routes for the the server
func (h *Handler) Routes(e *echo.Group) {
	e.Use(h.middleware...)

	e.POST("/"+graphPath, func(c echo.Context) error {
		h.graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	if h.playground != nil {
		handlers, err := h.playground.Handlers()
		if err != nil {
			h.r.logger.Fatal("error configuring playground handlers", "error", err)
			return
		}

		for i := range handlers {
			// with the function we need to dereference the handler so that it remains
			// the same in the function below
			hCopy := handlers[i].Handler

			e.GET(handlers[i].Path, func(c echo.Context) error {
				hCopy.ServeHTTP(c.Response(), c.Request())
				return nil
			})
		}
	}
}
