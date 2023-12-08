package auth

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/tokens"
)

// newValidSignedJWTWithClaims creates a jwt with a fake subject for testing purposes ONLY
func newValidSignedJWTWithClaims(subject string) (*tokens.Token, error) {
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
		OrgID:       "nano_id_of_org",
		ParentOrgID: "nano_id_of_parent_org",
		Tier:        "premium",
	}

	t := &tokens.Token{
		Claims: *claims,
	}

	return t, nil
}

// NewTestContextWithValidUser creates an echo context with a fake subject for testing purposes ONLY
func NewTestContextWithValidUser(subject string) (*echo.Context, error) {
	ec := echocontext.NewTestEchoContext()

	j, err := newValidSignedJWTWithClaims(subject)
	if err != nil {
		return nil, err
	}

	ec.Set(ContextUserClaims, j)

	return &ec, nil
}
