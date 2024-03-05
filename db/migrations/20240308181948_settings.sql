-- Create index "tfasettings_owner_id" to table: "tfa_settings"
CREATE UNIQUE INDEX `tfasettings_owner_id` ON `tfa_settings` (`owner_id`) WHERE deleted_at is NULL;
