-- Create "org_membership_history" table
CREATE TABLE "org_membership_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "role" character varying NOT NULL DEFAULT 'MEMBER', "organization_id" character varying NOT NULL, "user_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "orgmembershiphistory_history_time" to table: "org_membership_history"
CREATE INDEX "orgmembershiphistory_history_time" ON "org_membership_history" ("history_time");
-- Create "file_history" table
CREATE TABLE "file_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "file_name" character varying NOT NULL, "file_extension" character varying NOT NULL, "file_size" bigint NULL, "content_type" character varying NOT NULL, "store_key" character varying NOT NULL, "category" character varying NULL, "annotation" character varying NULL, PRIMARY KEY ("id"));
-- Create index "filehistory_history_time" to table: "file_history"
CREATE INDEX "filehistory_history_time" ON "file_history" ("history_time");
-- Create "document_data_history" table
CREATE TABLE "document_data_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "template_id" character varying NOT NULL, "data" jsonb NOT NULL, PRIMARY KEY ("id"));
-- Create index "documentdatahistory_history_time" to table: "document_data_history"
CREATE INDEX "documentdatahistory_history_time" ON "document_data_history" ("history_time");
-- Create "template_history" table
CREATE TABLE "template_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "owner_id" character varying NULL, "name" character varying NOT NULL, "template_type" character varying NOT NULL DEFAULT 'DOCUMENT', "description" character varying NULL, "jsonconfig" jsonb NOT NULL, "uischema" jsonb NULL, PRIMARY KEY ("id"));
-- Create index "templatehistory_history_time" to table: "template_history"
CREATE INDEX "templatehistory_history_time" ON "template_history" ("history_time");
-- Create "hush_history" table
CREATE TABLE "hush_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "kind" character varying NULL, "secret_name" character varying NULL, "secret_value" character varying NULL, PRIMARY KEY ("id"));
-- Create index "hushhistory_history_time" to table: "hush_history"
CREATE INDEX "hushhistory_history_time" ON "hush_history" ("history_time");
-- Create "entitlement_history" table
CREATE TABLE "entitlement_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "owner_id" character varying NULL, "tier" character varying NOT NULL DEFAULT 'FREE', "external_customer_id" character varying NULL, "external_subscription_id" character varying NULL, "expires" boolean NOT NULL DEFAULT false, "expires_at" timestamptz NULL, "cancelled" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id"));
-- Create index "entitlementhistory_history_time" to table: "entitlement_history"
CREATE INDEX "entitlementhistory_history_time" ON "entitlement_history" ("history_time");
-- Create "user_history" table
CREATE TABLE "user_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "email" character varying NOT NULL, "first_name" character varying NULL, "last_name" character varying NULL, "display_name" character varying NOT NULL, "avatar_remote_url" character varying NULL, "avatar_local_file" character varying NULL, "avatar_updated_at" timestamptz NULL, "last_seen" timestamptz NULL, "password" character varying NULL, "sub" character varying NULL, "auth_provider" character varying NOT NULL DEFAULT 'CREDENTIALS', "role" character varying NULL DEFAULT 'USER', PRIMARY KEY ("id"));
-- Create index "userhistory_history_time" to table: "user_history"
CREATE INDEX "userhistory_history_time" ON "user_history" ("history_time");
-- Create "event_history" table
CREATE TABLE "event_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "event_id" character varying NULL, "correlation_id" character varying NULL, "event_type" character varying NOT NULL, "metadata" jsonb NULL, PRIMARY KEY ("id"));
-- Create index "eventhistory_history_time" to table: "event_history"
CREATE INDEX "eventhistory_history_time" ON "event_history" ("history_time");
-- Create "user_setting_history" table
CREATE TABLE "user_setting_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "user_id" character varying NULL, "locked" boolean NOT NULL DEFAULT false, "silenced_at" timestamptz NULL, "suspended_at" timestamptz NULL, "status" character varying NOT NULL DEFAULT 'ACTIVE', "email_confirmed" boolean NOT NULL DEFAULT false, "tags" jsonb NULL, "is_webauthn_allowed" boolean NULL DEFAULT false, "is_tfa_enabled" boolean NULL DEFAULT false, "phone_number" character varying NULL, PRIMARY KEY ("id"));
-- Create index "usersettinghistory_history_time" to table: "user_setting_history"
CREATE INDEX "usersettinghistory_history_time" ON "user_setting_history" ("history_time");
-- Create "feature_history" table
CREATE TABLE "feature_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "name" character varying NOT NULL, "global" boolean NOT NULL DEFAULT true, "enabled" boolean NOT NULL DEFAULT false, "description" character varying NULL, PRIMARY KEY ("id"));
-- Create index "featurehistory_history_time" to table: "feature_history"
CREATE INDEX "featurehistory_history_time" ON "feature_history" ("history_time");
-- Create "webhook_history" table
CREATE TABLE "webhook_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "owner_id" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "destination_url" character varying NOT NULL, "enabled" boolean NOT NULL DEFAULT true, "callback" character varying NULL, "expires_at" timestamptz NULL, "secret" bytea NULL, "failures" bigint NULL DEFAULT 0, "last_error" character varying NULL, "last_response" character varying NULL, PRIMARY KEY ("id"));
-- Create index "webhookhistory_history_time" to table: "webhook_history"
CREATE INDEX "webhookhistory_history_time" ON "webhook_history" ("history_time");
-- Create "oauth_provider_history" table
CREATE TABLE "oauth_provider_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "client_id" character varying NOT NULL, "client_secret" character varying NOT NULL, "redirect_url" character varying NOT NULL, "scopes" character varying NOT NULL, "auth_url" character varying NOT NULL, "token_url" character varying NOT NULL, "auth_style" smallint NOT NULL, "info_url" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "oauthproviderhistory_history_time" to table: "oauth_provider_history"
CREATE INDEX "oauthproviderhistory_history_time" ON "oauth_provider_history" ("history_time");
-- Create "organization_setting_history" table
CREATE TABLE "organization_setting_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "domains" jsonb NULL, "billing_contact" character varying NULL, "billing_email" character varying NULL, "billing_phone" character varying NULL, "billing_address" character varying NULL, "tax_identifier" character varying NULL, "tags" jsonb NULL, "geo_location" character varying NULL DEFAULT 'AMER', "organization_id" character varying NULL, PRIMARY KEY ("id"));
-- Create index "organizationsettinghistory_history_time" to table: "organization_setting_history"
CREATE INDEX "organizationsettinghistory_history_time" ON "organization_setting_history" ("history_time");
-- Create "group_history" table
CREATE TABLE "group_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "owner_id" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "gravatar_logo_url" character varying NULL, "logo_url" character varying NULL, "display_name" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- Create index "grouphistory_history_time" to table: "group_history"
CREATE INDEX "grouphistory_history_time" ON "group_history" ("history_time");
-- Create "organization_history" table
CREATE TABLE "organization_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "display_name" character varying NOT NULL DEFAULT '', "description" character varying NULL, "parent_organization_id" character varying NULL, "personal_org" boolean NULL DEFAULT false, "avatar_remote_url" character varying NULL, "dedicated_db" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id"));
-- Create index "organizationhistory_history_time" to table: "organization_history"
CREATE INDEX "organizationhistory_history_time" ON "organization_history" ("history_time");
-- Create "group_membership_history" table
CREATE TABLE "group_membership_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "role" character varying NOT NULL DEFAULT 'MEMBER', "group_id" character varying NOT NULL, "user_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "groupmembershiphistory_history_time" to table: "group_membership_history"
CREATE INDEX "groupmembershiphistory_history_time" ON "group_membership_history" ("history_time");
-- Create "integration_history" table
CREATE TABLE "integration_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "owner_id" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "kind" character varying NULL, PRIMARY KEY ("id"));
-- Create index "integrationhistory_history_time" to table: "integration_history"
CREATE INDEX "integrationhistory_history_time" ON "integration_history" ("history_time");
-- Create "group_setting_history" table
CREATE TABLE "group_setting_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "visibility" character varying NOT NULL DEFAULT 'PUBLIC', "join_policy" character varying NOT NULL DEFAULT 'INVITE_OR_APPLICATION', "tags" jsonb NULL, "sync_to_slack" boolean NULL DEFAULT false, "sync_to_github" boolean NULL DEFAULT false, "group_id" character varying NULL, PRIMARY KEY ("id"));
-- Create index "groupsettinghistory_history_time" to table: "group_setting_history"
CREATE INDEX "groupsettinghistory_history_time" ON "group_setting_history" ("history_time");
-- Create "organizations" table
CREATE TABLE "organizations" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "display_name" character varying NOT NULL DEFAULT '', "description" character varying NULL, "personal_org" boolean NULL DEFAULT false, "avatar_remote_url" character varying NULL, "dedicated_db" boolean NOT NULL DEFAULT false, "parent_organization_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "organizations_organizations_children" FOREIGN KEY ("parent_organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "organization_name" to table: "organizations"
CREATE UNIQUE INDEX "organization_name" ON "organizations" ("name") WHERE (deleted_at IS NULL);
-- Create index "organizations_mapping_id_key" to table: "organizations"
CREATE UNIQUE INDEX "organizations_mapping_id_key" ON "organizations" ("mapping_id");
-- Create "api_tokens" table
CREATE TABLE "api_tokens" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "name" character varying NOT NULL, "token" character varying NOT NULL, "expires_at" timestamptz NULL, "description" character varying NULL, "scopes" jsonb NULL, "last_used_at" timestamptz NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "api_tokens_organizations_api_tokens" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "api_tokens_mapping_id_key" to table: "api_tokens"
CREATE UNIQUE INDEX "api_tokens_mapping_id_key" ON "api_tokens" ("mapping_id");
-- Create index "api_tokens_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX "api_tokens_token_key" ON "api_tokens" ("token");
-- Create index "apitoken_token" to table: "api_tokens"
CREATE INDEX "apitoken_token" ON "api_tokens" ("token");
-- Create "templates" table
CREATE TABLE "templates" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "name" character varying NOT NULL, "template_type" character varying NOT NULL DEFAULT 'DOCUMENT', "description" character varying NULL, "jsonconfig" jsonb NOT NULL, "uischema" jsonb NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "templates_organizations_templates" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "template_name_owner_id_template_type" to table: "templates"
CREATE UNIQUE INDEX "template_name_owner_id_template_type" ON "templates" ("name", "owner_id", "template_type") WHERE (deleted_at IS NULL);
-- Create index "templates_mapping_id_key" to table: "templates"
CREATE UNIQUE INDEX "templates_mapping_id_key" ON "templates" ("mapping_id");
-- Create "document_data" table
CREATE TABLE "document_data" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "data" jsonb NOT NULL, "template_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "document_data_templates_documents" FOREIGN KEY ("template_id") REFERENCES "templates" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "document_data_mapping_id_key" to table: "document_data"
CREATE UNIQUE INDEX "document_data_mapping_id_key" ON "document_data" ("mapping_id");
-- Create "users" table
CREATE TABLE "users" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "email" character varying NOT NULL, "first_name" character varying NULL, "last_name" character varying NULL, "display_name" character varying NOT NULL, "avatar_remote_url" character varying NULL, "avatar_local_file" character varying NULL, "avatar_updated_at" timestamptz NULL, "last_seen" timestamptz NULL, "password" character varying NULL, "sub" character varying NULL, "auth_provider" character varying NOT NULL DEFAULT 'CREDENTIALS', "role" character varying NULL DEFAULT 'USER', PRIMARY KEY ("id"));
-- Create index "user_email_auth_provider" to table: "users"
CREATE UNIQUE INDEX "user_email_auth_provider" ON "users" ("email", "auth_provider") WHERE (deleted_at IS NULL);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX "user_id" ON "users" ("id");
-- Create index "users_mapping_id_key" to table: "users"
CREATE UNIQUE INDEX "users_mapping_id_key" ON "users" ("mapping_id");
-- Create index "users_sub_key" to table: "users"
CREATE UNIQUE INDEX "users_sub_key" ON "users" ("sub");
-- Create "email_verification_tokens" table
CREATE TABLE "email_verification_tokens" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "token" character varying NOT NULL, "ttl" timestamptz NOT NULL, "email" character varying NOT NULL, "secret" bytea NOT NULL, "owner_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "email_verification_tokens_users_email_verification_tokens" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "email_verification_tokens_mapping_id_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX "email_verification_tokens_mapping_id_key" ON "email_verification_tokens" ("mapping_id");
-- Create index "email_verification_tokens_token_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX "email_verification_tokens_token_key" ON "email_verification_tokens" ("token");
-- Create index "emailverificationtoken_token" to table: "email_verification_tokens"
CREATE UNIQUE INDEX "emailverificationtoken_token" ON "email_verification_tokens" ("token") WHERE (deleted_at IS NULL);
-- Create "entitlements" table
CREATE TABLE "entitlements" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tier" character varying NOT NULL DEFAULT 'FREE', "external_customer_id" character varying NULL, "external_subscription_id" character varying NULL, "expires" boolean NOT NULL DEFAULT false, "expires_at" timestamptz NULL, "cancelled" boolean NOT NULL DEFAULT false, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "entitlements_organizations_entitlements" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "entitlements_mapping_id_key" to table: "entitlements"
CREATE UNIQUE INDEX "entitlements_mapping_id_key" ON "entitlements" ("mapping_id");
-- Create "events" table
CREATE TABLE "events" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "event_id" character varying NULL, "correlation_id" character varying NULL, "event_type" character varying NOT NULL, "metadata" jsonb NULL, PRIMARY KEY ("id"));
-- Create index "events_mapping_id_key" to table: "events"
CREATE UNIQUE INDEX "events_mapping_id_key" ON "events" ("mapping_id");
-- Create "entitlement_events" table
CREATE TABLE "entitlement_events" ("entitlement_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("entitlement_id", "event_id"), CONSTRAINT "entitlement_events_entitlement_id" FOREIGN KEY ("entitlement_id") REFERENCES "entitlements" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "entitlement_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "features" table
CREATE TABLE "features" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "name" character varying NOT NULL, "global" boolean NOT NULL DEFAULT true, "enabled" boolean NOT NULL DEFAULT false, "description" character varying NULL, PRIMARY KEY ("id"));
-- Create index "features_mapping_id_key" to table: "features"
CREATE UNIQUE INDEX "features_mapping_id_key" ON "features" ("mapping_id");
-- Create index "features_name_key" to table: "features"
CREATE UNIQUE INDEX "features_name_key" ON "features" ("name");
-- Create "entitlement_features" table
CREATE TABLE "entitlement_features" ("entitlement_id" character varying NOT NULL, "feature_id" character varying NOT NULL, PRIMARY KEY ("entitlement_id", "feature_id"), CONSTRAINT "entitlement_features_entitlement_id" FOREIGN KEY ("entitlement_id") REFERENCES "entitlements" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "entitlement_features_feature_id" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "feature_events" table
CREATE TABLE "feature_events" ("feature_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("feature_id", "event_id"), CONSTRAINT "feature_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "feature_events_feature_id" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "files" table
CREATE TABLE "files" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "file_name" character varying NOT NULL, "file_extension" character varying NOT NULL, "file_size" bigint NULL, "content_type" character varying NOT NULL, "store_key" character varying NOT NULL, "category" character varying NULL, "annotation" character varying NULL, "user_files" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "files_users_files" FOREIGN KEY ("user_files") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "files_mapping_id_key" to table: "files"
CREATE UNIQUE INDEX "files_mapping_id_key" ON "files" ("mapping_id");
-- Create "groups" table
CREATE TABLE "groups" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "name" character varying NOT NULL, "description" character varying NULL, "gravatar_logo_url" character varying NULL, "logo_url" character varying NULL, "display_name" character varying NOT NULL DEFAULT '', "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "groups_organizations_groups" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "group_name_owner_id" to table: "groups"
CREATE UNIQUE INDEX "group_name_owner_id" ON "groups" ("name", "owner_id") WHERE (deleted_at IS NULL);
-- Create index "groups_mapping_id_key" to table: "groups"
CREATE UNIQUE INDEX "groups_mapping_id_key" ON "groups" ("mapping_id");
-- Create "group_events" table
CREATE TABLE "group_events" ("group_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("group_id", "event_id"), CONSTRAINT "group_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "group_events_group_id" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "group_features" table
CREATE TABLE "group_features" ("group_id" character varying NOT NULL, "feature_id" character varying NOT NULL, PRIMARY KEY ("group_id", "feature_id"), CONSTRAINT "group_features_feature_id" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "group_features_group_id" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "group_files" table
CREATE TABLE "group_files" ("group_id" character varying NOT NULL, "file_id" character varying NOT NULL, PRIMARY KEY ("group_id", "file_id"), CONSTRAINT "group_files_file_id" FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "group_files_group_id" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "group_memberships" table
CREATE TABLE "group_memberships" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "role" character varying NOT NULL DEFAULT 'MEMBER', "group_id" character varying NOT NULL, "user_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "group_memberships_groups_group" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "group_memberships_users_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "group_memberships_mapping_id_key" to table: "group_memberships"
CREATE UNIQUE INDEX "group_memberships_mapping_id_key" ON "group_memberships" ("mapping_id");
-- Create index "groupmembership_user_id_group_id" to table: "group_memberships"
CREATE UNIQUE INDEX "groupmembership_user_id_group_id" ON "group_memberships" ("user_id", "group_id") WHERE (deleted_at IS NULL);
-- Create "group_membership_events" table
CREATE TABLE "group_membership_events" ("group_membership_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("group_membership_id", "event_id"), CONSTRAINT "group_membership_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "group_membership_events_group_membership_id" FOREIGN KEY ("group_membership_id") REFERENCES "group_memberships" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "group_settings" table
CREATE TABLE "group_settings" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "visibility" character varying NOT NULL DEFAULT 'PUBLIC', "join_policy" character varying NOT NULL DEFAULT 'INVITE_OR_APPLICATION', "tags" jsonb NULL, "sync_to_slack" boolean NULL DEFAULT false, "sync_to_github" boolean NULL DEFAULT false, "group_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "group_settings_groups_setting" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "group_settings_group_id_key" to table: "group_settings"
CREATE UNIQUE INDEX "group_settings_group_id_key" ON "group_settings" ("group_id");
-- Create index "group_settings_mapping_id_key" to table: "group_settings"
CREATE UNIQUE INDEX "group_settings_mapping_id_key" ON "group_settings" ("mapping_id");
-- Create "hushes" table
CREATE TABLE "hushes" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "kind" character varying NULL, "secret_name" character varying NULL, "secret_value" character varying NULL, PRIMARY KEY ("id"));
-- Create index "hushes_mapping_id_key" to table: "hushes"
CREATE UNIQUE INDEX "hushes_mapping_id_key" ON "hushes" ("mapping_id");
-- Create "hush_events" table
CREATE TABLE "hush_events" ("hush_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("hush_id", "event_id"), CONSTRAINT "hush_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "hush_events_hush_id" FOREIGN KEY ("hush_id") REFERENCES "hushes" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "integrations" table
CREATE TABLE "integrations" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "kind" character varying NULL, "group_integrations" character varying NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "integrations_groups_integrations" FOREIGN KEY ("group_integrations") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "integrations_organizations_integrations" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "integrations_mapping_id_key" to table: "integrations"
CREATE UNIQUE INDEX "integrations_mapping_id_key" ON "integrations" ("mapping_id");
-- Create "integration_events" table
CREATE TABLE "integration_events" ("integration_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("integration_id", "event_id"), CONSTRAINT "integration_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "integration_events_integration_id" FOREIGN KEY ("integration_id") REFERENCES "integrations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "oh_auth_too_tokens" table
CREATE TABLE "oh_auth_too_tokens" ("id" character varying NOT NULL, "mapping_id" character varying NOT NULL, "client_id" text NOT NULL, "scopes" jsonb NULL, "nonce" text NOT NULL, "claims_user_id" text NOT NULL, "claims_username" text NOT NULL, "claims_email" text NOT NULL, "claims_email_verified" boolean NOT NULL, "claims_groups" jsonb NULL, "claims_preferred_username" text NOT NULL, "connector_id" text NOT NULL, "connector_data" jsonb NULL, "last_used" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "oh_auth_too_tokens_mapping_id_key" to table: "oh_auth_too_tokens"
CREATE UNIQUE INDEX "oh_auth_too_tokens_mapping_id_key" ON "oh_auth_too_tokens" ("mapping_id");
-- Create "integration_oauth2tokens" table
CREATE TABLE "integration_oauth2tokens" ("integration_id" character varying NOT NULL, "oh_auth_too_token_id" character varying NOT NULL, PRIMARY KEY ("integration_id", "oh_auth_too_token_id"), CONSTRAINT "integration_oauth2tokens_integration_id" FOREIGN KEY ("integration_id") REFERENCES "integrations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "integration_oauth2tokens_oh_auth_too_token_id" FOREIGN KEY ("oh_auth_too_token_id") REFERENCES "oh_auth_too_tokens" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "integration_secrets" table
CREATE TABLE "integration_secrets" ("integration_id" character varying NOT NULL, "hush_id" character varying NOT NULL, PRIMARY KEY ("integration_id", "hush_id"), CONSTRAINT "integration_secrets_hush_id" FOREIGN KEY ("hush_id") REFERENCES "hushes" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "integration_secrets_integration_id" FOREIGN KEY ("integration_id") REFERENCES "integrations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "invites" table
CREATE TABLE "invites" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "token" character varying NOT NULL, "expires" timestamptz NOT NULL, "recipient" character varying NOT NULL, "status" character varying NOT NULL DEFAULT 'INVITATION_SENT', "role" character varying NOT NULL DEFAULT 'MEMBER', "send_attempts" bigint NOT NULL DEFAULT 0, "requestor_id" character varying NOT NULL, "secret" bytea NOT NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "invites_organizations_invites" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "invite_recipient_owner_id" to table: "invites"
CREATE UNIQUE INDEX "invite_recipient_owner_id" ON "invites" ("recipient", "owner_id") WHERE (deleted_at IS NULL);
-- Create index "invites_mapping_id_key" to table: "invites"
CREATE UNIQUE INDEX "invites_mapping_id_key" ON "invites" ("mapping_id");
-- Create index "invites_token_key" to table: "invites"
CREATE UNIQUE INDEX "invites_token_key" ON "invites" ("token");
-- Create "invite_events" table
CREATE TABLE "invite_events" ("invite_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("invite_id", "event_id"), CONSTRAINT "invite_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "invite_events_invite_id" FOREIGN KEY ("invite_id") REFERENCES "invites" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "oauth_providers" table
CREATE TABLE "oauth_providers" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "client_id" character varying NOT NULL, "client_secret" character varying NOT NULL, "redirect_url" character varying NOT NULL, "scopes" character varying NOT NULL, "auth_url" character varying NOT NULL, "token_url" character varying NOT NULL, "auth_style" smallint NOT NULL, "info_url" character varying NOT NULL, "organization_oauthprovider" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "oauth_providers_organizations_oauthprovider" FOREIGN KEY ("organization_oauthprovider") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "oauth_providers_mapping_id_key" to table: "oauth_providers"
CREATE UNIQUE INDEX "oauth_providers_mapping_id_key" ON "oauth_providers" ("mapping_id");
-- Create "oh_auth_too_token_events" table
CREATE TABLE "oh_auth_too_token_events" ("oh_auth_too_token_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("oh_auth_too_token_id", "event_id"), CONSTRAINT "oh_auth_too_token_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "oh_auth_too_token_events_oh_auth_too_token_id" FOREIGN KEY ("oh_auth_too_token_id") REFERENCES "oh_auth_too_tokens" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "org_memberships" table
CREATE TABLE "org_memberships" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "role" character varying NOT NULL DEFAULT 'MEMBER', "organization_id" character varying NOT NULL, "user_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "org_memberships_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "org_memberships_users_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "org_memberships_mapping_id_key" to table: "org_memberships"
CREATE UNIQUE INDEX "org_memberships_mapping_id_key" ON "org_memberships" ("mapping_id");
-- Create index "orgmembership_user_id_organization_id" to table: "org_memberships"
CREATE UNIQUE INDEX "orgmembership_user_id_organization_id" ON "org_memberships" ("user_id", "organization_id") WHERE (deleted_at IS NULL);
-- Create "org_membership_events" table
CREATE TABLE "org_membership_events" ("org_membership_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("org_membership_id", "event_id"), CONSTRAINT "org_membership_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "org_membership_events_org_membership_id" FOREIGN KEY ("org_membership_id") REFERENCES "org_memberships" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "organization_events" table
CREATE TABLE "organization_events" ("organization_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("organization_id", "event_id"), CONSTRAINT "organization_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "organization_events_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "organization_features" table
CREATE TABLE "organization_features" ("organization_id" character varying NOT NULL, "feature_id" character varying NOT NULL, PRIMARY KEY ("organization_id", "feature_id"), CONSTRAINT "organization_features_feature_id" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "organization_features_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "organization_files" table
CREATE TABLE "organization_files" ("organization_id" character varying NOT NULL, "file_id" character varying NOT NULL, PRIMARY KEY ("organization_id", "file_id"), CONSTRAINT "organization_files_file_id" FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "organization_files_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "personal_access_tokens" table
CREATE TABLE "personal_access_tokens" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "mapping_id" character varying NOT NULL, "name" character varying NOT NULL, "token" character varying NOT NULL, "expires_at" timestamptz NOT NULL, "description" character varying NULL, "scopes" jsonb NULL, "last_used_at" timestamptz NULL, "owner_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "personal_access_tokens_users_personal_access_tokens" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "personal_access_tokens_mapping_id_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX "personal_access_tokens_mapping_id_key" ON "personal_access_tokens" ("mapping_id");
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX "personal_access_tokens_token_key" ON "personal_access_tokens" ("token");
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX "personalaccesstoken_token" ON "personal_access_tokens" ("token");
-- Create "organization_personal_access_tokens" table
CREATE TABLE "organization_personal_access_tokens" ("organization_id" character varying NOT NULL, "personal_access_token_id" character varying NOT NULL, PRIMARY KEY ("organization_id", "personal_access_token_id"), CONSTRAINT "organization_personal_access_tokens_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "organization_personal_access_tokens_personal_access_token_id" FOREIGN KEY ("personal_access_token_id") REFERENCES "personal_access_tokens" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "organization_secrets" table
CREATE TABLE "organization_secrets" ("organization_id" character varying NOT NULL, "hush_id" character varying NOT NULL, PRIMARY KEY ("organization_id", "hush_id"), CONSTRAINT "organization_secrets_hush_id" FOREIGN KEY ("hush_id") REFERENCES "hushes" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "organization_secrets_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "organization_settings" table
CREATE TABLE "organization_settings" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "domains" jsonb NULL, "billing_contact" character varying NULL, "billing_email" character varying NULL, "billing_phone" character varying NULL, "billing_address" character varying NULL, "tax_identifier" character varying NULL, "tags" jsonb NULL, "geo_location" character varying NULL DEFAULT 'AMER', "organization_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "organization_settings_organizations_setting" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "organization_settings_mapping_id_key" to table: "organization_settings"
CREATE UNIQUE INDEX "organization_settings_mapping_id_key" ON "organization_settings" ("mapping_id");
-- Create index "organization_settings_organization_id_key" to table: "organization_settings"
CREATE UNIQUE INDEX "organization_settings_organization_id_key" ON "organization_settings" ("organization_id");
-- Create "password_reset_tokens" table
CREATE TABLE "password_reset_tokens" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "token" character varying NOT NULL, "ttl" timestamptz NOT NULL, "email" character varying NOT NULL, "secret" bytea NOT NULL, "owner_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "password_reset_tokens_users_password_reset_tokens" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "password_reset_tokens_mapping_id_key" to table: "password_reset_tokens"
CREATE UNIQUE INDEX "password_reset_tokens_mapping_id_key" ON "password_reset_tokens" ("mapping_id");
-- Create index "password_reset_tokens_token_key" to table: "password_reset_tokens"
CREATE UNIQUE INDEX "password_reset_tokens_token_key" ON "password_reset_tokens" ("token");
-- Create index "passwordresettoken_token" to table: "password_reset_tokens"
CREATE UNIQUE INDEX "passwordresettoken_token" ON "password_reset_tokens" ("token") WHERE (deleted_at IS NULL);
-- Create "personal_access_token_events" table
CREATE TABLE "personal_access_token_events" ("personal_access_token_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("personal_access_token_id", "event_id"), CONSTRAINT "personal_access_token_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "personal_access_token_events_personal_access_token_id" FOREIGN KEY ("personal_access_token_id") REFERENCES "personal_access_tokens" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "subscribers" table
CREATE TABLE "subscribers" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "email" character varying NOT NULL, "phone_number" character varying NULL, "verified_email" boolean NOT NULL DEFAULT false, "verified_phone" boolean NOT NULL DEFAULT false, "active" boolean NOT NULL DEFAULT false, "token" character varying NOT NULL, "ttl" timestamptz NOT NULL, "secret" bytea NOT NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "subscribers_organizations_subscribers" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "subscriber_email_owner_id" to table: "subscribers"
CREATE UNIQUE INDEX "subscriber_email_owner_id" ON "subscribers" ("email", "owner_id") WHERE (deleted_at IS NULL);
-- Create index "subscribers_mapping_id_key" to table: "subscribers"
CREATE UNIQUE INDEX "subscribers_mapping_id_key" ON "subscribers" ("mapping_id");
-- Create index "subscribers_token_key" to table: "subscribers"
CREATE UNIQUE INDEX "subscribers_token_key" ON "subscribers" ("token");
-- Create "subscriber_events" table
CREATE TABLE "subscriber_events" ("subscriber_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("subscriber_id", "event_id"), CONSTRAINT "subscriber_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "subscriber_events_subscriber_id" FOREIGN KEY ("subscriber_id") REFERENCES "subscribers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "tfa_settings" table
CREATE TABLE "tfa_settings" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tfa_secret" character varying NULL, "verified" boolean NOT NULL DEFAULT false, "recovery_codes" jsonb NULL, "phone_otp_allowed" boolean NULL DEFAULT false, "email_otp_allowed" boolean NULL DEFAULT false, "totp_allowed" boolean NULL DEFAULT false, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "tfa_settings_users_tfa_settings" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "tfa_settings_mapping_id_key" to table: "tfa_settings"
CREATE UNIQUE INDEX "tfa_settings_mapping_id_key" ON "tfa_settings" ("mapping_id");
-- Create index "tfasetting_owner_id" to table: "tfa_settings"
CREATE UNIQUE INDEX "tfasetting_owner_id" ON "tfa_settings" ("owner_id") WHERE (deleted_at IS NULL);
-- Create "user_events" table
CREATE TABLE "user_events" ("user_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("user_id", "event_id"), CONSTRAINT "user_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "user_events_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "user_features" table
CREATE TABLE "user_features" ("user_id" character varying NOT NULL, "feature_id" character varying NOT NULL, PRIMARY KEY ("user_id", "feature_id"), CONSTRAINT "user_features_feature_id" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "user_features_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "user_settings" table
CREATE TABLE "user_settings" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "locked" boolean NOT NULL DEFAULT false, "silenced_at" timestamptz NULL, "suspended_at" timestamptz NULL, "status" character varying NOT NULL DEFAULT 'ACTIVE', "email_confirmed" boolean NOT NULL DEFAULT false, "tags" jsonb NULL, "is_webauthn_allowed" boolean NULL DEFAULT false, "is_tfa_enabled" boolean NULL DEFAULT false, "phone_number" character varying NULL, "user_id" character varying NULL, "user_setting_default_org" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "user_settings_organizations_default_org" FOREIGN KEY ("user_setting_default_org") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "user_settings_users_setting" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "user_settings_mapping_id_key" to table: "user_settings"
CREATE UNIQUE INDEX "user_settings_mapping_id_key" ON "user_settings" ("mapping_id");
-- Create index "user_settings_user_id_key" to table: "user_settings"
CREATE UNIQUE INDEX "user_settings_user_id_key" ON "user_settings" ("user_id");
-- Create "webauthns" table
CREATE TABLE "webauthns" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "credential_id" bytea NULL, "public_key" bytea NULL, "attestation_type" character varying NULL, "aaguid" bytea NOT NULL, "sign_count" integer NOT NULL, "transports" jsonb NOT NULL, "backup_eligible" boolean NOT NULL DEFAULT false, "backup_state" boolean NOT NULL DEFAULT false, "user_present" boolean NOT NULL DEFAULT false, "user_verified" boolean NOT NULL DEFAULT false, "owner_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "webauthns_users_webauthn" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "webauthns_aaguid_key" to table: "webauthns"
CREATE UNIQUE INDEX "webauthns_aaguid_key" ON "webauthns" ("aaguid");
-- Create index "webauthns_credential_id_key" to table: "webauthns"
CREATE UNIQUE INDEX "webauthns_credential_id_key" ON "webauthns" ("credential_id");
-- Create index "webauthns_mapping_id_key" to table: "webauthns"
CREATE UNIQUE INDEX "webauthns_mapping_id_key" ON "webauthns" ("mapping_id");
-- Create "webhooks" table
CREATE TABLE "webhooks" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "destination_url" character varying NOT NULL, "enabled" boolean NOT NULL DEFAULT true, "callback" character varying NULL, "expires_at" timestamptz NULL, "secret" bytea NULL, "failures" bigint NULL DEFAULT 0, "last_error" character varying NULL, "last_response" character varying NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "webhooks_organizations_webhooks" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "webhook_name_owner_id" to table: "webhooks"
CREATE UNIQUE INDEX "webhook_name_owner_id" ON "webhooks" ("name", "owner_id") WHERE (deleted_at IS NULL);
-- Create index "webhooks_callback_key" to table: "webhooks"
CREATE UNIQUE INDEX "webhooks_callback_key" ON "webhooks" ("callback");
-- Create index "webhooks_mapping_id_key" to table: "webhooks"
CREATE UNIQUE INDEX "webhooks_mapping_id_key" ON "webhooks" ("mapping_id");
-- Create "webhook_events" table
CREATE TABLE "webhook_events" ("webhook_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("webhook_id", "event_id"), CONSTRAINT "webhook_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "webhook_events_webhook_id" FOREIGN KEY ("webhook_id") REFERENCES "webhooks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
