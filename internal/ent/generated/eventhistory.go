// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/eventhistory"
	"github.com/datumforge/enthistory"
)

// EventHistory is the model entity for the EventHistory schema.
type EventHistory struct {
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
	// EventID holds the value of the "event_id" field.
	EventID string `json:"event_id,omitempty"`
	// CorrelationID holds the value of the "correlation_id" field.
	CorrelationID string `json:"correlation_id,omitempty"`
	// EventType holds the value of the "event_type" field.
	EventType string `json:"event_type,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EventHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case eventhistory.FieldMetadata:
			values[i] = new([]byte)
		case eventhistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case eventhistory.FieldID, eventhistory.FieldRef, eventhistory.FieldCreatedBy, eventhistory.FieldUpdatedBy, eventhistory.FieldMappingID, eventhistory.FieldEventID, eventhistory.FieldCorrelationID, eventhistory.FieldEventType:
			values[i] = new(sql.NullString)
		case eventhistory.FieldHistoryTime, eventhistory.FieldCreatedAt, eventhistory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EventHistory fields.
func (eh *EventHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case eventhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				eh.ID = value.String
			}
		case eventhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				eh.HistoryTime = value.Time
			}
		case eventhistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				eh.Operation = *value
			}
		case eventhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				eh.Ref = value.String
			}
		case eventhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				eh.CreatedAt = value.Time
			}
		case eventhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				eh.UpdatedAt = value.Time
			}
		case eventhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				eh.CreatedBy = value.String
			}
		case eventhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				eh.UpdatedBy = value.String
			}
		case eventhistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				eh.MappingID = value.String
			}
		case eventhistory.FieldEventID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_id", values[i])
			} else if value.Valid {
				eh.EventID = value.String
			}
		case eventhistory.FieldCorrelationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field correlation_id", values[i])
			} else if value.Valid {
				eh.CorrelationID = value.String
			}
		case eventhistory.FieldEventType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_type", values[i])
			} else if value.Valid {
				eh.EventType = value.String
			}
		case eventhistory.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &eh.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		default:
			eh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EventHistory.
// This includes values selected through modifiers, order, etc.
func (eh *EventHistory) Value(name string) (ent.Value, error) {
	return eh.selectValues.Get(name)
}

// Update returns a builder for updating this EventHistory.
// Note that you need to call EventHistory.Unwrap() before calling this method if this EventHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (eh *EventHistory) Update() *EventHistoryUpdateOne {
	return NewEventHistoryClient(eh.config).UpdateOne(eh)
}

// Unwrap unwraps the EventHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eh *EventHistory) Unwrap() *EventHistory {
	_tx, ok := eh.config.driver.(*txDriver)
	if !ok {
		panic("generated: EventHistory is not a transactional entity")
	}
	eh.config.driver = _tx.drv
	return eh
}

// String implements the fmt.Stringer.
func (eh *EventHistory) String() string {
	var builder strings.Builder
	builder.WriteString("EventHistory(")
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
	builder.WriteString("event_id=")
	builder.WriteString(eh.EventID)
	builder.WriteString(", ")
	builder.WriteString("correlation_id=")
	builder.WriteString(eh.CorrelationID)
	builder.WriteString(", ")
	builder.WriteString("event_type=")
	builder.WriteString(eh.EventType)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", eh.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// EventHistories is a parsable slice of EventHistory.
type EventHistories []*EventHistory
