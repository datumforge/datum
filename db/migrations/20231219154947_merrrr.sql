-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_personal_access_tokens" table
CREATE TABLE `new_personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `abilities` json NULL, `expires_at` datetime NOT NULL, `description` text NOT NULL DEFAULT (''), `last_used_at` datetime NULL, `user_personal_access_tokens` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`user_personal_access_tokens`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "personal_access_tokens" to new temporary table "new_personal_access_tokens"
INSERT INTO `new_personal_access_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `token`, `abilities`, `expires_at`, `description`, `last_used_at`, `user_personal_access_tokens`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `token`, `abilities`, `expires_at`, `description`, `last_used_at`, `user_personal_access_tokens` FROM `personal_access_tokens`;
-- Drop "personal_access_tokens" table after copying rows
DROP TABLE `personal_access_tokens`;
-- Rename temporary table "new_personal_access_tokens" to "personal_access_tokens"
ALTER TABLE `new_personal_access_tokens` RENAME TO `personal_access_tokens`;
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
