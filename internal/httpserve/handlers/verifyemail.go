package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

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

	user, err := h.GetUserByToken(ctx, req.Token)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, auth.ErrorResponse("invalid token"))
			return err
		}

		ctx.JSON(http.StatusInternalServerError, auth.ErrorResponse("could not verify email"))
		return err
	}

	meowuser, err := h.DBClient.User.Query().WithSetting().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("email", user.Email))
	}).Only(ctx.Request().Context())

	// check to see if user is already confirmed
	if meowuser.Edges.Setting.EmailConfirmed {
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

	// Verify the token with the stored secret
	if err = token.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
		if errors.Is(err, auth.ErrExpiredCredentials) {
			// If expired, create a new token for the user
			if err = user.CreateVerificationToken(); err != nil {
				ctx.JSON(http.StatusInternalServerError, "could not verify email")
				return err
			}

			ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)

			tx, err := h.DBClient.Tx(ctx.Request().Context())
			if err != nil {
				return fmt.Errorf("error starting transaction: %v", err)
			}

			meowtoken, err := tx.EmailVerificationToken.Create().
				SetOwnerID(meowuser.ID).
				SetToken(user.EmailVerificationToken.String).
				SetTTL(ttl).
				SetEmail(user.Email).
				SetSecret(user.EmailVerificationSecret).
				Save(ctx.Request().Context())
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}

				return err
			}

			if err = tx.Commit(); err != nil {
				return err
			}

			if err = h.SendVerificationEmail(user); err != nil {
				return err
			}

			var out *RegisterReply

			out = &RegisterReply{
				ID:      meowuser.ID,
				Email:   meowuser.Email,
				Message: "token expired, a new verification token has been sent to the email associated with the account",
				Token:   meowtoken.Token,
			}

			// Send the new token to the user
			//s.tasks.Queue(marionette.TaskFunc(func(ctx context.Context) error {
			//	return s.SendVerificationEmail(user)
			//}),
			//	marionette.WithRetries(3),
			//	marionette.WithBackoff(backoff.NewExponentialBackOff()),
			//	marionette.WithErrorf("could not send verification email to user %s", user.ID.String()),
			//)

			ctx.JSON(http.StatusGone, out)
			return err
		}

		ctx.JSON(http.StatusBadRequest, "invalid token")
		return err
	}

	// does this actually set it?
	meowuser.Edges.Setting.EmailConfirmed = true

	claims := createClaims(user)

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

func (h *Handler) GetUserByToken(ctx echo.Context, token string) (u *User, err error) {
	u = &User{
		EmailVerificationToken: sql.NullString{String: token, Valid: true},
	}

	tx, err := h.DBClient.Tx(ctx.Request().Context())
	if err != nil {
		return u, fmt.Errorf("error starting transaction: %v", err)
	}

	meowuser, err := h.DBClient.User.Query().WithChildren().WithSetting().QueryEmailVerificationTokens().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("token", u.EmailVerificationToken))
	}).Only(ctx.Request().Context())

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return u, err
		}

		return u, err
	}

	u.EmailVerificationSecret = meowuser.Secret
	u.Email = meowuser.Email

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return u, nil
}
