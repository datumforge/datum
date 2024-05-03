-- Create "api_tokens" table
CREATE TABLE `api_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `api_tokens_organizations_api_tokens` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "api_tokens_mapping_id_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_mapping_id_key` ON `api_tokens` (`mapping_id`);
-- Create index "api_tokens_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_token_key` ON `api_tokens` (`token`);
-- Create index "apitoken_token" to table: "api_tokens"
CREATE INDEX `apitoken_token` ON `api_tokens` (`token`);
-- Create "document_data" table
CREATE TABLE `document_data` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `data` json NOT NULL, `template_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `document_data_templates_documents` FOREIGN KEY (`template_id`) REFERENCES `templates` (`id`) ON DELETE NO ACTION);
-- Create index "document_data_mapping_id_key" to table: "document_data"
CREATE UNIQUE INDEX `document_data_mapping_id_key` ON `document_data` (`mapping_id`);
-- Create "document_data_history" table
CREATE TABLE `document_data_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `template_id` text NOT NULL, `data` json NOT NULL, PRIMARY KEY (`id`));
-- Create index "documentdatahistory_history_time" to table: "document_data_history"
CREATE INDEX `documentdatahistory_history_time` ON `document_data_history` (`history_time`);
-- Create "email_verification_tokens" table
CREATE TABLE `email_verification_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `email_verification_tokens_users_email_verification_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "email_verification_tokens_mapping_id_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_mapping_id_key` ON `email_verification_tokens` (`mapping_id`);
-- Create index "email_verification_tokens_token_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_token_key` ON `email_verification_tokens` (`token`);
-- Create index "emailverificationtoken_token" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `emailverificationtoken_token` ON `email_verification_tokens` (`token`) WHERE deleted_at is NULL;
-- Create "entitlements" table
CREATE TABLE `entitlements` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "entitlements_mapping_id_key" to table: "entitlements"
CREATE UNIQUE INDEX `entitlements_mapping_id_key` ON `entitlements` (`mapping_id`);
-- Create "entitlement_history" table
CREATE TABLE `entitlement_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "entitlementhistory_history_time" to table: "entitlement_history"
CREATE INDEX `entitlementhistory_history_time` ON `entitlement_history` (`history_time`);
-- Create "events" table
CREATE TABLE `events` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `event_id` text NOT NULL, `correlation_id` text NOT NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- Create index "events_mapping_id_key" to table: "events"
CREATE UNIQUE INDEX `events_mapping_id_key` ON `events` (`mapping_id`);
-- Create "event_history" table
CREATE TABLE `event_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `event_id` text NOT NULL, `correlation_id` text NOT NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- Create index "eventhistory_history_time" to table: "event_history"
CREATE INDEX `eventhistory_history_time` ON `event_history` (`history_time`);
-- Create "features" table
CREATE TABLE `features` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `global` bool NOT NULL DEFAULT (true), `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, PRIMARY KEY (`id`));
-- Create index "features_mapping_id_key" to table: "features"
CREATE UNIQUE INDEX `features_mapping_id_key` ON `features` (`mapping_id`);
-- Create index "features_name_key" to table: "features"
CREATE UNIQUE INDEX `features_name_key` ON `features` (`name`);
-- Create "feature_history" table
CREATE TABLE `feature_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `global` bool NOT NULL DEFAULT (true), `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, PRIMARY KEY (`id`));
-- Create index "featurehistory_history_time" to table: "feature_history"
CREATE INDEX `featurehistory_history_time` ON `feature_history` (`history_time`);
-- Create "files" table
CREATE TABLE `files` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `file_name` text NOT NULL, `file_extension` text NOT NULL, `file_size` integer NULL, `content_type` text NOT NULL, `store_key` text NOT NULL, `category` text NULL, `annotation` text NULL, `user_files` text NULL, PRIMARY KEY (`id`), CONSTRAINT `files_users_files` FOREIGN KEY (`user_files`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "files_mapping_id_key" to table: "files"
CREATE UNIQUE INDEX `files_mapping_id_key` ON `files` (`mapping_id`);
-- Create "file_history" table
CREATE TABLE `file_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `file_name` text NOT NULL, `file_extension` text NOT NULL, `file_size` integer NULL, `content_type` text NOT NULL, `store_key` text NOT NULL, `category` text NULL, `annotation` text NULL, PRIMARY KEY (`id`));
-- Create index "filehistory_history_time" to table: "file_history"
CREATE INDEX `filehistory_history_time` ON `file_history` (`history_time`);
-- Create "groups" table
CREATE TABLE `groups` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "groups_mapping_id_key" to table: "groups"
CREATE UNIQUE INDEX `groups_mapping_id_key` ON `groups` (`mapping_id`);
-- Create index "group_name_owner_id" to table: "groups"
CREATE UNIQUE INDEX `group_name_owner_id` ON `groups` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- Create "group_history" table
CREATE TABLE `group_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), PRIMARY KEY (`id`));
-- Create index "grouphistory_history_time" to table: "group_history"
CREATE INDEX `grouphistory_history_time` ON `group_history` (`history_time`);
-- Create "group_memberships" table
CREATE TABLE `group_memberships` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `group_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `group_memberships_groups_group` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE NO ACTION, CONSTRAINT `group_memberships_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "group_memberships_mapping_id_key" to table: "group_memberships"
CREATE UNIQUE INDEX `group_memberships_mapping_id_key` ON `group_memberships` (`mapping_id`);
-- Create index "groupmembership_user_id_group_id" to table: "group_memberships"
CREATE UNIQUE INDEX `groupmembership_user_id_group_id` ON `group_memberships` (`user_id`, `group_id`) WHERE deleted_at is NULL;
-- Create "group_membership_history" table
CREATE TABLE `group_membership_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `group_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "groupmembershiphistory_history_time" to table: "group_membership_history"
CREATE INDEX `groupmembershiphistory_history_time` ON `group_membership_history` (`history_time`);
-- Create "group_settings" table
CREATE TABLE `group_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NULL, `sync_to_slack` bool NULL DEFAULT (false), `sync_to_github` bool NULL DEFAULT (false), `group_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Create index "group_settings_mapping_id_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_mapping_id_key` ON `group_settings` (`mapping_id`);
-- Create index "group_settings_group_id_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_id_key` ON `group_settings` (`group_id`);
-- Create "group_setting_history" table
CREATE TABLE `group_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NULL, `sync_to_slack` bool NULL DEFAULT (false), `sync_to_github` bool NULL DEFAULT (false), `group_id` text NULL, PRIMARY KEY (`id`));
-- Create index "groupsettinghistory_history_time" to table: "group_setting_history"
CREATE INDEX `groupsettinghistory_history_time` ON `group_setting_history` (`history_time`);
-- Create "hushes" table
CREATE TABLE `hushes` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NULL, `secret_value` text NULL, PRIMARY KEY (`id`));
-- Create index "hushes_mapping_id_key" to table: "hushes"
CREATE UNIQUE INDEX `hushes_mapping_id_key` ON `hushes` (`mapping_id`);
-- Create "hush_history" table
CREATE TABLE `hush_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NULL, `secret_value` text NULL, PRIMARY KEY (`id`));
-- Create index "hushhistory_history_time" to table: "hush_history"
CREATE INDEX `hushhistory_history_time` ON `hush_history` (`history_time`);
-- Create "integrations" table
CREATE TABLE `integrations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `group_integrations` text NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_groups_integrations` FOREIGN KEY (`group_integrations`) REFERENCES `groups` (`id`) ON DELETE SET NULL, CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "integrations_mapping_id_key" to table: "integrations"
CREATE UNIQUE INDEX `integrations_mapping_id_key` ON `integrations` (`mapping_id`);
-- Create "integration_history" table
CREATE TABLE `integration_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, PRIMARY KEY (`id`));
-- Create index "integrationhistory_history_time" to table: "integration_history"
CREATE INDEX `integrationhistory_history_time` ON `integration_history` (`history_time`);
-- Create "invites" table
CREATE TABLE `invites` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `expires` datetime NOT NULL, `recipient` text NOT NULL, `status` text NOT NULL DEFAULT ('INVITATION_SENT'), `role` text NOT NULL DEFAULT ('MEMBER'), `send_attempts` integer NOT NULL DEFAULT (0), `requestor_id` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `invites_organizations_invites` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "invites_mapping_id_key" to table: "invites"
CREATE UNIQUE INDEX `invites_mapping_id_key` ON `invites` (`mapping_id`);
-- Create index "invites_token_key" to table: "invites"
CREATE UNIQUE INDEX `invites_token_key` ON `invites` (`token`);
-- Create index "invite_recipient_owner_id" to table: "invites"
CREATE UNIQUE INDEX `invite_recipient_owner_id` ON `invites` (`recipient`, `owner_id`) WHERE deleted_at is NULL;
-- Create "oauth_providers" table
CREATE TABLE `oauth_providers` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, `organization_oauthprovider` text NULL, PRIMARY KEY (`id`), CONSTRAINT `oauth_providers_organizations_oauthprovider` FOREIGN KEY (`organization_oauthprovider`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "oauth_providers_mapping_id_key" to table: "oauth_providers"
CREATE UNIQUE INDEX `oauth_providers_mapping_id_key` ON `oauth_providers` (`mapping_id`);
-- Create "oauth_provider_history" table
CREATE TABLE `oauth_provider_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "oauthproviderhistory_history_time" to table: "oauth_provider_history"
CREATE INDEX `oauthproviderhistory_history_time` ON `oauth_provider_history` (`history_time`);
-- Create "oh_auth_too_tokens" table
CREATE TABLE `oh_auth_too_tokens` (`id` text NOT NULL, `mapping_id` text NOT NULL, `client_id` text NOT NULL, `scopes` json NULL, `nonce` text NOT NULL, `claims_user_id` text NOT NULL, `claims_username` text NOT NULL, `claims_email` text NOT NULL, `claims_email_verified` bool NOT NULL, `claims_groups` json NULL, `claims_preferred_username` text NOT NULL, `connector_id` text NOT NULL, `connector_data` json NULL, `last_used` datetime NOT NULL, PRIMARY KEY (`id`));
-- Create index "oh_auth_too_tokens_mapping_id_key" to table: "oh_auth_too_tokens"
CREATE UNIQUE INDEX `oh_auth_too_tokens_mapping_id_key` ON `oh_auth_too_tokens` (`mapping_id`);
-- Create "org_memberships" table
CREATE TABLE `org_memberships` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `org_memberships_organizations_organization` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION, CONSTRAINT `org_memberships_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "org_memberships_mapping_id_key" to table: "org_memberships"
CREATE UNIQUE INDEX `org_memberships_mapping_id_key` ON `org_memberships` (`mapping_id`);
-- Create index "orgmembership_user_id_organization_id" to table: "org_memberships"
CREATE UNIQUE INDEX `orgmembership_user_id_organization_id` ON `org_memberships` (`user_id`, `organization_id`) WHERE deleted_at is NULL;
-- Create "org_membership_history" table
CREATE TABLE `org_membership_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "orgmembershiphistory_history_time" to table: "org_membership_history"
CREATE INDEX `orgmembershiphistory_history_time` ON `org_membership_history` (`history_time`);
-- Create "organizations" table
CREATE TABLE `organizations` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `personal_org` bool NULL DEFAULT (false), `avatar_remote_url` text NULL, `dedicated_db` bool NOT NULL DEFAULT (false), `parent_organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "organizations_mapping_id_key" to table: "organizations"
CREATE UNIQUE INDEX `organizations_mapping_id_key` ON `organizations` (`mapping_id`);
-- Create index "organization_name" to table: "organizations"
CREATE UNIQUE INDEX `organization_name` ON `organizations` (`name`) WHERE deleted_at is NULL;
-- Create "organization_history" table
CREATE TABLE `organization_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `parent_organization_id` text NULL, `personal_org` bool NULL DEFAULT (false), `avatar_remote_url` text NULL, `dedicated_db` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "organizationhistory_history_time" to table: "organization_history"
CREATE INDEX `organizationhistory_history_time` ON `organization_history` (`history_time`);
-- Create "organization_settings" table
CREATE TABLE `organization_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `geo_location` text NULL DEFAULT ('AMER'), `organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "organization_settings_mapping_id_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_mapping_id_key` ON `organization_settings` (`mapping_id`);
-- Create index "organization_settings_organization_id_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_id_key` ON `organization_settings` (`organization_id`);
-- Create "organization_setting_history" table
CREATE TABLE `organization_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `geo_location` text NULL DEFAULT ('AMER'), `organization_id` text NULL, PRIMARY KEY (`id`));
-- Create index "organizationsettinghistory_history_time" to table: "organization_setting_history"
CREATE INDEX `organizationsettinghistory_history_time` ON `organization_setting_history` (`history_time`);
-- Create "password_reset_tokens" table
CREATE TABLE `password_reset_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `password_reset_tokens_users_password_reset_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "password_reset_tokens_mapping_id_key" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `password_reset_tokens_mapping_id_key` ON `password_reset_tokens` (`mapping_id`);
-- Create index "password_reset_tokens_token_key" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `password_reset_tokens_token_key` ON `password_reset_tokens` (`token`);
-- Create index "passwordresettoken_token" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `passwordresettoken_token` ON `password_reset_tokens` (`token`) WHERE deleted_at is NULL;
-- Create "personal_access_tokens" table
CREATE TABLE `personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `token` text NOT NULL, `expires_at` datetime NOT NULL, `description` text NULL, `scopes` json NULL, `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "personal_access_tokens_mapping_id_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_mapping_id_key` ON `personal_access_tokens` (`mapping_id`);
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Create "subscribers" table
CREATE TABLE `subscribers` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `phone_number` text NULL, `verified_email` bool NOT NULL DEFAULT (false), `verified_phone` bool NOT NULL DEFAULT (false), `active` bool NOT NULL DEFAULT (false), `token` text NOT NULL, `ttl` datetime NOT NULL, `secret` blob NOT NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `subscribers_organizations_subscribers` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "subscribers_mapping_id_key" to table: "subscribers"
CREATE UNIQUE INDEX `subscribers_mapping_id_key` ON `subscribers` (`mapping_id`);
-- Create index "subscribers_token_key" to table: "subscribers"
CREATE UNIQUE INDEX `subscribers_token_key` ON `subscribers` (`token`);
-- Create index "subscriber_email_owner_id" to table: "subscribers"
CREATE UNIQUE INDEX `subscriber_email_owner_id` ON `subscribers` (`email`, `owner_id`) WHERE deleted_at is NULL;
-- Create "tfa_settings" table
CREATE TABLE `tfa_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tfa_secret` text NULL, `verified` bool NOT NULL DEFAULT (false), `recovery_codes` json NULL, `phone_otp_allowed` bool NULL DEFAULT (false), `email_otp_allowed` bool NULL DEFAULT (false), `totp_allowed` bool NULL DEFAULT (false), `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `tfa_settings_users_tfa_settings` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "tfa_settings_mapping_id_key" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfa_settings_mapping_id_key` ON `tfa_settings` (`mapping_id`);
-- Create index "tfasetting_owner_id" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfasetting_owner_id` ON `tfa_settings` (`owner_id`) WHERE deleted_at is NULL;
-- Create "templates" table
CREATE TABLE `templates` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `name` text NOT NULL, `template_type` text NOT NULL DEFAULT ('DOCUMENT'), `description` text NULL, `jsonconfig` json NOT NULL, `uischema` json NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `templates_organizations_templates` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "templates_mapping_id_key" to table: "templates"
CREATE UNIQUE INDEX `templates_mapping_id_key` ON `templates` (`mapping_id`);
-- Create index "template_name_owner_id_template_type" to table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id_template_type` ON `templates` (`name`, `owner_id`, `template_type`) WHERE deleted_at is NULL;
-- Create "template_history" table
CREATE TABLE `template_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `template_type` text NOT NULL DEFAULT ('DOCUMENT'), `description` text NULL, `jsonconfig` json NOT NULL, `uischema` json NULL, PRIMARY KEY (`id`));
-- Create index "templatehistory_history_time" to table: "template_history"
CREATE INDEX `templatehistory_history_time` ON `template_history` (`history_time`);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `auth_provider` text NOT NULL DEFAULT ('CREDENTIALS'), `role` text NULL DEFAULT ('USER'), PRIMARY KEY (`id`));
-- Create index "users_mapping_id_key" to table: "users"
CREATE UNIQUE INDEX `users_mapping_id_key` ON `users` (`mapping_id`);
-- Create index "users_sub_key" to table: "users"
CREATE UNIQUE INDEX `users_sub_key` ON `users` (`sub`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create index "user_email_auth_provider" to table: "users"
CREATE UNIQUE INDEX `user_email_auth_provider` ON `users` (`email`, `auth_provider`) WHERE deleted_at is NULL;
-- Create "user_history" table
CREATE TABLE `user_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `auth_provider` text NOT NULL DEFAULT ('CREDENTIALS'), `role` text NULL DEFAULT ('USER'), PRIMARY KEY (`id`));
-- Create index "userhistory_history_time" to table: "user_history"
CREATE INDEX `userhistory_history_time` ON `user_history` (`history_time`);
-- Create "user_settings" table
CREATE TABLE `user_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NULL, `is_webauthn_allowed` bool NULL DEFAULT (false), `is_tfa_enabled` bool NULL DEFAULT (false), `phone_number` text NULL, `user_id` text NULL, `user_setting_default_org` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `user_settings_organizations_default_org` FOREIGN KEY (`user_setting_default_org`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "user_settings_mapping_id_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_mapping_id_key` ON `user_settings` (`mapping_id`);
-- Create index "user_settings_user_id_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_id_key` ON `user_settings` (`user_id`);
-- Create "user_setting_history" table
CREATE TABLE `user_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `user_id` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NULL, `is_webauthn_allowed` bool NULL DEFAULT (false), `is_tfa_enabled` bool NULL DEFAULT (false), `phone_number` text NULL, PRIMARY KEY (`id`));
-- Create index "usersettinghistory_history_time" to table: "user_setting_history"
CREATE INDEX `usersettinghistory_history_time` ON `user_setting_history` (`history_time`);
-- Create "webauthns" table
CREATE TABLE `webauthns` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `credential_id` blob NULL, `public_key` blob NULL, `attestation_type` text NULL, `aaguid` blob NOT NULL, `sign_count` integer NOT NULL, `transports` json NOT NULL, `backup_eligible` bool NOT NULL DEFAULT (false), `backup_state` bool NOT NULL DEFAULT (false), `user_present` bool NOT NULL DEFAULT (false), `user_verified` bool NOT NULL DEFAULT (false), `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `webauthns_users_webauthn` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "webauthns_mapping_id_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_mapping_id_key` ON `webauthns` (`mapping_id`);
-- Create index "webauthns_credential_id_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_credential_id_key` ON `webauthns` (`credential_id`);
-- Create index "webauthns_aaguid_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_aaguid_key` ON `webauthns` (`aaguid`);
-- Create "webhooks" table
CREATE TABLE `webhooks` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `destination_url` text NOT NULL, `enabled` bool NOT NULL DEFAULT (true), `callback` text NULL, `expires_at` datetime NULL, `secret` blob NULL, `failures` integer NULL DEFAULT (0), `last_error` text NULL, `last_response` text NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `webhooks_organizations_webhooks` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "webhooks_mapping_id_key" to table: "webhooks"
CREATE UNIQUE INDEX `webhooks_mapping_id_key` ON `webhooks` (`mapping_id`);
-- Create index "webhooks_callback_key" to table: "webhooks"
CREATE UNIQUE INDEX `webhooks_callback_key` ON `webhooks` (`callback`);
-- Create index "webhook_name_owner_id" to table: "webhooks"
CREATE UNIQUE INDEX `webhook_name_owner_id` ON `webhooks` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- Create "webhook_history" table
CREATE TABLE `webhook_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NULL, `name` text NOT NULL, `description` text NULL, `destination_url` text NOT NULL, `enabled` bool NOT NULL DEFAULT (true), `callback` text NULL, `expires_at` datetime NULL, `secret` blob NULL, `failures` integer NULL DEFAULT (0), `last_error` text NULL, `last_response` text NULL, PRIMARY KEY (`id`));
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
-- Create "integration_secrets" table
CREATE TABLE `integration_secrets` (`integration_id` text NOT NULL, `hush_id` text NOT NULL, PRIMARY KEY (`integration_id`, `hush_id`), CONSTRAINT `integration_secrets_integration_id` FOREIGN KEY (`integration_id`) REFERENCES `integrations` (`id`) ON DELETE CASCADE, CONSTRAINT `integration_secrets_hush_id` FOREIGN KEY (`hush_id`) REFERENCES `hushes` (`id`) ON DELETE CASCADE);
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
-- Create "organization_personal_access_tokens" table
CREATE TABLE `organization_personal_access_tokens` (`organization_id` text NOT NULL, `personal_access_token_id` text NOT NULL, PRIMARY KEY (`organization_id`, `personal_access_token_id`), CONSTRAINT `organization_personal_access_tokens_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_personal_access_tokens_personal_access_token_id` FOREIGN KEY (`personal_access_token_id`) REFERENCES `personal_access_tokens` (`id`) ON DELETE CASCADE);
-- Create "organization_events" table
CREATE TABLE `organization_events` (`organization_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`organization_id`, `event_id`), CONSTRAINT `organization_events_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- Create "organization_secrets" table
CREATE TABLE `organization_secrets` (`organization_id` text NOT NULL, `hush_id` text NOT NULL, PRIMARY KEY (`organization_id`, `hush_id`), CONSTRAINT `organization_secrets_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_secrets_hush_id` FOREIGN KEY (`hush_id`) REFERENCES `hushes` (`id`) ON DELETE CASCADE);
-- Create "organization_features" table
CREATE TABLE `organization_features` (`organization_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`organization_id`, `feature_id`), CONSTRAINT `organization_features_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `organization_features_feature_id` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE CASCADE);
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
