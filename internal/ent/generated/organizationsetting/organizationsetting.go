// Code generated by ent, DO NOT EDIT.

package organizationsetting

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the organizationsetting type in the database.
	Label = "organization_setting"
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
	// FieldDomains holds the string denoting the domains field in the database.
	FieldDomains = "domains"
	// FieldBillingContact holds the string denoting the billing_contact field in the database.
	FieldBillingContact = "billing_contact"
	// FieldBillingEmail holds the string denoting the billing_email field in the database.
	FieldBillingEmail = "billing_email"
	// FieldBillingPhone holds the string denoting the billing_phone field in the database.
	FieldBillingPhone = "billing_phone"
	// FieldBillingAddress holds the string denoting the billing_address field in the database.
	FieldBillingAddress = "billing_address"
	// FieldTaxIdentifier holds the string denoting the tax_identifier field in the database.
	FieldTaxIdentifier = "tax_identifier"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldAvatarRemoteURL holds the string denoting the avatar_remote_url field in the database.
	FieldAvatarRemoteURL = "avatar_remote_url"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// Table holds the table name of the organizationsetting in the database.
	Table = "organization_settings"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "organization_settings"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_setting"
)

// Columns holds all SQL columns for organizationsetting fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldDomains,
	FieldBillingContact,
	FieldBillingEmail,
	FieldBillingPhone,
	FieldBillingAddress,
	FieldTaxIdentifier,
	FieldTags,
	FieldAvatarRemoteURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "organization_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"organization_setting",
}

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
	Interceptors [1]ent.Interceptor
	Policy       ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// BillingEmailValidator is a validator for the "billing_email" field. It is called by the builders before save.
	BillingEmailValidator func(string) error
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// AvatarRemoteURLValidator is a validator for the "avatar_remote_url" field. It is called by the builders before save.
	AvatarRemoteURLValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the OrganizationSetting queries.
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

// ByBillingContact orders the results by the billing_contact field.
func ByBillingContact(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingContact, opts...).ToFunc()
}

// ByBillingEmail orders the results by the billing_email field.
func ByBillingEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingEmail, opts...).ToFunc()
}

// ByBillingPhone orders the results by the billing_phone field.
func ByBillingPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingPhone, opts...).ToFunc()
}

// ByBillingAddress orders the results by the billing_address field.
func ByBillingAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingAddress, opts...).ToFunc()
}

// ByTaxIdentifier orders the results by the tax_identifier field.
func ByTaxIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxIdentifier, opts...).ToFunc()
}

// ByAvatarRemoteURL orders the results by the avatar_remote_url field.
func ByAvatarRemoteURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarRemoteURL, opts...).ToFunc()
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OrganizationTable, OrganizationColumn),
	)
}
