// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupmembership"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// GroupMembership is the model entity for the GroupMembership schema.
type GroupMembership struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// the time the record was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// the time the record was last updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// the user who created the record
	CreatedBy string `json:"created_by,omitempty"`
	// the user who last updated the record
	UpdatedBy string `json:"updated_by,omitempty"`
	// the time the record was deleted
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// the user who deleted the record
	DeletedBy string `json:"deleted_by,omitempty"`
	// Role holds the value of the "role" field.
	Role enums.Role `json:"role,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID string `json:"group_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupMembershipQuery when eager-loading is set.
	Edges        GroupMembershipEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupMembershipEdges holds the relations/edges for other nodes in the graph.
type GroupMembershipEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMembershipEdges) GroupOrErr() (*Group, error) {
	if e.Group != nil {
		return e.Group, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: group.Label}
	}
	return nil, &NotLoadedError{edge: "group"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMembershipEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupMembership) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupmembership.FieldID, groupmembership.FieldCreatedBy, groupmembership.FieldUpdatedBy, groupmembership.FieldDeletedBy, groupmembership.FieldRole, groupmembership.FieldGroupID, groupmembership.FieldUserID:
			values[i] = new(sql.NullString)
		case groupmembership.FieldCreatedAt, groupmembership.FieldUpdatedAt, groupmembership.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupMembership fields.
func (gm *GroupMembership) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupmembership.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gm.ID = value.String
			}
		case groupmembership.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = value.Time
			}
		case groupmembership.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gm.UpdatedAt = value.Time
			}
		case groupmembership.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gm.CreatedBy = value.String
			}
		case groupmembership.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gm.UpdatedBy = value.String
			}
		case groupmembership.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gm.DeletedAt = value.Time
			}
		case groupmembership.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				gm.DeletedBy = value.String
			}
		case groupmembership.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				gm.Role = enums.Role(value.String)
			}
		case groupmembership.FieldGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gm.GroupID = value.String
			}
		case groupmembership.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				gm.UserID = value.String
			}
		default:
			gm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupMembership.
// This includes values selected through modifiers, order, etc.
func (gm *GroupMembership) Value(name string) (ent.Value, error) {
	return gm.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the GroupMembership entity.
func (gm *GroupMembership) QueryGroup() *GroupQuery {
	return NewGroupMembershipClient(gm.config).QueryGroup(gm)
}

// QueryUser queries the "user" edge of the GroupMembership entity.
func (gm *GroupMembership) QueryUser() *UserQuery {
	return NewGroupMembershipClient(gm.config).QueryUser(gm)
}

// Update returns a builder for updating this GroupMembership.
// Note that you need to call GroupMembership.Unwrap() before calling this method if this GroupMembership
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GroupMembership) Update() *GroupMembershipUpdateOne {
	return NewGroupMembershipClient(gm.config).UpdateOne(gm)
}

// Unwrap unwraps the GroupMembership entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GroupMembership) Unwrap() *GroupMembership {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("generated: GroupMembership is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GroupMembership) String() string {
	var builder strings.Builder
	builder.WriteString("GroupMembership(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(gm.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(gm.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gm.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(gm.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", gm.Role))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(gm.GroupID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(gm.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// GroupMemberships is a parsable slice of GroupMembership.
type GroupMemberships []*GroupMembership
