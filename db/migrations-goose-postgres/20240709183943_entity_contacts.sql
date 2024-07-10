-- +goose Up
-- create "contact_history" table
CREATE TABLE "contact_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "owner_id" character varying NULL, "full_name" character varying NOT NULL, "title" character varying NULL, "company" character varying NULL, "email" character varying NULL, "phone_number" character varying NULL, "address" character varying NULL, "status" character varying NOT NULL DEFAULT 'ACTIVE', PRIMARY KEY ("id"));
-- create index "contacthistory_history_time" to table: "contact_history"
CREATE INDEX "contacthistory_history_time" ON "contact_history" ("history_time");
-- create "entity_history" table
CREATE TABLE "entity_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "owner_id" character varying NULL, "name" character varying NOT NULL, "display_name" character varying NOT NULL DEFAULT '', "description" character varying NULL, "entity_type_id" character varying NULL, PRIMARY KEY ("id"));
-- create index "entityhistory_history_time" to table: "entity_history"
CREATE INDEX "entityhistory_history_time" ON "entity_history" ("history_time");
-- create "entity_type_history" table
CREATE TABLE "entity_type_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "owner_id" character varying NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "entitytypehistory_history_time" to table: "entity_type_history"
CREATE INDEX "entitytypehistory_history_time" ON "entity_type_history" ("history_time");
-- create "contacts" table
CREATE TABLE "contacts" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "full_name" character varying NOT NULL, "title" character varying NULL, "company" character varying NULL, "email" character varying NULL, "phone_number" character varying NULL, "address" character varying NULL, "status" character varying NOT NULL DEFAULT 'ACTIVE', "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "contacts_organizations_contacts" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "contacts_mapping_id_key" to table: "contacts"
CREATE UNIQUE INDEX "contacts_mapping_id_key" ON "contacts" ("mapping_id");
-- create "entity_types" table
CREATE TABLE "entity_types" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "name" character varying NOT NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "entity_types_organizations_entitytypes" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "entity_types_mapping_id_key" to table: "entity_types"
CREATE UNIQUE INDEX "entity_types_mapping_id_key" ON "entity_types" ("mapping_id");
-- create index "entitytype_name_owner_id" to table: "entity_types"
CREATE UNIQUE INDEX "entitytype_name_owner_id" ON "entity_types" ("name", "owner_id") WHERE (deleted_at IS NULL);
-- create "entities" table
CREATE TABLE "entities" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "name" character varying NOT NULL, "display_name" character varying NOT NULL DEFAULT '', "description" character varying NULL, "entity_type_id" character varying NULL, "entity_type_entities" character varying NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "entities_entity_types_entities" FOREIGN KEY ("entity_type_entities") REFERENCES "entity_types" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "entities_entity_types_entity_type" FOREIGN KEY ("entity_type_id") REFERENCES "entity_types" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "entities_organizations_entities" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "entities_mapping_id_key" to table: "entities"
CREATE UNIQUE INDEX "entities_mapping_id_key" ON "entities" ("mapping_id");
-- create index "entity_name_owner_id" to table: "entities"
CREATE UNIQUE INDEX "entity_name_owner_id" ON "entities" ("name", "owner_id") WHERE (deleted_at IS NULL);
-- create "entity_contacts" table
CREATE TABLE "entity_contacts" ("entity_id" character varying NOT NULL, "contact_id" character varying NOT NULL, PRIMARY KEY ("entity_id", "contact_id"), CONSTRAINT "entity_contacts_contact_id" FOREIGN KEY ("contact_id") REFERENCES "contacts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "entity_contacts_entity_id" FOREIGN KEY ("entity_id") REFERENCES "entities" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- create "entity_documents" table
CREATE TABLE "entity_documents" ("entity_id" character varying NOT NULL, "document_data_id" character varying NOT NULL, PRIMARY KEY ("entity_id", "document_data_id"), CONSTRAINT "entity_documents_document_data_id" FOREIGN KEY ("document_data_id") REFERENCES "document_data" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "entity_documents_entity_id" FOREIGN KEY ("entity_id") REFERENCES "entities" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);

-- +goose Down
-- reverse: create "entity_documents" table
DROP TABLE "entity_documents";
-- reverse: create "entity_contacts" table
DROP TABLE "entity_contacts";
-- reverse: create index "entity_name_owner_id" to table: "entities"
DROP INDEX "entity_name_owner_id";
-- reverse: create index "entities_mapping_id_key" to table: "entities"
DROP INDEX "entities_mapping_id_key";
-- reverse: create "entities" table
DROP TABLE "entities";
-- reverse: create index "entitytype_name_owner_id" to table: "entity_types"
DROP INDEX "entitytype_name_owner_id";
-- reverse: create index "entity_types_mapping_id_key" to table: "entity_types"
DROP INDEX "entity_types_mapping_id_key";
-- reverse: create "entity_types" table
DROP TABLE "entity_types";
-- reverse: create index "contacts_mapping_id_key" to table: "contacts"
DROP INDEX "contacts_mapping_id_key";
-- reverse: create "contacts" table
DROP TABLE "contacts";
-- reverse: create index "entitytypehistory_history_time" to table: "entity_type_history"
DROP INDEX "entitytypehistory_history_time";
-- reverse: create "entity_type_history" table
DROP TABLE "entity_type_history";
-- reverse: create index "entityhistory_history_time" to table: "entity_history"
DROP INDEX "entityhistory_history_time";
-- reverse: create "entity_history" table
DROP TABLE "entity_history";
-- reverse: create index "contacthistory_history_time" to table: "contact_history"
DROP INDEX "contacthistory_history_time";
-- reverse: create "contact_history" table
DROP TABLE "contact_history";
