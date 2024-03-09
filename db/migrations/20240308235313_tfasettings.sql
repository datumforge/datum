-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_users" table
CREATE TABLE `new_users` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `auth_provider` text NOT NULL DEFAULT ('CREDENTIALS'), PRIMARY KEY (`id`));
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `auth_provider`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `auth_provider` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "users_sub_key" to table: "users"
CREATE UNIQUE INDEX `users_sub_key` ON `users` (`sub`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create index "user_email_auth_provider" to table: "users"
CREATE UNIQUE INDEX `user_email_auth_provider` ON `users` (`email`, `auth_provider`) WHERE deleted_at is NULL;
-- Create "new_user_settings" table
CREATE TABLE `new_user_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NOT NULL, `is_webauthn_allowed` bool NULL DEFAULT (false), `is_tfa_enabled` bool NULL DEFAULT (false), `phone_number` text NULL, `user_id` text NULL, `user_setting_default_org` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `user_settings_organizations_default_org` FOREIGN KEY (`user_setting_default_org`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "user_settings" to new temporary table "new_user_settings"
INSERT INTO `new_user_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags`, `user_id`, `user_setting_default_org`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags`, `user_id`, `user_setting_default_org` FROM `user_settings`;
-- Drop "user_settings" table after copying rows
DROP TABLE `user_settings`;
-- Rename temporary table "new_user_settings" to "user_settings"
ALTER TABLE `new_user_settings` RENAME TO `user_settings`;
-- Create index "user_settings_user_id_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_id_key` ON `user_settings` (`user_id`);
-- Create "tfa_settings" table
CREATE TABLE `tfa_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tfa_secret` text NULL, `verified` bool NOT NULL DEFAULT (false), `recovery_codes` json NULL, `phone_otp_allowed` bool NULL DEFAULT (false), `email_otp_allowed` bool NULL DEFAULT (false), `totp_allowed` bool NULL DEFAULT (false), `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `tfa_settings_users_tfa_settings` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "tfasettings_owner_id" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfasettings_owner_id` ON `tfa_settings` (`owner_id`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
