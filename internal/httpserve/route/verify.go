package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getsentry/sentry-go"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	sentryecho "github.com/datumforge/datum/internal/utils/sentry"
)

// VerifyEmail verifies a user's email address by validating the token in the request.
// This endpoint is intended to be called by frontend applications after the user has
// followed the link in the verification email. If the user is not verified and the
// token is valid then the user is logged in. If the user is already verified then a
// 204 response is returned.
func registerVerifyHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/verify",
		Handler: func(c echo.Context) error {
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.WithScope(func(scope *sentry.Scope) {
					hub.CaptureMessage("verify handler")
				})
			}
			return h.VerifyEmail(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}
