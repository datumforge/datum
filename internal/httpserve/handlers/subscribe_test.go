package handlers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/rShetty/asyncwait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/emails/mock"
)

func (suite *HandlerTestSuite) TestSubscribeHandler() {
	t := suite.T()

	// add handler
	suite.e.GET("subscribe", suite.h.SubscribeHandler)

	org := suite.createTestOrg(t)

	testCases := []struct {
		name               string
		email              string
		orgID              string
		orgName            string
		emailExpected      bool
		expectedErrMessage string
		expectedStatus     int
		from               string
	}{
		{
			name:           "happy path, new subscriber",
			email:          "brax@datum.net",
			emailExpected:  true,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "subscriber already exists",
			email:          "brax@datum.net",
			emailExpected:  false,
			expectedStatus: http.StatusConflict,
		},
		{
			name:           "happy path, new subscriber for org",
			email:          "brax@datum.net",
			orgName:        org.Name,
			orgID:          org.ID,
			emailExpected:  true,
			expectedStatus: http.StatusCreated,
		},
		{
			name:               "missing email",
			emailExpected:      false,
			expectedStatus:     http.StatusBadRequest,
			expectedErrMessage: "email is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.fga)

			sent := time.Now()

			mock.ResetEmailMock()

			target := fmt.Sprintf("/subscribe?email=%s", tc.email)
			if tc.orgID != "" {
				target = fmt.Sprintf("%s&organization_id=%s", target, tc.orgID)
			}

			req := httptest.NewRequest(http.MethodGet, target, nil)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.SubscribeReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus == http.StatusOK {
				assert.NotEmpty(t, out.Message)
			} else {
				assert.Contains(t, out.Error, tc.expectedErrMessage)
			}

			// Test that one verify email was sent to each user
			orgName := "MITB"
			if tc.orgName != "" {
				orgName = tc.orgName
			}

			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      "mitb@datum.net",
					Subject:   fmt.Sprintf(emails.Subscribed, orgName),
					Timestamp: sent,
				},
			}

			// wait for messages
			predicate := func() bool {
				return suite.h.TaskMan.GetQueueLength() == 0
			}
			successful := asyncwait.NewAsyncWait(maxWaitInMillis, pollIntervalInMillis).Check(predicate)

			if successful != true {
				t.Errorf("max wait of email send")
			}

			if tc.emailExpected {
				mock.CheckEmails(t, messages)
			} else {
				mock.CheckEmails(t, nil)
			}
		})
	}
}

func (suite *HandlerTestSuite) createTestOrg(t *testing.T) *ent.Organization {
	mock_fga.WriteAny(t, suite.fga)
	// bypass auth
	ctx := context.Background()
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	// setup test data
	requestor := suite.db.User.Create().
		SetEmail("rocket@datum.net").
		SetFirstName("Rocket").
		SetLastName("Racoon").
		SaveX(ctx)

	ec, err := auth.NewTestContextWithValidUser(requestor.ID)
	require.NoError(t, err)

	newCtx := ec.Request().Context()
	newCtx = privacy.DecisionContext(newCtx, privacy.Allow)

	reqCtx := context.WithValue(newCtx, echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	return suite.db.Organization.Create().
		SetName("avengers").
		SaveX(reqCtx)
}
