package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/tokens"
)

func (h *Handler) VerifyEmail(ctx echo.Context) error {
	var (
		req *VerifyRequest
		err error
	)

	if err := json.NewDecoder(ctx.Request().Body).Decode(&req); err != nil {
		return auth.Unauthorized(ctx)
	}

	if req.Token == "" {
		return ctx.JSON(http.StatusBadRequest, "missing token")
	}

	user, err := h.GetUserByToken(ctx, req.Token)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return ctx.JSON(http.StatusBadRequest, auth.ErrorResponse("invalid token"))
		}

		return ctx.JSON(http.StatusInternalServerError, auth.ErrorResponse("could not verify email"))
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
		return ctx.JSON(http.StatusInternalServerError, "could not verify email")
	}

	// Verify the token with the stored secret
	if err = token.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
		if errors.Is(err, auth.ErrExpiredCredentials) {
			// If expired, create a new token for the user
			if err = user.CreateVerificationToken(); err != nil {
				return ctx.JSON(http.StatusInternalServerError, "could not verify email")
			}

			ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, "unable to parse ttl")
			}

			tx, err := h.DBClient.Tx(ctx.Request().Context())
			if err != nil {
				// TODO: error stuff
				return fmt.Errorf("error starting transaction: %v", err) //nolint:goerr113
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
				return ctx.JSON(http.StatusInternalServerError, auth.ErrorResponse(err))
			}

			if err = h.SendVerificationEmail(user); err != nil {
				return err
			}

			out := &RegisterReply{
				ID:      meowuser.ID,
				Email:   meowuser.Email,
				Message: "token expired, a new verification token has been sent to the email associated with the account",
				Token:   meowtoken.Token,
			}

			// Send the new token to the user
			// s.tasks.Queue(marionette.TaskFunc(func(ctx context.Context) error {
			//	return s.SendVerificationEmail(user)
			// }),
			//	marionette.WithRetries(3),
			//	marionette.WithBackoff(backoff.NewExponentialBackOff()),
			//	marionette.WithErrorf("could not send verification email to user %s", user.ID.String()),
			// )

			return ctx.JSON(http.StatusGone, out)
		}

		return ctx.JSON(http.StatusBadRequest, "invalid token")
	}

	// TODO: set email confirmed with db client
	meowuser.Edges.Setting.EmailConfirmed = true

	claims := createClaims(user)

	// Create a new access token/refresh token pair
	out := &LoginReply{}

	if out.AccessToken, out.RefreshToken, err = h.TM.CreateTokenPair(claims); err != nil {
		// TODO: go back and see why this is a huh?
		return ctx.JSON(http.StatusNoContent, "huh?") //nolint:stuff
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, out.AccessToken, out.RefreshToken, h.CookieDomain); err != nil {
		return auth.ErrorResponse(err)
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
		// TODO: error stuff
		return u, fmt.Errorf("error starting transaction: %v", err) //nolint:goerr113
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
