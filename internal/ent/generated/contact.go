// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/contact"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/pkg/enums"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
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
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// The organization id that owns the object
	OwnerID string `json:"owner_id,omitempty"`
	// the full name of the contact
	FullName string `json:"full_name,omitempty"`
	// the title of the contact
	Title string `json:"title,omitempty"`
	// the company of the contact
	Company string `json:"company,omitempty"`
	// the email of the contact
	Email string `json:"email,omitempty"`
	// the phone number of the contact
	PhoneNumber string `json:"phone_number,omitempty"`
	// the address of the contact
	Address string `json:"address,omitempty"`
	// status of the contact
	Status enums.UserStatus `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContactQuery when eager-loading is set.
	Edges        ContactEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContactEdges holds the relations/edges for other nodes in the graph.
type ContactEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// Entities holds the value of the entities edge.
	Entities []*Entity `json:"entities,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedEntities map[string][]*Entity
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// EntitiesOrErr returns the Entities value or an error if the edge
// was not loaded in eager-loading.
func (e ContactEdges) EntitiesOrErr() ([]*Entity, error) {
	if e.loadedTypes[1] {
		return e.Entities, nil
	}
	return nil, &NotLoadedError{edge: "entities"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldTags:
			values[i] = new([]byte)
		case contact.FieldID, contact.FieldCreatedBy, contact.FieldUpdatedBy, contact.FieldMappingID, contact.FieldDeletedBy, contact.FieldOwnerID, contact.FieldFullName, contact.FieldTitle, contact.FieldCompany, contact.FieldEmail, contact.FieldPhoneNumber, contact.FieldAddress, contact.FieldStatus:
			values[i] = new(sql.NullString)
		case contact.FieldCreatedAt, contact.FieldUpdatedAt, contact.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case contact.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case contact.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case contact.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = value.String
			}
		case contact.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = value.String
			}
		case contact.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				c.MappingID = value.String
			}
		case contact.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case contact.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				c.DeletedBy = value.String
			}
		case contact.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case contact.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				c.OwnerID = value.String
			}
		case contact.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				c.FullName = value.String
			}
		case contact.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case contact.FieldCompany:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company", values[i])
			} else if value.Valid {
				c.Company = value.String
			}
		case contact.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case contact.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				c.PhoneNumber = value.String
			}
		case contact.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		case contact.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = enums.UserStatus(value.String)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Contact.
// This includes values selected through modifiers, order, etc.
func (c *Contact) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Contact entity.
func (c *Contact) QueryOwner() *OrganizationQuery {
	return NewContactClient(c.config).QueryOwner(c)
}

// QueryEntities queries the "entities" edge of the Contact entity.
func (c *Contact) QueryEntities() *EntityQuery {
	return NewContactClient(c.config).QueryEntities(c)
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return NewContactClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("generated: Contact is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(c.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(c.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(c.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(c.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", c.Tags))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(c.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(c.FullName)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("company=")
	builder.WriteString(c.Company)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(c.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteByte(')')
	return builder.String()
}

// NamedEntities returns the Entities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Contact) NamedEntities(name string) ([]*Entity, error) {
	if c.Edges.namedEntities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedEntities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Contact) appendNamedEntities(name string, edges ...*Entity) {
	if c.Edges.namedEntities == nil {
		c.Edges.namedEntities = make(map[string][]*Entity)
	}
	if len(edges) == 0 {
		c.Edges.namedEntities[name] = []*Entity{}
	} else {
		c.Edges.namedEntities[name] = append(c.Edges.namedEntities[name], edges...)
	}
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact
