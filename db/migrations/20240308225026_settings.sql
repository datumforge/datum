-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_group_settings" table
CREATE TABLE `new_group_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NULL, `sync_to_slack` bool NULL DEFAULT (false), `sync_to_github` bool NULL DEFAULT (false), `group_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "group_settings" to new temporary table "new_group_settings"
INSERT INTO `new_group_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `visibility`, `join_policy`, `tags`, `sync_to_slack`, `sync_to_github`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `visibility`, `join_policy`, `tags`, `sync_to_slack`, `sync_to_github` FROM `group_settings`;
-- Drop "group_settings" table after copying rows
DROP TABLE `group_settings`;
-- Rename temporary table "new_group_settings" to "group_settings"
ALTER TABLE `new_group_settings` RENAME TO `group_settings`;
-- Create index "group_settings_group_id_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_id_key` ON `group_settings` (`group_id`);
-- Create "new_organization_settings" table
CREATE TABLE `new_organization_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "organization_settings" to new temporary table "new_organization_settings"
INSERT INTO `new_organization_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `domains`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `domains`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags` FROM `organization_settings`;
-- Drop "organization_settings" table after copying rows
DROP TABLE `organization_settings`;
-- Rename temporary table "new_organization_settings" to "organization_settings"
ALTER TABLE `new_organization_settings` RENAME TO `organization_settings`;
-- Create index "organization_settings_organization_id_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_id_key` ON `organization_settings` (`organization_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
