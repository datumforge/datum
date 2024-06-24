package authsessions

import (
	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/httpsling"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
)

type Auth struct {
	TokenManager  *tokens.TokenManager
	SessionConfig *sessions.SessionConfig
	Logger        *zap.SugaredLogger
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

// createTokenPair creates a new token pair for the user and the target organization id (or default org if none provided)
func (a *Auth) createTokenPair(user *generated.User, targetOrgID string) (*models.AuthData, error) {
	// create new claims for the user
	newClaims := createClaimsWithOrg(user, targetOrgID)

	// create a new token pair for the user
	access, refresh, err := a.TokenManager.CreateTokenPair(newClaims)
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
	if err := a.SessionConfig.CreateAndStoreSession(ctx, userID); err != nil {
		a.Logger.Errorw("unable to save session", "error", err)

		return "", err
	}

	// return the session value for the UI to use
	session, err := sessions.SessionToken(ctx.Request().Context())
	if err != nil {
		return "", err
	}

	return session, nil
}
