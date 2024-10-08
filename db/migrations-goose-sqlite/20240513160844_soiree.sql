-- +goose Up
-- create "integration_webhooks" table
CREATE TABLE `integration_webhooks` (`integration_id` text NOT NULL, `webhook_id` text NOT NULL, PRIMARY KEY (`integration_id`, `webhook_id`), CONSTRAINT `integration_webhooks_integration_id` FOREIGN KEY (`integration_id`) REFERENCES `integrations` (`id`) ON DELETE CASCADE, CONSTRAINT `integration_webhooks_webhook_id` FOREIGN KEY (`webhook_id`) REFERENCES `webhooks` (`id`) ON DELETE CASCADE);

-- +goose Down
-- reverse: create "integration_webhooks" table
DROP TABLE `integration_webhooks`;
