package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getsentry/sentry-go"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	sentryecho "github.com/datumforge/datum/internal/utils/sentry"
)

func registerInviteHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	// require authentication to accept an invitation
	authMW := mw
	authMW = append(authMW, h.AuthMiddleware...)
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/invite",
		Handler: func(c echo.Context) error {
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.WithScope(func(scope *sentry.Scope) {
					hub.CaptureMessage("invite handler")
				})
			}
			return h.OrganizationInviteAccept(c)
		},
	}.ForGroup(V1Version, authMW))

	return
}
