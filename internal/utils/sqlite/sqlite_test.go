package sqlite_test

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/utils/sqlite"
)

func TestDriver(t *testing.T) {
	db, err := sql.Open("datum_sqlite3", filepath.Join(t.TempDir(), "test.db"))
	require.NoError(t, err, "could not open connection to testdb")

	conn, err := db.Conn(context.Background())
	require.NoError(t, err, "could not create connection with custom driver")
	require.Equal(t, 1, sqlite.NumConns())

	// Get the underlying sqlite3 connection
	sqlc, ok := sqlite.GetLastConn()
	require.True(t, ok, "connection was not in connection map?")
	require.IsType(t, &sqlite.Conn{}, sqlc, "connection of wrong type returned")

	err = conn.Close()
	require.NoError(t, err, "could not close connection")

	err = db.Close()
	require.NoError(t, err, "could not close database")
	require.Equal(t, 0, sqlite.NumConns())
}

func TestOpenMany(t *testing.T) {
	tmpdir := t.TempDir()
	expectedConnections := 12
	closers := make([]io.Closer, expectedConnections)
	conns := make([]*sqlite.Conn, expectedConnections)

	for i := 0; i < expectedConnections; i++ {
		db, err := sql.Open("datum_sqlite3", filepath.Join(tmpdir, fmt.Sprintf("test-%d.db", i+1)))
		require.NoError(t, err, "could not open connection to database")
		require.NoError(t, db.Ping(), "could not ping database to establish a connection")
		closers[i] = db

		var ok bool
		conns[i], ok = sqlite.GetLastConn()
		require.True(t, ok, "expected new connection")
	}

	// Ensure that we created the expected number of connections
	require.Equal(t, expectedConnections, sqlite.NumConns())
	require.Len(t, closers, expectedConnections)
	require.Len(t, conns, expectedConnections)

	// Should have different connnections
	for i := 1; i < len(conns); i++ {
		require.NotSame(t, conns[i-1], conns[i], "expected connections to be different")
	}

	// Close each connection
	for _, closer := range closers {
		require.NoError(t, closer.Close(), "expected no error during close")
		expectedConnections--
		require.Equal(t, expectedConnections, sqlite.NumConns())
	}
}
