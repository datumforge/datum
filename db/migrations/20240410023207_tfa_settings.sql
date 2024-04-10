-- Drop index "tfasettings_owner_id" from table: "tfa_settings"
DROP INDEX `tfasettings_owner_id`;
-- Create index "tfasetting_owner_id" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfasetting_owner_id` ON `tfa_settings` (`owner_id`) WHERE deleted_at is NULL;
