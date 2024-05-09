-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_users" table
CREATE TABLE `new_users` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `email` text NOT NULL, `first_name` text NULL, `last_name` text NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `auth_provider` text NOT NULL DEFAULT ('CREDENTIALS'), `role` text NULL DEFAULT ('USER'), PRIMARY KEY (`id`));
-- copy rows from old table "users" to new temporary table "new_users"
INSERT INTO `new_users` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `auth_provider`, `role`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `auth_provider`, `role` FROM `users`;
-- drop "users" table after copying rows
DROP TABLE `users`;
-- rename temporary table "new_users" to "users"
ALTER TABLE `new_users` RENAME TO `users`;
-- create index "users_mapping_id_key" to table: "users"
CREATE UNIQUE INDEX `users_mapping_id_key` ON `users` (`mapping_id`);
-- create index "users_sub_key" to table: "users"
CREATE UNIQUE INDEX `users_sub_key` ON `users` (`sub`);
-- create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);
-- create index "user_email_auth_provider" to table: "users"
CREATE UNIQUE INDEX `user_email_auth_provider` ON `users` (`email`, `auth_provider`) WHERE deleted_at is NULL;
-- create "new_user_history" table
CREATE TABLE `new_user_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `email` text NOT NULL, `first_name` text NULL, `last_name` text NULL, `display_name` text NOT NULL, `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `last_seen` datetime NULL, `password` text NULL, `sub` text NULL, `auth_provider` text NOT NULL DEFAULT ('CREDENTIALS'), `role` text NULL DEFAULT ('USER'), PRIMARY KEY (`id`));
-- copy rows from old table "user_history" to new temporary table "new_user_history"
INSERT INTO `new_user_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `auth_provider`, `role`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `email`, `first_name`, `last_name`, `display_name`, `avatar_remote_url`, `avatar_local_file`, `avatar_updated_at`, `last_seen`, `password`, `sub`, `auth_provider`, `role` FROM `user_history`;
-- drop "user_history" table after copying rows
DROP TABLE `user_history`;
-- rename temporary table "new_user_history" to "user_history"
ALTER TABLE `new_user_history` RENAME TO `user_history`;
-- create index "userhistory_history_time" to table: "user_history"
CREATE INDEX `userhistory_history_time` ON `user_history` (`history_time`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "userhistory_history_time" to table: "user_history"
DROP INDEX `userhistory_history_time`;
-- reverse: create "new_user_history" table
DROP TABLE `new_user_history`;
-- reverse: create index "user_email_auth_provider" to table: "users"
DROP INDEX `user_email_auth_provider`;
-- reverse: create index "user_id" to table: "users"
DROP INDEX `user_id`;
-- reverse: create index "users_sub_key" to table: "users"
DROP INDEX `users_sub_key`;
-- reverse: create index "users_mapping_id_key" to table: "users"
DROP INDEX `users_mapping_id_key`;
-- reverse: create "new_users" table
DROP TABLE `new_users`;
