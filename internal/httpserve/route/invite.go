package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

//	@Summary		Register Invite Handler
//	@Description	Registers the invite handler used by the UI to accept an invitation
//	@Tags			Invite
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	handlers.InviteReply
//	@Failure		400	{object}	route.ErrorResponse.BadRequest
//	@Failure		500	{object}	route.ErrorResponse.InternalServerError
//	@Router			/invite [post]
func registerInviteHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	// require authentication to accept an invitation
	authMW := mw
	authMW = append(authMW, h.AuthMiddleware...)
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/invite",
		Handler: func(c echo.Context) error {
			return h.OrganizationInviteAccept(c)
		},
	}.ForGroup(V1Version, authMW))

	return
}
