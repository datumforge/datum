package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func TestResendHandler(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// add handler
	client.e.POST("resend", client.h.ResendEmail)

	ec := echocontext.NewTestEchoContext().Request().Context()

	ctx := privacy.DecisionContext(ec, privacy.Allow)

	// add mocks for writes
	mock_fga.WriteAny(t, client.fga)

	// create user in the database
	userSetting := client.db.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ctx)

	_ = client.db.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("bsanderson@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ctx)

	// create user in the database
	userSetting2 := client.db.UserSetting.Create().
		SetEmailConfirmed(true).
		SaveX(ctx)

	_ = client.db.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("dabraham@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting2).
		SaveX(ctx)

	testCases := []struct {
		name            string
		email           string
		expectedMessage string
		expectedStatus  int
	}{
		{
			name:            "happy path",
			email:           "bsanderson@datum.net",
			expectedStatus:  http.StatusOK,
			expectedMessage: "received your request to be resend",
		},
		{
			name:            "email does not exist, should still return 204",
			email:           "bsanderson1@datum.net",
			expectedStatus:  http.StatusOK,
			expectedMessage: "received your request to be resend",
		},
		{
			name:            "email confirmed",
			email:           "dabraham@datum.net",
			expectedStatus:  http.StatusOK,
			expectedMessage: "email is already confirmed",
		},
		{
			name:           "email not sent in request",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resendJSON := handlers.ResendRequest{
				Email: tc.email,
			}

			body, err := json.Marshal(resendJSON)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/resend", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			client.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.ResendReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			require.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus == http.StatusNoContent {
				require.NotEmpty(t, out)
				assert.NotEmpty(t, out.Message)
			} else {
				assert.Contains(t, out.Message, tc.expectedMessage)
			}
		})
	}
}
