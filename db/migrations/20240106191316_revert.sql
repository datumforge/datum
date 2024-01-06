-- Add column "deleted_at" to table: "personal_access_tokens"
ALTER TABLE `personal_access_tokens` ADD COLUMN `deleted_at` datetime NULL;
-- Add column "deleted_by" to table: "personal_access_tokens"
ALTER TABLE `personal_access_tokens` ADD COLUMN `deleted_by` text NULL;
