// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/entityhistory"
	"github.com/datumforge/datum/pkg/enums"
	"github.com/datumforge/enthistory"
)

// EntityHistory is the model entity for the EntityHistory schema.
type EntityHistory struct {
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
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// the name of the entity
	Name string `json:"name,omitempty"`
	// The entity's displayed 'friendly' name
	DisplayName string `json:"display_name,omitempty"`
	// An optional description of the entity
	Description string `json:"description,omitempty"`
	// the type of the entity
	EntityType   enums.EntityType `json:"entity_type,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntityHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entityhistory.FieldTags:
			values[i] = new([]byte)
		case entityhistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case entityhistory.FieldID, entityhistory.FieldRef, entityhistory.FieldCreatedBy, entityhistory.FieldUpdatedBy, entityhistory.FieldMappingID, entityhistory.FieldDeletedBy, entityhistory.FieldOwnerID, entityhistory.FieldName, entityhistory.FieldDisplayName, entityhistory.FieldDescription, entityhistory.FieldEntityType:
			values[i] = new(sql.NullString)
		case entityhistory.FieldHistoryTime, entityhistory.FieldCreatedAt, entityhistory.FieldUpdatedAt, entityhistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntityHistory fields.
func (eh *EntityHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entityhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				eh.ID = value.String
			}
		case entityhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				eh.HistoryTime = value.Time
			}
		case entityhistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				eh.Operation = *value
			}
		case entityhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				eh.Ref = value.String
			}
		case entityhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				eh.CreatedAt = value.Time
			}
		case entityhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				eh.UpdatedAt = value.Time
			}
		case entityhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				eh.CreatedBy = value.String
			}
		case entityhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				eh.UpdatedBy = value.String
			}
		case entityhistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				eh.MappingID = value.String
			}
		case entityhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				eh.DeletedAt = value.Time
			}
		case entityhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				eh.DeletedBy = value.String
			}
		case entityhistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &eh.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case entityhistory.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				eh.OwnerID = value.String
			}
		case entityhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				eh.Name = value.String
			}
		case entityhistory.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				eh.DisplayName = value.String
			}
		case entityhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				eh.Description = value.String
			}
		case entityhistory.FieldEntityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entity_type", values[i])
			} else if value.Valid {
				eh.EntityType = enums.EntityType(value.String)
			}
		default:
			eh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EntityHistory.
// This includes values selected through modifiers, order, etc.
func (eh *EntityHistory) Value(name string) (ent.Value, error) {
	return eh.selectValues.Get(name)
}

// Update returns a builder for updating this EntityHistory.
// Note that you need to call EntityHistory.Unwrap() before calling this method if this EntityHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (eh *EntityHistory) Update() *EntityHistoryUpdateOne {
	return NewEntityHistoryClient(eh.config).UpdateOne(eh)
}

// Unwrap unwraps the EntityHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eh *EntityHistory) Unwrap() *EntityHistory {
	_tx, ok := eh.config.driver.(*txDriver)
	if !ok {
		panic("generated: EntityHistory is not a transactional entity")
	}
	eh.config.driver = _tx.drv
	return eh
}

// String implements the fmt.Stringer.
func (eh *EntityHistory) String() string {
	var builder strings.Builder
	builder.WriteString("EntityHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", eh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(eh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", eh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(eh.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(eh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(eh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(eh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(eh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(eh.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(eh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(eh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", eh.Tags))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(eh.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(eh.Name)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(eh.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(eh.Description)
	builder.WriteString(", ")
	builder.WriteString("entity_type=")
	builder.WriteString(fmt.Sprintf("%v", eh.EntityType))
	builder.WriteByte(')')
	return builder.String()
}

// EntityHistories is a parsable slice of EntityHistory.
type EntityHistories []*EntityHistory
