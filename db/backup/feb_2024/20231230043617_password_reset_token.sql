-- Create "password_reset_tokens" table
CREATE TABLE `password_reset_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `user_reset_tokens` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `password_reset_tokens_users_reset_tokens` FOREIGN KEY (`user_reset_tokens`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "password_reset_tokens_token_key" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `password_reset_tokens_token_key` ON `password_reset_tokens` (`token`);
-- Create index "passwordresettoken_token" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `passwordresettoken_token` ON `password_reset_tokens` (`token`) WHERE deleted_at is NULL;
