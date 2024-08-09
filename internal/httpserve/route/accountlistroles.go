package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerAccountListRolesHandler registers the /account/list-roles handler
func registerAccountListRolesHandler(router *Router) (err error) {
	path := "/account/list-roles"
	method := http.MethodPost
	name := "AccountListRoles"

	route := echo.Route{
		Name:        name,
		Method:      method,
		Path:        path,
		Middlewares: authMW,
		Handler: func(c echo.Context) error {
			return router.Handler.AccountListRolesHandler(c)
		},
	}

	listRolesOperation := router.Handler.BindAccountListRoles()

	if err := router.Addv1Route(path, method, listRolesOperation, route); err != nil {
		return err
	}

	return nil
}
