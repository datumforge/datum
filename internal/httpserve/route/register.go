package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// Register creates a new user in the database with the specified password, allowing the
// user to login to Datum. This endpoint requires a "strong" password and a valid
// register request, otherwise a 400 reply is returned. The password is stored in the
// database as an argon2 derived key so it is impossible for a hacker to get access to
// raw passwordsf for that good good security.
//
// A "personal" organization is created for the user registering based on the organization data
// in the register request and the user is assigned the Owner role.

// @Summary		Register a new user
// @Description	Creates a new user in the database with the specified password
// @Tags			Register
// @Accept			json
// @Produce		json
// @Success		201	{object}	handlers.RegisterReply
// @Failure		400	{object}	route.ErrorResponse.BadRequest
// @Failure		409	{object}	route.ErrorResponse.Conflict
// @Failure		500	{object}	route.ErrorResponse.InternalServerError
// @Router			/register [post]
func registerRegisterHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/register",
		Handler: func(c echo.Context) error {
			return h.RegisterHandler(c)
		},
	}.ForGroup(V1Version, restrictedEndpointsMW))

	return
}
