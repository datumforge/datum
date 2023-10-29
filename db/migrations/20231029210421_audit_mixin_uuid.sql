-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_integrations" table
CREATE TABLE `new_integrations` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `kind` text NOT NULL, `description` text NULL, `secret_name` text NOT NULL, `organization_integrations` uuid NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`organization_integrations`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Set sequence for "new_integrations" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_integrations", 8589934592);
-- Copy rows from old table "integrations" to new temporary table "new_integrations"
INSERT INTO `new_integrations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `kind`, `description`, `secret_name`, `organization_integrations`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `kind`, `description`, `secret_name`, `organization_integrations` FROM `integrations`;
-- Drop "integrations" table after copying rows
DROP TABLE `integrations`;
-- Rename temporary table "new_integrations" to "integrations"
ALTER TABLE `new_integrations` RENAME TO `integrations`;
-- Create "new_organizations" table
CREATE TABLE `new_organizations` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- Set sequence for "new_organizations" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_organizations", 17179869184);
-- Copy rows from old table "organizations" to new temporary table "new_organizations"
INSERT INTO `new_organizations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name` FROM `organizations`;
-- Drop "organizations" table after copying rows
DROP TABLE `organizations`;
-- Rename temporary table "new_organizations" to "organizations"
ALTER TABLE `new_organizations` RENAME TO `organizations`;
-- Create index "organizations_name_key" to table: "organizations"
CREATE UNIQUE INDEX `organizations_name_key` ON `organizations` (`name`);
-- Create "new_sessions" table
CREATE TABLE `new_sessions` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `type` text NOT NULL, `disabled` bool NOT NULL, `token` text NOT NULL, `user_agent` text NULL, `ips` text NOT NULL, `session_users` uuid NULL, `user_sessions` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_users` FOREIGN KEY (`session_users`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_sessions`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Set sequence for "new_sessions" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_sessions", 21474836480);
-- Copy rows from old table "sessions" to new temporary table "new_sessions"
INSERT INTO `new_sessions` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `type`, `disabled`, `token`, `user_agent`, `ips`, `session_users`, `user_sessions`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `type`, `disabled`, `token`, `user_agent`, `ips`, `session_users`, `user_sessions` FROM `sessions`;
-- Drop "sessions" table after copying rows
DROP TABLE `sessions`;
-- Rename temporary table "new_sessions" to "sessions"
ALTER TABLE `new_sessions` RENAME TO `sessions`;
-- Create index "sessions_token_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_token_key` ON `sessions` (`token`);
-- Create index "session_id" to table: "sessions"
CREATE UNIQUE INDEX `session_id` ON `sessions` (`id`);
-- Create "new_memberships" table
CREATE TABLE `new_memberships` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `current` bool NOT NULL DEFAULT (false), `group_memberships` uuid NOT NULL, `organization_memberships` uuid NOT NULL, `user_memberships` uuid NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `memberships_groups_memberships` FOREIGN KEY (`group_memberships`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `memberships_organizations_memberships` FOREIGN KEY (`organization_memberships`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `memberships_users_memberships` FOREIGN KEY (`user_memberships`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Set sequence for "new_memberships" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_memberships", 12884901888);
-- Copy rows from old table "memberships" to new temporary table "new_memberships"
INSERT INTO `new_memberships` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `current`, `group_memberships`, `organization_memberships`, `user_memberships`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `current`, `group_memberships`, `organization_memberships`, `user_memberships` FROM `memberships`;
-- Drop "memberships" table after copying rows
DROP TABLE `memberships`;
-- Rename temporary table "new_memberships" to "memberships"
ALTER TABLE `new_memberships` RENAME TO `memberships`;
-- Create index "membership_organization_members_57e4e125ad3f56514a7fb2a9105c17d4" to table: "memberships"
CREATE UNIQUE INDEX `membership_organization_members_57e4e125ad3f56514a7fb2a9105c17d4` ON `memberships` (`organization_memberships`, `user_memberships`, `group_memberships`);
-- Create "new_group_settings" table
CREATE TABLE `new_group_settings` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('OPEN'), `group_setting` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Set sequence for "new_group_settings" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_group_settings", 4294967296);
-- Copy rows from old table "group_settings" to new temporary table "new_group_settings"
INSERT INTO `new_group_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, `join_policy`, `group_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, `join_policy`, `group_setting` FROM `group_settings`;
-- Drop "group_settings" table after copying rows
DROP TABLE `group_settings`;
-- Rename temporary table "new_group_settings" to "group_settings"
ALTER TABLE `new_group_settings` RENAME TO `group_settings`;
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Create "new_users" table
CREATE TABLE `new_users` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `locked` bool NOT NULL DEFAULT (false), `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, PRIMARY KEY (`id`));
-- Set sequence for "new_users" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_users", 25769803776);
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `email`, `first_name`, `last_name`, `display_name`, `locked`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `silenced_at`, `suspended_at`, `recovery_code`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `email`, `first_name`, `last_name`, `display_name`, `locked`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `silenced_at`, `suspended_at`, `recovery_code` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create "new_groups" table
CREATE TABLE `new_groups` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` uuid NULL, `updated_by` uuid NULL, `name` text NOT NULL, `description` text NOT NULL DEFAULT (''), `logo_url` text NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "groups" to new temporary table "new_groups"
INSERT INTO `new_groups` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url` FROM `groups`;
-- Drop "groups" table after copying rows
DROP TABLE `groups`;
-- Rename temporary table "new_groups" to "groups"
ALTER TABLE `new_groups` RENAME TO `groups`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
