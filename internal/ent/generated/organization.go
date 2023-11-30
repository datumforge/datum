// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// The organization's displayed 'friendly' name
	DisplayName string `json:"display_name,omitempty"`
	// An optional description of the Organization
	Description string `json:"description,omitempty"`
	// The ID of the parent organization for the organization.
	ParentOrganizationID string `json:"parent_organization_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges        OrganizationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Organization `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Organization `json:"children,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Integrations holds the value of the integrations edge.
	Integrations []*Integration `json:"integrations,omitempty"`
	// Setting holds the value of the setting edge.
	Setting *OrganizationSetting `json:"setting,omitempty"`
	// Entitlements holds the value of the entitlements edge.
	Entitlements []*Entitlement `json:"entitlements,omitempty"`
	// Oauthprovider holds the value of the oauthprovider edge.
	Oauthprovider []*OauthProvider `json:"oauthprovider,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
	// totalCount holds the count of the edges above.
	totalCount [8]map[string]int

	namedChildren      map[string][]*Organization
	namedUsers         map[string][]*User
	namedGroups        map[string][]*Group
	namedIntegrations  map[string][]*Integration
	namedEntitlements  map[string][]*Entitlement
	namedOauthprovider map[string][]*OauthProvider
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) ParentOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) ChildrenOrErr() ([]*Organization, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[3] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// IntegrationsOrErr returns the Integrations value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) IntegrationsOrErr() ([]*Integration, error) {
	if e.loadedTypes[4] {
		return e.Integrations, nil
	}
	return nil, &NotLoadedError{edge: "integrations"}
}

// SettingOrErr returns the Setting value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) SettingOrErr() (*OrganizationSetting, error) {
	if e.loadedTypes[5] {
		if e.Setting == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organizationsetting.Label}
		}
		return e.Setting, nil
	}
	return nil, &NotLoadedError{edge: "setting"}
}

// EntitlementsOrErr returns the Entitlements value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) EntitlementsOrErr() ([]*Entitlement, error) {
	if e.loadedTypes[6] {
		return e.Entitlements, nil
	}
	return nil, &NotLoadedError{edge: "entitlements"}
}

// OauthproviderOrErr returns the Oauthprovider value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) OauthproviderOrErr() ([]*OauthProvider, error) {
	if e.loadedTypes[7] {
		return e.Oauthprovider, nil
	}
	return nil, &NotLoadedError{edge: "oauthprovider"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldID, organization.FieldCreatedBy, organization.FieldUpdatedBy, organization.FieldDeletedBy, organization.FieldName, organization.FieldDisplayName, organization.FieldDescription, organization.FieldParentOrganizationID:
			values[i] = new(sql.NullString)
		case organization.FieldCreatedAt, organization.FieldUpdatedAt, organization.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				o.ID = value.String
			}
		case organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case organization.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case organization.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				o.CreatedBy = value.String
			}
		case organization.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				o.UpdatedBy = value.String
			}
		case organization.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = value.Time
			}
		case organization.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				o.DeletedBy = value.String
			}
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case organization.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				o.DisplayName = value.String
			}
		case organization.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				o.Description = value.String
			}
		case organization.FieldParentOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_organization_id", values[i])
			} else if value.Valid {
				o.ParentOrganizationID = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organization.
// This includes values selected through modifiers, order, etc.
func (o *Organization) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Organization entity.
func (o *Organization) QueryParent() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryParent(o)
}

// QueryChildren queries the "children" edge of the Organization entity.
func (o *Organization) QueryChildren() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryChildren(o)
}

// QueryUsers queries the "users" edge of the Organization entity.
func (o *Organization) QueryUsers() *UserQuery {
	return NewOrganizationClient(o.config).QueryUsers(o)
}

// QueryGroups queries the "groups" edge of the Organization entity.
func (o *Organization) QueryGroups() *GroupQuery {
	return NewOrganizationClient(o.config).QueryGroups(o)
}

// QueryIntegrations queries the "integrations" edge of the Organization entity.
func (o *Organization) QueryIntegrations() *IntegrationQuery {
	return NewOrganizationClient(o.config).QueryIntegrations(o)
}

// QuerySetting queries the "setting" edge of the Organization entity.
func (o *Organization) QuerySetting() *OrganizationSettingQuery {
	return NewOrganizationClient(o.config).QuerySetting(o)
}

// QueryEntitlements queries the "entitlements" edge of the Organization entity.
func (o *Organization) QueryEntitlements() *EntitlementQuery {
	return NewOrganizationClient(o.config).QueryEntitlements(o)
}

// QueryOauthprovider queries the "oauthprovider" edge of the Organization entity.
func (o *Organization) QueryOauthprovider() *OauthProviderQuery {
	return NewOrganizationClient(o.config).QueryOauthprovider(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("generated: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(o.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(o.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(o.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(o.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(o.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(o.Description)
	builder.WriteString(", ")
	builder.WriteString("parent_organization_id=")
	builder.WriteString(o.ParentOrganizationID)
	builder.WriteByte(')')
	return builder.String()
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedChildren(name string) ([]*Organization, error) {
	if o.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedChildren(name string, edges ...*Organization) {
	if o.Edges.namedChildren == nil {
		o.Edges.namedChildren = make(map[string][]*Organization)
	}
	if len(edges) == 0 {
		o.Edges.namedChildren[name] = []*Organization{}
	} else {
		o.Edges.namedChildren[name] = append(o.Edges.namedChildren[name], edges...)
	}
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedUsers(name string) ([]*User, error) {
	if o.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedUsers(name string, edges ...*User) {
	if o.Edges.namedUsers == nil {
		o.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		o.Edges.namedUsers[name] = []*User{}
	} else {
		o.Edges.namedUsers[name] = append(o.Edges.namedUsers[name], edges...)
	}
}

// NamedGroups returns the Groups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedGroups(name string) ([]*Group, error) {
	if o.Edges.namedGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedGroups(name string, edges ...*Group) {
	if o.Edges.namedGroups == nil {
		o.Edges.namedGroups = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		o.Edges.namedGroups[name] = []*Group{}
	} else {
		o.Edges.namedGroups[name] = append(o.Edges.namedGroups[name], edges...)
	}
}

// NamedIntegrations returns the Integrations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedIntegrations(name string) ([]*Integration, error) {
	if o.Edges.namedIntegrations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedIntegrations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedIntegrations(name string, edges ...*Integration) {
	if o.Edges.namedIntegrations == nil {
		o.Edges.namedIntegrations = make(map[string][]*Integration)
	}
	if len(edges) == 0 {
		o.Edges.namedIntegrations[name] = []*Integration{}
	} else {
		o.Edges.namedIntegrations[name] = append(o.Edges.namedIntegrations[name], edges...)
	}
}

// NamedEntitlements returns the Entitlements named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedEntitlements(name string) ([]*Entitlement, error) {
	if o.Edges.namedEntitlements == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedEntitlements[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedEntitlements(name string, edges ...*Entitlement) {
	if o.Edges.namedEntitlements == nil {
		o.Edges.namedEntitlements = make(map[string][]*Entitlement)
	}
	if len(edges) == 0 {
		o.Edges.namedEntitlements[name] = []*Entitlement{}
	} else {
		o.Edges.namedEntitlements[name] = append(o.Edges.namedEntitlements[name], edges...)
	}
}

// NamedOauthprovider returns the Oauthprovider named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedOauthprovider(name string) ([]*OauthProvider, error) {
	if o.Edges.namedOauthprovider == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedOauthprovider[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedOauthprovider(name string, edges ...*OauthProvider) {
	if o.Edges.namedOauthprovider == nil {
		o.Edges.namedOauthprovider = make(map[string][]*OauthProvider)
	}
	if len(edges) == 0 {
		o.Edges.namedOauthprovider[name] = []*OauthProvider{}
	} else {
		o.Edges.namedOauthprovider[name] = append(o.Edges.namedOauthprovider[name], edges...)
	}
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
