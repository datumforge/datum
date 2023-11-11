-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_group_settings" table
CREATE TABLE `new_group_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('OPEN'), `tags` json NOT NULL, `sync_to_slack` bool NOT NULL DEFAULT (false), `sync_to_github` bool NOT NULL DEFAULT (false), `group_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "group_settings" to new temporary table "new_group_settings"
INSERT INTO `new_group_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, `join_policy`, `tags`, `group_setting`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `visibility`, `join_policy`, `tags`, `group_setting` FROM `group_settings`;
-- Drop "group_settings" table after copying rows
DROP TABLE `group_settings`;
-- Rename temporary table "new_group_settings" to "group_settings"
ALTER TABLE `new_group_settings` RENAME TO `group_settings`;
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
