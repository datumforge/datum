package enums

import (
	"fmt"
	"io"
	"strings"
)

type InviteStatus string

var (
	InvitationSent     InviteStatus = "INVITATION_SENT"
	ApprovalRequired   InviteStatus = "APPROVAL_REQUIRED"
	InvitationAccepted InviteStatus = "INVITATION_ACCEPTED"
	InvitationExpired  InviteStatus = "INVITATION_EXPIRED"
	InviteInvalid      InviteStatus = "INVITE_INVALID"
)

// Values returns a slice of strings that represents all the possible values of the Role enum.
// Possible default values are "ADMIN", and "MEMBER".
func (InviteStatus) Values() (kinds []string) {
	for _, s := range []InviteStatus{InvitationSent} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the role as a string
func (r InviteStatus) String() string {
	return string(r)
}

// InviteEnum returns the Role based on string input
func InviteEnum(r string) InviteStatus {
	switch r := strings.ToUpper(r); r {
	case InvitationSent.String():
		return InvitationSent
	case ApprovalRequired.String():
		return ApprovalRequired
	case InvitationAccepted.String():
		return InvitationAccepted
	case InvitationExpired.String():
		return InvitationExpired
	default:
		return InviteInvalid
	}
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r InviteStatus) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *InviteStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for Role, got: %T", v) //nolint:goerr113
	}

	*r = InviteStatus(str)

	return nil
}
