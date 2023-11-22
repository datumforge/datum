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
    type = uuid
  }
  column "updated_by" {
    null = true
    type = uuid
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
  column "organization_groups" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "groups_organizations_groups" {
    columns     = [column.organization_groups]
    ref_columns = [table.organizations.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "group_name_organization_groups" {
    unique  = true
    columns = [column.name, column.organization_groups]
  }
}
table "group_setting" {
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
    type = uuid
  }
  column "updated_by" {
    null = true
    type = uuid
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
  foreign_key "group_setting_groups_setting" {
    columns     = [column.group_setting]
    ref_columns = [table.groups.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
  index "group_setting_group_setting_key" {
    unique  = true
    columns = [column.group_setting]
  }
}
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
    type = uuid
  }
  column "updated_by" {
    null = true
    type = uuid
  }
  column "name" {
    null = false
    type = text
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
    null = true
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
    type = uuid
  }
  column "updated_by" {
    null = true
    type = uuid
  }
  column "name" {
    null = false
    type = text
  }
  column "description" {
    null = true
    type = text
  }
  column "parent_organization_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "organizations_organizations_children" {
    columns     = [column.parent_organization_id]
    ref_columns = [table.organizations.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
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
    type = uuid
  }
  column "updated_by" {
    null = true
    type = uuid
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
    type = uuid
  }
  column "updated_by" {
    null = true
    type = uuid
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
table "group_users" {
  schema = schema.main
  column "group_id" {
    null = false
    type = uuid
  }
  column "user_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.group_id, column.user_id]
  }
  foreign_key "group_users_user_id" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "group_users_group_id" {
    columns     = [column.group_id]
    ref_columns = [table.groups.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "user_organizations" {
  schema = schema.main
  column "user_id" {
    null = false
    type = uuid
  }
  column "organization_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.user_id, column.organization_id]
  }
  foreign_key "user_organizations_organization_id" {
    columns     = [column.organization_id]
    ref_columns = [table.organizations.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "user_organizations_user_id" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
schema "main" {
}
