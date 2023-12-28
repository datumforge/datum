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

func (Role) Values() (kinds []string) {
	for _, s := range []Role{RoleOwner, RoleAdmin, RoleMember} {
		kinds = append(kinds, string(s))
	}

	return
}

func (r Role) String() string {
	return string(r)
}

func (r Role) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

func (r *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for Role, got: %T", v)
	}

	*r = Role(str)

	return nil
}
