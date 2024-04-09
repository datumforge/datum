-- Create "template_history" table
CREATE TABLE `template_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `owner_id` text NOT NULL, `name` text NOT NULL, `description` text NULL, `jsonconfig` json NULL, PRIMARY KEY (`id`));
-- Create index "templatehistory_history_time" to table: "template_history"
CREATE INDEX `templatehistory_history_time` ON `template_history` (`history_time`);
