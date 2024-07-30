-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_organizations" table
CREATE TABLE `new_organizations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NULL, `display_name` text NULL, `description` text NULL, `personal_org` bool NULL DEFAULT (false), `avatar_remote_url` text NULL, `dedicated_db` bool NOT NULL DEFAULT (false), `parent_organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- copy rows from old table "organizations" to new temporary table "new_organizations"
INSERT INTO `new_organizations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `personal_org`, `avatar_remote_url`, `dedicated_db`, `parent_organization_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `personal_org`, `avatar_remote_url`, `dedicated_db`, `parent_organization_id` FROM `organizations`;
-- drop "organizations" table after copying rows
DROP TABLE `organizations`;
-- rename temporary table "new_organizations" to "organizations"
ALTER TABLE `new_organizations` RENAME TO `organizations`;
-- create index "organizations_mapping_id_key" to table: "organizations"
CREATE UNIQUE INDEX `organizations_mapping_id_key` ON `organizations` (`mapping_id`);
-- create index "organization_name" to table: "organizations"
CREATE UNIQUE INDEX `organization_name` ON `organizations` (`name`) WHERE deleted_at is NULL;
-- create "new_organization_history" table
CREATE TABLE `new_organization_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NULL, `display_name` text NULL, `description` text NULL, `parent_organization_id` text NULL, `personal_org` bool NULL DEFAULT (false), `avatar_remote_url` text NULL, `dedicated_db` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- copy rows from old table "organization_history" to new temporary table "new_organization_history"
INSERT INTO `new_organization_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `parent_organization_id`, `personal_org`, `avatar_remote_url`, `dedicated_db`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `name`, `display_name`, `description`, `parent_organization_id`, `personal_org`, `avatar_remote_url`, `dedicated_db` FROM `organization_history`;
-- drop "organization_history" table after copying rows
DROP TABLE `organization_history`;
-- rename temporary table "new_organization_history" to "organization_history"
ALTER TABLE `new_organization_history` RENAME TO `organization_history`;
-- create index "organizationhistory_history_time" to table: "organization_history"
CREATE INDEX `organizationhistory_history_time` ON `organization_history` (`history_time`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "organizationhistory_history_time" to table: "organization_history"
DROP INDEX `organizationhistory_history_time`;
-- reverse: create "new_organization_history" table
DROP TABLE `new_organization_history`;
-- reverse: create index "organization_name" to table: "organizations"
DROP INDEX `organization_name`;
-- reverse: create index "organizations_mapping_id_key" to table: "organizations"
DROP INDEX `organizations_mapping_id_key`;
-- reverse: create "new_organizations" table
DROP TABLE `new_organizations`;
