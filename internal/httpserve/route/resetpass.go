package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// ResetPassword allows the user (after requesting a password reset) to
// set a new password - the password reset token needs to be set in the request
// and not expired. If the request is successful, a confirmation of the reset is sent
// to the user and a 204 no content is returned

// @Summary Reset Password
// @Description Allows the user to reset their password
// @Tags Password Reset
// @Accept json
// @Produce json
// @Success 200 {object} handlers.ResetPasswordReply
// @Failure 400 {object} route.ErrorResponse.BadRequest
// @Failure 500 {object} route.ErrorResponse.InternalServerError
// @Router /password-reset [post]

func registerResetPasswordHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/password-reset",
		Handler: func(c echo.Context) error {
			return h.ResetPassword(c)
		},
	}.ForGroup(V1Version, mw))

	return
}
