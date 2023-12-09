package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// Refresh re-authenticates users and api keys using a refresh token rather than
// requiring a username and password or API key credentials a second time and returns a
// new access and refresh token pair with the current credentials of the user. This
// endpoint is intended to facilitate long-running connections to datum systems that
// last longer than the duration of an access token; e.g. long sessions on the Datum UI
// or (especially) long running publishers and subscribers (machine users) that need to
// stay authenticated semi-permanently.

// TODO: implement refresh handler
func registerRefreshHandler(router *echo.Echo) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/refresh",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusNotImplemented, echo.Map{
				"error": "Not implemented",
			})
		},
	})

	return
}
