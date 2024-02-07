-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_user_settings" table
CREATE TABLE `new_user_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `role` text NOT NULL DEFAULT ('USER'), `permissions` json NOT NULL, `email_confirmed` bool NOT NULL DEFAULT (true), `tags` json NOT NULL, `user_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_setting`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "user_settings" to new temporary table "new_user_settings"
INSERT INTO `new_user_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `locked`, `silenced_at`, `suspended_at`, `recovery_code`, `status`, `role`, `permissions`, `email_confirmed`, `tags`, `user_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `locked`, `silenced_at`, `suspended_at`, `recovery_code`, `status`, `role`, `permissions`, IFNULL(`email_confirmed`, (true)) AS `email_confirmed`, `tags`, `user_setting` FROM `user_settings`;
-- Drop "user_settings" table after copying rows
DROP TABLE `user_settings`;
-- Rename temporary table "new_user_settings" to "user_settings"
ALTER TABLE `new_user_settings` RENAME TO `user_settings`;
-- Create index "user_settings_user_setting_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_setting_key` ON `user_settings` (`user_setting`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
