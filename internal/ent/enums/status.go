package enums

import (
	"fmt"
	"io"
)

type UserStatus string

var (
	Active      UserStatus = "ACTIVE"
	Inactive    UserStatus = "INACTIVE"
	Deactivated UserStatus = "DEACTIVATED"
	Suspended   UserStatus = "SUSPENDED"
)

// Values returns a slice of strings that represents all the possible values of the UserStatus enum.
// Possible default values are "ACTIVE", "INACTIVE", "DEACTIVATED", and "SUSPENDED".
func (UserStatus) Values() (kinds []string) {
	for _, s := range []UserStatus{Active, Inactive, Deactivated, Suspended} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the UserStatus as a string
func (r UserStatus) String() string {
	return string(r)
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r UserStatus) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for UserStatus, got: %T", v) //nolint:goerr113
	}

	*r = UserStatus(str)

	return nil
}
