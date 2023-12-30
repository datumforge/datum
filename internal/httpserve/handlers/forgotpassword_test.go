package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	echo "github.com/datumforge/echox"
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
			// create echo context
			e := echo.New()

			resendJSON := handlers.ForgotPasswordRequest{
				Email: tc.email,
			}

			body, err := json.Marshal(resendJSON)
			if err != nil {
				t.Error("error creating resend json")
			}

			req := httptest.NewRequest(http.MethodPost, "/forgot-password", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			ctx := e.NewContext(req, recorder)

			err = h.ForgotPassword(ctx)
			require.NoError(t, err)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.Response

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, ctx.Response().Status)

			if tc.expectedStatus == http.StatusNoContent {
				assert.Nil(t, out)
			} else {
				assert.Contains(t, out.Message, tc.expectedErrMessage)
			}
		})
	}

	// cleanup after
	EntClient.User.DeleteOneID(u.ID).ExecX(ec)
}
