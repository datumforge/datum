package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/marionette"
)

// ResetPasswordRequest contains user input required to reset a user's password
type ResetPasswordRequest struct {
	Password string `json:"password"`
}

// ResetPassword contains the full request to validate a password reset
type ResetPassword struct {
	Password string
	Token    string
}

// ResetPasswordReply is the response returned from a non-successful password reset request
// on success, no content is returned (204)
// swagger:response ResetPasswordReply
type ResetPasswordReply struct {
	rout.Reply
	Message string `json:"message"`
}

// ResetPassword allows the user (after requesting a password reset) to
// set a new password - the password reset token needs to be set in the request
// and not expired. If the request is successful, a confirmation of the reset is sent
// to the user and a 204 no content is returned
func (h *Handler) ResetPassword(ctx echo.Context) error {
	rp := &ResetPassword{
		Token: ctx.QueryParam("token"),
	}

	var in *ResetPasswordRequest

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		h.Logger.Errorw("error parsing request", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	// Add to the full request to be validated
	rp.Password = in.Password

	if err := rp.validateResetRequest(); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// setup viewer context
	ctxWithToken := token.NewContextWithResetToken(ctx.Request().Context(), rp.Token)

	// lookup user from db based on provided token
	entUser, err := h.getUserByResetToken(ctxWithToken, rp.Token)
	if err != nil {
		h.Logger.Errorf("error retrieving user token", "error", err)

		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(ErrPassWordResetTokenInvalid))
		}

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrUnableToVerifyEmail))
	}

	// ent user to &User for funcs
	user := &User{
		ID:    entUser.ID,
		Email: entUser.Email,
	}

	// set tokens for request
	if err := user.setResetTokens(entUser, rp.Token); err != nil {
		h.Logger.Errorw("unable to set reset tokens for request", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// Construct the user token from the database fields
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

	// Verify the token is valid with the stored secret
	if err = token.Verify(user.GetPasswordResetToken(), user.PasswordResetSecret); err != nil {
		if errors.Is(err, tokens.ErrTokenExpired) {
			errMsg := "reset token is expired, please request a new token using forgot-password"

			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(errMsg))
		}

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// make sure its not the same password as current
	valid, err := passwd.VerifyDerivedKey(*entUser.Password, rp.Password)
	if err != nil || valid {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(ErrNonUniquePassword))
	}

	// set context for remaining request based on logged in user
	userCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(user.ID, true))

	if err := h.updateUserPassword(userCtx, entUser.ID, rp.Password); err != nil {
		h.Logger.Errorw("error updating user password", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	if err := h.expireAllResetTokensUserByEmail(userCtx, user.Email); err != nil {
		h.Logger.Errorw("error expiring existing tokens", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	if err := h.TaskMan.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return h.SendPasswordResetSuccessEmail(user)
	}), marionette.WithRetries(3), // nolint: gomnd
		marionette.WithBackoff(backoff.NewExponentialBackOff()),
		marionette.WithErrorf("could not send password reset confirmation email to user %s", user.Email),
	); err != nil {
		h.Logger.Errorw("error sending confirmation email", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	out := &ResetPasswordReply{
		Reply:   rout.Reply{Success: true},
		Message: "password has been re-set successfully",
	}

	return ctx.JSON(http.StatusOK, out)
}

// validateVerifyRequest validates the required fields are set in the user request
func (r *ResetPassword) validateResetRequest() error {
	r.Password = strings.TrimSpace(r.Password)

	switch {
	case r.Token == "":
		return rout.NewMissingRequiredFieldError("token")
	case r.Password == "":
		return rout.NewMissingRequiredFieldError("password")
	case passwd.Strength(r.Password) < passwd.Moderate:
		return ErrPasswordTooWeak
	}

	return nil
}

// setResetTokens sets the fields for the password reset
func (u *User) setResetTokens(user *generated.User, reqToken string) error {
	tokens := user.Edges.PasswordResetTokens
	for _, t := range tokens {
		if t.Token == reqToken {
			u.PasswordResetToken = sql.NullString{String: t.Token, Valid: true}
			u.PasswordResetSecret = *t.Secret
			u.PasswordResetExpires = sql.NullString{String: t.TTL.Format(time.RFC3339Nano), Valid: true}

			return nil
		}
	}

	// This should only happen on a race condition with two request
	// otherwise, since we get the user by the token, it should always
	// be there
	return ErrPassWordResetTokenInvalid
}
