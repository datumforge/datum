-- Create "permissions" table
CREATE TABLE `permissions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `action` text NOT NULL, `description` text NULL, `is_disabled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "permissions_name_key" to table: "permissions"
CREATE UNIQUE INDEX `permissions_name_key` ON `permissions` (`name`);
-- Create index "permissions_action_key" to table: "permissions"
CREATE UNIQUE INDEX `permissions_action_key` ON `permissions` (`action`);
-- Create "roles" table
CREATE TABLE `roles` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `description` text NULL, `is_disabled` bool NOT NULL DEFAULT (false), `created_time` datetime NOT NULL, PRIMARY KEY (`id`));
-- Create index "roles_name_key" to table: "roles"
CREATE UNIQUE INDEX `roles_name_key` ON `roles` (`name`);
-- Create "role_permissions" table
CREATE TABLE `role_permissions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `role_id` text NOT NULL, `permission_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `role_permissions_roles_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE NO ACTION, CONSTRAINT `role_permissions_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE NO ACTION);
-- Create index "rolepermission_role_id_permission_id" to table: "role_permissions"
CREATE UNIQUE INDEX `rolepermission_role_id_permission_id` ON `role_permissions` (`role_id`, `permission_id`);
-- Create "user_roles" table
CREATE TABLE `user_roles` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `user_id` text NOT NULL, `role_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_roles_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION, CONSTRAINT `user_roles_roles_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE NO ACTION);
-- Create index "userrole_user_id_role_id" to table: "user_roles"
CREATE UNIQUE INDEX `userrole_user_id_role_id` ON `user_roles` (`user_id`, `role_id`);
