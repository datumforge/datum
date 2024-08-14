-- Modify "entities" table
ALTER TABLE "entities" ALTER COLUMN "name" DROP NOT NULL, ALTER COLUMN "display_name" DROP NOT NULL, ALTER COLUMN "display_name" DROP DEFAULT, ADD COLUMN "domains" jsonb NULL, ADD COLUMN "status" character varying NULL DEFAULT 'active';
-- Modify "entity_history" table
ALTER TABLE "entity_history" ALTER COLUMN "name" DROP NOT NULL, ALTER COLUMN "display_name" DROP NOT NULL, ALTER COLUMN "display_name" DROP DEFAULT, ADD COLUMN "domains" jsonb NULL, ADD COLUMN "status" character varying NULL DEFAULT 'active';
-- Create "note_history" table
CREATE TABLE "note_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "ref" character varying NULL, "operation" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "owner_id" character varying NULL, "text" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "notehistory_history_time" to table: "note_history"
CREATE INDEX "notehistory_history_time" ON "note_history" ("history_time");
-- Create "entity_files" table
CREATE TABLE "entity_files" ("entity_id" character varying NOT NULL, "file_id" character varying NOT NULL, PRIMARY KEY ("entity_id", "file_id"), CONSTRAINT "entity_files_entity_id" FOREIGN KEY ("entity_id") REFERENCES "entities" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "entity_files_file_id" FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "notes" table
CREATE TABLE "notes" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "text" character varying NOT NULL, "entity_notes" character varying NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "notes_entities_notes" FOREIGN KEY ("entity_notes") REFERENCES "entities" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "notes_organizations_notes" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "notes_mapping_id_key" to table: "notes"
CREATE UNIQUE INDEX "notes_mapping_id_key" ON "notes" ("mapping_id");
