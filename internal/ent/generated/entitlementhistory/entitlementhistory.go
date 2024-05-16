// Code generated by ent, DO NOT EDIT.

package entitlementhistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/enthistory"
)

const (
	// Label holds the string label denoting the entitlementhistory type in the database.
	Label = "entitlement_history"
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
	// FieldMappingID holds the string denoting the mapping_id field in the database.
	FieldMappingID = "mapping_id"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldTier holds the string denoting the tier field in the database.
	FieldTier = "tier"
	// FieldExternalCustomerID holds the string denoting the external_customer_id field in the database.
	FieldExternalCustomerID = "external_customer_id"
	// FieldExternalSubscriptionID holds the string denoting the external_subscription_id field in the database.
	FieldExternalSubscriptionID = "external_subscription_id"
	// FieldExpires holds the string denoting the expires field in the database.
	FieldExpires = "expires"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldCancelled holds the string denoting the cancelled field in the database.
	FieldCancelled = "cancelled"
	// Table holds the table name of the entitlementhistory in the database.
	Table = "entitlement_history"
)

// Columns holds all SQL columns for entitlementhistory fields.
var Columns = []string{
	FieldID,
	FieldHistoryTime,
	FieldOperation,
	FieldRef,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldMappingID,
	FieldTags,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldOwnerID,
	FieldTier,
	FieldExternalCustomerID,
	FieldExternalSubscriptionID,
	FieldExpires,
	FieldExpiresAt,
	FieldCancelled,
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
	// DefaultMappingID holds the default value on creation for the "mapping_id" field.
	DefaultMappingID func() string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultExpires holds the default value on creation for the "expires" field.
	DefaultExpires bool
	// DefaultCancelled holds the default value on creation for the "cancelled" field.
	DefaultCancelled bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("entitlementhistory: invalid enum value for operation field: %q", o)
	}
}

const DefaultTier enums.Tier = "FREE"

// TierValidator is a validator for the "tier" field enum values. It is called by the builders before save.
func TierValidator(t enums.Tier) error {
	switch t.String() {
	case "FREE", "PRO", "ENTERPRISE":
		return nil
	default:
		return fmt.Errorf("entitlementhistory: invalid enum value for tier field: %q", t)
	}
}

// OrderOption defines the ordering options for the EntitlementHistory queries.
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

// ByTier orders the results by the tier field.
func ByTier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTier, opts...).ToFunc()
}

// ByExternalCustomerID orders the results by the external_customer_id field.
func ByExternalCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalCustomerID, opts...).ToFunc()
}

// ByExternalSubscriptionID orders the results by the external_subscription_id field.
func ByExternalSubscriptionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalSubscriptionID, opts...).ToFunc()
}

// ByExpires orders the results by the expires field.
func ByExpires(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpires, opts...).ToFunc()
}

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}

// ByCancelled orders the results by the cancelled field.
func ByCancelled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCancelled, opts...).ToFunc()
}

var (
	// enthistory.OpType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enthistory.OpType)(nil)
	// enthistory.OpType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enthistory.OpType)(nil)
)

var (
	// enums.Tier must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.Tier)(nil)
	// enums.Tier must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.Tier)(nil)
)
