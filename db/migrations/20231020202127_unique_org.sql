-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_organizations" table
CREATE TABLE `new_organizations` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- copy rows from old table "organizations" to new temporary table "new_organizations"
INSERT INTO `new_organizations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name` FROM `organizations`;
-- drop "organizations" table after copying rows
DROP TABLE `organizations`;
-- rename temporary table "new_organizations" to "organizations"
ALTER TABLE `new_organizations` RENAME TO `organizations`;
-- create index "organizations_name_key" to table: "organizations"
CREATE UNIQUE INDEX `organizations_name_key` ON `organizations` (`name`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "organizations_name_key" to table: "organizations"
DROP INDEX `organizations_name_key`;
-- reverse: create "new_organizations" table
DROP TABLE `new_organizations`;
