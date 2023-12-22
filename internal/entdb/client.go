package entdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/internal/httpserve/config"
)

const (
	DefaultCacheTTL = 1 * time.Second
)

// EntClientConfig configures the entsql drivers
type EntClientConfig struct {
	// config contains the base database settings
	config config.DB
	// primaryDB contains the primary db connection
	primaryDB *entsql.Driver
	// secondaryDB contains the secondary db connection, if set
	secondaryDB *entsql.Driver
	// logger contains the zap logger
	logger *zap.SugaredLogger
}

// NewDBConfig returns a new database configuration
func NewDBConfig(c config.DB, l *zap.SugaredLogger) *EntClientConfig {
	return &EntClientConfig{
		config: c,
		logger: l,
	}
}

func (c *EntClientConfig) GetPrimaryDB() *entsql.Driver {
	return c.primaryDB
}

func (c *EntClientConfig) GetSecondaryDB() *entsql.Driver {
	return c.secondaryDB
}

func (c *EntClientConfig) newEntDB(dataSource string) (*entsql.Driver, error) {
	// setup db connection
	db, err := sql.Open(c.config.DriverName, dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to database: %w", err)
	}

	// verify db connection using ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed verifying database connection: %w", err)
	}

	return entsql.OpenDB(dialect.SQLite, db), nil
}

// NewMultiDriverDBClient returns a ent client with a primary and secondary, if configured, write database
func (c *EntClientConfig) NewMultiDriverDBClient(ctx context.Context, opts []ent.Option) (*ent.Client, error) {
	var err error

	c.primaryDB, err = c.newEntDB(c.config.PrimaryDBSource)
	if err != nil {
		return nil, err
	}

	// Decorates the sql.Driver with entcache.Driver on the primaryDB
	drvPrimary := entcache.NewDriver(
		c.primaryDB,
		entcache.TTL(c.config.CacheTTL), // set the TTL on the cache
	)

	if err := c.createSchema(ctx, c.primaryDB); err != nil {
		c.logger.Errorf("failed creating schema resources", zap.Error(err))

		return nil, err
	}

	var cOpts []ent.Option

	if c.config.MultiWrite {
		c.secondaryDB, err = c.newEntDB(c.config.SecondaryDBSource)
		if err != nil {
			return nil, err
		}

		// Decorates the sql.Driver with entcache.Driver on the primaryDB
		drvSecondary := entcache.NewDriver(
			c.secondaryDB,
			entcache.TTL(c.config.CacheTTL), // set the TTL on the cache
		)

		if err := c.createSchema(ctx, c.secondaryDB); err != nil {
			c.logger.Errorf("failed creating schema resources", zap.Error(err))

			return nil, err
		}

		// Create Multiwrite driver
		cOpts = []ent.Option{ent.Driver(&MultiWriteDriver{Wp: drvPrimary, Ws: drvSecondary})}
	} else {
		cOpts = []ent.Option{ent.Driver(drvPrimary)}
	}

	cOpts = append(cOpts, opts...)

	if c.config.Debug {
		cOpts = append(cOpts,
			ent.Log(c.logger.Named("ent").Debugln),
			ent.Debug(),
			ent.Driver(drvPrimary),
		)
	}

	client := ent.NewClient(cOpts...)

	client.Intercept(interceptors.QueryLogger(c.logger))

	return client, nil
}

func (c *EntClientConfig) createEntDBClient(db *entsql.Driver) *ent.Client {
	cOpts := []ent.Option{ent.Driver(db)}

	if c.config.Debug {
		cOpts = append(cOpts,
			ent.Log(c.logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	return ent.NewClient(cOpts...)
}

func (c *EntClientConfig) createSchema(ctx context.Context, db *entsql.Driver) error {
	client := c.createEntDBClient(db)

	// Run the automatic migration tool to create all schema resources.
	// entcache.Driver will skip the caching layer when running the schema migration
	if err := client.Schema.Create(entcache.Skip(ctx)); err != nil {
		c.logger.Errorf("failed creating schema resources", zap.Error(err))

		return err
	}

	return nil
}

// Healthcheck pings the DB to check if the connection is working
func Healthcheck(client *entsql.Driver) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		if err := client.DB().Ping(); err != nil {
			return errors.Wrap(err, "db connection failed")
		}

		return nil
	}
}
