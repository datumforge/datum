-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_users" table
CREATE TABLE `new_users` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `locked` bool NOT NULL DEFAULT (false), `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, PRIMARY KEY (`id`));
-- Set sequence for "new_users" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_users", 25769803776);
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `email`, `first_name`, `last_name`, `display_name`, `locked`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `silenced_at`, `suspended_at`, `recovery_code`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `email`, `first_name`, `last_name`, `display_name`, `locked`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `silenced_at`, `suspended_at`, `recovery_code` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create "new_groups" table
CREATE TABLE `new_groups` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `name` text NOT NULL, `description` text NOT NULL DEFAULT (''), `logo_url` text NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "groups" to new temporary table "new_groups"
INSERT INTO `new_groups` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `description`, `logo_url` FROM `groups`;
-- Drop "groups" table after copying rows
DROP TABLE `groups`;
-- Rename temporary table "new_groups" to "groups"
ALTER TABLE `new_groups` RENAME TO `groups`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
