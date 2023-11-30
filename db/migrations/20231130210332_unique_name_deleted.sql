-- Drop index "organizations_name_key" from table: "organizations"
DROP INDEX `organizations_name_key`;
-- Create index "organization_name" to table: "organizations"
CREATE UNIQUE INDEX `organization_name` ON `organizations` (`name`) WHERE deleted_at is NULL;
