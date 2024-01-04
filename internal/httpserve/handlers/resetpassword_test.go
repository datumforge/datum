package handlers_test

import (
	"context"
	"encoding/json"
	"fmt"
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

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/emails/mock"
)

func TestResetPassword(t *testing.T) {
	h := handlerSetup(t)

	ec := echocontext.NewTestEchoContext().Request().Context()

	testCases := []struct {
		name                 string
		email                string
		newPassword          string
		tokenSet             bool
		tokenProvided        string
		ttl                  string
		emailExpected        bool
		expectedEmailSubject string
		expectedResp         string
		expectedStatus       int
	}{
		{
			name:                 "happy path",
			email:                "kelsier@datum.net",
			tokenSet:             true,
			newPassword:          "6z9Fqc-E-9v32NsJzLNU",
			emailExpected:        true,
			expectedEmailSubject: emails.PasswordResetSuccessRE,
			expectedResp:         emptyResponse,
			expectedStatus:       http.StatusNoContent,
		},
		{
			name:           "bad token (user not found)",
			email:          "eventure@datum.net",
			tokenSet:       true,
			tokenProvided:  "thisisnotavalidtoken",
			newPassword:    "6z9Fqc-E-9v32NsJzLNU",
			emailExpected:  false,
			expectedResp:   "password reset token invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "weak password",
			email:          "sazed@datum.net",
			tokenSet:       true,
			newPassword:    "weak1",
			emailExpected:  false,
			expectedResp:   "password is too weak",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "same password",
			email:          "sventure@datum.net",
			tokenSet:       true,
			newPassword:    validPassword,
			emailExpected:  false,
			expectedResp:   "password was already used",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing token",
			email:          "dockson@datum.net",
			tokenSet:       false,
			newPassword:    "6z9Fqc-E-9v32NsJzLNU",
			emailExpected:  false,
			expectedResp:   "token is required",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "expired reset token",
			email:          "tensoon@datum.net",
			newPassword:    "6z9Fqc-E-9v32NsJzLNP",
			tokenSet:       true,
			emailExpected:  false,
			ttl:            "1987-08-16T03:04:11.169086-07:00",
			expectedResp:   "reset token is expired, please request a new token using forgot-password",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sent := time.Now()
			mock.ResetEmailMock()

			// create echo context with middleware
			e := setupEcho(h.SM)

			// create user in the database
			rt, userID, err := createUserWithResetToken(ec, tc.email, tc.ttl)
			require.NoError(t, err)

			// setup request request
			e.POST("reset-password", h.ResetPassword)

			pwResetJSON := handlers.ResetPasswordRequest{
				Password: tc.newPassword,
			}

			body, err := json.Marshal(pwResetJSON)
			if err != nil {
				require.NoError(t, err)
			}

			target := "/reset-password"
			if tc.tokenSet {
				token := rt.Token
				if tc.tokenProvided != "" {
					token = tc.tokenProvided
				}

				target = fmt.Sprintf("%s?token=%s", target, token)
			}

			req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			e.ServeHTTP(recorder, req)

			// get result
			res := recorder.Result()
			defer res.Body.Close()

			// check status
			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus != http.StatusNoContent {
				var out *handlers.Response

				// parse request body
				if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
					t.Error("error parsing response", err)
				}

				assert.Contains(t, out.Message, tc.expectedResp)
			}

			// Test that one verify email was sent to each user
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      h.SendGridConfig.FromEmail,
					Subject:   tc.expectedEmailSubject,
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

			// cleanup after
			ctx := privacy.DecisionContext(ec, privacy.Allow)
			EntClient.User.DeleteOneID(userID).ExecX(ctx)
		})
	}
}

// createUserWithResetToken creates a user with a valid reset token and returns the token, user id, and error if one occurred
func createUserWithResetToken(ec context.Context, email string, ttl string) (*ent.PasswordResetToken, string, error) {
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	userSetting := EntClient.UserSetting.Create().
		SetEmailConfirmed(true).
		SaveX(ctx)

	u := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail(email).
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
	if err := user.CreatePasswordResetToken(); err != nil {
		return nil, "", err
	}

	// set expiry if provided in test case
	if ttl != "" {
		user.PasswordResetExpires.String = ttl
	}

	exp, err := time.Parse(time.RFC3339Nano, user.PasswordResetExpires.String)
	if err != nil {
		return nil, "", err
	}

	// store token in db
	return EntClient.PasswordResetToken.Create().
		SetOwner(u).
		SetToken(user.PasswordResetToken.String).
		SetEmail(user.Email).
		SetSecret(user.PasswordResetSecret).
		SetTTL(exp).
		SaveX(ec), user.ID, nil
}
