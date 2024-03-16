package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/samber/lo"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/rout"
)

// VerifySubscribeReply holds the fields that are sent on a response to the `/subscribe/verify` endpoint
type VerifySubscribeReply struct {
	rout.Reply
	Message string `json:"message,omitempty"`
}

// VerifySubscription is the handler for the subscription verification endpoint
func (h *Handler) VerifySubscription(ctx echo.Context) error {
	reqToken := ctx.QueryParam("token")

	if err := validateVerifySubscriptionRequest(reqToken); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// setup viewer context
	ctxWithToken := token.NewContextWithVerifyToken(ctx.Request().Context(), reqToken)

	entSubscriber, err := h.getSubscriberByToken(ctxWithToken, reqToken)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving subscriber", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrUnableToVerifyEmail))
	}

	input := generated.UpdateSubscriberInput{
		Email:         &entSubscriber.Email,
		VerifiedEmail: lo.ToPtr(true),
		Active:        lo.ToPtr(true),
	}

	if err := h.updateSubscriber(ctxWithToken, entSubscriber.ID, input); err != nil {
		h.Logger.Errorf("error updating subscriber", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrUnableToVerifyEmail))
	}

	out := &VerifySubscribeReply{
		Reply:   rout.Reply{Success: true},
		Message: "subscription confirmed",
	}

	return ctx.JSON(http.StatusOK, out)
}

// validateVerifySubscriptionRequest validates the required fields are set in the user request
func validateVerifySubscriptionRequest(token string) error {
	if token == "" {
		return rout.NewMissingRequiredFieldError("token")
	}

	return nil
}
