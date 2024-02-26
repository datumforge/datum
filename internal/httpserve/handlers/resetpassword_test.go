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

	"github.com/brianvoe/gofakeit/v7"
	mock_fga "github.com/datumforge/fgax/mockery"
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
	client := setupTest(t)
	defer client.db.Close()

	// setup request request
	client.e.POST("password-reset", client.h.ResetPassword)

	ec := echocontext.NewTestEchoContext().Request().Context()

	var newPassword = "6z9Fqc-E-9v32NsJzLNU" //nolint:gosec

	expiredTTL := time.Now().AddDate(0, 0, -1).Format(time.RFC3339Nano)

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
		from                 string
	}{
		{
			name:                 "happy path",
			email:                "kelsier@datum.net",
			tokenSet:             true,
			newPassword:          newPassword,
			from:                 "mitb@datum.net",
			emailExpected:        true,
			expectedEmailSubject: emails.PasswordResetSuccessRE,
			expectedResp:         emptyResponse,
			expectedStatus:       http.StatusOK,
		},
		{
			name:           "bad token (user not found)",
			email:          "eventure@datum.net",
			tokenSet:       true,
			tokenProvided:  "thisisnotavalidtoken",
			newPassword:    newPassword,
			emailExpected:  false,
			from:           "notactuallyanemail",
			expectedResp:   "password reset token invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "weak password",
			email:          "sazed@datum.net",
			tokenSet:       true,
			newPassword:    "weak1",
			emailExpected:  false,
			from:           "nottodaysatan",
			expectedResp:   "password is too weak",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "same password",
			email:          "sventure@datum.net",
			tokenSet:       true,
			newPassword:    validPassword,
			emailExpected:  false,
			from:           "mmhmm",
			expectedResp:   "password was already used",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing token",
			email:          "dockson@datum.net",
			tokenSet:       false,
			newPassword:    newPassword,
			emailExpected:  false,
			from:           "yadayadayada",
			expectedResp:   "token is required",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "expired reset token",
			email:          "tensoon@datum.net",
			newPassword:    "6z9Fqc-E-9v32NsJzLNP",
			tokenSet:       true,
			emailExpected:  false,
			from:           "zonkertons",
			ttl:            expiredTTL,
			expectedResp:   "reset token is expired, please request a new token using forgot-password",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sent := time.Now()

			mock.ResetEmailMock()

			// create user in the database
			rt, _, err := createUserWithResetToken(t, client, ec, tc.email, tc.ttl)
			require.NoError(t, err)

			pwResetJSON := handlers.ResetPasswordRequest{
				Password: tc.newPassword,
			}

			body, err := json.Marshal(pwResetJSON)
			if err != nil {
				require.NoError(t, err)
			}

			target := "/password-reset"

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
			client.e.ServeHTTP(recorder, req)

			// get result
			res := recorder.Result()
			defer res.Body.Close()

			// check status
			assert.Equal(t, tc.expectedStatus, recorder.Code)

			var out *handlers.ResetPasswordReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			if tc.expectedStatus != http.StatusOK {
				assert.Contains(t, out.Error, tc.expectedResp)
			}

			// Test that one verify email was sent to each user
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      tc.from,
					Subject:   tc.expectedEmailSubject,
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

// createUserWithResetToken creates a user with a valid reset token and returns the token, user id, and error if one occurred
func createUserWithResetToken(t *testing.T, client *client, ec context.Context, email string, ttl string) (*ent.PasswordResetToken, string, error) {
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	// add mocks for writes
	mock_fga.WriteAny(t, client.fga)

	userSetting := client.db.UserSetting.Create().
		SetEmailConfirmed(true).
		SaveX(ctx)

	u := client.db.User.Create().
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

	if ttl != "" {
		user.PasswordResetExpires.String = ttl
	}

	// set expiry if provided in test case
	expiry, err := time.Parse(time.RFC3339Nano, user.PasswordResetExpires.String)
	if err != nil {
		return nil, "", err
	}

	// store token in db
	pr := client.db.PasswordResetToken.Create().
		SetOwner(u).
		SetToken(user.PasswordResetToken.String).
		SetEmail(user.Email).
		SetSecret(user.PasswordResetSecret).
		SetTTL(expiry).
		SaveX(ctx)

	// clear mocks
	mock_fga.ClearMocks(client.fga)

	return pr, u.ID, nil
}
