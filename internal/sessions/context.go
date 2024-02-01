package sessions

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"golang.org/x/oauth2"
)

// SessionContextKey is the context key for the user claims
var SessionContextKey = &ContextKey{"SessionContextKey"}

// ContextKey is the key name for the additional context
type ContextKey struct {
	name string
}

// OhAuthTokenFromContext returns the Token from the ctx
func OhAuthTokenFromContext(ctx context.Context) (*oauth2.Token, error) {
	token, ok := ctx.Value(SessionContextKey).(*oauth2.Token)
	if !ok {
		return nil, errors.New("context missing Token")
	}

	return token, nil
}

// ContextWithToken returns a copy of ctx that stores the Token
func ContextWithToken(ctx context.Context, token *oauth2.Token) context.Context {
	return context.WithValue(ctx, SessionContextKey, token)
}

// UserIDFromContext returns the user ID from the ctx
func UserIDFromContext(ctx context.Context) (userID any, err error) {
	sessionDetails, ok := ctx.Value(SessionContextKey).(*Session[any])
	if !ok {
		return nil, err
	}

	userID, ok = sessionDetails.GetOk("userID")
	if !ok {
		return nil, err
	}

	return userID, nil
}

// ContextWithUserID returns a copy of ctx that stores the user ID
func ContextWithUserID(ctx context.Context, userID string) context.Context {
	if strings.TrimSpace(userID) == "" {
		return ctx
	}

	return context.WithValue(ctx, SessionContextKey, userID)
}

// SessionToken returns the session token from the context maybe, unclear if this works
func SessionToken(ctx context.Context) map[string]any {
	sd := getSessionDataFromContext(ctx)

	sd.mu.Lock()
	defer sd.mu.Unlock()

	return sd.values
}

// addSessionDataToContext adds the session details to the context
func (s *Session[P]) addSessionDataToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, SessionContextKey, s)
}

// getSessionDataFromContext gets the session information from the context
func getSessionDataFromContext(ctx context.Context) *Session[any] {
	c, ok := ctx.Value(SessionContextKey).(*Session[any])
	if !ok {
		panic("no session data in context")
	}

	return c
}
