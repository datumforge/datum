-- +goose Up
-- modify "personal_access_tokens" table
ALTER TABLE "personal_access_tokens" ALTER COLUMN "expires_at" DROP NOT NULL;

-- +goose Down
-- reverse: modify "personal_access_tokens" table
ALTER TABLE "personal_access_tokens" ALTER COLUMN "expires_at" SET NOT NULL;
