-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_refresh_tokens" table
CREATE TABLE `new_refresh_tokens` (`id` text NOT NULL, `refresh_token` text NOT NULL, `expires_at` datetime NOT NULL, `issued_at` datetime NOT NULL, `organization_id` text NOT NULL, `user_id` text NOT NULL, `user_refresh_token` text NULL, PRIMARY KEY (`id`), CONSTRAINT `refresh_tokens_users_refresh_token` FOREIGN KEY (`user_refresh_token`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "refresh_tokens" to new temporary table "new_refresh_tokens"
INSERT INTO `new_refresh_tokens` (`id`) SELECT `id` FROM `refresh_tokens`;
-- Drop "refresh_tokens" table after copying rows
DROP TABLE `refresh_tokens`;
-- Rename temporary table "new_refresh_tokens" to "refresh_tokens"
ALTER TABLE `new_refresh_tokens` RENAME TO `refresh_tokens`;
-- Create index "refresh_tokens_refresh_token_key" to table: "refresh_tokens"
CREATE UNIQUE INDEX `refresh_tokens_refresh_token_key` ON `refresh_tokens` (`refresh_token`);
-- Create "new_sessions" table
CREATE TABLE `new_sessions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `session_token` text NOT NULL, `issued_at` datetime NOT NULL, `expires_at` datetime NULL, `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "sessions" to new temporary table "new_sessions"
INSERT INTO `new_sessions` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by` FROM `sessions`;
-- Drop "sessions" table after copying rows
DROP TABLE `sessions`;
-- Rename temporary table "new_sessions" to "sessions"
ALTER TABLE `new_sessions` RENAME TO `sessions`;
-- Create index "sessions_session_token_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_session_token_key` ON `sessions` (`session_token`);
-- Create index "session_session_token" to table: "sessions"
CREATE UNIQUE INDEX `session_session_token` ON `sessions` (`session_token`);
-- Create "access_tokens" table
CREATE TABLE `access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `access_token` text NOT NULL, `expires_at` datetime NOT NULL, `issued_at` datetime NOT NULL, `last_used_at` datetime NULL, `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `access_tokens_users_access_token` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "access_tokens_access_token_key" to table: "access_tokens"
CREATE UNIQUE INDEX `access_tokens_access_token_key` ON `access_tokens` (`access_token`);
-- Create index "accesstoken_access_token" to table: "access_tokens"
CREATE INDEX `accesstoken_access_token` ON `access_tokens` (`access_token`);
-- Create "oh_auth_too_tokens" table
CREATE TABLE `oh_auth_too_tokens` (`id` text NOT NULL, `client_id` text NOT NULL, `scopes` json NULL, `nonce` text NOT NULL, `claims_user_id` text NOT NULL, `claims_username` text NOT NULL, `claims_email` text NOT NULL, `claims_email_verified` bool NOT NULL, `claims_groups` json NULL, `claims_preferred_username` text NOT NULL, `connector_id` text NOT NULL, `connector_data` json NULL, `last_used` datetime NOT NULL, PRIMARY KEY (`id`));
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
