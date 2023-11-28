package entdb

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
)

// MultiWriteDriver allows you to write to a primary and secondary database
type MultiWriteDriver struct {
	// Wp (write-primary), Ws (write-secondary) Drivers
	Wp, Ws dialect.Driver
}

var _ dialect.Driver = (*MultiWriteDriver)(nil)

// Query will query the primary write database
func (d *MultiWriteDriver) Query(ctx context.Context, query string, args, v any) error {
	return d.Wp.Query(ctx, query, args, v)
}

// Exec logs its params and calls the underlying driver Exec method for both write drivers
func (d *MultiWriteDriver) Exec(ctx context.Context, query string, args, v any) error {
	err := d.Ws.Exec(ctx, query, args, v)
	if err != nil {
		return err
	}

	return d.Wp.Exec(ctx, query, args, v)
}

// Tx wraps the Exec and Query operations in transaction.
func (d *MultiWriteDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.Wp.Tx(ctx)
}

// BeginTx adds an log-id for the transaction and calls the underlying driver BeginTx command if it is supported.
func (d *MultiWriteDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return d.Wp.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

// Close the underlying connections
func (d *MultiWriteDriver) Close() error {
	wserr := d.Ws.Close()
	wperr := d.Wp.Close()

	if wperr != nil {
		return wserr
	}

	if wserr != nil {
		return wserr
	}

	return nil
}

// Dialect returns the dialect name of the primary driver
func (d *MultiWriteDriver) Dialect() string {
	return d.Wp.Dialect()
}
