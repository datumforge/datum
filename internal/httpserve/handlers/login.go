package handlers

import (
	"encoding/json"
	"net/http"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/tokens"
)

type User struct {
	Username string
	Password string
	userID   string
}

// LoginHandler validates the user credentials and returns a valid cookie
// this only supports username password login today (not oauth)
func (h *Handler) LoginHandler(ctx echo.Context) error {
	user, err := h.verifyCredentials(ctx)
	if err != nil {
		return err
	}

	claims := createClaims(user)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return auth.ErrorResponse(err)
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, access, refresh, h.CookieDomain); err != nil {
		return auth.ErrorResponse(err)
	}

	return ctx.JSON(http.StatusOK, "success")
}

func createClaims(u *User) *tokens.Claims {
	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.userID,
		},
		UserID: u.userID,
		Email:  u.Username,
	}
}

// verifyCredentials verifies the username and password are valid
func (h *Handler) verifyCredentials(ctx echo.Context) (*User, error) {
	var u User

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&u); err != nil {
		return nil, auth.ErrorResponse(err)
	}

	// check user in the database, username == email and ensure only one record is returned
	user, err := h.DBClient.User.Query().WithSetting().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("email", u.Username))
	}).Only(ctx.Request().Context())
	if err != nil {
		return nil, auth.ErrorResponse(err)
	}

	// verify the password is correct
	valid, err := passwd.VerifyDerivedKey(*user.Password, u.Password)
	if err != nil || !valid {
		return nil, auth.Unauthorized(ctx)
	}

	// verify email is verified
	if !user.Edges.Setting.EmailConfirmed {
		return nil, auth.Unverified(ctx)
	}

	u.userID = user.ID

	return &u, nil
}
