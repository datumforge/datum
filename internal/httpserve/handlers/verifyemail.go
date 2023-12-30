package handlers

import (
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
	if err := h.startTransaction(ctx.Request().Context()); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrProcessingRequest)
	}

	entUser, err := h.getUserByEVToken(ctx.Request().Context(), reqToken)
	if err != nil {
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
			if err := h.TXClient.Rollback(); err != nil {
				h.Logger.Errorw(rollbackErr, "error", err)
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
			if err := h.TXClient.Rollback(); err != nil {
				h.Logger.Errorw(rollbackErr, "error", err)
				return err
			}

			h.Logger.Errorw("unable to parse expiration", "error", err)

			return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
		}

		// Verify the token with the stored secret
		if err = token.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
			if errors.Is(err, tokens.ErrTokenExpired) {
				meowtoken, err := h.storeAndSendEmailVerificationToken(ctx.Request().Context(), user)
				if err != nil {
					if err := h.TXClient.Rollback(); err != nil {
						h.Logger.Errorw(rollbackErr, "error", err)
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
				if err := h.TXClient.Commit(); err != nil {
					h.Logger.Errorw("error committing transaction", "error", err)

					return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
				}

				return ctx.JSON(http.StatusCreated, out)
			}

			if err := h.TXClient.Rollback(); err != nil {
				h.Logger.Errorw(rollbackErr, "error", err)
				return err
			}

			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		if err := h.setEmailConfirmed(ctx.Request().Context(), entUser); err != nil {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}

	claims := createClaims(entUser)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		if err := h.TXClient.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return err
		}

		h.Logger.Errorw("error creating token pair", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, access, refresh, h.CookieDomain); err != nil {
		if err := h.TXClient.Rollback(); err != nil {
			h.Logger.Errorw(rollbackErr, "error", err)
			return err
		}

		return ErrorResponse(err)
	}

	// commit transaction at end of request
	if err := h.TXClient.Commit(); err != nil {
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
