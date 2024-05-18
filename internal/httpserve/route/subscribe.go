package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

func registerVerifySubscribeHandler(router *Router) (err error) {
	path := "/subscribe/verify"
	method := http.MethodGet

	route := echo.Route{
		Name:   "VerifySubscription",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.VerifySubscriptionHandler(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW)

	router.AddRoute(path, method, nil, route)

	return
}
