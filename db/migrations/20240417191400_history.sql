-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_templates" table
CREATE TABLE `new_templates` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `template_type` text NOT NULL DEFAULT ('DOCUMENT'), `description` text NULL, `jsonconfig` json NOT NULL, `uischema` json NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `templates_organizations_templates` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "templates" to new temporary table "new_templates"
INSERT INTO `new_templates` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `jsonconfig`, `uischema`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `jsonconfig`, `uischema`, `owner_id` FROM `templates`;
-- Drop "templates" table after copying rows
DROP TABLE `templates`;
-- Rename temporary table "new_templates" to "templates"
ALTER TABLE `new_templates` RENAME TO `templates`;
-- Create index "template_name_owner_id_template_type" to table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id_template_type` ON `templates` (`name`, `owner_id`, `template_type`) WHERE deleted_at is NULL;
-- Create "document_data_history" table
CREATE TABLE `document_data_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `template_id` text NOT NULL, `data` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "documentdatahistory_history_time" to table: "document_data_history"
CREATE INDEX `documentdatahistory_history_time` ON `document_data_history` (`history_time`);
-- Create "entitlement_history" table
CREATE TABLE `entitlement_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "entitlementhistory_history_time" to table: "entitlement_history"
CREATE INDEX `entitlementhistory_history_time` ON `entitlement_history` (`history_time`);
-- Create "group_history" table
CREATE TABLE `group_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), PRIMARY KEY (`id`));
-- Create index "grouphistory_history_time" to table: "group_history"
CREATE INDEX `grouphistory_history_time` ON `group_history` (`history_time`);
-- Create "group_membership_history" table
CREATE TABLE `group_membership_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `group_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "groupmembershiphistory_history_time" to table: "group_membership_history"
CREATE INDEX `groupmembershiphistory_history_time` ON `group_membership_history` (`history_time`);
-- Create "group_setting_history" table
CREATE TABLE `group_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NULL, `sync_to_slack` bool NULL DEFAULT (false), `sync_to_github` bool NULL DEFAULT (false), `group_id` text NULL, PRIMARY KEY (`id`));
-- Create index "groupsettinghistory_history_time" to table: "group_setting_history"
CREATE INDEX `groupsettinghistory_history_time` ON `group_setting_history` (`history_time`);
-- Create "integration_history" table
CREATE TABLE `integration_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NULL, PRIMARY KEY (`id`));
-- Create index "integrationhistory_history_time" to table: "integration_history"
CREATE INDEX `integrationhistory_history_time` ON `integration_history` (`history_time`);
-- Create "oauth_provider_history" table
CREATE TABLE `oauth_provider_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "oauthproviderhistory_history_time" to table: "oauth_provider_history"
CREATE INDEX `oauthproviderhistory_history_time` ON `oauth_provider_history` (`history_time`);
-- Create "org_membership_history" table
CREATE TABLE `org_membership_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "orgmembershiphistory_history_time" to table: "org_membership_history"
CREATE INDEX `orgmembershiphistory_history_time` ON `org_membership_history` (`history_time`);
-- Create "template_history" table
CREATE TABLE `template_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `template_type` text NOT NULL DEFAULT ('DOCUMENT'), `description` text NULL, `jsonconfig` json NOT NULL, `uischema` json NULL, PRIMARY KEY (`id`));
-- Create index "templatehistory_history_time" to table: "template_history"
CREATE INDEX `templatehistory_history_time` ON `template_history` (`history_time`);
-- Create "user_history" table
CREATE TABLE `user_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `auth_provider` text NOT NULL DEFAULT ('CREDENTIALS'), PRIMARY KEY (`id`));
-- Create index "userhistory_history_time" to table: "user_history"
CREATE INDEX `userhistory_history_time` ON `user_history` (`history_time`);
-- Create "user_setting_history" table
CREATE TABLE `user_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `user_id` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NOT NULL, `is_webauthn_allowed` bool NULL DEFAULT (false), `is_tfa_enabled` bool NULL DEFAULT (false), `phone_number` text NULL, PRIMARY KEY (`id`));
-- Create index "usersettinghistory_history_time" to table: "user_setting_history"
CREATE INDEX `usersettinghistory_history_time` ON `user_setting_history` (`history_time`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
