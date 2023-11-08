-- Add column "scopes" to table: "refresh_tokens"
ALTER TABLE `refresh_tokens` ADD COLUMN `scopes` text NULL;
-- Add column "claims_groups" to table: "refresh_tokens"
ALTER TABLE `refresh_tokens` ADD COLUMN `claims_groups` text NULL;
-- Add column "connector_data" to table: "refresh_tokens"
ALTER TABLE `refresh_tokens` ADD COLUMN `connector_data` text NULL;
