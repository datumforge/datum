-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_document_data" table
CREATE TABLE `new_document_data` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `data` json NOT NULL, `owner_id` text NULL, `template_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `document_data_organizations_documentdata` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL, CONSTRAINT `document_data_templates_documents` FOREIGN KEY (`template_id`) REFERENCES `templates` (`id`) ON DELETE NO ACTION);
-- copy rows from old table "document_data" to new temporary table "new_document_data"
INSERT INTO `new_document_data` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `data`, `template_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `data`, `template_id` FROM `document_data`;
-- drop "document_data" table after copying rows
DROP TABLE `document_data`;
-- rename temporary table "new_document_data" to "document_data"
ALTER TABLE `new_document_data` RENAME TO `document_data`;
-- create index "document_data_mapping_id_key" to table: "document_data"
CREATE UNIQUE INDEX `document_data_mapping_id_key` ON `document_data` (`mapping_id`);
-- add column "owner_id" to table: "document_data_history"
ALTER TABLE `document_data_history` ADD COLUMN `owner_id` text NULL;
-- create "new_oauth_providers" table
CREATE TABLE `new_oauth_providers` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `oauth_providers_organizations_oauthprovider` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- copy rows from old table "oauth_providers" to new temporary table "new_oauth_providers"
INSERT INTO `new_oauth_providers` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url` FROM `oauth_providers`;
-- drop "oauth_providers" table after copying rows
DROP TABLE `oauth_providers`;
-- rename temporary table "new_oauth_providers" to "oauth_providers"
ALTER TABLE `new_oauth_providers` RENAME TO `oauth_providers`;
-- create index "oauth_providers_mapping_id_key" to table: "oauth_providers"
CREATE UNIQUE INDEX `oauth_providers_mapping_id_key` ON `oauth_providers` (`mapping_id`);
-- add column "owner_id" to table: "oauth_provider_history"
ALTER TABLE `oauth_provider_history` ADD COLUMN `owner_id` text NULL;
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: add column "owner_id" to table: "oauth_provider_history"
ALTER TABLE `oauth_provider_history` DROP COLUMN `owner_id`;
-- reverse: create index "oauth_providers_mapping_id_key" to table: "oauth_providers"
DROP INDEX `oauth_providers_mapping_id_key`;
-- reverse: create "new_oauth_providers" table
DROP TABLE `new_oauth_providers`;
-- reverse: add column "owner_id" to table: "document_data_history"
ALTER TABLE `document_data_history` DROP COLUMN `owner_id`;
-- reverse: create index "document_data_mapping_id_key" to table: "document_data"
DROP INDEX `document_data_mapping_id_key`;
-- reverse: create "new_document_data" table
DROP TABLE `new_document_data`;
