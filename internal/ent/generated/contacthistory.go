// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/contacthistory"
	"github.com/datumforge/datum/pkg/enums"
	"github.com/datumforge/enthistory"
)

// ContactHistory is the model entity for the ContactHistory schema.
type ContactHistory struct {
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
	Status       enums.UserStatus `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ContactHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contacthistory.FieldTags:
			values[i] = new([]byte)
		case contacthistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case contacthistory.FieldID, contacthistory.FieldRef, contacthistory.FieldCreatedBy, contacthistory.FieldUpdatedBy, contacthistory.FieldMappingID, contacthistory.FieldDeletedBy, contacthistory.FieldOwnerID, contacthistory.FieldFullName, contacthistory.FieldTitle, contacthistory.FieldCompany, contacthistory.FieldEmail, contacthistory.FieldPhoneNumber, contacthistory.FieldAddress, contacthistory.FieldStatus:
			values[i] = new(sql.NullString)
		case contacthistory.FieldHistoryTime, contacthistory.FieldCreatedAt, contacthistory.FieldUpdatedAt, contacthistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ContactHistory fields.
func (ch *ContactHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contacthistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ch.ID = value.String
			}
		case contacthistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ch.HistoryTime = value.Time
			}
		case contacthistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				ch.Operation = *value
			}
		case contacthistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				ch.Ref = value.String
			}
		case contacthistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ch.CreatedAt = value.Time
			}
		case contacthistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ch.UpdatedAt = value.Time
			}
		case contacthistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ch.CreatedBy = value.String
			}
		case contacthistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ch.UpdatedBy = value.String
			}
		case contacthistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				ch.MappingID = value.String
			}
		case contacthistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ch.DeletedAt = value.Time
			}
		case contacthistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				ch.DeletedBy = value.String
			}
		case contacthistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ch.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case contacthistory.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				ch.OwnerID = value.String
			}
		case contacthistory.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				ch.FullName = value.String
			}
		case contacthistory.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ch.Title = value.String
			}
		case contacthistory.FieldCompany:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company", values[i])
			} else if value.Valid {
				ch.Company = value.String
			}
		case contacthistory.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				ch.Email = value.String
			}
		case contacthistory.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				ch.PhoneNumber = value.String
			}
		case contacthistory.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				ch.Address = value.String
			}
		case contacthistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ch.Status = enums.UserStatus(value.String)
			}
		default:
			ch.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ContactHistory.
// This includes values selected through modifiers, order, etc.
func (ch *ContactHistory) Value(name string) (ent.Value, error) {
	return ch.selectValues.Get(name)
}

// Update returns a builder for updating this ContactHistory.
// Note that you need to call ContactHistory.Unwrap() before calling this method if this ContactHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ch *ContactHistory) Update() *ContactHistoryUpdateOne {
	return NewContactHistoryClient(ch.config).UpdateOne(ch)
}

// Unwrap unwraps the ContactHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ch *ContactHistory) Unwrap() *ContactHistory {
	_tx, ok := ch.config.driver.(*txDriver)
	if !ok {
		panic("generated: ContactHistory is not a transactional entity")
	}
	ch.config.driver = _tx.drv
	return ch
}

// String implements the fmt.Stringer.
func (ch *ContactHistory) String() string {
	var builder strings.Builder
	builder.WriteString("ContactHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ch.ID))
	builder.WriteString("history_time=")
	builder.WriteString(ch.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ch.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(ch.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ch.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ch.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(ch.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(ch.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(ch.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ch.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(ch.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", ch.Tags))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(ch.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(ch.FullName)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(ch.Title)
	builder.WriteString(", ")
	builder.WriteString("company=")
	builder.WriteString(ch.Company)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(ch.Email)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(ch.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(ch.Address)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ch.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ContactHistories is a parsable slice of ContactHistory.
type ContactHistories []*ContactHistory
