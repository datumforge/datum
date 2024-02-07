-- Create "email_verification_tokens" table
CREATE TABLE `email_verification_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `user_email_verification_tokens` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `email_verification_tokens_users_email_verification_tokens` FOREIGN KEY (`user_email_verification_tokens`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "email_verification_tokens_token_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_token_key` ON `email_verification_tokens` (`token`);
-- Create index "emailverificationtoken_token" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `emailverificationtoken_token` ON `email_verification_tokens` (`token`) WHERE deleted_at is NULL;
