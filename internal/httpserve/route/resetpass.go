package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getsentry/sentry-go"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	sentryecho "github.com/datumforge/datum/internal/utils/sentry"
)

// ResetPassword allows the user (after requesting a password reset) to
// set a new password - the password reset token needs to be set in the request
// and not expired. If the request is successful, a confirmation of the reset is sent
// to the user and a 204 no content is returned
func registerResetPasswordHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/password-reset",
		Handler: func(c echo.Context) error {
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.WithScope(func(scope *sentry.Scope) {
					hub.CaptureMessage("resetpass handler")
				})
			}
			return h.ResetPassword(c)
		},
	}.ForGroup(V1Version, mw))

	return
}
