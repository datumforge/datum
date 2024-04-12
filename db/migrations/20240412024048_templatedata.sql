-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_templates" table
CREATE TABLE `new_templates` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `type` text NOT NULL DEFAULT ('DOCUMENT'), `description` text NULL, `jsonconfig` json NOT NULL, `uischema` json NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `templates_organizations_templates` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "templates" to new temporary table "new_templates"
INSERT INTO `new_templates` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `jsonconfig`, `owner_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `name`, `description`, `jsonconfig`, `owner_id` FROM `templates`;
-- Drop "templates" table after copying rows
DROP TABLE `templates`;
-- Rename temporary table "new_templates" to "templates"
ALTER TABLE `new_templates` RENAME TO `templates`;
-- Create index "template_name" to table: "templates"
CREATE UNIQUE INDEX `template_name` ON `templates` (`name`) WHERE deleted_at is NULL;
-- Create "document_data" table
CREATE TABLE `document_data` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `data` json NOT NULL, `template_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `document_data_templates_documents` FOREIGN KEY (`template_id`) REFERENCES `templates` (`id`) ON DELETE NO ACTION);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
