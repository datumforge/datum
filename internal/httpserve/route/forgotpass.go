package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/rout"
)

// @Summary Forgot Password
// @Description Allows the user to request a password reset email
// @Tags Forgot Password
// @Accept json
// @Produce json
// @Success 200 {object} handlers.ForgotPasswordReply
// @Failure 400 {object} route.ErrorResponse.BadRequest
// @Failure 500 {object} route.ErrorResponse.InternalServerError
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

type ErrorResponse struct {
	rout.StatusError
}

func (e *ErrorResponse) BadRequest() *ErrorResponse {
	out := &ErrorResponse{
		rout.StatusError{
			StatusCode: http.StatusBadRequest,
			Reply:      rout.Reply{Success: false, Error: "bad request"},
		}}
	return out
}

func (e *ErrorResponse) InternalServerError() *ErrorResponse {
	out := &ErrorResponse{
		rout.StatusError{
			StatusCode: http.StatusInternalServerError,
			Reply:      rout.Reply{Success: false, Error: "internal server error"},
		}}
	return out
}

func (e *ErrorResponse) Conflict() *ErrorResponse {
	out := &ErrorResponse{
		rout.StatusError{
			StatusCode: http.StatusConflict,
			Reply:      rout.Reply{Success: false, Error: "conflict"},
		}}
	return out
}

func (e *ErrorResponse) Unauthorized() *ErrorResponse {
	out := &ErrorResponse{
		rout.StatusError{
			StatusCode: http.StatusUnauthorized,
			Reply:      rout.Reply{Success: false, Error: "unauthorized"},
		}}
	return out
}
