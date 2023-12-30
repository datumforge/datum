package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/tokens"
)

func (h *Handler) VerifyEmail(ctx echo.Context) error {
	reqToken := ctx.QueryParam("token")

	if err := validateVerifyRequest(reqToken); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// starts db transaction for entire request
	tx, err := h.DBClient.Tx(ctx.Request().Context())
	if err != nil {
		h.Logger.Errorw("error starting transaction", "error", err)
		return err
	}

	entUser, err := h.getUserByToken(ctx.Request().Context(), tx, reqToken)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return err
		}

		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving user token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrUnableToVerifyEmail))
	}

	// create email verification
	user := &User{
		ID:    entUser.ID,
		Email: entUser.Email,
	}

	// check to see if user is already confirmed
	if !entUser.Edges.Setting.EmailConfirmed {
		// set tokens for request
		if err := user.setUserTokens(entUser, reqToken); err != nil {
			if err := tx.Rollback(); err != nil {
				h.Logger.Errorw("error rolling back transaction", "error", err)
				return err
			}

			h.Logger.Errorw("unable to set user tokens for request", "error", err)

			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		// Construct the user token from the database fields
		token := &tokens.VerificationToken{
			Email: entUser.Email,
		}

		if token.ExpiresAt, err = user.GetVerificationExpires(); err != nil {
			if err := tx.Rollback(); err != nil {
				h.Logger.Errorw("error rolling back transaction", "error", err)
				return err
			}

			h.Logger.Errorw("unable to parse expiration", "error", err)

			return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
		}

		// Verify the token with the stored secret
		if err = token.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
			if errors.Is(err, tokens.ErrTokenExpired) {
				meowtoken, err := h.storeAndSendEmailVerificationToken(ctx.Request().Context(), tx, user)
				if err != nil {
					if err := tx.Rollback(); err != nil {
						h.Logger.Errorw("error rolling back transaction", "error", err)
						return err
					}

					h.Logger.Errorw("unable to resend verification token", "error", err)

					return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
				}

				out := &RegisterReply{
					ID:      meowtoken.ID,
					Email:   user.Email,
					Message: "Token expired, a new token has been issued. Please try again.",
					Token:   meowtoken.Token,
				}

				// commit transaction at end of request
				if err := tx.Commit(); err != nil {
					h.Logger.Errorw("error committing transaction", "error", err)

					return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
				}

				return ctx.JSON(http.StatusCreated, out)
			}

			if err := tx.Rollback(); err != nil {
				h.Logger.Errorw("error rolling back transaction", "error", err)
				return err
			}

			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		if err := h.setEmailConfirmed(ctx.Request().Context(), tx, entUser); err != nil {
			if err := tx.Rollback(); err != nil {
				h.Logger.Errorw("error rolling back transaction", "error", err)
				return err
			}

			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}

	claims := createClaims(entUser)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return err
		}

		h.Logger.Errorw("error creating token pair", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, access, refresh, h.CookieDomain); err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return err
		}

		return ErrorResponse(err)
	}

	// commit transaction at end of request
	if err := tx.Commit(); err != nil {
		h.Logger.Errorw("error committing transaction", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

// validateVerifyRequest validates the required fields are set in the user request
func validateVerifyRequest(token string) error {
	if token == "" {
		return newMissingRequiredFieldError("token")
	}

	return nil
}

// getUserByToken returns the ent user with the user settings and email verification token fields based on the
// token in the request
func (h *Handler) getUserByToken(ctx context.Context, tx *generated.Tx, token string) (*generated.User, error) {
	user, err := tx.EmailVerificationToken.Query().WithOwner().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("token", token))
	}).QueryOwner().WithSetting().WithEmailVerificationTokens().Only(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return nil, err
		}

		h.Logger.Errorw("error obtaining user from email verification token", "error", err)

		return nil, err
	}

	return user, nil
}

// setUserTokens sets the fields to verify the email
func (u *User) setUserTokens(user *generated.User, reqToken string) error {
	tokens := user.Edges.EmailVerificationTokens
	for _, t := range tokens {
		if t.Token == reqToken {
			u.EmailVerificationToken = sql.NullString{String: t.Token, Valid: true}
			u.EmailVerificationSecret = *t.Secret
			u.EmailVerificationExpires = sql.NullString{String: t.TTL.Format(time.RFC3339Nano), Valid: true}

			return nil
		}
	}

	return ErrNotFound
}

// setEmailConfirmed sets the user setting field email_confirmed to true within a transaction
func (h *Handler) setEmailConfirmed(ctx context.Context, tx *generated.Tx, user *generated.User) error {
	if _, err := tx.UserSetting.Update().SetEmailConfirmed(true).Where(func(s *sql.Selector) {
		s.Where(sql.EQ("id", user.Edges.Setting.ID))
	}).Save(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return err
		}

		return err
	}

	return nil
}
