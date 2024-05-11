-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_api_tokens" table
CREATE TABLE `new_api_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `api_tokens_organizations_api_tokens` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "api_tokens" to new temporary table "new_api_tokens"
INSERT INTO `new_api_tokens` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `name`, `token`, `expires_at`, `description`, `scopes`, `last_used_at`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `name`, `token`, `expires_at`, `description`, `scopes`, `last_used_at`, `owner_id` FROM `api_tokens`;
-- Drop "api_tokens" table after copying rows
DROP TABLE `api_tokens`;
-- Rename temporary table "new_api_tokens" to "api_tokens"
ALTER TABLE `new_api_tokens` RENAME TO `api_tokens`;
-- Create index "api_tokens_mapping_id_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_mapping_id_key` ON `api_tokens` (`mapping_id`);
-- Create index "api_tokens_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_token_key` ON `api_tokens` (`token`);
-- Create index "apitoken_token" to table: "api_tokens"
CREATE INDEX `apitoken_token` ON `api_tokens` (`token`);
-- Create "new_entitlements" table
CREATE TABLE `new_entitlements` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "entitlements" to new temporary table "new_entitlements"
INSERT INTO `new_entitlements` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `tier`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `tier`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`, `owner_id` FROM `entitlements`;
-- Drop "entitlements" table after copying rows
DROP TABLE `entitlements`;
-- Rename temporary table "new_entitlements" to "entitlements"
ALTER TABLE `new_entitlements` RENAME TO `entitlements`;
-- Create index "entitlements_mapping_id_key" to table: "entitlements"
CREATE UNIQUE INDEX `entitlements_mapping_id_key` ON `entitlements` (`mapping_id`);
-- Create "new_entitlement_history" table
CREATE TABLE `new_entitlement_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Copy rows from old table "entitlement_history" to new temporary table "new_entitlement_history"
INSERT INTO `new_entitlement_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `owner_id`, `tier`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `owner_id`, `tier`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled` FROM `entitlement_history`;
-- Drop "entitlement_history" table after copying rows
DROP TABLE `entitlement_history`;
-- Rename temporary table "new_entitlement_history" to "entitlement_history"
ALTER TABLE `new_entitlement_history` RENAME TO `entitlement_history`;
-- Create index "entitlementhistory_history_time" to table: "entitlement_history"
CREATE INDEX `entitlementhistory_history_time` ON `entitlement_history` (`history_time`);
-- Create "new_groups" table
CREATE TABLE `new_groups` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "groups" to new temporary table "new_groups"
INSERT INTO `new_groups` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `name`, `description`, `gravatar_logo_url`, `logo_url`, `display_name`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `name`, `description`, `gravatar_logo_url`, `logo_url`, `display_name`, `owner_id` FROM `groups`;
-- Drop "groups" table after copying rows
DROP TABLE `groups`;
-- Rename temporary table "new_groups" to "groups"
ALTER TABLE `new_groups` RENAME TO `groups`;
-- Create index "groups_mapping_id_key" to table: "groups"
CREATE UNIQUE INDEX `groups_mapping_id_key` ON `groups` (`mapping_id`);
-- Create index "group_name_owner_id" to table: "groups"
CREATE UNIQUE INDEX `group_name_owner_id` ON `groups` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- Create "new_group_history" table
CREATE TABLE `new_group_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `owner_id` text NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), PRIMARY KEY (`id`));
-- Copy rows from old table "group_history" to new temporary table "new_group_history"
INSERT INTO `new_group_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `owner_id`, `name`, `description`, `gravatar_logo_url`, `logo_url`, `display_name`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `owner_id`, `name`, `description`, `gravatar_logo_url`, `logo_url`, `display_name` FROM `group_history`;
-- Drop "group_history" table after copying rows
DROP TABLE `group_history`;
-- Rename temporary table "new_group_history" to "group_history"
ALTER TABLE `new_group_history` RENAME TO `group_history`;
-- Create index "grouphistory_history_time" to table: "group_history"
CREATE INDEX `grouphistory_history_time` ON `group_history` (`history_time`);
-- Create "new_integrations" table
CREATE TABLE `new_integrations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `group_integrations` text NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_groups_integrations` FOREIGN KEY (`group_integrations`) REFERENCES `groups` (`id`) ON DELETE SET NULL, CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "integrations" to new temporary table "new_integrations"
INSERT INTO `new_integrations` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `group_integrations`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `name`, `description`, `kind`, `group_integrations`, `owner_id` FROM `integrations`;
-- Drop "integrations" table after copying rows
DROP TABLE `integrations`;
-- Rename temporary table "new_integrations" to "integrations"
ALTER TABLE `new_integrations` RENAME TO `integrations`;
-- Create index "integrations_mapping_id_key" to table: "integrations"
CREATE UNIQUE INDEX `integrations_mapping_id_key` ON `integrations` (`mapping_id`);
-- Create "new_integration_history" table
CREATE TABLE `new_integration_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "integration_history" to new temporary table "new_integration_history"
INSERT INTO `new_integration_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `owner_id`, `name`, `description`, `kind`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `owner_id`, `name`, `description`, `kind` FROM `integration_history`;
-- Drop "integration_history" table after copying rows
DROP TABLE `integration_history`;
-- Rename temporary table "new_integration_history" to "integration_history"
ALTER TABLE `new_integration_history` RENAME TO `integration_history`;
-- Create index "integrationhistory_history_time" to table: "integration_history"
CREATE INDEX `integrationhistory_history_time` ON `integration_history` (`history_time`);
-- Create "new_invites" table
CREATE TABLE `new_invites` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `expires` datetime NOT NULL, `recipient` text NOT NULL, `status` text NOT NULL DEFAULT ('INVITATION_SENT'), `role` text NOT NULL DEFAULT ('MEMBER'), `send_attempts` integer NOT NULL DEFAULT (0), `requestor_id` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `invites_organizations_invites` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "invites" to new temporary table "new_invites"
INSERT INTO `new_invites` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `token`, `expires`, `recipient`, `status`, `role`, `send_attempts`, `requestor_id`, `secret`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `token`, `expires`, `recipient`, `status`, `role`, `send_attempts`, `requestor_id`, `secret`, `owner_id` FROM `invites`;
-- Drop "invites" table after copying rows
DROP TABLE `invites`;
-- Rename temporary table "new_invites" to "invites"
ALTER TABLE `new_invites` RENAME TO `invites`;
-- Create index "invites_mapping_id_key" to table: "invites"
CREATE UNIQUE INDEX `invites_mapping_id_key` ON `invites` (`mapping_id`);
-- Create index "invites_token_key" to table: "invites"
CREATE UNIQUE INDEX `invites_token_key` ON `invites` (`token`);
-- Create index "invite_recipient_owner_id" to table: "invites"
CREATE UNIQUE INDEX `invite_recipient_owner_id` ON `invites` (`recipient`, `owner_id`) WHERE deleted_at is NULL;
-- Create "new_templates" table
CREATE TABLE `new_templates` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `template_type` text NOT NULL DEFAULT ('DOCUMENT'), `description` text NULL, `jsonconfig` json NOT NULL, `uischema` json NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `templates_organizations_templates` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "templates" to new temporary table "new_templates"
INSERT INTO `new_templates` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `name`, `template_type`, `description`, `jsonconfig`, `uischema`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `name`, `template_type`, `description`, `jsonconfig`, `uischema`, `owner_id` FROM `templates`;
-- Drop "templates" table after copying rows
DROP TABLE `templates`;
-- Rename temporary table "new_templates" to "templates"
ALTER TABLE `new_templates` RENAME TO `templates`;
-- Create index "templates_mapping_id_key" to table: "templates"
CREATE UNIQUE INDEX `templates_mapping_id_key` ON `templates` (`mapping_id`);
-- Create index "template_name_owner_id_template_type" to table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id_template_type` ON `templates` (`name`, `owner_id`, `template_type`) WHERE deleted_at is NULL;
-- Create "new_template_history" table
CREATE TABLE `new_template_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `owner_id` text NULL, `name` text NOT NULL, `template_type` text NOT NULL DEFAULT ('DOCUMENT'), `description` text NULL, `jsonconfig` json NOT NULL, `uischema` json NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "template_history" to new temporary table "new_template_history"
INSERT INTO `new_template_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `owner_id`, `name`, `template_type`, `description`, `jsonconfig`, `uischema`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `owner_id`, `name`, `template_type`, `description`, `jsonconfig`, `uischema` FROM `template_history`;
-- Drop "template_history" table after copying rows
DROP TABLE `template_history`;
-- Rename temporary table "new_template_history" to "template_history"
ALTER TABLE `new_template_history` RENAME TO `template_history`;
-- Create index "templatehistory_history_time" to table: "template_history"
CREATE INDEX `templatehistory_history_time` ON `template_history` (`history_time`);
-- Create "new_subscribers" table
CREATE TABLE `new_subscribers` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `phone_number` text NULL, `verified_email` bool NOT NULL DEFAULT (false), `verified_phone` bool NOT NULL DEFAULT (false), `active` bool NOT NULL DEFAULT (false), `token` text NOT NULL, `ttl` datetime NOT NULL, `secret` blob NOT NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `subscribers_organizations_subscribers` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "subscribers" to new temporary table "new_subscribers"
INSERT INTO `new_subscribers` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `email`, `phone_number`, `verified_email`, `verified_phone`, `active`, `token`, `ttl`, `secret`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `email`, `phone_number`, `verified_email`, `verified_phone`, `active`, `token`, `ttl`, `secret`, `owner_id` FROM `subscribers`;
-- Drop "subscribers" table after copying rows
DROP TABLE `subscribers`;
-- Rename temporary table "new_subscribers" to "subscribers"
ALTER TABLE `new_subscribers` RENAME TO `subscribers`;
-- Create index "subscribers_mapping_id_key" to table: "subscribers"
CREATE UNIQUE INDEX `subscribers_mapping_id_key` ON `subscribers` (`mapping_id`);
-- Create index "subscribers_token_key" to table: "subscribers"
CREATE UNIQUE INDEX `subscribers_token_key` ON `subscribers` (`token`);
-- Create index "subscriber_email_owner_id" to table: "subscribers"
CREATE UNIQUE INDEX `subscriber_email_owner_id` ON `subscribers` (`email`, `owner_id`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
