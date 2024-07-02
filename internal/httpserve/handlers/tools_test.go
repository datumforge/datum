package handlers_test

import (
	"context"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/fgax"
	mock_fga "github.com/datumforge/fgax/mockery"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/httpserve/authmanager"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/middleware/transaction"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/testutils"
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
	e     *echo.Echo
	db    *ent.Client
	datum *datumclient.DatumClient
	h     *handlers.Handler
	fga   *mock_fga.MockSdkClient
	tf    *testutils.TestFixture
}

// TestHandlerTestSuite runs all the tests in the HandlerTestSuite
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) SetupSuite() {
	suite.tf = entdb.NewTestFixture()
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

	// create database connection
	db, err := entdb.NewTestClient(ctx, suite.tf, opts)
	require.NoError(t, err, "failed opening connection to database")

	// add db to test client
	suite.db = db

	// add the datum client
	suite.datum, err = testutils.DatumTestClient(t, suite.db)
	require.NoError(t, err)

	// setup handler
	suite.h = handlerSetup(t, suite.db, em, taskMan)

	// setup echo router
	suite.e = setupEcho(suite.db)
}

func (suite *HandlerTestSuite) TearDownTest() {
	// clear all fga mocks
	mock_fga.ClearMocks(suite.fga)

	if suite.db != nil {
		if err := suite.db.Close(); err != nil {
			log.Fatalf("failed to close database: %s", err)
		}
	}
}

func (suite *HandlerTestSuite) TearDownSuite() {
	testutils.TeardownFixture(suite.tf)
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
	tm, err := testutils.CreateTokenManager(15 * time.Minute) //nolint:mnd
	if err != nil {
		t.Fatal("error creating token manager")
	}

	sm := testutils.CreateSessionManager()
	rc := testutils.NewRedisClient()
	logger := zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar()

	sessionConfig := sessions.NewSessionConfig(
		sm,
		sessions.WithPersistence(rc),
		sessions.WithLogger(logger),
	)

	sessionConfig.CookieConfig = &sessions.DebugOnlyCookieConfig

	as := authmanager.New()
	as.SetTokenManager(tm)
	as.SetSessionConfig(&sessionConfig)

	h := &handlers.Handler{
		IsTest:        true,
		TokenManager:  tm,
		DBClient:      ent,
		RedisClient:   rc,
		Logger:        logger,
		SessionConfig: &sessionConfig,
		AuthManager:   as,
		EmailManager:  em,
		TaskMan:       taskMan,
		AnalyticsClient: &analytics.EventManager{
			Enabled: false,
		},
	}

	return h
}
