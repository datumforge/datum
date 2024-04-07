-- Create "organization_history" table
CREATE TABLE `organization_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `parent_organization_id` text NULL, `personal_org` bool NULL DEFAULT (false), `avatar_remote_url` text NULL, `dedicated_db` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "organizationhistory_history_time" to table: "organization_history"
CREATE INDEX `organizationhistory_history_time` ON `organization_history` (`history_time`);
-- Create "organization_setting_history" table
CREATE TABLE `organization_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `geo_location` text NULL, `organization_id` text NULL, PRIMARY KEY (`id`));
-- Create index "organizationsettinghistory_history_time" to table: "organization_setting_history"
CREATE INDEX `organizationsettinghistory_history_time` ON `organization_setting_history` (`history_time`);
