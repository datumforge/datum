package auth

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/tokens"
)

// newValidClaims returns claims with a fake subject for testing purposes ONLY
func newValidClaims(subject string) *tokens.Claims {
	iat := time.Now()
	nbf := iat
	exp := time.Now().Add(time.Hour)

	claims := &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			Issuer:    "test suite",
			IssuedAt:  jwt.NewNumericDate(iat),
			NotBefore: jwt.NewNumericDate(nbf),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		UserID:      subject,
		Email:       "rustys@datum.net",
		OrgID:       "nano_id_of_org",
		ParentOrgID: "nano_id_of_parent_org",
		Tier:        "premium",
	}

	return claims
}

// NewTestContextWithValidUser creates an echo context with a fake subject for testing purposes ONLY
func NewTestContextWithValidUser(subject string) (echo.Context, error) {
	ec := echocontext.NewTestEchoContext()

	claims := newValidClaims(subject)

	ec.Set(ContextUserClaims.name, claims)

	return ec, nil
}
