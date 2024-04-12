-- Drop index "template_name" from table: "templates"
DROP INDEX `template_name`;
-- Create index "template_name_owner_id" to table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id` ON `templates` (`name`, `owner_id`) WHERE deleted_at is NULL;
