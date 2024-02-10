package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

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

func registerWebauthnVerificationsHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/registration/verifications",
		Handler: func(c echo.Context) error {
			return h.FinishWebauthnRegistration(c)
		},
	}.ForGroup(V1Version, mw))

	return
}

func registerWebauthnAuthenticationHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/authentication/options",
		Handler: func(c echo.Context) error {
			return h.BeginWebauthnLogin(c)
		},
	}.ForGroup(V1Version, mw))

	return
}

func registerWebauthnAuthVerificationHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/authentication/verifications",
		Handler: func(c echo.Context) error {
			return h.FinishWebauthnLogin(c)
		},
	}.ForGroup(V1Version, mw))

	return
}
