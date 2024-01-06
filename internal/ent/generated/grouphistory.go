// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/grouphistory"
	"github.com/flume/enthistory"
)

// GroupHistory is the model entity for the GroupHistory schema.
type GroupHistory struct {
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
	// the name of the group - must be unique within the organization
	Name string `json:"name,omitempty"`
	// the groups description
	Description string `json:"description,omitempty"`
	// the URL to an auto generated gravatar image for the group
	GravatarLogoURL string `json:"gravatar_logo_url,omitempty"`
	// the URL to an image uploaded by the customer for the groups avatar image
	LogoURL string `json:"logo_url,omitempty"`
	// The group's displayed 'friendly' name
	DisplayName  string `json:"display_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case grouphistory.FieldID, grouphistory.FieldOperation, grouphistory.FieldRef, grouphistory.FieldCreatedBy, grouphistory.FieldUpdatedBy, grouphistory.FieldDeletedBy, grouphistory.FieldName, grouphistory.FieldDescription, grouphistory.FieldGravatarLogoURL, grouphistory.FieldLogoURL, grouphistory.FieldDisplayName:
			values[i] = new(sql.NullString)
		case grouphistory.FieldHistoryTime, grouphistory.FieldCreatedAt, grouphistory.FieldUpdatedAt, grouphistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupHistory fields.
func (gh *GroupHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case grouphistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gh.ID = value.String
			}
		case grouphistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				gh.HistoryTime = value.Time
			}
		case grouphistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				gh.Operation = enthistory.OpType(value.String)
			}
		case grouphistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				gh.Ref = value.String
			}
		case grouphistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gh.CreatedAt = value.Time
			}
		case grouphistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gh.UpdatedAt = value.Time
			}
		case grouphistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gh.CreatedBy = value.String
			}
		case grouphistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gh.UpdatedBy = value.String
			}
		case grouphistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gh.DeletedAt = value.Time
			}
		case grouphistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				gh.DeletedBy = value.String
			}
		case grouphistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gh.Name = value.String
			}
		case grouphistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gh.Description = value.String
			}
		case grouphistory.FieldGravatarLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gravatar_logo_url", values[i])
			} else if value.Valid {
				gh.GravatarLogoURL = value.String
			}
		case grouphistory.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				gh.LogoURL = value.String
			}
		case grouphistory.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				gh.DisplayName = value.String
			}
		default:
			gh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupHistory.
// This includes values selected through modifiers, order, etc.
func (gh *GroupHistory) Value(name string) (ent.Value, error) {
	return gh.selectValues.Get(name)
}

// Update returns a builder for updating this GroupHistory.
// Note that you need to call GroupHistory.Unwrap() before calling this method if this GroupHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (gh *GroupHistory) Update() *GroupHistoryUpdateOne {
	return NewGroupHistoryClient(gh.config).UpdateOne(gh)
}

// Unwrap unwraps the GroupHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gh *GroupHistory) Unwrap() *GroupHistory {
	_tx, ok := gh.config.driver.(*txDriver)
	if !ok {
		panic("generated: GroupHistory is not a transactional entity")
	}
	gh.config.driver = _tx.drv
	return gh
}

// String implements the fmt.Stringer.
func (gh *GroupHistory) String() string {
	var builder strings.Builder
	builder.WriteString("GroupHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(gh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", gh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(gh.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(gh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(gh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(gh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gh.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(gh.Description)
	builder.WriteString(", ")
	builder.WriteString("gravatar_logo_url=")
	builder.WriteString(gh.GravatarLogoURL)
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(gh.LogoURL)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(gh.DisplayName)
	builder.WriteByte(')')
	return builder.String()
}

// GroupHistories is a parsable slice of GroupHistory.
type GroupHistories []*GroupHistory
