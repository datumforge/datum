// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
)

// GroupSettings is the model entity for the GroupSettings schema.
type GroupSettings struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Visibility holds the value of the "visibility" field.
	Visibility groupsettings.Visibility `json:"visibility,omitempty"`
	// JoinPolicy holds the value of the "join_policy" field.
	JoinPolicy groupsettings.JoinPolicy `json:"join_policy,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupSettingsQuery when eager-loading is set.
	Edges         GroupSettingsEdges `json:"edges"`
	group_setting *string
	selectValues  sql.SelectValues
}

// GroupSettingsEdges holds the relations/edges for other nodes in the graph.
type GroupSettingsEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupSettingsEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupSettings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupsettings.FieldTags:
			values[i] = new([]byte)
		case groupsettings.FieldID, groupsettings.FieldCreatedBy, groupsettings.FieldUpdatedBy, groupsettings.FieldVisibility, groupsettings.FieldJoinPolicy:
			values[i] = new(sql.NullString)
		case groupsettings.FieldCreatedAt, groupsettings.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case groupsettings.ForeignKeys[0]: // group_setting
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupSettings fields.
func (gs *GroupSettings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupsettings.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gs.ID = value.String
			}
		case groupsettings.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gs.CreatedAt = value.Time
			}
		case groupsettings.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gs.UpdatedAt = value.Time
			}
		case groupsettings.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gs.CreatedBy = value.String
			}
		case groupsettings.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gs.UpdatedBy = value.String
			}
		case groupsettings.FieldVisibility:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				gs.Visibility = groupsettings.Visibility(value.String)
			}
		case groupsettings.FieldJoinPolicy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field join_policy", values[i])
			} else if value.Valid {
				gs.JoinPolicy = groupsettings.JoinPolicy(value.String)
			}
		case groupsettings.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gs.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case groupsettings.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_setting", values[i])
			} else if value.Valid {
				gs.group_setting = new(string)
				*gs.group_setting = value.String
			}
		default:
			gs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupSettings.
// This includes values selected through modifiers, order, etc.
func (gs *GroupSettings) Value(name string) (ent.Value, error) {
	return gs.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the GroupSettings entity.
func (gs *GroupSettings) QueryGroup() *GroupQuery {
	return NewGroupSettingsClient(gs.config).QueryGroup(gs)
}

// Update returns a builder for updating this GroupSettings.
// Note that you need to call GroupSettings.Unwrap() before calling this method if this GroupSettings
// was returned from a transaction, and the transaction was committed or rolled back.
func (gs *GroupSettings) Update() *GroupSettingsUpdateOne {
	return NewGroupSettingsClient(gs.config).UpdateOne(gs)
}

// Unwrap unwraps the GroupSettings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gs *GroupSettings) Unwrap() *GroupSettings {
	_tx, ok := gs.config.driver.(*txDriver)
	if !ok {
		panic("generated: GroupSettings is not a transactional entity")
	}
	gs.config.driver = _tx.drv
	return gs
}

// String implements the fmt.Stringer.
func (gs *GroupSettings) String() string {
	var builder strings.Builder
	builder.WriteString("GroupSettings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gs.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(gs.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(gs.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", gs.Visibility))
	builder.WriteString(", ")
	builder.WriteString("join_policy=")
	builder.WriteString(fmt.Sprintf("%v", gs.JoinPolicy))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", gs.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// GroupSettingsSlice is a parsable slice of GroupSettings.
type GroupSettingsSlice []*GroupSettings
