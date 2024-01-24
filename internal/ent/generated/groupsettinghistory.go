// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/groupsettinghistory"
	"github.com/flume/enthistory"
)

// GroupSettingHistory is the model entity for the GroupSettingHistory schema.
type GroupSettingHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// whether the group is visible to it's members / owners only or if it's searchable by anyone within the organization
	Visibility enums.Visibility `json:"visibility,omitempty"`
	// the policy governing ability to freely join a group, whether it requires an invitation, application, or either
	JoinPolicy enums.JoinPolicy `json:"join_policy,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// SyncToSlack holds the value of the "sync_to_slack" field.
	SyncToSlack bool `json:"sync_to_slack,omitempty"`
	// SyncToGithub holds the value of the "sync_to_github" field.
	SyncToGithub bool `json:"sync_to_github,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupSettingHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupsettinghistory.FieldTags:
			values[i] = new([]byte)
		case groupsettinghistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case groupsettinghistory.FieldSyncToSlack, groupsettinghistory.FieldSyncToGithub:
			values[i] = new(sql.NullBool)
		case groupsettinghistory.FieldID, groupsettinghistory.FieldRef, groupsettinghistory.FieldCreatedBy, groupsettinghistory.FieldUpdatedBy, groupsettinghistory.FieldDeletedBy, groupsettinghistory.FieldVisibility, groupsettinghistory.FieldJoinPolicy:
			values[i] = new(sql.NullString)
		case groupsettinghistory.FieldHistoryTime, groupsettinghistory.FieldCreatedAt, groupsettinghistory.FieldUpdatedAt, groupsettinghistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupSettingHistory fields.
func (gsh *GroupSettingHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupsettinghistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gsh.ID = value.String
			}
		case groupsettinghistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				gsh.HistoryTime = value.Time
			}
		case groupsettinghistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				gsh.Operation = *value
			}
		case groupsettinghistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				gsh.Ref = value.String
			}
		case groupsettinghistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gsh.CreatedAt = value.Time
			}
		case groupsettinghistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gsh.UpdatedAt = value.Time
			}
		case groupsettinghistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gsh.CreatedBy = value.String
			}
		case groupsettinghistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gsh.UpdatedBy = value.String
			}
		case groupsettinghistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gsh.DeletedAt = value.Time
			}
		case groupsettinghistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				gsh.DeletedBy = value.String
			}
		case groupsettinghistory.FieldVisibility:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				gsh.Visibility = enums.Visibility(value.String)
			}
		case groupsettinghistory.FieldJoinPolicy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field join_policy", values[i])
			} else if value.Valid {
				gsh.JoinPolicy = enums.JoinPolicy(value.String)
			}
		case groupsettinghistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gsh.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case groupsettinghistory.FieldSyncToSlack:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sync_to_slack", values[i])
			} else if value.Valid {
				gsh.SyncToSlack = value.Bool
			}
		case groupsettinghistory.FieldSyncToGithub:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sync_to_github", values[i])
			} else if value.Valid {
				gsh.SyncToGithub = value.Bool
			}
		default:
			gsh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupSettingHistory.
// This includes values selected through modifiers, order, etc.
func (gsh *GroupSettingHistory) Value(name string) (ent.Value, error) {
	return gsh.selectValues.Get(name)
}

// Update returns a builder for updating this GroupSettingHistory.
// Note that you need to call GroupSettingHistory.Unwrap() before calling this method if this GroupSettingHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (gsh *GroupSettingHistory) Update() *GroupSettingHistoryUpdateOne {
	return NewGroupSettingHistoryClient(gsh.config).UpdateOne(gsh)
}

// Unwrap unwraps the GroupSettingHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gsh *GroupSettingHistory) Unwrap() *GroupSettingHistory {
	_tx, ok := gsh.config.driver.(*txDriver)
	if !ok {
		panic("generated: GroupSettingHistory is not a transactional entity")
	}
	gsh.config.driver = _tx.drv
	return gsh
}

// String implements the fmt.Stringer.
func (gsh *GroupSettingHistory) String() string {
	var builder strings.Builder
	builder.WriteString("GroupSettingHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gsh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(gsh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", gsh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(gsh.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gsh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gsh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(gsh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(gsh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gsh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(gsh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", gsh.Visibility))
	builder.WriteString(", ")
	builder.WriteString("join_policy=")
	builder.WriteString(fmt.Sprintf("%v", gsh.JoinPolicy))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", gsh.Tags))
	builder.WriteString(", ")
	builder.WriteString("sync_to_slack=")
	builder.WriteString(fmt.Sprintf("%v", gsh.SyncToSlack))
	builder.WriteString(", ")
	builder.WriteString("sync_to_github=")
	builder.WriteString(fmt.Sprintf("%v", gsh.SyncToGithub))
	builder.WriteByte(')')
	return builder.String()
}

// GroupSettingHistories is a parsable slice of GroupSettingHistory.
type GroupSettingHistories []*GroupSettingHistory
