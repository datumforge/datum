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

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

func TestLoginHandler(t *testing.T) {
	h := handlerSetup(t)

	ec := echocontext.NewTestEchoContext().Request().Context()

	// set privacy allow in order to allow the creation of the users without
	// authentication in the tests
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	// create user in the database
	validConfirmedUser := "rsanchez@datum.net"
	validPassword := "sup3rs3cu7e!"

	userSetting := EntClient.UserSetting.Create().
		SetEmailConfirmed(true).
		SaveX(ec)

	userConfirmed := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail(validConfirmedUser).
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ctx)

	validUnconfirmedUser := "msmith@datum.net"

	userSetting = EntClient.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ctx)

	userUnconfirmed := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail(validUnconfirmedUser).
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ctx)

	testCases := []struct {
		name           string
		username       string
		password       string
		expectedErr    error
		expectedStatus int
	}{
		{
			name:           "happy path, valid credentials",
			username:       validConfirmedUser,
			password:       validPassword,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "email unverified",
			username:       validUnconfirmedUser,
			password:       validPassword,
			expectedStatus: http.StatusBadRequest,
			expectedErr:    auth.ErrUnverifiedUser,
		},
		{
			name:           "invalid password",
			username:       validConfirmedUser,
			password:       "thisisnottherightone",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    auth.ErrInvalidCredentials,
		},
		{
			name:           "user not found",
			username:       "rick.sanchez@datum.net",
			password:       validPassword,
			expectedStatus: http.StatusBadRequest,
			expectedErr:    auth.ErrNoAuthUser,
		},
		{
			name:           "empty username",
			username:       "",
			password:       validPassword,
			expectedStatus: http.StatusBadRequest,
			expectedErr:    handlers.ErrMissingRequiredFields,
		},
		{
			name:           "empty username",
			username:       validConfirmedUser,
			password:       "",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    handlers.ErrMissingRequiredFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create echo context with middleware
			e := setupEcho(h.SM)
			e.POST("login", h.LoginHandler)

			loginJSON := handlers.LoginRequest{
				Username: tc.username,
				Password: tc.password,
			}

			body, err := json.Marshal(loginJSON)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

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

			if tc.expectedStatus == http.StatusOK {
				assert.Equal(t, out.Message, "success")
			} else {
				assert.Contains(t, out.Message, tc.expectedErr.Error())
			}
		})
	}

	// cleanup after
	EntClient.User.DeleteOneID(userConfirmed.ID).ExecX(ctx)
	EntClient.User.DeleteOneID(userUnconfirmed.ID).ExecX(ctx)
}
