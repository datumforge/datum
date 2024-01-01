package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
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
	h := handlerSetup(t)

	ec := echocontext.NewTestEchoContext().Request().Context()

	// create user in the database
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	userSetting := EntClient.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ctx)

	u := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("asandler@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ctx)

	testCases := []struct {
		name               string
		email              string
		emailExpected      bool
		expectedErrMessage string
		expectedStatus     int
	}{
		{
			name:           "happy path",
			email:          "asandler@datum.net",
			emailExpected:  true,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "email does not exist, should still return 204",
			email:          "asandler1@datum.net",
			emailExpected:  false,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "email not sent in request",
			emailExpected:  false,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sent := time.Now()
			mock.ResetEmailMock()

			e := setupEcho(h.SM)
			e.POST("forgot-password", h.ForgotPassword)

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

				assert.Contains(t, out.Message, tc.expectedErrMessage)
			}

			// Test that one verify email was sent to each user
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      h.SendGridConfig.FromEmail,
					Subject:   emails.PasswordResetRequestRE,
					Timestamp: sent,
				},
			}

			// wait for messages
			predicate := func() bool {
				return h.TaskMan.GetQueueLength() == 0
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

	// cleanup after
	EntClient.User.DeleteOneID(u.ID).ExecX(ctx)
}
