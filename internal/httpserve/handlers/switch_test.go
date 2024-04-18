package handlers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	echo "github.com/datumforge/echox"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/emails/mock"
)

type Organizations struct {
	Org1 *ent.Organization
	Org2 *ent.Organization
}

func (suite *HandlerTestSuite) TestSwitchHandler() {
	t := suite.T()

	suite.e.POST("switch", suite.h.SwitchHandler)

	tm, err := createTokenManager(-60 * time.Minute) //nolint:gomnd
	if err != nil {
		t.Error("error creating token manager")
	}

	suite.h.TM = tm

	ctx := context.Background()
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	mock_fga.WriteAny(t, suite.fga)

	// setup test data
	user := suite.db.User.Create().
		SetEmail("rocket@datum.net").
		SetFirstName("Rocket").
		SetLastName("Racoon").
		SaveX(ctx)

	ec, err := auth.NewTestContextWithValidUser(user.ID)
	require.NoError(t, err)

	newCtx := ec.Request().Context()
	newCtx = privacy.DecisionContext(newCtx, privacy.Allow)

	reqCtx := context.WithValue(newCtx, echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

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

	org1 := suite.db.Organization.Create().
		SetName("avengerscrew").
		SaveX(reqCtx)

	org2 := suite.db.Organization.Create().
		SetName("thanoscrew").
		SaveX(reqCtx)

	orgs := &Organizations{
		Org1: org1,
		Org2: org2,
	}

	testCases := []struct {
		name               string
		orgID1             string
		orgID2             string
		orgName1           string
		orgName2           string
		expectedErrMessage string
		expectedStatus     int
		refresh            string
	}{
		{
			name:           "happy path, new subscriber for org",
			orgName1:       orgs.Org1.Name,
			orgID1:         orgs.Org1.ID,
			orgName2:       orgs.Org2.Name,
			orgID2:         orgs.Org2.ID,
			expectedStatus: http.StatusOK,
			refresh:        refresh,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.fga)

			mock.ResetEmailMock()

			// mock auth
			mock_fga.ListAny(t, suite.fga, []string{fmt.Sprintf("organization:%s", orgs.Org2.ID)})

			switchJSON := handlers.SwitchOrganizationRequest{
				TargetOrganizationID: tc.orgID2,
			}

			body, err := json.Marshal(switchJSON)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/switch", strings.NewReader(string(body)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.SwitchOrganizationReply

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
