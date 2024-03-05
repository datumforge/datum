-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_webauthns" table
CREATE TABLE `new_webauthns` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `credential_id` blob NULL, `public_key` blob NULL, `attestation_type` text NULL, `aaguid` blob NOT NULL, `sign_count` integer NOT NULL, `transports` json NOT NULL, `backup_eligible` bool NOT NULL DEFAULT (false), `backup_state` bool NOT NULL DEFAULT (false), `user_present` bool NOT NULL DEFAULT (false), `user_verified` bool NOT NULL DEFAULT (false), `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `webauthns_users_webauthn` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "webauthns" to new temporary table "new_webauthns"
INSERT INTO `new_webauthns` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `credential_id`, `public_key`, `attestation_type`, `aaguid`, `sign_count`, `transports`, `backup_eligible`, `backup_state`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `credential_id`, `public_key`, `attestation_type`, `aaguid`, `sign_count`, `transports`, IFNULL(`backup_eligible`, (false)) AS `backup_eligible`, IFNULL(`backup_state`, (false)) AS `backup_state`, `owner_id` FROM `webauthns`;
-- Drop "webauthns" table after copying rows
DROP TABLE `webauthns`;
-- Rename temporary table "new_webauthns" to "webauthns"
ALTER TABLE `new_webauthns` RENAME TO `webauthns`;
-- Create index "webauthns_credential_id_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_credential_id_key` ON `webauthns` (`credential_id`);
-- Create index "webauthns_aaguid_key" to table: "webauthns"
CREATE UNIQUE INDEX `webauthns_aaguid_key` ON `webauthns` (`aaguid`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
