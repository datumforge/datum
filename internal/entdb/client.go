package entdb

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"ariga.io/entcache"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/datumforge/entx"
	"github.com/pressly/goose/v3"

	"go.uber.org/zap"

	migratedb "github.com/datumforge/datum/db"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/pkg/testutils"
)

type client struct {
	// config is the entdb configuration
	config *entx.Config
	// pc is the primary ent client
	pc *ent.Client
	// sc is the secondary ent client
	sc *ent.Client
	// logger holds the zap logger
	logger *zap.SugaredLogger
}

// NewMultiDriverDBClient returns a ent client with a primary and secondary, if configured, write database
func NewMultiDriverDBClient(ctx context.Context, c entx.Config, l *zap.SugaredLogger, opts []ent.Option) (*ent.Client, *entx.EntClientConfig, error) {
	client := &client{
		config: &c,
		logger: l,
	}

	dbOpts := []entx.DBOption{
		entx.WithLogger(l),
	}

	if c.MultiWrite {
		dbOpts = append(dbOpts, entx.WithSecondaryDB())
	}

	entConfig := entx.NewDBConfig(c, dbOpts...)

	// Decorates the sql.Driver with entcache.Driver on the primaryDB
	drvPrimary := entcache.NewDriver(
		entConfig.GetPrimaryDB(),
		entcache.TTL(c.CacheTTL), // set the TTL on the cache
	)

	client.pc = client.createEntDBClient(entConfig.GetPrimaryDB())

	if c.RunMigrations {
		if err := client.runMigrations(ctx); err != nil {
			client.logger.Errorf("failed running migrations", zap.Error(err))

			return nil, nil, err
		}
	}

	var cOpts []ent.Option

	if c.MultiWrite {
		// Decorates the sql.Driver with entcache.Driver on the primaryDB
		drvSecondary := entcache.NewDriver(
			entConfig.GetSecondaryDB(),
			entcache.TTL(c.CacheTTL), // set the TTL on the cache
		)

		client.sc = client.createEntDBClient(entConfig.GetSecondaryDB())

		if c.RunMigrations {
			if err := client.runMigrations(ctx); err != nil {
				client.logger.Errorf("failed running migrations", zap.Error(err))

				return nil, nil, err
			}
		}

		// Create Multiwrite driver
		cOpts = []ent.Option{ent.Driver(&entx.MultiWriteDriver{Wp: drvPrimary, Ws: drvSecondary})}
	} else {
		cOpts = []ent.Option{ent.Driver(drvPrimary)}
	}

	cOpts = append(cOpts, opts...)

	if c.Debug {
		cOpts = append(cOpts,
			ent.Log(client.logger.Named("ent").Debugln),
			ent.Debug(),
			ent.Driver(drvPrimary),
		)
	}

	ec := ent.NewClient(cOpts...)

	// add authz hooks
	ec.WithAuthz()

	if c.EnableHistory {
		// add history hooks
		ec.WithHistory()
	}

	ec.Intercept(interceptors.QueryLogger(client.logger))

	return ec, entConfig, nil
}

// runMigrations runs the migrations based on the configured migration provider on startup
func (c *client) runMigrations(ctx context.Context) error {
	switch c.config.MigrationProvider {
	case "goose":
		return c.runGooseMigrations()
	default: // atlas
		return c.runAtlasMigrations(ctx)
	}
}

// runGooseMigrations runs the goose migrations
func (c *client) runGooseMigrations() error {
	driver, err := entx.CheckEntDialect(c.config.DriverName)
	if err != nil {
		return err
	}

	drv, err := sql.Open(c.config.DriverName, c.config.PrimaryDBSource)
	if err != nil {
		return err
	}
	defer drv.Close()

	if _, err := drv.Exec("PRAGMA foreign_keys = off;", nil); err != nil {
		drv.Close()

		return fmt.Errorf("failed to disable foreign keys: %w", err)
	}

	goose.SetBaseFS(migratedb.GooseMigrations)

	if err := goose.SetDialect(driver); err != nil {
		return err
	}

	if err := goose.Up(drv, "migrations-goose"); err != nil {
		return err
	}

	if _, err := drv.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		drv.Close()

		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	return nil
}

// runAtlasMigrations runs the atlas auto-migrations
// this do not use the generated versioned migrations files from ent
func (c *client) runAtlasMigrations(ctx context.Context) error {
	// Run the automatic migration tool to create all schema resources.
	// entcache.Driver will skip the caching layer when running the schema migration
	if err := c.pc.Schema.Create(entcache.Skip(ctx)); err != nil {
		c.logger.Errorf("failed creating schema resources", zap.Error(err))

		return err
	}

	return nil
}

// createEntDBClient creates a new ent client with configured options
func (c *client) createEntDBClient(db *entsql.Driver) *ent.Client {
	cOpts := []ent.Option{ent.Driver(db)}

	if c.config.Debug {
		cOpts = append(cOpts,
			ent.Log(c.logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	return ent.NewClient(cOpts...)
}

// NewTestContainer creates a test container for testing purposes
func NewTestContainer(ctx context.Context) *testutils.TC {
	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")

	return testutils.GetTestURI(ctx, testDBURI)
}

// NewTestClient creates a entdb client that can be used for TEST purposes ONLY
func NewTestClient(ctx context.Context, ctr *testutils.TC, entOpts []ent.Option) (*ent.Client, error) {
	// setup logger
	logger := zap.NewNop().Sugar()

	dbconf := entx.Config{
		Debug:           true,
		DriverName:      ctr.Dialect,
		PrimaryDBSource: ctr.URI,
		EnableHistory:   true,             // enable history so the code path is checked during unit tests
		CacheTTL:        -1 * time.Second, // do not cache results in tests
	}

	entOpts = append(entOpts, ent.Logger(*logger))

	db, _, err := NewMultiDriverDBClient(ctx, dbconf, logger, entOpts)
	if err != nil {
		return nil, err
	}

	if err := db.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
