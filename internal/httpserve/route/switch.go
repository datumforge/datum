package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerSwitchRoute registers the switch route to switch the user's logged in organization context
func registerSwitchRoute(router *Router) (err error) {
	authMW := mw
	authMW = append(authMW, router.Handler.AuthMiddleware...)

	path := "/switch"
	method := http.MethodPost

	route := echo.Route{
		Name:   "Switch",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.SwitchHandler(c)
		},
	}.ForGroup(V1Version, authMW)

	router.AddRoute(path, method, nil, route)

	return
}
