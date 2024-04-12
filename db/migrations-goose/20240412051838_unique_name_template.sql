-- +goose Up
-- drop index "template_name" from table: "templates"
DROP INDEX `template_name`;
-- create index "template_name_owner_id" to table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id` ON `templates` (`name`, `owner_id`) WHERE deleted_at is NULL;

-- +goose Down
-- reverse: create index "template_name_owner_id" to table: "templates"
DROP INDEX `template_name_owner_id`;
-- reverse: drop index "template_name" from table: "templates"
CREATE UNIQUE INDEX `template_name` ON `templates` (`name`) WHERE deleted_at is NULL;
