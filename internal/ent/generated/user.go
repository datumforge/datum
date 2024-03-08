// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
)

// User is the model entity for the User schema.
type User struct {
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
	Password *string `json:"-"`
	// the Subject of the user JWT
	Sub string `json:"sub,omitempty"`
	// auth provider used to register the account
	AuthProvider enums.AuthProvider `json:"auth_provider,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// PersonalAccessTokens holds the value of the personal_access_tokens edge.
	PersonalAccessTokens []*PersonalAccessToken `json:"personal_access_tokens,omitempty"`
	// Setting holds the value of the setting edge.
	Setting *UserSetting `json:"setting,omitempty"`
	// EmailVerificationTokens holds the value of the email_verification_tokens edge.
	EmailVerificationTokens []*EmailVerificationToken `json:"email_verification_tokens,omitempty"`
	// PasswordResetTokens holds the value of the password_reset_tokens edge.
	PasswordResetTokens []*PasswordResetToken `json:"password_reset_tokens,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Organizations holds the value of the organizations edge.
	Organizations []*Organization `json:"organizations,omitempty"`
	// Webauthn holds the value of the webauthn edge.
	Webauthn []*Webauthn `json:"webauthn,omitempty"`
	// GroupMemberships holds the value of the group_memberships edge.
	GroupMemberships []*GroupMembership `json:"group_memberships,omitempty"`
	// OrgMemberships holds the value of the org_memberships edge.
	OrgMemberships []*OrgMembership `json:"org_memberships,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedPersonalAccessTokens    map[string][]*PersonalAccessToken
	namedEmailVerificationTokens map[string][]*EmailVerificationToken
	namedPasswordResetTokens     map[string][]*PasswordResetToken
	namedGroups                  map[string][]*Group
	namedOrganizations           map[string][]*Organization
	namedWebauthn                map[string][]*Webauthn
	namedGroupMemberships        map[string][]*GroupMembership
	namedOrgMemberships          map[string][]*OrgMembership
}

// PersonalAccessTokensOrErr returns the PersonalAccessTokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PersonalAccessTokensOrErr() ([]*PersonalAccessToken, error) {
	if e.loadedTypes[0] {
		return e.PersonalAccessTokens, nil
	}
	return nil, &NotLoadedError{edge: "personal_access_tokens"}
}

// SettingOrErr returns the Setting value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SettingOrErr() (*UserSetting, error) {
	if e.Setting != nil {
		return e.Setting, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: usersetting.Label}
	}
	return nil, &NotLoadedError{edge: "setting"}
}

// EmailVerificationTokensOrErr returns the EmailVerificationTokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EmailVerificationTokensOrErr() ([]*EmailVerificationToken, error) {
	if e.loadedTypes[2] {
		return e.EmailVerificationTokens, nil
	}
	return nil, &NotLoadedError{edge: "email_verification_tokens"}
}

// PasswordResetTokensOrErr returns the PasswordResetTokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PasswordResetTokensOrErr() ([]*PasswordResetToken, error) {
	if e.loadedTypes[3] {
		return e.PasswordResetTokens, nil
	}
	return nil, &NotLoadedError{edge: "password_reset_tokens"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[4] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrganizationsOrErr() ([]*Organization, error) {
	if e.loadedTypes[5] {
		return e.Organizations, nil
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// WebauthnOrErr returns the Webauthn value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WebauthnOrErr() ([]*Webauthn, error) {
	if e.loadedTypes[6] {
		return e.Webauthn, nil
	}
	return nil, &NotLoadedError{edge: "webauthn"}
}

// GroupMembershipsOrErr returns the GroupMemberships value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupMembershipsOrErr() ([]*GroupMembership, error) {
	if e.loadedTypes[7] {
		return e.GroupMemberships, nil
	}
	return nil, &NotLoadedError{edge: "group_memberships"}
}

// OrgMembershipsOrErr returns the OrgMemberships value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrgMembershipsOrErr() ([]*OrgMembership, error) {
	if e.loadedTypes[8] {
		return e.OrgMemberships, nil
	}
	return nil, &NotLoadedError{edge: "org_memberships"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldCreatedBy, user.FieldUpdatedBy, user.FieldDeletedBy, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldDisplayName, user.FieldAvatarRemoteURL, user.FieldAvatarLocalFile, user.FieldPassword, user.FieldSub, user.FieldAuthProvider:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt, user.FieldAvatarUpdatedAt, user.FieldLastSeen:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				u.CreatedBy = value.String
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				u.UpdatedBy = value.String
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				u.DeletedBy = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldAvatarRemoteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_remote_url", values[i])
			} else if value.Valid {
				u.AvatarRemoteURL = new(string)
				*u.AvatarRemoteURL = value.String
			}
		case user.FieldAvatarLocalFile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_local_file", values[i])
			} else if value.Valid {
				u.AvatarLocalFile = new(string)
				*u.AvatarLocalFile = value.String
			}
		case user.FieldAvatarUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_updated_at", values[i])
			} else if value.Valid {
				u.AvatarUpdatedAt = new(time.Time)
				*u.AvatarUpdatedAt = value.Time
			}
		case user.FieldLastSeen:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_seen", values[i])
			} else if value.Valid {
				u.LastSeen = new(time.Time)
				*u.LastSeen = value.Time
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = new(string)
				*u.Password = value.String
			}
		case user.FieldSub:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub", values[i])
			} else if value.Valid {
				u.Sub = value.String
			}
		case user.FieldAuthProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_provider", values[i])
			} else if value.Valid {
				u.AuthProvider = enums.AuthProvider(value.String)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryPersonalAccessTokens queries the "personal_access_tokens" edge of the User entity.
func (u *User) QueryPersonalAccessTokens() *PersonalAccessTokenQuery {
	return NewUserClient(u.config).QueryPersonalAccessTokens(u)
}

// QuerySetting queries the "setting" edge of the User entity.
func (u *User) QuerySetting() *UserSettingQuery {
	return NewUserClient(u.config).QuerySetting(u)
}

// QueryEmailVerificationTokens queries the "email_verification_tokens" edge of the User entity.
func (u *User) QueryEmailVerificationTokens() *EmailVerificationTokenQuery {
	return NewUserClient(u.config).QueryEmailVerificationTokens(u)
}

// QueryPasswordResetTokens queries the "password_reset_tokens" edge of the User entity.
func (u *User) QueryPasswordResetTokens() *PasswordResetTokenQuery {
	return NewUserClient(u.config).QueryPasswordResetTokens(u)
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return NewUserClient(u.config).QueryGroups(u)
}

// QueryOrganizations queries the "organizations" edge of the User entity.
func (u *User) QueryOrganizations() *OrganizationQuery {
	return NewUserClient(u.config).QueryOrganizations(u)
}

// QueryWebauthn queries the "webauthn" edge of the User entity.
func (u *User) QueryWebauthn() *WebauthnQuery {
	return NewUserClient(u.config).QueryWebauthn(u)
}

// QueryGroupMemberships queries the "group_memberships" edge of the User entity.
func (u *User) QueryGroupMemberships() *GroupMembershipQuery {
	return NewUserClient(u.config).QueryGroupMemberships(u)
}

// QueryOrgMemberships queries the "org_memberships" edge of the User entity.
func (u *User) QueryOrgMemberships() *OrgMembershipQuery {
	return NewUserClient(u.config).QueryOrgMemberships(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("generated: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(u.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(u.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(u.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", ")
	if v := u.AvatarRemoteURL; v != nil {
		builder.WriteString("avatar_remote_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.AvatarLocalFile; v != nil {
		builder.WriteString("avatar_local_file=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.AvatarUpdatedAt; v != nil {
		builder.WriteString("avatar_updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.LastSeen; v != nil {
		builder.WriteString("last_seen=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("sub=")
	builder.WriteString(u.Sub)
	builder.WriteString(", ")
	builder.WriteString("auth_provider=")
	builder.WriteString(fmt.Sprintf("%v", u.AuthProvider))
	builder.WriteByte(')')
	return builder.String()
}

// NamedPersonalAccessTokens returns the PersonalAccessTokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPersonalAccessTokens(name string) ([]*PersonalAccessToken, error) {
	if u.Edges.namedPersonalAccessTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPersonalAccessTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPersonalAccessTokens(name string, edges ...*PersonalAccessToken) {
	if u.Edges.namedPersonalAccessTokens == nil {
		u.Edges.namedPersonalAccessTokens = make(map[string][]*PersonalAccessToken)
	}
	if len(edges) == 0 {
		u.Edges.namedPersonalAccessTokens[name] = []*PersonalAccessToken{}
	} else {
		u.Edges.namedPersonalAccessTokens[name] = append(u.Edges.namedPersonalAccessTokens[name], edges...)
	}
}

// NamedEmailVerificationTokens returns the EmailVerificationTokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedEmailVerificationTokens(name string) ([]*EmailVerificationToken, error) {
	if u.Edges.namedEmailVerificationTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedEmailVerificationTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedEmailVerificationTokens(name string, edges ...*EmailVerificationToken) {
	if u.Edges.namedEmailVerificationTokens == nil {
		u.Edges.namedEmailVerificationTokens = make(map[string][]*EmailVerificationToken)
	}
	if len(edges) == 0 {
		u.Edges.namedEmailVerificationTokens[name] = []*EmailVerificationToken{}
	} else {
		u.Edges.namedEmailVerificationTokens[name] = append(u.Edges.namedEmailVerificationTokens[name], edges...)
	}
}

// NamedPasswordResetTokens returns the PasswordResetTokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPasswordResetTokens(name string) ([]*PasswordResetToken, error) {
	if u.Edges.namedPasswordResetTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPasswordResetTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPasswordResetTokens(name string, edges ...*PasswordResetToken) {
	if u.Edges.namedPasswordResetTokens == nil {
		u.Edges.namedPasswordResetTokens = make(map[string][]*PasswordResetToken)
	}
	if len(edges) == 0 {
		u.Edges.namedPasswordResetTokens[name] = []*PasswordResetToken{}
	} else {
		u.Edges.namedPasswordResetTokens[name] = append(u.Edges.namedPasswordResetTokens[name], edges...)
	}
}

// NamedGroups returns the Groups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedGroups(name string) ([]*Group, error) {
	if u.Edges.namedGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedGroups(name string, edges ...*Group) {
	if u.Edges.namedGroups == nil {
		u.Edges.namedGroups = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		u.Edges.namedGroups[name] = []*Group{}
	} else {
		u.Edges.namedGroups[name] = append(u.Edges.namedGroups[name], edges...)
	}
}

// NamedOrganizations returns the Organizations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOrganizations(name string) ([]*Organization, error) {
	if u.Edges.namedOrganizations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOrganizations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOrganizations(name string, edges ...*Organization) {
	if u.Edges.namedOrganizations == nil {
		u.Edges.namedOrganizations = make(map[string][]*Organization)
	}
	if len(edges) == 0 {
		u.Edges.namedOrganizations[name] = []*Organization{}
	} else {
		u.Edges.namedOrganizations[name] = append(u.Edges.namedOrganizations[name], edges...)
	}
}

// NamedWebauthn returns the Webauthn named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedWebauthn(name string) ([]*Webauthn, error) {
	if u.Edges.namedWebauthn == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedWebauthn[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedWebauthn(name string, edges ...*Webauthn) {
	if u.Edges.namedWebauthn == nil {
		u.Edges.namedWebauthn = make(map[string][]*Webauthn)
	}
	if len(edges) == 0 {
		u.Edges.namedWebauthn[name] = []*Webauthn{}
	} else {
		u.Edges.namedWebauthn[name] = append(u.Edges.namedWebauthn[name], edges...)
	}
}

// NamedGroupMemberships returns the GroupMemberships named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedGroupMemberships(name string) ([]*GroupMembership, error) {
	if u.Edges.namedGroupMemberships == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedGroupMemberships[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedGroupMemberships(name string, edges ...*GroupMembership) {
	if u.Edges.namedGroupMemberships == nil {
		u.Edges.namedGroupMemberships = make(map[string][]*GroupMembership)
	}
	if len(edges) == 0 {
		u.Edges.namedGroupMemberships[name] = []*GroupMembership{}
	} else {
		u.Edges.namedGroupMemberships[name] = append(u.Edges.namedGroupMemberships[name], edges...)
	}
}

// NamedOrgMemberships returns the OrgMemberships named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOrgMemberships(name string) ([]*OrgMembership, error) {
	if u.Edges.namedOrgMemberships == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOrgMemberships[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOrgMemberships(name string, edges ...*OrgMembership) {
	if u.Edges.namedOrgMemberships == nil {
		u.Edges.namedOrgMemberships = make(map[string][]*OrgMembership)
	}
	if len(edges) == 0 {
		u.Edges.namedOrgMemberships[name] = []*OrgMembership{}
	} else {
		u.Edges.namedOrgMemberships[name] = append(u.Edges.namedOrgMemberships[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
