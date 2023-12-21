package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	dump "github.com/datumforge/datum/internal/httpserve/middleware/debug"
)

// registerJwksWellKnownHandler supplies the JWKS endpoint.
// This endpoint will contain the JWK used to verify all Datum JWTs
func registerJwksWellKnownHandler(router *echo.Echo, h *handlers.Handler) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/.well-known/jwks.json",
		Handler: func(c echo.Context) error {
			return h.JWKSWellKnownHandler(c)
		},
		Middlewares: []echo.MiddlewareFunc{middleware.Recover(), dump.BodyDump(logger)},
	})

	return
}
