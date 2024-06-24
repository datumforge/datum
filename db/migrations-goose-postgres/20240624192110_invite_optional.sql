-- +goose Up
-- modify "invites" table
ALTER TABLE "invites" ALTER COLUMN "expires" DROP NOT NULL, ALTER COLUMN "requestor_id" DROP NOT NULL;

-- +goose Down
-- reverse: modify "invites" table
ALTER TABLE "invites" ALTER COLUMN "requestor_id" SET NOT NULL, ALTER COLUMN "expires" SET NOT NULL;
