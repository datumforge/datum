-- +goose Up
-- drop index "template_name_owner_id" from table: "templates"
DROP INDEX `template_name_owner_id`;
-- create index "template_name_owner_id_type" to table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id_type` ON `templates` (`name`, `owner_id`, `type`) WHERE deleted_at is NULL;

-- +goose Down
-- reverse: create index "template_name_owner_id_type" to table: "templates"
DROP INDEX `template_name_owner_id_type`;
-- reverse: drop index "template_name_owner_id" from table: "templates"
CREATE UNIQUE INDEX `template_name_owner_id` ON `templates` (`name`, `owner_id`) WHERE deleted_at is NULL;
