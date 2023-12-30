package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/utils/marionette"
)

var (
	emptyResponse = "null\n"
	validPassword = "sup3rs3cu7e!"
)

// handlerSetup to be used for required references in the handler tests
func handlerSetup(t *testing.T) *handlers.Handler {
	tm, err := createTokenManager(15 * time.Minute) //nolint:gomnd
	if err != nil {
		t.Fatal("error creating token manager")
	}

	h := &handlers.Handler{
		TM:           tm,
		DBClient:     EntClient,
		Logger:       zap.NewNop().Sugar(),
		CookieDomain: "datum.net",
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

func TestVerifyHandler(t *testing.T) {
	h := handlerSetup(t)

	ec := echocontext.NewTestEchoContext().Request().Context()

	testCases := []struct {
		name           string
		userConfirmed  bool
		email          string
		ttl            string
		tokenSet       bool
		expectedResp   string
		expectedStatus int
	}{
		{
			name:           "happy path, unconfirmed user",
			userConfirmed:  false,
			email:          "mitb@datum.net",
			tokenSet:       true,
			expectedResp:   emptyResponse,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "happy path, already confirmed user",
			userConfirmed:  true,
			email:          "sitb@datum.net",
			tokenSet:       true,
			expectedResp:   emptyResponse,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "missing token",
			userConfirmed:  true,
			email:          "santa@datum.net",
			tokenSet:       false,
			expectedResp:   "token is required",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "expired token, but not already confirmed",
			userConfirmed:  false,
			email:          "elf@datum.net",
			tokenSet:       true,
			ttl:            "1987-08-16T03:04:11.169086-07:00",
			expectedResp:   "Token expired, a new token has been issued. Please try again",
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create echo context with middleware
			e := setupEcho()
			e.GET("verify", h.VerifyEmail)

			// create user in the database
			userSetting := EntClient.UserSetting.Create().
				SetEmailConfirmed(tc.userConfirmed).
				SaveX(ec)

			u := EntClient.User.Create().
				SetFirstName(gofakeit.FirstName()).
				SetLastName(gofakeit.LastName()).
				SetEmail(tc.email).
				SetPassword(validPassword).
				SetSetting(userSetting).
				SaveX(ec)

			user := handlers.User{
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Email:     u.Email,
				ID:        u.ID,
			}

			if err := user.CreateVerificationToken(); err != nil {
				require.NoError(t, err)
			}

			if tc.ttl != "" {
				user.EmailVerificationExpires.String = tc.ttl
			}

			ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
			if err != nil {
				require.NoError(t, err)
			}

			et := EntClient.EmailVerificationToken.Create().
				SetOwner(u).
				SetToken(user.EmailVerificationToken.String).
				SetEmail(user.Email).
				SetSecret(user.EmailVerificationSecret).
				SetTTL(ttl).
				SaveX(ec)

			target := "/verify"
			if tc.tokenSet {
				target = fmt.Sprintf("/verify?token=%s", et.Token)
			}

			req := httptest.NewRequest(http.MethodGet, target, nil)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus != http.StatusNoContent {
				var out *handlers.Response

				// parse request body
				if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
					t.Error("error parsing response", err)
				}

				assert.Contains(t, out.Message, tc.expectedResp)
			}

			// cleanup after
			EntClient.User.DeleteOneID(u.ID).ExecX(ec)
		})
	}
}
