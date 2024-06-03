//go:build ignore

package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	// supported ent database drivers
	_ "github.com/datumforge/entx" // overlay for sqlite
	_ "github.com/lib/pq"          // postgres driver
	"github.com/ory/dockertest"
	_ "github.com/tursodatabase/libsql-client-go/libsql" // libsql driver
	_ "modernc.org/sqlite"                               // sqlite driver (non-cgo)

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	atlas "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"github.com/datumforge/datum/internal/ent/generated/migrate"
	"github.com/datumforge/datum/pkg/testutils"
)

func main() {
	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	atlasDir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	gooseDirSqlite, err := sqltool.NewGooseDir("migrations-goose-sqlite")
	if err != nil {
		log.Fatalf("failed creating goose migration directory: %v", err)
	}

	gooseDirPG, err := sqltool.NewGooseDir("migrations-goose-postgres")
	if err != nil {
		log.Fatalf("failed creating goose migration directory: %v", err)
	}

	// Migrate diff options.
	baseOpts := []schema.MigrateOption{
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	sqliteOpts := append(baseOpts, schema.WithDialect(dialect.SQLite))
	postgresOpts := append(baseOpts, schema.WithDialect(dialect.Postgres))

	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod create_migration.go <name>'")
	}

	sqliteDBURI, ok := os.LookupEnv("ATLAS_SQLITE_DB_URI")
	if !ok {
		log.Fatalln("failed to load the ATLAS_SQLITE_DB_URI env var")
	}

	tf := createPostgresTestContainer()
	defer testutils.TeardownFixture(tf)

	// Generate migrations using Atlas support for sqlite (note the Ent dialect option passed above).
	atlasOpts := append(baseOpts,
		schema.WithDialect(dialect.Postgres),
		schema.WithDir(atlasDir),
		schema.WithFormatter(atlas.DefaultFormatter),
	)

	if err := migrate.NamedDiff(ctx, tf.URI, os.Args[1], atlasOpts...); err != nil {
		log.Fatalf("failed generating atlas migration file: %v", err)
	}

	// Generate migrations using Goose support for sqlite
	gooseOptsSQLite := append(sqliteOpts, schema.WithDir(gooseDirSqlite))

	if err = migrate.NamedDiff(ctx, sqliteDBURI, os.Args[1], gooseOptsSQLite...); err != nil {
		log.Fatalf("failed generating goose migration file for sqlite: %v", err)
	}

	// Generate migrations using Goose support for postgres
	gooseOptsPG := append(postgresOpts, schema.WithDir(gooseDirPG))

	if err = migrate.NamedDiff(ctx, tf.URI, os.Args[1], gooseOptsPG...); err != nil {
		log.Fatalf("failed generating goose migration file for postgres: %v", err)
	}
}

// createPostgresTestContainer creates a test postgres container and waits for it to be ready to accept connections
func createPostgresTestContainer() *testutils.TestFixture {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	pgDBURI, ok := os.LookupEnv("ATLAS_POSTGRES_DB_URI")
	if !ok {
		log.Fatalln("failed to load the ATLAS_POSTGRES_DB_URI env var")
	}

	// create a test postgres container
	tf := testutils.GetTestURI(pgDBURI, 5) // allow the container to live for 5 minutes

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		db, err := sql.Open("postgres", tf.URI)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	return tf
}
