-- Create "user_groups" table
CREATE TABLE `user_groups` (`user_id` uuid NOT NULL, `group_id` uuid NOT NULL, PRIMARY KEY (`user_id`, `group_id`), CONSTRAINT `user_groups_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_groups_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE);
