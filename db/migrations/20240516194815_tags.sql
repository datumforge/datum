-- Modify "api_tokens" table
ALTER TABLE "api_tokens" ADD COLUMN "tags" jsonb NULL;
-- Modify "document_data" table
ALTER TABLE "document_data" ADD COLUMN "tags" jsonb NULL;
-- Modify "document_data_history" table
ALTER TABLE "document_data_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "entitlement_history" table
ALTER TABLE "entitlement_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "entitlements" table
ALTER TABLE "entitlements" ADD COLUMN "tags" jsonb NULL;
-- Modify "event_history" table
ALTER TABLE "event_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "events" table
ALTER TABLE "events" ADD COLUMN "tags" jsonb NULL;
-- Modify "feature_history" table
ALTER TABLE "feature_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "features" table
ALTER TABLE "features" ADD COLUMN "tags" jsonb NULL;
-- Modify "file_history" table
ALTER TABLE "file_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "files" table
ALTER TABLE "files" ADD COLUMN "tags" jsonb NULL;
-- Modify "group_history" table
ALTER TABLE "group_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "groups" table
ALTER TABLE "groups" ADD COLUMN "tags" jsonb NULL;
-- Modify "integration_history" table
ALTER TABLE "integration_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "integrations" table
ALTER TABLE "integrations" ADD COLUMN "tags" jsonb NULL;
-- Modify "oauth_provider_history" table
ALTER TABLE "oauth_provider_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "oauth_providers" table
ALTER TABLE "oauth_providers" ADD COLUMN "tags" jsonb NULL;
-- Modify "oh_auth_too_tokens" table
ALTER TABLE "oh_auth_too_tokens" ADD COLUMN "tags" jsonb NULL;
-- Modify "organization_history" table
ALTER TABLE "organization_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "organizations" table
ALTER TABLE "organizations" ADD COLUMN "tags" jsonb NULL;
-- Modify "personal_access_tokens" table
ALTER TABLE "personal_access_tokens" ADD COLUMN "tags" jsonb NULL;
-- Modify "subscribers" table
ALTER TABLE "subscribers" ADD COLUMN "tags" jsonb NULL;
-- Modify "template_history" table
ALTER TABLE "template_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "templates" table
ALTER TABLE "templates" ADD COLUMN "tags" jsonb NULL;
-- Modify "tfa_settings" table
ALTER TABLE "tfa_settings" ADD COLUMN "tags" jsonb NULL;
-- Modify "user_history" table
ALTER TABLE "user_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "tags" jsonb NULL;
-- Modify "webauthns" table
ALTER TABLE "webauthns" ADD COLUMN "tags" jsonb NULL;
-- Modify "webhook_history" table
ALTER TABLE "webhook_history" ADD COLUMN "tags" jsonb NULL;
-- Modify "webhooks" table
ALTER TABLE "webhooks" ADD COLUMN "tags" jsonb NULL;
