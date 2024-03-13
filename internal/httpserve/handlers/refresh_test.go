package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

func (suite *HandlerTestSuite) TestRefreshHandler() {
	t := suite.T()

	// add handler
	suite.client.e.POST("refresh", suite.client.h.RefreshHandler)

	// Set full overlap of the refresh and access token so the refresh token is immediately valid
	tm, err := createTokenManager(-60 * time.Minute) //nolint:gomnd
	if err != nil {
		t.Error("error creating token manager")
	}

	suite.client.h.TM = tm

	ec := echocontext.NewTestEchoContext().Request().Context()

	// set privacy allow in order to allow the creation of the users without
	// authentication in the tests
	ec = privacy.DecisionContext(ec, privacy.Allow)

	// add mocks for writes
	mock_fga.WriteAny(t, suite.client.fga)

	// create user in the database
	validUser := gofakeit.Email()
	validPassword := gofakeit.Password(true, true, true, true, false, 20)

	userID := ulids.New().String()

	userSetting := suite.client.db.UserSetting.Create().
		SetEmailConfirmed(true).
		SaveX(ec)

	user := suite.client.db.User.Create().
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
		name               string
		refresh            string
		expectedErrMessage string
		expectedStatus     int
	}{
		{
			name:           "happy path, valid credentials",
			refresh:        refresh,
			expectedStatus: http.StatusOK,
		},
		{
			name:               "empty refresh",
			refresh:            "",
			expectedStatus:     http.StatusBadRequest,
			expectedErrMessage: "refresh_token is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			refreshJSON := handlers.RefreshRequest{
				RefreshToken: tc.refresh,
			}

			body, err := json.Marshal(refreshJSON)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/refresh", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.client.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.RefreshReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus == http.StatusOK {
				assert.True(t, out.Success)
			} else {
				assert.Contains(t, out.Error, tc.expectedErrMessage)
			}
		})
	}
}
