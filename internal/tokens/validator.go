package tokens

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Validator are able to verify that access and refresh tokens were issued by
// Datum and that their claims are valid (e.g. not expired).
type Validator interface {
	// Verify an access or a refresh token after parsing and return its claims
	Verify(tks string) (claims *Claims, err error)

	// Parse an access or refresh token without verifying claims (e.g. to check an expired token)
	Parse(tks string) (claims *Claims, err error)
}

// validator implements the Validator interface, allowing structs in this package to
// embed the validation code base and supply their own keyFunc; unifying functionality
type validator struct {
	audience string
	issuer   string
	keyFunc  jwt.Keyfunc
}

// Verify an access or a refresh token after parsing and return its claims
func (v *validator) Verify(tks string) (claims *Claims, err error) {
	var token *jwt.Token

	if token, err = jwt.ParseWithClaims(tks, &Claims{}, v.keyFunc); err != nil {
		return nil, err
	}

	//TODO figure out how to use the v5 validators

	return nil, fmt.Errorf("could not parse or verify claims from %T", token.Claims) //nolint:goerr113
}

// Parse an access or refresh token verifying its signature but without verifying its
// claims. This ensures that valid JWT tokens are still accepted but claims can be
// handled on a case-by-case basis; for example by validating an expired access token
// during reauthentication
func (v *validator) Parse(tks string) (claims *Claims, err error) {
	claims = &Claims{}
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())

	if _, err = parser.ParseWithClaims(tks, claims, v.keyFunc); err != nil {
		return nil, err
	}

	return claims, nil
}
