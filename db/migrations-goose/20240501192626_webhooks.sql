-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_integrations" table
CREATE TABLE `new_integrations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `group_integrations` text NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_groups_integrations` FOREIGN KEY (`group_integrations`) REFERENCES `groups` (`id`) ON DELETE SET NULL, CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- copy rows from old table "integrations" to new temporary table "new_integrations"
INSERT INTO `new_integrations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `owner_id` FROM `integrations`;
-- drop "integrations" table after copying rows
DROP TABLE `integrations`;
-- rename temporary table "new_integrations" to "integrations"
ALTER TABLE `new_integrations` RENAME TO `integrations`;
-- create "events" table
CREATE TABLE `events` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `event_id` text NOT NULL, `correlation_id` text NOT NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- create "event_history" table
CREATE TABLE `event_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `event_id` text NOT NULL, `correlation_id` text NOT NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- create index "eventhistory_history_time" to table: "event_history"
CREATE INDEX `eventhistory_history_time` ON `event_history` (`history_time`);
-- create "features" table
CREATE TABLE `features` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `global` bool NOT NULL DEFAULT (true), `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, PRIMARY KEY (`id`));
-- create index "features_name_key" to table: "features"
CREATE UNIQUE INDEX `features_name_key` ON `features` (`name`);
-- create "feature_history" table
CREATE TABLE `feature_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `global` bool NOT NULL DEFAULT (true), `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, PRIMARY KEY (`id`));
-- create index "featurehistory_history_time" to table: "feature_history"
CREATE INDEX `featurehistory_history_time` ON `feature_history` (`history_time`);
-- create "files" table
CREATE TABLE `files` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `file_name` text NOT NULL, `file_extension` text NOT NULL, `file_size` integer NULL, `content_type` text NOT NULL, `store_key` text NOT NULL, `category` text NULL, `annotation` text NULL, `user_files` text NULL, PRIMARY KEY (`id`), CONSTRAINT `files_users_files` FOREIGN KEY (`user_files`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- create "file_history" table
CREATE TABLE `file_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `file_name` text NOT NULL, `file_extension` text NOT NULL, `file_size` integer NULL, `content_type` text NOT NULL, `store_key` text NOT NULL, `category` text NULL, `annotation` text NULL, PRIMARY KEY (`id`));
-- create index "filehistory_history_time" to table: "file_history"
CREATE INDEX `filehistory_history_time` ON `file_history` (`history_time`);
-- create "webhooks" table
CREATE TABLE `webhooks` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `destination_url` text NOT NULL, `enabled` bool NOT NULL DEFAULT (true), `callback` text NULL, `expires_at` datetime NULL, `secret` blob NULL, `failures` integer NULL DEFAULT (0), `last_error` text NULL, `last_response` text NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `webhooks_organizations_webhooks` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- create index "webhooks_callback_key" to table: "webhooks"
CREATE UNIQUE INDEX `webhooks_callback_key` ON `webhooks` (`callback`);
-- create index "webhook_name_owner_id" to table: "webhooks"
CREATE UNIQUE INDEX `webhook_name_owner_id` ON `webhooks` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- create "webhook_history" table
CREATE TABLE `webhook_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NULL, `name` text NOT NULL, `description` text NULL, `destination_url` text NOT NULL, `enabled` bool NOT NULL DEFAULT (true), `callback` text NULL, `expires_at` datetime NULL, `secret` blob NULL, `failures` integer NULL DEFAULT (0), `last_error` text NULL, `last_response` text NULL, PRIMARY KEY (`id`));
-- create index "webhookhistory_history_time" to table: "webhook_history"
CREATE INDEX `webhookhistory_history_time` ON `webhook_history` (`history_time`);
-- create "entitlement_features" table
CREATE TABLE `entitlement_features` (`entitlement_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`entitlement_id`, `feature_id`), CONSTRAINT `entitlement_features_entitlement_id` FOREIGN KEY (`entitlement_id`) REFERENCES `entitlements` (`id`) ON DELETE CASCADE, CONSTRAINT `entitlement_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
-- create "entitlement_events" table
CREATE TABLE `entitlement_events` (`entitlement_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`entitlement_id`, `event_id`), CONSTRAINT `entitlement_events_entitlement_id` FOREIGN KEY (`entitlement_id`) REFERENCES `entitlements` (`id`) ON DELETE CASCADE, CONSTRAINT `entitlement_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "feature_events" table
CREATE TABLE `feature_events` (`feature_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`feature_id`, `event_id`), CONSTRAINT `feature_events_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE, CONSTRAINT `feature_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "group_features" table
CREATE TABLE `group_features` (`group_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`group_id`, `feature_id`), CONSTRAINT `group_features_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
-- create "group_events" table
CREATE TABLE `group_events` (`group_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`group_id`, `event_id`), CONSTRAINT `group_events_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "group_files" table
CREATE TABLE `group_files` (`group_id` text NOT NULL, `file_id` text NOT NULL, PRIMARY KEY (`group_id`, `file_id`), CONSTRAINT `group_files_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_files_file_id` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE CASCADE);
-- create "group_membership_events" table
CREATE TABLE `group_membership_events` (`group_membership_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`group_membership_id`, `event_id`), CONSTRAINT `group_membership_events_group_membership_id` FOREIGN KEY (`group_membership_id`) REFERENCES `group_memberships` (`id`) ON DELETE CASCADE, CONSTRAINT `group_membership_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "hush_events" table
CREATE TABLE `hush_events` (`hush_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`hush_id`, `event_id`), CONSTRAINT `hush_events_hush_id` FOREIGN KEY (`hush_id`) REFERENCES `hushes` (`id`) ON DELETE CASCADE, CONSTRAINT `hush_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "integration_oauth2tokens" table
CREATE TABLE `integration_oauth2tokens` (`integration_id` text NOT NULL, `oh_auth_too_token_id` text NOT NULL, PRIMARY KEY (`integration_id`, `oh_auth_too_token_id`), CONSTRAINT `integration_oauth2tokens_integration_id` FOREIGN KEY (`integration_id`) REFERENCES `integrations` (`id`) ON DELETE CASCADE, CONSTRAINT `integration_oauth2tokens_oh_auth_too_token_id` FOREIGN KEY (`oh_auth_too_token_id`) REFERENCES `oh_auth_too_tokens` (`id`) ON DELETE CASCADE);
-- create "integration_events" table
CREATE TABLE `integration_events` (`integration_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`integration_id`, `event_id`), CONSTRAINT `integration_events_integration_id` FOREIGN KEY (`integration_id`) REFERENCES `integrations` (`id`) ON DELETE CASCADE, CONSTRAINT `integration_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "invite_events" table
CREATE TABLE `invite_events` (`invite_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`invite_id`, `event_id`), CONSTRAINT `invite_events_invite_id` FOREIGN KEY (`invite_id`) REFERENCES `invites` (`id`) ON DELETE CASCADE, CONSTRAINT `invite_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "oh_auth_too_token_events" table
CREATE TABLE `oh_auth_too_token_events` (`oh_auth_too_token_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`oh_auth_too_token_id`, `event_id`), CONSTRAINT `oh_auth_too_token_events_oh_auth_too_token_id` FOREIGN KEY (`oh_auth_too_token_id`) REFERENCES `oh_auth_too_tokens` (`id`) ON DELETE CASCADE, CONSTRAINT `oh_auth_too_token_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "org_membership_events" table
CREATE TABLE `org_membership_events` (`org_membership_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`org_membership_id`, `event_id`), CONSTRAINT `org_membership_events_org_membership_id` FOREIGN KEY (`org_membership_id`) REFERENCES `org_memberships` (`id`) ON DELETE CASCADE, CONSTRAINT `org_membership_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "organization_events" table
CREATE TABLE `organization_events` (`organization_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`organization_id`, `event_id`), CONSTRAINT `organization_events_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "organization_secrets" table
CREATE TABLE `organization_secrets` (`organization_id` text NOT NULL, `hush_id` text NOT NULL, PRIMARY KEY (`organization_id`, `hush_id`), CONSTRAINT `organization_secrets_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_secrets_hush_id` FOREIGN KEY (`hush_id`) REFERENCES `hushes` (`id`) ON DELETE CASCADE);
-- create "organization_features" table
CREATE TABLE `organization_features` (`organization_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`organization_id`, `feature_id`), CONSTRAINT `organization_features_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
-- create "organization_files" table
CREATE TABLE `organization_files` (`organization_id` text NOT NULL, `file_id` text NOT NULL, PRIMARY KEY (`organization_id`, `file_id`), CONSTRAINT `organization_files_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_files_file_id` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE CASCADE);
-- create "personal_access_token_events" table
CREATE TABLE `personal_access_token_events` (`personal_access_token_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`personal_access_token_id`, `event_id`), CONSTRAINT `personal_access_token_events_personal_access_token_id` FOREIGN KEY (`personal_access_token_id`) REFERENCES `personal_access_tokens` (`id`) ON DELETE CASCADE, CONSTRAINT `personal_access_token_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "subscriber_events" table
CREATE TABLE `subscriber_events` (`subscriber_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`subscriber_id`, `event_id`), CONSTRAINT `subscriber_events_subscriber_id` FOREIGN KEY (`subscriber_id`) REFERENCES `subscribers` (`id`) ON DELETE CASCADE, CONSTRAINT `subscriber_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "user_events" table
CREATE TABLE `user_events` (`user_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`user_id`, `event_id`), CONSTRAINT `user_events_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "user_features" table
CREATE TABLE `user_features` (`user_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`user_id`, `feature_id`), CONSTRAINT `user_features_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
-- create "webhook_events" table
CREATE TABLE `webhook_events` (`webhook_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`webhook_id`, `event_id`), CONSTRAINT `webhook_events_webhook_id` FOREIGN KEY (`webhook_id`) REFERENCES `webhooks` (`id`) ON DELETE CASCADE, CONSTRAINT `webhook_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create "webhook_events" table
DROP TABLE `webhook_events`;
-- reverse: create "user_features" table
DROP TABLE `user_features`;
-- reverse: create "user_events" table
DROP TABLE `user_events`;
-- reverse: create "subscriber_events" table
DROP TABLE `subscriber_events`;
-- reverse: create "personal_access_token_events" table
DROP TABLE `personal_access_token_events`;
-- reverse: create "organization_files" table
DROP TABLE `organization_files`;
-- reverse: create "organization_features" table
DROP TABLE `organization_features`;
-- reverse: create "organization_secrets" table
DROP TABLE `organization_secrets`;
-- reverse: create "organization_events" table
DROP TABLE `organization_events`;
-- reverse: create "org_membership_events" table
DROP TABLE `org_membership_events`;
-- reverse: create "oh_auth_too_token_events" table
DROP TABLE `oh_auth_too_token_events`;
-- reverse: create "invite_events" table
DROP TABLE `invite_events`;
-- reverse: create "integration_events" table
DROP TABLE `integration_events`;
-- reverse: create "integration_oauth2tokens" table
DROP TABLE `integration_oauth2tokens`;
-- reverse: create "hush_events" table
DROP TABLE `hush_events`;
-- reverse: create "group_membership_events" table
DROP TABLE `group_membership_events`;
-- reverse: create "group_files" table
DROP TABLE `group_files`;
-- reverse: create "group_events" table
DROP TABLE `group_events`;
-- reverse: create "group_features" table
DROP TABLE `group_features`;
-- reverse: create "feature_events" table
DROP TABLE `feature_events`;
-- reverse: create "entitlement_events" table
DROP TABLE `entitlement_events`;
-- reverse: create "entitlement_features" table
DROP TABLE `entitlement_features`;
-- reverse: create index "webhookhistory_history_time" to table: "webhook_history"
DROP INDEX `webhookhistory_history_time`;
-- reverse: create "webhook_history" table
DROP TABLE `webhook_history`;
-- reverse: create index "webhook_name_owner_id" to table: "webhooks"
DROP INDEX `webhook_name_owner_id`;
-- reverse: create index "webhooks_callback_key" to table: "webhooks"
DROP INDEX `webhooks_callback_key`;
-- reverse: create "webhooks" table
DROP TABLE `webhooks`;
-- reverse: create index "filehistory_history_time" to table: "file_history"
DROP INDEX `filehistory_history_time`;
-- reverse: create "file_history" table
DROP TABLE `file_history`;
-- reverse: create "files" table
DROP TABLE `files`;
-- reverse: create index "featurehistory_history_time" to table: "feature_history"
DROP INDEX `featurehistory_history_time`;
-- reverse: create "feature_history" table
DROP TABLE `feature_history`;
-- reverse: create index "features_name_key" to table: "features"
DROP INDEX `features_name_key`;
-- reverse: create "features" table
DROP TABLE `features`;
-- reverse: create index "eventhistory_history_time" to table: "event_history"
DROP INDEX `eventhistory_history_time`;
-- reverse: create "event_history" table
DROP TABLE `event_history`;
-- reverse: create "events" table
DROP TABLE `events`;
-- reverse: create "new_integrations" table
DROP TABLE `new_integrations`;
