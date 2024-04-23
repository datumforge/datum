// Code generated by ent, DO NOT EDIT.

package usersettinghistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/enthistory"
)

const (
	// Label holds the string label denoting the usersettinghistory type in the database.
	Label = "user_setting_history"
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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldSilencedAt holds the string denoting the silenced_at field in the database.
	FieldSilencedAt = "silenced_at"
	// FieldSuspendedAt holds the string denoting the suspended_at field in the database.
	FieldSuspendedAt = "suspended_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldEmailConfirmed holds the string denoting the email_confirmed field in the database.
	FieldEmailConfirmed = "email_confirmed"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldIsWebauthnAllowed holds the string denoting the is_webauthn_allowed field in the database.
	FieldIsWebauthnAllowed = "is_webauthn_allowed"
	// FieldIsTfaEnabled holds the string denoting the is_tfa_enabled field in the database.
	FieldIsTfaEnabled = "is_tfa_enabled"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// Table holds the table name of the usersettinghistory in the database.
	Table = "user_setting_history"
)

// Columns holds all SQL columns for usersettinghistory fields.
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
	FieldUserID,
	FieldLocked,
	FieldSilencedAt,
	FieldSuspendedAt,
	FieldStatus,
	FieldEmailConfirmed,
	FieldTags,
	FieldIsWebauthnAllowed,
	FieldIsTfaEnabled,
	FieldPhoneNumber,
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
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked bool
	// DefaultEmailConfirmed holds the default value on creation for the "email_confirmed" field.
	DefaultEmailConfirmed bool
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultIsWebauthnAllowed holds the default value on creation for the "is_webauthn_allowed" field.
	DefaultIsWebauthnAllowed bool
	// DefaultIsTfaEnabled holds the default value on creation for the "is_tfa_enabled" field.
	DefaultIsTfaEnabled bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("usersettinghistory: invalid enum value for operation field: %q", o)
	}
}

const DefaultStatus enums.UserStatus = "ACTIVE"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.UserStatus) error {
	switch s.String() {
	case "ACTIVE", "INACTIVE", "DEACTIVATED", "SUSPENDED":
		return nil
	default:
		return fmt.Errorf("usersettinghistory: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the UserSettingHistory queries.
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

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByLocked orders the results by the locked field.
func ByLocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocked, opts...).ToFunc()
}

// BySilencedAt orders the results by the silenced_at field.
func BySilencedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSilencedAt, opts...).ToFunc()
}

// BySuspendedAt orders the results by the suspended_at field.
func BySuspendedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuspendedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByEmailConfirmed orders the results by the email_confirmed field.
func ByEmailConfirmed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailConfirmed, opts...).ToFunc()
}

// ByIsWebauthnAllowed orders the results by the is_webauthn_allowed field.
func ByIsWebauthnAllowed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsWebauthnAllowed, opts...).ToFunc()
}

// ByIsTfaEnabled orders the results by the is_tfa_enabled field.
func ByIsTfaEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsTfaEnabled, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

var (
	// enthistory.OpType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enthistory.OpType)(nil)
	// enthistory.OpType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enthistory.OpType)(nil)
)

var (
	// enums.UserStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.UserStatus)(nil)
	// enums.UserStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.UserStatus)(nil)
)