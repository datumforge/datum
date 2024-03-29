package enums

import (
	"fmt"
	"io"
	"strings"
)

type JoinPolicy string

var (
	Open                JoinPolicy = "OPEN"
	InviteOnly          JoinPolicy = "INVITE_ONLY"
	ApplicationOnly     JoinPolicy = "APPLICATION_ONLY"
	InviteOrApplication JoinPolicy = "INVITE_OR_APPLICATION"
)

// Values returns a slice of strings that represents all the possible values of the JoinPolicy enum.
// Possible default values are "OPEN", "INVITE_ONLY", "APPLICATION_ONLY", and "INVITE_OR_APPLICATION".
func (JoinPolicy) Values() (kinds []string) {
	for _, s := range []JoinPolicy{Open, InviteOnly, ApplicationOnly, InviteOrApplication} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the JoinPolicy as a string
func (r JoinPolicy) String() string {
	return string(r)
}

// ToGroupJoinPolicy returns the user status enum based on string input
func ToGroupJoinPolicy(r string) *JoinPolicy {
	switch r := strings.ToUpper(r); r {
	case Open.String():
		return &Open
	case InviteOnly.String():
		return &InviteOnly
	case ApplicationOnly.String():
		return &ApplicationOnly
	case InviteOrApplication.String():
		return &InviteOrApplication
	default:
		return &InviteOnly
	}
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r JoinPolicy) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *JoinPolicy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for JoinPolicy, got: %T", v) //nolint:goerr113
	}

	*r = JoinPolicy(str)

	return nil
}
