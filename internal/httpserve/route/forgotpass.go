package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// ForgotPassword is a service for users to request a password reset email. The email
// address must be provided in the POST request and the user must exist in the
// database. This endpoint always returns 204 regardless of whether the user exists or
// not to avoid leaking information about users in the database.

// TODO: implement forgotpass handler
func registerForgotPasswordHandler(router *echo.Echo) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/forgot-password",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusNotImplemented, echo.Map{
				"error": "Not implemented",
			})
		},
	})

	return
}
