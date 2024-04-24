-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_user_settings" table
CREATE TABLE `new_user_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NULL, `is_webauthn_allowed` bool NULL DEFAULT (false), `is_tfa_enabled` bool NULL DEFAULT (false), `phone_number` text NULL, `user_id` text NULL, `user_setting_default_org` text NULL, PRIMARY KEY (`id`), CONSTRAINT `user_settings_users_setting` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `user_settings_organizations_default_org` FOREIGN KEY (`user_setting_default_org`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- copy rows from old table "user_settings" to new temporary table "new_user_settings"
INSERT INTO `new_user_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags`, `is_webauthn_allowed`, `is_tfa_enabled`, `phone_number`, `user_id`, `user_setting_default_org`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags`, `is_webauthn_allowed`, `is_tfa_enabled`, `phone_number`, `user_id`, `user_setting_default_org` FROM `user_settings`;
-- drop "user_settings" table after copying rows
DROP TABLE `user_settings`;
-- rename temporary table "new_user_settings" to "user_settings"
ALTER TABLE `new_user_settings` RENAME TO `user_settings`;
-- create index "user_settings_user_id_key" to table: "user_settings"
CREATE UNIQUE INDEX `user_settings_user_id_key` ON `user_settings` (`user_id`);
-- create "new_user_setting_history" table
CREATE TABLE `new_user_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `user_id` text NULL, `locked` bool NOT NULL DEFAULT (false), `silenced_at` datetime NULL, `suspended_at` datetime NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `email_confirmed` bool NOT NULL DEFAULT (false), `tags` json NULL, `is_webauthn_allowed` bool NULL DEFAULT (false), `is_tfa_enabled` bool NULL DEFAULT (false), `phone_number` text NULL, PRIMARY KEY (`id`));
-- copy rows from old table "user_setting_history" to new temporary table "new_user_setting_history"
INSERT INTO `new_user_setting_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `user_id`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags`, `is_webauthn_allowed`, `is_tfa_enabled`, `phone_number`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `user_id`, `locked`, `silenced_at`, `suspended_at`, `status`, `email_confirmed`, `tags`, `is_webauthn_allowed`, `is_tfa_enabled`, `phone_number` FROM `user_setting_history`;
-- drop "user_setting_history" table after copying rows
DROP TABLE `user_setting_history`;
-- rename temporary table "new_user_setting_history" to "user_setting_history"
ALTER TABLE `new_user_setting_history` RENAME TO `user_setting_history`;
-- create index "usersettinghistory_history_time" to table: "user_setting_history"
CREATE INDEX `usersettinghistory_history_time` ON `user_setting_history` (`history_time`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "usersettinghistory_history_time" to table: "user_setting_history"
DROP INDEX `usersettinghistory_history_time`;
-- reverse: create "new_user_setting_history" table
DROP TABLE `new_user_setting_history`;
-- reverse: create index "user_settings_user_id_key" to table: "user_settings"
DROP INDEX `user_settings_user_id_key`;
-- reverse: create "new_user_settings" table
DROP TABLE `new_user_settings`;
