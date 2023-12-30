package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// ForgotPassword is a service for users to request a password reset email. The email
// address must be provided in the POST request and the user must exist in the
// database. This endpoint always returns 204 regardless of whether the user exists or
// not to avoid leaking information about users in the database.
func registerForgotPasswordHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/forgot-password",
		Handler: func(c echo.Context) error {
			return h.ForgotPassword(c)
		},
	}.ForGroup(V1Version, mw))

	return
}
