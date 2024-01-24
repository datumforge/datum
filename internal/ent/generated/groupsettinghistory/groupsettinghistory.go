// Code generated by ent, DO NOT EDIT.

package groupsettinghistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/flume/enthistory"
)

const (
	// Label holds the string label denoting the groupsettinghistory type in the database.
	Label = "group_setting_history"
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
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldJoinPolicy holds the string denoting the join_policy field in the database.
	FieldJoinPolicy = "join_policy"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldSyncToSlack holds the string denoting the sync_to_slack field in the database.
	FieldSyncToSlack = "sync_to_slack"
	// FieldSyncToGithub holds the string denoting the sync_to_github field in the database.
	FieldSyncToGithub = "sync_to_github"
	// Table holds the table name of the groupsettinghistory in the database.
	Table = "group_setting_history"
)

// Columns holds all SQL columns for groupsettinghistory fields.
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
	FieldVisibility,
	FieldJoinPolicy,
	FieldTags,
	FieldSyncToSlack,
	FieldSyncToGithub,
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
	// DefaultSyncToSlack holds the default value on creation for the "sync_to_slack" field.
	DefaultSyncToSlack bool
	// DefaultSyncToGithub holds the default value on creation for the "sync_to_github" field.
	DefaultSyncToGithub bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("groupsettinghistory: invalid enum value for operation field: %q", o)
	}
}

const DefaultVisibility enums.Visibility = "PUBLIC"

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v enums.Visibility) error {
	switch v.String() {
	case "PUBLIC", "PRIVATE":
		return nil
	default:
		return fmt.Errorf("groupsettinghistory: invalid enum value for visibility field: %q", v)
	}
}

const DefaultJoinPolicy enums.JoinPolicy = "INVITE_OR_APPLICATION"

// JoinPolicyValidator is a validator for the "join_policy" field enum values. It is called by the builders before save.
func JoinPolicyValidator(jp enums.JoinPolicy) error {
	switch jp.String() {
	case "OPEN", "INVITE_ONLY", "APPLICATION_ONLY", "INVITE_OR_APPLICATION":
		return nil
	default:
		return fmt.Errorf("groupsettinghistory: invalid enum value for join_policy field: %q", jp)
	}
}

// OrderOption defines the ordering options for the GroupSettingHistory queries.
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

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByJoinPolicy orders the results by the join_policy field.
func ByJoinPolicy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJoinPolicy, opts...).ToFunc()
}

// BySyncToSlack orders the results by the sync_to_slack field.
func BySyncToSlack(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSyncToSlack, opts...).ToFunc()
}

// BySyncToGithub orders the results by the sync_to_github field.
func BySyncToGithub(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSyncToGithub, opts...).ToFunc()
}

var (
	// enthistory.OpType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enthistory.OpType)(nil)
	// enthistory.OpType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enthistory.OpType)(nil)
)

var (
	// enums.Visibility must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.Visibility)(nil)
	// enums.Visibility must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.Visibility)(nil)
)

var (
	// enums.JoinPolicy must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enums.JoinPolicy)(nil)
	// enums.JoinPolicy must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.JoinPolicy)(nil)
)
