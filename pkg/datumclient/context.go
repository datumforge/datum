package datumclient

import "context"

// API-specific context keys for passing values to Datum requests. These keys are
// unexported to reduce the public interface of the API package and to ensure that there
// is no incorrect handling of the API context. Users should use the helper methods
// to manage the API context for specific requests
type contextKey uint8

// Allocate context keys to simplify context key usage in helper functions
const (
	contextKeyUnknown contextKey = iota
	contextKeyCreds
	contextKeyRequestID
)

// ContextWithToken returns a copy of the parent with the access token stored as a
// value on the new context. Passing in a context with an access token overrides the
// default credentials of the API to make per-request authenticated requests and is
// primarily used by clients that need to passthrough a user's credentials to each
// request so that the API call can be authenticated correctly
func ContextWithToken(parent context.Context, token string) context.Context {
	return context.WithValue(parent, contextKeyCreds, Token(token))
}

// CredsFromContext returns the Credentials from the provided context along with a
// boolean describing if credentials were available on the specified context
func CredsFromContext(ctx context.Context) (Credentials, bool) {
	creds, ok := ctx.Value(contextKeyCreds).(Credentials)
	return creds, ok
}

func ContextWithRequestID(parent context.Context, requestID string) context.Context {
	return context.WithValue(parent, contextKeyRequestID, requestID)
}

func RequestIDFromContext(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(contextKeyRequestID).(string)
	return requestID, ok
}

var contextKeyNames = []string{"unknown", "creds", "requestID"}

// String returns a human readable representation of the context key for easier debugging
func (c contextKey) String() string {
	if int(c) < len(contextKeyNames) {
		return contextKeyNames[c]
	}
	return contextKeyNames[0]
}
