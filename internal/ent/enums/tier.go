package enums

import (
	"fmt"
	"io"
)

type Tier string

var (
	Free       Tier = "FREE"
	Pro        Tier = "PRO"
	Enterprise Tier = "ENTERPRISE"
)

// Values returns a slice of strings that represents all the possible values of the Tier enum.
// Possible default values are "FREE", "PRO" and "ENTERPRISE".
func (Tier) Values() (kinds []string) {
	for _, s := range []Tier{Free, Pro, Enterprise} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the Tier as a string
func (r Tier) String() string {
	return string(r)
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r Tier) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *Tier) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for Tier, got: %T", v) //nolint:goerr113
	}

	*r = Tier(str)

	return nil
}
