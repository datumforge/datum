-- Create "ent_types" table
CREATE TABLE `ent_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `type` text NOT NULL);
-- Create index "ent_types_type_key" to table: "ent_types"
CREATE UNIQUE INDEX `ent_types_type_key` ON `ent_types` (`type`);
-- Add pk ranges for ('groups'),('group_settings'),('integrations'),('memberships'),('organizations'),('sessions'),('users') tables
INSERT INTO `ent_types` (`type`) VALUES ('groups'), ('group_settings'), ('integrations'), ('memberships'), ('organizations'), ('sessions'), ('users');
