// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/generated/documentdatahistory"
	"github.com/datumforge/enthistory"
)

// DocumentDataHistory is the model entity for the DocumentDataHistory schema.
type DocumentDataHistory struct {
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
	// the template id of the document
	TemplateID string `json:"template_id,omitempty"`
	// the json data of the document
	Data         customtypes.JSONObject `json:"data,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DocumentDataHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case documentdatahistory.FieldData:
			values[i] = new([]byte)
		case documentdatahistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case documentdatahistory.FieldID, documentdatahistory.FieldRef, documentdatahistory.FieldCreatedBy, documentdatahistory.FieldUpdatedBy, documentdatahistory.FieldDeletedBy, documentdatahistory.FieldTemplateID:
			values[i] = new(sql.NullString)
		case documentdatahistory.FieldHistoryTime, documentdatahistory.FieldCreatedAt, documentdatahistory.FieldUpdatedAt, documentdatahistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DocumentDataHistory fields.
func (ddh *DocumentDataHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case documentdatahistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ddh.ID = value.String
			}
		case documentdatahistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ddh.HistoryTime = value.Time
			}
		case documentdatahistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				ddh.Operation = *value
			}
		case documentdatahistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				ddh.Ref = value.String
			}
		case documentdatahistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ddh.CreatedAt = value.Time
			}
		case documentdatahistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ddh.UpdatedAt = value.Time
			}
		case documentdatahistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ddh.CreatedBy = value.String
			}
		case documentdatahistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ddh.UpdatedBy = value.String
			}
		case documentdatahistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ddh.DeletedAt = value.Time
			}
		case documentdatahistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				ddh.DeletedBy = value.String
			}
		case documentdatahistory.FieldTemplateID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_id", values[i])
			} else if value.Valid {
				ddh.TemplateID = value.String
			}
		case documentdatahistory.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ddh.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		default:
			ddh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DocumentDataHistory.
// This includes values selected through modifiers, order, etc.
func (ddh *DocumentDataHistory) Value(name string) (ent.Value, error) {
	return ddh.selectValues.Get(name)
}

// Update returns a builder for updating this DocumentDataHistory.
// Note that you need to call DocumentDataHistory.Unwrap() before calling this method if this DocumentDataHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ddh *DocumentDataHistory) Update() *DocumentDataHistoryUpdateOne {
	return NewDocumentDataHistoryClient(ddh.config).UpdateOne(ddh)
}

// Unwrap unwraps the DocumentDataHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ddh *DocumentDataHistory) Unwrap() *DocumentDataHistory {
	_tx, ok := ddh.config.driver.(*txDriver)
	if !ok {
		panic("generated: DocumentDataHistory is not a transactional entity")
	}
	ddh.config.driver = _tx.drv
	return ddh
}

// String implements the fmt.Stringer.
func (ddh *DocumentDataHistory) String() string {
	var builder strings.Builder
	builder.WriteString("DocumentDataHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ddh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(ddh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ddh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(ddh.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ddh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ddh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(ddh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(ddh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ddh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(ddh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("template_id=")
	builder.WriteString(ddh.TemplateID)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", ddh.Data))
	builder.WriteByte(')')
	return builder.String()
}

// DocumentDataHistories is a parsable slice of DocumentDataHistory.
type DocumentDataHistories []*DocumentDataHistory
