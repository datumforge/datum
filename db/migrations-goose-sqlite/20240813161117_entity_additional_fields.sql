-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_entities" table
CREATE TABLE `new_entities` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `name` text NULL, `display_name` text NULL, `description` text NULL, `domains` json NULL, `status` text NULL DEFAULT ('active'), `entity_type_id` text NULL, `entity_type_entities` text NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entities_entity_types_entity_type` FOREIGN KEY (`entity_type_id`) REFERENCES `entity_types` (`id`) ON DELETE SET NULL, CONSTRAINT `entities_entity_types_entities` FOREIGN KEY (`entity_type_entities`) REFERENCES `entity_types` (`id`) ON DELETE SET NULL, CONSTRAINT `entities_organizations_entities` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- copy rows from old table "entities" to new temporary table "new_entities"
INSERT INTO `new_entities` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `tags`, `name`, `display_name`, `description`, `entity_type_id`, `entity_type_entities`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `tags`, `name`, `display_name`, `description`, `entity_type_id`, `entity_type_entities`, `owner_id` FROM `entities`;
-- drop "entities" table after copying rows
DROP TABLE `entities`;
-- rename temporary table "new_entities" to "entities"
ALTER TABLE `new_entities` RENAME TO `entities`;
-- create index "entities_mapping_id_key" to table: "entities"
CREATE UNIQUE INDEX `entities_mapping_id_key` ON `entities` (`mapping_id`);
-- create index "entity_name_owner_id" to table: "entities"
CREATE UNIQUE INDEX `entity_name_owner_id` ON `entities` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- create "new_entity_history" table
CREATE TABLE `new_entity_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `ref` text NULL, `operation` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `owner_id` text NULL, `name` text NULL, `display_name` text NULL, `description` text NULL, `domains` json NULL, `entity_type_id` text NULL, `status` text NULL DEFAULT ('active'), PRIMARY KEY (`id`));
-- copy rows from old table "entity_history" to new temporary table "new_entity_history"
INSERT INTO `new_entity_history` (`id`, `history_time`, `ref`, `operation`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `tags`, `owner_id`, `name`, `display_name`, `description`, `entity_type_id`) SELECT `id`, `history_time`, `ref`, `operation`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `deleted_at`, `deleted_by`, `tags`, `owner_id`, `name`, `display_name`, `description`, `entity_type_id` FROM `entity_history`;
-- drop "entity_history" table after copying rows
DROP TABLE `entity_history`;
-- rename temporary table "new_entity_history" to "entity_history"
ALTER TABLE `new_entity_history` RENAME TO `entity_history`;
-- create index "entityhistory_history_time" to table: "entity_history"
CREATE INDEX `entityhistory_history_time` ON `entity_history` (`history_time`);
-- create "notes" table
CREATE TABLE `notes` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `text` text NOT NULL, `entity_notes` text NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `notes_entities_notes` FOREIGN KEY (`entity_notes`) REFERENCES `entities` (`id`) ON DELETE SET NULL, CONSTRAINT `notes_organizations_notes` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- create index "notes_mapping_id_key" to table: "notes"
CREATE UNIQUE INDEX `notes_mapping_id_key` ON `notes` (`mapping_id`);
-- create "note_history" table
CREATE TABLE `note_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `ref` text NULL, `operation` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `owner_id` text NULL, `text` text NOT NULL, PRIMARY KEY (`id`));
-- create index "notehistory_history_time" to table: "note_history"
CREATE INDEX `notehistory_history_time` ON `note_history` (`history_time`);
-- create "entity_files" table
CREATE TABLE `entity_files` (`entity_id` text NOT NULL, `file_id` text NOT NULL, PRIMARY KEY (`entity_id`, `file_id`), CONSTRAINT `entity_files_entity_id` FOREIGN KEY (`entity_id`) REFERENCES `entities` (`id`) ON DELETE CASCADE, CONSTRAINT `entity_files_file_id` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE CASCADE);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create "entity_files" table
DROP TABLE `entity_files`;
-- reverse: create index "notehistory_history_time" to table: "note_history"
DROP INDEX `notehistory_history_time`;
-- reverse: create "note_history" table
DROP TABLE `note_history`;
-- reverse: create index "notes_mapping_id_key" to table: "notes"
DROP INDEX `notes_mapping_id_key`;
-- reverse: create "notes" table
DROP TABLE `notes`;
-- reverse: create index "entityhistory_history_time" to table: "entity_history"
DROP INDEX `entityhistory_history_time`;
-- reverse: create "new_entity_history" table
DROP TABLE `new_entity_history`;
-- reverse: create index "entity_name_owner_id" to table: "entities"
DROP INDEX `entity_name_owner_id`;
-- reverse: create index "entities_mapping_id_key" to table: "entities"
DROP INDEX `entities_mapping_id_key`;
-- reverse: create "new_entities" table
DROP TABLE `new_entities`;
