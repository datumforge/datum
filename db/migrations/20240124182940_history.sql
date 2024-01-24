-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_users" table
CREATE TABLE `new_users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `oauth` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `oauth`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `oauth` FROM `users`;
-- Drop "users" table after copying rows
DROP TABLE `users`;
-- Rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- Create index "users_sub_key" to table: "users"
CREATE UNIQUE INDEX `users_sub_key` ON `users` (`sub`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create index "user_email" to table: "users"
CREATE UNIQUE INDEX `user_email` ON `users` (`email`) WHERE deleted_at is NULL;
-- Create "new_entitlements" table
CREATE TABLE `new_entitlements` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "entitlements" to new temporary table "new_entitlements"
INSERT INTO `new_entitlements` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `tier`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, IFNULL(`tier`, ('FREE')) AS `tier`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`, `owner_id` FROM `entitlements`;
-- Drop "entitlements" table after copying rows
DROP TABLE `entitlements`;
-- Rename temporary table "new_entitlements" to "entitlements"
ALTER TABLE `new_entitlements` RENAME TO `entitlements`;
-- Create "group_history" table
CREATE TABLE `group_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), PRIMARY KEY (`id`));
-- Create index "grouphistory_history_time" to table: "group_history"
CREATE INDEX `grouphistory_history_time` ON `group_history` (`history_time`);
-- Create "group_setting_history" table
CREATE TABLE `group_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NOT NULL, `sync_to_slack` bool NOT NULL DEFAULT (false), `sync_to_github` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "groupsettinghistory_history_time" to table: "group_setting_history"
CREATE INDEX `groupsettinghistory_history_time` ON `group_setting_history` (`history_time`);
-- Create "organization_history" table
CREATE TABLE `organization_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `parent_organization_id` text NULL, `personal_org` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "organizationhistory_history_time" to table: "organization_history"
CREATE INDEX `organizationhistory_history_time` ON `organization_history` (`history_time`);
-- Create "organization_setting_history" table
CREATE TABLE `organization_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `sso_cert` text NULL, `sso_entrypoint` text NULL, `sso_issuer` text NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, PRIMARY KEY (`id`));
-- Create index "organizationsettinghistory_history_time" to table: "organization_setting_history"
CREATE INDEX `organizationsettinghistory_history_time` ON `organization_setting_history` (`history_time`);
-- Create "user_history" table
CREATE TABLE `user_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `oauth` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "userhistory_history_time" to table: "user_history"
CREATE INDEX `userhistory_history_time` ON `user_history` (`history_time`);
-- Create "user_setting_history" table
CREATE TABLE `user_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `default_org` text NULL, `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "usersettinghistory_history_time" to table: "user_setting_history"
CREATE INDEX `usersettinghistory_history_time` ON `user_setting_history` (`history_time`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
