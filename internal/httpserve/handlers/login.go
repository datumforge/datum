package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
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
		return err
	}

	claims := createClaims(user)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return ErrorResponse(err)
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, access, refresh, h.CookieDomain); err != nil {
		return ErrorResponse(err)
	}

	tx, err := h.DBClient.Tx(ctx.Request().Context())
	if err != nil {
		h.Logger.Errorw("error starting transaction", "error", err)
		return ctx.JSON(http.StatusInternalServerError, ErrProcessingRequest)
	}

	if _, err := tx.User.Update().SetLastSeen(time.Now()).Where(func(s *sql.Selector) {
		s.Where(sql.EQ("id", user.ID))
	}).Save(ctx.Request().Context()); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	if err = tx.Commit(); err != nil {
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
		auth.Unauthorized(ctx) //nolint:errcheck
		return nil, ErrBadRequest
	}

	if l.Username == "" || l.Password == "" {
		auth.Unauthorized(ctx) //nolint:errcheck
		return nil, ErrMissingRequiredFields
	}

	// check user in the database, username == email and ensure only one record is returned
	user, err := h.DBClient.User.Query().WithSetting().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("email", l.Username))
	}).Only(ctx.Request().Context())
	if err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return nil, ErrNoAuthUser
	}

	if user.Edges.Setting.Status != "ACTIVE" {
		auth.Unauthorized(ctx) //nolint:errcheck
		return nil, ErrNoAuthUser
	}

	// verify the password is correct
	valid, err := passwd.VerifyDerivedKey(*user.Password, l.Password)
	if err != nil || !valid {
		auth.Unauthorized(ctx) //nolint:errcheck
		return nil, ErrInvalidCredentials
	}

	// verify email is verified
	if !user.Edges.Setting.EmailConfirmed {
		auth.Unverified(ctx) //nolint:errcheck
		return nil, ErrUnverifiedUser
	}

	return user, nil
}
