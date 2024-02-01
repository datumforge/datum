// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/oauthproviderhistory"
	"github.com/datumforge/enthistory"
)

// OauthProviderHistory is the model entity for the OauthProviderHistory schema.
type OauthProviderHistory struct {
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
	// the oauth provider's name
	Name string `json:"name,omitempty"`
	// the client id for the oauth provider
	ClientID string `json:"client_id,omitempty"`
	// the client secret
	ClientSecret string `json:"client_secret,omitempty"`
	// the redirect url
	RedirectURL string `json:"redirect_url,omitempty"`
	// the scopes
	Scopes string `json:"scopes,omitempty"`
	// the auth url of the provider
	AuthURL string `json:"auth_url,omitempty"`
	// the token url of the provider
	TokenURL string `json:"token_url,omitempty"`
	// the auth style, 0: auto detect 1: third party log in 2: log in with username and password
	AuthStyle uint8 `json:"auth_style,omitempty"`
	// the URL to request user information by token
	InfoURL      string `json:"info_url,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OauthProviderHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthproviderhistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case oauthproviderhistory.FieldAuthStyle:
			values[i] = new(sql.NullInt64)
		case oauthproviderhistory.FieldID, oauthproviderhistory.FieldRef, oauthproviderhistory.FieldCreatedBy, oauthproviderhistory.FieldUpdatedBy, oauthproviderhistory.FieldDeletedBy, oauthproviderhistory.FieldName, oauthproviderhistory.FieldClientID, oauthproviderhistory.FieldClientSecret, oauthproviderhistory.FieldRedirectURL, oauthproviderhistory.FieldScopes, oauthproviderhistory.FieldAuthURL, oauthproviderhistory.FieldTokenURL, oauthproviderhistory.FieldInfoURL:
			values[i] = new(sql.NullString)
		case oauthproviderhistory.FieldHistoryTime, oauthproviderhistory.FieldCreatedAt, oauthproviderhistory.FieldUpdatedAt, oauthproviderhistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OauthProviderHistory fields.
func (oph *OauthProviderHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthproviderhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				oph.ID = value.String
			}
		case oauthproviderhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				oph.HistoryTime = value.Time
			}
		case oauthproviderhistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				oph.Operation = *value
			}
		case oauthproviderhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				oph.Ref = value.String
			}
		case oauthproviderhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oph.CreatedAt = value.Time
			}
		case oauthproviderhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oph.UpdatedAt = value.Time
			}
		case oauthproviderhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				oph.CreatedBy = value.String
			}
		case oauthproviderhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				oph.UpdatedBy = value.String
			}
		case oauthproviderhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				oph.DeletedAt = value.Time
			}
		case oauthproviderhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				oph.DeletedBy = value.String
			}
		case oauthproviderhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				oph.Name = value.String
			}
		case oauthproviderhistory.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				oph.ClientID = value.String
			}
		case oauthproviderhistory.FieldClientSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_secret", values[i])
			} else if value.Valid {
				oph.ClientSecret = value.String
			}
		case oauthproviderhistory.FieldRedirectURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_url", values[i])
			} else if value.Valid {
				oph.RedirectURL = value.String
			}
		case oauthproviderhistory.FieldScopes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value.Valid {
				oph.Scopes = value.String
			}
		case oauthproviderhistory.FieldAuthURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_url", values[i])
			} else if value.Valid {
				oph.AuthURL = value.String
			}
		case oauthproviderhistory.FieldTokenURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_url", values[i])
			} else if value.Valid {
				oph.TokenURL = value.String
			}
		case oauthproviderhistory.FieldAuthStyle:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field auth_style", values[i])
			} else if value.Valid {
				oph.AuthStyle = uint8(value.Int64)
			}
		case oauthproviderhistory.FieldInfoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info_url", values[i])
			} else if value.Valid {
				oph.InfoURL = value.String
			}
		default:
			oph.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OauthProviderHistory.
// This includes values selected through modifiers, order, etc.
func (oph *OauthProviderHistory) Value(name string) (ent.Value, error) {
	return oph.selectValues.Get(name)
}

// Update returns a builder for updating this OauthProviderHistory.
// Note that you need to call OauthProviderHistory.Unwrap() before calling this method if this OauthProviderHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (oph *OauthProviderHistory) Update() *OauthProviderHistoryUpdateOne {
	return NewOauthProviderHistoryClient(oph.config).UpdateOne(oph)
}

// Unwrap unwraps the OauthProviderHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oph *OauthProviderHistory) Unwrap() *OauthProviderHistory {
	_tx, ok := oph.config.driver.(*txDriver)
	if !ok {
		panic("generated: OauthProviderHistory is not a transactional entity")
	}
	oph.config.driver = _tx.drv
	return oph
}

// String implements the fmt.Stringer.
func (oph *OauthProviderHistory) String() string {
	var builder strings.Builder
	builder.WriteString("OauthProviderHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oph.ID))
	builder.WriteString("history_time=")
	builder.WriteString(oph.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", oph.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(oph.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(oph.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oph.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(oph.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(oph.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(oph.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(oph.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(oph.Name)
	builder.WriteString(", ")
	builder.WriteString("client_id=")
	builder.WriteString(oph.ClientID)
	builder.WriteString(", ")
	builder.WriteString("client_secret=")
	builder.WriteString(oph.ClientSecret)
	builder.WriteString(", ")
	builder.WriteString("redirect_url=")
	builder.WriteString(oph.RedirectURL)
	builder.WriteString(", ")
	builder.WriteString("scopes=")
	builder.WriteString(oph.Scopes)
	builder.WriteString(", ")
	builder.WriteString("auth_url=")
	builder.WriteString(oph.AuthURL)
	builder.WriteString(", ")
	builder.WriteString("token_url=")
	builder.WriteString(oph.TokenURL)
	builder.WriteString(", ")
	builder.WriteString("auth_style=")
	builder.WriteString(fmt.Sprintf("%v", oph.AuthStyle))
	builder.WriteString(", ")
	builder.WriteString("info_url=")
	builder.WriteString(oph.InfoURL)
	builder.WriteByte(')')
	return builder.String()
}

// OauthProviderHistories is a parsable slice of OauthProviderHistory.
type OauthProviderHistories []*OauthProviderHistory
