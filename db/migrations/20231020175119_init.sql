-- +goose Up
-- create "integrations" table
CREATE TABLE `integrations` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `kind` text NOT NULL, `description` text NULL, `secret_name` text NOT NULL, `organization_integrations` uuid NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `integrations_organizations_integrations` FOREIGN KEY (`organization_integrations`) REFERENCES `organizations` (`id`) ON DELETE CASCADE);
-- create "memberships" table
CREATE TABLE `memberships` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `current` bool NOT NULL DEFAULT (false), `organization_memberships` uuid NOT NULL, `user_memberships` uuid NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `memberships_organizations_memberships` FOREIGN KEY (`organization_memberships`) REFERENCES `organizations` (`id`) ON DELETE CASCADE, CONSTRAINT `memberships_users_memberships` FOREIGN KEY (`user_memberships`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- create index "membership_organization_memberships_user_memberships" to table: "memberships"
CREATE UNIQUE INDEX `membership_organization_memberships_user_memberships` ON `memberships` (`organization_memberships`, `user_memberships`);
-- create "organizations" table
CREATE TABLE `organizations` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `name` text NOT NULL DEFAULT ('default'), PRIMARY KEY (`id`));
-- create "sessions" table
CREATE TABLE `sessions` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `type` text NOT NULL, `disabled` bool NOT NULL, `token` text NOT NULL, `user_agent` text NULL, `ips` text NOT NULL, `session_users` uuid NULL, `user_sessions` uuid NULL, PRIMARY KEY (`id`), CONSTRAINT `sessions_users_users` FOREIGN KEY (`session_users`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `sessions_users_sessions` FOREIGN KEY (`user_sessions`) REFERENCES `users` (`id`) ON DELETE CASCADE);
-- create index "sessions_token_key" to table: "sessions"
CREATE UNIQUE INDEX `sessions_token_key` ON `sessions` (`token`);
-- create index "session_id" to table: "sessions"
CREATE UNIQUE INDEX `session_id` ON `sessions` (`id`);
-- create "users" table
CREATE TABLE `users` (`id` uuid NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` integer NULL, `updated_by` integer NULL, `email` text NOT NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `display_name` text NOT NULL DEFAULT ('unknown'), `locked` bool NOT NULL DEFAULT (false), `avatar_remote_url` text NULL, `avatar_local_file` text NULL, `avatar_updated_at` datetime NULL, `silenced_at` datetime NULL, `suspended_at` datetime NULL, `recovery_code` text NULL, PRIMARY KEY (`id`));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- create index "user_id" to table: "users"
CREATE UNIQUE INDEX `user_id` ON `users` (`id`);

-- +goose Down
-- reverse: create index "user_id" to table: "users"
DROP INDEX `user_id`;
-- reverse: create index "users_email_key" to table: "users"
DROP INDEX `users_email_key`;
-- reverse: create "users" table
DROP TABLE `users`;
-- reverse: create index "session_id" to table: "sessions"
DROP INDEX `session_id`;
-- reverse: create index "sessions_token_key" to table: "sessions"
DROP INDEX `sessions_token_key`;
-- reverse: create "sessions" table
DROP TABLE `sessions`;
-- reverse: create "organizations" table
DROP TABLE `organizations`;
-- reverse: create index "membership_organization_memberships_user_memberships" to table: "memberships"
DROP INDEX `membership_organization_memberships_user_memberships`;
-- reverse: create "memberships" table
DROP TABLE `memberships`;
-- reverse: create "integrations" table
DROP TABLE `integrations`;
