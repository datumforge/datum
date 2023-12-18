-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_integrations" table
CREATE TABLE `new_integrations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `kind` text NOT NULL, `description` text NULL, `secret_name` text NOT NULL, `organization_integrations` text NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`organization_integrations`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "integrations" to new temporary table "new_integrations"
INSERT INTO `new_integrations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `kind`, `description`, `secret_name`, `organization_integrations`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `kind`, `description`, `secret_name`, `organization_integrations` FROM `integrations`;
-- Drop "integrations" table after copying rows
DROP TABLE `integrations`;
-- Rename temporary table "new_integrations" to "integrations"
ALTER TABLE `new_integrations` RENAME TO `integrations`;
-- Create "new_organization_settings" table
CREATE TABLE `new_organization_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `domains` json NULL, `sso_cert` text NULL, `sso_entrypoint` text NULL, `sso_issuer` text NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `organization_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_setting`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "organization_settings" to new temporary table "new_organization_settings"
INSERT INTO `new_organization_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `domains`, `sso_cert`, `sso_entrypoint`, `sso_issuer`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `organization_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `domains`, `sso_cert`, `sso_entrypoint`, `sso_issuer`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `organization_setting` FROM `organization_settings`;
-- Drop "organization_settings" table after copying rows
DROP TABLE `organization_settings`;
-- Rename temporary table "new_organization_settings" to "organization_settings"
ALTER TABLE `new_organization_settings` RENAME TO `organization_settings`;
-- Create index "organization_settings_organization_setting_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_setting_key` ON `organization_settings` (`organization_setting`);
-- Create "new_groups" table
CREATE TABLE `new_groups` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), `organization_groups` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`organization_groups`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "groups" to new temporary table "new_groups"
INSERT INTO `new_groups` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `logo_url`, `display_name`, `organization_groups`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `logo_url`, `display_name`, `organization_groups` FROM `groups`;
-- Drop "groups" table after copying rows
DROP TABLE `groups`;
-- Rename temporary table "new_groups" to "groups"
ALTER TABLE `new_groups` RENAME TO `groups`;
-- Create index "group_name_organization_groups" to table: "groups"
CREATE UNIQUE INDEX `group_name_organization_groups` ON `groups` (`name`, `organization_groups`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
