package token

import (
	"context"
)

// WebauthnToken that implements the PrivacyToken interface
type WebauthnToken struct {
	PrivacyToken
	email string
}

type webauthnTokenKey struct{}

// NewWebauthnWithEmail creates a new PrivacyToken of type WebauthnToken with
// email set
func NewWebauthnWithEmail(email string) WebauthnToken {
	return WebauthnToken{
		email: email,
	}
}

// GetEmail from webauthn token
func (token *WebauthnToken) GetEmail() string {
	return token.email
}

// SetEmail on the webauthn token
func (token *WebauthnToken) SetEmail(email string) {
	token.email = email
}

// GetContextKey from WebauthnToken
func (WebauthnToken) GetContextKey() interface{} {
	return webauthnTokenKey{}
}

// NewContextWithWebauthnToken creates a new context with a webauthn token. It takes a
// parent context and a webauthn token as parameters and returns a new context with
// the webauthn token added
func NewContextWithWebauthnToken(parent context.Context, email string) context.Context {
	return context.WithValue(parent, webauthnTokenKey{}, &WebauthnToken{
		email: email,
	})
}

// WebauthnTokenFromContext retrieves the value associated with the
// webauthnTokenKey key from the context.
// It then type asserts the value to an WebauthnToken and returns it. If the
// value is not of type WebauthnToken, it returns nil
func WebauthnTokenFromContext(ctx context.Context) *WebauthnToken {
	token, _ := ctx.Value(webauthnTokenKey{}).(*WebauthnToken)
	return token
}
