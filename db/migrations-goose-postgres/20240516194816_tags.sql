-- +goose Up
-- modify "api_tokens" table
ALTER TABLE "api_tokens" ADD COLUMN "tags" jsonb NULL;
-- modify "document_data" table
ALTER TABLE "document_data" ADD COLUMN "tags" jsonb NULL;
-- modify "document_data_history" table
ALTER TABLE "document_data_history" ADD COLUMN "tags" jsonb NULL;
-- modify "entitlement_history" table
ALTER TABLE "entitlement_history" ADD COLUMN "tags" jsonb NULL;
-- modify "entitlements" table
ALTER TABLE "entitlements" ADD COLUMN "tags" jsonb NULL;
-- modify "event_history" table
ALTER TABLE "event_history" ADD COLUMN "tags" jsonb NULL;
-- modify "events" table
ALTER TABLE "events" ADD COLUMN "tags" jsonb NULL;
-- modify "feature_history" table
ALTER TABLE "feature_history" ADD COLUMN "tags" jsonb NULL;
-- modify "features" table
ALTER TABLE "features" ADD COLUMN "tags" jsonb NULL;
-- modify "file_history" table
ALTER TABLE "file_history" ADD COLUMN "tags" jsonb NULL;
-- modify "files" table
ALTER TABLE "files" ADD COLUMN "tags" jsonb NULL;
-- modify "group_history" table
ALTER TABLE "group_history" ADD COLUMN "tags" jsonb NULL;
-- modify "groups" table
ALTER TABLE "groups" ADD COLUMN "tags" jsonb NULL;
-- modify "integration_history" table
ALTER TABLE "integration_history" ADD COLUMN "tags" jsonb NULL;
-- modify "integrations" table
ALTER TABLE "integrations" ADD COLUMN "tags" jsonb NULL;
-- modify "oauth_provider_history" table
ALTER TABLE "oauth_provider_history" ADD COLUMN "tags" jsonb NULL;
-- modify "oauth_providers" table
ALTER TABLE "oauth_providers" ADD COLUMN "tags" jsonb NULL;
-- modify "oh_auth_too_tokens" table
ALTER TABLE "oh_auth_too_tokens" ADD COLUMN "tags" jsonb NULL;
-- modify "organization_history" table
ALTER TABLE "organization_history" ADD COLUMN "tags" jsonb NULL;
-- modify "organizations" table
ALTER TABLE "organizations" ADD COLUMN "tags" jsonb NULL;
-- modify "personal_access_tokens" table
ALTER TABLE "personal_access_tokens" ADD COLUMN "tags" jsonb NULL;
-- modify "subscribers" table
ALTER TABLE "subscribers" ADD COLUMN "tags" jsonb NULL;
-- modify "template_history" table
ALTER TABLE "template_history" ADD COLUMN "tags" jsonb NULL;
-- modify "templates" table
ALTER TABLE "templates" ADD COLUMN "tags" jsonb NULL;
-- modify "tfa_settings" table
ALTER TABLE "tfa_settings" ADD COLUMN "tags" jsonb NULL;
-- modify "user_history" table
ALTER TABLE "user_history" ADD COLUMN "tags" jsonb NULL;
-- modify "users" table
ALTER TABLE "users" ADD COLUMN "tags" jsonb NULL;
-- modify "webauthns" table
ALTER TABLE "webauthns" ADD COLUMN "tags" jsonb NULL;
-- modify "webhook_history" table
ALTER TABLE "webhook_history" ADD COLUMN "tags" jsonb NULL;
-- modify "webhooks" table
ALTER TABLE "webhooks" ADD COLUMN "tags" jsonb NULL;

-- +goose Down
-- reverse: modify "webhooks" table
ALTER TABLE "webhooks" DROP COLUMN "tags";
-- reverse: modify "webhook_history" table
ALTER TABLE "webhook_history" DROP COLUMN "tags";
-- reverse: modify "webauthns" table
ALTER TABLE "webauthns" DROP COLUMN "tags";
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "tags";
-- reverse: modify "user_history" table
ALTER TABLE "user_history" DROP COLUMN "tags";
-- reverse: modify "tfa_settings" table
ALTER TABLE "tfa_settings" DROP COLUMN "tags";
-- reverse: modify "templates" table
ALTER TABLE "templates" DROP COLUMN "tags";
-- reverse: modify "template_history" table
ALTER TABLE "template_history" DROP COLUMN "tags";
-- reverse: modify "subscribers" table
ALTER TABLE "subscribers" DROP COLUMN "tags";
-- reverse: modify "personal_access_tokens" table
ALTER TABLE "personal_access_tokens" DROP COLUMN "tags";
-- reverse: modify "organizations" table
ALTER TABLE "organizations" DROP COLUMN "tags";
-- reverse: modify "organization_history" table
ALTER TABLE "organization_history" DROP COLUMN "tags";
-- reverse: modify "oh_auth_too_tokens" table
ALTER TABLE "oh_auth_too_tokens" DROP COLUMN "tags";
-- reverse: modify "oauth_providers" table
ALTER TABLE "oauth_providers" DROP COLUMN "tags";
-- reverse: modify "oauth_provider_history" table
ALTER TABLE "oauth_provider_history" DROP COLUMN "tags";
-- reverse: modify "integrations" table
ALTER TABLE "integrations" DROP COLUMN "tags";
-- reverse: modify "integration_history" table
ALTER TABLE "integration_history" DROP COLUMN "tags";
-- reverse: modify "groups" table
ALTER TABLE "groups" DROP COLUMN "tags";
-- reverse: modify "group_history" table
ALTER TABLE "group_history" DROP COLUMN "tags";
-- reverse: modify "files" table
ALTER TABLE "files" DROP COLUMN "tags";
-- reverse: modify "file_history" table
ALTER TABLE "file_history" DROP COLUMN "tags";
-- reverse: modify "features" table
ALTER TABLE "features" DROP COLUMN "tags";
-- reverse: modify "feature_history" table
ALTER TABLE "feature_history" DROP COLUMN "tags";
-- reverse: modify "events" table
ALTER TABLE "events" DROP COLUMN "tags";
-- reverse: modify "event_history" table
ALTER TABLE "event_history" DROP COLUMN "tags";
-- reverse: modify "entitlements" table
ALTER TABLE "entitlements" DROP COLUMN "tags";
-- reverse: modify "entitlement_history" table
ALTER TABLE "entitlement_history" DROP COLUMN "tags";
-- reverse: modify "document_data_history" table
ALTER TABLE "document_data_history" DROP COLUMN "tags";
-- reverse: modify "document_data" table
ALTER TABLE "document_data" DROP COLUMN "tags";
-- reverse: modify "api_tokens" table
ALTER TABLE "api_tokens" DROP COLUMN "tags";
