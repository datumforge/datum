-- Create "email_verification_tokens" table
CREATE TABLE `email_verification_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `email_verification_tokens_users_email_verification_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "email_verification_tokens_token_key" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `email_verification_tokens_token_key` ON `email_verification_tokens` (`token`);
-- Create index "emailverificationtoken_token" to table: "email_verification_tokens"
CREATE UNIQUE INDEX `emailverificationtoken_token` ON `email_verification_tokens` (`token`) WHERE deleted_at is NULL;
-- Create "entitlements" table
CREATE TABLE `entitlements` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tier` text NOT NULL DEFAULT ('FREE'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create "groups" table
CREATE TABLE `groups` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `gravatar_logo_url` text NULL, `logo_url` text NULL, `display_name` text NOT NULL DEFAULT (''), `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `groups_organizations_groups` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "group_name_owner_id" to table: "groups"
CREATE UNIQUE INDEX `group_name_owner_id` ON `groups` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- Create "group_memberships" table
CREATE TABLE `group_memberships` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `group_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `group_memberships_groups_group` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE NO ACTION, CONSTRAINT `group_memberships_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "groupmembership_user_id_group_id" to table: "group_memberships"
CREATE UNIQUE INDEX `groupmembership_user_id_group_id` ON `group_memberships` (`user_id`, `group_id`) WHERE deleted_at is NULL;
-- Create "group_settings" table
CREATE TABLE `group_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `visibility` text NOT NULL DEFAULT ('PUBLIC'), `join_policy` text NOT NULL DEFAULT ('INVITE_OR_APPLICATION'), `tags` json NOT NULL, `sync_to_slack` bool NOT NULL DEFAULT (false), `sync_to_github` bool NOT NULL DEFAULT (false), `group_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `group_settings_groups_setting` FOREIGN KEY (`group_setting`) REFERENCES `groups` (`id`) ON DELETE SET NULL);
-- Create index "group_settings_group_setting_key" to table: "group_settings"
CREATE UNIQUE INDEX `group_settings_group_setting_key` ON `group_settings` (`group_setting`);
-- Create "integrations" table
CREATE TABLE `integrations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `kind` text NULL, `secret_name` text NOT NULL, `organization_integrations` text NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`organization_integrations`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create "invites" table
CREATE TABLE `invites` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `expires` datetime NOT NULL, `recipient` text NOT NULL, `status` text NOT NULL DEFAULT ('INVITATION_SENT'), `role` text NOT NULL DEFAULT ('MEMBER'), `send_attempts` integer NOT NULL DEFAULT (0), `requestor_id` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `invites_organizations_invites` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Create index "invites_token_key" to table: "invites"
CREATE UNIQUE INDEX `invites_token_key` ON `invites` (`token`);
-- Create index "invite_recipient_owner_id" to table: "invites"
CREATE UNIQUE INDEX `invite_recipient_owner_id` ON `invites` (`recipient`, `owner_id`) WHERE deleted_at is NULL;
-- Create "oauth_providers" table
CREATE TABLE `oauth_providers` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, `organization_oauthprovider` text NULL, PRIMARY KEY (`id`), CONSTRAINT `oauth_providers_organizations_oauthprovider` FOREIGN KEY (`organization_oauthprovider`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create "oh_auth_too_tokens" table
CREATE TABLE `oh_auth_too_tokens` (`id` text NOT NULL, `client_id` text NOT NULL, `scopes` json NULL, `nonce` text NOT NULL, `claims_user_id` text NOT NULL, `claims_username` text NOT NULL, `claims_email` text NOT NULL, `claims_email_verified` bool NOT NULL, `claims_groups` json NULL, `claims_preferred_username` text NOT NULL, `connector_id` text NOT NULL, `connector_data` json NULL, `last_used` datetime NOT NULL, PRIMARY KEY (`id`));
-- Create "org_memberships" table
CREATE TABLE `org_memberships` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `role` text NOT NULL DEFAULT ('MEMBER'), `organization_id` text NOT NULL, `user_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `org_memberships_organizations_organization` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION, CONSTRAINT `org_memberships_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "orgmembership_user_id_organization_id" to table: "org_memberships"
CREATE UNIQUE INDEX `orgmembership_user_id_organization_id` ON `org_memberships` (`user_id`, `organization_id`) WHERE deleted_at is NULL;
-- Create "organizations" table
CREATE TABLE `organizations` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `personal_org` bool NOT NULL DEFAULT (false), `parent_organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organizations_organizations_children` FOREIGN KEY (`parent_organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "organization_name" to table: "organizations"
CREATE UNIQUE INDEX `organization_name` ON `organizations` (`name`) WHERE deleted_at is NULL;
-- Create "organization_settings" table
CREATE TABLE `organization_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `sso_cert` text NULL, `sso_entrypoint` text NULL, `sso_issuer` text NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `organization_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_setting`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "organization_settings_organization_setting_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_setting_key` ON `organization_settings` (`organization_setting`);
-- Create "password_reset_tokens" table
CREATE TABLE `password_reset_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `ttl` datetime NOT NULL, `email` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `password_reset_tokens_users_password_reset_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "password_reset_tokens_token_key" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `password_reset_tokens_token_key` ON `password_reset_tokens` (`token`);
-- Create index "passwordresettoken_token" to table: "password_reset_tokens"
CREATE UNIQUE INDEX `passwordresettoken_token` ON `password_reset_tokens` (`token`) WHERE deleted_at is NULL;
-- Create "personal_access_tokens" table
CREATE TABLE `personal_access_tokens` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `token` text NOT NULL, `abilities` json NULL, `expires_at` datetime NOT NULL, `description` text NULL DEFAULT (''), `last_used_at` datetime NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `personal_access_tokens_users_personal_access_tokens` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "personal_access_tokens_token_key" to table: "personal_access_tokens"
CREATE UNIQUE INDEX `personal_access_tokens_token_key` ON `personal_access_tokens` (`token`);
-- Create index "personalaccesstoken_token" to table: "personal_access_tokens"
CREATE INDEX `personalaccesstoken_token` ON `personal_access_tokens` (`token`);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `oauth` bool NOT NULL DEFAULT (false), `auth_provider` text NOT NULL DEFAULT ('CREDENTIALS'), `tfa_secret` text NULL, `is_phone_otp_allowed` bool NULL DEFAULT (true), `is_email_otp_allowed` bool NULL DEFAULT (true), `is_totp_allowed` bool NULL DEFAULT (true), `is_webauthn_allowed` bool NULL DEFAULT (true), `is_tfa_enabled` bool NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "users_sub_key" to table: "users"
CREATE UNIQUE INDEX `users_sub_key` ON `users` (`sub`);
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- Create index "user_email_auth_provider" to table: "users"
CREATE UNIQUE INDEX `user_email_auth_provider` ON `users` (`email`, `auth_provider`) WHERE deleted_at is NULL;
-- Create "user_settings" table
CREATE TABLE `user_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `default_org` text NULL, `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NOT NULL, `user_setting` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_setting`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "user_settings_user_setting_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_setting_key` ON `user_settings` (`user_setting`);
-- Create "webauthns" table
CREATE TABLE `webauthns` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `user_id` text NOT NULL, `credential_id` text NULL, `public_key` blob NULL, `attestation_type` text NULL, `aaguid` text NULL, `sign_count` integer NULL, `transports` json NULL, `flags` json NULL, `authenticator` json NULL, `backup_eligible` bool NULL, `backup_state` bool NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `webauthns_users_webauthn` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "webauthns_user_id_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_user_id_key` ON `webauthns` (`user_id`);
-- Create index "webauthns_credential_id_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_credential_id_key` ON `webauthns` (`credential_id`);
