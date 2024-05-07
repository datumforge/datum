package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

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
