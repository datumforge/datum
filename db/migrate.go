// Package db provides an embedded filesystem containing all the database migrations
package db

import (
	"embed"
)

// Migrations contain an embedded filesystem with all the sql migration files
//
//go:embed migrations/*.sql
var Migrations embed.FS

// GooseMigrations contain an embedded filesystem with all the goose migration files
//
//go:embed migrations-goose/*.sql
var GooseMigrations embed.FS
