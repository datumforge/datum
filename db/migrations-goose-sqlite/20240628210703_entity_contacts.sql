-- +goose Up
-- create "contacts" table
CREATE TABLE `contacts` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `full_name` text NOT NULL, `title` text NULL, `company` text NULL, `email` text NULL, `phone_number` text NULL, `address` text NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `contacts_organizations_contacts` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- create index "contacts_mapping_id_key" to table: "contacts"
CREATE UNIQUE INDEX `contacts_mapping_id_key` ON `contacts` (`mapping_id`);
-- create "contact_history" table
CREATE TABLE `contact_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `owner_id` text NULL, `full_name` text NOT NULL, `title` text NULL, `company` text NULL, `email` text NULL, `phone_number` text NULL, `address` text NULL, `status` text NOT NULL DEFAULT ('ACTIVE'), PRIMARY KEY (`id`));
-- create index "contacthistory_history_time" to table: "contact_history"
CREATE INDEX `contacthistory_history_time` ON `contact_history` (`history_time`);
-- create "entities" table
CREATE TABLE `entities` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `entity_type` text NOT NULL DEFAULT ('ORGANIZATION'), `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entities_organizations_entities` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- create index "entities_mapping_id_key" to table: "entities"
CREATE UNIQUE INDEX `entities_mapping_id_key` ON `entities` (`mapping_id`);
-- create index "entity_name_owner_id" to table: "entities"
CREATE UNIQUE INDEX `entity_name_owner_id` ON `entities` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- create "entity_history" table
CREATE TABLE `entity_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `owner_id` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `entity_type` text NOT NULL DEFAULT ('ORGANIZATION'), PRIMARY KEY (`id`));
-- create index "entityhistory_history_time" to table: "entity_history"
CREATE INDEX `entityhistory_history_time` ON `entity_history` (`history_time`);
-- create "entity_contacts" table
CREATE TABLE `entity_contacts` (`entity_id` text NOT NULL, `contact_id` text NOT NULL, PRIMARY KEY (`entity_id`, `contact_id`), CONSTRAINT `entity_contacts_entity_id` FOREIGN KEY (`entity_id`) REFERENCES `entities` (`id`) ON DELETE CASCADE, CONSTRAINT `entity_contacts_contact_id` FOREIGN KEY (`contact_id`) REFERENCES `contacts` (`id`) ON DELETE CASCADE);
-- create "entity_documents" table
CREATE TABLE `entity_documents` (`entity_id` text NOT NULL, `document_data_id` text NOT NULL, PRIMARY KEY (`entity_id`, `document_data_id`), CONSTRAINT `entity_documents_entity_id` FOREIGN KEY (`entity_id`) REFERENCES `entities` (`id`) ON DELETE CASCADE, CONSTRAINT `entity_documents_document_data_id` FOREIGN KEY (`document_data_id`) REFERENCES `document_data` (`id`) ON DELETE CASCADE);

-- +goose Down
-- reverse: create "entity_documents" table
DROP TABLE `entity_documents`;
-- reverse: create "entity_contacts" table
DROP TABLE `entity_contacts`;
-- reverse: create index "entityhistory_history_time" to table: "entity_history"
DROP INDEX `entityhistory_history_time`;
-- reverse: create "entity_history" table
DROP TABLE `entity_history`;
-- reverse: create index "entity_name_owner_id" to table: "entities"
DROP INDEX `entity_name_owner_id`;
-- reverse: create index "entities_mapping_id_key" to table: "entities"
DROP INDEX `entities_mapping_id_key`;
-- reverse: create "entities" table
DROP TABLE `entities`;
-- reverse: create index "contacthistory_history_time" to table: "contact_history"
DROP INDEX `contacthistory_history_time`;
-- reverse: create "contact_history" table
DROP TABLE `contact_history`;
-- reverse: create index "contacts_mapping_id_key" to table: "contacts"
DROP INDEX `contacts_mapping_id_key`;
-- reverse: create "contacts" table
DROP TABLE `contacts`;
