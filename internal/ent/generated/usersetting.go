// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
)

// UserSetting is the model entity for the UserSetting schema.
type UserSetting struct {
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
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// user account is locked if unconfirmed or explicitly locked
	Locked bool `json:"locked,omitempty"`
	// The time notifications regarding the user were silenced
	SilencedAt *time.Time `json:"silenced_at,omitempty"`
	// The time the user was suspended
	SuspendedAt *time.Time `json:"suspended_at,omitempty"`
	// Status holds the value of the "status" field.
	Status enums.UserStatus `json:"status,omitempty"`
	// EmailConfirmed holds the value of the "email_confirmed" field.
	EmailConfirmed bool `json:"email_confirmed,omitempty"`
	// tags associated with the user
	Tags []string `json:"tags,omitempty"`
	// TFA secret for the user
	TfaSecret *string `json:"-"`
	// recovery codes for 2fa
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
	// specifies a user may complete authentication by verifying an OTP code delivered through SMS
	IsPhoneOtpAllowed bool `json:"is_phone_otp_allowed,omitempty"`
	// specifies a user may complete authentication by verifying an OTP code delivered through email
	IsEmailOtpAllowed bool `json:"is_email_otp_allowed,omitempty"`
	// specifies a user may complete authentication by verifying a TOTP code delivered through an authenticator app
	IsTotpAllowed bool `json:"is_totp_allowed,omitempty"`
	// specifies a user may complete authentication by verifying a WebAuthn capable device
	IsWebauthnAllowed bool `json:"is_webauthn_allowed,omitempty"`
	// whether the user has two factor authentication enabled
	IsTfaEnabled bool `json:"is_tfa_enabled,omitempty"`
	// phone number associated with the account, used 2factor SMS authentication
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSettingQuery when eager-loading is set.
	Edges                    UserSettingEdges `json:"edges"`
	user_setting_default_org *string
	selectValues             sql.SelectValues
}

// UserSettingEdges holds the relations/edges for other nodes in the graph.
type UserSettingEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// organization to load on user login
	DefaultOrg *Organization `json:"default_org,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSettingEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// DefaultOrgOrErr returns the DefaultOrg value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSettingEdges) DefaultOrgOrErr() (*Organization, error) {
	if e.DefaultOrg != nil {
		return e.DefaultOrg, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "default_org"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersetting.FieldTags, usersetting.FieldRecoveryCodes:
			values[i] = new([]byte)
		case usersetting.FieldLocked, usersetting.FieldEmailConfirmed, usersetting.FieldIsPhoneOtpAllowed, usersetting.FieldIsEmailOtpAllowed, usersetting.FieldIsTotpAllowed, usersetting.FieldIsWebauthnAllowed, usersetting.FieldIsTfaEnabled:
			values[i] = new(sql.NullBool)
		case usersetting.FieldID, usersetting.FieldCreatedBy, usersetting.FieldUpdatedBy, usersetting.FieldDeletedBy, usersetting.FieldUserID, usersetting.FieldStatus, usersetting.FieldTfaSecret, usersetting.FieldPhoneNumber:
			values[i] = new(sql.NullString)
		case usersetting.FieldCreatedAt, usersetting.FieldUpdatedAt, usersetting.FieldDeletedAt, usersetting.FieldSilencedAt, usersetting.FieldSuspendedAt:
			values[i] = new(sql.NullTime)
		case usersetting.ForeignKeys[0]: // user_setting_default_org
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSetting fields.
func (us *UserSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersetting.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				us.ID = value.String
			}
		case usersetting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				us.CreatedAt = value.Time
			}
		case usersetting.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				us.UpdatedAt = value.Time
			}
		case usersetting.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				us.CreatedBy = value.String
			}
		case usersetting.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				us.UpdatedBy = value.String
			}
		case usersetting.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				us.DeletedAt = value.Time
			}
		case usersetting.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				us.DeletedBy = value.String
			}
		case usersetting.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				us.UserID = value.String
			}
		case usersetting.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				us.Locked = value.Bool
			}
		case usersetting.FieldSilencedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field silenced_at", values[i])
			} else if value.Valid {
				us.SilencedAt = new(time.Time)
				*us.SilencedAt = value.Time
			}
		case usersetting.FieldSuspendedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field suspended_at", values[i])
			} else if value.Valid {
				us.SuspendedAt = new(time.Time)
				*us.SuspendedAt = value.Time
			}
		case usersetting.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				us.Status = enums.UserStatus(value.String)
			}
		case usersetting.FieldEmailConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmed", values[i])
			} else if value.Valid {
				us.EmailConfirmed = value.Bool
			}
		case usersetting.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &us.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case usersetting.FieldTfaSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tfa_secret", values[i])
			} else if value.Valid {
				us.TfaSecret = new(string)
				*us.TfaSecret = value.String
			}
		case usersetting.FieldRecoveryCodes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field recovery_codes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &us.RecoveryCodes); err != nil {
					return fmt.Errorf("unmarshal field recovery_codes: %w", err)
				}
			}
		case usersetting.FieldIsPhoneOtpAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_phone_otp_allowed", values[i])
			} else if value.Valid {
				us.IsPhoneOtpAllowed = value.Bool
			}
		case usersetting.FieldIsEmailOtpAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_email_otp_allowed", values[i])
			} else if value.Valid {
				us.IsEmailOtpAllowed = value.Bool
			}
		case usersetting.FieldIsTotpAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_totp_allowed", values[i])
			} else if value.Valid {
				us.IsTotpAllowed = value.Bool
			}
		case usersetting.FieldIsWebauthnAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_webauthn_allowed", values[i])
			} else if value.Valid {
				us.IsWebauthnAllowed = value.Bool
			}
		case usersetting.FieldIsTfaEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_tfa_enabled", values[i])
			} else if value.Valid {
				us.IsTfaEnabled = value.Bool
			}
		case usersetting.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				us.PhoneNumber = new(string)
				*us.PhoneNumber = value.String
			}
		case usersetting.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_setting_default_org", values[i])
			} else if value.Valid {
				us.user_setting_default_org = new(string)
				*us.user_setting_default_org = value.String
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSetting.
// This includes values selected through modifiers, order, etc.
func (us *UserSetting) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserSetting entity.
func (us *UserSetting) QueryUser() *UserQuery {
	return NewUserSettingClient(us.config).QueryUser(us)
}

// QueryDefaultOrg queries the "default_org" edge of the UserSetting entity.
func (us *UserSetting) QueryDefaultOrg() *OrganizationQuery {
	return NewUserSettingClient(us.config).QueryDefaultOrg(us)
}

// Update returns a builder for updating this UserSetting.
// Note that you need to call UserSetting.Unwrap() before calling this method if this UserSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSetting) Update() *UserSettingUpdateOne {
	return NewUserSettingClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSetting) Unwrap() *UserSetting {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("generated: UserSetting is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSetting) String() string {
	var builder strings.Builder
	builder.WriteString("UserSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("created_at=")
	builder.WriteString(us.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(us.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(us.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(us.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(us.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(us.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(us.UserID)
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", us.Locked))
	builder.WriteString(", ")
	if v := us.SilencedAt; v != nil {
		builder.WriteString("silenced_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := us.SuspendedAt; v != nil {
		builder.WriteString("suspended_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", us.Status))
	builder.WriteString(", ")
	builder.WriteString("email_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", us.EmailConfirmed))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", us.Tags))
	builder.WriteString(", ")
	builder.WriteString("tfa_secret=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("recovery_codes=")
	builder.WriteString(fmt.Sprintf("%v", us.RecoveryCodes))
	builder.WriteString(", ")
	builder.WriteString("is_phone_otp_allowed=")
	builder.WriteString(fmt.Sprintf("%v", us.IsPhoneOtpAllowed))
	builder.WriteString(", ")
	builder.WriteString("is_email_otp_allowed=")
	builder.WriteString(fmt.Sprintf("%v", us.IsEmailOtpAllowed))
	builder.WriteString(", ")
	builder.WriteString("is_totp_allowed=")
	builder.WriteString(fmt.Sprintf("%v", us.IsTotpAllowed))
	builder.WriteString(", ")
	builder.WriteString("is_webauthn_allowed=")
	builder.WriteString(fmt.Sprintf("%v", us.IsWebauthnAllowed))
	builder.WriteString(", ")
	builder.WriteString("is_tfa_enabled=")
	builder.WriteString(fmt.Sprintf("%v", us.IsTfaEnabled))
	builder.WriteString(", ")
	if v := us.PhoneNumber; v != nil {
		builder.WriteString("phone_number=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// UserSettings is a parsable slice of UserSetting.
type UserSettings []*UserSetting
