package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerResendEmailHandler registers the resend email handler and route
func registerResendEmailHandler(router *Router) (err error) {
	path := "/resend"
	method := http.MethodPost

	route := echo.Route{
		Name:   "ResendEmail",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.ResendEmail(c)
		},
	}.ForGroup(V1Version, mw)

	resendOperation := router.Handler.BindResendEmailHandler()

	if err := router.AddRoute(path, method, resendOperation, route); err != nil {
		return err
	}

	return nil
}
