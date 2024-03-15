// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldParentOrganizationID holds the string denoting the parent_organization_id field in the database.
	FieldParentOrganizationID = "parent_organization_id"
	// FieldPersonalOrg holds the string denoting the personal_org field in the database.
	FieldPersonalOrg = "personal_org"
	// FieldAvatarRemoteURL holds the string denoting the avatar_remote_url field in the database.
	FieldAvatarRemoteURL = "avatar_remote_url"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeIntegrations holds the string denoting the integrations edge name in mutations.
	EdgeIntegrations = "integrations"
	// EdgeSetting holds the string denoting the setting edge name in mutations.
	EdgeSetting = "setting"
	// EdgeEntitlements holds the string denoting the entitlements edge name in mutations.
	EdgeEntitlements = "entitlements"
	// EdgePersonalAccessTokens holds the string denoting the personal_access_tokens edge name in mutations.
	EdgePersonalAccessTokens = "personal_access_tokens"
	// EdgeOauthprovider holds the string denoting the oauthprovider edge name in mutations.
	EdgeOauthprovider = "oauthprovider"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeInvites holds the string denoting the invites edge name in mutations.
	EdgeInvites = "invites"
	// EdgeSubscribers holds the string denoting the subscribers edge name in mutations.
	EdgeSubscribers = "subscribers"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "organizations"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_organization_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "organizations"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_organization_id"
	// GroupsTable is the table that holds the groups relation/edge.
	GroupsTable = "groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "owner_id"
	// IntegrationsTable is the table that holds the integrations relation/edge.
	IntegrationsTable = "integrations"
	// IntegrationsInverseTable is the table name for the Integration entity.
	// It exists in this package in order to avoid circular dependency with the "integration" package.
	IntegrationsInverseTable = "integrations"
	// IntegrationsColumn is the table column denoting the integrations relation/edge.
	IntegrationsColumn = "owner_id"
	// SettingTable is the table that holds the setting relation/edge.
	SettingTable = "organization_settings"
	// SettingInverseTable is the table name for the OrganizationSetting entity.
	// It exists in this package in order to avoid circular dependency with the "organizationsetting" package.
	SettingInverseTable = "organization_settings"
	// SettingColumn is the table column denoting the setting relation/edge.
	SettingColumn = "organization_id"
	// EntitlementsTable is the table that holds the entitlements relation/edge.
	EntitlementsTable = "entitlements"
	// EntitlementsInverseTable is the table name for the Entitlement entity.
	// It exists in this package in order to avoid circular dependency with the "entitlement" package.
	EntitlementsInverseTable = "entitlements"
	// EntitlementsColumn is the table column denoting the entitlements relation/edge.
	EntitlementsColumn = "owner_id"
	// PersonalAccessTokensTable is the table that holds the personal_access_tokens relation/edge. The primary key declared below.
	PersonalAccessTokensTable = "organization_personal_access_tokens"
	// PersonalAccessTokensInverseTable is the table name for the PersonalAccessToken entity.
	// It exists in this package in order to avoid circular dependency with the "personalaccesstoken" package.
	PersonalAccessTokensInverseTable = "personal_access_tokens"
	// OauthproviderTable is the table that holds the oauthprovider relation/edge.
	OauthproviderTable = "oauth_providers"
	// OauthproviderInverseTable is the table name for the OauthProvider entity.
	// It exists in this package in order to avoid circular dependency with the "oauthprovider" package.
	OauthproviderInverseTable = "oauth_providers"
	// OauthproviderColumn is the table column denoting the oauthprovider relation/edge.
	OauthproviderColumn = "organization_oauthprovider"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "org_memberships"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// InvitesTable is the table that holds the invites relation/edge.
	InvitesTable = "invites"
	// InvitesInverseTable is the table name for the Invite entity.
	// It exists in this package in order to avoid circular dependency with the "invite" package.
	InvitesInverseTable = "invites"
	// InvitesColumn is the table column denoting the invites relation/edge.
	InvitesColumn = "owner_id"
	// SubscribersTable is the table that holds the subscribers relation/edge.
	SubscribersTable = "subscribers"
	// SubscribersInverseTable is the table name for the Subscribers entity.
	// It exists in this package in order to avoid circular dependency with the "subscribers" package.
	SubscribersInverseTable = "subscribers"
	// SubscribersColumn is the table column denoting the subscribers relation/edge.
	SubscribersColumn = "owner_id"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "org_memberships"
	// MembersInverseTable is the table name for the OrgMembership entity.
	// It exists in this package in order to avoid circular dependency with the "orgmembership" package.
	MembersInverseTable = "org_memberships"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "organization_id"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldName,
	FieldDisplayName,
	FieldDescription,
	FieldParentOrganizationID,
	FieldPersonalOrg,
	FieldAvatarRemoteURL,
}

var (
	// PersonalAccessTokensPrimaryKey and PersonalAccessTokensColumn2 are the table columns denoting the
	// primary key for the personal_access_tokens relation (M2M).
	PersonalAccessTokensPrimaryKey = []string{"organization_id", "personal_access_token_id"}
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "organization_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/datumforge/datum/internal/ent/generated/runtime"
var (
	Hooks        [4]ent.Hook
	Interceptors [2]ent.Interceptor
	Policy       ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDisplayName holds the default value on creation for the "display_name" field.
	DefaultDisplayName string
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultPersonalOrg holds the default value on creation for the "personal_org" field.
	DefaultPersonalOrg bool
	// AvatarRemoteURLValidator is a validator for the "avatar_remote_url" field. It is called by the builders before save.
	AvatarRemoteURLValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Organization queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByParentOrganizationID orders the results by the parent_organization_id field.
func ByParentOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentOrganizationID, opts...).ToFunc()
}

// ByPersonalOrg orders the results by the personal_org field.
func ByPersonalOrg(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPersonalOrg, opts...).ToFunc()
}

// ByAvatarRemoteURL orders the results by the avatar_remote_url field.
func ByAvatarRemoteURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarRemoteURL, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIntegrationsCount orders the results by integrations count.
func ByIntegrationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIntegrationsStep(), opts...)
	}
}

// ByIntegrations orders the results by integrations terms.
func ByIntegrations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIntegrationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySettingField orders the results by setting field.
func BySettingField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSettingStep(), sql.OrderByField(field, opts...))
	}
}

// ByEntitlementsCount orders the results by entitlements count.
func ByEntitlementsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEntitlementsStep(), opts...)
	}
}

// ByEntitlements orders the results by entitlements terms.
func ByEntitlements(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEntitlementsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPersonalAccessTokensCount orders the results by personal_access_tokens count.
func ByPersonalAccessTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPersonalAccessTokensStep(), opts...)
	}
}

// ByPersonalAccessTokens orders the results by personal_access_tokens terms.
func ByPersonalAccessTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPersonalAccessTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOauthproviderCount orders the results by oauthprovider count.
func ByOauthproviderCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOauthproviderStep(), opts...)
	}
}

// ByOauthprovider orders the results by oauthprovider terms.
func ByOauthprovider(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOauthproviderStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInvitesCount orders the results by invites count.
func ByInvitesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvitesStep(), opts...)
	}
}

// ByInvites orders the results by invites terms.
func ByInvites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvitesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubscribersCount orders the results by subscribers count.
func BySubscribersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscribersStep(), opts...)
	}
}

// BySubscribers orders the results by subscribers terms.
func BySubscribers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscribersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMembersCount orders the results by members count.
func ByMembersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMembersStep(), opts...)
	}
}

// ByMembers orders the results by members terms.
func ByMembers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GroupsTable, GroupsColumn),
	)
}
func newIntegrationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IntegrationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IntegrationsTable, IntegrationsColumn),
	)
}
func newSettingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SettingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SettingTable, SettingColumn),
	)
}
func newEntitlementsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EntitlementsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EntitlementsTable, EntitlementsColumn),
	)
}
func newPersonalAccessTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PersonalAccessTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PersonalAccessTokensTable, PersonalAccessTokensPrimaryKey...),
	)
}
func newOauthproviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OauthproviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OauthproviderTable, OauthproviderColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
func newInvitesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvitesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InvitesTable, InvitesColumn),
	)
}
func newSubscribersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscribersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubscribersTable, SubscribersColumn),
	)
}
func newMembersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, MembersTable, MembersColumn),
	)
}
