-- Create "subscribers" table
CREATE TABLE `subscribers` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `email` text NOT NULL, `phone_number` text NULL, `verified_email` bool NOT NULL DEFAULT (false), `verified_phone` bool NOT NULL DEFAULT (false), `active` bool NOT NULL DEFAULT (false), `token` text NOT NULL, `ttl` datetime NOT NULL, `secret` blob NOT NULL, `owner_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `subscribers_organizations_subscribers` FOREIGN KEY (`owner_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Create index "subscribers_token_key" to table: "subscribers"
CREATE UNIQUE INDEX `subscribers_token_key` ON `subscribers` (`token`);
-- Create index "subscriber_email_active_owner_id" to table: "subscribers"
CREATE UNIQUE INDEX `subscriber_email_active_owner_id` ON `subscribers` (`email`, `active`, `owner_id`) WHERE deleted_at is NULL;
