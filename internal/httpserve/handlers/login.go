package handlers

import (
	"encoding/json"
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/tokens"
)

// LoginRequest to authenticate with the Datum Sever
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler validates the user credentials and returns a valid cookie
// this only supports username password login today (not oauth)
func (h *Handler) LoginHandler(ctx echo.Context) error {
	user, err := h.verifyUserPassword(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	claims := createClaims(user)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, access, refresh, h.CookieDomain); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	if err := h.startTransaction(ctx.Request().Context()); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrProcessingRequest)
	}

	if err := h.updateUserLastSeen(ctx.Request().Context(), user.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	if err = h.TXClient.Commit(); err != nil {
		h.Logger.Errorw(transactionCommitErr, "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, Response{Message: "success"})
}

func createClaims(u *generated.User) *tokens.Claims {
	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.ID,
		},
		UserID: u.ID,
		Email:  u.Email,
	}
}

// verifyUserPassword verifies the username and password are valid
func (h *Handler) verifyUserPassword(ctx echo.Context) (*generated.User, error) {
	var l LoginRequest

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&l); err != nil {
		return nil, ErrBadRequest
	}

	if l.Username == "" || l.Password == "" {
		return nil, ErrMissingRequiredFields
	}

	if err := h.startTransaction(ctx.Request().Context()); err != nil {
		return nil, err
	}

	// check user in the database, username == email and ensure only one record is returned
	user, err := h.getUserByEmail(ctx.Request().Context(), l.Username)
	if err != nil {
		return nil, ErrNoAuthUser
	}

	if err = h.TXClient.Commit(); err != nil {
		return nil, err
	}

	if user.Edges.Setting.Status != "ACTIVE" {
		return nil, ErrNoAuthUser
	}

	// verify the password is correct
	valid, err := passwd.VerifyDerivedKey(*user.Password, l.Password)
	if err != nil || !valid {
		return nil, ErrInvalidCredentials
	}

	// verify email is verified
	if !user.Edges.Setting.EmailConfirmed {
		return nil, ErrUnverifiedUser
	}

	return user, nil
}
