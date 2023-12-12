package handlers_test

import (
	"crypto/rand"
	"crypto/rsa"
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
	"go.uber.org/zap"

	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/tokens"
)

func createTokenManager() (*tokens.TokenManager, error) {
	conf := tokens.Config{
		Audience:        "http://localhost:17608",
		Issuer:          "http://localhost:17608",
		AccessDuration:  1 * time.Hour,     // nolint: gomnd
		RefreshDuration: 2 * time.Hour,     // nolint: gomnd
		RefreshOverlap:  -15 * time.Minute, // nolint: gomnd
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048) // nolint: gomnd
	if err != nil {
		return nil, err
	}

	return tokens.NewWithKey(key, conf)
}

func TestLoginHandler(t *testing.T) {
	tm, err := createTokenManager()
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
		name     string
		username string
		password string
		err      error
	}{
		{
			name:     "happy path, valid credentials",
			username: validConfirmedUser,
			password: validPassword,
			err:      nil,
		},
		{
			name:     "email unverified",
			username: validUnconfirmedUser,
			password: validPassword,
			err:      auth.ErrUnverifiedUser,
		},
		{
			name:     "invalid password",
			username: validConfirmedUser,
			password: "thisisnottherightone",
			err:      auth.ErrInvalidCredentials,
		},
		{
			name:     "user not found",
			username: "rick.sanchez@datum.net",
			password: validPassword,
			err:      auth.ErrNoAuthUser,
		},
		{
			name:     "empty username",
			username: "",
			password: validPassword,
			err:      handlers.ErrMissingRequiredFields,
		},
		{
			name:     "empty username",
			username: validConfirmedUser,
			password: "",
			err:      handlers.ErrMissingRequiredFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create echo context
			e := echo.New()

			userJSON := handlers.User{
				Username: tc.username,
				Password: tc.password,
			}

			body, err := json.Marshal(userJSON)
			if err != nil {
				t.Error("error creating user json")
			}

			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			ctx := e.NewContext(req, recorder)

			err = h.LoginHandler(ctx)

			if tc.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.err.Error())

				return
			}

			assert.NoError(t, err)
		})
	}

	// cleanup after
	EntClient.User.DeleteOneID(userConfirmed.ID).ExecX(ec)
	EntClient.User.DeleteOneID(userUnconfirmed.ID).ExecX(ec)
}
