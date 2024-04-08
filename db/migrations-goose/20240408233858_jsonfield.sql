-- +goose Up
-- add column "thatjsonbaby" to table: "templates"
ALTER TABLE `templates` ADD COLUMN `thatjsonbaby` json NULL;

-- +goose Down
-- reverse: add column "thatjsonbaby" to table: "templates"
ALTER TABLE `templates` DROP COLUMN `thatjsonbaby`;
