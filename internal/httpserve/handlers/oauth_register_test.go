package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	echo "github.com/datumforge/echox"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/tokens"
)

func (suite *HandlerTestSuite) TestOauthRegister() {
	t := suite.T()

	// add login handler
	suite.e.POST("oauth/register", suite.h.OauthRegister)

	type args struct {
		name     string
		email    string
		provider enums.AuthProvider
		username string
		userID   string
		token    string
	}

	tests := []struct {
		name           string
		args           args
		writes         bool
		expectedStatus int
		expectedErr    string
		wantErr        bool
	}{
		{
			name: "happy path, github",
			args: args{
				name:     "Ant Man",
				email:    "antman@datum.net",
				provider: enums.GitHub,
				username: "scarletwitch",
				userID:   "123456",
				token:    "gh_thistokenisvalid",
			},
			expectedStatus: http.StatusOK,
			writes:         true,
		},
		{
			name: "happy path, github, same user",
			args: args{
				name:     "Ant Man",
				email:    "antman@datum.net",
				provider: enums.GitHub,
				username: "scarletwitch",
				userID:   "123456",
				token:    "gh_thistokenisvalid",
			},
			expectedStatus: http.StatusOK,
			writes:         false, // user already created, no FGA writes this time
		},
		{
			name: "mismatch email",
			args: args{
				name:     "Ant Man",
				email:    "antman@marvel.com",
				provider: enums.GitHub,
				username: "scarletwitch",
				userID:   "123456",
				token:    "gh_thistokenisvalid",
			},
			expectedStatus: http.StatusBadRequest,
			writes:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.writes {
				// add mocks for writes when a new user is created
				mock_fga.WriteOnce(t, suite.fga)
			}

			if !tt.writes && tt.expectedStatus == http.StatusOK {
				// required to list objects to get the default org in claims
				// when the user is not created in the call
				mock_fga.ListAny(t, suite.fga, []string{"organization:test"})
			}

			registerJSON := handlers.OauthTokenRequest{
				Name:             tt.args.name,
				Email:            tt.args.email,
				AuthProvider:     tt.args.provider.String(),
				ExternalUserID:   tt.args.userID,
				ExternalUserName: tt.args.username,
				ClientToken:      tt.args.token,
			}

			body, err := json.Marshal(registerJSON)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/oauth/register", strings.NewReader(string(body)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *tokens.TokenResponse

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tt.expectedStatus, recorder.Code)

			if tt.expectedStatus == http.StatusOK {
				assert.NotNil(t, out.AccessToken)
				assert.NotNil(t, out.RefreshToken)
				assert.NotNil(t, out.ExpiresIn)
				assert.Equal(t, "Bearer", out.TokenType)
			}
		})
	}
}
