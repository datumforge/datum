// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "logo_url", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeUUID},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_tenants_tenant",
				Columns:    []*schema.Column{GroupsColumns[8]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupSettingsColumns holds the columns for the "group_settings" table.
	GroupSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "visibility", Type: field.TypeEnum, Enums: []string{"PUBLIC", "PRIVATE"}, Default: "PUBLIC"},
		{Name: "join_policy", Type: field.TypeEnum, Enums: []string{"OPEN", "INVITE_ONLY", "APPLICATION_ONLY", "INVITE_OR_APPLICATION"}, Default: "OPEN"},
		{Name: "group_setting", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// GroupSettingsTable holds the schema information for the "group_settings" table.
	GroupSettingsTable = &schema.Table{
		Name:       "group_settings",
		Columns:    GroupSettingsColumns,
		PrimaryKey: []*schema.Column{GroupSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_settings_groups_setting",
				Columns:    []*schema.Column{GroupSettingsColumns[7]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// IntegrationsColumns holds the columns for the "integrations" table.
	IntegrationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "kind", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "secret_name", Type: field.TypeString},
		{Name: "organization_integrations", Type: field.TypeUUID},
	}
	// IntegrationsTable holds the schema information for the "integrations" table.
	IntegrationsTable = &schema.Table{
		Name:       "integrations",
		Columns:    IntegrationsColumns,
		PrimaryKey: []*schema.Column{IntegrationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "integrations_organizations_integrations",
				Columns:    []*schema.Column{IntegrationsColumns[8]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MembershipsColumns holds the columns for the "memberships" table.
	MembershipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "current", Type: field.TypeBool, Default: false},
		{Name: "group_memberships", Type: field.TypeUUID},
		{Name: "organization_memberships", Type: field.TypeUUID},
		{Name: "user_memberships", Type: field.TypeUUID},
	}
	// MembershipsTable holds the schema information for the "memberships" table.
	MembershipsTable = &schema.Table{
		Name:       "memberships",
		Columns:    MembershipsColumns,
		PrimaryKey: []*schema.Column{MembershipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "memberships_groups_memberships",
				Columns:    []*schema.Column{MembershipsColumns[6]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "memberships_organizations_memberships",
				Columns:    []*schema.Column{MembershipsColumns[7]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "memberships_users_memberships",
				Columns:    []*schema.Column{MembershipsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "membership_organization_memberships_user_memberships_group_memberships",
				Unique:  true,
				Columns: []*schema.Column{MembershipsColumns[7], MembershipsColumns[8], MembershipsColumns[6]},
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 160},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"local", "oauth", "app_password"}},
		{Name: "disabled", Type: field.TypeBool},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "user_agent", Type: field.TypeString, Nullable: true},
		{Name: "ips", Type: field.TypeString},
		{Name: "session_users", Type: field.TypeUUID, Nullable: true},
		{Name: "user_sessions", Type: field.TypeUUID, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_users",
				Columns:    []*schema.Column{SessionsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "session_id",
				Unique:  true,
				Columns: []*schema.Column{SessionsColumns[0]},
			},
		},
	}
	// TenantsColumns holds the columns for the "tenants" table.
	TenantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
	}
	// TenantsTable holds the schema information for the "tenants" table.
	TenantsTable = &schema.Table{
		Name:       "tenants",
		Columns:    TenantsColumns,
		PrimaryKey: []*schema.Column{TenantsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "first_name", Type: field.TypeString, Size: 64},
		{Name: "last_name", Type: field.TypeString, Size: 64},
		{Name: "display_name", Type: field.TypeString, Size: 64, Default: "unknown"},
		{Name: "locked", Type: field.TypeBool, Default: false},
		{Name: "avatar_remote_url", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "avatar_local_file", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "avatar_updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "silenced_at", Type: field.TypeTime, Nullable: true},
		{Name: "suspended_at", Type: field.TypeTime, Nullable: true},
		{Name: "recovery_code", Type: field.TypeString, Nullable: true},
		{Name: "tenant_id", Type: field.TypeUUID},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_tenants_tenant",
				Columns:    []*schema.Column{UsersColumns[16]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		GroupSettingsTable,
		IntegrationsTable,
		MembershipsTable,
		OrganizationsTable,
		SessionsTable,
		TenantsTable,
		UsersTable,
		UserGroupsTable,
	}
)

func init() {
	GroupsTable.ForeignKeys[0].RefTable = TenantsTable
	GroupSettingsTable.ForeignKeys[0].RefTable = GroupsTable
	IntegrationsTable.ForeignKeys[0].RefTable = OrganizationsTable
	MembershipsTable.ForeignKeys[0].RefTable = GroupsTable
	MembershipsTable.ForeignKeys[1].RefTable = OrganizationsTable
	MembershipsTable.ForeignKeys[2].RefTable = UsersTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = TenantsTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
}
