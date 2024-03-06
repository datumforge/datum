// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"time"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateTFAInput is used for create tfa settings.
type CreateTFAInput struct {
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	ClearUpdatedAt *bool      `json:"clearUpdatedAt,omitempty"`
	UpdatedBy      *string    `json:"updatedBy,omitempty"`
	ClearUpdatedBy *bool      `json:"clearUpdatedBy,omitempty"`
	// TFA secret for the user
	TfaSecret      *string `json:"tfaSecret,omitempty"`
	ResetTfaSecret *bool   `json:"resetTfaSecret,omitempty"`
	// recovery codes for 2fa
	RecoveryCodes       []string `json:"recoveryCodes,omitempty"`
	AppendRecoveryCodes []string `json:"appendRecoveryCodes,omitempty"`
	ResetRecoveryCodes  *bool    `json:"resetRecoveryCodes,omitempty"`
	// specifies a user may complete authentication by verifying an OTP code delivered through SMS
	IsPhoneOtpAllowed      *bool `json:"isPhoneOtpAllowed,omitempty"`
	ClearIsPhoneOtpAllowed *bool `json:"clearIsPhoneOtpAllowed,omitempty"`
	// specifies a user may complete authentication by verifying an OTP code delivered through email
	IsEmailOtpAllowed      *bool `json:"isEmailOtpAllowed,omitempty"`
	ClearIsEmailOtpAllowed *bool `json:"clearIsEmailOtpAllowed,omitempty"`
	// specifies a user may complete authentication by verifying a TOTP code delivered through an authenticator app
	IsTotpAllowed      *bool   `json:"isTotpAllowed,omitempty"`
	ClearIsTotpAllowed *bool   `json:"clearIsTotpAllowed,omitempty"`
	IsTfaEnabled       *bool   `json:"isTfaEnabled,omitempty"`
	ClearIsTfaEnabled  *bool   `json:"clearIsTfaEnabled,omitempty"`
	UserID             *string `json:"userID,omitempty"`
}

type Dummy struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// Return response for createEntitlement mutation
type EntitlementCreatePayload struct {
	// Created entitlement
	Entitlement *generated.Entitlement `json:"entitlement"`
}

// Return response for deleteEntitlement mutation
type EntitlementDeletePayload struct {
	// Deleted entitlement ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateEntitlement mutation
type EntitlementUpdatePayload struct {
	// Updated entitlement
	Entitlement *generated.Entitlement `json:"entitlement"`
}

// Return response for createGroup mutation
type GroupCreatePayload struct {
	// Created group
	Group *generated.Group `json:"group"`
}

// Return response for deleteGroup mutation
type GroupDeletePayload struct {
	// Deleted group ID
	DeletedID string `json:"deletedID"`
}

// Return response for createGroupMembership mutation
type GroupMembershipCreatePayload struct {
	// Created groupMembership
	GroupMembership *generated.GroupMembership `json:"groupMembership"`
}

// Return response for deleteGroupMembership mutation
type GroupMembershipDeletePayload struct {
	// Deleted groupMembership ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateGroupMembership mutation
type GroupMembershipUpdatePayload struct {
	// Updated groupMembership
	GroupMembership *generated.GroupMembership `json:"groupMembership"`
}

// Return response for createGroupSetting mutation
type GroupSettingCreatePayload struct {
	// Created groupSetting
	GroupSetting *generated.GroupSetting `json:"groupSetting"`
}

// Return response for deleteGroupSetting mutation
type GroupSettingDeletePayload struct {
	// Deleted groupSetting ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateGroupSetting mutation
type GroupSettingUpdatePayload struct {
	// Updated groupSetting
	GroupSetting *generated.GroupSetting `json:"groupSetting"`
}

// Return response for updateGroup mutation
type GroupUpdatePayload struct {
	// Updated group
	Group *generated.Group `json:"group"`
}

// Return response for createIntegration mutation
type IntegrationCreatePayload struct {
	// Created integration
	Integration *generated.Integration `json:"integration"`
}

// Return response for deleteIntegration mutation
type IntegrationDeletePayload struct {
	// Deleted integration ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateIntegration mutation
type IntegrationUpdatePayload struct {
	// Updated integration
	Integration *generated.Integration `json:"integration"`
}

// Return response for createInvite mutation
type InviteCreatePayload struct {
	// Created invite
	Invite *generated.Invite `json:"invite"`
}

// Return response for deleteInvite mutation
type InviteDeletePayload struct {
	// Deleted invite ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateInvite mutation
type InviteUpdatePayload struct {
	// Updated invite
	Invite *generated.Invite `json:"invite"`
}

// Return response for createOauthProvider mutation
type OauthProviderCreatePayload struct {
	// Created oauthProvider
	OauthProvider *generated.OauthProvider `json:"oauthProvider"`
}

// Return response for deleteOauthProvider mutation
type OauthProviderDeletePayload struct {
	// Deleted oauthProvider ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOauthProvider mutation
type OauthProviderUpdatePayload struct {
	// Updated oauthProvider
	OauthProvider *generated.OauthProvider `json:"oauthProvider"`
}

// Return response for createOhAuthTooToken mutation
type OhAuthTooTokenCreatePayload struct {
	// Created ohAuthTooToken
	OhAuthTooToken *generated.OhAuthTooToken `json:"ohAuthTooToken"`
}

// Return response for deleteOhAuthTooToken mutation
type OhAuthTooTokenDeletePayload struct {
	// Deleted ohAuthTooToken ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOhAuthTooToken mutation
type OhAuthTooTokenUpdatePayload struct {
	// Updated ohAuthTooToken
	OhAuthTooToken *generated.OhAuthTooToken `json:"ohAuthTooToken"`
}

// Return response for createOrgMembership mutation
type OrgMembershipCreatePayload struct {
	// Created orgMembership
	OrgMembership *generated.OrgMembership `json:"orgMembership"`
}

// Return response for deleteOrgMembership mutation
type OrgMembershipDeletePayload struct {
	// Deleted orgMembership ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOrgMembership mutation
type OrgMembershipUpdatePayload struct {
	// Updated orgMembership
	OrgMembership *generated.OrgMembership `json:"orgMembership"`
}

// Return response for createOrganization mutation
type OrganizationCreatePayload struct {
	// Created organization
	Organization *generated.Organization `json:"organization"`
}

// Return response for deleteOrganization mutation
type OrganizationDeletePayload struct {
	// Deleted organization ID
	DeletedID string `json:"deletedID"`
}

// Return response for createOrganizationSetting mutation
type OrganizationSettingCreatePayload struct {
	// Created organizationSetting
	OrganizationSetting *generated.OrganizationSetting `json:"organizationSetting"`
}

// Return response for deleteOrganizationSetting mutation
type OrganizationSettingDeletePayload struct {
	// Deleted organizationSetting ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOrganizationSetting mutation
type OrganizationSettingUpdatePayload struct {
	// Updated organizationSetting
	OrganizationSetting *generated.OrganizationSetting `json:"organizationSetting"`
}

// Return response for updateOrganization mutation
type OrganizationUpdatePayload struct {
	// Updated organization
	Organization *generated.Organization `json:"organization"`
}

// Return response for createPersonalAccessToken mutation
type PersonalAccessTokenCreatePayload struct {
	// Created personalAccessToken
	PersonalAccessToken *generated.PersonalAccessToken `json:"personalAccessToken"`
}

// Return response for deletePersonalAccessToken mutation
type PersonalAccessTokenDeletePayload struct {
	// Deleted personalAccessToken ID
	DeletedID string `json:"deletedID"`
}

// Return response for updatePersonalAccessToken mutation
type PersonalAccessTokenUpdatePayload struct {
	// Updated personalAccessToken
	PersonalAccessToken *generated.PersonalAccessToken `json:"personalAccessToken"`
}

type Subscription struct {
}

// Return response for createTFASetting mutation
type TFASettingCreatePayload struct {
	// Create TFASetting
	TfaSetting *TfaSetting `json:"tfaSetting"`
}

type TfaSetting struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	// recovery codes for 2fa
	RecoveryCodes []string `json:"recoveryCodes,omitempty"`
	// specifies a user may complete authentication by verifying an OTP code delivered through SMS
	IsPhoneOtpAllowed *bool `json:"isPhoneOtpAllowed,omitempty"`
	// specifies a user may complete authentication by verifying an OTP code delivered through email
	IsEmailOtpAllowed *bool `json:"isEmailOtpAllowed,omitempty"`
	// specifies a user may complete authentication by verifying a TOTP code delivered through an authenticator app
	IsTotpAllowed *bool `json:"isTotpAllowed,omitempty"`
	// whether the user has two factor authentication enabled
	IsTfaEnabled *bool `json:"isTfaEnabled,omitempty"`
}

func (TfaSetting) IsNode() {}

// Return response for createUser mutation
type UserCreatePayload struct {
	// Created user
	User *generated.User `json:"user"`
}

// Return response for deleteUser mutation
type UserDeletePayload struct {
	// Deleted user ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateUserSetting mutation
type UserSettingUpdatePayload struct {
	// Updated userSetting
	UserSetting *generated.UserSetting `json:"userSetting"`
}

// Return response for updateUser mutation
type UserUpdatePayload struct {
	// Updated user
	User *generated.User `json:"user"`
}
