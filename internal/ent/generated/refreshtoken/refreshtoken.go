// Code generated by ent, DO NOT EDIT.

package refreshtoken

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the refreshtoken type in the database.
	Label = "refresh_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRefreshToken holds the string denoting the refresh_token field in the database.
	FieldRefreshToken = "refresh_token"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldIssuedAt holds the string denoting the issued_at field in the database.
	FieldIssuedAt = "issued_at"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// Table holds the table name of the refreshtoken in the database.
	Table = "refresh_tokens"
)

// Columns holds all SQL columns for refreshtoken fields.
var Columns = []string{
	FieldID,
	FieldRefreshToken,
	FieldExpiresAt,
	FieldIssuedAt,
	FieldOrganizationID,
	FieldUserID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "refresh_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_refresh_token",
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
	Hooks [1]ent.Hook
	// DefaultExpiresAt holds the default value on creation for the "expires_at" field.
	DefaultExpiresAt time.Time
	// DefaultIssuedAt holds the default value on creation for the "issued_at" field.
	DefaultIssuedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the RefreshToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRefreshToken orders the results by the refresh_token field.
func ByRefreshToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshToken, opts...).ToFunc()
}

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}

// ByIssuedAt orders the results by the issued_at field.
func ByIssuedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIssuedAt, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}
