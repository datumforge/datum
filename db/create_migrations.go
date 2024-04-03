//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	// supported ent database drivers
	_ "github.com/datumforge/entx"                       // overlay for sqlite
	_ "github.com/lib/pq"                                // postgres driver
	_ "github.com/tursodatabase/libsql-client-go/libsql" // libsql driver
	_ "modernc.org/sqlite"                               // sqlite driver (non-cgo)

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	atlas "ariga.io/atlas/sql/migrate"
	"github.com/datumforge/datum/internal/ent/generated/migrate"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
)

func main() {
	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.SQLite),          // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod create_migration.go <name>'")
	}

	dbURI, ok := os.LookupEnv("ATLAS_DB_URI")
	if !ok {
		log.Fatalln("failed to load the ATLAS_DB_URI env var")
	}

	// Generate migrations using Atlas support for sqlite (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dbURI, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
