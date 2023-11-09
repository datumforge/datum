-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_refresh_tokens" table
CREATE TABLE `new_refresh_tokens` (`id` text NOT NULL, `client_id` text NOT NULL, `scopes` json NULL, `nonce` text NOT NULL, `claims_user_id` text NOT NULL, `claims_username` text NOT NULL, `claims_email` text NOT NULL, `claims_email_verified` bool NOT NULL, `claims_groups` json NULL, `claims_preferred_username` text NOT NULL, `connector_id` text NOT NULL, `connector_data` json NULL, `token` text NOT NULL, `obsolete_token` text NOT NULL, `last_used` datetime NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "refresh_tokens" to new temporary table "new_refresh_tokens"
INSERT INTO `new_refresh_tokens` (`id`, `client_id`, `scopes`, `nonce`, `claims_user_id`, `claims_username`, `claims_email`, `claims_email_verified`, `claims_groups`, `claims_preferred_username`, `connector_id`, `connector_data`, `token`, `obsolete_token`, `last_used`) SELECT `id`, `client_id`, `scopes`, `nonce`, `claims_user_id`, `claims_username`, `claims_email`, `claims_email_verified`, `claims_groups`, `claims_preferred_username`, `connector_id`, `connector_data`, `token`, `obsolete_token`, `last_used` FROM `refresh_tokens`;
-- Drop "refresh_tokens" table after copying rows
DROP TABLE `refresh_tokens`;
-- Rename temporary table "new_refresh_tokens" to "refresh_tokens"
ALTER TABLE `new_refresh_tokens` RENAME TO `refresh_tokens`;
-- Create "organization_settings" table
CREATE TABLE `organization_settings` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `domains` json NOT NULL, `sso_cert` text NOT NULL DEFAULT (''), `sso_entrypoint` text NOT NULL DEFAULT (''), `sso_issuer` text NOT NULL DEFAULT (''), `billing_contact` text NOT NULL, `billing_email` text NOT NULL, `billing_phone` text NOT NULL, `billing_address` text NOT NULL, `tax_identifier` text NOT NULL, PRIMARY KEY (`id`));
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
