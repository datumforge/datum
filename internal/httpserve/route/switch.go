package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerSwitchRoute registers the switch route to switch the user's logged in organization context
func registerSwitchRoute(router *echo.Echo, h *handlers.Handler) (err error) {
	authMW := mw
	authMW = append(authMW, h.AuthMiddleware...)
	_, err = router.AddRoute(echo.Route{
		Name:   "Switch",
		Method: http.MethodPost,
		Path:   "/switch",
		Handler: func(c echo.Context) error {
			return h.SwitchHandler(c)
		},
	}.ForGroup(V1Version, authMW))

	return
}
