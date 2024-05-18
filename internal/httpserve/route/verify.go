package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// VerifyEmail verifies a user's email address by validating the token in the request.
// This endpoint is intended to be called by frontend applications after the user has
// followed the link in the verification email. If the user is not verified and the
// token is valid then the user is logged in. If the user is already verified then a
// 204 response is returned.
func registerVerifyHandler(router *Router) (err error) {
	path := "/verify"
	method := http.MethodGet

	route := echo.Route{
		Name:   "VerifyEmail",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.VerifyEmail(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW)

	router.AddRoute(path, method, nil, route)

	return
}
