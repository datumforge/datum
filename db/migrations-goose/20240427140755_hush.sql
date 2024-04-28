-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_integrations" table
CREATE TABLE `new_integrations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- copy rows from old table "integrations" to new temporary table "new_integrations"
INSERT INTO `new_integrations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `owner_id` FROM `integrations`;
-- drop "integrations" table after copying rows
DROP TABLE `integrations`;
-- rename temporary table "new_integrations" to "integrations"
ALTER TABLE `new_integrations` RENAME TO `integrations`;
-- create "new_integration_history" table
CREATE TABLE `new_integration_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, PRIMARY KEY (`id`));
-- copy rows from old table "integration_history" to new temporary table "new_integration_history"
INSERT INTO `new_integration_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `owner_id`, `name`, `description`, `kind`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `owner_id`, `name`, `description`, `kind` FROM `integration_history`;
-- drop "integration_history" table after copying rows
DROP TABLE `integration_history`;
-- rename temporary table "new_integration_history" to "integration_history"
ALTER TABLE `new_integration_history` RENAME TO `integration_history`;
-- create index "integrationhistory_history_time" to table: "integration_history"
CREATE INDEX `integrationhistory_history_time` ON `integration_history` (`history_time`);
-- create "hushes" table
CREATE TABLE `hushes` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NULL, `secret_value` text NULL, PRIMARY KEY (`id`));
-- create "hush_history" table
CREATE TABLE `hush_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NULL, `secret_value` text NULL, PRIMARY KEY (`id`));
-- create index "hushhistory_history_time" to table: "hush_history"
CREATE INDEX `hushhistory_history_time` ON `hush_history` (`history_time`);
-- create "integration_secrets" table
CREATE TABLE `integration_secrets` (`integration_id` text NOT NULL, `hush_id` text NOT NULL, PRIMARY KEY (`integration_id`, `hush_id`), CONSTRAINT `integration_secrets_integration_id` FOREIGN KEY (`integration_id`) REFERENCES `integrations` (`id`) ON DELETE CASCADE, CONSTRAINT `integration_secrets_hush_id` FOREIGN KEY (`hush_id`) REFERENCES `hushes` (`id`) ON DELETE CASCADE);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create "integration_secrets" table
DROP TABLE `integration_secrets`;
-- reverse: create index "hushhistory_history_time" to table: "hush_history"
DROP INDEX `hushhistory_history_time`;
-- reverse: create "hush_history" table
DROP TABLE `hush_history`;
-- reverse: create "hushes" table
DROP TABLE `hushes`;
-- reverse: create index "integrationhistory_history_time" to table: "integration_history"
DROP INDEX `integrationhistory_history_time`;
-- reverse: create "new_integration_history" table
DROP TABLE `new_integration_history`;
-- reverse: create "new_integrations" table
DROP TABLE `new_integrations`;
