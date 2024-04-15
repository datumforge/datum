-- Drop index "template_name_owner_id" from table: "templates"
DROP INDEX `template_name_owner_id`;
-- Create index "template_name_owner_id_type" to table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id_type` ON `templates` (`name`, `owner_id`, `type`) WHERE deleted_at is NULL;
