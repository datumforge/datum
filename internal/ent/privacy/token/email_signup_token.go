package token

import (
	"context"
)

type PrivacyToken interface {
	GetContextKey() interface{}
}

type EmailSignUpToken struct {
	BaseSignUpToken
	Email string
}

type signUpTokenKey struct{}

type (
	SignUpToken interface {
		PrivacyToken
	}

	BaseSignUpToken struct {
		SignUpToken
	}
)

func (token EmailSignUpToken) GetEmail() string {
	return token.Email
}

func (BaseSignUpToken) GetContextKey() interface{} {
	return signUpTokenKey{}
}

// NewContextWithSignUpToken creates a new context with a sign-up token. It takes a
// parent context and a sign-up token as parameters and returns a new context with
// the sign-up token added
func NewContextWithSignUpToken(parent context.Context, token SignUpToken) context.Context {
	return context.WithValue(parent, signUpTokenKey{}, token)
}

// EmailSignUpTokenFromContext retrieves the value associated with the
// signUpTokenKey key from the context.
// It then type asserts the value to an EmailSignUpToken and returns it. If the
// value is not of type EmailSignUpToken, it returns nil
func EmailSignUpTokenFromContext(ctx context.Context) *EmailSignUpToken {
	token, _ := ctx.Value(signUpTokenKey{}).(*EmailSignUpToken)
	return token
}
