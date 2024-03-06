-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Add column "avatar_remote_url" to table: "organizations"
ALTER TABLE `organizations` ADD COLUMN `avatar_remote_url` text NULL;
-- Create "new_organization_settings" table
CREATE TABLE `new_organization_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `organization_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_setting`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "organization_settings" to new temporary table "new_organization_settings"
INSERT INTO `new_organization_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `domains`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `organization_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `domains`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `organization_setting` FROM `organization_settings`;
-- Drop "organization_settings" table after copying rows
DROP TABLE `organization_settings`;
-- Rename temporary table "new_organization_settings" to "organization_settings"
ALTER TABLE `new_organization_settings` RENAME TO `organization_settings`;
-- Create index "organization_settings_organization_setting_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_setting_key` ON `organization_settings` (`organization_setting`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
