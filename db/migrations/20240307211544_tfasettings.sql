-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_tfa_settings" table
CREATE TABLE `new_tfa_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tfa_secret` text NULL, `verified` bool NOT NULL DEFAULT (false), `recovery_codes` json NULL, `phone_otp_allowed` bool NULL DEFAULT (false), `email_otp_allowed` bool NULL DEFAULT (false), `totp_allowed` bool NULL DEFAULT (false), `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `tfa_settings_users_tfa_settings` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "tfa_settings" to new temporary table "new_tfa_settings"
INSERT INTO `new_tfa_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `tfa_secret`, `recovery_codes`, `phone_otp_allowed`, `email_otp_allowed`, `totp_allowed`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `tfa_secret`, `recovery_codes`, `phone_otp_allowed`, `email_otp_allowed`, `totp_allowed`, `owner_id` FROM `tfa_settings`;
-- Drop "tfa_settings" table after copying rows
DROP TABLE `tfa_settings`;
-- Rename temporary table "new_tfa_settings" to "tfa_settings"
ALTER TABLE `new_tfa_settings` RENAME TO `tfa_settings`;
-- Create index "tfa_settings_owner_id_key" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfa_settings_owner_id_key` ON `tfa_settings` (`owner_id`);
-- Create index "tfasettings_owner_id" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfasettings_owner_id` ON `tfa_settings` (`owner_id`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
