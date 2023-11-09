-- Create "tacos" table
CREATE TABLE `tacos` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `created_by` text NULL, `updated_by` text NULL, `tier` text NOT NULL DEFAULT ('free'), `stripe_customer_id` text NULL, `stripe_subscription_id` text NULL, `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
