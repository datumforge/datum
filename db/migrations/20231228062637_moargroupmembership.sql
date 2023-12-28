-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_group_memberships" table
CREATE TABLE `new_group_memberships` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `group_role` text NOT NULL, `group_members` text NOT NULL, `role_id` text NULL, `user_group_memberships` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `group_memberships_groups_members` FOREIGN KEY (`group_members`) REFERENCES `groups` (`id`) ON DELETE NO ACTION, CONSTRAINT `group_memberships_roles_group_roles` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE SET NULL, CONSTRAINT `group_memberships_users_group_memberships` FOREIGN KEY (`user_group_memberships`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "group_memberships" to new temporary table "new_group_memberships"
INSERT INTO `new_group_memberships` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `group_members`, `user_group_memberships`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `group_members`, `user_group_memberships` FROM `group_memberships`;
-- Drop "group_memberships" table after copying rows
DROP TABLE `group_memberships`;
-- Rename temporary table "new_group_memberships" to "group_memberships"
ALTER TABLE `new_group_memberships` RENAME TO `group_memberships`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
