-- +goose Up
-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_entitlements" table
CREATE TABLE `new_entitlements` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), `plan_id` text NOT NULL, `owner_id` text NULL, `organization_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlements_entitlement_plans_entitlements` FOREIGN KEY (`plan_id`) REFERENCES `entitlement_plans` (`id`) ON DELETE NO ACTION, CONSTRAINT `entitlements_organizations_entitlements` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL, CONSTRAINT `entitlements_organizations_organization_entitlement` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- copy rows from old table "entitlements" to new temporary table "new_entitlements"
INSERT INTO `new_entitlements` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`, `owner_id` FROM `entitlements`;
-- drop "entitlements" table after copying rows
DROP TABLE `entitlements`;
-- rename temporary table "new_entitlements" to "entitlements"
ALTER TABLE `new_entitlements` RENAME TO `entitlements`;
-- create index "entitlements_mapping_id_key" to table: "entitlements"
CREATE UNIQUE INDEX `entitlements_mapping_id_key` ON `entitlements` (`mapping_id`);
-- create index "entitlement_organization_id_owner_id" to table: "entitlements"
CREATE UNIQUE INDEX `entitlement_organization_id_owner_id` ON `entitlements` (`organization_id`, `owner_id`) WHERE deleted_at is NULL and cancelled = false;
-- create "new_entitlement_history" table
CREATE TABLE `new_entitlement_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NULL, `plan_id` text NOT NULL, `organization_id` text NOT NULL, `external_customer_id` text NULL, `external_subscription_id` text NULL, `expires` bool NOT NULL DEFAULT (false), `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- copy rows from old table "entitlement_history" to new temporary table "new_entitlement_history"
INSERT INTO `new_entitlement_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `owner_id`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `mapping_id`, `tags`, `deleted_at`, `deleted_by`, `owner_id`, `external_customer_id`, `external_subscription_id`, `expires`, `expires_at`, `cancelled` FROM `entitlement_history`;
-- drop "entitlement_history" table after copying rows
DROP TABLE `entitlement_history`;
-- rename temporary table "new_entitlement_history" to "entitlement_history"
ALTER TABLE `new_entitlement_history` RENAME TO `entitlement_history`;
-- create index "entitlementhistory_history_time" to table: "entitlement_history"
CREATE INDEX `entitlementhistory_history_time` ON `entitlement_history` (`history_time`);
-- create "new_features" table
CREATE TABLE `new_features` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `name` text NOT NULL, `display_name` text NULL, `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, `metadata` json NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `features_organizations_features` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- copy rows from old table "features" to new temporary table "new_features"
INSERT INTO `new_features` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `tags`, `name`, `enabled`, `description`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `tags`, `name`, `enabled`, `description` FROM `features`;
-- drop "features" table after copying rows
DROP TABLE `features`;
-- rename temporary table "new_features" to "features"
ALTER TABLE `new_features` RENAME TO `features`;
-- create index "features_mapping_id_key" to table: "features"
CREATE UNIQUE INDEX `features_mapping_id_key` ON `features` (`mapping_id`);
-- create index "feature_name_owner_id" to table: "features"
CREATE UNIQUE INDEX `feature_name_owner_id` ON `features` (`name`, `owner_id`) WHERE deleted_at is NULL;
-- create "new_feature_history" table
CREATE TABLE `new_feature_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `mapping_id` text NOT NULL, `tags` json NULL, `owner_id` text NULL, `name` text NOT NULL, `display_name` text NULL, `enabled` bool NOT NULL DEFAULT (false), `description` text NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- copy rows from old table "feature_history" to new temporary table "new_feature_history"
INSERT INTO `new_feature_history` (`id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `tags`, `name`, `enabled`, `description`) SELECT `id`, `history_time`, `operation`, `ref`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `mapping_id`, `tags`, `name`, `enabled`, `description` FROM `feature_history`;
-- drop "feature_history" table after copying rows
DROP TABLE `feature_history`;
-- rename temporary table "new_feature_history" to "feature_history"
ALTER TABLE `new_feature_history` RENAME TO `feature_history`;
-- create index "featurehistory_history_time" to table: "feature_history"
CREATE INDEX `featurehistory_history_time` ON `feature_history` (`history_time`);
-- create "entitlement_plans" table
CREATE TABLE `entitlement_plans` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `display_name` text NULL, `name` text NOT NULL, `description` text NULL, `version` text NOT NULL, `metadata` json NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlement_plans_organizations_entitlementplans` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- create index "entitlement_plans_mapping_id_key" to table: "entitlement_plans"
CREATE UNIQUE INDEX `entitlement_plans_mapping_id_key` ON `entitlement_plans` (`mapping_id`);
-- create index "entitlementplan_name_version_owner_id" to table: "entitlement_plans"
CREATE UNIQUE INDEX `entitlementplan_name_version_owner_id` ON `entitlement_plans` (`name`, `version`, `owner_id`) WHERE deleted_at is NULL;
-- create "entitlement_plan_history" table
CREATE TABLE `entitlement_plan_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `owner_id` text NULL, `display_name` text NULL, `name` text NOT NULL, `description` text NULL, `version` text NOT NULL, `metadata` json NULL, PRIMARY KEY (`id`));
-- create index "entitlementplanhistory_history_time" to table: "entitlement_plan_history"
CREATE INDEX `entitlementplanhistory_history_time` ON `entitlement_plan_history` (`history_time`);
-- create "entitlement_plan_features" table
CREATE TABLE `entitlement_plan_features` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `metadata` json NULL, `plan_id` text NOT NULL, `feature_id` text NOT NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `entitlement_plan_features_entitlement_plans_plan` FOREIGN KEY (`plan_id`) REFERENCES `entitlement_plans` (`id`) ON DELETE NO ACTION, CONSTRAINT `entitlement_plan_features_features_feature` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`) ON DELETE NO ACTION, CONSTRAINT `entitlement_plan_features_organizations_entitlementplanfeatures` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- create index "entitlement_plan_features_mapping_id_key" to table: "entitlement_plan_features"
CREATE UNIQUE INDEX `entitlement_plan_features_mapping_id_key` ON `entitlement_plan_features` (`mapping_id`);
-- create index "entitlementplanfeature_feature_id_plan_id" to table: "entitlement_plan_features"
CREATE UNIQUE INDEX `entitlementplanfeature_feature_id_plan_id` ON `entitlement_plan_features` (`feature_id`, `plan_id`) WHERE deleted_at is NULL;
-- create "entitlement_plan_feature_history" table
CREATE TABLE `entitlement_plan_feature_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `mapping_id` text NOT NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `tags` json NULL, `owner_id` text NULL, `metadata` json NULL, `plan_id` text NOT NULL, `feature_id` text NOT NULL, PRIMARY KEY (`id`));
-- create index "entitlementplanfeaturehistory_history_time" to table: "entitlement_plan_feature_history"
CREATE INDEX `entitlementplanfeaturehistory_history_time` ON `entitlement_plan_feature_history` (`history_time`);
-- create "entitlement_plan_events" table
CREATE TABLE `entitlement_plan_events` (`entitlement_plan_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`entitlement_plan_id`, `event_id`), CONSTRAINT `entitlement_plan_events_entitlement_plan_id` FOREIGN KEY (`entitlement_plan_id`) REFERENCES `entitlement_plans` (`id`) ON DELETE CASCADE, CONSTRAINT `entitlement_plan_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- create "entitlement_plan_feature_events" table
CREATE TABLE `entitlement_plan_feature_events` (`entitlement_plan_feature_id` text NOT NULL, `event_id` text NOT NULL, PRIMARY KEY (`entitlement_plan_feature_id`, `event_id`), CONSTRAINT `entitlement_plan_feature_events_entitlement_plan_feature_id` FOREIGN KEY (`entitlement_plan_feature_id`) REFERENCES `entitlement_plan_features` (`id`) ON DELETE CASCADE, CONSTRAINT `entitlement_plan_feature_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE CASCADE);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;

-- +goose Down
-- reverse: create "entitlement_plan_feature_events" table
DROP TABLE `entitlement_plan_feature_events`;
-- reverse: create "entitlement_plan_events" table
DROP TABLE `entitlement_plan_events`;
-- reverse: create index "entitlementplanfeaturehistory_history_time" to table: "entitlement_plan_feature_history"
DROP INDEX `entitlementplanfeaturehistory_history_time`;
-- reverse: create "entitlement_plan_feature_history" table
DROP TABLE `entitlement_plan_feature_history`;
-- reverse: create index "entitlementplanfeature_feature_id_plan_id" to table: "entitlement_plan_features"
DROP INDEX `entitlementplanfeature_feature_id_plan_id`;
-- reverse: create index "entitlement_plan_features_mapping_id_key" to table: "entitlement_plan_features"
DROP INDEX `entitlement_plan_features_mapping_id_key`;
-- reverse: create "entitlement_plan_features" table
DROP TABLE `entitlement_plan_features`;
-- reverse: create index "entitlementplanhistory_history_time" to table: "entitlement_plan_history"
DROP INDEX `entitlementplanhistory_history_time`;
-- reverse: create "entitlement_plan_history" table
DROP TABLE `entitlement_plan_history`;
-- reverse: create index "entitlementplan_name_version_owner_id" to table: "entitlement_plans"
DROP INDEX `entitlementplan_name_version_owner_id`;
-- reverse: create index "entitlement_plans_mapping_id_key" to table: "entitlement_plans"
DROP INDEX `entitlement_plans_mapping_id_key`;
-- reverse: create "entitlement_plans" table
DROP TABLE `entitlement_plans`;
-- reverse: create index "featurehistory_history_time" to table: "feature_history"
DROP INDEX `featurehistory_history_time`;
-- reverse: create "new_feature_history" table
DROP TABLE `new_feature_history`;
-- reverse: create index "feature_name_owner_id" to table: "features"
DROP INDEX `feature_name_owner_id`;
-- reverse: create index "features_mapping_id_key" to table: "features"
DROP INDEX `features_mapping_id_key`;
-- reverse: create "new_features" table
DROP TABLE `new_features`;
-- reverse: create index "entitlementhistory_history_time" to table: "entitlement_history"
DROP INDEX `entitlementhistory_history_time`;
-- reverse: create "new_entitlement_history" table
DROP TABLE `new_entitlement_history`;
-- reverse: create index "entitlement_organization_id_owner_id" to table: "entitlements"
DROP INDEX `entitlement_organization_id_owner_id`;
-- reverse: create index "entitlements_mapping_id_key" to table: "entitlements"
DROP INDEX `entitlements_mapping_id_key`;
-- reverse: create "new_entitlements" table
DROP TABLE `new_entitlements`;
