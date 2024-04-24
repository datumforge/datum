package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerEventPublisher registers the event publisher endpoint
func registerEventPublisher(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Name:   "EventPublisher",
		Method: http.MethodPost,
		Path:   "/event/publish",
		Handler: func(c echo.Context) error {
			return h.PublishEvent(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}
