package token

import "context"

type (
	RefreshToken struct {
		PrivacyToken
		loginSessionID string
	}

	refreshTokenKey struct{}
)

func (RefreshToken) GetContextKey() interface{} {
	return refreshTokenKey{}
}

// NewContextWithRefreshToken returns a new context with the refresh token inside
func NewContextWithRefreshToken(parent context.Context, loginSessionID string) context.Context {
	return context.WithValue(parent, refreshTokenKey{}, &RefreshToken{
		loginSessionID: loginSessionID,
	})
}

// RefreshTokenFromContext parses a context for a refresh token and returns the token
func RefreshTokenFromContext(ctx context.Context) *RefreshToken {
	token, _ := ctx.Value(refreshTokenKey{}).(*RefreshToken)
	return token
}
