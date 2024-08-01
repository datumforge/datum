// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// PasswordResetToken is the model entity for the PasswordResetToken schema.
type PasswordResetToken struct {
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
	// The user id that owns the object
	OwnerID string `json:"owner_id,omitempty"`
	// the reset token sent to the user via email which should only be provided to the /forgot-password endpoint + handler
	Token string `json:"token,omitempty"`
	// the ttl of the reset token which defaults to 15 minutes
	TTL *time.Time `json:"ttl,omitempty"`
	// the email used as input to generate the reset token; this is used to verify that the token when regenerated within the server matches the token emailed
	Email string `json:"email,omitempty"`
	// the comparison secret to verify the token's signature
	Secret *[]byte `json:"secret,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PasswordResetTokenQuery when eager-loading is set.
	Edges        PasswordResetTokenEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PasswordResetTokenEdges holds the relations/edges for other nodes in the graph.
type PasswordResetTokenEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PasswordResetTokenEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PasswordResetToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case passwordresettoken.FieldSecret:
			values[i] = new([]byte)
		case passwordresettoken.FieldID, passwordresettoken.FieldCreatedBy, passwordresettoken.FieldUpdatedBy, passwordresettoken.FieldMappingID, passwordresettoken.FieldDeletedBy, passwordresettoken.FieldOwnerID, passwordresettoken.FieldToken, passwordresettoken.FieldEmail:
			values[i] = new(sql.NullString)
		case passwordresettoken.FieldCreatedAt, passwordresettoken.FieldUpdatedAt, passwordresettoken.FieldDeletedAt, passwordresettoken.FieldTTL:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PasswordResetToken fields.
func (prt *PasswordResetToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case passwordresettoken.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				prt.ID = value.String
			}
		case passwordresettoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				prt.CreatedAt = value.Time
			}
		case passwordresettoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				prt.UpdatedAt = value.Time
			}
		case passwordresettoken.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				prt.CreatedBy = value.String
			}
		case passwordresettoken.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				prt.UpdatedBy = value.String
			}
		case passwordresettoken.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				prt.MappingID = value.String
			}
		case passwordresettoken.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				prt.DeletedAt = value.Time
			}
		case passwordresettoken.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				prt.DeletedBy = value.String
			}
		case passwordresettoken.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				prt.OwnerID = value.String
			}
		case passwordresettoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				prt.Token = value.String
			}
		case passwordresettoken.FieldTTL:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ttl", values[i])
			} else if value.Valid {
				prt.TTL = new(time.Time)
				*prt.TTL = value.Time
			}
		case passwordresettoken.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				prt.Email = value.String
			}
		case passwordresettoken.FieldSecret:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value != nil {
				prt.Secret = value
			}
		default:
			prt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PasswordResetToken.
// This includes values selected through modifiers, order, etc.
func (prt *PasswordResetToken) Value(name string) (ent.Value, error) {
	return prt.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the PasswordResetToken entity.
func (prt *PasswordResetToken) QueryOwner() *UserQuery {
	return NewPasswordResetTokenClient(prt.config).QueryOwner(prt)
}

// Update returns a builder for updating this PasswordResetToken.
// Note that you need to call PasswordResetToken.Unwrap() before calling this method if this PasswordResetToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (prt *PasswordResetToken) Update() *PasswordResetTokenUpdateOne {
	return NewPasswordResetTokenClient(prt.config).UpdateOne(prt)
}

// Unwrap unwraps the PasswordResetToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (prt *PasswordResetToken) Unwrap() *PasswordResetToken {
	_tx, ok := prt.config.driver.(*txDriver)
	if !ok {
		panic("generated: PasswordResetToken is not a transactional entity")
	}
	prt.config.driver = _tx.drv
	return prt
}

// String implements the fmt.Stringer.
func (prt *PasswordResetToken) String() string {
	var builder strings.Builder
	builder.WriteString("PasswordResetToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", prt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(prt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(prt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(prt.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(prt.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(prt.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(prt.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(prt.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(prt.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(prt.Token)
	builder.WriteString(", ")
	if v := prt.TTL; v != nil {
		builder.WriteString("ttl=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(prt.Email)
	builder.WriteString(", ")
	if v := prt.Secret; v != nil {
		builder.WriteString("secret=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// PasswordResetTokens is a parsable slice of PasswordResetToken.
type PasswordResetTokens []*PasswordResetToken
