package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerRefreshHandler registers the refresh handler and route
func registerRefreshHandler(router *Router) (err error) {
	path := "/refresh"
	method := http.MethodPost

	route := echo.Route{
		Name:   "Refresh",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.RefreshHandler(c)
		},
	}.ForGroup(V1Version, mw)

	refreshOperation := router.Handler.BindRefreshHandler()

	if err := router.AddRoute(path, method, refreshOperation, route); err != nil {
		return err
	}

	return nil
}
