package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

func registerInviteHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	// require authentication to accept an invitation
	authMW := mw
	authMW = append(authMW, h.AuthMiddleware...)
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/invite",
		Handler: func(c echo.Context) error {
			return h.OrganizationInviteAccept(c)
		},
	}.ForGroup(V1Version, authMW))

	return
}
