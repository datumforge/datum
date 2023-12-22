-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_organizations" table
CREATE TABLE `new_organizations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `personal_org` bool NOT NULL DEFAULT (false), `parent_organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "organizations" to new temporary table "new_organizations"
INSERT INTO `new_organizations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `parent_organization_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `parent_organization_id` FROM `organizations`;
-- Drop "organizations" table after copying rows
DROP TABLE `organizations`;
-- Rename temporary table "new_organizations" to "organizations"
ALTER TABLE `new_organizations` RENAME TO `organizations`;
-- Create index "organization_name" to table: "organizations"
CREATE UNIQUE INDEX `organization_name` ON `organizations` (`name`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
