-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop index "users_email_key" from table: "users"
DROP INDEX `users_email_key`;
-- Create index "user_email" to table: "users"
CREATE UNIQUE INDEX `user_email` ON `users` (`email`) WHERE deleted_at is NULL;
-- Create "new_group_memberships" table
CREATE TABLE `new_group_memberships` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `group_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `group_memberships_groups_group` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE NO ACTION, CONSTRAINT `group_memberships_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "group_memberships" to new temporary table "new_group_memberships"
INSERT INTO `new_group_memberships` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `role`, `group_id`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `role`, `group_id`, `user_id` FROM `group_memberships`;
-- Drop "group_memberships" table after copying rows
DROP TABLE `group_memberships`;
-- Rename temporary table "new_group_memberships" to "group_memberships"
ALTER TABLE `new_group_memberships` RENAME TO `group_memberships`;
-- Create index "groupmembership_user_id_group_id" to table: "group_memberships"
CREATE UNIQUE INDEX `groupmembership_user_id_group_id` ON `group_memberships` (`user_id`, `group_id`) WHERE deleted_at is NULL;
-- Create "new_org_memberships" table
CREATE TABLE `new_org_memberships` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `org_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `org_memberships_organizations_org` FOREIGN KEY (`org_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION, CONSTRAINT `org_memberships_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "org_memberships" to new temporary table "new_org_memberships"
INSERT INTO `new_org_memberships` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `role`, `org_id`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `role`, `org_id`, `user_id` FROM `org_memberships`;
-- Drop "org_memberships" table after copying rows
DROP TABLE `org_memberships`;
-- Rename temporary table "new_org_memberships" to "org_memberships"
ALTER TABLE `new_org_memberships` RENAME TO `org_memberships`;
-- Create index "orgmembership_user_id_org_id" to table: "org_memberships"
CREATE UNIQUE INDEX `orgmembership_user_id_org_id` ON `org_memberships` (`user_id`, `org_id`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
