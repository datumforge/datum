package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/rout"
)

// UnsubscribeReply holds the fields that are sent on a response to the `/unsubscribe` endpoint
type UnsubscribeReply struct {
	rout.Reply
	Message string `json:"message"`
}

// UnsubscribeHandler is responsible for handling requests to the `/unsubscribe` endpoint
// and removes subscribers from the mailing list
func (h *Handler) UnsubscribeHandler(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	if email == "" {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse("email is required"))
	}

	// organization, if null defaults to root level datum subscribers
	organizationID := ctx.QueryParam("organization_id")

	// set viewer context
	ctxWithToken := token.NewContextWithSignUpToken(ctx.Request().Context(), email)

	if err := h.deleteSubscriber(ctxWithToken, email, organizationID); err != nil {
		h.Logger.Errorw("error un user", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	out := &UnsubscribeReply{
		Reply:   rout.Reply{Success: true},
		Message: "We are sorry to see you go. You have been unsubscribed.",
	}

	return ctx.JSON(http.StatusCreated, out)
}
