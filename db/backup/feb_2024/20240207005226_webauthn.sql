-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_webauthns" table
CREATE TABLE `new_webauthns` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `user_id` text NOT NULL, `credential_id` text NULL, `public_key` blob NULL, `attestation_type` text NULL, `aaguid` text NULL, `sign_count` integer NULL, `transports` json NULL, `flags` json NULL, `authenticator` json NULL, `backup_eligible` bool NULL, `backup_state` bool NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `webauthns_users_webauthn` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "webauthns" to new temporary table "new_webauthns"
INSERT INTO `new_webauthns` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `user_id`, `credential_id`, `public_key`, `attestation_type`, `aaguid`, `sign_count`, `transports`, `flags`, `authenticator`, `backup_eligible`, `backup_state`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `user_id`, `credential_id`, `public_key`, `attestation_type`, `aaguid`, `sign_count`, `transports`, `flags`, `authenticator`, `backup_eligible`, `backup_state`, `owner_id` FROM `webauthns`;
-- Drop "webauthns" table after copying rows
DROP TABLE `webauthns`;
-- Rename temporary table "new_webauthns" to "webauthns"
ALTER TABLE `new_webauthns` RENAME TO `webauthns`;
-- Create index "webauthns_user_id_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_user_id_key` ON `webauthns` (`user_id`);
-- Create index "webauthns_credential_id_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_credential_id_key` ON `webauthns` (`credential_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
