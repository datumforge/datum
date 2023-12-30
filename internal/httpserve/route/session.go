package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

func registerSessionPutHandler(router *echo.Echo, h *handlers.Handler) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/sessionput",
		Handler: func(c echo.Context) error {
			return h.SessionPutHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

func registerSessionGetHandler(router *echo.Echo, h *handlers.Handler) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/sessionget",
		Handler: func(c echo.Context) error {
			return h.SessionGetHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}
