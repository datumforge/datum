package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// ResetPassword allows the user (after requesting a password reset) to
// set a new password - the password reset token needs to be set in the request
// and not expired. If the request is successful, a confirmation of the reset is sent
// to the user and a 204 no content is returned
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
