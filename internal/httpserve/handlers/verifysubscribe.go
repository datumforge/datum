package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/marionette"
)

// VerifySubscriptionHandler is the handler for the subscription verification endpoint
func (h *Handler) VerifySubscriptionHandler(ctx echo.Context) error {
	var in models.VerifySubscribeRequest
	if err := ctx.Bind(&in); err != nil {
		return h.BadRequest(ctx, err)
	}

	if err := in.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponseWithCode(err, InvalidInputErrCode))
	}

	// setup viewer context
	ctxWithToken := token.NewContextWithVerifyToken(ctx.Request().Context(), in.Token)

	entSubscriber, err := h.getSubscriberByToken(ctxWithToken, in.Token)
	if err != nil {
		if generated.IsNotFound(err) {
			return h.BadRequest(ctx, err)
		}

		h.Logger.Errorf("error retrieving subscriber", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrUnableToVerifyEmail))
	}

	// add org to the authenticated context
	reqCtx := auth.AddAuthenticatedUserContext(ctx, &auth.AuthenticatedUser{
		OrganizationID:  entSubscriber.OwnerID,
		OrganizationIDs: []string{entSubscriber.OwnerID},
	})

	ctxWithToken = token.NewContextWithVerifyToken(reqCtx, in.Token)

	if !entSubscriber.VerifiedEmail {
		if err := h.verifySubscriberToken(ctxWithToken, entSubscriber); err != nil {
			if errors.Is(err, ErrExpiredToken) {
				out := &models.VerifySubscribeReply{
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

	out := &models.VerifySubscribeReply{
		Reply:   rout.Reply{Success: true},
		Message: "Subscription confirmed, looking forward to sending you updates!",
	}

	return h.Success(ctx, out)
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

// BindVerifySubscriberHandler creates the openapi operation for the subscription verification endpoint
func (h *Handler) BindVerifySubscriberHandler() *openapi3.Operation {
	verify := openapi3.NewOperation()
	verify.Description = "Verify an email address for a subscription"
	verify.OperationID = "VerifySubscriberEmail"
	verify.Security = &openapi3.SecurityRequirements{}

	h.AddRequestBody("VerifySubscriptionEmail", models.ExampleVerifySubscriptionSuccessRequest, verify)
	h.AddResponse("VerifySubscriptionReply", "success", models.ExampleVerifySubscriptionResponse, verify, http.StatusOK)
	verify.AddResponse(http.StatusInternalServerError, internalServerError())
	verify.AddResponse(http.StatusBadRequest, badRequest())
	verify.AddResponse(http.StatusCreated, created())

	return verify
}
