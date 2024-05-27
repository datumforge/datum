package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// Register creates a new user in the database with the specified password, allowing the
// user to login to Datum. This endpoint requires a "strong" password and a valid
// register request, otherwise a 400 reply is returned. The password is stored in the
// database as an argon2 derived key so it is impossible for a hacker to get access to
// raw passwordsf for that good good security.
//
// A "personal" organization is created for the user registering based on the organization data
// in the register request and the user is assigned the Owner role.
func registerRegisterHandler(router *Router) (err error) {
	path := "/register"
	method := http.MethodPost

	route := echo.Route{
		Name:   "Register",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.RegisterHandler(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW)

	registerOperation := router.Handler.BindRegisterHandler()

	if err := router.AddRoute(path, method, registerOperation, route); err != nil {
		return err
	}

	return nil
}
