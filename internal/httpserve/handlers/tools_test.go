package handlers_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"log"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	echo "github.com/datumforge/echox"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/middleware/transaction"
	"github.com/datumforge/datum/internal/tokens"
)

var (
	defaultDBURI = "file:ent?mode=memory&cache=shared&_fk=1"
	EntClient    *ent.Client
)

func TestMain(m *testing.M) {
	// setup the database if needed
	setupDB()
	// run the tests
	code := m.Run()
	// teardown the database
	teardownDB()
	// return the test response code
	os.Exit(code)
}

func setupEcho() *echo.Echo {
	// create echo context with middleware
	e := echo.New()
	transactionConfig := transaction.Client{
		EntDBClient: EntClient,
		Logger:      zap.NewNop().Sugar(),
	}

	e.Use(transactionConfig.Middleware)

	return e
}

func setupDB() {
	// don't setup the datastore if we already have one
	if EntClient != nil {
		return
	}

	logger := zap.NewNop().Sugar()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")
	if testDBURI == "" {
		testDBURI = defaultDBURI
	}

	dbconf := config.DB{
		Debug:           true,
		DriverName:      dialect.SQLite,
		PrimaryDBSource: testDBURI,
	}

	entConfig := entdb.NewDBConfig(dbconf, logger)

	ctx := context.Background()

	opts := []ent.Option{ent.Logger(*logger)}

	c, err := entConfig.NewMultiDriverDBClient(ctx, opts)
	if err != nil {
		errPanic("failed opening connection to database:", err)
	}

	errPanic("failed creating db schema", c.Schema.Create(ctx))

	EntClient = c
}

func teardownDB() {
	if EntClient != nil {
		errPanic("teardown failed to close database connection", EntClient.Close())
	}
}

func errPanic(msg string, err error) {
	if err != nil {
		log.Panicf("%s err: %s", msg, err.Error())
	}
}

func createTokenManager(refreshOverlap time.Duration) (*tokens.TokenManager, error) {
	conf := tokens.Config{
		Audience:        "http://localhost:17608",
		Issuer:          "http://localhost:17608",
		AccessDuration:  1 * time.Hour, // nolint: gomnd
		RefreshDuration: 2 * time.Hour, // nolint: gomnd
		RefreshOverlap:  refreshOverlap,
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048) // nolint: gomnd
	if err != nil {
		return nil, err
	}

	return tokens.NewWithKey(key, conf)
}
