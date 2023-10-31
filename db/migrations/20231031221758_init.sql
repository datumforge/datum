-- Create "groups" table
CREATE TABLE `groups` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `name` text NOT NULL, `description` text NOT NULL DEFAULT (''), `logo_url` text NOT NULL, `organization_groups` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`organization_groups`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Create index "group_name_organization_groups" to table: "groups"
CREATE UNIQUE INDEX `group_name_organization_groups` ON `groups` (`name`, `organization_groups`);
-- Create "group_settings" table
CREATE TABLE `group_settings` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('OPEN'), `group_setting` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Create "integrations" table
CREATE TABLE `integrations` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `name` text NOT NULL, `kind` text NOT NULL, `description` text NULL, `secret_name` text NOT NULL, `organization_integrations` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`organization_integrations`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Create "organizations" table
CREATE TABLE `organizations` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `name` text NOT NULL, `description` text NULL, `parent_organization_id` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "organizations_name_key" to table: "organizations"
CREATE UNIQUE INDEX `organizations_name_key` ON `organizations` (`name`);
-- Create "sessions" table
CREATE TABLE `sessions` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `type` text NOT NULL, `disabled` bool NOT NULL, `token` text NOT NULL, `user_agent` text NULL, `ips` text NOT NULL, `session_users` uuid NULL, `user_sessions` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_users` FOREIGN KEY (`session_users`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_sessions`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create index "sessions_token_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_token_key` ON `sessions` (`token`);
-- Create index "session_id" to table: "sessions"
CREATE UNIQUE INDEX `session_id` ON `sessions` (`id`);
-- Create "users" table
CREATE TABLE `users` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `locked` bool NOT NULL DEFAULT (false), `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, PRIMARY KEY (`id`));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create "group_users" table
CREATE TABLE `group_users` (`group_id` uuid NOT NULL, `user_id` uuid NOT NULL, PRIMARY KEY (`group_id`, `user_id`), CONSTRAINT `group_users_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Create "user_organizations" table
CREATE TABLE `user_organizations` (`user_id` uuid NOT NULL, `organization_id` uuid NOT NULL, PRIMARY KEY (`user_id`, `organization_id`), CONSTRAINT `user_organizations_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_organizations_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
