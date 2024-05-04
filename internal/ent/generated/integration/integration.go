// Code generated by ent, DO NOT EDIT.

package integration

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the integration type in the database.
	Label = "integration"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeSecrets holds the string denoting the secrets edge name in mutations.
	EdgeSecrets = "secrets"
	// EdgeOauth2tokens holds the string denoting the oauth2tokens edge name in mutations.
	EdgeOauth2tokens = "oauth2tokens"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// Table holds the table name of the integration in the database.
	Table = "integrations"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "integrations"
	// OwnerInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OwnerInverseTable = "organizations"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
	// SecretsTable is the table that holds the secrets relation/edge. The primary key declared below.
	SecretsTable = "integration_secrets"
	// SecretsInverseTable is the table name for the Hush entity.
	// It exists in this package in order to avoid circular dependency with the "hush" package.
	SecretsInverseTable = "hushes"
	// Oauth2tokensTable is the table that holds the oauth2tokens relation/edge. The primary key declared below.
	Oauth2tokensTable = "integration_oauth2tokens"
	// Oauth2tokensInverseTable is the table name for the OhAuthTooToken entity.
	// It exists in this package in order to avoid circular dependency with the "ohauthtootoken" package.
	Oauth2tokensInverseTable = "oh_auth_too_tokens"
	// EventsTable is the table that holds the events relation/edge. The primary key declared below.
	EventsTable = "integration_events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
)

// Columns holds all SQL columns for integration fields.
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
	FieldName,
	FieldDescription,
	FieldKind,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "integrations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_integrations",
}

var (
	// SecretsPrimaryKey and SecretsColumn2 are the table columns denoting the
	// primary key for the secrets relation (M2M).
	SecretsPrimaryKey = []string{"integration_id", "hush_id"}
	// Oauth2tokensPrimaryKey and Oauth2tokensColumn2 are the table columns denoting the
	// primary key for the oauth2tokens relation (M2M).
	Oauth2tokensPrimaryKey = []string{"integration_id", "oh_auth_too_token_id"}
	// EventsPrimaryKey and EventsColumn2 are the table columns denoting the
	// primary key for the events relation (M2M).
	EventsPrimaryKey = []string{"integration_id", "event_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	Hooks        [4]ent.Hook
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Integration queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByKind orders the results by the kind field.
func ByKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKind, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// BySecretsCount orders the results by secrets count.
func BySecretsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSecretsStep(), opts...)
	}
}

// BySecrets orders the results by secrets terms.
func BySecrets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSecretsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOauth2tokensCount orders the results by oauth2tokens count.
func ByOauth2tokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOauth2tokensStep(), opts...)
	}
}

// ByOauth2tokens orders the results by oauth2tokens terms.
func ByOauth2tokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOauth2tokensStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newSecretsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SecretsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SecretsTable, SecretsPrimaryKey...),
	)
}
func newOauth2tokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Oauth2tokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, Oauth2tokensTable, Oauth2tokensPrimaryKey...),
	)
}
func newEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, EventsTable, EventsPrimaryKey...),
	)
}
