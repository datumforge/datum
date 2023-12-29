package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	echo "github.com/datumforge/echox"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

func TestLoginHandler(t *testing.T) {
	tm, err := createTokenManager(15 * time.Minute) //nolint:gomnd
	if err != nil {
		t.Error("error creating token manager")
	}

	h := handlers.Handler{
		TM:           tm,
		DBClient:     EntClient,
		Logger:       zap.NewNop().Sugar(),
		CookieDomain: "datum.net",
	}

	ec := echocontext.NewTestEchoContext().Request().Context()

	// set privacy allow in order to allow the creation of the users without
	// authentication in the tests
	ec = privacy.DecisionContext(ec, privacy.Allow)

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
		SaveX(ec)

	validUnconfirmedUser := "msmith@datum.net"

	userSetting = EntClient.UserSetting.Create().
		SetEmailConfirmed(false).
		SaveX(ec)

	userUnconfirmed := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail(validUnconfirmedUser).
		SetPassword(validPassword).
		SetSetting(userSetting).
		SaveX(ec)

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
			// create echo context
			e := echo.New()

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

			ctx := e.NewContext(req, recorder)

			err = h.LoginHandler(ctx)
			require.NoError(t, err)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.Response

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, ctx.Response().Status)

			if tc.expectedStatus == http.StatusOK {
				assert.Equal(t, out.Message, "success")
			} else {
				assert.Contains(t, out.Message, tc.expectedErr.Error())
			}
		})
	}

	// cleanup after
	EntClient.User.DeleteOneID(userConfirmed.ID).ExecX(ec)
	EntClient.User.DeleteOneID(userUnconfirmed.ID).ExecX(ec)
}
