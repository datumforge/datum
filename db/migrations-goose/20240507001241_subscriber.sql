-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_subscribers" table
CREATE TABLE `new_subscribers` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `phone_number` text NULL, `verified_email` bool NOT NULL DEFAULT (false), `verified_phone` bool NOT NULL DEFAULT (false), `active` bool NOT NULL DEFAULT (false), `token` text NOT NULL, `ttl` datetime NOT NULL, `secret` blob NOT NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `subscribers_organizations_subscribers` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- copy rows from old table "subscribers" to new temporary table "new_subscribers"
INSERT INTO `new_subscribers` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `email`, `phone_number`, `verified_email`, `verified_phone`, `active`, `token`, `ttl`, `secret`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `email`, `phone_number`, `verified_email`, `verified_phone`, `active`, `token`, `ttl`, `secret`, `owner_id` FROM `subscribers`;
-- drop "subscribers" table after copying rows
DROP TABLE `subscribers`;
-- rename temporary table "new_subscribers" to "subscribers"
ALTER TABLE `new_subscribers` RENAME TO `subscribers`;
-- create index "subscribers_mapping_id_key" to table: "subscribers"
CREATE UNIQUE INDEX `subscribers_mapping_id_key` ON `subscribers` (`mapping_id`);
-- create index "subscribers_token_key" to table: "subscribers"
CREATE UNIQUE INDEX `subscribers_token_key` ON `subscribers` (`token`);
-- create index "subscriber_email_owner_id" to table: "subscribers"
CREATE UNIQUE INDEX `subscriber_email_owner_id` ON `subscribers` (`email`, `owner_id`) WHERE deleted_at is NULL;
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "subscriber_email_owner_id" to table: "subscribers"
DROP INDEX `subscriber_email_owner_id`;
-- reverse: create index "subscribers_token_key" to table: "subscribers"
DROP INDEX `subscribers_token_key`;
-- reverse: create index "subscribers_mapping_id_key" to table: "subscribers"
DROP INDEX `subscribers_mapping_id_key`;
-- reverse: create "new_subscribers" table
DROP TABLE `new_subscribers`;
