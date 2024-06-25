package authsessions

import (
	"context"
	"fmt"
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/httpsling"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
)

type Auth struct {
	tokenManager  *tokens.TokenManager
	sessionConfig *sessions.SessionConfig
}

func NewAuthManager() *Auth {
	return &Auth{}
}

// SetSessionConfig sets the session config for the auth session
func (a *Auth) SetSessionConfig(sc *sessions.SessionConfig) {
	if a == nil {
		a = &Auth{}
	}

	a.sessionConfig = sc
}

func (a *Auth) GetSessionConfig() *sessions.SessionConfig {
	return a.sessionConfig
}

// SetTokenManager sets the token manager for the auth session
func (a *Auth) SetTokenManager(tm *tokens.TokenManager) {
	if a == nil {
		a = &Auth{}
	}

	a.tokenManager = tm
}

func (a *Auth) GetTokenManager() *tokens.TokenManager {
	return a.tokenManager
}

// createClaims creates the claims for the JWT token using the id for the user and organization
func createClaimsWithOrg(u *generated.User, targetOrgID string) *tokens.Claims {
	if targetOrgID == "" {
		if u.Edges.Setting.Edges.DefaultOrg != nil {
			targetOrgID = u.Edges.Setting.Edges.DefaultOrg.ID
		}
	}

	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.ID,
		},
		UserID: u.ID,
		OrgID:  targetOrgID,
	}
}

// generateNewAuthSession creates a new auth session for the user and their default organization id
func (a *Auth) GenerateUserAuthSession(ctx echo.Context, user *generated.User) (*models.AuthData, error) {
	return a.GenerateUserAuthSessionWithOrg(ctx, user, "")
}

// generateUserAuthSessionWithOrg creates a new auth session for the user and the new target organization id
func (a *Auth) GenerateUserAuthSessionWithOrg(ctx echo.Context, user *generated.User, targetOrgID string) (*models.AuthData, error) {
	auth, err := a.createTokenPair(user, targetOrgID)
	if err != nil {
		return nil, err
	}

	auth.Session, err = a.generateUserSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	auth.TokenType = string(httpsling.BearerAuthType)

	return auth, nil
}

// GenerateOauthAuthSession creates a new auth session for the oauth user and their default organization id
func (a *Auth) GenerateOauthAuthSession(ctx context.Context, w http.ResponseWriter, user *generated.User, oauthRequest models.OauthTokenRequest) (*models.AuthData, error) {
	auth, err := a.createTokenPair(user, "")
	if err != nil {
		return nil, err
	}

	auth.Session, err = a.generateOauthUserSession(ctx, w, user.ID, oauthRequest)
	if err != nil {
		return nil, err
	}

	auth.TokenType = string(httpsling.BearerAuthType)

	return auth, nil
}

// createTokenPair creates a new token pair for the user and the target organization id (or default org if none provided)
func (a *Auth) createTokenPair(user *generated.User, targetOrgID string) (*models.AuthData, error) {
	// create new claims for the user
	newClaims := createClaimsWithOrg(user, targetOrgID)

	// create a new token pair for the user
	access, refresh, err := a.tokenManager.CreateTokenPair(newClaims)
	if err != nil {
		return nil, err
	}

	return &models.AuthData{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// GenerateUserSession creates a new session for the user and stores it in the response
func (a *Auth) generateUserSession(ctx echo.Context, userID string) (string, error) {
	// set sessions in response
	if err := a.sessionConfig.CreateAndStoreSession(ctx, userID); err != nil {
		return "", err
	}

	// return the session value for the UI to use
	session, err := sessions.SessionToken(ctx.Request().Context())
	if err != nil {
		return "", err
	}

	return session, nil
}

// generateOauthUserSession creates a new session for the oauth user and stores it in the response
func (a *Auth) generateOauthUserSession(ctx context.Context, w http.ResponseWriter, userID string, oauthRequest models.OauthTokenRequest) (string, error) {
	setSessionMap := map[string]any{}
	setSessionMap[sessions.ExternalUserIDKey] = fmt.Sprintf("%v", oauthRequest.ExternalUserID)
	setSessionMap[sessions.UsernameKey] = oauthRequest.ExternalUserName
	setSessionMap[sessions.UserTypeKey] = oauthRequest.AuthProvider
	setSessionMap[sessions.EmailKey] = oauthRequest.Email
	setSessionMap[sessions.UserIDKey] = userID

	c, err := a.sessionConfig.SaveAndStoreSession(ctx, w, setSessionMap, userID)
	if err != nil {
		return "", err
	}

	session, err := sessions.SessionToken(c)
	if err != nil {
		return "", err
	}

	return session, nil
}
