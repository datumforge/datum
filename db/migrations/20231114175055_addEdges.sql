-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_refresh_tokens" table
CREATE TABLE `new_refresh_tokens` (`id` text NOT NULL, `client_id` text NOT NULL, `scopes` json NULL, `nonce` text NOT NULL, `claims_user_id` text NOT NULL, `claims_username` text NOT NULL, `claims_email` text NOT NULL, `claims_email_verified` bool NOT NULL, `claims_groups` json NULL, `claims_preferred_username` text NOT NULL, `connector_id` text NOT NULL, `connector_data` json NULL, `token` text NOT NULL, `obsolete_token` text NOT NULL, `last_used` datetime NOT NULL, `user_refreshtoken` text NULL, PRIMARY KEY (`id`), CONSTRAINT `refresh_tokens_users_refreshtoken` FOREIGN KEY (`user_refreshtoken`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "refresh_tokens" to new temporary table "new_refresh_tokens"
INSERT INTO `new_refresh_tokens` (`id`, `client_id`, `scopes`, `nonce`, `claims_user_id`, `claims_username`, `claims_email`, `claims_email_verified`, `claims_groups`, `claims_preferred_username`, `connector_id`, `connector_data`, `token`, `obsolete_token`, `last_used`) SELECT `id`, `client_id`, `scopes`, `nonce`, `claims_user_id`, `claims_username`, `claims_email`, `claims_email_verified`, `claims_groups`, `claims_preferred_username`, `connector_id`, `connector_data`, `token`, `obsolete_token`, `last_used` FROM `refresh_tokens`;
-- Drop "refresh_tokens" table after copying rows
DROP TABLE `refresh_tokens`;
-- Rename temporary table "new_refresh_tokens" to "refresh_tokens"
ALTER TABLE `new_refresh_tokens` RENAME TO `refresh_tokens`;
-- Create "new_entitlements" table
CREATE TABLE `new_entitlements` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `tier` text NOT NULL DEFAULT ('free'), `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires_at` datetime NULL, `upgraded_at` datetime NULL, `upgraded_tier` text NULL, `downgraded_at` datetime NULL, `downgraded_tier` text NULL, `cancelled` bool NOT NULL DEFAULT (false), `organization_entitlements` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`organization_entitlements`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "entitlements" to new temporary table "new_entitlements"
INSERT INTO `new_entitlements` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `tier`, `expires_at`, `cancelled`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `tier`, `expires_at`, `cancelled` FROM `entitlements`;
-- Drop "entitlements" table after copying rows
DROP TABLE `entitlements`;
-- Rename temporary table "new_entitlements" to "entitlements"
ALTER TABLE `new_entitlements` RENAME TO `entitlements`;
-- Create "new_oauth_providers" table
CREATE TABLE `new_oauth_providers` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, `user_oauthprovider` text NULL, PRIMARY KEY (`id`), CONSTRAINT `oauth_providers_users_oauthprovider` FOREIGN KEY (`user_oauthprovider`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "oauth_providers" to new temporary table "new_oauth_providers"
INSERT INTO `new_oauth_providers` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url` FROM `oauth_providers`;
-- Drop "oauth_providers" table after copying rows
DROP TABLE `oauth_providers`;
-- Rename temporary table "new_oauth_providers" to "oauth_providers"
ALTER TABLE `new_oauth_providers` RENAME TO `oauth_providers`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
