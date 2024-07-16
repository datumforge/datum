// Code generated by ent, DO NOT EDIT.

package tfasetting

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tfasetting type in the database.
	Label = "tfa_setting"
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
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldTfaSecret holds the string denoting the tfa_secret field in the database.
	FieldTfaSecret = "tfa_secret"
	// FieldVerified holds the string denoting the verified field in the database.
	FieldVerified = "verified"
	// FieldRecoveryCodes holds the string denoting the recovery_codes field in the database.
	FieldRecoveryCodes = "recovery_codes"
	// FieldPhoneOtpAllowed holds the string denoting the phone_otp_allowed field in the database.
	FieldPhoneOtpAllowed = "phone_otp_allowed"
	// FieldEmailOtpAllowed holds the string denoting the email_otp_allowed field in the database.
	FieldEmailOtpAllowed = "email_otp_allowed"
	// FieldTotpAllowed holds the string denoting the totp_allowed field in the database.
	FieldTotpAllowed = "totp_allowed"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the tfasetting in the database.
	Table = "tfa_settings"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "tfa_settings"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
)

// Columns holds all SQL columns for tfasetting fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldMappingID,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldTags,
	FieldOwnerID,
	FieldTfaSecret,
	FieldVerified,
	FieldRecoveryCodes,
	FieldPhoneOtpAllowed,
	FieldEmailOtpAllowed,
	FieldTotpAllowed,
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
	Hooks        [4]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMappingID holds the default value on creation for the "mapping_id" field.
	DefaultMappingID func() string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultVerified holds the default value on creation for the "verified" field.
	DefaultVerified bool
	// DefaultPhoneOtpAllowed holds the default value on creation for the "phone_otp_allowed" field.
	DefaultPhoneOtpAllowed bool
	// DefaultEmailOtpAllowed holds the default value on creation for the "email_otp_allowed" field.
	DefaultEmailOtpAllowed bool
	// DefaultTotpAllowed holds the default value on creation for the "totp_allowed" field.
	DefaultTotpAllowed bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the TFASetting queries.
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

// ByTfaSecret orders the results by the tfa_secret field.
func ByTfaSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTfaSecret, opts...).ToFunc()
}

// ByVerified orders the results by the verified field.
func ByVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerified, opts...).ToFunc()
}

// ByPhoneOtpAllowed orders the results by the phone_otp_allowed field.
func ByPhoneOtpAllowed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneOtpAllowed, opts...).ToFunc()
}

// ByEmailOtpAllowed orders the results by the email_otp_allowed field.
func ByEmailOtpAllowed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailOtpAllowed, opts...).ToFunc()
}

// ByTotpAllowed orders the results by the totp_allowed field.
func ByTotpAllowed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotpAllowed, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
