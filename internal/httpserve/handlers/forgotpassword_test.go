package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	mock_fga "github.com/datumforge/fgax/mockery"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/rShetty/asyncwait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/emails/mock"
)

func TestForgotPasswordHandler(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup handler
	client.e.POST("forgot-password", client.h.ForgotPassword)

	ec := echocontext.NewTestEchoContext().Request().Context()

	// create user in the database
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	// add mocks for writes
	mock_fga.WriteAny(t, client.fga)

	userSetting := client.db.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ctx)

	_ = client.db.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("asandler@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ctx)

	testCases := []struct {
		name               string
		from               string
		email              string
		emailExpected      bool
		expectedErrMessage string
		expectedStatus     int
	}{
		{
			name:           "happy path",
			email:          "asandler@datum.net",
			from:           "mitb@datum.net",
			emailExpected:  true,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "email does not exist, should still return 204",
			email:          "asandler1@datum.net",
			from:           "mitb@datum.net",
			emailExpected:  false,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "email not sent in request",
			from:           "mitb@datum.net",
			emailExpected:  false,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sent := time.Now()

			mock.ResetEmailMock()

			resendJSON := handlers.ForgotPasswordRequest{
				Email: tc.email,
			}

			body, err := json.Marshal(resendJSON)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/forgot-password", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			client.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus != http.StatusNoContent {
				var out *handlers.Response

				// parse request body
				if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
					t.Error("error parsing response", err)
				}

				assert.Contains(t, out.Message, tc.expectedErrMessage)
			}

			// Test that one verify email was sent to each user
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      tc.from,
					Subject:   emails.PasswordResetRequestRE,
					Timestamp: sent,
				},
			}

			// wait for messages
			predicate := func() bool {
				return client.h.TaskMan.GetQueueLength() == 0
			}
			successful := asyncwait.NewAsyncWait(maxWaitInMillis, pollIntervalInMillis).Check(predicate)

			if successful != true {
				t.Errorf("max wait of email send")
			}

			if tc.emailExpected {
				mock.CheckEmails(t, messages)
			} else {
				mock.CheckEmails(t, nil)
			}
		})
	}
}
