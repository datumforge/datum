package token

import (
	"context"
)

type PrivacyToken interface {
	GetContextKey() interface{}
}

type EmailSignupToken struct {
	BaseSignupToken
	Email string
}

type signupTokenKey struct{}

type (
	SignupToken interface {
		PrivacyToken
	}

	BaseSignupToken struct {
		SignupToken
	}
)

func (token EmailSignupToken) GetEmail() string {
	return token.Email
}

func (BaseSignupToken) GetContextKey() interface{} {
	return signupTokenKey{}
}

// NewContextWithSignupToken createsg a new context with a signup token. It takes a
// parent context and a signup token as parameters and returns a new context with
// the signup token added
func NewContextWithSignupToken(parent context.Context, token SignupToken) context.Context {
	return context.WithValue(parent, signupTokenKey{}, token)
}

// EmailSignupTokenFromContext retrieves the value associated with the
// signupTokenKey key from the context.
// It then type asserts the value to an EmailSignupToken and returns it. If the
// value is not of type EmailSignupToken`, it returns nil
func EmailSignupTokenFromContext(ctx context.Context) *EmailSignupToken {
	token, _ := ctx.Value(signupTokenKey{}).(*EmailSignupToken)
	return token
}
