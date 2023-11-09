-- Create "entitlements" table
CREATE TABLE `entitlements` (`id` text NOT NULL, `tier` text NOT NULL DEFAULT ('free'), `stripe_customer_id` text NULL, `stripe_subscription_id` text NULL, `expires_at` datetime NULL, `cancelled` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
