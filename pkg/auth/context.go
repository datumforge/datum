package auth

const (
	// JWTAuthentication is the authentication type for JWT tokens
	JWTAuthentication = "jwt"
	// PATAuthentication is the authentication type for personal access tokens
	PATAuthentication = "pat"
	// APITokenAuthentication is the authentication type for API tokens
	APITokenAuthentication = "api_token"
)

// ContextUserClaims is the context key for the user claims
var ContextUserClaims = &ContextKey{"user_claims"}

// ContextAccessToken is the context key for the access token
var ContextAccessToken = &ContextKey{"access_token"}

// ContextRequestID is the context key for the request ID
var ContextRequestID = &ContextKey{"request_id"}

// ContextAuthType is the context key for the authentication type
var ContextAuthType = &ContextKey{"auth_type"}

// ContextKey is the key name for the additional context
type ContextKey struct {
	name string
}

// GetContextName returns the name of the context key
func GetContextName(key *ContextKey) string {
	return key.name
}
