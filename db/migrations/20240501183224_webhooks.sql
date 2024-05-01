-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_integrations" table
CREATE TABLE `new_integrations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `group_integrations` text NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_groups_integrations` FOREIGN KEY (`group_integrations`) REFERENCES `groups` (`id`) ON DELETE SET NULL, CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "integrations" to new temporary table "new_integrations"
INSERT INTO `new_integrations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `owner_id` FROM `integrations`;
-- Drop "integrations" table after copying rows
DROP TABLE `integrations`;
-- Rename temporary table "new_integrations" to "integrations"
ALTER TABLE `new_integrations` RENAME TO `integrations`;
-- Create "events" table
CREATE TABLE `events` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `event_id` text NOT NULL, `correlation_id` text NOT NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- Create "event_history" table
CREATE TABLE `event_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `event_id` text NOT NULL, `correlation_id` text NOT NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- Create index "eventhistory_history_time" to table: "event_history"
CREATE INDEX `eventhistory_history_time` ON `event_history` (`history_time`);
-- Create "features" table
CREATE TABLE `features` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `global` bool NOT NULL DEFAULT (true), `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, PRIMARY KEY (`id`));
-- Create index "features_name_key" to table: "features"
CREATE UNIQUE INDEX `features_name_key` ON `features` (`name`);
-- Create "feature_history" table
CREATE TABLE `feature_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `global` bool NOT NULL DEFAULT (true), `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, PRIMARY KEY (`id`));
-- Create index "featurehistory_history_time" to table: "feature_history"
CREATE INDEX `featurehistory_history_time` ON `feature_history` (`history_time`);
-- Create "files" table
CREATE TABLE `files` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `file_name` text NOT NULL, `file_extension` text NOT NULL, `file_size` integer NULL, `content_type` text NOT NULL, `store_key` text NOT NULL, `category` text NULL, `annotation` text NULL, `user_files` text NULL, PRIMARY KEY (`id`), CONSTRAINT `files_users_files` FOREIGN KEY (`user_files`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create "file_history" table
CREATE TABLE `file_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `file_name` text NOT NULL, `file_extension` text NOT NULL, `file_size` integer NULL, `content_type` text NOT NULL, `store_key` text NOT NULL, `category` text NULL, `annotation` text NULL, PRIMARY KEY (`id`));
-- Create index "filehistory_history_time" to table: "file_history"
CREATE INDEX `filehistory_history_time` ON `file_history` (`history_time`);
-- Create "webhooks" table
CREATE TABLE `webhooks` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `destination_url` text NOT NULL, `enabled` bool NOT NULL DEFAULT (true), `callback` text NULL, `expires_at` datetime NULL, `secret` blob NULL, `failures` integer NULL DEFAULT (0), `last_error` text NULL, `last_response` text NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `webhooks_organizations_webhooks` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "webhooks_callback_key" to table: "webhooks"
CREATE UNIQUE INDEX `webhooks_callback_key` ON `webhooks` (`callback`);
-- Create index "webhook_name_owner_id" to table: "webhooks"
CREATE UNIQUE INDEX `webhook_name_owner_id` ON `webhooks` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- Create "webhook_history" table
CREATE TABLE `webhook_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NULL, `name` text NOT NULL, `description` text NULL, `destination_url` text NOT NULL, `enabled` bool NOT NULL DEFAULT (true), `callback` text NULL, `expires_at` datetime NULL, `secret` blob NULL, `failures` integer NULL DEFAULT (0), `last_error` text NULL, `last_response` text NULL, PRIMARY KEY (`id`));
-- Create index "webhookhistory_history_time" to table: "webhook_history"
CREATE INDEX `webhookhistory_history_time` ON `webhook_history` (`history_time`);
-- Create "entitlement_features" table
CREATE TABLE `entitlement_features` (`entitlement_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`entitlement_id`, `feature_id`), CONSTRAINT `entitlement_features_entitlement_id` FOREIGN KEY (`entitlement_id`) REFERENCES `entitlements` (`id`) ON DELETE CASCADE, CONSTRAINT `entitlement_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
-- Create "entitlement_events" table
CREATE TABLE `entitlement_events` (`entitlement_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`entitlement_id`, `event_id`), CONSTRAINT `entitlement_events_entitlement_id` FOREIGN KEY (`entitlement_id`) REFERENCES `entitlements` (`id`) ON DELETE CASCADE, CONSTRAINT `entitlement_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "feature_events" table
CREATE TABLE `feature_events` (`feature_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`feature_id`, `event_id`), CONSTRAINT `feature_events_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE, CONSTRAINT `feature_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "group_features" table
CREATE TABLE `group_features` (`group_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`group_id`, `feature_id`), CONSTRAINT `group_features_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
-- Create "group_events" table
CREATE TABLE `group_events` (`group_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`group_id`, `event_id`), CONSTRAINT `group_events_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "group_files" table
CREATE TABLE `group_files` (`group_id` text NOT NULL, `file_id` text NOT NULL, PRIMARY KEY (`group_id`, `file_id`), CONSTRAINT `group_files_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_files_file_id` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE CASCADE);
-- Create "group_membership_events" table
CREATE TABLE `group_membership_events` (`group_membership_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`group_membership_id`, `event_id`), CONSTRAINT `group_membership_events_group_membership_id` FOREIGN KEY (`group_membership_id`) REFERENCES `group_memberships` (`id`) ON DELETE CASCADE, CONSTRAINT `group_membership_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "hush_events" table
CREATE TABLE `hush_events` (`hush_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`hush_id`, `event_id`), CONSTRAINT `hush_events_hush_id` FOREIGN KEY (`hush_id`) REFERENCES `hushes` (`id`) ON DELETE CASCADE, CONSTRAINT `hush_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "integration_oauth2tokens" table
CREATE TABLE `integration_oauth2tokens` (`integration_id` text NOT NULL, `oh_auth_too_token_id` text NOT NULL, PRIMARY KEY (`integration_id`, `oh_auth_too_token_id`), CONSTRAINT `integration_oauth2tokens_integration_id` FOREIGN KEY (`integration_id`) REFERENCES `integrations` (`id`) ON DELETE CASCADE, CONSTRAINT `integration_oauth2tokens_oh_auth_too_token_id` FOREIGN KEY (`oh_auth_too_token_id`) REFERENCES `oh_auth_too_tokens` (`id`) ON DELETE CASCADE);
-- Create "integration_events" table
CREATE TABLE `integration_events` (`integration_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`integration_id`, `event_id`), CONSTRAINT `integration_events_integration_id` FOREIGN KEY (`integration_id`) REFERENCES `integrations` (`id`) ON DELETE CASCADE, CONSTRAINT `integration_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "invite_events" table
CREATE TABLE `invite_events` (`invite_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`invite_id`, `event_id`), CONSTRAINT `invite_events_invite_id` FOREIGN KEY (`invite_id`) REFERENCES `invites` (`id`) ON DELETE CASCADE, CONSTRAINT `invite_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "oh_auth_too_token_events" table
CREATE TABLE `oh_auth_too_token_events` (`oh_auth_too_token_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`oh_auth_too_token_id`, `event_id`), CONSTRAINT `oh_auth_too_token_events_oh_auth_too_token_id` FOREIGN KEY (`oh_auth_too_token_id`) REFERENCES `oh_auth_too_tokens` (`id`) ON DELETE CASCADE, CONSTRAINT `oh_auth_too_token_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "org_membership_events" table
CREATE TABLE `org_membership_events` (`org_membership_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`org_membership_id`, `event_id`), CONSTRAINT `org_membership_events_org_membership_id` FOREIGN KEY (`org_membership_id`) REFERENCES `org_memberships` (`id`) ON DELETE CASCADE, CONSTRAINT `org_membership_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "organization_events" table
CREATE TABLE `organization_events` (`organization_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`organization_id`, `event_id`), CONSTRAINT `organization_events_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "organization_secrets" table
CREATE TABLE `organization_secrets` (`organization_id` text NOT NULL, `hush_id` text NOT NULL, PRIMARY KEY (`organization_id`, `hush_id`), CONSTRAINT `organization_secrets_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_secrets_hush_id` FOREIGN KEY (`hush_id`) REFERENCES `hushes` (`id`) ON DELETE CASCADE);
-- Create "organization_files" table
CREATE TABLE `organization_files` (`organization_id` text NOT NULL, `file_id` text NOT NULL, PRIMARY KEY (`organization_id`, `file_id`), CONSTRAINT `organization_files_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_files_file_id` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE CASCADE);
-- Create "personal_access_token_events" table
CREATE TABLE `personal_access_token_events` (`personal_access_token_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`personal_access_token_id`, `event_id`), CONSTRAINT `personal_access_token_events_personal_access_token_id` FOREIGN KEY (`personal_access_token_id`) REFERENCES `personal_access_tokens` (`id`) ON DELETE CASCADE, CONSTRAINT `personal_access_token_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "subscriber_events" table
CREATE TABLE `subscriber_events` (`subscriber_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`subscriber_id`, `event_id`), CONSTRAINT `subscriber_events_subscriber_id` FOREIGN KEY (`subscriber_id`) REFERENCES `subscribers` (`id`) ON DELETE CASCADE, CONSTRAINT `subscriber_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "user_events" table
CREATE TABLE `user_events` (`user_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`user_id`, `event_id`), CONSTRAINT `user_events_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "user_features" table
CREATE TABLE `user_features` (`user_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`user_id`, `feature_id`), CONSTRAINT `user_features_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
-- Create "webhook_events" table
CREATE TABLE `webhook_events` (`webhook_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`webhook_id`, `event_id`), CONSTRAINT `webhook_events_webhook_id` FOREIGN KEY (`webhook_id`) REFERENCES `webhooks` (`id`) ON DELETE CASCADE, CONSTRAINT `webhook_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
