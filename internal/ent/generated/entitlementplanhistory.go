// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/entitlementplanhistory"
	"github.com/datumforge/enthistory"
)

// EntitlementPlanHistory is the model entity for the EntitlementPlanHistory schema.
type EntitlementPlanHistory struct {
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
	// the displayed 'friendly' name of the plan
	DisplayName string `json:"display_name,omitempty"`
	// the unique name of the plan
	Name string `json:"name,omitempty"`
	// a description of the plan
	Description string `json:"description,omitempty"`
	// the version of the plan
	Version string `json:"version,omitempty"`
	// metadata for the plan
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntitlementPlanHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entitlementplanhistory.FieldTags, entitlementplanhistory.FieldMetadata:
			values[i] = new([]byte)
		case entitlementplanhistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case entitlementplanhistory.FieldID, entitlementplanhistory.FieldRef, entitlementplanhistory.FieldCreatedBy, entitlementplanhistory.FieldUpdatedBy, entitlementplanhistory.FieldMappingID, entitlementplanhistory.FieldDeletedBy, entitlementplanhistory.FieldOwnerID, entitlementplanhistory.FieldDisplayName, entitlementplanhistory.FieldName, entitlementplanhistory.FieldDescription, entitlementplanhistory.FieldVersion:
			values[i] = new(sql.NullString)
		case entitlementplanhistory.FieldHistoryTime, entitlementplanhistory.FieldCreatedAt, entitlementplanhistory.FieldUpdatedAt, entitlementplanhistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntitlementPlanHistory fields.
func (eph *EntitlementPlanHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entitlementplanhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				eph.ID = value.String
			}
		case entitlementplanhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				eph.HistoryTime = value.Time
			}
		case entitlementplanhistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				eph.Operation = *value
			}
		case entitlementplanhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				eph.Ref = value.String
			}
		case entitlementplanhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				eph.CreatedAt = value.Time
			}
		case entitlementplanhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				eph.UpdatedAt = value.Time
			}
		case entitlementplanhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				eph.CreatedBy = value.String
			}
		case entitlementplanhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				eph.UpdatedBy = value.String
			}
		case entitlementplanhistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				eph.MappingID = value.String
			}
		case entitlementplanhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				eph.DeletedAt = value.Time
			}
		case entitlementplanhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				eph.DeletedBy = value.String
			}
		case entitlementplanhistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &eph.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case entitlementplanhistory.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				eph.OwnerID = value.String
			}
		case entitlementplanhistory.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				eph.DisplayName = value.String
			}
		case entitlementplanhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				eph.Name = value.String
			}
		case entitlementplanhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				eph.Description = value.String
			}
		case entitlementplanhistory.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				eph.Version = value.String
			}
		case entitlementplanhistory.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &eph.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		default:
			eph.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EntitlementPlanHistory.
// This includes values selected through modifiers, order, etc.
func (eph *EntitlementPlanHistory) Value(name string) (ent.Value, error) {
	return eph.selectValues.Get(name)
}

// Update returns a builder for updating this EntitlementPlanHistory.
// Note that you need to call EntitlementPlanHistory.Unwrap() before calling this method if this EntitlementPlanHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (eph *EntitlementPlanHistory) Update() *EntitlementPlanHistoryUpdateOne {
	return NewEntitlementPlanHistoryClient(eph.config).UpdateOne(eph)
}

// Unwrap unwraps the EntitlementPlanHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eph *EntitlementPlanHistory) Unwrap() *EntitlementPlanHistory {
	_tx, ok := eph.config.driver.(*txDriver)
	if !ok {
		panic("generated: EntitlementPlanHistory is not a transactional entity")
	}
	eph.config.driver = _tx.drv
	return eph
}

// String implements the fmt.Stringer.
func (eph *EntitlementPlanHistory) String() string {
	var builder strings.Builder
	builder.WriteString("EntitlementPlanHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", eph.ID))
	builder.WriteString("history_time=")
	builder.WriteString(eph.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", eph.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(eph.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(eph.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(eph.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(eph.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(eph.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(eph.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(eph.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(eph.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", eph.Tags))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(eph.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(eph.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(eph.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(eph.Description)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(eph.Version)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", eph.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// EntitlementPlanHistories is a parsable slice of EntitlementPlanHistory.
type EntitlementPlanHistories []*EntitlementPlanHistory