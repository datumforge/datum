package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerStaticHandler
func registerStaticHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/security.txt",
		Handler: func(c echo.Context) error {
			return h.SecurityHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}
