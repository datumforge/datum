package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerRefreshHandler re-authenticates users and api keys using a refresh token rather than
// requiring a username and password or API key credentials a second time and returns a
// new access and refresh token pair with the current credentials of the user. This
// endpoint is intended to facilitate long-running connections to datum systems that
// last longer than the duration of an access token; e.g. long sessions on the Datum UI
// or (especially) long running publishers and subscribers (machine users) that need to
// stay authenticated semi-permanently.

//	@Summary		Refresh authentication tokens
//	@Description	Re-authenticates users and API keys using a refresh token
//	@Tags			Refresh
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	handlers.RefreshReply
//	@Failure		400	{object}	route.ErrorResponse.BadRequest
//	@Failure		500	{object}	route.ErrorResponse.InternalServerError
//	@Router			/refresh [post]
func registerRefreshHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/refresh",
		Handler: func(c echo.Context) error {
			return h.RefreshHandler(c)
		},
	}.ForGroup(V1Version, mw))

	return
}
