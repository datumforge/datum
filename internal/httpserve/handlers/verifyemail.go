package handlers

import (
	"errors"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/tokens"
)

// VerifyRequest holds the fields that should be included on a request to the `/verify` endpoint
type VerifyRequest struct {
	Token string `json:"token"`
}

// VerifyReply holds the fields that are sent on a response to the `/verify` endpoint
type VerifyReply struct {
	rout.Reply
	ID           string `json:"user_id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Message      string `json:"message,omitempty"`
}

// VerifyEmail is the handler for the email verification endpoint
func (h *Handler) VerifyEmail(ctx echo.Context) error {
	reqToken := ctx.QueryParam("token")

	if err := validateVerifyRequest(reqToken); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// setup viewer context
	ctxWithToken := token.NewContextWithVerifyToken(ctx.Request().Context(), reqToken)

	entUser, err := h.getUserByEVToken(ctxWithToken, reqToken)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving user token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrUnableToVerifyEmail))
	}

	// create email verification
	user := &User{
		ID:    entUser.ID,
		Email: entUser.Email,
	}

	viewerCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(entUser.ID, true))

	// check to see if user is already confirmed
	if !entUser.Edges.Setting.EmailConfirmed {
		// set tokens for request
		if err := user.setUserTokens(entUser, reqToken); err != nil {
			h.Logger.Errorw("unable to set user tokens for request", "error", err)

			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
		}

		// Construct the user token from the database fields
		t := &tokens.VerificationToken{
			Email: entUser.Email,
		}

		if t.ExpiresAt, err = user.GetVerificationExpires(); err != nil {
			h.Logger.Errorw("unable to parse expiration", "error", err)

			return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
		}

		// Verify the token with the stored secret
		if err = t.Verify(user.GetVerificationToken(), user.EmailVerificationSecret); err != nil {
			if errors.Is(err, tokens.ErrTokenExpired) {
				viewerCtx = token.NewContextWithSignUpToken(viewerCtx, user.Email)

				meowtoken, err := h.storeAndSendEmailVerificationToken(viewerCtx, user)
				if err != nil {
					h.Logger.Errorw("unable to resend verification token", "error", err)

					return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
				}

				out := &VerifyReply{
					Reply:   rout.Reply{Success: false},
					ID:      meowtoken.ID,
					Email:   user.Email,
					Message: "Token expired, a new token has been issued. Please try again.",
					Token:   meowtoken.Token,
				}

				return ctx.JSON(http.StatusCreated, out)
			}

			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
		}

		if err := h.setEmailConfirmed(viewerCtx, entUser); err != nil {
			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
		}
	}

	claims := createClaims(entUser)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		h.Logger.Errorw("error creating token pair", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// set cookies on request with the access and refresh token
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	out := &VerifyReply{
		Reply:        rout.Reply{Success: true},
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "access_token",
		ExpiresIn:    claims.ExpiresAt.Unix(),
	}

	return ctx.JSON(http.StatusOK, out)
}

// validateVerifyRequest validates the required fields are set in the user request
func validateVerifyRequest(token string) error {
	if token == "" {
		return rout.NewMissingRequiredFieldError("token")
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
