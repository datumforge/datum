package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerResetPasswordHandler registers the reset password handler and route
func registerResetPasswordHandler(router *Router) (err error) {
	path := "/password-reset"
	method := http.MethodPost

	route := echo.Route{
		Name:   "ResetPassword",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.ResetPassword(c)
		},
	}.ForGroup(V1Version, mw)

	resetOperation := router.Handler.BindResetPasswordHandler()

	if err := router.AddRoute(path, method, resetOperation, route); err != nil {
		return err
	}

	return nil
}
