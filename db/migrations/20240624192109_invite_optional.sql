-- Modify "invites" table
ALTER TABLE "invites" ALTER COLUMN "expires" DROP NOT NULL, ALTER COLUMN "requestor_id" DROP NOT NULL;
