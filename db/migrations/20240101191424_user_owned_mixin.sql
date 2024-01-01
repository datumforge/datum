-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Add column "deleted_at" to table: "personal_access_tokens"
ALTER TABLE `personal_access_tokens` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "personal_access_tokens"
ALTER TABLE `personal_access_tokens` ADD COLUMN `deleted_by` text NULL;
-- Create "new_email_verification_tokens" table
CREATE TABLE `new_email_verification_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `email_verification_tokens_users_email_verification_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "email_verification_tokens" to new temporary table "new_email_verification_tokens"
INSERT INTO `new_email_verification_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `ttl`, `email`, `secret`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `ttl`, `email`, `secret` FROM `email_verification_tokens`;
-- Drop "email_verification_tokens" table after copying rows
DROP TABLE `email_verification_tokens`;
-- Rename temporary table "new_email_verification_tokens" to "email_verification_tokens"
ALTER TABLE `new_email_verification_tokens` RENAME TO `email_verification_tokens`;
-- Create index "email_verification_tokens_token_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_token_key` ON `email_verification_tokens` (`token`);
-- Create index "emailverificationtoken_token" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `emailverificationtoken_token` ON `email_verification_tokens` (`token`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
