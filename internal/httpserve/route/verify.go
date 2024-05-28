package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerVerifyHandler registers the verify handler and route which handles email verification
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

	verifyOperation := router.Handler.BindVerifyEmailHandler()

	if err := router.AddRoute(path, method, verifyOperation, route); err != nil {
		return err
	}

	return nil
}
