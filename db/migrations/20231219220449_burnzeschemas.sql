-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_sessions" table
CREATE TABLE `new_sessions` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `session_token` text NOT NULL, `issued_at` datetime NOT NULL, `expires_at` datetime NOT NULL, `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "sessions" to new temporary table "new_sessions"
INSERT INTO `new_sessions` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `session_token`, `issued_at`, `expires_at`, `organization_id`, `user_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `session_token`, `issued_at`, `expires_at`, `organization_id`, `user_id` FROM `sessions`;
-- Drop "sessions" table after copying rows
DROP TABLE `sessions`;
-- Rename temporary table "new_sessions" to "sessions"
ALTER TABLE `new_sessions` RENAME TO `sessions`;
-- Create index "sessions_session_token_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_session_token_key` ON `sessions` (`session_token`);
-- Create index "session_session_token" to table: "sessions"
CREATE UNIQUE INDEX `session_session_token` ON `sessions` (`session_token`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
