package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerWebauthnRegistrationHandler registers the webauthn registration handler
func registerWebauthnRegistrationHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/registration/options",
		Handler: func(c echo.Context) error {
			return h.BeginWebauthnRegistration(c)
		},
	}.ForGroup(V1Version, mw))

	return
}

// registerWebauthnVerificationsHandler registers the webauthn registration verification handler
func registerWebauthnVerificationsHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/registration/verification",
		Handler: func(c echo.Context) error {
			return h.FinishWebauthnRegistration(c)
		},
	}.ForGroup(V1Version, mw))

	return
}

// registerWebauthnAuthenticationHandler registers the webauthn authentication handler
func registerWebauthnAuthenticationHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/authentication/options",
		Handler: func(c echo.Context) error {
			return h.BeginWebauthnLogin(c)
		},
	}.ForGroup(V1Version, mw))

	return
}

// registerWebauthnAuthVerificationHandler registers the webauthn authentication verification handler
func registerWebauthnAuthVerificationHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/authentication/verification",
		Handler: func(c echo.Context) error {
			return h.FinishWebauthnLogin(c)
		},
	}.ForGroup(V1Version, mw))

	return
}
