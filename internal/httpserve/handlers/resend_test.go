package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

func TestResendHandler(t *testing.T) {
	h := handlerSetup(t)

	ec := echocontext.NewTestEchoContext().Request().Context()

	// create user in the database
	userSetting := EntClient.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ec)

	u := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("bsanderson@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ec)

	// create user in the database
	userSetting2 := EntClient.UserSetting.Create().
		SetEmailConfirmed(true).
		SaveX(ec)

	u2 := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("dabraham@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting2).
		SaveX(ec)

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
			// create echo context with middleware
			e := setupEcho(h.SM)
			e.POST("resend", h.ResendEmail)

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
			e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.Response

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus == http.StatusNoContent {
				require.NotEmpty(t, out)
				assert.NotEmpty(t, out.Message)
			} else {
				assert.Contains(t, out.Message, tc.expectedMessage)
			}
		})
	}

	// cleanup after
	EntClient.User.DeleteOneID(u.ID).ExecX(ec)
	EntClient.User.DeleteOneID(u2.ID).ExecX(ec)
}
