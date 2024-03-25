package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

func registerSubscribeHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Name:   "Subscribe",
		Method: http.MethodGet,
		Path:   "/subscribe",
		Handler: func(c echo.Context) error {
			return h.SubscribeHandler(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}

func registerVerifySubscribeHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Name:   "VerifySubscription",
		Method: http.MethodGet,
		Path:   "/subscribe/verify",
		Handler: func(c echo.Context) error {
			return h.VerifySubscriptionHandler(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}

func registerUnsubscribeHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Name:   "Unsubscribe",
		Method: http.MethodGet,
		Path:   "/unsubscribe",
		Handler: func(c echo.Context) error {
			return h.UnsubscribeHandler(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}
