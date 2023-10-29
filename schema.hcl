table "integrations" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "created_by" {
    null = true
    type = integer
  }
  column "updated_by" {
    null = true
    type = integer
  }
  column "kind" {
    null = false
    type = text
  }
  column "description" {
    null = true
    type = text
  }
  column "secret_name" {
    null = false
    type = text
  }
  column "organization_integrations" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "integrations_organizations_integrations" {
    columns     = [column.organization_integrations]
    ref_columns = [table.organizations.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "organizations" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "created_by" {
    null = true
    type = integer
  }
  column "updated_by" {
    null = true
    type = integer
  }
  column "name" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "organizations_name_key" {
    unique  = true
    columns = [column.name]
  }
}
table "sessions" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "created_by" {
    null = true
    type = integer
  }
  column "updated_by" {
    null = true
    type = integer
  }
  column "type" {
    null = false
    type = text
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "token" {
    null = false
    type = text
  }
  column "user_agent" {
    null = true
    type = text
  }
  column "ips" {
    null = false
    type = text
  }
  column "session_users" {
    null = true
    type = uuid
  }
  column "user_sessions" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "sessions_users_sessions" {
    columns     = [column.user_sessions]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "sessions_users_users" {
    columns     = [column.session_users]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
  index "sessions_token_key" {
    unique  = true
    columns = [column.token]
  }
  index "session_id" {
    unique  = true
    columns = [column.id]
  }
}
table "users" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "created_by" {
    null = true
    type = integer
  }
  column "updated_by" {
    null = true
    type = integer
  }
  column "email" {
    null = false
    type = text
  }
  column "first_name" {
    null = false
    type = text
  }
  column "last_name" {
    null = false
    type = text
  }
  column "display_name" {
    null    = false
    type    = text
    default = "unknown"
  }
  column "locked" {
    null    = false
    type    = bool
    default = false
  }
  column "avatar_remote_url" {
    null = true
    type = text
  }
  column "avatar_local_file" {
    null = true
    type = text
  }
  column "avatar_updated_at" {
    null = true
    type = datetime
  }
  column "silenced_at" {
    null = true
    type = datetime
  }
  column "suspended_at" {
    null = true
    type = datetime
  }
  column "recovery_code" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "users_email_key" {
    unique  = true
    columns = [column.email]
  }
  index "user_id" {
    unique  = true
    columns = [column.id]
  }
}
table "memberships" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "created_by" {
    null = true
    type = integer
  }
  column "updated_by" {
    null = true
    type = integer
  }
  column "current" {
    null    = false
    type    = bool
    default = false
  }
  column "group_memberships" {
    null = false
    type = uuid
  }
  column "organization_memberships" {
    null = false
    type = uuid
  }
  column "user_memberships" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "memberships_users_memberships" {
    columns     = [column.user_memberships]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "memberships_organizations_memberships" {
    columns     = [column.organization_memberships]
    ref_columns = [table.organizations.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "memberships_groups_memberships" {
    columns     = [column.group_memberships]
    ref_columns = [table.groups.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "membership_organization_members_57e4e125ad3f56514a7fb2a9105c17d4" {
    unique  = true
    columns = [column.organization_memberships, column.user_memberships, column.group_memberships]
  }
}
table "groups" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "created_by" {
    null = true
    type = integer
  }
  column "updated_by" {
    null = true
    type = integer
  }
  column "name" {
    null = false
    type = text
  }
  column "description" {
    null    = false
    type    = text
    default = ""
  }
  column "logo_url" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "groups_name_key" {
    unique  = true
    columns = [column.name]
  }
}
table "group_settings" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  column "created_by" {
    null = true
    type = integer
  }
  column "updated_by" {
    null = true
    type = integer
  }
  column "visibility" {
    null    = false
    type    = text
    default = "PUBLIC"
  }
  column "join_policy" {
    null    = false
    type    = text
    default = "OPEN"
  }
  column "group_setting" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "group_settings_groups_setting" {
    columns     = [column.group_setting]
    ref_columns = [table.groups.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
  index "group_settings_group_setting_key" {
    unique  = true
    columns = [column.group_setting]
  }
}
table "ent_types" {
  schema = schema.main
  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }
  column "type" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "ent_types_type_key" {
    unique  = true
    columns = [column.type]
  }
}
schema "main" {
}
