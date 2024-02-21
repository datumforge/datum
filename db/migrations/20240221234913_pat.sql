-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_personal_access_tokens" table
CREATE TABLE `new_personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NOT NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "personal_access_tokens" to new temporary table "new_personal_access_tokens"
INSERT INTO `new_personal_access_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `token`, `expires_at`, `description`, `last_used_at`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `token`, `expires_at`, `description`, `last_used_at`, `owner_id` FROM `personal_access_tokens`;
-- Drop "personal_access_tokens" table after copying rows
DROP TABLE `personal_access_tokens`;
-- Rename temporary table "new_personal_access_tokens" to "personal_access_tokens"
ALTER TABLE `new_personal_access_tokens` RENAME TO `personal_access_tokens`;
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Create "organization_personal_access_tokens" table
CREATE TABLE `organization_personal_access_tokens` (`organization_id` text NOT NULL, `personal_access_token_id` text NOT NULL, PRIMARY KEY (`organization_id`, `personal_access_token_id`), CONSTRAINT `organization_personal_access_tokens_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_personal_access_tokens_personal_access_token_id` FOREIGN KEY (`personal_access_token_id`) REFERENCES `personal_access_tokens` (`id`) ON DELETE CASCADE);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
