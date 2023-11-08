package testutils

import (
	"context"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"entgo.io/ent/dialect"
)

// SetupDB sets up in-memory nats server/conn, database and ent client to interact with db
func SetupDB() {
	ctx := context.Background()

	// NATS setup
	nats, err := eventtools.NewNatsServer()
	IfErrPanic("failed to start nats server", err)

	conn, err := events.NewConnection(nats.Config)

	IfErrPanic("failed to create events connection", err)

	// DB and EntClient setup
	dia, uri, cntr := ParseDBURI(ctx)

	c, err := ent.Open(dia, uri, ent.Debug(), ent.EventsPublisher(conn))
	if err != nil {
		log.Println(err)
		IfErrPanic("failed terminating test db container after failing to connect to the db", cntr.Container.Terminate(ctx))
		IfErrPanic("failed opening connection to database:", err)
	}

	switch dia {
	case dialect.SQLite:
		// Run automatic migrations for SQLite
		IfErrPanic("failed creating db schema", c.Schema.Create(ctx))
	case dialect.Postgres:
		log.Println("Running database migrations")
		goosex.MigrateUp(uri, db.Migrations)
	}

	EventsConn = conn
	EntClient = c
	DBContainer = cntr
	NATSConn = nats
}

// TeardownDB used for clean up test setup
func TeardownDB() {
	ctx := context.Background()

	if EntClient != nil {
		IfErrPanic("teardown failed to close database connection", EntClient.Close())
	}

	if DBContainer != nil && DBContainer.Container.IsRunning() {
		IfErrPanic("teardown failed to terminate test db container", DBContainer.Container.Terminate(ctx))
	}

	_ = EventsConn.Shutdown(ctx)

	NATSConn.Close()
}

// ParseDBURI parses the kind of query language from TESTDB_URI env var and initializes DBContainer as required
func ParseDBURI(ctx context.Context) (string, string, *testcontainersx.DBContainer) {
	switch {
	// if you don't pass in a database we default to an in memory sqlite
	case testDBURI == "":
		return dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", nil
	case strings.HasPrefix(testDBURI, "sqlite://"):
		return dialect.SQLite, strings.TrimPrefix(testDBURI, "sqlite://"), nil
	case strings.HasPrefix(testDBURI, "postgres://"), strings.HasPrefix(testDBURI, "postgresql://"):
		return dialect.Postgres, testDBURI, nil
	case strings.HasPrefix(testDBURI, "docker://"):
		dbImage := strings.TrimPrefix(testDBURI, "docker://")

		switch {
		case strings.HasPrefix(dbImage, "cockroach"), strings.HasPrefix(dbImage, "cockroachdb"), strings.HasPrefix(dbImage, "crdb"):
			cntr, err := testcontainersx.NewCockroachDB(ctx, dbImage)
			IfErrPanic("error starting db test container", err)

			return dialect.Postgres, cntr.URI, cntr
		case strings.HasPrefix(dbImage, "postgres"):
			_, b, _, _ := runtime.Caller(0)
			initScriptPath := filepath.Join(filepath.Dir(b), "testdata", "postgres_init.sh")

			cntr, err := testcontainersx.NewPostgresDB(ctx, dbImage,
				postgres.WithInitScripts(initScriptPath),
			)
			IfErrPanic("error starting db test container", err)

			return dialect.Postgres, cntr.URI, cntr
		default:
			panic("invalid testcontainer URI, uri: " + testDBURI)
		}

	default:
		panic("invalid DB URI, uri: " + testDBURI)
	}
}
