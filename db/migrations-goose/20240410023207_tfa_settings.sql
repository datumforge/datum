-- +goose Up
-- drop index "tfasettings_owner_id" from table: "tfa_settings"
DROP INDEX `tfasettings_owner_id`;
-- create index "tfasetting_owner_id" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfasetting_owner_id` ON `tfa_settings` (`owner_id`) WHERE deleted_at is NULL;

-- +goose Down
-- reverse: create index "tfasetting_owner_id" to table: "tfa_settings"
DROP INDEX `tfasetting_owner_id`;
-- reverse: drop index "tfasettings_owner_id" from table: "tfa_settings"
CREATE UNIQUE INDEX `tfasettings_owner_id` ON `tfa_settings` (`owner_id`) WHERE deleted_at is NULL;
