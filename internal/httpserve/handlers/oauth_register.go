package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	echo "github.com/datumforge/echox"
	ph "github.com/posthog/posthog-go"
	"golang.org/x/oauth2"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/providers/github"
	"github.com/datumforge/datum/pkg/providers/google"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/sessions"
)

// OauthTokenRequest to authenticate an oauth user with the Datum Server
type OauthTokenRequest struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	AuthProvider     string `json:"authProvider"`
	ExternalUserID   string `json:"externalUserId"`
	ExternalUserName string `json:"externalUserName"`
	ClientToken      string `json:"clientToken"`
}

// OauthRegister returns the TokenResponse for a verified authenticated external oauth user
func (h *Handler) OauthRegister(ctx echo.Context) error {
	var r OauthTokenRequest
	if err := ctx.Bind(&r); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	ctxWithToken := token.NewContextWithOauthTooToken(ctx.Request().Context(), r.Email)

	// create oauth2 token from request input
	tok := &oauth2.Token{
		AccessToken: r.ClientToken,
	}

	// verify the token provided to ensure the user is valid
	if err := h.verifyClientToken(ctxWithToken, r.AuthProvider, tok, r.Email); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// check if users exists and create if not, updates last seen of existing user
	user, err := h.CheckAndCreateUser(ctxWithToken, r.Name, r.Email, enums.AuthProvider(strings.ToUpper(r.AuthProvider)))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// set context for remaining request based on logged in user
	userCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(user.ID, true))

	if err := h.addDefaultOrgToUserQuery(userCtx, user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// create claims for verified user
	claims := createClaims(user)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// set cookies for the user
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	setSessionMap := map[string]any{}
	setSessionMap[sessions.ExternalUserIDKey] = fmt.Sprintf("%v", r.ExternalUserID)
	setSessionMap[sessions.UsernameKey] = r.ExternalUserName
	setSessionMap[sessions.UserTypeKey] = r.AuthProvider
	setSessionMap[sessions.EmailKey] = r.Email
	setSessionMap[sessions.UserIDKey] = user.ID

	c, err := h.SessionConfig.SaveAndStoreSession(ctx.Request().Context(), ctx.Response().Writer, setSessionMap, user.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	ctx.SetRequest(ctx.Request().WithContext(c))

	props := ph.NewProperties().
		Set("user_id", user.ID).
		Set("email", user.Email).
		Set("organization_id", claims.OrgID).
		Set("auth_provider", r.AuthProvider)

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
		Message:      "success",
		AccessToken:  access,
		RefreshToken: refresh,
		Session:      s,
		TokenType:    "Bearer",
		ExpiresIn:    claims.ExpiresAt.Unix(),
	}

	// Return the access token
	return ctx.JSON(http.StatusOK, out)
}

// verifyClientToken verifies the provided access token from an external oauth2 provider is valid and matches the user's email
// supported providers are Github and Google
func (h *Handler) verifyClientToken(ctx context.Context, provider string, token *oauth2.Token, email string) error {
	switch strings.ToLower(provider) {
	case githubProvider:
		config := h.getGithubOauth2Config()
		cc := github.ClientConfig{IsEnterprise: false, IsMock: h.IsTest}

		return github.VerifyClientToken(ctx, token, config, email, &cc)
	case googleProvider:
		config := h.getGoogleOauth2Config()
		return google.VerifyClientToken(ctx, token, config, email)
	default:
		return ErrInvalidProvider
	}
}
