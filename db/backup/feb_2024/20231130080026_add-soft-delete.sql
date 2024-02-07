-- Add column "deleted_at" to table: "organizations"
ALTER TABLE `organizations` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "organizations"
ALTER TABLE `organizations` ADD COLUMN `deleted_by` text NULL;
