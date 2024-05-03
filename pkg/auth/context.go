package auth

import (
	"context"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

type AuthenticationType string

const (
	// JWTAuthentication is the authentication type for JWT tokens
	JWTAuthentication AuthenticationType = "jwt"
	// PATAuthentication is the authentication type for personal access tokens
	PATAuthentication AuthenticationType = "pat"
	// APITokenAuthentication is the authentication type for API tokens
	APITokenAuthentication AuthenticationType = "api_token"
)

// ContextAuthenticatedUser is the context key for the user claims
var ContextAuthenticatedUser = &ContextKey{"authenticated_user"}

// ContextAccessToken is the context key for the access token
var ContextAccessToken = &ContextKey{"access_token"}

// ContextRequestID is the context key for the request ID
var ContextRequestID = &ContextKey{"request_id"}

// ContextKey is the key name for the additional context
type ContextKey struct {
	name string
}

// AuthenticatedUser contains the user and organization ID for the authenticated user
type AuthenticatedUser struct {
	// SubjectID is the user ID of the authenticated user or the api token ID if the user is an API token
	SubjectID string
	// OrganizationID is the organization ID of the authenticated user
	OrganizationID string
	// AuthenticationType is the type of authentication used to authenticate the user (JWT, PAT, API Token)
	AuthenticationType AuthenticationType
}

// GetContextName returns the name of the context key
func GetContextName(key *ContextKey) string {
	return key.name
}

// SetAuthenticatedUserContext sets the authenticated user context in the echo context
func SetAuthenticatedUserContext(c echo.Context, user *AuthenticatedUser) {
	c.Set(ContextAuthenticatedUser.name, user)
}

// GetAuthTypeFromEchoContext retrieves the authentication type from the echo context
func GetAuthTypeFromEchoContext(c echo.Context) AuthenticationType {
	if v := c.Get(ContextAuthenticatedUser.name); v != nil {
		a, ok := v.(*AuthenticatedUser)
		if ok {
			return a.AuthenticationType
		}
	}

	return ""
}

// GetSubjectIDFromContext retrieves the subject ID from the context
func getSubjectIDFromEchoContext(c echo.Context) (string, error) {
	if v := c.Get(ContextAuthenticatedUser.name); v != nil {
		a, ok := v.(*AuthenticatedUser)
		if !ok {
			return "", ErrNoAuthUser
		}

		uid, err := ulids.Parse(a.SubjectID)
		if err != nil {
			return "", err
		}

		if ulids.IsZero(uid) {
			return "", ErrNoAuthUser
		}

		return a.SubjectID, nil
	}

	return "", ErrNoAuthUser
}

func getOrganizationIDFromEchoContext(c echo.Context) (string, error) {
	if v := c.Get(ContextAuthenticatedUser.name); v != nil {
		a, ok := v.(*AuthenticatedUser)
		if !ok {
			return "", ErrNoAuthUser
		}

		uid, err := ulids.Parse(a.OrganizationID)
		if err != nil {
			return "", err
		}

		if ulids.IsZero(uid) {
			return "", ErrNoAuthUser
		}

		return a.OrganizationID, nil
	}

	return "", ErrNoAuthUser
}

// GetOrganizationIDFromContext returns the organization ID from context from context
func GetOrganizationIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return getOrganizationIDFromEchoContext(ec)
}

// GetUserIDFromContext returns the actor subject from the echo context
func GetUserIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return getSubjectIDFromEchoContext(ec)
}
