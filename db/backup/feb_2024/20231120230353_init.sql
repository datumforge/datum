-- Create "entitlements" table
CREATE TABLE `entitlements` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `tier` text NOT NULL DEFAULT ('free'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires_at` datetime NULL, `upgraded_at` datetime NULL, `upgraded_tier` text NULL, `downgraded_at` datetime NULL, `downgraded_tier` text NULL, `cancelled` bool NOT NULL DEFAULT (false), `organization_entitlements` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`organization_entitlements`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create "groups" table
CREATE TABLE `groups` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `description` text NOT NULL DEFAULT (''), `logo_url` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `organization_groups` text NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`organization_groups`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Create index "group_name_organization_groups" to table: "groups"
CREATE UNIQUE INDEX `group_name_organization_groups` ON `groups` (`name`, `organization_groups`);
-- Create "group_settings" table
CREATE TABLE `group_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('OPEN'), `tags` json NOT NULL, `sync_to_slack` bool NOT NULL DEFAULT (false), `sync_to_github` bool NOT NULL DEFAULT (false), `group_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Create "integrations" table
CREATE TABLE `integrations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `kind` text NOT NULL, `description` text NULL, `secret_name` text NOT NULL, `organization_integrations` text NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`organization_integrations`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Create "oauth_providers" table
CREATE TABLE `oauth_providers` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, `organization_oauthprovider` text NULL, PRIMARY KEY (`id`), CONSTRAINT `oauth_providers_organizations_oauthprovider` FOREIGN KEY (`organization_oauthprovider`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create "organizations" table
CREATE TABLE `organizations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `description` text NULL, `parent_organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "organizations_name_key" to table: "organizations"
CREATE UNIQUE INDEX `organizations_name_key` ON `organizations` (`name`);
-- Create "organization_settings" table
CREATE TABLE `organization_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `domains` json NOT NULL, `sso_cert` text NOT NULL DEFAULT (''), `sso_entrypoint` text NOT NULL DEFAULT (''), `sso_issuer` text NOT NULL DEFAULT (''), `billing_contact` text NOT NULL, `billing_email` text NOT NULL, `billing_phone` text NOT NULL, `billing_address` text NOT NULL, `tax_identifier` text NOT NULL, `tags` json NULL, `organization_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_setting`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Create index "organization_settings_organization_setting_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_setting_key` ON `organization_settings` (`organization_setting`);
-- Create "personal_access_tokens" table
CREATE TABLE `personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `abilities` json NULL, `expiration_at` datetime NOT NULL, `last_used_at` datetime NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Create "refresh_tokens" table
CREATE TABLE `refresh_tokens` (`id` text NOT NULL, `client_id` text NOT NULL, `scopes` json NULL, `nonce` text NOT NULL, `claims_user_id` text NOT NULL, `claims_username` text NOT NULL, `claims_email` text NOT NULL, `claims_email_verified` bool NOT NULL, `claims_groups` json NULL, `claims_preferred_username` text NOT NULL, `connector_id` text NOT NULL, `connector_data` json NULL, `token` text NOT NULL, `obsolete_token` text NOT NULL, `last_used` datetime NOT NULL, `user_refreshtoken` text NULL, PRIMARY KEY (`id`), CONSTRAINT `refresh_tokens_users_refreshtoken` FOREIGN KEY (`user_refreshtoken`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create "sessions" table
CREATE TABLE `sessions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `type` text NOT NULL, `disabled` bool NOT NULL, `token` text NOT NULL, `user_agent` text NULL, `ips` text NOT NULL, `session_users` text NULL, `user_sessions` text NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_users` FOREIGN KEY (`session_users`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_sessions`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "sessions_token_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_token_key` ON `sessions` (`token`);
-- Create index "session_id" to table: "sessions"
CREATE UNIQUE INDEX `session_id` ON `sessions` (`id`);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password_hash` text NULL, PRIMARY KEY (`id`));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create "user_settings" table
CREATE TABLE `user_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `role` text NOT NULL DEFAULT ('USER'), `permissions` json NOT NULL, `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NOT NULL, `user_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_setting`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "user_settings_user_setting_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_setting_key` ON `user_settings` (`user_setting`);
-- Create "group_users" table
CREATE TABLE `group_users` (`group_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`group_id`, `user_id`), CONSTRAINT `group_users_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create "user_organizations" table
CREATE TABLE `user_organizations` (`user_id` text NOT NULL, `organization_id` text NOT NULL, PRIMARY KEY (`user_id`, `organization_id`), CONSTRAINT `user_organizations_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_organizations_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
