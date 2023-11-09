// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "logo_url", Type: field.TypeString},
		{Name: "organization_groups", Type: field.TypeString, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_organizations_groups",
				Columns:    []*schema.Column{GroupsColumns[8]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "group_name_organization_groups",
				Unique:  true,
				Columns: []*schema.Column{GroupsColumns[5], GroupsColumns[8]},
			},
		},
	}
	// GroupSettingsColumns holds the columns for the "group_settings" table.
	GroupSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "visibility", Type: field.TypeEnum, Enums: []string{"PUBLIC", "PRIVATE"}, Default: "PUBLIC"},
		{Name: "join_policy", Type: field.TypeEnum, Enums: []string{"OPEN", "INVITE_ONLY", "APPLICATION_ONLY", "INVITE_OR_APPLICATION"}, Default: "OPEN"},
		{Name: "group_setting", Type: field.TypeString, Unique: true, Nullable: true},
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
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "kind", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "secret_name", Type: field.TypeString},
		{Name: "organization_integrations", Type: field.TypeString, Nullable: true},
	}
	// IntegrationsTable holds the schema information for the "integrations" table.
	IntegrationsTable = &schema.Table{
		Name:       "integrations",
		Columns:    IntegrationsColumns,
		PrimaryKey: []*schema.Column{IntegrationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "integrations_organizations_integrations",
				Columns:    []*schema.Column{IntegrationsColumns[9]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 160},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "parent_organization_id", Type: field.TypeString, Nullable: true},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organizations_organizations_children",
				Columns:    []*schema.Column{OrganizationsColumns[7]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrganizationSettingsColumns holds the columns for the "organization_settings" table.
	OrganizationSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "domains", Type: field.TypeJSON},
		{Name: "sso_cert", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "sso_entrypoint", Type: field.TypeString, Default: ""},
		{Name: "sso_issuer", Type: field.TypeString, Default: ""},
		{Name: "billing_contact", Type: field.TypeString},
		{Name: "billing_email", Type: field.TypeString},
		{Name: "billing_phone", Type: field.TypeString},
		{Name: "billing_address", Type: field.TypeString},
		{Name: "tax_identifier", Type: field.TypeString},
	}
	// OrganizationSettingsTable holds the schema information for the "organization_settings" table.
	OrganizationSettingsTable = &schema.Table{
		Name:       "organization_settings",
		Columns:    OrganizationSettingsColumns,
		PrimaryKey: []*schema.Column{OrganizationSettingsColumns[0]},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "client_id", Type: field.TypeString, Size: 2147483647},
		{Name: "scopes", Type: field.TypeJSON, Nullable: true},
		{Name: "nonce", Type: field.TypeString, Size: 2147483647},
		{Name: "claims_user_id", Type: field.TypeString, Size: 2147483647},
		{Name: "claims_username", Type: field.TypeString, Size: 2147483647},
		{Name: "claims_email", Type: field.TypeString, Size: 2147483647},
		{Name: "claims_email_verified", Type: field.TypeBool},
		{Name: "claims_groups", Type: field.TypeJSON, Nullable: true},
		{Name: "claims_preferred_username", Type: field.TypeString, Size: 2147483647},
		{Name: "connector_id", Type: field.TypeString, Size: 2147483647},
		{Name: "connector_data", Type: field.TypeJSON, Nullable: true},
		{Name: "token", Type: field.TypeString, Size: 2147483647},
		{Name: "obsolete_token", Type: field.TypeString, Size: 2147483647},
		{Name: "last_used", Type: field.TypeTime},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"local", "oauth", "app_password"}},
		{Name: "disabled", Type: field.TypeBool},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "user_agent", Type: field.TypeString, Nullable: true},
		{Name: "ips", Type: field.TypeString},
		{Name: "session_users", Type: field.TypeString, Nullable: true},
		{Name: "user_sessions", Type: field.TypeString, Nullable: true},
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
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
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
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
		},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserOrganizationsColumns holds the columns for the "user_organizations" table.
	UserOrganizationsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString},
		{Name: "organization_id", Type: field.TypeString},
	}
	// UserOrganizationsTable holds the schema information for the "user_organizations" table.
	UserOrganizationsTable = &schema.Table{
		Name:       "user_organizations",
		Columns:    UserOrganizationsColumns,
		PrimaryKey: []*schema.Column{UserOrganizationsColumns[0], UserOrganizationsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_organizations_user_id",
				Columns:    []*schema.Column{UserOrganizationsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_organizations_organization_id",
				Columns:    []*schema.Column{UserOrganizationsColumns[1]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		GroupSettingsTable,
		IntegrationsTable,
		OrganizationsTable,
		OrganizationSettingsTable,
		RefreshTokensTable,
		SessionsTable,
		UsersTable,
		GroupUsersTable,
		UserOrganizationsTable,
	}
)

func init() {
	GroupsTable.ForeignKeys[0].RefTable = OrganizationsTable
	GroupSettingsTable.ForeignKeys[0].RefTable = GroupsTable
	IntegrationsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationsTable.ForeignKeys[0].RefTable = OrganizationsTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.ForeignKeys[1].RefTable = UsersTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
	UserOrganizationsTable.ForeignKeys[0].RefTable = UsersTable
	UserOrganizationsTable.ForeignKeys[1].RefTable = OrganizationsTable
}
