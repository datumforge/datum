package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/tokens"
	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"
)

func (h *Handler) VerifyEmail(ctx echo.Context) error {
	var (
		req *VerifyRequest
		err error
	)

	if err := json.NewDecoder(ctx.Request().Body).Decode(&req); err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return err
	}

	if req.Token == "" {
		ctx.JSON(http.StatusBadRequest, "missing token")
		return nil
	}

	user := &User{}

	// Look up the user by the token

	if user, err = GetUserByToken(ctx, req.Token, req.OrgID); err != nil {
		if errors.Is(err, ErrNotFound) {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse("invalid token"))
			return err
		}

		ctx.JSON(http.StatusInternalServerError, ErrorResponse("could not verify email"))
		return err
	}

	// check to see if user is already confirmed?
	if user.Edges.Setting.EmailConfirmed {
		auth.Unverified(ctx) //nolint:errcheck
		return err
	}

	// Construct the user token from the database fields
	token := &tokens.VerificationToken{
		Email: user.Email,
	}
	if token.ExpiresAt, err = user.GetVerificationExpires(); err != nil {
		ctx.JSON(http.StatusInternalServerError, "could not verify email")
		return err
	}

	// not sure if there's a better way to do this yet
	user := &User{}

	queryUser, err := h.DBClient.User.Query().WithSetting().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("token", user.Username))
	}).Only(ctx.Request().Context())
	if err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return err
	}

	// Verify the token with the stored secret
	if err = token.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
		if errors.Is(err, auth.ErrExpiredCredentials) {
			// If expired, create a new token for the user
			if err = user.CreateVerificationToken(); err != nil {
				ctx.JSON(http.StatusInternalServerError, "could not verify email")
				return err
			}

			if err = user.Save(ctx.Request.Context()); err != nil {
				ctx.JSON(http.StatusInternalServerError, "could not verify email")
				return err
			}

			// Send the new token to the user
			//s.tasks.Queue(marionette.TaskFunc(func(ctx context.Context) error {
			//	return s.SendVerificationEmail(user)
			//}),
			//	marionette.WithRetries(3),
			//	marionette.WithBackoff(backoff.NewExponentialBackOff()),
			//	marionette.WithErrorf("could not send verification email to user %s", user.ID.String()),
			//)

			ctx.JSON(http.StatusGone, "token expired, a new verification token has been sent to the email associated with the account")
			return err
		}

		ctx.JSON(http.StatusBadRequest, "invalid token")
		return err
	}

	// Mark user as verified so they can login
	user.EmailVerified = true
	if err = user.Save(c.Request.Context()); err != nil {
		ctx.JSON(http.StatusInternalServerError, "could not verify email")
		return err
	}

	// Issue claims to the user to log them in, this skips the password check so it
	// only happens the first time a user is verified
	// need to figure out how to do this
	var claims *tokens.Claims
	if claims, err = user.NewClaims(c.Request.Context()); err != nil {
		ctx.JSON(http.StatusNoContent, "huh?")
		return err
	}

	// Create a new access token/refresh token pair
	out := &LoginReply{}
	if out.AccessToken, out.RefreshToken, err = h.TM.CreateTokenPair(claims); err != nil {
		ctx.JSON(http.StatusNoContent, "huh?")
		return err
	}

	return ctx.JSON(http.StatusOK, out)
}

type VerifyRequest struct {
	Token string    `json:"token"`
	OrgID ulid.ULID `json:"org_id,omitempty"`
}

type LoginReply struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	LastLogin    string `json:"last_login,omitempty"`
}
