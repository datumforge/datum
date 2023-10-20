// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
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
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldAvatarRemoteURL holds the string denoting the avatar_remote_url field in the database.
	FieldAvatarRemoteURL = "avatar_remote_url"
	// FieldAvatarLocalFile holds the string denoting the avatar_local_file field in the database.
	FieldAvatarLocalFile = "avatar_local_file"
	// FieldAvatarUpdatedAt holds the string denoting the avatar_updated_at field in the database.
	FieldAvatarUpdatedAt = "avatar_updated_at"
	// FieldSilencedAt holds the string denoting the silenced_at field in the database.
	FieldSilencedAt = "silenced_at"
	// FieldSuspendedAt holds the string denoting the suspended_at field in the database.
	FieldSuspendedAt = "suspended_at"
	// FieldRecoveryCode holds the string denoting the recovery_code field in the database.
	FieldRecoveryCode = "recovery_code"
	// EdgeMemberships holds the string denoting the memberships edge name in mutations.
	EdgeMemberships = "memberships"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MembershipsTable is the table that holds the memberships relation/edge.
	MembershipsTable = "memberships"
	// MembershipsInverseTable is the table name for the Membership entity.
	// It exists in this package in order to avoid circular dependency with the "membership" package.
	MembershipsInverseTable = "memberships"
	// MembershipsColumn is the table column denoting the memberships relation/edge.
	MembershipsColumn = "user_memberships"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "user_sessions"
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
	FieldLocked,
	FieldAvatarRemoteURL,
	FieldAvatarLocalFile,
	FieldAvatarUpdatedAt,
	FieldSilencedAt,
	FieldSuspendedAt,
	FieldRecoveryCode,
}

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
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked bool
	// AvatarRemoteURLValidator is a validator for the "avatar_remote_url" field. It is called by the builders before save.
	AvatarRemoteURLValidator func(string) error
	// AvatarLocalFileValidator is a validator for the "avatar_local_file" field. It is called by the builders before save.
	AvatarLocalFileValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
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

// ByLocked orders the results by the locked field.
func ByLocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocked, opts...).ToFunc()
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

// BySilencedAt orders the results by the silenced_at field.
func BySilencedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSilencedAt, opts...).ToFunc()
}

// BySuspendedAt orders the results by the suspended_at field.
func BySuspendedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuspendedAt, opts...).ToFunc()
}

// ByRecoveryCode orders the results by the recovery_code field.
func ByRecoveryCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecoveryCode, opts...).ToFunc()
}

// ByMembershipsCount orders the results by memberships count.
func ByMembershipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMembershipsStep(), opts...)
	}
}

// ByMemberships orders the results by memberships terms.
func ByMemberships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembershipsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newMembershipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembershipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MembershipsTable, MembershipsColumn),
	)
}
func newSessionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SessionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
	)
}
