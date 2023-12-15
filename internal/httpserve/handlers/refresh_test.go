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
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func TestRefreshHandler(t *testing.T) {
	// Set full overlap of the refresh and access token so the refresh token is immediately valid
	tm, err := createTokenManager(-60 * time.Minute) //nolint:gomnd
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
	validUser := gofakeit.Email()
	validPassword := gofakeit.Password(true, true, true, true, false, 20)

	userID := ulids.New().String()

	userSetting := EntClient.UserSetting.Create().
		SetEmailConfirmed(true).
		SaveX(ec)

	user := EntClient.User.Create().
		SetFirstName(gofakeit.FirstName()).
		SetLastName(gofakeit.LastName()).
		SetEmail(validUser).
		SetPassword(validPassword).
		SetSetting(userSetting).
		SetID(userID).
		SetSub(userID). // this is required to parse the refresh token
		SaveX(ec)

	claims := &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.ID,
		},
		UserID: user.ID,
		Email:  user.Email,
	}

	_, refresh, err := tm.CreateTokenPair(claims)
	if err != nil {
		t.Error("error creating token pair")
	}

	testCases := []struct {
		name    string
		refresh string
		err     error
	}{
		{
			name:    "happy path, valid credentials",
			refresh: refresh,
			err:     nil,
		},
		{
			name:    "empty refresh",
			refresh: "",
			err:     handlers.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create echo context
			e := echo.New()

			refreshJSON := handlers.RefreshRequest{
				RefreshToken: tc.refresh,
			}

			body, err := json.Marshal(refreshJSON)
			if err != nil {
				t.Error("error creating refresh json")
			}

			req := httptest.NewRequest(http.MethodPost, "/refresh", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			ctx := e.NewContext(req, recorder)

			err = h.RefreshHandler(ctx)

			if tc.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.err.Error())

				return
			}

			assert.NoError(t, err)
		})
	}

	// cleanup after
	EntClient.User.DeleteOneID(user.ID).ExecX(ec)
}
