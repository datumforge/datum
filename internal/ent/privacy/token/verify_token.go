package token

import "context"

type (
	VerifyToken struct {
		PrivacyToken
		VerifyToken string
	}

	verifyTokenKey struct{}
)

func (VerifyToken) GetContextKey() interface{} {
	return verifyTokenKey{}
}

// NewContextWithVerifyToken returns a new context with the verify token inside
func NewContextWithVerifyToken(parent context.Context, verifyToken string) context.Context {
	return context.WithValue(parent, verifyTokenKey{}, &VerifyToken{
		VerifyToken: verifyToken,
	})
}

// VerifyTokenFromContext parses a context for a verify token and returns the token
func VerifyTokenFromContext(ctx context.Context) *VerifyToken {
	token, _ := ctx.Value(verifyTokenKey{}).(*VerifyToken)
	return token
}
