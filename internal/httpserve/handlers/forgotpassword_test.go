package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	echo "github.com/datumforge/echox"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/rShetty/asyncwait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/emails/mock"
)

func (suite *HandlerTestSuite) TestForgotPasswordHandler() {
	t := suite.T()

	// setup handler
	suite.e.POST("forgot-password", suite.h.ForgotPassword)

	ec := echocontext.NewTestEchoContext().Request().Context()

	// create user in the database
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	// add mocks for writes
	mock_fga.WriteAny(t, suite.fga)

	userSetting := suite.db.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ctx)

	_ = suite.db.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("asandler@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ctx)

	var mitb = "mitb@datum.net"

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
			from:           mitb,
			emailExpected:  true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "email does not exist, should still return 200",
			email:          "asandler1@datum.net",
			from:           mitb,
			emailExpected:  false,
			expectedStatus: http.StatusOK,
		},
		{
			name:               "email not sent in request",
			from:               mitb,
			emailExpected:      false,
			expectedStatus:     http.StatusBadRequest,
			expectedErrMessage: "email is required",
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
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus != http.StatusOK {
				var out *handlers.ForgotPasswordReply

				// parse request body
				if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
					t.Error("error parsing response", err)
				}

				assert.Contains(t, out.Error, tc.expectedErrMessage)
				assert.False(t, out.Success)
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
				return suite.h.TaskMan.GetQueueLength() == 0
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
