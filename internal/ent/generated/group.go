// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
	"github.com/datumforge/datum/internal/ent/generated/organization"
)

// Group is the model entity for the Group schema.
type Group struct {
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
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// the name of the group - must be unique within the organization
	Name string `json:"name,omitempty"`
	// the groups description
	Description string `json:"description,omitempty"`
	// the URL to an auto generated gravatar image for the group
	GravatarLogoURL string `json:"gravatar_logo_url,omitempty"`
	// the URL to an image uploaded by the customer for the groups avatar image
	LogoURL string `json:"logo_url,omitempty"`
	// The group's displayed 'friendly' name
	DisplayName string `json:"display_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges        GroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// Setting holds the value of the setting edge.
	Setting *GroupSetting `json:"setting,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Members holds the value of the members edge.
	Members []*GroupMembership `json:"members,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedUsers   map[string][]*User
	namedMembers map[string][]*GroupMembership
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// SettingOrErr returns the Setting value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) SettingOrErr() (*GroupSetting, error) {
	if e.Setting != nil {
		return e.Setting, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: groupsetting.Label}
	}
	return nil, &NotLoadedError{edge: "setting"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) MembersOrErr() ([]*GroupMembership, error) {
	if e.loadedTypes[3] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldID, group.FieldCreatedBy, group.FieldUpdatedBy, group.FieldDeletedBy, group.FieldOwnerID, group.FieldName, group.FieldDescription, group.FieldGravatarLogoURL, group.FieldLogoURL, group.FieldDisplayName:
			values[i] = new(sql.NullString)
		case group.FieldCreatedAt, group.FieldUpdatedAt, group.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gr.ID = value.String
			}
		case group.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case group.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gr.UpdatedAt = value.Time
			}
		case group.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gr.CreatedBy = value.String
			}
		case group.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gr.UpdatedBy = value.String
			}
		case group.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gr.DeletedAt = value.Time
			}
		case group.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				gr.DeletedBy = value.String
			}
		case group.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				gr.OwnerID = value.String
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gr.Description = value.String
			}
		case group.FieldGravatarLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gravatar_logo_url", values[i])
			} else if value.Valid {
				gr.GravatarLogoURL = value.String
			}
		case group.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				gr.LogoURL = value.String
			}
		case group.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				gr.DisplayName = value.String
			}
		default:
			gr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (gr *Group) Value(name string) (ent.Value, error) {
	return gr.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Group entity.
func (gr *Group) QueryOwner() *OrganizationQuery {
	return NewGroupClient(gr.config).QueryOwner(gr)
}

// QuerySetting queries the "setting" edge of the Group entity.
func (gr *Group) QuerySetting() *GroupSettingQuery {
	return NewGroupClient(gr.config).QuerySetting(gr)
}

// QueryUsers queries the "users" edge of the Group entity.
func (gr *Group) QueryUsers() *UserQuery {
	return NewGroupClient(gr.config).QueryUsers(gr)
}

// QueryMembers queries the "members" edge of the Group entity.
func (gr *Group) QueryMembers() *GroupMembershipQuery {
	return NewGroupClient(gr.config).QueryMembers(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("generated: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(gr.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(gr.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(gr.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(gr.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(gr.Description)
	builder.WriteString(", ")
	builder.WriteString("gravatar_logo_url=")
	builder.WriteString(gr.GravatarLogoURL)
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(gr.LogoURL)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(gr.DisplayName)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (gr *Group) NamedUsers(name string) ([]*User, error) {
	if gr.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := gr.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (gr *Group) appendNamedUsers(name string, edges ...*User) {
	if gr.Edges.namedUsers == nil {
		gr.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		gr.Edges.namedUsers[name] = []*User{}
	} else {
		gr.Edges.namedUsers[name] = append(gr.Edges.namedUsers[name], edges...)
	}
}

// NamedMembers returns the Members named value or an error if the edge was not
// loaded in eager-loading with this name.
func (gr *Group) NamedMembers(name string) ([]*GroupMembership, error) {
	if gr.Edges.namedMembers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := gr.Edges.namedMembers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (gr *Group) appendNamedMembers(name string, edges ...*GroupMembership) {
	if gr.Edges.namedMembers == nil {
		gr.Edges.namedMembers = make(map[string][]*GroupMembership)
	}
	if len(edges) == 0 {
		gr.Edges.namedMembers[name] = []*GroupMembership{}
	} else {
		gr.Edges.namedMembers[name] = append(gr.Edges.namedMembers[name], edges...)
	}
}

// Groups is a parsable slice of Group.
type Groups []*Group
