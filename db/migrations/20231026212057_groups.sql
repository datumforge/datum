-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_memberships" table
CREATE TABLE `new_memberships` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `current` bool NOT NULL DEFAULT (false), `group_memberships` uuid NOT NULL, `organization_memberships` uuid NOT NULL, `user_memberships` uuid NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `memberships_groups_memberships` FOREIGN KEY (`group_memberships`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `memberships_organizations_memberships` FOREIGN KEY (`organization_memberships`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `memberships_users_memberships` FOREIGN KEY (`user_memberships`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "memberships" to new temporary table "new_memberships"
INSERT INTO `new_memberships` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `current`, `organization_memberships`, `user_memberships`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `current`, `organization_memberships`, `user_memberships` FROM `memberships`;
-- Drop "memberships" table after copying rows
DROP TABLE `memberships`;
-- Rename temporary table "new_memberships" to "memberships"
ALTER TABLE `new_memberships` RENAME TO `memberships`;
-- Create index "membership_organization_members_57e4e125ad3f56514a7fb2a9105c17d4" to table: "memberships"
CREATE UNIQUE INDEX `membership_organization_members_57e4e125ad3f56514a7fb2a9105c17d4` ON `memberships` (`organization_memberships`, `user_memberships`, `group_memberships`);
-- Create "groups" table
CREATE TABLE `groups` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `name` text NOT NULL, `description` text NOT NULL DEFAULT (''), `logo_url` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "groups_name_key" to table: "groups"
CREATE UNIQUE INDEX `groups_name_key` ON `groups` (`name`);
-- Create "group_settings" table
CREATE TABLE `group_settings` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('OPEN'), `group_setting` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
