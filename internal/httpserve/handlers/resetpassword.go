package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/marionette"
)

// ResetPasswordRequest contains fields required to reset a user's password
type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type ResetPasswordReply struct {
	Message string `json:"message"`
}

// ResetPassword allows the user (after requesting a password reset) to
// set a new password - the password reset token needs to be set in the request
// and not expired. If the request is successful, a confirmation of the reset is sent
// to the user and a 204 no content is returned
func (h *Handler) ResetPassword(ctx echo.Context) error {
	var in *ResetPasswordRequest

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		h.Logger.Errorw("error parsing request", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	if err := in.validateResetToken(); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	reqToken := ctx.QueryParam("token")

	// lookup user from db based on provided token
	entUser, err := h.getUserByResetToken(ctx.Request().Context(), reqToken)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving user token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrUnableToVerifyEmail))
	}

	// ent user to &User for funcs
	user := &User{
		ID:    entUser.ID,
		Email: entUser.Email,
	}

	// type ulid to string
	uid, err := ulid.Parse(entUser.ID)
	if err != nil {
		return err
	}

	// construct token from db fields
	token := &tokens.ResetToken{
		UserID: uid,
	}

	if token.ExpiresAt, err = user.GetPasswordResetExpires(); err != nil {
		h.Logger.Errorw("unable to parse expiration", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
	}

	// Verify the token with the stored secret
	if err = token.Verify(user.GetPasswordResetToken(), user.PasswordResetSecret); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrPassWordResetTokenInvalid)
	}

	if err := h.updateUserPassword(ctx.Request().Context(), entUser.ID, in.Password); err != nil {
		return err
	}

	if err := h.expireAllResetTokensUserByEmail(ctx.Request().Context(), user.Email); err != nil {
		h.Logger.Errorw("error expiring existing tokens", "error", err)

		return err
	}

	if err := h.TaskMan.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return h.SendPasswordResetSuccessEmail(user)
	}), marionette.WithRetries(3), // nolint: gomnd
		marionette.WithBackoff(backoff.NewExponentialBackOff()),
		marionette.WithErrorf("could not send password reset confirmation email to user %s", user.Email),
	); err != nil {
		return err
	}

	out := &ResetPasswordReply{
		Message: "password reset successfully",
	}

	return ctx.JSON(http.StatusOK, out)
}

// validateVerifyRequest validates the required fields are set in the user request
func (r *ResetPasswordRequest) validateResetToken() error {
	r.Password = strings.TrimSpace(r.Password)

	switch {
	case r.Token == "":
		return newMissingRequiredFieldError("token")
	case r.Password == "":
		return newMissingRequiredFieldError("password")
	case passwd.Strength(r.Password) < passwd.Moderate:
		return auth.ErrPasswordTooWeak
	}

	return nil
}
