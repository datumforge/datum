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

func TestForgotPasswordHandler(t *testing.T) {
	h := handlerSetup(t)

	ec := echocontext.NewTestEchoContext().Request().Context()

	// create user in the database
	userSetting := EntClient.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ec)

	u := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail("asandler@datum.net").
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ec)

	testCases := []struct {
		name               string
		email              string
		expectedErrMessage string
		expectedStatus     int
	}{
		{
			name:           "happy path",
			email:          "asandler@datum.net",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "email does not exist, should still return 204",
			email:          "asandler1@datum.net",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "email not sent in request",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := setupEcho()
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
		})
	}

	// cleanup after
	EntClient.User.DeleteOneID(u.ID).ExecX(ec)
}
