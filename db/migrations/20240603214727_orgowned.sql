-- Modify "document_data_history" table
ALTER TABLE "document_data_history" ADD COLUMN "owner_id" character varying NULL;
-- Modify "oauth_provider_history" table
ALTER TABLE "oauth_provider_history" ADD COLUMN "owner_id" character varying NULL;
-- Modify "document_data" table
ALTER TABLE "document_data" ADD COLUMN "owner_id" character varying NULL, ADD CONSTRAINT "document_data_organizations_documentdata" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "oauth_providers" table
ALTER TABLE "oauth_providers" DROP CONSTRAINT "oauth_providers_organizations_oauthprovider", DROP COLUMN "organization_oauthprovider", ADD COLUMN "owner_id" character varying NULL, ADD CONSTRAINT "oauth_providers_organizations_oauthprovider" FOREIGN KEY ("owner_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
