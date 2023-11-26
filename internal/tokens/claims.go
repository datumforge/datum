package tokens

import (
	"crypto/subtle"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/utils/ulids"
)

// Claims implements custom claims and extends the `jwt.RegisteredClaims` struct; we will store user-related elements here (and thus in the JWT Token) for reference / validation
type Claims struct {

	// UserID is the internal generated ID for the user
	UserID string `json:"user_id,omitempty"`
	// the email associated with the user
	Email string `json:"email,omitempty"`
	// the organization ID the JWT token is valid for
	OrgID string `json:"org,omitempty"`
	// the ID of the parent organization, if a child
	ParentOrgID string `json:"project,omitempty"`
	// the entitlement tier the token is valid for
	Tier string `json:"tier,omitempty"`
	jwt.RegisteredClaims
}

// RegisteredClaims are a structured version of the JWT Claims Set restricted to Registered Claim Names
// https://datatracker.ietf.org/doc/html/rfc7519#section-4.1
type RegisteredClaims struct {
	// the `iss` (Issuer) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1
	Issuer string `json:"iss,omitempty"`

	// the `sub` (Subject) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2
	Subject string `json:"sub,omitempty"`

	// the `aud` (Audience) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.3
	Audience ClaimStrings `json:"aud,omitempty"`

	// the `exp` (Expiration Time) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
	ExpiresAt *NumericDate `json:"exp,omitempty"`

	// the `nbf` (Not Before) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5
	NotBefore *NumericDate `json:"nbf,omitempty"`

	// the `iat` (Issued At) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6
	IssuedAt *NumericDate `json:"iat,omitempty"`

	// the `jti` (JWT ID) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7
	ID string `json:"jti,omitempty"`
}

// Valid validates time based claims "exp, iat, nbf"
func (c RegisteredClaims) Valid() error {
	vErr := new(ValidationError)
	now := TimeFunc()

	// don't fail verification if optional claims use default values
	if !c.VerifyExpiresAt(now, false) {
		delta := now.Sub(c.ExpiresAt.Time)
		vErr.Inner = fmt.Errorf("%s by %s", ErrTokenExpired, delta) //nolint:goerr113
		vErr.Errors |= ValidationErrorExpired
	}

	if !c.VerifyIssuedAt(now, false) {
		vErr.Inner = ErrTokenUsedBeforeIssued
		vErr.Errors |= ValidationErrorIssuedAt
	}

	if !c.VerifyNotBefore(now, false) {
		vErr.Inner = ErrTokenNotValidYet
		vErr.Errors |= ValidationErrorNotValidYet
	}

	if vErr.valid() {
		return nil
	}

	return vErr
}

// VerifyAudience compares the aud claim against cmp
func (c *RegisteredClaims) VerifyAudience(cmp string, req bool) bool {
	return verifyAud(c.Audience, cmp, req)
}

// VerifyExpiresAt compares the exp claim against cmp (cmp < exp)
func (c *RegisteredClaims) VerifyExpiresAt(cmp time.Time, req bool) bool {
	if c.ExpiresAt == nil {
		return verifyExp(nil, cmp, req)
	}

	return verifyExp(&c.ExpiresAt.Time, cmp, req)
}

// VerifyIssuedAt compares the iat claim against cmp (cmp >= iat)
func (c *RegisteredClaims) VerifyIssuedAt(cmp time.Time, req bool) bool {
	if c.IssuedAt == nil {
		return verifyIat(nil, cmp, req)
	}

	return verifyIat(&c.IssuedAt.Time, cmp, req)
}

// VerifyNotBefore compares the nbf claim against cmp (cmp >= nbf)
func (c *RegisteredClaims) VerifyNotBefore(cmp time.Time, req bool) bool {
	if c.NotBefore == nil {
		return verifyNbf(nil, cmp, req)
	}

	return verifyNbf(&c.NotBefore.Time, cmp, req)
}

// VerifyIssuer compares the iss claim against cmp
func (c *RegisteredClaims) VerifyIssuer(cmp string, req bool) bool {
	return verifyIss(c.Issuer, cmp, req)
}

// ParseOrgID returns the organization ID in the claims
func (c Claims) ParseOrgID() ulid.ULID {
	orgID, err := ulid.Parse(c.OrgID)
	if err != nil {
		return ulids.Null
	}

	return orgID
}

// ParseUserID returns the ID of the user from the Subject of the claims
func (c Claims) ParseUserID() ulid.ULID {
	userID, err := ulid.Parse(c.Subject)
	if err != nil {
		return ulids.Null
	}

	return userID
}

func verifyAud(aud []string, cmp string, required bool) bool {
	if len(aud) == 0 {
		return !required
	}
	// use a var here to keep constant time compare when looping over a number of claims
	result := false

	var stringClaims string

	for _, a := range aud {
		if subtle.ConstantTimeCompare([]byte(a), []byte(cmp)) != 0 {
			result = true
		}

		stringClaims += a
	}

	// case where "" is sent in one or many aud claims
	if len(stringClaims) == 0 {
		return !required
	}

	return result
}

func verifyExp(exp *time.Time, now time.Time, required bool) bool {
	if exp == nil {
		return !required
	}

	return now.Before(*exp)
}

func verifyIat(iat *time.Time, now time.Time, required bool) bool {
	if iat == nil {
		return !required
	}

	return now.After(*iat) || now.Equal(*iat)
}

func verifyNbf(nbf *time.Time, now time.Time, required bool) bool {
	if nbf == nil {
		return !required
	}

	return now.After(*nbf) || now.Equal(*nbf)
}

func verifyIss(iss string, cmp string, required bool) bool {
	if iss == "" {
		return !required
	}

	if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) != 0 {
		return true
	} else {
		return false
	}
}
