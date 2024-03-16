package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"
	"github.com/samber/lo"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/tokens"
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

	// create email verification
	user := &User{
		ID:    entSubscriber.ID,
		Email: entSubscriber.Email,
	}

	if !entSubscriber.VerifiedEmail {
		// Construct the user token from the database fields
		t := &tokens.VerificationToken{
			Email: entSubscriber.Email,
		}

		user.EmailVerificationSecret = *entSubscriber.Secret
		user.EmailVerificationToken = sql.NullString{String: entSubscriber.Token, Valid: true}
		user.EmailVerificationExpires = sql.NullString{String: entSubscriber.TTL.Format(time.RFC3339Nano), Valid: true}

		if t.ExpiresAt, err = user.GetVerificationExpires(); err != nil {
			h.Logger.Errorw("unable to parse expiration", "error", err)

			return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
		}

		if err := h.verifySubscriberToken(ctxWithToken, t, user); err != nil {
			if errors.Is(err, ErrExpiredToken) {
				out := &VerifySubscribeReply{
					Reply:   rout.Reply{Success: false},
					Message: "The verification link has expired, a new one has been sent to your email.",
				}

				return ctx.JSON(http.StatusCreated, out)
			}

			h.Logger.Errorf("error verifying subscriber token", "error", err)

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
	}

	out := &VerifySubscribeReply{
		Reply:   rout.Reply{Success: true},
		Message: "Subscription confirmed, looking forward to sending you updates!",
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

func (h *Handler) verifySubscriberToken(ctx context.Context, t *tokens.VerificationToken, user *User) error {
	if err := t.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
		if errors.Is(err, tokens.ErrTokenExpired) {
			if err := user.CreateVerificationToken(); err != nil {
				h.Logger.Errorw("error creating verification token", "error", err)

				return err
			}

			if err := h.updateSubscriberVerificationToken(ctx, user); err != nil {
				h.Logger.Errorw("error updating subscriber verification token", "error", err)

				return err
			}

			if err := h.sendEmail(user); err != nil {
				h.Logger.Errorw("error sending subscriber email", "error", err)

				return err
			}

		}

		return ErrExpiredToken
	}

	return nil
}
