package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// @Summary Forgot Password
// @Description Allows the user to request a password reset email
// @Tags Forgot Password
// @Accept json
// @Produce json
// @Success 200 {string} handlers.ForgotPasswordReply
// @Failure 400 {object} handlers.StatusError "Status error object"
// @Failure 500 {object} handlers.StatusError "Status error object"
// @Router /forgot-password [get]

func registerForgotPasswordHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/forgot-password",
		Handler: func(c echo.Context) error {
			return h.ForgotPassword(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	if err != nil {
		return err
	}

	// we need to handle the redirect from /.well-known 302 -> /v1/forgot-password with a 200
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/forgot-password",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return c.JSON(http.StatusOK, "OK")
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	if err != nil {
		return err
	}

	return
}
