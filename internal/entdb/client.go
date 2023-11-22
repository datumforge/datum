package entdb

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
)

// EntClientConfig configures the entsql drivers
type EntClientConfig struct {
	// Debug to print debug database logs
	Debug bool
	// SQL Driver name from dialect.Driver
	DriverName string
	// Logger used for debug logs
	Logger zap.SugaredLogger
	// Primary write database source (required)
	PrimaryDBSource string
	// Secondary write databsae source (optional)
	SecondaryDBSource string
}

func (c *EntClientConfig) newEntDB(dataSource string) (*entsql.Driver, error) {
	// setup db connection
	db, err := sql.Open(c.DriverName, dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to database: %w", err)
	}

	// verify db connection using ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed verifying database connection: %w", err)
	}

	return entsql.OpenDB(dialect.SQLite, db), nil
}

// NewEntDBDriver returns a ent db client
func (c *EntClientConfig) NewEntDBDriver(ctx context.Context) (*ent.Client, error) {
	db, err := c.newEntDB(c.PrimaryDBSource)
	if err != nil {
		return nil, err
	}

	cOpts := []ent.Option{ent.Driver(db)}

	if c.Debug {
		cOpts = append(cOpts,
			ent.Log(c.Logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		c.Logger.Errorf("failed creating schema resources", zap.Error(err))

		return nil, err
	}

	return client, nil
}

// NewMultiDriverDBClient returns a ent client with a primary and secondary write database
func (c *EntClientConfig) NewMultiDriverDBClient(ctx context.Context) (*ent.Client, error) {
	primaryDB, err := c.newEntDB(c.PrimaryDBSource)
	if err != nil {
		return nil, err
	}

	if err := c.createSchema(ctx, primaryDB); err != nil {
		return nil, err
	}

	secondaryDB, err := c.newEntDB(c.SecondaryDBSource)
	if err != nil {
		return nil, err
	}

	if err := c.createSchema(ctx, secondaryDB); err != nil {
		return nil, err
	}

	// Create Multiwrite driver
	cOpts := []ent.Option{ent.Driver(&MultiWriteDriver{Wp: primaryDB, Ws: secondaryDB})}
	if c.Debug {
		cOpts = append(cOpts,
			ent.Log(c.Logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)

	return client, nil
}

func (c *EntClientConfig) createEntDBClient(db *entsql.Driver) *ent.Client {
	cOpts := []ent.Option{ent.Driver(db)}

	if c.Debug {
		cOpts = append(cOpts,
			ent.Log(c.Logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	return ent.NewClient(cOpts...)
}

func (c *EntClientConfig) createSchema(ctx context.Context, db *entsql.Driver) error {
	client := c.createEntDBClient(db)

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		c.Logger.Errorf("failed creating schema resources", zap.Error(err))

		return err
	}

	return nil
}
