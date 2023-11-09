-- Add column "created_at" to table: "entitlements"
ALTER TABLE `entitlements` ADD COLUMN `created_at` datetime NOT NULL;
-- Add column "updated_at" to table: "entitlements"
ALTER TABLE `entitlements` ADD COLUMN `updated_at` datetime NOT NULL;
-- Add column "created_by" to table: "entitlements"
ALTER TABLE `entitlements` ADD COLUMN `created_by` text NULL;
-- Add column "updated_by" to table: "entitlements"
ALTER TABLE `entitlements` ADD COLUMN `updated_by` text NULL;
