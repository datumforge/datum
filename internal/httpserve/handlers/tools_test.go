package handlers_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"

	echo "github.com/datumforge/echox"
	ofgaclient "github.com/openfga/go-sdk/client"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"github.com/datumforge/datum/internal/cookies"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/transaction"
	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/testutils"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/marionette"
)

var (
	EntClient   *ent.Client
	DBContainer *testutils.TC

	// commonly used vars in tests
	emptyResponse = "null\n"
	validPassword = "sup3rs3cu7e!"

	// mock email send settings
	maxWaitInMillis      = 2000
	pollIntervalInMillis = 50
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

func setupEchoAuth(entClient *ent.Client) *echo.Echo {
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

	h := &handlers.Handler{
		TM:          tm,
		DBClient:    ent,
		RedisClient: rc,
		Logger:      zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)).Sugar(),
		SM:          createSessionManager(),
	}

	if err := h.NewTestEmailManager(); err != nil {
		t.Fatalf("error creating email manager: %v", err)
	}

	// Start task manager
	tmConfig := marionette.Config{
		Logger: zap.NewNop().Sugar(),
	}

	h.TaskMan = marionette.New(tmConfig)

	h.TaskMan.Start()

	return h
}

func setupDB() {
	ctx := context.Background()

	// don't setup the datastore if we already have one
	if EntClient != nil {
		return
	}

	logger := zap.NewNop().Sugar()

	// Grab the DB environment variable or use the default
	testDBURI := os.Getenv("TEST_DB_URL")

	ctr := testutils.GetTestURI(ctx, testDBURI)
	DBContainer = ctr

	dbconf := entdb.Config{
		Debug:           true,
		DriverName:      ctr.Dialect,
		PrimaryDBSource: ctr.URI,
	}

	entConfig := entdb.NewDBConfig(dbconf, logger)

	opts := []ent.Option{ent.Logger(*logger)}

	c, err := entConfig.NewMultiDriverDBClient(ctx, opts)
	if err != nil {
		errPanic("failed opening connection to database:", err)
	}

	errPanic("failed creating db schema", c.Schema.Create(ctx))

	EntClient = c
}

func setupAuthEntDB(t *testing.T, mockCtrl *gomock.Controller, mc *mock_client.MockSdkClient) *ent.Client {
	fc, err := fga.NewTestFGAClient(t, mockCtrl, mc)
	if err != nil {
		t.Fatalf("enable to create test FGA client")
	}

	logger := zap.NewNop().Sugar()

	if DBContainer == nil {
		t.Fatalf("DBContainer is nil")
	}

	dbconf := entdb.Config{
		Debug:           true,
		DriverName:      DBContainer.Dialect,
		PrimaryDBSource: DBContainer.URI,
		CacheTTL:        -1 * time.Second, // do not cache results in tests
	}

	entConfig := entdb.NewDBConfig(dbconf, logger)

	ctx := context.Background()

	opts := []ent.Option{ent.Logger(*logger), ent.Authz(*fc)}

	c, err := entConfig.NewMultiDriverDBClient(ctx, opts)
	if err != nil {
		errPanic("failed opening connection to database:", err)
	}

	errPanic("failed creating db schema", c.Schema.Create(ctx))

	return c
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

// mockWriteAny creates mock responses based on the mock FGA client
func mockWriteAny(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, errMsg error) {
	mockExecute := mock_client.NewMockSdkClientWriteRequestInterface(mockCtrl)

	if errMsg == nil {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteRequestWriteResponse{
				{
					Status: ofgaclient.SUCCESS,
				},
			},
			Deletes: []ofgaclient.ClientWriteRequestDeleteResponse{
				{
					Status: ofgaclient.SUCCESS,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, nil)
	} else {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteRequestWriteResponse{
				{
					Status: ofgaclient.FAILURE,
				},
			},
			Deletes: []ofgaclient.ClientWriteRequestDeleteResponse{
				{
					Status: ofgaclient.FAILURE,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, errMsg)
	}

	mockRequest := mock_client.NewMockSdkClientWriteRequestInterface(mockCtrl)

	mockRequest.EXPECT().Options(gomock.Any()).Return(mockExecute)

	mockBody := mock_client.NewMockSdkClientWriteRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(gomock.Any()).Return(mockRequest)

	c.EXPECT().Write(gomock.Any()).Return(mockBody)
}
