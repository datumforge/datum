package handlers_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/fgax"
	mock_fga "github.com/datumforge/fgax/mockery"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/middleware/transaction"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/testutils"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/marionette"
)

var (
	// commonly used vars in tests
	emptyResponse = "null\n"
	validPassword = "sup3rs3cu7e!"

	// mock email send settings
	maxWaitInMillis      = 2000
	pollIntervalInMillis = 100
)

// HandlerTestSuite handles the setup and teardown between tests
type HandlerTestSuite struct {
	suite.Suite
	e   *echo.Echo
	db  *ent.Client
	h   *handlers.Handler
	fga *mock_fga.MockSdkClient
	tc  *testutils.TC
}

// TestHandlerTestSuite runs all the tests in the HandlerTestSuite
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) SetupSuite() {
	ctx := context.Background()

	suite.tc = entdb.NewTestContainer(ctx)
}

func (suite *HandlerTestSuite) SetupTest() {
	t := suite.T()

	ctx := context.Background()

	suite.fga = mock_fga.NewMockSdkClient(t)

	// create mock FGA client
	fc := fgax.NewMockFGAClient(t, suite.fga)

	// setup logger
	logger := zap.NewNop().Sugar()

	emConfig := emails.Config{
		Testing:   true,
		Archive:   filepath.Join("fixtures", "emails"),
		FromEmail: "mitb@datum.net",
	}

	em, err := emails.New(emConfig)
	if err != nil {
		t.Fatal("error creating email manager")
	}

	em.MarketingURLConfig = emails.MarketingURLConfig{
		DefaultSubscriptionOrg: "MITB",
	}

	// Start task manager
	tmConfig := marionette.Config{
		Logger: zap.NewNop().Sugar(),
	}

	taskMan := marionette.New(tmConfig)

	taskMan.Start()

	opts := []ent.Option{
		ent.Logger(*logger),
		ent.Authz(*fc),
		ent.Marionette(taskMan),
		ent.Emails(em),
		ent.Analytics(&analytics.EventManager{Enabled: false}),
	}

	ctr := entdb.NewTestContainer(ctx)

	// create database connection
	db, err := entdb.NewTestClient(ctx, ctr, opts)
	if err != nil {
		require.NoError(t, err, "failed opening connection to database")
	}

	// add db to test client
	suite.db = db

	// setup handler
	suite.h = handlerSetup(t, suite.db, em, taskMan)

	// setup echo router
	suite.e = setupEcho(suite.db)
}

func (suite *HandlerTestSuite) TearDownTest() {
	// clear all fga mocks
	mock_fga.ClearMocks(suite.fga)

	if err := suite.db.Close(); err != nil {
		log.Fatalf("failed to close database: %s", err)
	}
}

func (suite *HandlerTestSuite) TearDownSuite() {
	if suite.tc.Container != nil {
		if err := suite.tc.Container.Terminate(context.Background()); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}
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
func handlerSetup(t *testing.T, ent *ent.Client, em *emails.EmailManager, taskMan *marionette.TaskManager) *handlers.Handler {
	tm, err := createTokenManager(15 * time.Minute) //nolint:gomnd
	if err != nil {
		t.Fatal("error creating token manager")
	}

	sm := createSessionManager()
	rc := newRedisClient()
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	sessionConfig := sessions.NewSessionConfig(
		sm,
		sessions.WithPersistence(rc),
		sessions.WithLogger(logger),
	)

	sessionConfig.CookieConfig = &sessions.DebugOnlyCookieConfig

	h := &handlers.Handler{
		IsTest:        true,
		TM:            tm,
		DBClient:      ent,
		RedisClient:   rc,
		Logger:        logger,
		SessionConfig: &sessionConfig,
		EmailManager:  em,
		TaskMan:       taskMan,
		AnalyticsClient: &analytics.EventManager{
			Enabled: false,
		},
	}

	return h
}

func newRedisClient() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:             mr.Addr(),
		DisableIndentity: true,
	})

	return client
}

func createSessionManager() sessions.Store[map[string]any] {
	hashKey := randomString(32)  //nolint:gomnd
	blockKey := randomString(32) //nolint:gomnd

	sm := sessions.NewCookieStore[map[string]any](sessions.DebugCookieConfig,
		hashKey, blockKey,
	)

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
