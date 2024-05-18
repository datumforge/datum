package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// ForgotPassword is a service for users to request a password reset email. The email
// address must be provided in the POST request and the user must exist in the
// database. This endpoint always returns 204 regardless of whether the user exists or
// not to avoid leaking information about users in the database.
func registerForgotPasswordHandler(router *Router) (err error) {
	path := "/forgot-password"
	method := http.MethodPost

	route := echo.Route{
		Name:   "ForgotPassword",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.ForgotPassword(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW)

	router.AddRoute(path, method, nil, route)

	return
}
