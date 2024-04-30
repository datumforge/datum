-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_api_tokens" table
CREATE TABLE `new_api_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `api_tokens_organizations_api_tokens` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "api_tokens" to new temporary table "new_api_tokens"
INSERT INTO `new_api_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `token`, `expires_at`, `description`, `scopes`, `last_used_at`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `token`, `expires_at`, `description`, `scopes`, `last_used_at`, `owner_id` FROM `api_tokens`;
-- Drop "api_tokens" table after copying rows
DROP TABLE `api_tokens`;
-- Rename temporary table "new_api_tokens" to "api_tokens"
ALTER TABLE `new_api_tokens` RENAME TO `api_tokens`;
-- Create index "api_tokens_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_token_key` ON `api_tokens` (`token`);
-- Create index "apitoken_token" to table: "api_tokens"
CREATE INDEX `apitoken_token` ON `api_tokens` (`token`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
