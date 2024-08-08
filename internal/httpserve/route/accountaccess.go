package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerAccountAccessHandler registers the /account/access handler
func registerAccountAccessHandler(router *Router) (err error) {
	path := "/account/access"
	method := http.MethodPost
	name := "AccountAccess"

	route := echo.Route{
		Name:        name,
		Method:      method,
		Path:        path,
		Middlewares: authMW,
		Handler: func(c echo.Context) error {
			return router.Handler.AccountAccessHandler(c)
		},
	}

	checkAccessOperation := router.Handler.BindCheckAccess()

	if err := router.Addv1Route(path, method, checkAccessOperation, route); err != nil {
		return err
	}

	return nil
}
