package graphapi_test

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/datumforge/fgax"
	mock_fga "github.com/datumforge/fgax/mockery"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/datumforge/datum/internal/ent/entconfig"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/testutils"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/marionette"
	"github.com/datumforge/datum/pkg/utils/totp"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

var (
	testUser          *ent.User
	testPersonalOrgID string
	testOrgID         string
)

// TestGraphTestSuite runs all the tests in the GraphTestSuite
func TestGraphTestSuite(t *testing.T) {
	suite.Run(t, new(GraphTestSuite))
}

// GraphTestSuite handles the setup and teardown between tests
type GraphTestSuite struct {
	suite.Suite
	client *client
	tf     *testutils.TestFixture
}

// client contains all the clients the test need to interact with
type client struct {
	db                *ent.Client
	datum             *datumclient.DatumClient
	datumWithPAT      *datumclient.DatumClient
	datumWithAPIToken *datumclient.DatumClient
	fga               *mock_fga.MockSdkClient
}

func (suite *GraphTestSuite) SetupSuite() {
	suite.tf = entdb.NewTestFixture()
}

func (suite *GraphTestSuite) SetupTest() {
	t := suite.T()

	ctx := context.Background()

	// setup fga mock
	c := &client{
		fga: mock_fga.NewMockSdkClient(t),
	}

	// create mock FGA client
	fc := fgax.NewMockFGAClient(t, c.fga)

	// setup logger
	logger := zap.NewNop().Sugar()

	// setup email manager
	emConfig := emails.Config{
		Testing:   true,
		Archive:   filepath.Join("fixtures", "emails"),
		FromEmail: "mitb@datum.net",
	}

	em, err := emails.New(emConfig)
	if err != nil {
		t.Fatal("error creating email manager")
	}

	// setup task manager
	tmConfig := marionette.Config{
		Logger: zap.NewNop().Sugar(),
	}

	taskMan := marionette.New(tmConfig)

	taskMan.Start()

	// setup otp manager
	otpOpts := []totp.ConfigOption{
		totp.WithCodeLength(6),
		totp.WithIssuer("datum"),
		totp.WithSecret(totp.Secret{
			Version: 0,
			Key:     ulids.New().String(),
		}),
	}

	tm, err := testutils.CreateTokenManager(15 * time.Minute) //nolint:mnd
	if err != nil {
		t.Fatal("error creating token manager")
	}

	sm := testutils.CreateSessionManager()
	rc := testutils.NewRedisClient()

	sessionConfig := sessions.NewSessionConfig(
		sm,
		sessions.WithPersistence(rc),
		sessions.WithLogger(logger),
	)

	sessionConfig.CookieConfig = &sessions.DebugOnlyCookieConfig

	otpMan := totp.NewOTP(otpOpts...)

	opts := []ent.Option{
		ent.Logger(*logger),
		ent.Authz(*fc),
		ent.Emails(em),
		ent.Marionette(taskMan),
		ent.Analytics(&analytics.EventManager{Enabled: false}),
		ent.TOTP(&totp.Manager{
			TOTPManager: otpMan,
		}),
		ent.TokenManager(tm),
		ent.SessionConfig(&sessionConfig),
		ent.EntConfig(&entconfig.Config{
			Flags: entconfig.Flags{
				UseListUserService:   false,
				UseListObjectService: false,
			},
		}),
	}

	// create database connection
	db, err := entdb.NewTestClient(ctx, suite.tf, opts)
	require.NoError(t, err, "failed opening connection to database")

	// assign values
	c.db = db
	c.datum, err = testutils.DatumTestClient(t, c.db)
	require.NoError(t, err)

	// create test user
	ctx = echocontext.NewTestContext()
	testUser = (&UserBuilder{client: c}).MustNew(ctx, t)
	testPersonalOrg, err := testUser.Edges.Setting.DefaultOrg(ctx)
	require.NoError(t, err)

	testPersonalOrgID = testPersonalOrg.ID

	userCtx, err := auth.NewTestContextWithOrgID(testUser.ID, testPersonalOrgID)
	require.NoError(t, err)

	// setup a non-personal org
	testOrg := (&OrganizationBuilder{client: c}).MustNew(userCtx, t)
	testOrgID = testOrg.ID

	userCtx, err = userContext()
	require.NoError(t, err)

	// setup client with a personal access token
	pat := (&PersonalAccessTokenBuilder{client: c, OwnerID: testUser.ID, OrganizationIDs: []string{testOrgID, testPersonalOrgID}}).MustNew(userCtx, t)
	authHeaderPAT := datumclient.Authorization{
		BearerToken: pat.Token,
	}

	c.datumWithPAT, err = testutils.DatumTestClientWithAuth(t, c.db, datumclient.WithCredentials(authHeaderPAT))
	require.NoError(t, err)

	// setup client with an API token
	apiToken := (&APITokenBuilder{client: c}).MustNew(userCtx, t)

	authHeaderAPIToken := datumclient.Authorization{
		BearerToken: apiToken.Token,
	}
	c.datumWithAPIToken, err = testutils.DatumTestClientWithAuth(t, c.db, datumclient.WithCredentials(authHeaderAPIToken))
	require.NoError(t, err)

	suite.client = c
}

func (suite *GraphTestSuite) TearDownTest() {
	// clear all fga mocks
	mock_fga.ClearMocks(suite.client.fga)

	if suite.client.db != nil {
		if err := suite.client.db.Close(); err != nil {
			log.Fatalf("failed to close database: %s", err)
		}
	}
}

func (suite *GraphTestSuite) TearDownSuite() {
	testutils.TeardownFixture(suite.tf)
}

// userContext creates a new user in the database and returns a context with
// the user claims attached
func userContext() (context.Context, error) {
	return auth.NewTestContextWithOrgID(testUser.ID, testOrgID)
}

// userContextWithID creates a new user context with the provided user ID
func userContextWithID(userID string) (context.Context, error) {
	// Use that user to create the organization
	ec, err := auth.NewTestEchoContextWithValidUser(userID)
	if err != nil {
		return nil, err
	}

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	return reqCtx, nil
}

func (suite *GraphTestSuite) captureOutput(funcToRun func()) string {
	var buffer bytes.Buffer

	oldLogger := suite.client.db.Logger
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buffer)

	logger := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel)).
		Sugar()

	suite.client.db.Logger = *logger

	funcToRun()
	writer.Flush()

	suite.client.db.Logger = oldLogger

	return buffer.String()
}
