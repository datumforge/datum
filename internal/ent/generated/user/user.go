// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldPasswordHash holds the string denoting the passwordhash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldSub holds the string denoting the sub field in the database.
	FieldSub = "sub"
	// FieldOauth holds the string denoting the oauth field in the database.
	FieldOauth = "oauth"
	// EdgeOrganizations holds the string denoting the organizations edge name in mutations.
	EdgeOrganizations = "organizations"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgePersonalAccessTokens holds the string denoting the personal_access_tokens edge name in mutations.
	EdgePersonalAccessTokens = "personal_access_tokens"
	// EdgeSetting holds the string denoting the setting edge name in mutations.
	EdgeSetting = "setting"
	// EdgeRefreshtoken holds the string denoting the refreshtoken edge name in mutations.
	EdgeRefreshtoken = "refreshtoken"
	// Table holds the table name of the user in the database.
	Table = "users"
	// OrganizationsTable is the table that holds the organizations relation/edge. The primary key declared below.
	OrganizationsTable = "user_organizations"
	// OrganizationsInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationsInverseTable = "organizations"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "user_sessions"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_users"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// PersonalAccessTokensTable is the table that holds the personal_access_tokens relation/edge.
	PersonalAccessTokensTable = "personal_access_tokens"
	// PersonalAccessTokensInverseTable is the table name for the PersonalAccessToken entity.
	// It exists in this package in order to avoid circular dependency with the "personalaccesstoken" package.
	PersonalAccessTokensInverseTable = "personal_access_tokens"
	// PersonalAccessTokensColumn is the table column denoting the personal_access_tokens relation/edge.
	PersonalAccessTokensColumn = "user_id"
	// SettingTable is the table that holds the setting relation/edge.
	SettingTable = "user_settings"
	// SettingInverseTable is the table name for the UserSetting entity.
	// It exists in this package in order to avoid circular dependency with the "usersetting" package.
	SettingInverseTable = "user_settings"
	// SettingColumn is the table column denoting the setting relation/edge.
	SettingColumn = "user_setting"
	// RefreshtokenTable is the table that holds the refreshtoken relation/edge.
	RefreshtokenTable = "refresh_tokens"
	// RefreshtokenInverseTable is the table name for the RefreshToken entity.
	// It exists in this package in order to avoid circular dependency with the "refreshtoken" package.
	RefreshtokenInverseTable = "refresh_tokens"
	// RefreshtokenColumn is the table column denoting the refreshtoken relation/edge.
	RefreshtokenColumn = "user_refreshtoken"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldEmail,
	FieldFirstName,
	FieldLastName,
	FieldDisplayName,
	FieldAvatarRemoteURL,
	FieldAvatarLocalFile,
	FieldAvatarUpdatedAt,
	FieldLastSeen,
	FieldPasswordHash,
	FieldSub,
	FieldOauth,
}

var (
	// OrganizationsPrimaryKey and OrganizationsColumn2 are the table columns denoting the
	// primary key for the organizations relation (M2M).
	OrganizationsPrimaryKey = []string{"user_id", "organization_id"}
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "user_id"}
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
	Hooks [1]ent.Hook
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
	// DefaultDisplayName holds the default value on creation for the "display_name" field.
	DefaultDisplayName string
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
	// DefaultOauth holds the default value on creation for the "oauth" field.
	DefaultOauth bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

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

// ByPasswordHash orders the results by the passwordHash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// BySub orders the results by the sub field.
func BySub(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSub, opts...).ToFunc()
}

// ByOauth orders the results by the oauth field.
func ByOauth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOauth, opts...).ToFunc()
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

// BySessionsCount orders the results by sessions count.
func BySessionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSessionsStep(), opts...)
	}
}

// BySessions orders the results by sessions terms.
func BySessions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSessionsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// BySettingField orders the results by setting field.
func BySettingField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSettingStep(), sql.OrderByField(field, opts...))
	}
}

// ByRefreshtokenCount orders the results by refreshtoken count.
func ByRefreshtokenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRefreshtokenStep(), opts...)
	}
}

// ByRefreshtoken orders the results by refreshtoken terms.
func ByRefreshtoken(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRefreshtokenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrganizationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, OrganizationsTable, OrganizationsPrimaryKey...),
	)
}
func newSessionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SessionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
	)
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
	)
}
func newPersonalAccessTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PersonalAccessTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PersonalAccessTokensTable, PersonalAccessTokensColumn),
	)
}
func newSettingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SettingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SettingTable, SettingColumn),
	)
}
func newRefreshtokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RefreshtokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RefreshtokenTable, RefreshtokenColumn),
	)
}
