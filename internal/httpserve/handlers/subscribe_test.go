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

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/emails/mock"
)

func (suite *HandlerTestSuite) TestSubscribeHandler() {
	t := suite.T()

	// add handler
	suite.e.POST("subscribe", suite.h.RegisterHandler)

	ctx := context.Background()
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	org := suite.db.Organization.Create().
		SetName("avengers").
		SaveX(ctx)

	testCases := []struct {
		name               string
		email              string
		orgID              string
		emailExpected      bool
		expectedErrMessage string
		expectedStatus     int
		from               string
	}{
		{
			name:           "happy path, root subscriber",
			email:          "brax@datum.net",
			emailExpected:  true,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "happy path, root subscriber",
			email:          "brax@datum.net",
			orgID:          org.ID,
			emailExpected:  true,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "missing email",
			orgID:          org.ID,
			emailExpected:  true,
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.fga)

			sent := time.Now()

			mock.ResetEmailMock()

			// setup mock authz writes
			if tc.expectedErrMessage == "" {
				mock_fga.WriteAny(t, suite.fga)
			}

			target := fmt.Sprintf("/subscribe?email=%s", tc.email)
			if tc.orgID != "" {
				target = fmt.Sprintf("%s&org=%s", target, tc.orgID)
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
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      "mitb@datum.net",
					Subject:   emails.Subscribed,
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
