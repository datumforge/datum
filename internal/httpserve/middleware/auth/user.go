package auth

import (
	"context"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

// GetActorSubject returns the user from the echo.Context
func GetActorSubject(c echo.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}

	return claims.ParseUserID().String(), nil
}

// GetUserIDFromContext returns the actor subject from the echo context
func GetUserIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return GetActorSubject(*ec)
}
