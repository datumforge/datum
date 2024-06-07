package handlers

import (
	"context"
	"fmt"
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
)

// authSession is a struct that holds the access, refresh tokens and session for a user
type authSession struct {
	accessToken  string
	refreshToken string
	session      string
}

const bearerToken = "Bearer"

// createClaims creates the claims for the JWT token using the mapping ids for the user
// and the user's default organization id
func createClaims(u *generated.User) *tokens.Claims {
	return createClaimsWithOrg(u, "")
}

// createClaims creates the claims for the JWT token using the mapping ids for the user and organization
func createClaimsWithOrg(u *generated.User, targetOrgMappingID string) *tokens.Claims {
	if targetOrgMappingID == "" {
		if u.Edges.Setting.Edges.DefaultOrg != nil {
			targetOrgMappingID = u.Edges.Setting.Edges.DefaultOrg.MappingID
		}
	}

	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.MappingID,
		},
		UserID: u.MappingID,
		OrgID:  targetOrgMappingID,
	}
}

func (h *Handler) generateUserAuthSession(ctx echo.Context, user *generated.User) (*authSession, error) {
	return h.generateUserAuthSessionWithOrg(ctx, user, "")
}

// generateNewAuthSession creates a new auth session for the user
func (h *Handler) generateUserAuthSessionWithOrg(ctx echo.Context, user *generated.User, targetOrgMappingID string) (*authSession, error) {
	auth, err := h.createTokenPair(user, targetOrgMappingID)
	if err != nil {
		return nil, err
	}

	auth.session, err = h.generateUserSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

// generateNewAuthSession creates a new auth session for the user
func (h *Handler) generateOauthAuthSession(ctx context.Context, w http.ResponseWriter, user *generated.User, oauthRequest models.OauthTokenRequest) (*authSession, error) {
	auth, err := h.createTokenPair(user, "")
	if err != nil {
		return nil, err
	}

	auth.session, err = h.generateOauthUserSession(ctx, w, user.ID, oauthRequest)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (h *Handler) createTokenPair(user *generated.User, targetOrgMappingID string) (*authSession, error) {
	// create new claims for the user
	newClaims := createClaimsWithOrg(user, targetOrgMappingID)

	// create a new token pair for the user
	access, refresh, err := h.TM.CreateTokenPair(newClaims)
	if err != nil {
		return nil, err
	}

	return &authSession{
		accessToken:  access,
		refreshToken: refresh,
	}, nil
}

func (h *Handler) generateUserSession(ctx echo.Context, userID string) (string, error) {
	// set sessions in response
	if err := h.SessionConfig.CreateAndStoreSession(ctx, userID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return "", h.InternalServerError(ctx, err)
	}

	// return the session value for the UI to use
	session, err := sessions.SessionToken(ctx.Request().Context())
	if err != nil {
		return "", h.InternalServerError(ctx, err)
	}

	return session, nil
}

func (h *Handler) generateOauthUserSession(ctx context.Context, w http.ResponseWriter, userID string, oauthRequest models.OauthTokenRequest) (string, error) {
	setSessionMap := map[string]any{}
	setSessionMap[sessions.ExternalUserIDKey] = fmt.Sprintf("%v", oauthRequest.ExternalUserID)
	setSessionMap[sessions.UsernameKey] = oauthRequest.ExternalUserName
	setSessionMap[sessions.UserTypeKey] = oauthRequest.AuthProvider
	setSessionMap[sessions.EmailKey] = oauthRequest.Email
	setSessionMap[sessions.UserIDKey] = user.ID

	c, err := h.SessionConfig.SaveAndStoreSession(ctx, w, setSessionMap, userID)
	if err != nil {
		return "", err
	}

	session, err := sessions.SessionToken(c)
	if err != nil {
		return "", err
	}

	return session, nil
}
