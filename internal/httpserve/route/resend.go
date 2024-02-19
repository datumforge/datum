package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getsentry/sentry-go"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	sentryecho "github.com/datumforge/datum/internal/utils/sentry"
)

// ResendEmail accepts an email address via a POST request and always returns a 204
// response, no matter the input or result of the processing. This is to ensure that
// no secure information is leaked from this unauthenticated endpoint. If the email
// address belongs to a user who has not been verified, another verification email is
// sent. If the post request contains an orgID and the user is invited to that
// organization but hasn't accepted the invite, then the invite is resent.
func registerResendEmailHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/resend",
		Handler: func(c echo.Context) error {
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.WithScope(func(scope *sentry.Scope) {
					hub.CaptureMessage("resend handler")
				})
			}
			return h.ResendEmail(c)
		},
	}.ForGroup(V1Version, mw))

	return
}
