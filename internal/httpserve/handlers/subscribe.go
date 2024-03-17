package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/utils/marionette"
)

// SubscribeReply holds the fields that are sent on a response to the `/subscribe` endpoint
type SubscribeReply struct {
	rout.Reply
	Message string `json:"message"`
}

// SubscribeHandler is responsible for handling requests to the `/subscribe` endpoint
// It creates a new subscriber and sends a verification email to the subscriber
// this current only supports email subscriptions
func (h *Handler) SubscribeHandler(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	if email == "" {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse("email is required"))
	}

	// organization, if null defaults to root level datum subscribers
	organizationID := ctx.QueryParam("organization_id")

	// create user input for subscriber verification token
	user := &User{
		Email: email,
	}

	if err := user.CreateVerificationToken(); err != nil {
		h.Logger.Errorw("error creating verification token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse("could not create verification token"))
	}

	// create subscriber input
	input := generated.CreateSubscriberInput{
		Email: email,
	}

	if organizationID != "" {
		input.OwnerID = &organizationID
	}

	// set viewer context
	ctxWithToken := token.NewContextWithSignUpToken(ctx.Request().Context(), email)

	if _, err := h.createSubscriber(ctxWithToken, input, user); err != nil {
		h.Logger.Errorw("error creating new subscriber", "error", err)

		if IsConstraintError(err) || errors.Is(err, hooks.ErrUserAlreadySubscriber) {
			return ctx.JSON(http.StatusConflict, rout.ErrorResponse("email address is already subscribed"))
		}

		if generated.IsValidationError(err) {
			field := err.(*generated.ValidationError).Name
			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(fmt.Sprintf("%s was invalid", field)))
		}

		return err
	}

	if err := h.sendSubscriberEmail(ctxWithToken, user, organizationID); err != nil {
		h.Logger.Errorw("error sending subscriber email", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse("could not send subscriber email"))
	}

	out := &SubscribeReply{
		Reply:   rout.Reply{Success: true},
		Message: "Thank you for subscribing. Please check your email and click on the super sweet verification link.",
	}

	return ctx.JSON(http.StatusCreated, out)
}

func (h *Handler) sendSubscriberEmail(ctx context.Context, user *User, orgID string) error {
	// get org name if not root level (Datum)
	orgName := h.EmailManager.DefaultSubscriptionOrg

	if orgID != "" {
		org, err := h.getOrgByID(ctx, orgID)
		if err != nil {
			return err
		}

		orgName = org.Name
	}

	// send emails via TaskMan as to not create blocking operations in the server
	if err := h.TaskMan.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return h.SendSubscriberEmail(user, orgName)
	}), marionette.WithRetries(3), // nolint: gomnd
		marionette.WithBackoff(backoff.NewExponentialBackOff()),
		marionette.WithErrorf("could not send subscriber verification email to user %s", user.Email),
	); err != nil {
		return err
	}

	return nil
}
