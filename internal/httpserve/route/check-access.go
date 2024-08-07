package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerCheckAccessHandler registers the check-access handler
func registerCheckAccessHandler(router *Router) (err error) {
	path := "/check-access"
	method := http.MethodPost
	name := "CheckAccess"

	route := echo.Route{
		Name:        name,
		Method:      method,
		Path:        path,
		Middlewares: authMW,
		Handler: func(c echo.Context) error {
			return router.Handler.CheckAccessHandler(c)
		},
	}

	checkAccessOperation := router.Handler.BindCheckAccess()

	if err := router.Addv1Route(path, method, checkAccessOperation, route); err != nil {
		return err
	}

	return nil
}
