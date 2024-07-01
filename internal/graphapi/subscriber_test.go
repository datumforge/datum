package graphapi_test

import (
	"context"
	"testing"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *GraphTestSuite) TestQuerySubscriber() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// Org to subscribe users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org
	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org.ID)
	require.NoError(t, err)

	subscriber := (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org2
	reqCtx2, err := auth.NewTestContextWithOrgID(testUser.ID, org2.ID)
	require.NoError(t, err)

	sub := (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx2, t)

	testCases := []struct {
		name        string
		email       string
		shouldCheck bool
		expected    *ent.Subscriber
		wantErr     bool
	}{
		{
			name:        "happy path",
			email:       subscriber.Email,
			shouldCheck: true,
			expected:    subscriber,
			wantErr:     false,
		},
		{
			name:        "invalid email",
			email:       "beep@boop.com",
			shouldCheck: false,
			expected:    nil,
			wantErr:     true,
		},
		{
			name:        "subscriber for another org",
			email:       sub.Email,
			shouldCheck: false,
			expected:    nil,
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.shouldCheck {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			resp, err := suite.client.datum.GetSubscriberByEmail(reqCtx, tc.email)

			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Subscriber)
		})
	}
}

func (suite *GraphTestSuite) TestQuerySubscribers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// Org to subscribe users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org
	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org.ID)
	require.NoError(t, err)

	_ = (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org2
	reqCtx2, err := auth.NewTestContextWithOrgID(testUser.ID, org2.ID)
	require.NoError(t, err)

	_ = (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx2, t)

	testCases := []struct {
		name        string
		context     context.Context
		numExpected int
		check       bool
	}{
		{
			name:        "happy path, multiple subscribers",
			context:     reqCtx,
			check:       false,
			numExpected: 2,
		},
		{
			name:        "happy path, one subscriber",
			context:     reqCtx2,
			check:       true,
			numExpected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// TODO (sfunk): this is because of 1 vs multiple returned, look at the filter check
			if tc.check {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			resp, err := suite.client.datum.GetAllSubscribers(tc.context)

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Len(t, resp.Subscribers.Edges, tc.numExpected)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateSubscriber() {
	t := suite.T()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	testCases := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "happy path, new subscriber",
			email:   "c.stark@example.com",
			wantErr: false,
		},
		{
			name:    "missing email",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, true)

			input := datumclient.CreateSubscriberInput{
				Email: tc.email,
			}

			resp, err := suite.client.datum.CreateSubscriber(ctx, input)

			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			// Assert matching fields
			assert.Equal(t, tc.email, resp.CreateSubscriber.Subscriber.Email)

		})
	}
}

func (suite *GraphTestSuite) TestUpdateSubscriber() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// Org to subscribe users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org
	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org.ID)
	require.NoError(t, err)

	subscriber := (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org2
	reqCtx2, err := auth.NewTestContextWithOrgID(testUser.ID, org2.ID)
	require.NoError(t, err)

	sub := (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx2, t)

	testCases := []struct {
		name        string
		email       string
		updateInput datumclient.UpdateSubscriberInput
		shouldCheck bool
		expected    *ent.Subscriber
		wantErr     bool
	}{
		{
			name:  "happy path",
			email: subscriber.Email,
			updateInput: datumclient.UpdateSubscriberInput{
				PhoneNumber: lo.ToPtr("+1-555-867-5309"),
			},
			expected: subscriber,
			wantErr:  false,
		},
		{
			name:  "invalid email",
			email: "beep@boop.com",
			updateInput: datumclient.UpdateSubscriberInput{
				PhoneNumber: lo.ToPtr("+1-555-867-5309"),
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name:  "subscriber for another org",
			email: sub.Email,
			updateInput: datumclient.UpdateSubscriberInput{
				PhoneNumber: lo.ToPtr("+1-555-867-5309"),
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if !tc.wantErr {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			resp, err := suite.client.datum.UpdateSubscriber(reqCtx, tc.email, tc.updateInput)

			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Equal(t, tc.email, resp.UpdateSubscriber.Subscriber.Email)
			require.Equal(t, tc.updateInput.PhoneNumber, resp.UpdateSubscriber.Subscriber.PhoneNumber)
		})
	}
}

func (suite *GraphTestSuite) TestDeleteSubscriber() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// Org to subscribe users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org
	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org.ID)
	require.NoError(t, err)

	subscriber := (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org2
	reqCtx2, err := auth.NewTestContextWithOrgID(testUser.ID, org2.ID)
	require.NoError(t, err)

	sub := (&SubscriberBuilder{client: suite.client}).MustNew(reqCtx2, t)

	testCases := []struct {
		name        string
		email       string
		shouldCheck bool
		expected    *ent.Subscriber
		wantErr     bool
	}{
		{
			name:     "happy path",
			email:    subscriber.Email,
			expected: subscriber,
			wantErr:  false,
		},
		{
			name:     "invalid email",
			email:    "beep@boop.com",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "subscriber for another org",
			email:    sub.Email,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, true)

			resp, err := suite.client.datum.DeleteSubscriber(reqCtx, tc.email)

			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Equal(t, tc.email, resp.DeleteSubscriber.Email)
		})
	}
}
