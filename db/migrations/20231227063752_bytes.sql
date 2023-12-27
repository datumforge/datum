-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_email_verification_tokens" table
CREATE TABLE `new_email_verification_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NULL, `ttl` datetime NULL, `email` text NULL, `secret` blob NULL, `user_email_verification_tokens` text NOT NULL, `user_children` text NULL, PRIMARY KEY (`id`), CONSTRAINT `email_verification_tokens_users_email_verification_tokens` FOREIGN KEY (`user_email_verification_tokens`) REFERENCES `users` (`id`) ON DELETE NO ACTION, CONSTRAINT `email_verification_tokens_users_children` FOREIGN KEY (`user_children`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "email_verification_tokens" to new temporary table "new_email_verification_tokens"
INSERT INTO `new_email_verification_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `ttl`, `email`, `secret`, `user_email_verification_tokens`, `user_children`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `ttl`, `email`, `secret`, `user_email_verification_tokens`, `user_children` FROM `email_verification_tokens`;
-- Drop "email_verification_tokens" table after copying rows
DROP TABLE `email_verification_tokens`;
-- Rename temporary table "new_email_verification_tokens" to "email_verification_tokens"
ALTER TABLE `new_email_verification_tokens` RENAME TO `email_verification_tokens`;
-- Create index "email_verification_tokens_user_email_verification_tokens_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_user_email_verification_tokens_key` ON `email_verification_tokens` (`user_email_verification_tokens`);
-- Create index "emailverificationtoken_token" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `emailverificationtoken_token` ON `email_verification_tokens` (`token`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
