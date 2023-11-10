-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_groups" table
CREATE TABLE `new_groups` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `description` text NOT NULL DEFAULT (''), `logo_url` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `organization_groups` text NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`organization_groups`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "groups" to new temporary table "new_groups"
INSERT INTO `new_groups` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url`, `organization_groups`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url`, `organization_groups` FROM `groups`;
-- Drop "groups" table after copying rows
DROP TABLE `groups`;
-- Rename temporary table "new_groups" to "groups"
ALTER TABLE `new_groups` RENAME TO `groups`;
-- Create index "group_name_organization_groups" to table: "groups"
CREATE UNIQUE INDEX `group_name_organization_groups` ON `groups` (`name`, `organization_groups`);
-- Add column "tags" to table: "group_settings"
ALTER TABLE `group_settings` ADD COLUMN `tags` json NOT NULL;
-- Create "new_organizations" table
CREATE TABLE `new_organizations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `description` text NULL, `parent_organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "organizations" to new temporary table "new_organizations"
INSERT INTO `new_organizations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `parent_organization_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `parent_organization_id` FROM `organizations`;
-- Drop "organizations" table after copying rows
DROP TABLE `organizations`;
-- Rename temporary table "new_organizations" to "organizations"
ALTER TABLE `new_organizations` RENAME TO `organizations`;
-- Create index "organizations_name_key" to table: "organizations"
CREATE UNIQUE INDEX `organizations_name_key` ON `organizations` (`name`);
-- Create "new_users" table
CREATE TABLE `new_users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password_hash` text NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create "new_organization_settings" table
CREATE TABLE `new_organization_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `domains` json NOT NULL, `sso_cert` text NOT NULL DEFAULT (''), `sso_entrypoint` text NOT NULL DEFAULT (''), `sso_issuer` text NOT NULL DEFAULT (''), `billing_contact` text NOT NULL, `billing_email` text NOT NULL, `billing_phone` text NOT NULL, `billing_address` text NOT NULL, `tax_identifier` text NOT NULL, `tags` json NOT NULL, `organization_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_setting`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "organization_settings" to new temporary table "new_organization_settings"
INSERT INTO `new_organization_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `domains`, `sso_cert`, `sso_entrypoint`, `sso_issuer`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `domains`, `sso_cert`, `sso_entrypoint`, `sso_issuer`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier` FROM `organization_settings`;
-- Drop "organization_settings" table after copying rows
DROP TABLE `organization_settings`;
-- Rename temporary table "new_organization_settings" to "organization_settings"
ALTER TABLE `new_organization_settings` RENAME TO `organization_settings`;
-- Create index "organization_settings_organization_setting_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_setting_key` ON `organization_settings` (`organization_setting`);
-- Create "user_settings" table
CREATE TABLE `user_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `role` text NOT NULL DEFAULT ('USER'), `permissions` json NOT NULL, `email_confirmed` bool NOT NULL DEFAULT (true), `tags` json NOT NULL, `user_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_setting`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "user_settings_user_setting_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_setting_key` ON `user_settings` (`user_setting`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
