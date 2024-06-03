-- +goose Up
-- modify "document_data_history" table
ALTER TABLE "document_data_history" ADD COLUMN "owner_id" character varying NULL;
-- modify "oauth_provider_history" table
ALTER TABLE "oauth_provider_history" ADD COLUMN "owner_id" character varying NULL;
-- modify "document_data" table
ALTER TABLE "document_data" ADD COLUMN "owner_id" character varying NULL, ADD CONSTRAINT "document_data_organizations_documentdata" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- modify "oauth_providers" table
ALTER TABLE "oauth_providers" DROP CONSTRAINT "oauth_providers_organizations_oauthprovider", DROP COLUMN "organization_oauthprovider", ADD COLUMN "owner_id" character varying NULL, ADD CONSTRAINT "oauth_providers_organizations_oauthprovider" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;

-- +goose Down
-- reverse: modify "oauth_providers" table
ALTER TABLE "oauth_providers" DROP CONSTRAINT "oauth_providers_organizations_oauthprovider", DROP COLUMN "owner_id", ADD COLUMN "organization_oauthprovider" character varying NULL, ADD CONSTRAINT "oauth_providers_organizations_oauthprovider" FOREIGN KEY ("organization_oauthprovider") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- reverse: modify "document_data" table
ALTER TABLE "document_data" DROP CONSTRAINT "document_data_organizations_documentdata", DROP COLUMN "owner_id";
-- reverse: modify "oauth_provider_history" table
ALTER TABLE "oauth_provider_history" DROP COLUMN "owner_id";
-- reverse: modify "document_data_history" table
ALTER TABLE "document_data_history" DROP COLUMN "owner_id";
