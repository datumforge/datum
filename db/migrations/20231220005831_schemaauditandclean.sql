-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_entitlements" table
CREATE TABLE `new_entitlements` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tier` text NOT NULL DEFAULT ('free'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), `organization_id` text NOT NULL, `organization_entitlements` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`organization_entitlements`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "entitlements" to new temporary table "new_entitlements"
INSERT INTO `new_entitlements` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `tier`, `external_customer_id`, `external_subscription_id`, `expires_at`, `cancelled`, `organization_entitlements`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `tier`, `external_customer_id`, `external_subscription_id`, `expires_at`, `cancelled`, `organization_entitlements` FROM `entitlements`;
-- Drop "entitlements" table after copying rows
DROP TABLE `entitlements`;
-- Rename temporary table "new_entitlements" to "entitlements"
ALTER TABLE `new_entitlements` RENAME TO `entitlements`;
-- Add column "deleted_at" to table: "oauth_providers"
ALTER TABLE `oauth_providers` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "oauth_providers"
ALTER TABLE `oauth_providers` ADD COLUMN `deleted_by` text NULL;
-- Add column "deleted_at" to table: "user_settings"
ALTER TABLE `user_settings` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "user_settings"
ALTER TABLE `user_settings` ADD COLUMN `deleted_by` text NULL;
-- Create "new_integrations" table
CREATE TABLE `new_integrations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NOT NULL, `organization_integrations` text NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`organization_integrations`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "integrations" to new temporary table "new_integrations"
INSERT INTO `new_integrations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `kind`, `secret_name`, `organization_integrations`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `kind`, `secret_name`, `organization_integrations` FROM `integrations`;
-- Drop "integrations" table after copying rows
DROP TABLE `integrations`;
-- Rename temporary table "new_integrations" to "integrations"
ALTER TABLE `new_integrations` RENAME TO `integrations`;
-- Add column "deleted_at" to table: "organization_settings"
ALTER TABLE `organization_settings` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "organization_settings"
ALTER TABLE `organization_settings` ADD COLUMN `deleted_by` text NULL;
-- Add column "gravatar_logo_url" to table: "groups"
ALTER TABLE `groups` ADD COLUMN `gravatar_logo_url` text NULL;
-- Add column "organization_id" to table: "groups"
ALTER TABLE `groups` ADD COLUMN `organization_id` text NOT NULL;
-- Create "new_group_settings" table
CREATE TABLE `new_group_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NOT NULL, `sync_to_slack` bool NOT NULL DEFAULT (false), `sync_to_github` bool NOT NULL DEFAULT (false), `group_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "group_settings" to new temporary table "new_group_settings"
INSERT INTO `new_group_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, `join_policy`, `tags`, `sync_to_slack`, `sync_to_github`, `group_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, IFNULL(`join_policy`, ('INVITE_OR_APPLICATION')) AS `join_policy`, `tags`, `sync_to_slack`, `sync_to_github`, `group_setting` FROM `group_settings`;
-- Drop "group_settings" table after copying rows
DROP TABLE `group_settings`;
-- Rename temporary table "new_group_settings" to "group_settings"
ALTER TABLE `new_group_settings` RENAME TO `group_settings`;
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Create "new_personal_access_tokens" table
CREATE TABLE `new_personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `abilities` json NULL, `expires_at` datetime NOT NULL, `description` text NULL DEFAULT (''), `last_used_at` datetime NULL, `user_personal_access_tokens` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`user_personal_access_tokens`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "personal_access_tokens" to new temporary table "new_personal_access_tokens"
INSERT INTO `new_personal_access_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `token`, `abilities`, `expires_at`, `description`, `last_used_at`, `user_personal_access_tokens`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `token`, `abilities`, `expires_at`, `description`, `last_used_at`, `user_personal_access_tokens` FROM `personal_access_tokens`;
-- Drop "personal_access_tokens" table after copying rows
DROP TABLE `personal_access_tokens`;
-- Rename temporary table "new_personal_access_tokens" to "personal_access_tokens"
ALTER TABLE `new_personal_access_tokens` RENAME TO `personal_access_tokens`;
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
