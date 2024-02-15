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
//
// @route POST /forgot-password
// @tags Forgot Password
// @summary Request a password reset email
// @response 200 {object} ForgotPasswordResponse
// @response 400 {string} Bad Request
// @response 500 {string} Internal Server Error
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
