package graphapi_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/emails/mock"
	echo "github.com/datumforge/echox"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/rShetty/asyncwait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *GraphTestSuite) TestMutationCreateSubscriber() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	validCtx, err := auth.NewTestContextWithOrgID(org.ID)
	require.NoError(t, err)

	reqCtx = context.WithValue(validCtx.Request().Context(), echocontext.EchoContextKey, validCtx)

	validCtx.SetRequest(validCtx.Request().WithContext(reqCtx))

	testCases := []struct {
		name          string
		input         datumclient.CreateSubscriberInput
		ctx           echo.Context
		emailExpected bool
		errorMsg      string
	}{
		{
			name: "happy path",
			input: datumclient.CreateSubscriberInput{
				Email: "brax@datum.net",
			},
			emailExpected: true,
			ctx:           validCtx,
		},
		{
			name:          "email missing",
			input:         datumclient.CreateSubscriberInput{},
			ctx:           validCtx,
			emailExpected: false,
			errorMsg:      "email is required",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			sent := time.Now()

			mock.ResetEmailMock()

			if tc.errorMsg == "" {
				// mock a call to check orgs
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", org.ID)})
			}

			resp, err := suite.client.datum.CreateSubscriber(tc.ctx.Request().Context(), tc.input)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.input.Email, resp.CreateSubscriber.Subscriber.Email)

			// Test that one email was sent for to the subscriber
			messages := []*mock.EmailMetadata{
				{
					To:        tc.input.Email,
					From:      "mitb@datum.net",
					Subject:   fmt.Sprintf(emails.Subscribed, org.Name),
					Timestamp: sent,
				},
			}

			// wait for messages
			predicate := func() bool {
				return suite.client.db.Marionette.GetQueueLength() == 0
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

func (suite *GraphTestSuite) TestMutationDeleteSubscriber() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	sub := (&SubscriberBuilder{client: suite.client, OrgID: org.ID}).MustNew(reqCtx, t)

	// create a user context with org set in claims
	validCtx, err := auth.NewTestContextWithOrgID(org.ID)
	require.NoError(t, err)

	reqCtx = context.WithValue(validCtx.Request().Context(), echocontext.EchoContextKey, validCtx)

	validCtx.SetRequest(validCtx.Request().WithContext(reqCtx))

	// create a user context with org set in claims, but for another org than the subscriber
	otherOrgCtx, err := auth.NewTestContextWithOrgID(org2.ID)
	require.NoError(t, err)

	reqCtx2 := context.WithValue(validCtx.Request().Context(), echocontext.EchoContextKey, otherOrgCtx)

	otherOrgCtx.SetRequest(validCtx.Request().WithContext(reqCtx2))

	testCases := []struct {
		name     string
		email    string
		ctx      echo.Context
		errorMsg string
	}{
		{
			name:  "happy path",
			email: sub.Email,
			ctx:   validCtx,
		},
		{
			name:     "subscriber not found",
			email:    gofakeit.Email(),
			ctx:      validCtx,
			errorMsg: "subscriber not found",
		},
		{
			name:     "subscriber for another org",
			email:    sub.Email,
			ctx:      otherOrgCtx,
			errorMsg: "subscriber not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := suite.client.datum.DeleteSubscriber(tc.ctx.Request().Context(), tc.email)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.email, resp.DeleteSubscriber.Email)
		})
	}
}
