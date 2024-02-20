package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// VerifyEmail verifies a user's email address by validating the token in the request.
// This endpoint is intended to be called by frontend applications after the user has
// followed the link in the verification email. If the user is not verified and the
// token is valid then the user is logged in. If the user is already verified then a
// 204 response is returned.
// @Summary		Verify Email
// @Description	Verifies a user's email address by validating the token in the request
// @Tags			Verify
// @Accept			json
// @Produce		json
// @Success		200	{object}	handlers.VerifyReply
// @Success		201	{object}	handlers.VerifyReply
// @Failure		400	{object}	route.ErrorResponse.BadRequest
// @Failure		500	{object}	route.ErrorResponse.InternalServerError
// @Router			/verify [get]
func registerVerifyHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/verify",
		Handler: func(c echo.Context) error {
			return h.VerifyEmail(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}
