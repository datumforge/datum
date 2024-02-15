package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	mock_fga "github.com/datumforge/fgax/mockery"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

func TestVerifyHandler(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// add handler
	client.e.GET("verify", client.h.VerifyEmail)

	ec := echocontext.NewTestEchoContext().Request().Context()

	testCases := []struct {
		name            string
		userConfirmed   bool
		email           string
		ttl             string
		tokenSet        bool
		expectedMessage string
		expectedStatus  int
	}{
		{
			name:            "happy path, unconfirmed user",
			userConfirmed:   false,
			email:           "mitb@datum.net",
			tokenSet:        true,
			expectedMessage: "success",
			expectedStatus:  http.StatusOK,
		},
		{
			name:            "happy path, already confirmed user",
			userConfirmed:   true,
			email:           "sitb@datum.net",
			tokenSet:        true,
			expectedMessage: "success",
			expectedStatus:  http.StatusOK,
		},
		{
			name:            "missing token",
			userConfirmed:   true,
			email:           "santa@datum.net",
			tokenSet:        false,
			expectedMessage: "token is required",
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name:            "expired token, but not already confirmed",
			userConfirmed:   false,
			email:           "elf@datum.net",
			tokenSet:        true,
			ttl:             "1987-08-16T03:04:11.169086-07:00",
			expectedMessage: "Token expired, a new token has been issued. Please try again",
			expectedStatus:  http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)

			// set privacy allow in order to allow the creation of the users without
			// authentication in the tests
			ctx := privacy.DecisionContext(ec, privacy.Allow)

			// create user in the database
			userSetting := client.db.UserSetting.Create().
				SetEmailConfirmed(tc.userConfirmed).
				SaveX(ctx)

			// mock writes for user creation
			mock_fga.WriteAny(t, client.fga)

			u := client.db.User.Create().
				SetFirstName(gofakeit.FirstName()).
				SetLastName(gofakeit.LastName()).
				SetEmail(tc.email).
				SetPassword(validPassword).
				SetSetting(userSetting).
				SaveX(ctx)

			user := handlers.User{
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Email:     u.Email,
				ID:        u.ID,
			}

			// create token
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

			// store token in db
			et := client.db.EmailVerificationToken.Create().
				SetOwner(u).
				SetToken(user.EmailVerificationToken.String).
				SetEmail(user.Email).
				SetSecret(user.EmailVerificationSecret).
				SetTTL(ttl).
				SaveX(ctx)

			target := "/verify"
			if tc.tokenSet {
				target = fmt.Sprintf("/verify?token=%s", et.Token)
			}

			req := httptest.NewRequest(http.MethodGet, target, nil)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			client.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			var out *handlers.VerifyReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			if tc.expectedStatus == http.StatusOK {
				assert.True(t, out.Success)
			} else if tc.expectedStatus == http.StatusCreated {
				assert.False(t, out.Success)
			} else {
				assert.Contains(t, out.Error, tc.expectedMessage)
				assert.False(t, out.Success)
			}
		})
	}
}
