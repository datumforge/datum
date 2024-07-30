-- +goose Up
-- modify "organization_history" table
ALTER TABLE "organization_history" ALTER COLUMN "name" DROP NOT NULL, ALTER COLUMN "display_name" DROP NOT NULL, ALTER COLUMN "display_name" DROP DEFAULT;
-- modify "organizations" table
ALTER TABLE "organizations" ALTER COLUMN "name" DROP NOT NULL, ALTER COLUMN "display_name" DROP NOT NULL, ALTER COLUMN "display_name" DROP DEFAULT;

-- +goose Down
-- reverse: modify "organizations" table
ALTER TABLE "organizations" ALTER COLUMN "display_name" SET NOT NULL, ALTER COLUMN "display_name" SET DEFAULT '', ALTER COLUMN "name" SET NOT NULL;
-- reverse: modify "organization_history" table
ALTER TABLE "organization_history" ALTER COLUMN "display_name" SET NOT NULL, ALTER COLUMN "display_name" SET DEFAULT '', ALTER COLUMN "name" SET NOT NULL;
