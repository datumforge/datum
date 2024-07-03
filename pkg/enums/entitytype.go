package enums

import (
	"fmt"
	"io"
	"strings"
)

type EntityType string

var (
	// Organization are entities that represent an organization for the end user customer
	Organization EntityType = "ORGANIZATION"
	// Vendors are entities that represent a vendor for the end user customer
	Vendor EntityType = "VENDOR"
	// EntityTypeInvalid is the default value for the EntityType enum
	EntityTypeInvalid EntityType = "INVALID"
)

// Values returns a slice of strings that represents all the possible values of the EntityType enum.
// Possible default values are "ORGANIZATION", "VENDOR"
func (EntityType) Values() (kinds []string) {
	for _, s := range []EntityType{Organization, Vendor} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the EntityType as a string
func (r EntityType) String() string {
	return string(r)
}

// ToEntityType returns the user status enum based on string input
func ToEntityType(r string) *EntityType {
	switch r := strings.ToUpper(r); r {
	case Organization.String():
		return &Organization
	case Vendor.String():
		return &Vendor
	default:
		return &EntityTypeInvalid
	}
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r EntityType) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *EntityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for EntityType, got: %T", v) //nolint:err113
	}

	*r = EntityType(str)

	return nil
}
