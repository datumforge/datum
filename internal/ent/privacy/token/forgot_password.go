package token

import "context"

type (
	ForgotPasswordToken struct {
		PrivacyToken
		Email string
	}

	ForgotPasswordTokenKey struct{}
)

func (ForgotPasswordToken) GetContextKey() interface{} {
	return &ForgotPasswordTokenKey{}
}

// NewContextWithForgotPasswordToken returns a new context with the email for forgot password inside
func NewContextWithForgotPasswordToken(parent context.Context, email string) context.Context {
	return context.WithValue(parent, &ForgotPasswordTokenKey{}, &ForgotPasswordToken{Email: email})
}

// ForgotPasswordTokenFromContext parses the context for a forgot password and returns it
func ForgotPasswordTokenFromContext(ctx context.Context) *ForgotPasswordToken {
	token, _ := ctx.Value(&ForgotPasswordTokenKey{}).(*ForgotPasswordToken)

	return token
}
