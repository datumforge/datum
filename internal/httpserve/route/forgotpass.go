package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerForgotPasswordHandler registers the forgot password handler and route
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

	forgotPasswordOperation := router.Handler.BindForgotPassword()

	if err := router.AddRoute(path, method, forgotPasswordOperation, route); err != nil {
		return err
	}

	return nil
}
