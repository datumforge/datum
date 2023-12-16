package graphapi_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

func TestQuery_GroupsNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)

	group1 := (&GroupBuilder{Owner: org1.ID}).MustNew(reqCtx)
	group2 := (&GroupBuilder{Owner: org1.ID}).MustNew(reqCtx)

	t.Run("Get Groups", func(t *testing.T) {
		resp, err := client.GetAllGroups(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Groups.Edges)

		// make sure at least two groups are returned
		assert.GreaterOrEqual(t, len(resp.Groups.Edges), 2)

		group1Found := false
		group2Found := false
		for _, o := range resp.Groups.Edges {
			if o.Node.ID == group1.ID {
				group1Found = true
			} else if o.Node.ID == group2.ID {
				group2Found = true
			}
		}

		// if one of the orgs isn't found, fail the test
		if !group1Found || !group2Found {
			t.Fail()
		}
	})

	// delete created orgs and groups
	(&GroupCleanup{GroupID: group1.ID}).MustDelete(reqCtx)
	(&GroupCleanup{GroupID: group2.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
}

func TestQuery_GroupsByOwnerNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	org2 := (&OrganizationBuilder{}).MustNew(reqCtx)

	group1 := (&GroupBuilder{Owner: org1.ID}).MustNew(reqCtx)
	group2 := (&GroupBuilder{Owner: org2.ID}).MustNew(reqCtx)

	t.Run("Get Groups By Owner", func(t *testing.T) {
		whereInput := &datumclient.GroupWhereInput{
			HasOwnerWith: []*datumclient.OrganizationWhereInput{
				{
					ID: &org1.ID,
				},
			},
		}

		resp, err := client.GroupsWhere(reqCtx, whereInput)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Groups.Edges)

		// make sure 1 group is returned
		assert.Equal(t, 1, len(resp.Groups.Edges))

		group1Found := false
		group2Found := false
		for _, o := range resp.Groups.Edges {
			if o.Node.ID == group1.ID {
				group1Found = true
			} else if o.Node.ID == group2.ID {
				group2Found = true
			}
		}

		// group1 should be returned, group 2 should not be returned
		if !group1Found || group2Found {
			t.Fail()
		}
	})

	// delete created orgs and groups
	(&GroupCleanup{GroupID: group1.ID}).MustDelete(reqCtx)
	(&GroupCleanup{GroupID: group2.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org2.ID}).MustDelete(reqCtx)
}

func TestQuery_GroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	group1 := (&GroupBuilder{Owner: org1.ID}).MustNew(reqCtx)

	testCases := []struct {
		name     string
		queryID  string
		expected *ent.Group
		errorMsg string
	}{
		{
			name:     "happy path organization",
			queryID:  group1.ID,
			expected: group1,
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			errorMsg: "group not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetGroupByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Group)
		})
	}

	// delete created orgs and groups
	(&GroupCleanup{GroupID: group1.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
}

func TestMutation_CreateGroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	org := (&OrganizationBuilder{}).MustNew(reqCtx)

	testCases := []struct {
		name        string
		groupName   string
		description string
		displayName string
		logoURL     string
		owner       string
		errorMsg    string
	}{
		{
			name:        "happy path group",
			groupName:   gofakeit.Name(),
			displayName: gofakeit.LetterN(50),
			description: gofakeit.HipsterSentence(10),
			logoURL:     "http://mitb.net/logo.png",
			owner:       org.ID,
		},
		{
			name:      "happy path group, minimum fields",
			groupName: gofakeit.Name(),
			owner:     org.ID,
		},
		{
			name:      "missing owner",
			groupName: gofakeit.Name(),
			errorMsg:  "constraint failed", // TODO: better error messaging
		},
		{
			name:     "missing name",
			owner:    org.ID,
			errorMsg: "validator failed",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			tc := tc
			input := datumclient.CreateGroupInput{
				Name:        tc.groupName,
				Description: &tc.description,
				DisplayName: &tc.displayName,
				LogoURL:     &tc.logoURL,
				OwnerID:     tc.owner,
			}

			if tc.displayName != "" {
				input.DisplayName = &tc.displayName
			}

			resp, err := client.CreateGroup(reqCtx, input)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateGroup.Group)

			// Make sure provided values match
			assert.Equal(t, tc.groupName, resp.CreateGroup.Group.Name)
			assert.Equal(t, tc.description, *resp.CreateGroup.Group.Description)
			assert.Equal(t, tc.logoURL, *resp.CreateGroup.Group.LogoURL)
			assert.Equal(t, tc.owner, resp.CreateGroup.Group.Owner.ID)

			if tc.displayName != "" {
				assert.Equal(t, tc.displayName, resp.CreateGroup.Group.DisplayName)
			} else {
				// display name defaults to the name if not set
				assert.Equal(t, tc.groupName, resp.CreateGroup.Group.DisplayName)
			}

			// cleanup group
			(&GroupCleanup{GroupID: resp.CreateGroup.Group.ID}).MustDelete(reqCtx)
		})
	}

	(&OrganizationCleanup{OrgID: org.ID}).MustDelete(reqCtx)
}

func TestMutation_UpdateGroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	group := (&GroupBuilder{}).MustNew(reqCtx)

	reqCtx = privacy.DecisionContext(reqCtx, privacy.Allow)

	nameUpdate := gofakeit.Name()
	nameUpdate2 := gofakeit.Name()
	displayNameUpdate := gofakeit.LetterN(40)
	displayNameUpdate2 := gofakeit.LetterN(20)

	descriptionUpdate := gofakeit.HipsterSentence(10)

	testCases := []struct {
		name        string
		updateInput datumclient.UpdateGroupInput
		expectedRes datumclient.UpdateGroup_UpdateGroup_Group
		errorMsg    string
	}{
		{
			name: "update name, happy path",
			updateInput: datumclient.UpdateGroupInput{
				Name: &nameUpdate,
			},
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate,
				DisplayName: nameUpdate, // display name should update if name is updated without display name
				Description: &group.Description,
			},
		},
		{
			name: "update name and display name",
			updateInput: datumclient.UpdateGroupInput{
				Name:        &nameUpdate2,
				DisplayName: &displayNameUpdate,
			},
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate2,
				DisplayName: displayNameUpdate,
				Description: &group.Description,
			},
		},
		{
			name: "update just display name",
			updateInput: datumclient.UpdateGroupInput{
				DisplayName: &displayNameUpdate2,
			},
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate2,
				DisplayName: displayNameUpdate2,
				Description: &group.Description,
			},
		},
		{
			name: "update description",
			updateInput: datumclient.UpdateGroupInput{
				Description: &descriptionUpdate,
			},
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate2,
				DisplayName: displayNameUpdate2,
				Description: &descriptionUpdate,
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			// update group
			resp, err := client.UpdateGroup(reqCtx, group.ID, tc.updateInput)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateGroup.Group)

			// Make sure provided values match
			updatedGroup := resp.GetUpdateGroup().Group
			assert.Equal(t, tc.expectedRes, updatedGroup)
		})
	}

	owner, _ := group.Owner(reqCtx)
	(&OrganizationCleanup{OrgID: owner.ID}).MustDelete(reqCtx)
	(&GroupCleanup{GroupID: group.ID}).MustDelete(reqCtx)
}

func TestMutation_DeleteGroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	group := (&GroupBuilder{}).MustNew(reqCtx)

	reqCtx = privacy.DecisionContext(reqCtx, privacy.Allow)

	testCases := []struct {
		name     string
		groupID  string
		errorMsg string
	}{
		{
			name:    "delete group, happy path",
			groupID: group.ID,
		},
		{
			name:     "delete org, not found",
			groupID:  "tacos-tuesday",
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			// delete group
			resp, err := client.DeleteGroup(reqCtx, tc.groupID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.DeleteGroup.DeletedID)

			// make sure the deletedID matches the ID we wanted to delete
			assert.Equal(t, tc.groupID, resp.DeleteGroup.DeletedID)

			o, err := client.GetGroupByID(reqCtx, tc.groupID)

			require.Nil(t, o)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")
		})
	}

	owner, _ := group.Owner(reqCtx)
	(&OrganizationCleanup{OrgID: owner.ID}).MustDelete(reqCtx)
}
