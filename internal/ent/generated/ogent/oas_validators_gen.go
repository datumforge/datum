// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"github.com/go-faster/errors"
)

func (s ListIntegrationOKApplicationJSON) Validate() error {
	alias := ([]IntegrationList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s ListMembershipOKApplicationJSON) Validate() error {
	alias := ([]MembershipList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s ListOrganizationIntegrationsOKApplicationJSON) Validate() error {
	alias := ([]OrganizationIntegrationsList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s ListOrganizationMembershipsOKApplicationJSON) Validate() error {
	alias := ([]OrganizationMembershipsList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s ListOrganizationOKApplicationJSON) Validate() error {
	alias := ([]OrganizationList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s ListUserMembershipsOKApplicationJSON) Validate() error {
	alias := ([]UserMembershipsList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s ListUserOKApplicationJSON) Validate() error {
	alias := ([]UserList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
