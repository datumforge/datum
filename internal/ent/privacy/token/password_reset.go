package token

import "context"

type (
	PasswordResetToken struct {
		PrivacyToken
		ResetToken string
	}

	PasswordResetTokenKey struct{}
)

func (PasswordResetToken) GetContextKey() interface{} {
	return &PasswordResetTokenKey{}
}

// NewContextWithPasswordResetToken returns a new context with the password reset token inside
func NewContextWithPasswordResetToken(parent context.Context, token string) context.Context {
	return context.WithValue(parent, &PasswordResetTokenKey{}, &PasswordResetToken{ResetToken: token})
}

// PasswordResetTokenFromContext parses the context for a password reset token and returns it
func PasswordResetTokenFromContext(ctx context.Context) *PasswordResetToken {
	token, _ := ctx.Value(&PasswordResetTokenKey{}).(*PasswordResetToken)

	return token
}
