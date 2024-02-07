-- Add column "deleted_at" to table: "group_memberships"
ALTER TABLE `group_memberships` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "group_memberships"
ALTER TABLE `group_memberships` ADD COLUMN `deleted_by` text NULL;
-- Add column "deleted_at" to table: "org_memberships"
ALTER TABLE `org_memberships` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "org_memberships"
ALTER TABLE `org_memberships` ADD COLUMN `deleted_by` text NULL;
