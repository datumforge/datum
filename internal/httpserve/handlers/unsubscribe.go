package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/rout"
)

// UnsubscribeRequest holds the fields that should be included on a request to the `/unsubscribe` endpoint
type UnsubscribeRequest struct {
	Email          string `query:"email"`
	OrganizationID string `query:"organization_id" json:",omitempty"`
}

// UnsubscribeReply holds the fields that are sent on a response to the `/unsubscribe` endpoint
type UnsubscribeReply struct {
	rout.Reply
	Message string `json:"message"`
}

// UnsubscribeHandler is responsible for handling requests to the `/unsubscribe` endpoint
// and removes subscribers from the mailing list
func (h *Handler) UnsubscribeHandler(ctx echo.Context) error {
	var req UnsubscribeRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	if req.Email == "" {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse("email is required"))
	}

	// set viewer context
	ctxWithToken := token.NewContextWithSignUpToken(ctx.Request().Context(), req.Email)

	if err := h.deleteSubscriber(ctxWithToken, req.Email, req.OrganizationID); err != nil {
		h.Logger.Errorw("error un user", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	props := ph.NewProperties().
		Set("email", req.Email).
		Set("organization_id", req.OrganizationID)

	h.AnalyticsClient.Event("unsubscribe", props)

	out := &UnsubscribeReply{
		Reply:   rout.Reply{Success: true},
		Message: "We are sorry to see you go. You have been unsubscribed.",
	}

	return ctx.JSON(http.StatusOK, out)
}
