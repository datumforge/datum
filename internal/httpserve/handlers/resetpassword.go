package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/passwd"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/marionette"
)

// ResetPassword allows the user (after requesting a password reset) to
// set a new password - the password reset token needs to be set in the request
// and not expired. If the request is successful, a confirmation of the reset is sent
// to the user and a 204 no content is returned
func (h *Handler) ResetPassword(ctx echo.Context) error {
	var in models.ResetPasswordRequest
	if err := ctx.Bind(&in); err != nil {
		return h.BadRequest(ctx, err)
	}

	if err := in.ValidateResetRequest(); err != nil {
		return h.BadRequest(ctx, err)
	}

	// setup viewer context
	ctxWithToken := token.NewContextWithResetToken(ctx.Request().Context(), in.Token)

	// lookup user from db based on provided token
	entUser, err := h.getUserByResetToken(ctxWithToken, in.Token)
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
	if err := user.setResetTokens(entUser, in.Token); err != nil {
		h.Logger.Errorw("unable to set reset tokens for request", "error", err)

		return h.BadRequest(ctx, err)
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

		return h.BadRequest(ctx, err)
	}

	// make sure its not the same password as current
	valid, err := passwd.VerifyDerivedKey(*entUser.Password, in.Password)
	if err != nil || valid {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(ErrNonUniquePassword))
	}

	// set context for remaining request based on logged in user
	userCtx := auth.AddAuthenticatedUserContext(ctx, &auth.AuthenticatedUser{
		SubjectID: entUser.ID,
	})

	if err := h.updateUserPassword(userCtx, entUser.ID, in.Password); err != nil {
		h.Logger.Errorw("error updating user password", "error", err)

		return h.BadRequest(ctx, err)
	}

	if err := h.expireAllResetTokensUserByEmail(userCtx, user.Email); err != nil {
		h.Logger.Errorw("error expiring existing tokens", "error", err)

		return h.BadRequest(ctx, err)
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

	out := &models.ResetPasswordReply{
		Reply:   rout.Reply{Success: true},
		Message: "password has been re-set successfully",
	}

	return h.Success(ctx, out)
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

// BindResetPasswordHandler binds the reset password handler to the OpenAPI schema
func (h *Handler) BindResetPasswordHandler() *openapi3.Operation {
	resetPassword := openapi3.NewOperation()
	resetPassword.Description = "ResetPassword allows the user (after requesting a password reset) to set a new password - the password reset token needs to be set in the request and not expired. If the request is successful, a confirmation of the reset is sent to the user and a 200 StatusOK is returned"
	resetPassword.OperationID = "PasswordReset"
	resetPassword.Security = &openapi3.SecurityRequirements{}

	h.AddRequestBody("ResetPasswordRequest", models.ExampleResetPasswordSuccessRequest, resetPassword)
	h.AddResponse("ResetPasswordReply", "success", models.ExampleResetPasswordSuccessResponse, resetPassword, http.StatusOK)
	resetPassword.AddResponse(http.StatusInternalServerError, internalServerError())
	resetPassword.AddResponse(http.StatusBadRequest, badRequest())

	return resetPassword
}
