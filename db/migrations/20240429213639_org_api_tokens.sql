-- Create "api_tokens" table
CREATE TABLE `api_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `organization_id` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NOT NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `api_tokens_organizations_api_tokens` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "api_tokens_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_token_key` ON `api_tokens` (`token`);
-- Create index "apitoken_token" to table: "api_tokens"
CREATE INDEX `apitoken_token` ON `api_tokens` (`token`);
