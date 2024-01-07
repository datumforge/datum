// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// PersonalAccessToken is the model entity for the PersonalAccessToken schema.
type PersonalAccessToken struct {
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
	// the name associated with the token
	Name string `json:"name,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"-"`
	// what abilites the token should have
	Abilities []string `json:"abilities,omitempty"`
	// when the token expires
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// a description of the token's purpose
	Description string `json:"description,omitempty"`
	// LastUsedAt holds the value of the "last_used_at" field.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonalAccessTokenQuery when eager-loading is set.
	Edges        PersonalAccessTokenEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PersonalAccessTokenEdges holds the relations/edges for other nodes in the graph.
type PersonalAccessTokenEdges struct {
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
func (e PersonalAccessTokenEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PersonalAccessToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case personalaccesstoken.FieldAbilities:
			values[i] = new([]byte)
		case personalaccesstoken.FieldID, personalaccesstoken.FieldCreatedBy, personalaccesstoken.FieldUpdatedBy, personalaccesstoken.FieldDeletedBy, personalaccesstoken.FieldOwnerID, personalaccesstoken.FieldName, personalaccesstoken.FieldToken, personalaccesstoken.FieldDescription:
			values[i] = new(sql.NullString)
		case personalaccesstoken.FieldCreatedAt, personalaccesstoken.FieldUpdatedAt, personalaccesstoken.FieldDeletedAt, personalaccesstoken.FieldExpiresAt, personalaccesstoken.FieldLastUsedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PersonalAccessToken fields.
func (pat *PersonalAccessToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personalaccesstoken.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pat.ID = value.String
			}
		case personalaccesstoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pat.CreatedAt = value.Time
			}
		case personalaccesstoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pat.UpdatedAt = value.Time
			}
		case personalaccesstoken.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pat.CreatedBy = value.String
			}
		case personalaccesstoken.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pat.UpdatedBy = value.String
			}
		case personalaccesstoken.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pat.DeletedAt = value.Time
			}
		case personalaccesstoken.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				pat.DeletedBy = value.String
			}
		case personalaccesstoken.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				pat.OwnerID = value.String
			}
		case personalaccesstoken.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pat.Name = value.String
			}
		case personalaccesstoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				pat.Token = value.String
			}
		case personalaccesstoken.FieldAbilities:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field abilities", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pat.Abilities); err != nil {
					return fmt.Errorf("unmarshal field abilities: %w", err)
				}
			}
		case personalaccesstoken.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				pat.ExpiresAt = new(time.Time)
				*pat.ExpiresAt = value.Time
			}
		case personalaccesstoken.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pat.Description = value.String
			}
		case personalaccesstoken.FieldLastUsedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used_at", values[i])
			} else if value.Valid {
				pat.LastUsedAt = new(time.Time)
				*pat.LastUsedAt = value.Time
			}
		default:
			pat.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PersonalAccessToken.
// This includes values selected through modifiers, order, etc.
func (pat *PersonalAccessToken) Value(name string) (ent.Value, error) {
	return pat.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the PersonalAccessToken entity.
func (pat *PersonalAccessToken) QueryOwner() *UserQuery {
	return NewPersonalAccessTokenClient(pat.config).QueryOwner(pat)
}

// Update returns a builder for updating this PersonalAccessToken.
// Note that you need to call PersonalAccessToken.Unwrap() before calling this method if this PersonalAccessToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (pat *PersonalAccessToken) Update() *PersonalAccessTokenUpdateOne {
	return NewPersonalAccessTokenClient(pat.config).UpdateOne(pat)
}

// Unwrap unwraps the PersonalAccessToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pat *PersonalAccessToken) Unwrap() *PersonalAccessToken {
	_tx, ok := pat.config.driver.(*txDriver)
	if !ok {
		panic("generated: PersonalAccessToken is not a transactional entity")
	}
	pat.config.driver = _tx.drv
	return pat
}

// String implements the fmt.Stringer.
func (pat *PersonalAccessToken) String() string {
	var builder strings.Builder
	builder.WriteString("PersonalAccessToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pat.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pat.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pat.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pat.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pat.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pat.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(pat.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(pat.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pat.Name)
	builder.WriteString(", ")
	builder.WriteString("token=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("abilities=")
	builder.WriteString(fmt.Sprintf("%v", pat.Abilities))
	builder.WriteString(", ")
	if v := pat.ExpiresAt; v != nil {
		builder.WriteString("expires_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pat.Description)
	builder.WriteString(", ")
	if v := pat.LastUsedAt; v != nil {
		builder.WriteString("last_used_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// PersonalAccessTokens is a parsable slice of PersonalAccessToken.
type PersonalAccessTokens []*PersonalAccessToken
