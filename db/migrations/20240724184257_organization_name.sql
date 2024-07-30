-- Modify "organization_history" table
ALTER TABLE "organization_history" ALTER COLUMN "name" DROP NOT NULL, ALTER COLUMN "display_name" DROP NOT NULL, ALTER COLUMN "display_name" DROP DEFAULT;
-- Modify "organizations" table
ALTER TABLE "organizations" ALTER COLUMN "name" DROP NOT NULL, ALTER COLUMN "display_name" DROP NOT NULL, ALTER COLUMN "display_name" DROP DEFAULT;
