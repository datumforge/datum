-- Modify "entitlement_history" table
ALTER TABLE "entitlement_history" DROP COLUMN "tier", ADD COLUMN "plan_id" character varying NOT NULL, ADD COLUMN "organization_id" character varying NOT NULL;
-- Modify "feature_history" table
ALTER TABLE "feature_history" DROP COLUMN "global", ADD COLUMN "owner_id" character varying NULL, ADD COLUMN "display_name" character varying NULL, ADD COLUMN "metadata" jsonb NULL;
-- Create "entitlement_plan_history" table
CREATE TABLE "entitlement_plan_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "owner_id" character varying NULL, "display_name" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "version" character varying NOT NULL, "metadata" jsonb NULL, PRIMARY KEY ("id"));
-- Create index "entitlementplanhistory_history_time" to table: "entitlement_plan_history"
CREATE INDEX "entitlementplanhistory_history_time" ON "entitlement_plan_history" ("history_time");
-- Create "entitlement_plan_feature_history" table
CREATE TABLE "entitlement_plan_feature_history" ("id" character varying NOT NULL, "history_time" timestamptz NOT NULL, "operation" character varying NOT NULL, "ref" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "owner_id" character varying NULL, "metadata" jsonb NULL, "plan_id" character varying NOT NULL, "feature_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "entitlementplanfeaturehistory_history_time" to table: "entitlement_plan_feature_history"
CREATE INDEX "entitlementplanfeaturehistory_history_time" ON "entitlement_plan_feature_history" ("history_time");
-- Create "entitlement_plans" table
CREATE TABLE "entitlement_plans" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "display_name" character varying NULL, "name" character varying NOT NULL, "description" character varying NULL, "version" character varying NOT NULL, "metadata" jsonb NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "entitlement_plans_organizations_entitlementplans" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "entitlement_plans_mapping_id_key" to table: "entitlement_plans"
CREATE UNIQUE INDEX "entitlement_plans_mapping_id_key" ON "entitlement_plans" ("mapping_id");
-- Create index "entitlementplan_name_version_owner_id" to table: "entitlement_plans"
CREATE UNIQUE INDEX "entitlementplan_name_version_owner_id" ON "entitlement_plans" ("name", "version", "owner_id") WHERE (deleted_at IS NULL);
-- Create "entitlement_plan_events" table
CREATE TABLE "entitlement_plan_events" ("entitlement_plan_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("entitlement_plan_id", "event_id"), CONSTRAINT "entitlement_plan_events_entitlement_plan_id" FOREIGN KEY ("entitlement_plan_id") REFERENCES "entitlement_plans" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "entitlement_plan_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Drop index "features_name_key" from table: "features"
DROP INDEX "features_name_key";
-- Modify "features" table
ALTER TABLE "features" DROP COLUMN "global", ADD COLUMN "display_name" character varying NULL, ADD COLUMN "metadata" jsonb NULL, ADD COLUMN "owner_id" character varying NULL, ADD CONSTRAINT "features_organizations_features" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Create index "feature_name_owner_id" to table: "features"
CREATE UNIQUE INDEX "feature_name_owner_id" ON "features" ("name", "owner_id") WHERE (deleted_at IS NULL);
-- Create "entitlement_plan_features" table
CREATE TABLE "entitlement_plan_features" ("id" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "created_by" character varying NULL, "updated_by" character varying NULL, "mapping_id" character varying NOT NULL, "deleted_at" timestamptz NULL, "deleted_by" character varying NULL, "tags" jsonb NULL, "metadata" jsonb NULL, "plan_id" character varying NOT NULL, "feature_id" character varying NOT NULL, "owner_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "entitlement_plan_features_entitlement_plans_plan" FOREIGN KEY ("plan_id") REFERENCES "entitlement_plans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "entitlement_plan_features_features_feature" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "entitlement_plan_features_organizations_entitlementplanfeatures" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "entitlement_plan_features_mapping_id_key" to table: "entitlement_plan_features"
CREATE UNIQUE INDEX "entitlement_plan_features_mapping_id_key" ON "entitlement_plan_features" ("mapping_id");
-- Create index "entitlementplanfeature_feature_id_plan_id" to table: "entitlement_plan_features"
CREATE UNIQUE INDEX "entitlementplanfeature_feature_id_plan_id" ON "entitlement_plan_features" ("feature_id", "plan_id") WHERE (deleted_at IS NULL);
-- Create "entitlement_plan_feature_events" table
CREATE TABLE "entitlement_plan_feature_events" ("entitlement_plan_feature_id" character varying NOT NULL, "event_id" character varying NOT NULL, PRIMARY KEY ("entitlement_plan_feature_id", "event_id"), CONSTRAINT "entitlement_plan_feature_events_entitlement_plan_feature_id" FOREIGN KEY ("entitlement_plan_feature_id") REFERENCES "entitlement_plan_features" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "entitlement_plan_feature_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Modify "entitlements" table
ALTER TABLE "entitlements" DROP COLUMN "tier", ADD COLUMN "plan_id" character varying NOT NULL, ADD COLUMN "organization_id" character varying NOT NULL, ADD CONSTRAINT "entitlements_entitlement_plans_entitlements" FOREIGN KEY ("plan_id") REFERENCES "entitlement_plans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "entitlements_organizations_organization_entitlement" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Create index "entitlement_organization_id_owner_id" to table: "entitlements"
CREATE UNIQUE INDEX "entitlement_organization_id_owner_id" ON "entitlements" ("organization_id", "owner_id") WHERE ((deleted_at IS NULL) AND (cancelled = false));
