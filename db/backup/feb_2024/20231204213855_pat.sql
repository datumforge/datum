-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_groups" table
CREATE TABLE `new_groups` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NOT NULL DEFAULT (''), `logo_url` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `organization_groups` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`organization_groups`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "groups" to new temporary table "new_groups"
INSERT INTO `new_groups` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url`, `display_name`, `organization_groups`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url`, `display_name`, `organization_groups` FROM `groups`;
-- Drop "groups" table after copying rows
DROP TABLE `groups`;
-- Rename temporary table "new_groups" to "groups"
ALTER TABLE `new_groups` RENAME TO `groups`;
-- Create index "group_name_organization_groups" to table: "groups"
CREATE UNIQUE INDEX `group_name_organization_groups` ON `groups` (`name`, `organization_groups`) WHERE deleted_at is NULL;
-- Create "new_group_settings" table
CREATE TABLE `new_group_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('OPEN'), `tags` json NOT NULL, `sync_to_slack` bool NOT NULL DEFAULT (false), `sync_to_github` bool NOT NULL DEFAULT (false), `group_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "group_settings" to new temporary table "new_group_settings"
INSERT INTO `new_group_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, `join_policy`, `tags`, `sync_to_slack`, `sync_to_github`, `group_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, `join_policy`, `tags`, `sync_to_slack`, `sync_to_github`, `group_setting` FROM `group_settings`;
-- Drop "group_settings" table after copying rows
DROP TABLE `group_settings`;
-- Rename temporary table "new_group_settings" to "group_settings"
ALTER TABLE `new_group_settings` RENAME TO `group_settings`;
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Create "new_personal_access_tokens" table
CREATE TABLE `new_personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `abilities` json NULL, `expiration_at` datetime NOT NULL, `description` text NOT NULL DEFAULT (''), `last_used_at` datetime NULL, `user_personal_access_tokens` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`user_personal_access_tokens`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "personal_access_tokens" to new temporary table "new_personal_access_tokens"
INSERT INTO `new_personal_access_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `token`, `abilities`, `expiration_at`, `last_used_at`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `token`, `abilities`, `expiration_at`, `last_used_at` FROM `personal_access_tokens`;
-- Drop "personal_access_tokens" table after copying rows
DROP TABLE `personal_access_tokens`;
-- Rename temporary table "new_personal_access_tokens" to "personal_access_tokens"
ALTER TABLE `new_personal_access_tokens` RENAME TO `personal_access_tokens`;
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Add column "deleted_at" to table: "users"
ALTER TABLE `users` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "users"
ALTER TABLE `users` ADD COLUMN `deleted_by` text NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
