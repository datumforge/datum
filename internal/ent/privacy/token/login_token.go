package token

import (
	"context"
)

type LoginToken struct {
	PrivacyToken
	Email string
}

type LoginTokenKey struct{}

func (LoginToken) GetContextKey() interface{} {
	return LoginTokenKey{}
}

// NewContextWithLoginToken takes the parent context and returns a new context with the login token present
func NewContextWithLoginToken(parent context.Context, email string) context.Context {
	return context.WithValue(parent, LoginTokenKey{}, &LoginToken{Email: email})
}

// LoginTokenFromContext parses the context and returns the token
func LoginTokenFromContext(ctx context.Context) *LoginToken {
	token, _ := ctx.Value(LoginTokenKey{}).(*LoginToken)
	return token
}

// GetLoginTokenValidatorFunc can be used to validate a PrivacyToken
func GetLoginTokenValidatorFunc(email string) func(PrivacyToken) error {
	return func(t PrivacyToken) error {
		actualToken, ok := t.(*LoginToken)
		if !ok {
			return ErrInvalidTokenType
		}

		if actualToken.Email == email {
			return nil
		}

		return ErrIncorrectEmail
	}
}
