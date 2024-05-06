-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_events" table
CREATE TABLE `new_events` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `event_id` text NULL, `correlation_id` text NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- copy rows from old table "events" to new temporary table "new_events"
INSERT INTO `new_events` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `event_id`, `correlation_id`, `event_type`, `metadata`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `event_id`, `correlation_id`, `event_type`, `metadata` FROM `events`;
-- drop "events" table after copying rows
DROP TABLE `events`;
-- rename temporary table "new_events" to "events"
ALTER TABLE `new_events` RENAME TO `events`;
-- create "new_event_history" table
CREATE TABLE `new_event_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `event_id` text NULL, `correlation_id` text NULL, `event_type` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- copy rows from old table "event_history" to new temporary table "new_event_history"
INSERT INTO `new_event_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `event_id`, `correlation_id`, `event_type`, `metadata`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `event_id`, `correlation_id`, `event_type`, `metadata` FROM `event_history`;
-- drop "event_history" table after copying rows
DROP TABLE `event_history`;
-- rename temporary table "new_event_history" to "event_history"
ALTER TABLE `new_event_history` RENAME TO `event_history`;
-- create index "eventhistory_history_time" to table: "event_history"
CREATE INDEX `eventhistory_history_time` ON `event_history` (`history_time`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create index "eventhistory_history_time" to table: "event_history"
DROP INDEX `eventhistory_history_time`;
-- reverse: create "new_event_history" table
DROP TABLE `new_event_history`;
-- reverse: create "new_events" table
DROP TABLE `new_events`;
