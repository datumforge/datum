package tokens

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/utils/ulids"
)

// Claims implements custom claims and extends the `jwt.RegisteredClaims` struct; we will store user-related elements here (and thus in the JWT Token) for reference / validation
type Claims struct {
	jwt.RegisteredClaims
	// UserID is the internal generated ID for the user
	UserID string `json:"user_id,omitempty"`
	// the email associated with the user
	Email string `json:"email,omitempty"`
	// the organization ID the JWT token is valid for
	OrgID string `json:"org,omitempty"`
	// the ID of the parent organization, if a child
	ParentOrgID string `json:"parentorg,omitempty"`
	// the entitlement tier the token is valid for
	Tier string `json:"tier,omitempty"`
}

// ParseUserID returns the ID of the user from the Subject of the claims
func (c Claims) ParseUserID() ulid.ULID {
	userID, err := ulid.Parse(c.Subject)
	if err != nil {
		return ulids.Null
	}

	return userID
}

// ParseOrgID parses and return the organization ID from the `OrgID` field of the claims
func (c Claims) ParseOrgID() ulid.ULID {
	orgID, err := ulid.Parse(c.OrgID)
	if err != nil {
		return ulids.Null
	}

	return orgID
}

// ParseParentOrgID parses and returns the parent organization ID from the ParentOrgID field of the claims
func (c Claims) ParseParentOrgID() ulid.ULID {
	parentOrgID, err := ulid.Parse(c.ParentOrgID)
	if err != nil {
		return ulids.Null
	}

	return parentOrgID
}
