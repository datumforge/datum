package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerVerifySubscribeHandler registers the verify subscription handler and route
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

	subscribeOperation := router.Handler.BindVerifySubscriberHandler()

	if err := router.AddRoute(path, method, subscribeOperation, route); err != nil {
		return err
	}

	return nil
}
