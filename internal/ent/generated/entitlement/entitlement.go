// Code generated by ent, DO NOT EDIT.

package entitlement

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the entitlement type in the database.
	Label = "entitlement"
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
	// FieldTier holds the string denoting the tier field in the database.
	FieldTier = "tier"
	// FieldStripeCustomerID holds the string denoting the stripe_customer_id field in the database.
	FieldStripeCustomerID = "stripe_customer_id"
	// FieldStripeSubscriptionID holds the string denoting the stripe_subscription_id field in the database.
	FieldStripeSubscriptionID = "stripe_subscription_id"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldCancelled holds the string denoting the cancelled field in the database.
	FieldCancelled = "cancelled"
	// Table holds the table name of the entitlement in the database.
	Table = "entitlements"
)

// Columns holds all SQL columns for entitlement fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldTier,
	FieldStripeCustomerID,
	FieldStripeSubscriptionID,
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
	// DefaultCancelled holds the default value on creation for the "cancelled" field.
	DefaultCancelled bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// Tier defines the type for the "tier" enum field.
type Tier string

// TierFree is the default value of the Tier enum.
const DefaultTier = TierFree

// Tier values.
const (
	TierFree       Tier = "free"
	TierPro        Tier = "pro"
	TierEnterprise Tier = "enterprise"
)

func (t Tier) String() string {
	return string(t)
}

// TierValidator is a validator for the "tier" field enum values. It is called by the builders before save.
func TierValidator(t Tier) error {
	switch t {
	case TierFree, TierPro, TierEnterprise:
		return nil
	default:
		return fmt.Errorf("entitlement: invalid enum value for tier field: %q", t)
	}
}

// OrderOption defines the ordering options for the Entitlement queries.
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

// ByTier orders the results by the tier field.
func ByTier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTier, opts...).ToFunc()
}

// ByStripeCustomerID orders the results by the stripe_customer_id field.
func ByStripeCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStripeCustomerID, opts...).ToFunc()
}

// ByStripeSubscriptionID orders the results by the stripe_subscription_id field.
func ByStripeSubscriptionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStripeSubscriptionID, opts...).ToFunc()
}

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}

// ByCancelled orders the results by the cancelled field.
func ByCancelled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCancelled, opts...).ToFunc()
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Tier) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Tier) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Tier(str)
	if err := TierValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Tier", str)
	}
	return nil
}
