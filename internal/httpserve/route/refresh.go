package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerRefreshHandler re-authenticates users and api keys using a refresh token rather than
// requiring a username and password or API key credentials a second time and returns a
// new access and refresh token pair with the current credentials of the user. This
// endpoint is intended to facilitate long-running connections to datum systems that
// last longer than the duration of an access token; e.g. long sessions on the Datum UI
// or (especially) long running publishers and subscribers (machine users) that need to
// stay authenticated semi-permanently.
func registerRefreshHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/refresh",
		Handler: func(c echo.Context) error {
			return h.RefreshHandler(c)
		},
		Middlewares: []echo.MiddlewareFunc{middleware.Recover()},
	})

	return
}
