// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/userhistory"
	"github.com/flume/enthistory"
)

// UserHistory is the model entity for the UserHistory schema.
type UserHistory struct {
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
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// The user's displayed 'friendly' name
	DisplayName string `json:"display_name,omitempty"`
	// URL of the user's remote avatar
	AvatarRemoteURL *string `json:"avatar_remote_url,omitempty"`
	// The user's local avatar file
	AvatarLocalFile *string `json:"avatar_local_file,omitempty"`
	// The time the user's (local) avatar was last updated
	AvatarUpdatedAt *time.Time `json:"avatar_updated_at,omitempty"`
	// the time the user was last seen
	LastSeen *time.Time `json:"last_seen,omitempty"`
	// user password hash
	Password *string `json:"password,omitempty"`
	// the Subject of the user JWT
	Sub string `json:"sub,omitempty"`
	// whether the user uses oauth for login or not
	Oauth        bool `json:"oauth,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userhistory.FieldOauth:
			values[i] = new(sql.NullBool)
		case userhistory.FieldID, userhistory.FieldOperation, userhistory.FieldRef, userhistory.FieldCreatedBy, userhistory.FieldUpdatedBy, userhistory.FieldDeletedBy, userhistory.FieldEmail, userhistory.FieldFirstName, userhistory.FieldLastName, userhistory.FieldDisplayName, userhistory.FieldAvatarRemoteURL, userhistory.FieldAvatarLocalFile, userhistory.FieldPassword, userhistory.FieldSub:
			values[i] = new(sql.NullString)
		case userhistory.FieldHistoryTime, userhistory.FieldCreatedAt, userhistory.FieldUpdatedAt, userhistory.FieldDeletedAt, userhistory.FieldAvatarUpdatedAt, userhistory.FieldLastSeen:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserHistory fields.
func (uh *UserHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				uh.ID = value.String
			}
		case userhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				uh.HistoryTime = value.Time
			}
		case userhistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				uh.Operation = enthistory.OpType(value.String)
			}
		case userhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				uh.Ref = value.String
			}
		case userhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uh.CreatedAt = value.Time
			}
		case userhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				uh.UpdatedAt = value.Time
			}
		case userhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				uh.CreatedBy = value.String
			}
		case userhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				uh.UpdatedBy = value.String
			}
		case userhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				uh.DeletedAt = value.Time
			}
		case userhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				uh.DeletedBy = value.String
			}
		case userhistory.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				uh.Email = value.String
			}
		case userhistory.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				uh.FirstName = value.String
			}
		case userhistory.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				uh.LastName = value.String
			}
		case userhistory.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				uh.DisplayName = value.String
			}
		case userhistory.FieldAvatarRemoteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_remote_url", values[i])
			} else if value.Valid {
				uh.AvatarRemoteURL = new(string)
				*uh.AvatarRemoteURL = value.String
			}
		case userhistory.FieldAvatarLocalFile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_local_file", values[i])
			} else if value.Valid {
				uh.AvatarLocalFile = new(string)
				*uh.AvatarLocalFile = value.String
			}
		case userhistory.FieldAvatarUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_updated_at", values[i])
			} else if value.Valid {
				uh.AvatarUpdatedAt = new(time.Time)
				*uh.AvatarUpdatedAt = value.Time
			}
		case userhistory.FieldLastSeen:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_seen", values[i])
			} else if value.Valid {
				uh.LastSeen = new(time.Time)
				*uh.LastSeen = value.Time
			}
		case userhistory.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				uh.Password = new(string)
				*uh.Password = value.String
			}
		case userhistory.FieldSub:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub", values[i])
			} else if value.Valid {
				uh.Sub = value.String
			}
		case userhistory.FieldOauth:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field oauth", values[i])
			} else if value.Valid {
				uh.Oauth = value.Bool
			}
		default:
			uh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserHistory.
// This includes values selected through modifiers, order, etc.
func (uh *UserHistory) Value(name string) (ent.Value, error) {
	return uh.selectValues.Get(name)
}

// Update returns a builder for updating this UserHistory.
// Note that you need to call UserHistory.Unwrap() before calling this method if this UserHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (uh *UserHistory) Update() *UserHistoryUpdateOne {
	return NewUserHistoryClient(uh.config).UpdateOne(uh)
}

// Unwrap unwraps the UserHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uh *UserHistory) Unwrap() *UserHistory {
	_tx, ok := uh.config.driver.(*txDriver)
	if !ok {
		panic("generated: UserHistory is not a transactional entity")
	}
	uh.config.driver = _tx.drv
	return uh
}

// String implements the fmt.Stringer.
func (uh *UserHistory) String() string {
	var builder strings.Builder
	builder.WriteString("UserHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(uh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", uh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(uh.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(uh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(uh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(uh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(uh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(uh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(uh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(uh.Email)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(uh.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(uh.LastName)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(uh.DisplayName)
	builder.WriteString(", ")
	if v := uh.AvatarRemoteURL; v != nil {
		builder.WriteString("avatar_remote_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := uh.AvatarLocalFile; v != nil {
		builder.WriteString("avatar_local_file=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := uh.AvatarUpdatedAt; v != nil {
		builder.WriteString("avatar_updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := uh.LastSeen; v != nil {
		builder.WriteString("last_seen=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := uh.Password; v != nil {
		builder.WriteString("password=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("sub=")
	builder.WriteString(uh.Sub)
	builder.WriteString(", ")
	builder.WriteString("oauth=")
	builder.WriteString(fmt.Sprintf("%v", uh.Oauth))
	builder.WriteByte(')')
	return builder.String()
}

// UserHistories is a parsable slice of UserHistory.
type UserHistories []*UserHistory
