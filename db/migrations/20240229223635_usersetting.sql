-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_user_settings" table
CREATE TABLE `new_user_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NOT NULL, `user_id` text NULL, `user_setting_default_org` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `user_settings_organizations_default_org` FOREIGN KEY (`user_setting_default_org`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "user_settings" to new temporary table "new_user_settings"
INSERT INTO `new_user_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags` FROM `user_settings`;
-- Drop "user_settings" table after copying rows
DROP TABLE `user_settings`;
-- Rename temporary table "new_user_settings" to "user_settings"
ALTER TABLE `new_user_settings` RENAME TO `user_settings`;
-- Create index "user_settings_user_id_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_id_key` ON `user_settings` (`user_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
