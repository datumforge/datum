package auth

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/tokens"
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
func // `NewTestContextWithValidUser` is a function that creates an Echo context for testing purposes.
// It sets a fake subject and other claims in the context to simulate a valid user during testing.
// The function returns the created Echo context.
// `NewTestContextWithValidUser` is a function that creates an Echo context for testing purposes.
// It sets up a fake user with specific claims (such as user ID, email, organization ID, etc.) for
// the purpose of simulating a user session during testing. This function is used to generate a
// test context with predefined user information.
NewTestContextWithValidUser(subject string) (echo.Context, error) {
	ec := echocontext.NewTestEchoContext()

	claims := newValidClaims(subject)

	ec.Set(ContextUserClaims.name, claims)

	return ec, nil
}

// newValidClaims returns claims with a fake orgID for testing purposes ONLY
func newValidClaimsOrgID(sub, orgID string) *tokens.Claims {
	iat := time.Now()
	nbf := iat
	exp := time.Now().Add(time.Hour)

	claims := &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   sub,
			Issuer:    "test suite",
			IssuedAt:  jwt.NewNumericDate(iat),
			NotBefore: jwt.NewNumericDate(nbf),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		UserID:      sub,
		Email:       "rustys@datum.net",
		OrgID:       orgID,
		ParentOrgID: "01HWRCWA74GJ7F4092CRNXT7MD",
		Tier:        "premium",
	}

	return claims
}

// NewTestContextWithOrgID creates an echo context with a fake orgID for testing purposes ONLY
func NewTestContextWithOrgID(sub, orgID string) (echo.Context, error) {
	ec := echocontext.NewTestEchoContext()

	claims := newValidClaimsOrgID(sub, orgID)

	ec.Set(ContextUserClaims.name, claims)

	return ec, nil
}
