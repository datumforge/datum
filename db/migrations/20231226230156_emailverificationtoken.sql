-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_users" table
CREATE TABLE `new_users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `oauth` bool NOT NULL DEFAULT (false), `agree_tos` bool NOT NULL DEFAULT (false), `agree_privacy` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `oauth`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `oauth` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "users_sub_key" to table: "users"
CREATE UNIQUE INDEX `users_sub_key` ON `users` (`sub`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create "email_verification_tokens" table
CREATE TABLE `email_verification_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` text NOT NULL, `user_email_verification_tokens` text NULL, PRIMARY KEY (`id`), CONSTRAINT `email_verification_tokens_users_email_verification_tokens` FOREIGN KEY (`user_email_verification_tokens`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "email_verification_tokens_email_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_email_key` ON `email_verification_tokens` (`email`);
-- Create index "email_verification_tokens_secret_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_secret_key` ON `email_verification_tokens` (`secret`);
-- Create index "email_verification_tokens_user_email_verification_tokens_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_user_email_verification_tokens_key` ON `email_verification_tokens` (`user_email_verification_tokens`);
-- Create index "emailverificationtoken_token" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `emailverificationtoken_token` ON `email_verification_tokens` (`token`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
