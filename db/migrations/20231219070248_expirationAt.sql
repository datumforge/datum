-- Add column "expiration_at" to table: "personal_access_tokens"
ALTER TABLE `personal_access_tokens` ADD COLUMN `expiration_at` datetime NOT NULL;
