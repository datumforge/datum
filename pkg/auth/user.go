package auth

import (
	"context"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

// GetActorUserID returns the user from the echo.Context
func GetActorUserID(c echo.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}

	// check for null ulid
	userID := claims.ParseUserID()
	if ulids.IsZero(userID) {
		return "", ErrNoUserInfo
	}

	return claims.ParseUserID().String(), nil
}

// GetUserIDFromContext returns the actor subject from the echo context
func GetUserIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return GetActorUserID(*ec)
}

// GetOrganizationID returns the organization ID from the echo.Context
func GetOrganizationID(c echo.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}

	// check for null ulid
	orgID := claims.ParseOrgID()
	if ulids.IsZero(orgID) {
		return "", ErrNoUserInfo
	}

	return claims.ParseOrgID().String(), nil
}

// GetOrganizationIDFromContext returns the organization ID from context from context
func GetOrganizationIDFromContext(ctx context.Context) (string, error) {
	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	return GetOrganizationID(*ec)
}
