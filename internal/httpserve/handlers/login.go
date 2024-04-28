package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/passwd"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
)

// LoginRequest to authenticate with the Datum Sever
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OTPCode  string `json:"otp_code,omitempty"`
}

// LoginReply holds response to successful authentication
type LoginReply struct {
	rout.Reply
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Session      string `json:"session,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Message      string `json:"message"`
}

// LoginHandler validates the user credentials and returns a valid cookie
// this handler only supports username password login
func (h *Handler) LoginHandler(ctx echo.Context) error {
	user, err := h.verifyUserPassword(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// set context for remaining request based on logged in user
	userCtx := viewer.NewContext(ctx.Request().Context(), viewer.NewUserViewerFromID(user.ID, true))

	if err := h.addDefaultOrgToUserQuery(userCtx, user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

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

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	if err := h.updateUserLastSeen(userCtx, user.ID); err != nil {
		h.Logger.Errorw("unable to update last seen", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	props := ph.NewProperties().
		Set("user_id", user.ID).
		Set("email", user.Email).
		Set("organization_id", claims.OrgID).
		Set("auth_provider", user.AuthProvider)

	h.AnalyticsClient.Event("user_authenticated", props)
	h.AnalyticsClient.UserProperties(user.ID, props)

	// return the session value for the UI to use
	// the UI will need to set the cookie because authentication is handled
	// server side
	s, err := sessions.SessionToken(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	out := LoginReply{
		Reply:        rout.Reply{Success: true},
		Message:      "success",
		AccessToken:  access,
		RefreshToken: refresh,
		Session:      s,
		TokenType:    "access_token",
		ExpiresIn:    claims.ExpiresAt.Unix(),
	}

	return ctx.JSON(http.StatusOK, out)
}

func createClaims(u *generated.User) *tokens.Claims {
	orgID := ""
	if u.Edges.Setting.Edges.DefaultOrg != nil {
		orgID = u.Edges.Setting.Edges.DefaultOrg.ID
	}

	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.ID,
		},
		UserID:      u.ID,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		AvatarURL:   *u.AvatarRemoteURL,
		OrgID:       orgID,
	}
}

// verifyUserPassword verifies the username and password are valid
func (h *Handler) verifyUserPassword(ctx echo.Context) (*generated.User, error) {
	var l LoginRequest
	if err := ctx.Bind(&l); err != nil {
		return nil, ErrBadRequest
	}

	if l.Username == "" || l.Password == "" {
		return nil, ErrMissingRequiredFields
	}

	// check user in the database, username == email and ensure only one record is returned
	user, err := h.getUserByEmail(ctx.Request().Context(), l.Username, enums.AuthProviderCredentials)
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
