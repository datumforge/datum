package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// ResetPassword allows users to set a new password after requesting a password reset.
// A token must be provided in the request and must not be expired. On success this
// endpoint sends a confirmation email to the user and returns a 204 No Content.

// TODO: implement login handler
func registerResetPasswordHandler(router *echo.Echo) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/reset-password",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusNotImplemented, echo.Map{
				"error": "Not implemented",
			})
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}
