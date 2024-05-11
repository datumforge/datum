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
	// OrganizationIDs is the list of organization IDs the user is authorized to access
	OrganizationIDs []string
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

// GetAuthenticatedUserContext gets the authenticated user context
func GetAuthenticatedUserContext(c context.Context) (*AuthenticatedUser, error) {
	ec, err := echocontext.EchoContextFromContext(c)
	if err != nil {
		return nil, err
	}

	result := ec.Get(ContextAuthenticatedUser.name)

	au, ok := result.(*AuthenticatedUser)
	if !ok {
		return nil, ErrNoAuthUser
	}

	return au, nil
}

// AddAuthenticatedUserContext adds the authenticated user context and returns the context
func AddAuthenticatedUserContext(c echo.Context, user *AuthenticatedUser) context.Context {
	c.Set(ContextAuthenticatedUser.name, user)

	ctx := context.WithValue(c.Request().Context(), echocontext.EchoContextKey, c)

	c.SetRequest(c.Request().WithContext(ctx))

	return ctx
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

// getSubjectIDFromEchoContext retrieves the subject ID from the echo context
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

// getOrganizationIDFromEchoContext returns the organization ID from the echo context
func getOrganizationIDFromEchoContext(c echo.Context) (string, error) {
	if v := c.Get(ContextAuthenticatedUser.name); v != nil {
		a, ok := v.(*AuthenticatedUser)
		if !ok {
			return "", ErrNoAuthUser
		}

		oID, err := ulids.Parse(a.OrganizationID)
		if err != nil {
			return "", err
		}

		if ulids.IsZero(oID) {
			return "", ErrNoAuthUser
		}

		return a.OrganizationID, nil
	}

	return "", ErrNoAuthUser
}

// getOrganizationIDsFromEchoContext returns the list of organization IDs from the echo context
func getOrganizationIDsFromEchoContext(c echo.Context) ([]string, error) {
	if v := c.Get(ContextAuthenticatedUser.name); v != nil {
		a, ok := v.(*AuthenticatedUser)
		if !ok {
			return []string{}, ErrNoAuthUser
		}

		return a.OrganizationIDs, nil
	}

	return []string{}, ErrNoAuthUser
}

// GetOrganizationIDFromContext returns the organization ID from context
func GetOrganizationIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return getOrganizationIDFromEchoContext(ec)
}

// GetOrganizationIDsFromContext returns the list of organization IDs from context
func GetOrganizationIDsFromContext(ctx context.Context) ([]string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return []string{}, err
	}

	return getOrganizationIDsFromEchoContext(ec)
}

// GetUserIDFromContext returns the actor subject from the context
func GetUserIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return getSubjectIDFromEchoContext(ec)
}

// GetAuthTypeFromEchoContext retrieves the authentication type from the context
func GetAuthTypeFromContext(ctx context.Context) AuthenticationType {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return ""
	}

	return GetAuthTypeFromEchoContext(ec)
}

func IsAPITokenAuthentication(ctx context.Context) bool {
	return GetAuthTypeFromContext(ctx) == APITokenAuthentication
}

const (
	// UserSubjectType is the subject type for user accounts
	UserSubjectType = "user"
	// ServiceSubjectType is the subject type for service accounts
	ServiceSubjectType = "service"
)

// GetAuthzSubjectType returns the subject type based on the authentication type
func GetAuthzSubjectType(ctx context.Context) string {
	subjectType := UserSubjectType

	authType := GetAuthTypeFromContext(ctx)
	if authType == APITokenAuthentication {
		subjectType = ServiceSubjectType
	}

	return subjectType
}
