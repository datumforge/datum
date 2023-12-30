package handlers

import (
	"context"
	"encoding/json"
	"net/http"

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

	entUser, err := h.getUserByEmail(ctx.Request().Context(), in.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			// return a 204 response even if user is not found to avoid
			// exposing confidential information
			return ctx.NoContent(http.StatusNoContent)
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

	if _, err = h.storeAndSendPasswordResetToken(ctx.Request().Context(), user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	return ctx.NoContent(http.StatusNoContent)
}

// validateResendRequest validates the required fields are set in the user request
func validateForgotPasswordRequest(req *ForgotPasswordRequest) error {
	if req.Email == "" {
		return newMissingRequiredFieldError("email")
	}

	return nil
}

func (h *Handler) storeAndSendPasswordResetToken(ctx context.Context, user *User) (*ent.PasswordResetToken, error) {
	if err := h.expireAllResetTokensUserByEmail(ctx, user.Email); err != nil {
		h.Logger.Errorw("error expiring existing tokens", "error", err)

		return nil, err
	}

	if err := user.CreateResetToken(); err != nil {
		h.Logger.Errorw("unable to create password reset token", "error", err)
		return nil, err
	}

	meowtoken, err := h.createPasswordResetToken(ctx, user)
	if err != nil {
		return nil, err
	}

	// send emails via TaskMan as to not create blocking operations in the server
	if err := h.TaskMan.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return h.SendPasswordResetRequestEmail(user)
	}), marionette.WithRetries(3), //nolint: gomnd
		marionette.WithBackoff(backoff.NewExponentialBackOff()),
		marionette.WithErrorf("could not send password reset email to user %s", user.ID),
	); err != nil {
		return nil, err
	}

	return meowtoken, nil
}
