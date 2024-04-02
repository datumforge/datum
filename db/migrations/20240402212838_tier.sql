-- Create "tiers" table
CREATE TABLE `tiers` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `organization_id` text NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `tiers_organizations_tiers` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "tier_name_owner_id" to table: "tiers"
CREATE UNIQUE INDEX `tier_name_owner_id` ON `tiers` (`name`, `owner_id`) WHERE deleted_at is NULL;
