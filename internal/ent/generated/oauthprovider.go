// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
	"github.com/datumforge/datum/internal/ent/generated/organization"
)

// OauthProvider is the model entity for the OauthProvider schema.
type OauthProvider struct {
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
	AuthStyle customtypes.Uint8 `json:"auth_style,omitempty"`
	// the URL to request user information by token
	InfoURL string `json:"info_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OauthProviderQuery when eager-loading is set.
	Edges                      OauthProviderEdges `json:"edges"`
	organization_oauthprovider *string
	selectValues               sql.SelectValues
}

// OauthProviderEdges holds the relations/edges for other nodes in the graph.
type OauthProviderEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OauthProviderEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OauthProvider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthprovider.FieldAuthStyle:
			values[i] = new(sql.NullInt64)
		case oauthprovider.FieldID, oauthprovider.FieldCreatedBy, oauthprovider.FieldUpdatedBy, oauthprovider.FieldMappingID, oauthprovider.FieldDeletedBy, oauthprovider.FieldName, oauthprovider.FieldClientID, oauthprovider.FieldClientSecret, oauthprovider.FieldRedirectURL, oauthprovider.FieldScopes, oauthprovider.FieldAuthURL, oauthprovider.FieldTokenURL, oauthprovider.FieldInfoURL:
			values[i] = new(sql.NullString)
		case oauthprovider.FieldCreatedAt, oauthprovider.FieldUpdatedAt, oauthprovider.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case oauthprovider.ForeignKeys[0]: // organization_oauthprovider
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OauthProvider fields.
func (op *OauthProvider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthprovider.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				op.ID = value.String
			}
		case oauthprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				op.CreatedAt = value.Time
			}
		case oauthprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				op.UpdatedAt = value.Time
			}
		case oauthprovider.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				op.CreatedBy = value.String
			}
		case oauthprovider.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				op.UpdatedBy = value.String
			}
		case oauthprovider.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				op.MappingID = value.String
			}
		case oauthprovider.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				op.DeletedAt = value.Time
			}
		case oauthprovider.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				op.DeletedBy = value.String
			}
		case oauthprovider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				op.Name = value.String
			}
		case oauthprovider.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				op.ClientID = value.String
			}
		case oauthprovider.FieldClientSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_secret", values[i])
			} else if value.Valid {
				op.ClientSecret = value.String
			}
		case oauthprovider.FieldRedirectURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_url", values[i])
			} else if value.Valid {
				op.RedirectURL = value.String
			}
		case oauthprovider.FieldScopes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value.Valid {
				op.Scopes = value.String
			}
		case oauthprovider.FieldAuthURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_url", values[i])
			} else if value.Valid {
				op.AuthURL = value.String
			}
		case oauthprovider.FieldTokenURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_url", values[i])
			} else if value.Valid {
				op.TokenURL = value.String
			}
		case oauthprovider.FieldAuthStyle:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field auth_style", values[i])
			} else if value.Valid {
				op.AuthStyle = customtypes.Uint8(value.Int64)
			}
		case oauthprovider.FieldInfoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info_url", values[i])
			} else if value.Valid {
				op.InfoURL = value.String
			}
		case oauthprovider.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_oauthprovider", values[i])
			} else if value.Valid {
				op.organization_oauthprovider = new(string)
				*op.organization_oauthprovider = value.String
			}
		default:
			op.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OauthProvider.
// This includes values selected through modifiers, order, etc.
func (op *OauthProvider) Value(name string) (ent.Value, error) {
	return op.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the OauthProvider entity.
func (op *OauthProvider) QueryOwner() *OrganizationQuery {
	return NewOauthProviderClient(op.config).QueryOwner(op)
}

// Update returns a builder for updating this OauthProvider.
// Note that you need to call OauthProvider.Unwrap() before calling this method if this OauthProvider
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OauthProvider) Update() *OauthProviderUpdateOne {
	return NewOauthProviderClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OauthProvider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OauthProvider) Unwrap() *OauthProvider {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("generated: OauthProvider is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OauthProvider) String() string {
	var builder strings.Builder
	builder.WriteString("OauthProvider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("created_at=")
	builder.WriteString(op.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(op.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(op.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(op.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(op.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(op.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(op.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(op.Name)
	builder.WriteString(", ")
	builder.WriteString("client_id=")
	builder.WriteString(op.ClientID)
	builder.WriteString(", ")
	builder.WriteString("client_secret=")
	builder.WriteString(op.ClientSecret)
	builder.WriteString(", ")
	builder.WriteString("redirect_url=")
	builder.WriteString(op.RedirectURL)
	builder.WriteString(", ")
	builder.WriteString("scopes=")
	builder.WriteString(op.Scopes)
	builder.WriteString(", ")
	builder.WriteString("auth_url=")
	builder.WriteString(op.AuthURL)
	builder.WriteString(", ")
	builder.WriteString("token_url=")
	builder.WriteString(op.TokenURL)
	builder.WriteString(", ")
	builder.WriteString("auth_style=")
	builder.WriteString(fmt.Sprintf("%v", op.AuthStyle))
	builder.WriteString(", ")
	builder.WriteString("info_url=")
	builder.WriteString(op.InfoURL)
	builder.WriteByte(')')
	return builder.String()
}

// OauthProviders is a parsable slice of OauthProvider.
type OauthProviders []*OauthProvider
