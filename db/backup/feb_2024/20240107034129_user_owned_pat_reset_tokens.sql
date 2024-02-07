-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_personal_access_tokens" table
CREATE TABLE `new_personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `abilities` json NULL, `expires_at` datetime NOT NULL, `description` text NULL DEFAULT (''), `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "personal_access_tokens" to new temporary table "new_personal_access_tokens"
INSERT INTO `new_personal_access_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `token`, `abilities`, `expires_at`, `description`, `last_used_at`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `token`, `abilities`, `expires_at`, `description`, `last_used_at` FROM `personal_access_tokens`;
-- Drop "personal_access_tokens" table after copying rows
DROP TABLE `personal_access_tokens`;
-- Rename temporary table "new_personal_access_tokens" to "personal_access_tokens"
ALTER TABLE `new_personal_access_tokens` RENAME TO `personal_access_tokens`;
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Create "new_password_reset_tokens" table
CREATE TABLE `new_password_reset_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `password_reset_tokens_users_password_reset_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "password_reset_tokens" to new temporary table "new_password_reset_tokens"
INSERT INTO `new_password_reset_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `ttl`, `email`, `secret`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `ttl`, `email`, `secret` FROM `password_reset_tokens`;
-- Drop "password_reset_tokens" table after copying rows
DROP TABLE `password_reset_tokens`;
-- Rename temporary table "new_password_reset_tokens" to "password_reset_tokens"
ALTER TABLE `new_password_reset_tokens` RENAME TO `password_reset_tokens`;
-- Create index "password_reset_tokens_token_key" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `password_reset_tokens_token_key` ON `password_reset_tokens` (`token`);
-- Create index "passwordresettoken_token" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `passwordresettoken_token` ON `password_reset_tokens` (`token`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
