// Code generated by ent, DO NOT EDIT.

package organizationsettinghistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/enthistory"
)

const (
	// Label holds the string label denoting the organizationsettinghistory type in the database.
	Label = "organization_setting_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHistoryTime holds the string denoting the history_time field in the database.
	FieldHistoryTime = "history_time"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
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
	// FieldSSOCert holds the string denoting the sso_cert field in the database.
	FieldSSOCert = "sso_cert"
	// FieldSSOEntrypoint holds the string denoting the sso_entrypoint field in the database.
	FieldSSOEntrypoint = "sso_entrypoint"
	// FieldSSOIssuer holds the string denoting the sso_issuer field in the database.
	FieldSSOIssuer = "sso_issuer"
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
	// Table holds the table name of the organizationsettinghistory in the database.
	Table = "organization_setting_history"
)

// Columns holds all SQL columns for organizationsettinghistory fields.
var Columns = []string{
	FieldID,
	FieldHistoryTime,
	FieldOperation,
	FieldRef,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldDomains,
	FieldSSOCert,
	FieldSSOEntrypoint,
	FieldSSOIssuer,
	FieldBillingContact,
	FieldBillingEmail,
	FieldBillingPhone,
	FieldBillingAddress,
	FieldTaxIdentifier,
	FieldTags,
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

var (
	// DefaultHistoryTime holds the default value on creation for the "history_time" field.
	DefaultHistoryTime func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("organizationsettinghistory: invalid enum value for operation field: %q", o)
	}
}

// OrderOption defines the ordering options for the OrganizationSettingHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHistoryTime orders the results by the history_time field.
func ByHistoryTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHistoryTime, opts...).ToFunc()
}

// ByOperation orders the results by the operation field.
func ByOperation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperation, opts...).ToFunc()
}

// ByRef orders the results by the ref field.
func ByRef(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRef, opts...).ToFunc()
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

// BySSOCert orders the results by the sso_cert field.
func BySSOCert(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSSOCert, opts...).ToFunc()
}

// BySSOEntrypoint orders the results by the sso_entrypoint field.
func BySSOEntrypoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSSOEntrypoint, opts...).ToFunc()
}

// BySSOIssuer orders the results by the sso_issuer field.
func BySSOIssuer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSSOIssuer, opts...).ToFunc()
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

var (
	// enthistory.OpType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enthistory.OpType)(nil)
	// enthistory.OpType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enthistory.OpType)(nil)
)
