-- +goose Up
-- create "templates" table
CREATE TABLE `templates` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `jsonconfig` json NULL, `owner_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `templates_organizations_templates` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE NO ACTION);

-- +goose Down
-- reverse: create "templates" table
DROP TABLE `templates`;
