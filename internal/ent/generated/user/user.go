// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/enums"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldAvatarRemoteURL holds the string denoting the avatar_remote_url field in the database.
	FieldAvatarRemoteURL = "avatar_remote_url"
	// FieldAvatarLocalFile holds the string denoting the avatar_local_file field in the database.
	FieldAvatarLocalFile = "avatar_local_file"
	// FieldAvatarUpdatedAt holds the string denoting the avatar_updated_at field in the database.
	FieldAvatarUpdatedAt = "avatar_updated_at"
	// FieldLastSeen holds the string denoting the last_seen field in the database.
	FieldLastSeen = "last_seen"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSub holds the string denoting the sub field in the database.
	FieldSub = "sub"
	// FieldAuthProvider holds the string denoting the auth_provider field in the database.
	FieldAuthProvider = "auth_provider"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgePersonalAccessTokens holds the string denoting the personal_access_tokens edge name in mutations.
	EdgePersonalAccessTokens = "personal_access_tokens"
	// EdgeTfaSettings holds the string denoting the tfa_settings edge name in mutations.
	EdgeTfaSettings = "tfa_settings"
	// EdgeSetting holds the string denoting the setting edge name in mutations.
	EdgeSetting = "setting"
	// EdgeEmailVerificationTokens holds the string denoting the email_verification_tokens edge name in mutations.
	EdgeEmailVerificationTokens = "email_verification_tokens"
	// EdgePasswordResetTokens holds the string denoting the password_reset_tokens edge name in mutations.
	EdgePasswordResetTokens = "password_reset_tokens"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeOrganizations holds the string denoting the organizations edge name in mutations.
	EdgeOrganizations = "organizations"
	// EdgeWebauthn holds the string denoting the webauthn edge name in mutations.
	EdgeWebauthn = "webauthn"
	// EdgeGroupMemberships holds the string denoting the group_memberships edge name in mutations.
	EdgeGroupMemberships = "group_memberships"
	// EdgeOrgMemberships holds the string denoting the org_memberships edge name in mutations.
	EdgeOrgMemberships = "org_memberships"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PersonalAccessTokensTable is the table that holds the personal_access_tokens relation/edge.
	PersonalAccessTokensTable = "personal_access_tokens"
	// PersonalAccessTokensInverseTable is the table name for the PersonalAccessToken entity.
	// It exists in this package in order to avoid circular dependency with the "personalaccesstoken" package.
	PersonalAccessTokensInverseTable = "personal_access_tokens"
	// PersonalAccessTokensColumn is the table column denoting the personal_access_tokens relation/edge.
	PersonalAccessTokensColumn = "owner_id"
	// TfaSettingsTable is the table that holds the tfa_settings relation/edge.
	TfaSettingsTable = "tfa_settings"
	// TfaSettingsInverseTable is the table name for the TFASetting entity.
	// It exists in this package in order to avoid circular dependency with the "tfasetting" package.
	TfaSettingsInverseTable = "tfa_settings"
	// TfaSettingsColumn is the table column denoting the tfa_settings relation/edge.
	TfaSettingsColumn = "owner_id"
	// SettingTable is the table that holds the setting relation/edge.
	SettingTable = "user_settings"
	// SettingInverseTable is the table name for the UserSetting entity.
	// It exists in this package in order to avoid circular dependency with the "usersetting" package.
	SettingInverseTable = "user_settings"
	// SettingColumn is the table column denoting the setting relation/edge.
	SettingColumn = "user_id"
	// EmailVerificationTokensTable is the table that holds the email_verification_tokens relation/edge.
	EmailVerificationTokensTable = "email_verification_tokens"
	// EmailVerificationTokensInverseTable is the table name for the EmailVerificationToken entity.
	// It exists in this package in order to avoid circular dependency with the "emailverificationtoken" package.
	EmailVerificationTokensInverseTable = "email_verification_tokens"
	// EmailVerificationTokensColumn is the table column denoting the email_verification_tokens relation/edge.
	EmailVerificationTokensColumn = "owner_id"
	// PasswordResetTokensTable is the table that holds the password_reset_tokens relation/edge.
	PasswordResetTokensTable = "password_reset_tokens"
	// PasswordResetTokensInverseTable is the table name for the PasswordResetToken entity.
	// It exists in this package in order to avoid circular dependency with the "passwordresettoken" package.
	PasswordResetTokensInverseTable = "password_reset_tokens"
	// PasswordResetTokensColumn is the table column denoting the password_reset_tokens relation/edge.
	PasswordResetTokensColumn = "owner_id"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_memberships"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// OrganizationsTable is the table that holds the organizations relation/edge. The primary key declared below.
	OrganizationsTable = "org_memberships"
	// OrganizationsInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationsInverseTable = "organizations"
	// WebauthnTable is the table that holds the webauthn relation/edge.
	WebauthnTable = "webauthns"
	// WebauthnInverseTable is the table name for the Webauthn entity.
	// It exists in this package in order to avoid circular dependency with the "webauthn" package.
	WebauthnInverseTable = "webauthns"
	// WebauthnColumn is the table column denoting the webauthn relation/edge.
	WebauthnColumn = "owner_id"
	// GroupMembershipsTable is the table that holds the group_memberships relation/edge.
	GroupMembershipsTable = "group_memberships"
	// GroupMembershipsInverseTable is the table name for the GroupMembership entity.
	// It exists in this package in order to avoid circular dependency with the "groupmembership" package.
	GroupMembershipsInverseTable = "group_memberships"
	// GroupMembershipsColumn is the table column denoting the group_memberships relation/edge.
	GroupMembershipsColumn = "user_id"
	// OrgMembershipsTable is the table that holds the org_memberships relation/edge.
	OrgMembershipsTable = "org_memberships"
	// OrgMembershipsInverseTable is the table name for the OrgMembership entity.
	// It exists in this package in order to avoid circular dependency with the "orgmembership" package.
	OrgMembershipsInverseTable = "org_memberships"
	// OrgMembershipsColumn is the table column denoting the org_memberships relation/edge.
	OrgMembershipsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldEmail,
	FieldFirstName,
	FieldLastName,
	FieldDisplayName,
	FieldAvatarRemoteURL,
	FieldAvatarLocalFile,
	FieldAvatarUpdatedAt,
	FieldLastSeen,
	FieldPassword,
	FieldSub,
	FieldAuthProvider,
	FieldRole,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"user_id", "group_id"}
	// OrganizationsPrimaryKey and OrganizationsColumn2 are the table columns denoting the
	// primary key for the organizations relation (M2M).
	OrganizationsPrimaryKey = []string{"user_id", "organization_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/datumforge/datum/internal/ent/generated/runtime"
var (
	Hooks        [5]ent.Hook
	Interceptors [2]ent.Interceptor
	Policy       ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// AvatarRemoteURLValidator is a validator for the "avatar_remote_url" field. It is called by the builders before save.
	AvatarRemoteURLValidator func(string) error
	// AvatarLocalFileValidator is a validator for the "avatar_local_file" field. It is called by the builders before save.
	AvatarLocalFileValidator func(string) error
	// UpdateDefaultAvatarUpdatedAt holds the default value on update for the "avatar_updated_at" field.
	UpdateDefaultAvatarUpdatedAt func() time.Time
	// UpdateDefaultLastSeen holds the default value on update for the "last_seen" field.
	UpdateDefaultLastSeen func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

const DefaultAuthProvider enums.AuthProvider = "CREDENTIALS"

// AuthProviderValidator is a validator for the "auth_provider" field enum values. It is called by the builders before save.
func AuthProviderValidator(ap enums.AuthProvider) error {
	switch ap.String() {
	case "CREDENTIALS", "GOOGLE", "GITHUB", "WEBAUTHN", "INVALID":
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for auth_provider field: %q", ap)
	}
}

const DefaultRole enums.Role = "USER"

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r enums.Role) error {
	switch r.String() {
	case "ADMIN", "MEMBER", "OWNER", "USER":
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByAvatarRemoteURL orders the results by the avatar_remote_url field.
func ByAvatarRemoteURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarRemoteURL, opts...).ToFunc()
}

// ByAvatarLocalFile orders the results by the avatar_local_file field.
func ByAvatarLocalFile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarLocalFile, opts...).ToFunc()
}

// ByAvatarUpdatedAt orders the results by the avatar_updated_at field.
func ByAvatarUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarUpdatedAt, opts...).ToFunc()
}

// ByLastSeen orders the results by the last_seen field.
func ByLastSeen(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastSeen, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// BySub orders the results by the sub field.
func BySub(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSub, opts...).ToFunc()
}

// ByAuthProvider orders the results by the auth_provider field.
func ByAuthProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthProvider, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByPersonalAccessTokensCount orders the results by personal_access_tokens count.
func ByPersonalAccessTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPersonalAccessTokensStep(), opts...)
	}
}

// ByPersonalAccessTokens orders the results by personal_access_tokens terms.
func ByPersonalAccessTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPersonalAccessTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTfaSettingsCount orders the results by tfa_settings count.
func ByTfaSettingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTfaSettingsStep(), opts...)
	}
}

// ByTfaSettings orders the results by tfa_settings terms.
func ByTfaSettings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTfaSettingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySettingField orders the results by setting field.
func BySettingField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSettingStep(), sql.OrderByField(field, opts...))
	}
}

// ByEmailVerificationTokensCount orders the results by email_verification_tokens count.
func ByEmailVerificationTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEmailVerificationTokensStep(), opts...)
	}
}

// ByEmailVerificationTokens orders the results by email_verification_tokens terms.
func ByEmailVerificationTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmailVerificationTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPasswordResetTokensCount orders the results by password_reset_tokens count.
func ByPasswordResetTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPasswordResetTokensStep(), opts...)
	}
}

// ByPasswordResetTokens orders the results by password_reset_tokens terms.
func ByPasswordResetTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPasswordResetTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrganizationsCount orders the results by organizations count.
func ByOrganizationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrganizationsStep(), opts...)
	}
}

// ByOrganizations orders the results by organizations terms.
func ByOrganizations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWebauthnCount orders the results by webauthn count.
func ByWebauthnCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWebauthnStep(), opts...)
	}
}

// ByWebauthn orders the results by webauthn terms.
func ByWebauthn(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWebauthnStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupMembershipsCount orders the results by group_memberships count.
func ByGroupMembershipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupMembershipsStep(), opts...)
	}
}

// ByGroupMemberships orders the results by group_memberships terms.
func ByGroupMemberships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupMembershipsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrgMembershipsCount orders the results by org_memberships count.
func ByOrgMembershipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrgMembershipsStep(), opts...)
	}
}

// ByOrgMemberships orders the results by org_memberships terms.
func ByOrgMemberships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrgMembershipsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPersonalAccessTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PersonalAccessTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PersonalAccessTokensTable, PersonalAccessTokensColumn),
	)
}
func newTfaSettingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TfaSettingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TfaSettingsTable, TfaSettingsColumn),
	)
}
func newSettingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SettingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SettingTable, SettingColumn),
	)
}
func newEmailVerificationTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmailVerificationTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EmailVerificationTokensTable, EmailVerificationTokensColumn),
	)
}
func newPasswordResetTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PasswordResetTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PasswordResetTokensTable, PasswordResetTokensColumn),
	)
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, GroupsTable, GroupsPrimaryKey...),
	)
}
func newOrganizationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, OrganizationsTable, OrganizationsPrimaryKey...),
	)
}
func newWebauthnStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WebauthnInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WebauthnTable, WebauthnColumn),
	)
}
func newGroupMembershipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupMembershipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, GroupMembershipsTable, GroupMembershipsColumn),
	)
}
func newOrgMembershipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrgMembershipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OrgMembershipsTable, OrgMembershipsColumn),
	)
}

var (
	// enums.AuthProvider must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.AuthProvider)(nil)
	// enums.AuthProvider must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.AuthProvider)(nil)
)

var (
	// enums.Role must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.Role)(nil)
	// enums.Role must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.Role)(nil)
)
