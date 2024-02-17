package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/utils/marionette"
)

// ForgotPasswordRequest contains fields for a forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// ForgotPasswordReply contains fields for a forgot password response
// swagger:response ForgotPasswordReply
type ForgotPasswordReply struct {
	rout.Reply `json:"reply"`
	Message    string `json:"message,omitempty"`
}

// ForgotPassword will send an forgot password email if the provided
// email exists
func (h *Handler) ForgotPassword(ctx echo.Context) error {
	var in *ForgotPasswordRequest

	out := &ForgotPasswordReply{
		Reply: rout.Reply{
			Success: true,
		},
		Message: "We've received your request to have the password associated with this email reset. Please check your email.",
	}

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		h.Logger.Errorw("error parsing request", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	if err := validateForgotPasswordRequest(in); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	entUser, err := h.getUserByEmail(ctx.Request().Context(), in.Email, enums.Credentials)
	if err != nil {
		if ent.IsNotFound(err) {
			// return a 200 response even if user is not found to avoid
			// exposing confidential information
			return ctx.JSON(http.StatusOK, out)
		}

		h.Logger.Errorf("error retrieving user email", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	// create password reset email token
	user := &User{
		FirstName: entUser.FirstName,
		LastName:  entUser.LastName,
		Email:     entUser.Email,
		ID:        entUser.ID,
	}

	viewerCtx := viewer.NewContext(ctx.Request().Context(), viewer.NewUserViewerFromID(entUser.ID, true))

	if _, err = h.storeAndSendPasswordResetToken(viewerCtx, user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	return ctx.JSON(http.StatusOK, out)
}

// validateResendRequest validates the required fields are set in the user request
func validateForgotPasswordRequest(req *ForgotPasswordRequest) error {
	if req.Email == "" {
		return rout.NewMissingRequiredFieldError("email")
	}

	return nil
}

func (h *Handler) storeAndSendPasswordResetToken(ctx context.Context, user *User) (*ent.PasswordResetToken, error) {
	if err := h.expireAllResetTokensUserByEmail(ctx, user.Email); err != nil {
		h.Logger.Errorw("error expiring existing tokens", "error", err)

		return nil, err
	}

	if err := user.CreatePasswordResetToken(); err != nil {
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
