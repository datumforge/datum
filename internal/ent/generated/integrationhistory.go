// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/integrationhistory"
	"github.com/datumforge/enthistory"
)

// IntegrationHistory is the model entity for the IntegrationHistory schema.
type IntegrationHistory struct {
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
	// the name of the integration - must be unique within the organization
	Name string `json:"name,omitempty"`
	// a description of the integration
	Description string `json:"description,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// SecretName holds the value of the "secret_name" field.
	SecretName   string `json:"secret_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IntegrationHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case integrationhistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case integrationhistory.FieldID, integrationhistory.FieldRef, integrationhistory.FieldCreatedBy, integrationhistory.FieldUpdatedBy, integrationhistory.FieldDeletedBy, integrationhistory.FieldName, integrationhistory.FieldDescription, integrationhistory.FieldKind, integrationhistory.FieldSecretName:
			values[i] = new(sql.NullString)
		case integrationhistory.FieldHistoryTime, integrationhistory.FieldCreatedAt, integrationhistory.FieldUpdatedAt, integrationhistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IntegrationHistory fields.
func (ih *IntegrationHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case integrationhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ih.ID = value.String
			}
		case integrationhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ih.HistoryTime = value.Time
			}
		case integrationhistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				ih.Operation = *value
			}
		case integrationhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				ih.Ref = value.String
			}
		case integrationhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ih.CreatedAt = value.Time
			}
		case integrationhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ih.UpdatedAt = value.Time
			}
		case integrationhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ih.CreatedBy = value.String
			}
		case integrationhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ih.UpdatedBy = value.String
			}
		case integrationhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ih.DeletedAt = value.Time
			}
		case integrationhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				ih.DeletedBy = value.String
			}
		case integrationhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ih.Name = value.String
			}
		case integrationhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ih.Description = value.String
			}
		case integrationhistory.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				ih.Kind = value.String
			}
		case integrationhistory.FieldSecretName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_name", values[i])
			} else if value.Valid {
				ih.SecretName = value.String
			}
		default:
			ih.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IntegrationHistory.
// This includes values selected through modifiers, order, etc.
func (ih *IntegrationHistory) Value(name string) (ent.Value, error) {
	return ih.selectValues.Get(name)
}

// Update returns a builder for updating this IntegrationHistory.
// Note that you need to call IntegrationHistory.Unwrap() before calling this method if this IntegrationHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ih *IntegrationHistory) Update() *IntegrationHistoryUpdateOne {
	return NewIntegrationHistoryClient(ih.config).UpdateOne(ih)
}

// Unwrap unwraps the IntegrationHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ih *IntegrationHistory) Unwrap() *IntegrationHistory {
	_tx, ok := ih.config.driver.(*txDriver)
	if !ok {
		panic("generated: IntegrationHistory is not a transactional entity")
	}
	ih.config.driver = _tx.drv
	return ih
}

// String implements the fmt.Stringer.
func (ih *IntegrationHistory) String() string {
	var builder strings.Builder
	builder.WriteString("IntegrationHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ih.ID))
	builder.WriteString("history_time=")
	builder.WriteString(ih.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ih.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(ih.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ih.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ih.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(ih.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(ih.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ih.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(ih.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ih.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ih.Description)
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(ih.Kind)
	builder.WriteString(", ")
	builder.WriteString("secret_name=")
	builder.WriteString(ih.SecretName)
	builder.WriteByte(')')
	return builder.String()
}

// IntegrationHistories is a parsable slice of IntegrationHistory.
type IntegrationHistories []*IntegrationHistory
