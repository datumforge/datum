package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/marionette"
)

// VerifySubscribeRequest holds the fields that should be included on a request to the `/subscribe/verify` endpoint
type VerifySubscribeRequest struct {
	Token string `query:"token"`
}

// VerifySubscribeReply holds the fields that are sent on a response to the `/subscribe/verify` endpoint
type VerifySubscribeReply struct {
	rout.Reply
	Message string `json:"message,omitempty"`
}

// VerifySubscriptionHandler is the handler for the subscription verification endpoint
func (h *Handler) VerifySubscriptionHandler(ctx echo.Context) error {
	var req VerifySubscribeRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	if err := validateVerifySubscriptionRequest(req.Token); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// setup viewer context
	ctxWithToken := token.NewContextWithVerifyToken(ctx.Request().Context(), req.Token)

	entSubscriber, err := h.getSubscriberByToken(ctxWithToken, req.Token)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving subscriber", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrUnableToVerifyEmail))
	}

	// add org to the authenticated context
	reqCtx := auth.AddAuthenticatedUserContext(ctx, &auth.AuthenticatedUser{
		OrganizationID:  entSubscriber.OwnerID,
		OrganizationIDs: []string{entSubscriber.OwnerID},
	})

	ctxWithToken = token.NewContextWithVerifyToken(reqCtx, req.Token)

	if !entSubscriber.VerifiedEmail {
		if err := h.verifySubscriberToken(ctxWithToken, entSubscriber); err != nil {
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
			Email: &entSubscriber.Email,
		}

		if err := h.updateSubscriberVerifiedEmail(ctxWithToken, entSubscriber.ID, input); err != nil {
			h.Logger.Errorf("error updating subscriber", "error", err)

			return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrUnableToVerifyEmail))
		}
	}

	props := ph.NewProperties().
		Set("email", entSubscriber.Email)

	h.AnalyticsClient.Event("subscriber_verified", props)

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

// verifySubscriberToken checks the token provided by the user and verifies it against the database
func (h *Handler) verifySubscriberToken(ctx context.Context, entSubscriber *generated.Subscriber) error {
	// create User struct from entSubscriber
	user := &User{
		ID:                       entSubscriber.ID,
		Email:                    entSubscriber.Email,
		EmailVerificationSecret:  *entSubscriber.Secret,
		EmailVerificationToken:   sql.NullString{String: entSubscriber.Token, Valid: true},
		EmailVerificationExpires: sql.NullString{String: entSubscriber.TTL.Format(time.RFC3339Nano), Valid: true},
	}

	// setup token to be validated
	t := &tokens.VerificationToken{
		Email: entSubscriber.Email,
	}

	var err error
	t.ExpiresAt, err = user.GetVerificationExpires()

	if err != nil {
		h.Logger.Errorw("unable to parse expiration", "error", err)

		return ErrUnableToVerifyEmail
	}

	// verify token is valid, otherwise reset and send new token
	if err := t.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
		// if token is expired, create new token and send email
		if errors.Is(err, tokens.ErrTokenExpired) {
			if err := user.CreateVerificationToken(); err != nil {
				h.Logger.Errorw("error creating verification token", "error", err)

				return err
			}

			// update token settings in the database
			if err := h.updateSubscriberVerificationToken(ctx, user); err != nil {
				h.Logger.Errorw("error updating subscriber verification token", "error", err)

				return err
			}

			// set viewer context
			ctxWithToken := token.NewContextWithSignUpToken(ctx, entSubscriber.Email)

			// resend email with new token to the subscriber
			if err := h.sendSubscriberEmail(ctxWithToken, user, entSubscriber.OwnerID); err != nil {
				h.Logger.Errorw("error sending subscriber email", "error", err)

				return err
			}
		}

		return ErrExpiredToken
	}

	return nil
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

	props := ph.NewProperties().
		Set("user_id", user.ID).
		Set("email", user.Email).
		Set("first_name", user.FirstName).
		Set("last_name", user.LastName).
		Set("organization_name", orgName)

	h.AnalyticsClient.Event("email_verification_sent", props)

	return nil
}
