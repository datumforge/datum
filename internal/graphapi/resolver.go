package graphapi

import (
	"fmt"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	echo "github.com/datumforge/echox"
	"github.com/wundergraph/graphql-go-tools/pkg/playground"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

const (
	ActionGet    = "get"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionCreate = "create"
)

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
func NewResolver(client *ent.Client, authEnabled bool) *Resolver {
	return &Resolver{
		client: client,
		// do not disable auth by default
		authDisabled: !authEnabled,
	}
}

func (r Resolver) WithLogger(l *zap.SugaredLogger) *Resolver {
	r.logger = l

	return &r
}

// Handler is an http handler wrapping a Resolver
type Handler struct {
	r              *Resolver
	graphqlHandler *handler.Server
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

	// add transactional db client
	WithTransactions(srv, r.client)

	srv.Use(extension.Introspection{})

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

// WithTransactions adds the transactioner to the ent db client
func WithTransactions(h *handler.Server, c *ent.Client) {
	// setup transactional db client
	h.AroundOperations(injectClient(c))
	h.Use(entgql.Transactioner{TxOpener: c})
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
