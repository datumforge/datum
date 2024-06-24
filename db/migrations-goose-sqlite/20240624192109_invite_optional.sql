-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_invites" table
CREATE TABLE `new_invites` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `token` text NOT NULL, `expires` datetime NULL, `recipient` text NOT NULL, `status` text NOT NULL DEFAULT ('INVITATION_SENT'), `role` text NOT NULL DEFAULT ('MEMBER'), `send_attempts` integer NOT NULL DEFAULT (0), `requestor_id` text NULL, `secret` blob NOT NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `invites_organizations_invites` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- copy rows from old table "invites" to new temporary table "new_invites"
INSERT INTO `new_invites` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `token`, `expires`, `recipient`, `status`, `role`, `send_attempts`, `requestor_id`, `secret`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `token`, `expires`, `recipient`, `status`, `role`, `send_attempts`, `requestor_id`, `secret`, `owner_id` FROM `invites`;
-- drop "invites" table after copying rows
DROP TABLE `invites`;
-- rename temporary table "new_invites" to "invites"
ALTER TABLE `new_invites` RENAME TO `invites`;
-- create index "invites_mapping_id_key" to table: "invites"
CREATE UNIQUE INDEX `invites_mapping_id_key` ON `invites` (`mapping_id`);
-- create index "invites_token_key" to table: "invites"
CREATE UNIQUE INDEX `invites_token_key` ON `invites` (`token`);
-- create index "invite_recipient_owner_id" to table: "invites"
CREATE UNIQUE INDEX `invite_recipient_owner_id` ON `invites` (`recipient`, `owner_id`) WHERE deleted_at is NULL;
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "invite_recipient_owner_id" to table: "invites"
DROP INDEX `invite_recipient_owner_id`;
-- reverse: create index "invites_token_key" to table: "invites"
DROP INDEX `invites_token_key`;
-- reverse: create index "invites_mapping_id_key" to table: "invites"
DROP INDEX `invites_mapping_id_key`;
-- reverse: create "new_invites" table
DROP TABLE `new_invites`;
