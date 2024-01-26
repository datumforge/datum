-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_invites" table
CREATE TABLE `new_invites` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `expires` datetime NOT NULL, `recipient` text NOT NULL, `status` text NOT NULL DEFAULT ('INVITATION_SENT'), `send_attempts` integer NOT NULL DEFAULT (0), `requestor_id` text NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `invites_organizations_invites` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "invites" to new temporary table "new_invites"
INSERT INTO `new_invites` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `expires`, `recipient`, `status`, `requestor_id`, `secret`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `token`, `expires`, `recipient`, `status`, `requestor_id`, `secret`, `owner_id` FROM `invites`;
-- Drop "invites" table after copying rows
DROP TABLE `invites`;
-- Rename temporary table "new_invites" to "invites"
ALTER TABLE `new_invites` RENAME TO `invites`;
-- Create index "invites_token_key" to table: "invites"
CREATE UNIQUE INDEX `invites_token_key` ON `invites` (`token`);
-- Create index "invite_recipient_owner_id" to table: "invites"
CREATE UNIQUE INDEX `invite_recipient_owner_id` ON `invites` (`recipient`, `owner_id`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
