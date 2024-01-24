package handlers_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"

	echo "github.com/datumforge/echox"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"github.com/datumforge/datum/internal/cookies"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/transaction"
	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/testutils"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/marionette"
)

var (
	dbContainer *testutils.TC

	// commonly used vars in tests
	emptyResponse = "null\n"
	validPassword = "sup3rs3cu7e!"

	// mock email send settings
	maxWaitInMillis      = 2000
	pollIntervalInMillis = 50
)

type client struct {
	e   *echo.Echo
	db  *ent.Client
	h   *handlers.Handler
	fga *mock_fga.MockSdkClient
}

func setupEcho(entClient *ent.Client) *echo.Echo {
	// create echo context with middleware
	e := echo.New()
	transactionConfig := transaction.Client{
		EntDBClient: entClient,
		Logger:      zap.NewNop().Sugar(),
	}

	e.Use(transactionConfig.Middleware)

	return e
}

// handlerSetup to be used for required references in the handler tests
func handlerSetup(t *testing.T, ent *ent.Client) *handlers.Handler {
	tm, err := createTokenManager(15 * time.Minute) //nolint:gomnd
	if err != nil {
		t.Fatal("error creating token manager")
	}

	rc := newRedisClient()

	emConfig := emails.Config{
		Testing:   true,
		Archive:   filepath.Join("fixtures", "emails"),
		FromEmail: "mitb@datum.net",
	}

	em, err := emails.New(emConfig)
	if err != nil {
		t.Fatal("error creating email manager")
	}

	h := &handlers.Handler{
		TM:           tm,
		DBClient:     ent,
		RedisClient:  rc,
		Logger:       zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar(),
		SM:           createSessionManager(),
		EmailManager: em,
	}

	// Start task manager
	tmConfig := marionette.Config{
		Logger: zap.NewNop().Sugar(),
	}

	h.TaskMan = marionette.New(tmConfig)

	h.TaskMan.Start()

	return h
}

func setupTest(t *testing.T) *client {
	ctx := context.Background()

	c := &client{
		fga: mock_fga.NewMockSdkClient(t),
	}

	// create mock FGA client
	fc := fga.NewMockFGAClient(t, c.fga)

	// setup logger
	logger := zap.NewNop().Sugar()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")

	ctr := testutils.GetTestURI(ctx, testDBURI)
	dbContainer = ctr

	dbconf := entdb.Config{
		Debug:           true,
		DriverName:      dbContainer.Dialect,
		PrimaryDBSource: dbContainer.URI,
		CacheTTL:        -1 * time.Second, // do not cache results in tests
	}

	entConfig := entdb.NewDBConfig(dbconf, logger)

	opts := []ent.Option{ent.Logger(*logger), ent.Authz(*fc)}

	db, err := entConfig.NewMultiDriverDBClient(ctx, opts)
	if err != nil {
		require.NoError(t, err, "failed opening connection to database")
	}

	if err := db.Schema.Create(ctx); err != nil {
		require.NoError(t, err, "failed creating database schema")
	}

	// add db to test client
	c.db = db

	// setup handler
	c.h = handlerSetup(t, c.db)

	// setup echo router
	c.e = setupEcho(c.db)

	return c
}

func newRedisClient() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client
}

func createSessionManager() sessions.CookieStore {
	hashKey := randomString(32)  //nolint:gomnd
	blockKey := randomString(32) //nolint:gomnd

	sm := sessions.NewCookieStore(&cookies.DebugOnlyCookieConfig,
		hashKey, blockKey)

	return sm
}

func randomString(n int) []byte {
	id := make([]byte, n)

	if _, err := io.ReadFull(rand.Reader, id); err != nil {
		panic(err) // This shouldn't happen
	}

	return id
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
