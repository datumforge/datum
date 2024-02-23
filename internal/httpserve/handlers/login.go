package handlers

import (
	"encoding/json"
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/analytics"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/pkg/auth"
)

// LoginRequest to authenticate with the Datum Sever
type LoginRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	InviteToken string `json:"invite_token,omitempty"`
}

// LoginReply holds response to successful authentication
type LoginReply struct {
	rout.Reply
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Message      string `json:"message"`
}

// LoginHandler validates the user credentials and returns a valid cookie
// this only supports username password login today (not oauth)
func (h *Handler) LoginHandler(ctx echo.Context) error {
	user, err := h.verifyUserPassword(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// set context for remaining request based on logged in user
	userCtx := viewer.NewContext(ctx.Request().Context(), viewer.NewUserViewerFromID(user.ID, true))

	claims := createClaims(user)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// set cookies on request with the access and refresh token
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	// set sessions in response
	if err := h.SessionConfig.CreateAndStoreSession(ctx, user.ID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return err
	}

	if err := h.updateUserLastSeen(userCtx, user.ID); err != nil {
		h.Logger.Errorw("unable to update last seen", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	analytics.AssociateUser(user.ID, claims.OrgID)

	out := LoginReply{
		Reply:        rout.Reply{Success: true},
		Message:      "success",
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "access_token",
		ExpiresIn:    claims.ExpiresAt.Unix(),
	}

	return ctx.JSON(http.StatusOK, out)
}

func createClaims(u *generated.User) *tokens.Claims {
	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.ID,
		},
		UserID:      u.ID,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		AvatarURL:   *u.AvatarRemoteURL,
		OrgID:       u.Edges.Setting.DefaultOrg,
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

	// check user in the database, username == email and ensure only one record is returned
	user, err := h.getUserByEmail(ctx.Request().Context(), l.Username, enums.Credentials)
	if err != nil {
		return nil, ErrNoAuthUser
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
