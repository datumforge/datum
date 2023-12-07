package auth

import (
	"context"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

// GetActorSubject returns the user from the echo.Context
func GetActorSubject(c echo.Context) (string, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return "", ErrJWTMissingInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return "", ErrJWTClaimsInvalid
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", ErrSubjectNotFound
	}

	return sub, nil
}

// GetUserIDFromContext returns the actor subject from the echo context
func GetUserIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return GetActorSubject(*ec)
}
