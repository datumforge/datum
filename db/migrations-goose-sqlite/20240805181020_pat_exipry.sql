-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_personal_access_tokens" table
CREATE TABLE `new_personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `name` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- copy rows from old table "personal_access_tokens" to new temporary table "new_personal_access_tokens"
INSERT INTO `new_personal_access_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `tags`, `name`, `token`, `expires_at`, `description`, `scopes`, `last_used_at`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `tags`, `name`, `token`, `expires_at`, `description`, `scopes`, `last_used_at`, `owner_id` FROM `personal_access_tokens`;
-- drop "personal_access_tokens" table after copying rows
DROP TABLE `personal_access_tokens`;
-- rename temporary table "new_personal_access_tokens" to "personal_access_tokens"
ALTER TABLE `new_personal_access_tokens` RENAME TO `personal_access_tokens`;
-- create index "personal_access_tokens_mapping_id_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_mapping_id_key` ON `personal_access_tokens` (`mapping_id`);
-- create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "personalaccesstoken_token" to table: "personal_access_tokens"
DROP INDEX `personalaccesstoken_token`;
-- reverse: create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
DROP INDEX `personal_access_tokens_token_key`;
-- reverse: create index "personal_access_tokens_mapping_id_key" to table: "personal_access_tokens"
DROP INDEX `personal_access_tokens_mapping_id_key`;
-- reverse: create "new_personal_access_tokens" table
DROP TABLE `new_personal_access_tokens`;
