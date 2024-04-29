-- +goose Up
-- create "api_tokens" table
CREATE TABLE `api_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `organization_id` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NOT NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `api_tokens_organizations_api_tokens` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- create index "api_tokens_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_token_key` ON `api_tokens` (`token`);
-- create index "apitoken_token" to table: "api_tokens"
CREATE INDEX `apitoken_token` ON `api_tokens` (`token`);

-- +goose Down
-- reverse: create index "apitoken_token" to table: "api_tokens"
DROP INDEX `apitoken_token`;
-- reverse: create index "api_tokens_token_key" to table: "api_tokens"
DROP INDEX `api_tokens_token_key`;
-- reverse: create "api_tokens" table
DROP TABLE `api_tokens`;
