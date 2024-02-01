-- Create "entitlement_history" table
CREATE TABLE `entitlement_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "entitlementhistory_history_time" to table: "entitlement_history"
CREATE INDEX `entitlementhistory_history_time` ON `entitlement_history` (`history_time`);
-- Create "group_history" table
CREATE TABLE `group_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), PRIMARY KEY (`id`));
-- Create index "grouphistory_history_time" to table: "group_history"
CREATE INDEX `grouphistory_history_time` ON `group_history` (`history_time`);
-- Create "group_membership_history" table
CREATE TABLE `group_membership_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `group_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "groupmembershiphistory_history_time" to table: "group_membership_history"
CREATE INDEX `groupmembershiphistory_history_time` ON `group_membership_history` (`history_time`);
-- Create "group_setting_history" table
CREATE TABLE `group_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NOT NULL, `sync_to_slack` bool NOT NULL DEFAULT (false), `sync_to_github` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "groupsettinghistory_history_time" to table: "group_setting_history"
CREATE INDEX `groupsettinghistory_history_time` ON `group_setting_history` (`history_time`);
-- Create "integration_history" table
CREATE TABLE `integration_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "integrationhistory_history_time" to table: "integration_history"
CREATE INDEX `integrationhistory_history_time` ON `integration_history` (`history_time`);
-- Create "oauth_provider_history" table
CREATE TABLE `oauth_provider_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "oauthproviderhistory_history_time" to table: "oauth_provider_history"
CREATE INDEX `oauthproviderhistory_history_time` ON `oauth_provider_history` (`history_time`);
-- Create "org_membership_history" table
CREATE TABLE `org_membership_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "orgmembershiphistory_history_time" to table: "org_membership_history"
CREATE INDEX `orgmembershiphistory_history_time` ON `org_membership_history` (`history_time`);
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
