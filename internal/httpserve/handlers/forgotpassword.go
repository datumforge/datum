package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/utils/marionette"
)

// ForgotPasswordRequest contains fields for a forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// ForgotPassword will send an forgot password email if the provided
// email exists
func (h *Handler) ForgotPassword(ctx echo.Context) error {
	var in *ForgotPasswordRequest

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		h.Logger.Errorw("error parsing request", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	if err := validateForgotPasswordRequest(in); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// start transaction
	tx, err := h.DBClient.Tx(ctx.Request().Context())
	if err != nil {
		h.Logger.Errorw("error starting transaction", "error", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	entUser, err := h.getTokenUserByEmail(ctx.Request().Context(), tx, in.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			// return a 204 response even if user is not found to avoid
			// exposing confidential information
			return ctx.JSON(http.StatusNoContent, nil)
		}

		h.Logger.Errorf("error retrieving user email", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	// create password reset email token
	user := &User{
		FirstName: entUser.FirstName,
		LastName:  entUser.LastName,
		Email:     entUser.Email,
		ID:        entUser.ID,
	}

	if _, err = h.sentPasswordResetEmail(ctx.Request().Context(), tx, user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	// TODO: this will rollback on email failure, but FGA tuples will not get rolled back
	if err = tx.Commit(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

// validateResendRequest validates the required fields are set in the user request
func validateForgotPasswordRequest(req *ForgotPasswordRequest) error {
	if req.Email == "" {
		return newMissingRequiredFieldError("email")
	}

	return nil
}

func (h *Handler) sentPasswordResetEmail(ctx context.Context, tx *ent.Tx, user *User) (*ent.EmailVerificationToken, error) {
	if err := user.CreateVerificationToken(); err != nil {
		h.Logger.Errorw("unable to create verification token", "error", err)
		return nil, err
	}

	ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
	if err != nil {
		h.Logger.Errorw("unable to parse ttl", "error", err)
		return nil, err
	}

	meowtoken, err := tx.EmailVerificationToken.Create().
		SetOwnerID(user.ID).
		SetToken(user.EmailVerificationToken.String).
		SetTTL(ttl).
		SetEmail(user.Email).
		SetSecret(user.EmailVerificationSecret).
		Save(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return nil, err
		}

		h.Logger.Errorw("error creating email verification token", "error", err)

		return nil, err
	}

	// send emails via TaskMan as to not create blocking operations in the server
	if err := h.TaskMan.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return h.SendPasswordResetRequestEmail(user)
	}), marionette.WithRetries(3), // nolint: gomnd
		marionette.WithBackoff(backoff.NewExponentialBackOff()),
		marionette.WithErrorf("could not send verification email to user %s", user.ID),
	); err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return nil, err
		}

		return nil, err
	}

	return meowtoken, nil
}
