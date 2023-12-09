package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// Login is oriented towards human users who use their email and password for
// authentication (whereas authenticate is used for machine access using API keys).
// Login verifies the password submitted for the user is correct by looking up the user
// by email and using the argon2 derived key verification process to confirm the
// password matches. Upon authentication an access token and a refresh token with the
// authorized claims of the user are returned. The user can use the
// access token to authenticate to Datum systems. The access token has an expiration and
// the refresh token can be used with the refresh endpoint to get a new access token
// without the user having to log in again. The refresh token overlaps with the access
// token to provide a seamless authentication experience and the user can refresh their
// access token so long as the refresh token is valid.

// TODO: implement login handler
func registerLoginHandler(router *echo.Echo) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/login",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusNotImplemented, echo.Map{
				"error": "Not implemented",
			})
		},
	})

	return
}
