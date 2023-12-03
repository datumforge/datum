-- Add column "user_type" to table: "users"
ALTER TABLE `users` ADD COLUMN `user_type` text NOT NULL;
-- Drop index "user_id" from table: "users"
DROP INDEX `user_id`;
-- Add column "code" to table: "organizations"
ALTER TABLE `organizations` ADD COLUMN `code` text NULL;
