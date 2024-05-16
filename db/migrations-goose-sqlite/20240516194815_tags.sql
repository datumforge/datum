-- +goose Up
-- add column "tags" to table: "api_tokens"
ALTER TABLE `api_tokens` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "document_data"
ALTER TABLE `document_data` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "document_data_history"
ALTER TABLE `document_data_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "entitlements"
ALTER TABLE `entitlements` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "entitlement_history"
ALTER TABLE `entitlement_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "events"
ALTER TABLE `events` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "event_history"
ALTER TABLE `event_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "features"
ALTER TABLE `features` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "feature_history"
ALTER TABLE `feature_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "files"
ALTER TABLE `files` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "file_history"
ALTER TABLE `file_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "groups"
ALTER TABLE `groups` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "group_history"
ALTER TABLE `group_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "integrations"
ALTER TABLE `integrations` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "integration_history"
ALTER TABLE `integration_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "oauth_providers"
ALTER TABLE `oauth_providers` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "oauth_provider_history"
ALTER TABLE `oauth_provider_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "oh_auth_too_tokens"
ALTER TABLE `oh_auth_too_tokens` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "organizations"
ALTER TABLE `organizations` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "organization_history"
ALTER TABLE `organization_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "personal_access_tokens"
ALTER TABLE `personal_access_tokens` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "subscribers"
ALTER TABLE `subscribers` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "tfa_settings"
ALTER TABLE `tfa_settings` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "templates"
ALTER TABLE `templates` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "template_history"
ALTER TABLE `template_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "users"
ALTER TABLE `users` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "user_history"
ALTER TABLE `user_history` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "webauthns"
ALTER TABLE `webauthns` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "webhooks"
ALTER TABLE `webhooks` ADD COLUMN `tags` json NULL;
-- add column "tags" to table: "webhook_history"
ALTER TABLE `webhook_history` ADD COLUMN `tags` json NULL;

-- +goose Down
-- reverse: add column "tags" to table: "webhook_history"
ALTER TABLE `webhook_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "webhooks"
ALTER TABLE `webhooks` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "webauthns"
ALTER TABLE `webauthns` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "user_history"
ALTER TABLE `user_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "users"
ALTER TABLE `users` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "template_history"
ALTER TABLE `template_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "templates"
ALTER TABLE `templates` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "tfa_settings"
ALTER TABLE `tfa_settings` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "subscribers"
ALTER TABLE `subscribers` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "personal_access_tokens"
ALTER TABLE `personal_access_tokens` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "organization_history"
ALTER TABLE `organization_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "organizations"
ALTER TABLE `organizations` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "oh_auth_too_tokens"
ALTER TABLE `oh_auth_too_tokens` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "oauth_provider_history"
ALTER TABLE `oauth_provider_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "oauth_providers"
ALTER TABLE `oauth_providers` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "integration_history"
ALTER TABLE `integration_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "integrations"
ALTER TABLE `integrations` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "group_history"
ALTER TABLE `group_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "groups"
ALTER TABLE `groups` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "file_history"
ALTER TABLE `file_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "files"
ALTER TABLE `files` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "feature_history"
ALTER TABLE `feature_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "features"
ALTER TABLE `features` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "event_history"
ALTER TABLE `event_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "events"
ALTER TABLE `events` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "entitlement_history"
ALTER TABLE `entitlement_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "entitlements"
ALTER TABLE `entitlements` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "document_data_history"
ALTER TABLE `document_data_history` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "document_data"
ALTER TABLE `document_data` DROP COLUMN `tags`;
-- reverse: add column "tags" to table: "api_tokens"
ALTER TABLE `api_tokens` DROP COLUMN `tags`;
