package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
)

func (suite *HandlerTestSuite) TestUserInfoHandler() {
	t := suite.T()

	// bypass auth
	ctx := context.Background()
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	mock_fga.WriteAny(t, suite.fga)

	// setup test data
	user := suite.db.User.Create().
		SetEmail("juju@datum.net").
		SetFirstName("Juju").
		SetLastName("Bee").
		SaveX(ctx)

	ec, err := auth.NewTestEchoContextWithValidUser(user.ID)
	require.NoError(t, err)

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	suite.e.GET("oauth/userinfo", suite.h.UserInfo)

	tests := []struct {
		name    string
		ctx     context.Context
		wantErr bool
	}{
		{
			name:    "happy path",
			ctx:     reqCtx,
			wantErr: false,
		},
		{
			name:    "empty context",
			ctx:     context.Background(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new request
			req := httptest.NewRequest(http.MethodGet, "/oauth/userinfo", nil)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req.WithContext(tt.ctx))

			res := recorder.Result()
			defer res.Body.Close()

			var out *ent.User

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			if tt.wantErr {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)

				return
			}

			assert.Equal(t, http.StatusOK, recorder.Code)
			require.NotNil(t, out)

			assert.Equal(t, user.ID, out.ID)
		})
	}
}
