-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_organization_settings" table
CREATE TABLE `new_organization_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `domains` json NOT NULL, `sso_cert` text NOT NULL DEFAULT (''), `sso_entrypoint` text NOT NULL DEFAULT (''), `sso_issuer` text NOT NULL DEFAULT (''), `billing_contact` text NOT NULL, `billing_email` text NOT NULL, `billing_phone` text NOT NULL, `billing_address` text NOT NULL, `tax_identifier` text NOT NULL, `tags` json NULL, `organization_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_setting`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "organization_settings" to new temporary table "new_organization_settings"
INSERT INTO `new_organization_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `domains`, `sso_cert`, `sso_entrypoint`, `sso_issuer`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `organization_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `domains`, `sso_cert`, `sso_entrypoint`, `sso_issuer`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `organization_setting` FROM `organization_settings`;
-- Drop "organization_settings" table after copying rows
DROP TABLE `organization_settings`;
-- Rename temporary table "new_organization_settings" to "organization_settings"
ALTER TABLE `new_organization_settings` RENAME TO `organization_settings`;
-- Create index "organization_settings_organization_setting_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_setting_key` ON `organization_settings` (`organization_setting`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
