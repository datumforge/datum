// Code generated by ent, DO NOT EDIT.

package invite

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
	// Label holds the string label denoting the invite type in the database.
	Label = "invite"
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
	// FieldMappingID holds the string denoting the mapping_id field in the database.
	FieldMappingID = "mapping_id"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldExpires holds the string denoting the expires field in the database.
	FieldExpires = "expires"
	// FieldRecipient holds the string denoting the recipient field in the database.
	FieldRecipient = "recipient"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldSendAttempts holds the string denoting the send_attempts field in the database.
	FieldSendAttempts = "send_attempts"
	// FieldRequestorID holds the string denoting the requestor_id field in the database.
	FieldRequestorID = "requestor_id"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// Table holds the table name of the invite in the database.
	Table = "invites"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "invites"
	// OwnerInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OwnerInverseTable = "organizations"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
	// EventsTable is the table that holds the events relation/edge. The primary key declared below.
	EventsTable = "invite_events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
)

// Columns holds all SQL columns for invite fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldMappingID,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldOwnerID,
	FieldToken,
	FieldExpires,
	FieldRecipient,
	FieldStatus,
	FieldRole,
	FieldSendAttempts,
	FieldRequestorID,
	FieldSecret,
}

var (
	// EventsPrimaryKey and EventsColumn2 are the table columns denoting the
	// primary key for the events relation (M2M).
	EventsPrimaryKey = []string{"invite_id", "event_id"}
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
	Hooks        [6]ent.Hook
	Interceptors [2]ent.Interceptor
	Policy       ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMappingID holds the default value on creation for the "mapping_id" field.
	DefaultMappingID func() string
	// OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	OwnerIDValidator func(string) error
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func(string) error
	// RecipientValidator is a validator for the "recipient" field. It is called by the builders before save.
	RecipientValidator func(string) error
	// DefaultSendAttempts holds the default value on creation for the "send_attempts" field.
	DefaultSendAttempts int
	// RequestorIDValidator is a validator for the "requestor_id" field. It is called by the builders before save.
	RequestorIDValidator func(string) error
	// SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	SecretValidator func([]byte) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

const DefaultStatus enums.InviteStatus = "INVITATION_SENT"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.InviteStatus) error {
	switch s.String() {
	case "INVITATION_SENT", "APPROVAL_REQUIRED", "INVITATION_ACCEPTED", "INVITATION_EXPIRED":
		return nil
	default:
		return fmt.Errorf("invite: invalid enum value for status field: %q", s)
	}
}

const DefaultRole enums.Role = "MEMBER"

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r enums.Role) error {
	switch r.String() {
	case "ADMIN", "MEMBER", "USER":
		return nil
	default:
		return fmt.Errorf("invite: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the Invite queries.
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

// ByMappingID orders the results by the mapping_id field.
func ByMappingID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMappingID, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByExpires orders the results by the expires field.
func ByExpires(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpires, opts...).ToFunc()
}

// ByRecipient orders the results by the recipient field.
func ByRecipient(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecipient, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// BySendAttempts orders the results by the send_attempts field.
func BySendAttempts(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSendAttempts, opts...).ToFunc()
}

// ByRequestorID orders the results by the requestor_id field.
func ByRequestorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestorID, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByEventsCount orders the results by events count.
func ByEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventsStep(), opts...)
	}
}

// ByEvents orders the results by events terms.
func ByEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, EventsTable, EventsPrimaryKey...),
	)
}

var (
	// enums.InviteStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.InviteStatus)(nil)
	// enums.InviteStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.InviteStatus)(nil)
)

var (
	// enums.Role must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.Role)(nil)
	// enums.Role must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.Role)(nil)
)
