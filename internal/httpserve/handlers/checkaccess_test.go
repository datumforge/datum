package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/pkg/httpsling"
	"github.com/datumforge/datum/pkg/models"
)

func (suite *HandlerTestSuite) TestCheckAccessHandler() {
	t := suite.T()

	// add handler
	suite.e.POST("check-access", suite.h.CheckAccessHandler)

	// bypass auth
	ctx := context.Background()
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	mock_fga.WriteAny(t, suite.fga)

	// setup test data
	requestor := suite.db.User.Create().
		SetEmail("marco@datum.net").
		SetFirstName("Marco").
		SetLastName("Polo").
		SaveX(ctx)

	reqCtx, err := userContextWithID(requestor.ID)
	require.NoError(t, err)

	mock_fga.ClearMocks(suite.fga)

	testCases := []struct {
		name      string
		request   models.CheckAccessRequest
		mockAllow bool
		errMsg    string
	}{
		{
			name:      "happy path, allow access",
			mockAllow: true,
			request: models.CheckAccessRequest{
				ObjectID:   "org-id",
				ObjectType: "organization",
				Relation:   "can_view",
			},
		},
		{
			name:      "access denied",
			mockAllow: false,
			request: models.CheckAccessRequest{
				ObjectID:   "another-org-id",
				ObjectType: "organization",
				Relation:   "can_delete",
			},
		},
		{
			name: "missing object id",
			request: models.CheckAccessRequest{
				ObjectType: "organization",
				Relation:   "can_delete",
			},
			errMsg: "objectId is required",
		},
		{
			name: "missing object type",
			request: models.CheckAccessRequest{
				ObjectID: "org-id",
				Relation: "can_delete",
			},
			errMsg: "objectType is required",
		},
		{
			name: "missing relation",
			request: models.CheckAccessRequest{
				ObjectID:   "org-id",
				ObjectType: "organization",
			},
			errMsg: "relation is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.fga)

			if tc.errMsg == "" {
				mock_fga.CheckAny(t, suite.fga, tc.mockAllow)
			}

			target := "/check-access"

			body, err := json.Marshal(tc.request)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(string(body)))
			req.Header.Set(httpsling.HeaderContentType, httpsling.ContentTypeJSONUTF8)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req.WithContext(reqCtx))

			res := recorder.Result()
			defer res.Body.Close()

			var out *models.CheckAccessReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			if tc.errMsg != "" {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				assert.False(t, out.Success)
				assert.Equal(t, tc.errMsg, out.Error)

				return
			}

			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.True(t, out.Success)
			assert.Equal(t, tc.mockAllow, out.Allowed)
		})
	}
}
