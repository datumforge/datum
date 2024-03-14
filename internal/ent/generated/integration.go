// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
)

// Integration is the model entity for the Integration schema.
type Integration struct {
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
	// the name of the integration - must be unique within the organization
	Name string `json:"name,omitempty"`
	// a description of the integration
	Description string `json:"description,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// SecretName holds the value of the "secret_name" field.
	SecretName string `json:"secret_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IntegrationQuery when eager-loading is set.
	Edges        IntegrationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// IntegrationEdges holds the relations/edges for other nodes in the graph.
type IntegrationEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IntegrationEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Integration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case integration.FieldID, integration.FieldCreatedBy, integration.FieldUpdatedBy, integration.FieldDeletedBy, integration.FieldOwnerID, integration.FieldName, integration.FieldDescription, integration.FieldKind, integration.FieldSecretName:
			values[i] = new(sql.NullString)
		case integration.FieldCreatedAt, integration.FieldUpdatedAt, integration.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Integration fields.
func (i *Integration) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case integration.FieldID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value.Valid {
				i.ID = value.String
			}
		case integration.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case integration.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case integration.FieldCreatedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[j])
			} else if value.Valid {
				i.CreatedBy = value.String
			}
		case integration.FieldUpdatedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[j])
			} else if value.Valid {
				i.UpdatedBy = value.String
			}
		case integration.FieldDeletedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[j])
			} else if value.Valid {
				i.DeletedAt = value.Time
			}
		case integration.FieldDeletedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[j])
			} else if value.Valid {
				i.DeletedBy = value.String
			}
		case integration.FieldOwnerID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[j])
			} else if value.Valid {
				i.OwnerID = value.String
			}
		case integration.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case integration.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case integration.FieldKind:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[j])
			} else if value.Valid {
				i.Kind = value.String
			}
		case integration.FieldSecretName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_name", values[j])
			} else if value.Valid {
				i.SecretName = value.String
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Integration.
// This includes values selected through modifiers, order, etc.
func (i *Integration) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Integration entity.
func (i *Integration) QueryOwner() *OrganizationQuery {
	return NewIntegrationClient(i.config).QueryOwner(i)
}

// Update returns a builder for updating this Integration.
// Note that you need to call Integration.Unwrap() before calling this method if this Integration
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Integration) Update() *IntegrationUpdateOne {
	return NewIntegrationClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Integration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Integration) Unwrap() *Integration {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("generated: Integration is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Integration) String() string {
	var builder strings.Builder
	builder.WriteString("Integration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(i.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(i.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(i.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(i.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(i.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(i.Kind)
	builder.WriteString(", ")
	builder.WriteString("secret_name=")
	builder.WriteString(i.SecretName)
	builder.WriteByte(')')
	return builder.String()
}

// Integrations is a parsable slice of Integration.
type Integrations []*Integration
