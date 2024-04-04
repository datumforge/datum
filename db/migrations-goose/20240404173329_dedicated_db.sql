-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_organizations" table
CREATE TABLE `new_organizations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `personal_org` bool NULL DEFAULT (false), `avatar_remote_url` text NULL, `dedicated_db` bool NOT NULL DEFAULT (false), `parent_organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- copy rows from old table "organizations" to new temporary table "new_organizations"
INSERT INTO `new_organizations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `personal_org`, `avatar_remote_url`, `parent_organization_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `personal_org`, `avatar_remote_url`, `parent_organization_id` FROM `organizations`;
-- drop "organizations" table after copying rows
DROP TABLE `organizations`;
-- rename temporary table "new_organizations" to "organizations"
ALTER TABLE `new_organizations` RENAME TO `organizations`;
-- create index "organization_name" to table: "organizations"
CREATE UNIQUE INDEX `organization_name` ON `organizations` (`name`) WHERE deleted_at is NULL;
-- add column "geo_location" to table: "organization_settings"
ALTER TABLE `organization_settings` ADD COLUMN `geo_location` text NULL;
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: add column "geo_location" to table: "organization_settings"
ALTER TABLE `organization_settings` DROP COLUMN `geo_location`;
-- reverse: create index "organization_name" to table: "organizations"
DROP INDEX `organization_name`;
-- reverse: create "new_organizations" table
DROP TABLE `new_organizations`;
