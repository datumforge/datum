-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_oauth_providers" table
CREATE TABLE `new_oauth_providers` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `name` text NOT NULL, `client_id` text NOT NULL, `client_secret` text NOT NULL, `redirect_url` text NOT NULL, `scopes` text NOT NULL, `auth_url` text NOT NULL, `token_url` text NOT NULL, `auth_style` integer NOT NULL, `info_url` text NOT NULL, `organization_oauthprovider` text NULL, PRIMARY KEY (`id`), CONSTRAINT `oauth_providers_organizations_oauthprovider` FOREIGN KEY (`organization_oauthprovider`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "oauth_providers" to new temporary table "new_oauth_providers"
INSERT INTO `new_oauth_providers` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `name`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url` FROM `oauth_providers`;
-- Drop "oauth_providers" table after copying rows
DROP TABLE `oauth_providers`;
-- Rename temporary table "new_oauth_providers" to "oauth_providers"
ALTER TABLE `new_oauth_providers` RENAME TO `oauth_providers`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
