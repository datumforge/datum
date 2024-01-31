// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EmailVerificationTokensColumns holds the columns for the "email_verification_tokens" table.
	EmailVerificationTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "ttl", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString},
		{Name: "secret", Type: field.TypeBytes},
		{Name: "owner_id", Type: field.TypeString},
	}
	// EmailVerificationTokensTable holds the schema information for the "email_verification_tokens" table.
	EmailVerificationTokensTable = &schema.Table{
		Name:       "email_verification_tokens",
		Columns:    EmailVerificationTokensColumns,
		PrimaryKey: []*schema.Column{EmailVerificationTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "email_verification_tokens_users_email_verification_tokens",
				Columns:    []*schema.Column{EmailVerificationTokensColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "emailverificationtoken_token",
				Unique:  true,
				Columns: []*schema.Column{EmailVerificationTokensColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
		},
	}
	// EntitlementsColumns holds the columns for the "entitlements" table.
	EntitlementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "tier", Type: field.TypeEnum, Enums: []string{"FREE", "PRO", "ENTERPRISE"}, Default: "PRO"},
		{Name: "external_customer_id", Type: field.TypeString, Nullable: true},
		{Name: "external_subscription_id", Type: field.TypeString, Nullable: true},
		{Name: "expires", Type: field.TypeBool, Default: false},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "cancelled", Type: field.TypeBool, Default: false},
		{Name: "owner_id", Type: field.TypeString},
	}
	// EntitlementsTable holds the schema information for the "entitlements" table.
	EntitlementsTable = &schema.Table{
		Name:       "entitlements",
		Columns:    EntitlementsColumns,
		PrimaryKey: []*schema.Column{EntitlementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entitlements_organizations_entitlements",
				Columns:    []*schema.Column{EntitlementsColumns[13]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "gravatar_logo_url", Type: field.TypeString, Nullable: true},
		{Name: "logo_url", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "owner_id", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_organizations_groups",
				Columns:    []*schema.Column{GroupsColumns[12]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "group_name_owner_id",
				Unique:  true,
				Columns: []*schema.Column{GroupsColumns[7], GroupsColumns[12]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
		},
	}
	// GroupMembershipsColumns holds the columns for the "group_memberships" table.
	GroupMembershipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "MEMBER"}, Default: "MEMBER"},
		{Name: "group_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// GroupMembershipsTable holds the schema information for the "group_memberships" table.
	GroupMembershipsTable = &schema.Table{
		Name:       "group_memberships",
		Columns:    GroupMembershipsColumns,
		PrimaryKey: []*schema.Column{GroupMembershipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_memberships_groups_group",
				Columns:    []*schema.Column{GroupMembershipsColumns[8]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_memberships_users_user",
				Columns:    []*schema.Column{GroupMembershipsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "groupmembership_user_id_group_id",
				Unique:  true,
				Columns: []*schema.Column{GroupMembershipsColumns[9], GroupMembershipsColumns[8]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
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
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "visibility", Type: field.TypeEnum, Enums: []string{"PUBLIC", "PRIVATE"}, Default: "PUBLIC"},
		{Name: "join_policy", Type: field.TypeEnum, Enums: []string{"OPEN", "INVITE_ONLY", "APPLICATION_ONLY", "INVITE_OR_APPLICATION"}, Default: "INVITE_OR_APPLICATION"},
		{Name: "tags", Type: field.TypeJSON},
		{Name: "sync_to_slack", Type: field.TypeBool, Default: false},
		{Name: "sync_to_github", Type: field.TypeBool, Default: false},
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
				Columns:    []*schema.Column{GroupSettingsColumns[12]},
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
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "kind", Type: field.TypeString, Nullable: true},
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
				Columns:    []*schema.Column{IntegrationsColumns[11]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InvitesColumns holds the columns for the "invites" table.
	InvitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "expires", Type: field.TypeTime},
		{Name: "recipient", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"INVITATION_SENT", "APPROVAL_REQUIRED", "INVITATION_ACCEPTED", "INVITATION_EXPIRED"}, Default: "INVITATION_SENT"},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "MEMBER", "OWNER"}, Default: "MEMBER"},
		{Name: "send_attempts", Type: field.TypeInt, Default: 0},
		{Name: "requestor_id", Type: field.TypeString},
		{Name: "secret", Type: field.TypeBytes},
		{Name: "owner_id", Type: field.TypeString},
	}
	// InvitesTable holds the schema information for the "invites" table.
	InvitesTable = &schema.Table{
		Name:       "invites",
		Columns:    InvitesColumns,
		PrimaryKey: []*schema.Column{InvitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invites_organizations_invites",
				Columns:    []*schema.Column{InvitesColumns[15]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "invite_recipient_owner_id",
				Unique:  true,
				Columns: []*schema.Column{InvitesColumns[9], InvitesColumns[15]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
		},
	}
	// OauthProvidersColumns holds the columns for the "oauth_providers" table.
	OauthProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "client_id", Type: field.TypeString},
		{Name: "client_secret", Type: field.TypeString},
		{Name: "redirect_url", Type: field.TypeString},
		{Name: "scopes", Type: field.TypeString},
		{Name: "auth_url", Type: field.TypeString},
		{Name: "token_url", Type: field.TypeString},
		{Name: "auth_style", Type: field.TypeUint8},
		{Name: "info_url", Type: field.TypeString},
		{Name: "organization_oauthprovider", Type: field.TypeString, Nullable: true},
	}
	// OauthProvidersTable holds the schema information for the "oauth_providers" table.
	OauthProvidersTable = &schema.Table{
		Name:       "oauth_providers",
		Columns:    OauthProvidersColumns,
		PrimaryKey: []*schema.Column{OauthProvidersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "oauth_providers_organizations_oauthprovider",
				Columns:    []*schema.Column{OauthProvidersColumns[16]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OhAuthTooTokensColumns holds the columns for the "oh_auth_too_tokens" table.
	OhAuthTooTokensColumns = []*schema.Column{
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
		{Name: "last_used", Type: field.TypeTime},
	}
	// OhAuthTooTokensTable holds the schema information for the "oh_auth_too_tokens" table.
	OhAuthTooTokensTable = &schema.Table{
		Name:       "oh_auth_too_tokens",
		Columns:    OhAuthTooTokensColumns,
		PrimaryKey: []*schema.Column{OhAuthTooTokensColumns[0]},
	}
	// OrgMembershipsColumns holds the columns for the "org_memberships" table.
	OrgMembershipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "MEMBER", "OWNER"}, Default: "MEMBER"},
		{Name: "organization_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// OrgMembershipsTable holds the schema information for the "org_memberships" table.
	OrgMembershipsTable = &schema.Table{
		Name:       "org_memberships",
		Columns:    OrgMembershipsColumns,
		PrimaryKey: []*schema.Column{OrgMembershipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "org_memberships_organizations_organization",
				Columns:    []*schema.Column{OrgMembershipsColumns[8]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "org_memberships_users_user",
				Columns:    []*schema.Column{OrgMembershipsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "orgmembership_user_id_organization_id",
				Unique:  true,
				Columns: []*schema.Column{OrgMembershipsColumns[9], OrgMembershipsColumns[8]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
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
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 160},
		{Name: "display_name", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "personal_org", Type: field.TypeBool, Default: false},
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
				Columns:    []*schema.Column{OrganizationsColumns[11]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "organization_name",
				Unique:  true,
				Columns: []*schema.Column{OrganizationsColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
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
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "domains", Type: field.TypeJSON, Nullable: true},
		{Name: "sso_cert", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "sso_entrypoint", Type: field.TypeString, Nullable: true},
		{Name: "sso_issuer", Type: field.TypeString, Nullable: true},
		{Name: "billing_contact", Type: field.TypeString, Nullable: true},
		{Name: "billing_email", Type: field.TypeString, Nullable: true},
		{Name: "billing_phone", Type: field.TypeString, Nullable: true},
		{Name: "billing_address", Type: field.TypeString, Nullable: true},
		{Name: "tax_identifier", Type: field.TypeString, Nullable: true},
		{Name: "tags", Type: field.TypeJSON, Nullable: true},
		{Name: "organization_setting", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// OrganizationSettingsTable holds the schema information for the "organization_settings" table.
	OrganizationSettingsTable = &schema.Table{
		Name:       "organization_settings",
		Columns:    OrganizationSettingsColumns,
		PrimaryKey: []*schema.Column{OrganizationSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_settings_organizations_setting",
				Columns:    []*schema.Column{OrganizationSettingsColumns[17]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PasswordResetTokensColumns holds the columns for the "password_reset_tokens" table.
	PasswordResetTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "ttl", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString},
		{Name: "secret", Type: field.TypeBytes},
		{Name: "owner_id", Type: field.TypeString},
	}
	// PasswordResetTokensTable holds the schema information for the "password_reset_tokens" table.
	PasswordResetTokensTable = &schema.Table{
		Name:       "password_reset_tokens",
		Columns:    PasswordResetTokensColumns,
		PrimaryKey: []*schema.Column{PasswordResetTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "password_reset_tokens_users_password_reset_tokens",
				Columns:    []*schema.Column{PasswordResetTokensColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "passwordresettoken_token",
				Unique:  true,
				Columns: []*schema.Column{PasswordResetTokensColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
		},
	}
	// PersonalAccessTokensColumns holds the columns for the "personal_access_tokens" table.
	PersonalAccessTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "abilities", Type: field.TypeJSON, Nullable: true},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "last_used_at", Type: field.TypeTime, Nullable: true},
		{Name: "owner_id", Type: field.TypeString},
	}
	// PersonalAccessTokensTable holds the schema information for the "personal_access_tokens" table.
	PersonalAccessTokensTable = &schema.Table{
		Name:       "personal_access_tokens",
		Columns:    PersonalAccessTokensColumns,
		PrimaryKey: []*schema.Column{PersonalAccessTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "personal_access_tokens_users_personal_access_tokens",
				Columns:    []*schema.Column{PersonalAccessTokensColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "personalaccesstoken_token",
				Unique:  false,
				Columns: []*schema.Column{PersonalAccessTokensColumns[8]},
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
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString, Size: 64},
		{Name: "last_name", Type: field.TypeString, Size: 64},
		{Name: "display_name", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "avatar_remote_uri", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "avatar_local_file", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "avatar_updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_seen", Type: field.TypeTime, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "sub", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "oauth", Type: field.TypeBool, Default: false},
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
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[7]},
				Annotation: &entsql.IndexAnnotation{
					Where: "deleted_at is NULL",
				},
			},
		},
	}
	// UserSettingsColumns holds the columns for the "user_settings" table.
	UserSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "locked", Type: field.TypeBool, Default: false},
		{Name: "silenced_at", Type: field.TypeTime, Nullable: true},
		{Name: "suspended_at", Type: field.TypeTime, Nullable: true},
		{Name: "recovery_code", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACTIVE", "INACTIVE", "DEACTIVATED", "SUSPENDED"}, Default: "ACTIVE"},
		{Name: "default_org", Type: field.TypeString, Nullable: true},
		{Name: "email_confirmed", Type: field.TypeBool, Default: false},
		{Name: "tags", Type: field.TypeJSON},
		{Name: "user_setting", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// UserSettingsTable holds the schema information for the "user_settings" table.
	UserSettingsTable = &schema.Table{
		Name:       "user_settings",
		Columns:    UserSettingsColumns,
		PrimaryKey: []*schema.Column{UserSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_settings_users_setting",
				Columns:    []*schema.Column{UserSettingsColumns[15]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EmailVerificationTokensTable,
		EntitlementsTable,
		GroupsTable,
		GroupMembershipsTable,
		GroupSettingsTable,
		IntegrationsTable,
		InvitesTable,
		OauthProvidersTable,
		OhAuthTooTokensTable,
		OrgMembershipsTable,
		OrganizationsTable,
		OrganizationSettingsTable,
		PasswordResetTokensTable,
		PersonalAccessTokensTable,
		UsersTable,
		UserSettingsTable,
	}
)

func init() {
	EmailVerificationTokensTable.ForeignKeys[0].RefTable = UsersTable
	EntitlementsTable.ForeignKeys[0].RefTable = OrganizationsTable
	GroupsTable.ForeignKeys[0].RefTable = OrganizationsTable
	GroupMembershipsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupMembershipsTable.ForeignKeys[1].RefTable = UsersTable
	GroupSettingsTable.ForeignKeys[0].RefTable = GroupsTable
	IntegrationsTable.ForeignKeys[0].RefTable = OrganizationsTable
	InvitesTable.ForeignKeys[0].RefTable = OrganizationsTable
	OauthProvidersTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrgMembershipsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrgMembershipsTable.ForeignKeys[1].RefTable = UsersTable
	OrganizationsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationSettingsTable.ForeignKeys[0].RefTable = OrganizationsTable
	PasswordResetTokensTable.ForeignKeys[0].RefTable = UsersTable
	PersonalAccessTokensTable.ForeignKeys[0].RefTable = UsersTable
	UserSettingsTable.ForeignKeys[0].RefTable = UsersTable
}
