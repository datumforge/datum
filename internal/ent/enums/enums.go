package enums

import (
	"fmt"
	"io"
)

type Role string

const (
	RoleOwner    Role = "OWNER"
	RoleAdmin    Role = "ADMIN"
	RoleMember   Role = "MEMBER"
	RoleCustomer Role = "CUSTOMER"
)

// Values returns a slice of strings that represents all the possible values of the Role enum.
// Possible values are "OWNER", "ADMIN", and "MEMBER".
func (Role) Values() (kinds []string) {
	for _, s := range []Role{RoleOwner, RoleAdmin, RoleMember} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the role as a string
func (r Role) String() string {
	return string(r)
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r Role) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for Role, got: %T", v) //nolint:goerr113
	}

	*r = Role(str)

	return nil
}
