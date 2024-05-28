package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerRegisterHandler registers the register handler and route
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
