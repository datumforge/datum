// Code generated by ent, DO NOT EDIT.

package generated

import (
	"time"

	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/google/uuid"
)

// CreateGroupInput represents a mutation input for creating groups.
type CreateGroupInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	CreatedBy     *int
	UpdatedBy     *int
	Name          string
	Description   *string
	LogoURL       string
	SettingID     uuid.UUID
	MembershipIDs []uuid.UUID
	UserIDs       []uuid.UUID
}

// Mutate applies the CreateGroupInput on the GroupMutation builder.
func (i *CreateGroupInput) Mutate(m *GroupMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetLogoURL(i.LogoURL)
	m.SetSettingID(i.SettingID)
	if v := i.MembershipIDs; len(v) > 0 {
		m.AddMembershipIDs(v...)
	}
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreateGroupInput on the GroupCreate builder.
func (c *GroupCreate) SetInput(i CreateGroupInput) *GroupCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateGroupInput represents a mutation input for updating groups.
type UpdateGroupInput struct {
	UpdatedAt           *time.Time
	ClearCreatedBy      bool
	CreatedBy           *int
	ClearUpdatedBy      bool
	UpdatedBy           *int
	Name                *string
	Description         *string
	LogoURL             *string
	SettingID           *uuid.UUID
	ClearMemberships    bool
	AddMembershipIDs    []uuid.UUID
	RemoveMembershipIDs []uuid.UUID
	ClearUsers          bool
	AddUserIDs          []uuid.UUID
	RemoveUserIDs       []uuid.UUID
}

// Mutate applies the UpdateGroupInput on the GroupMutation builder.
func (i *UpdateGroupInput) Mutate(m *GroupMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearCreatedBy {
		m.ClearCreatedBy()
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if i.ClearUpdatedBy {
		m.ClearUpdatedBy()
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.LogoURL; v != nil {
		m.SetLogoURL(*v)
	}
	if v := i.SettingID; v != nil {
		m.SetSettingID(*v)
	}
	if i.ClearMemberships {
		m.ClearMemberships()
	}
	if v := i.AddMembershipIDs; len(v) > 0 {
		m.AddMembershipIDs(v...)
	}
	if v := i.RemoveMembershipIDs; len(v) > 0 {
		m.RemoveMembershipIDs(v...)
	}
	if i.ClearUsers {
		m.ClearUsers()
	}
	if v := i.AddUserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.RemoveUserIDs; len(v) > 0 {
		m.RemoveUserIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateGroupInput on the GroupUpdate builder.
func (c *GroupUpdate) SetInput(i UpdateGroupInput) *GroupUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateGroupInput on the GroupUpdateOne builder.
func (c *GroupUpdateOne) SetInput(i UpdateGroupInput) *GroupUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateGroupSettingsInput represents a mutation input for creating groupsettingsslice.
type CreateGroupSettingsInput struct {
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	CreatedBy  *int
	UpdatedBy  *int
	Visibility *groupsettings.Visibility
	JoinPolicy *groupsettings.JoinPolicy
}

// Mutate applies the CreateGroupSettingsInput on the GroupSettingsMutation builder.
func (i *CreateGroupSettingsInput) Mutate(m *GroupSettingsMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Visibility; v != nil {
		m.SetVisibility(*v)
	}
	if v := i.JoinPolicy; v != nil {
		m.SetJoinPolicy(*v)
	}
}

// SetInput applies the change-set in the CreateGroupSettingsInput on the GroupSettingsCreate builder.
func (c *GroupSettingsCreate) SetInput(i CreateGroupSettingsInput) *GroupSettingsCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateGroupSettingsInput represents a mutation input for updating groupsettingsslice.
type UpdateGroupSettingsInput struct {
	UpdatedAt      *time.Time
	ClearCreatedBy bool
	CreatedBy      *int
	ClearUpdatedBy bool
	UpdatedBy      *int
	Visibility     *groupsettings.Visibility
	JoinPolicy     *groupsettings.JoinPolicy
}

// Mutate applies the UpdateGroupSettingsInput on the GroupSettingsMutation builder.
func (i *UpdateGroupSettingsInput) Mutate(m *GroupSettingsMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearCreatedBy {
		m.ClearCreatedBy()
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if i.ClearUpdatedBy {
		m.ClearUpdatedBy()
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Visibility; v != nil {
		m.SetVisibility(*v)
	}
	if v := i.JoinPolicy; v != nil {
		m.SetJoinPolicy(*v)
	}
}

// SetInput applies the change-set in the UpdateGroupSettingsInput on the GroupSettingsUpdate builder.
func (c *GroupSettingsUpdate) SetInput(i UpdateGroupSettingsInput) *GroupSettingsUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateGroupSettingsInput on the GroupSettingsUpdateOne builder.
func (c *GroupSettingsUpdateOne) SetInput(i UpdateGroupSettingsInput) *GroupSettingsUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateIntegrationInput represents a mutation input for creating integrations.
type CreateIntegrationInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	CreatedBy      *int
	UpdatedBy      *int
	Kind           string
	Description    *string
	SecretName     string
	OrganizationID uuid.UUID
}

// Mutate applies the CreateIntegrationInput on the IntegrationMutation builder.
func (i *CreateIntegrationInput) Mutate(m *IntegrationMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	m.SetKind(i.Kind)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetSecretName(i.SecretName)
	m.SetOrganizationID(i.OrganizationID)
}

// SetInput applies the change-set in the CreateIntegrationInput on the IntegrationCreate builder.
func (c *IntegrationCreate) SetInput(i CreateIntegrationInput) *IntegrationCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateIntegrationInput represents a mutation input for updating integrations.
type UpdateIntegrationInput struct {
	UpdatedAt        *time.Time
	ClearCreatedBy   bool
	CreatedBy        *int
	ClearUpdatedBy   bool
	UpdatedBy        *int
	ClearDescription bool
	Description      *string
	OrganizationID   *uuid.UUID
}

// Mutate applies the UpdateIntegrationInput on the IntegrationMutation builder.
func (i *UpdateIntegrationInput) Mutate(m *IntegrationMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearCreatedBy {
		m.ClearCreatedBy()
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if i.ClearUpdatedBy {
		m.ClearUpdatedBy()
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.OrganizationID; v != nil {
		m.SetOrganizationID(*v)
	}
}

// SetInput applies the change-set in the UpdateIntegrationInput on the IntegrationUpdate builder.
func (c *IntegrationUpdate) SetInput(i UpdateIntegrationInput) *IntegrationUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateIntegrationInput on the IntegrationUpdateOne builder.
func (c *IntegrationUpdateOne) SetInput(i UpdateIntegrationInput) *IntegrationUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMembershipInput represents a mutation input for creating memberships.
type CreateMembershipInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	CreatedBy      *int
	UpdatedBy      *int
	Current        *bool
	OrganizationID uuid.UUID
	UserID         uuid.UUID
	GroupID        uuid.UUID
}

// Mutate applies the CreateMembershipInput on the MembershipMutation builder.
func (i *CreateMembershipInput) Mutate(m *MembershipMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Current; v != nil {
		m.SetCurrent(*v)
	}
	m.SetOrganizationID(i.OrganizationID)
	m.SetUserID(i.UserID)
	m.SetGroupID(i.GroupID)
}

// SetInput applies the change-set in the CreateMembershipInput on the MembershipCreate builder.
func (c *MembershipCreate) SetInput(i CreateMembershipInput) *MembershipCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMembershipInput represents a mutation input for updating memberships.
type UpdateMembershipInput struct {
	UpdatedAt      *time.Time
	ClearCreatedBy bool
	CreatedBy      *int
	ClearUpdatedBy bool
	UpdatedBy      *int
	Current        *bool
	OrganizationID *uuid.UUID
	UserID         *uuid.UUID
	GroupID        *uuid.UUID
}

// Mutate applies the UpdateMembershipInput on the MembershipMutation builder.
func (i *UpdateMembershipInput) Mutate(m *MembershipMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearCreatedBy {
		m.ClearCreatedBy()
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if i.ClearUpdatedBy {
		m.ClearUpdatedBy()
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Current; v != nil {
		m.SetCurrent(*v)
	}
	if v := i.OrganizationID; v != nil {
		m.SetOrganizationID(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.GroupID; v != nil {
		m.SetGroupID(*v)
	}
}

// SetInput applies the change-set in the UpdateMembershipInput on the MembershipUpdate builder.
func (c *MembershipUpdate) SetInput(i UpdateMembershipInput) *MembershipUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMembershipInput on the MembershipUpdateOne builder.
func (c *MembershipUpdateOne) SetInput(i UpdateMembershipInput) *MembershipUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateOrganizationInput represents a mutation input for creating organizations.
type CreateOrganizationInput struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	CreatedBy      *int
	UpdatedBy      *int
	Name           string
	MembershipIDs  []uuid.UUID
	IntegrationIDs []uuid.UUID
}

// Mutate applies the CreateOrganizationInput on the OrganizationMutation builder.
func (i *CreateOrganizationInput) Mutate(m *OrganizationMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	m.SetName(i.Name)
	if v := i.MembershipIDs; len(v) > 0 {
		m.AddMembershipIDs(v...)
	}
	if v := i.IntegrationIDs; len(v) > 0 {
		m.AddIntegrationIDs(v...)
	}
}

// SetInput applies the change-set in the CreateOrganizationInput on the OrganizationCreate builder.
func (c *OrganizationCreate) SetInput(i CreateOrganizationInput) *OrganizationCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrganizationInput represents a mutation input for updating organizations.
type UpdateOrganizationInput struct {
	UpdatedAt            *time.Time
	ClearCreatedBy       bool
	CreatedBy            *int
	ClearUpdatedBy       bool
	UpdatedBy            *int
	Name                 *string
	ClearMemberships     bool
	AddMembershipIDs     []uuid.UUID
	RemoveMembershipIDs  []uuid.UUID
	ClearIntegrations    bool
	AddIntegrationIDs    []uuid.UUID
	RemoveIntegrationIDs []uuid.UUID
}

// Mutate applies the UpdateOrganizationInput on the OrganizationMutation builder.
func (i *UpdateOrganizationInput) Mutate(m *OrganizationMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearCreatedBy {
		m.ClearCreatedBy()
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if i.ClearUpdatedBy {
		m.ClearUpdatedBy()
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearMemberships {
		m.ClearMemberships()
	}
	if v := i.AddMembershipIDs; len(v) > 0 {
		m.AddMembershipIDs(v...)
	}
	if v := i.RemoveMembershipIDs; len(v) > 0 {
		m.RemoveMembershipIDs(v...)
	}
	if i.ClearIntegrations {
		m.ClearIntegrations()
	}
	if v := i.AddIntegrationIDs; len(v) > 0 {
		m.AddIntegrationIDs(v...)
	}
	if v := i.RemoveIntegrationIDs; len(v) > 0 {
		m.RemoveIntegrationIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdate builder.
func (c *OrganizationUpdate) SetInput(i UpdateOrganizationInput) *OrganizationUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdateOne builder.
func (c *OrganizationUpdateOne) SetInput(i UpdateOrganizationInput) *OrganizationUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateSessionInput represents a mutation input for creating sessions.
type CreateSessionInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	CreatedBy *int
	UpdatedBy *int
	Type      session.Type
	Disabled  bool
	Token     *string
	UserAgent *string
	Ips       string
	UsersID   *uuid.UUID
}

// Mutate applies the CreateSessionInput on the SessionMutation builder.
func (i *CreateSessionInput) Mutate(m *SessionMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	m.SetType(i.Type)
	m.SetDisabled(i.Disabled)
	if v := i.Token; v != nil {
		m.SetToken(*v)
	}
	if v := i.UserAgent; v != nil {
		m.SetUserAgent(*v)
	}
	m.SetIps(i.Ips)
	if v := i.UsersID; v != nil {
		m.SetUsersID(*v)
	}
}

// SetInput applies the change-set in the CreateSessionInput on the SessionCreate builder.
func (c *SessionCreate) SetInput(i CreateSessionInput) *SessionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateSessionInput represents a mutation input for updating sessions.
type UpdateSessionInput struct {
	UpdatedAt      *time.Time
	ClearCreatedBy bool
	CreatedBy      *int
	ClearUpdatedBy bool
	UpdatedBy      *int
	Disabled       *bool
	ClearUserAgent bool
	UserAgent      *string
	Ips            *string
	ClearUsers     bool
	UsersID        *uuid.UUID
}

// Mutate applies the UpdateSessionInput on the SessionMutation builder.
func (i *UpdateSessionInput) Mutate(m *SessionMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearCreatedBy {
		m.ClearCreatedBy()
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if i.ClearUpdatedBy {
		m.ClearUpdatedBy()
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Disabled; v != nil {
		m.SetDisabled(*v)
	}
	if i.ClearUserAgent {
		m.ClearUserAgent()
	}
	if v := i.UserAgent; v != nil {
		m.SetUserAgent(*v)
	}
	if v := i.Ips; v != nil {
		m.SetIps(*v)
	}
	if i.ClearUsers {
		m.ClearUsers()
	}
	if v := i.UsersID; v != nil {
		m.SetUsersID(*v)
	}
}

// SetInput applies the change-set in the UpdateSessionInput on the SessionUpdate builder.
func (c *SessionUpdate) SetInput(i UpdateSessionInput) *SessionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateSessionInput on the SessionUpdateOne builder.
func (c *SessionUpdateOne) SetInput(i UpdateSessionInput) *SessionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	CreatedBy       *int
	UpdatedBy       *int
	Email           string
	FirstName       string
	LastName        string
	DisplayName     *string
	Locked          *bool
	AvatarRemoteURL *string
	AvatarLocalFile *string
	AvatarUpdatedAt *time.Time
	SilencedAt      *time.Time
	SuspendedAt     *time.Time
	RecoveryCode    *string
	MembershipIDs   []uuid.UUID
	SessionIDs      []uuid.UUID
	GroupIDs        []uuid.UUID
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	m.SetEmail(i.Email)
	m.SetFirstName(i.FirstName)
	m.SetLastName(i.LastName)
	if v := i.DisplayName; v != nil {
		m.SetDisplayName(*v)
	}
	if v := i.Locked; v != nil {
		m.SetLocked(*v)
	}
	if v := i.AvatarRemoteURL; v != nil {
		m.SetAvatarRemoteURL(*v)
	}
	if v := i.AvatarLocalFile; v != nil {
		m.SetAvatarLocalFile(*v)
	}
	if v := i.AvatarUpdatedAt; v != nil {
		m.SetAvatarUpdatedAt(*v)
	}
	if v := i.SilencedAt; v != nil {
		m.SetSilencedAt(*v)
	}
	if v := i.SuspendedAt; v != nil {
		m.SetSuspendedAt(*v)
	}
	if v := i.RecoveryCode; v != nil {
		m.SetRecoveryCode(*v)
	}
	if v := i.MembershipIDs; len(v) > 0 {
		m.AddMembershipIDs(v...)
	}
	if v := i.SessionIDs; len(v) > 0 {
		m.AddSessionIDs(v...)
	}
	if v := i.GroupIDs; len(v) > 0 {
		m.AddGroupIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	UpdatedAt            *time.Time
	ClearCreatedBy       bool
	CreatedBy            *int
	ClearUpdatedBy       bool
	UpdatedBy            *int
	Email                *string
	FirstName            *string
	LastName             *string
	DisplayName          *string
	Locked               *bool
	ClearAvatarRemoteURL bool
	AvatarRemoteURL      *string
	ClearAvatarLocalFile bool
	AvatarLocalFile      *string
	ClearAvatarUpdatedAt bool
	AvatarUpdatedAt      *time.Time
	ClearSilencedAt      bool
	SilencedAt           *time.Time
	ClearSuspendedAt     bool
	SuspendedAt          *time.Time
	ClearRecoveryCode    bool
	RecoveryCode         *string
	ClearMemberships     bool
	AddMembershipIDs     []uuid.UUID
	RemoveMembershipIDs  []uuid.UUID
	ClearSessions        bool
	AddSessionIDs        []uuid.UUID
	RemoveSessionIDs     []uuid.UUID
	ClearGroups          bool
	AddGroupIDs          []uuid.UUID
	RemoveGroupIDs       []uuid.UUID
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearCreatedBy {
		m.ClearCreatedBy()
	}
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if i.ClearUpdatedBy {
		m.ClearUpdatedBy()
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.FirstName; v != nil {
		m.SetFirstName(*v)
	}
	if v := i.LastName; v != nil {
		m.SetLastName(*v)
	}
	if v := i.DisplayName; v != nil {
		m.SetDisplayName(*v)
	}
	if v := i.Locked; v != nil {
		m.SetLocked(*v)
	}
	if i.ClearAvatarRemoteURL {
		m.ClearAvatarRemoteURL()
	}
	if v := i.AvatarRemoteURL; v != nil {
		m.SetAvatarRemoteURL(*v)
	}
	if i.ClearAvatarLocalFile {
		m.ClearAvatarLocalFile()
	}
	if v := i.AvatarLocalFile; v != nil {
		m.SetAvatarLocalFile(*v)
	}
	if i.ClearAvatarUpdatedAt {
		m.ClearAvatarUpdatedAt()
	}
	if v := i.AvatarUpdatedAt; v != nil {
		m.SetAvatarUpdatedAt(*v)
	}
	if i.ClearSilencedAt {
		m.ClearSilencedAt()
	}
	if v := i.SilencedAt; v != nil {
		m.SetSilencedAt(*v)
	}
	if i.ClearSuspendedAt {
		m.ClearSuspendedAt()
	}
	if v := i.SuspendedAt; v != nil {
		m.SetSuspendedAt(*v)
	}
	if i.ClearRecoveryCode {
		m.ClearRecoveryCode()
	}
	if v := i.RecoveryCode; v != nil {
		m.SetRecoveryCode(*v)
	}
	if i.ClearMemberships {
		m.ClearMemberships()
	}
	if v := i.AddMembershipIDs; len(v) > 0 {
		m.AddMembershipIDs(v...)
	}
	if v := i.RemoveMembershipIDs; len(v) > 0 {
		m.RemoveMembershipIDs(v...)
	}
	if i.ClearSessions {
		m.ClearSessions()
	}
	if v := i.AddSessionIDs; len(v) > 0 {
		m.AddSessionIDs(v...)
	}
	if v := i.RemoveSessionIDs; len(v) > 0 {
		m.RemoveSessionIDs(v...)
	}
	if i.ClearGroups {
		m.ClearGroups()
	}
	if v := i.AddGroupIDs; len(v) > 0 {
		m.AddGroupIDs(v...)
	}
	if v := i.RemoveGroupIDs; len(v) > 0 {
		m.RemoveGroupIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
