package auth

import (
	echo "github.com/datumforge/echox"
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

const (
	JWTAuthentication = "jwt"
	PATAuthentication = "pat"
)

// SessionSkipperFunc is the function that determines if the session check should be skipped
// due to the request being a PAT auth request
var SessionSkipperFunc = func(c echo.Context) bool {
	return c.Get(ContextAuthType.name) == PATAuthentication
}
