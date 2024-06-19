package graphapi_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

func (suite *GraphTestSuite) TestQueryGroup() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group1 := (&GroupBuilder{client: suite.client}).MustNew(reqCtx, t)

	listGroups := []string{fmt.Sprintf("group:%s", group1.ID)}

	testCases := []struct {
		name     string
		queryID  string
		allowed  bool
		expected *ent.Group
		errorMsg string
	}{
		{
			name:     "happy path group",
			allowed:  true,
			queryID:  group1.ID,
			expected: group1,
		},
		{
			name:     "no access",
			allowed:  false,
			queryID:  group1.ID,
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			// second check won't happen if org does not exist
			if tc.errorMsg == "" {
				mock_fga.ListTimes(t, suite.client.fga, listGroups, 1)
			}

			resp, err := suite.client.datum.GetGroupByID(reqCtx, tc.queryID)

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

	// delete created org and group
	(&GroupCleanup{client: suite.client, GroupID: group1.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryGroupsByOwner() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org1.ID)
	require.NoError(t, err)

	group1 := (&GroupBuilder{client: suite.client, Owner: org1.ID}).MustNew(reqCtx, t)
	group2 := (&GroupBuilder{client: suite.client, Owner: org2.ID}).MustNew(reqCtx, t)

	t.Run("Get Groups By Owner", func(t *testing.T) {
		defer mock_fga.ClearMocks(suite.client.fga)

		// check tuple per org
		listGroups := []string{fmt.Sprintf("group:%s", group1.ID)}

		mock_fga.ListAny(t, suite.client.fga, listGroups)

		whereInput := &datumclient.GroupWhereInput{
			HasOwnerWith: []*datumclient.OrganizationWhereInput{
				{
					ID: &org1.ID,
				},
			},
		}

		resp, err := suite.client.datum.GroupsWhere(reqCtx, whereInput)

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

		whereInput = &datumclient.GroupWhereInput{
			HasOwnerWith: []*datumclient.OrganizationWhereInput{
				{
					ID: &org2.ID,
				},
			},
		}

		resp, err = suite.client.datum.GroupsWhere(reqCtx, whereInput)

		require.NoError(t, err)
		require.Empty(t, resp.Groups.Edges)
	})

	// delete created orgs and groups
	reqCtx2, err := auth.NewTestContextWithOrgID(testUser.ID, org2.ID)
	require.NoError(t, err)

	(&OrganizationCleanup{client: suite.client, OrgID: org1.ID}).MustDelete(reqCtx, t)
	(&OrganizationCleanup{client: suite.client, OrgID: org2.ID}).MustDelete(reqCtx2, t)
}

func (suite *GraphTestSuite) TestQueryGroups() {
	t := suite.T()

	// setup user context
	reqCtx2, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx2, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx2, t)

	reqCtx1, err := auth.NewTestContextWithOrgID(testUser.ID, org1.ID)
	require.NoError(t, err)

	reqCtx2, err = auth.NewTestContextWithOrgID(testUser.ID, org2.ID)
	require.NoError(t, err)

	group1 := (&GroupBuilder{client: suite.client, Owner: org1.ID}).MustNew(reqCtx1, t)
	group2 := (&GroupBuilder{client: suite.client, Owner: org2.ID}).MustNew(reqCtx2, t)
	group3 := (&GroupBuilder{client: suite.client, Owner: org2.ID}).MustNew(reqCtx2, t)

	t.Run("Get Groups", func(t *testing.T) {
		defer mock_fga.ClearMocks(suite.client.fga)

		// check org tuples
		listGroups := []string{fmt.Sprintf("group:%s", group2.ID), fmt.Sprintf("group:%s", group3.ID)}

		mock_fga.ListTimes(t, suite.client.fga, listGroups, 1)

		resp, err := suite.client.datum.GetAllGroups(reqCtx2)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Groups.Edges)

		// make sure two organizations are returned (group 2 and group 3)
		assert.Equal(t, 2, len(resp.Groups.Edges))

		group1Found := false
		group2Found := false
		group3Found := false

		for _, o := range resp.Groups.Edges {
			switch id := o.Node.ID; id {
			case group1.ID:
				group1Found = true
			case group2.ID:
				group2Found = true
			case group3.ID:
				group3Found = true
			}
		}

		// if one of the groups isn't found, fail the test
		if !group2Found || !group3Found {
			t.Fail()
		}

		// if group 1 (which belongs to an unauthorized org) is found, fail the test
		if group1Found {
			t.Fail()
		}

		// Check user with no relations, gets no groups back
		mock_fga.ListAny(t, suite.client.fga, []string{})

		resp, err = suite.client.datum.GetAllGroups(reqCtx2)

		require.NoError(t, err)
		require.NotNil(t, resp)

		// make sure no organizations are returned
		assert.Equal(t, 0, len(resp.Groups.Edges))
	})

	// delete created orgs and groups
	(&GroupCleanup{client: suite.client, GroupID: group1.ID}).MustDelete(reqCtx1, t)
	(&GroupCleanup{client: suite.client, GroupID: group2.ID}).MustDelete(reqCtx2, t)
	(&GroupCleanup{client: suite.client, GroupID: group3.ID}).MustDelete(reqCtx2, t)
	(&OrganizationCleanup{client: suite.client, OrgID: org1.ID}).MustDelete(reqCtx1, t)
	(&OrganizationCleanup{client: suite.client, OrgID: org2.ID}).MustDelete(reqCtx2, t)
}

func (suite *GraphTestSuite) TestMutationCreateGroup() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	owner1 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	owner2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	listObjects := []string{fmt.Sprintf("organization:%s", owner1.ID)}

	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, owner1.ID)
	require.NoError(t, err)

	testCases := []struct {
		name        string
		groupName   string
		description string
		displayName string
		owner       string
		settings    *datumclient.CreateGroupSettingInput
		allowed     bool
		check       bool
		errorMsg    string
	}{
		{
			name:        "happy path group",
			groupName:   gofakeit.Name(),
			displayName: gofakeit.LetterN(50),
			description: gofakeit.HipsterSentence(10),
			owner:       owner1.ID,
			allowed:     true,
			check:       true,
		},
		{
			name:        "happy path group with settings",
			groupName:   gofakeit.Name(),
			displayName: gofakeit.LetterN(50),
			description: gofakeit.HipsterSentence(10),
			owner:       owner1.ID,
			settings: &datumclient.CreateGroupSettingInput{
				JoinPolicy: &enums.JoinPolicyInviteOnly,
			},
			allowed: true,
			check:   true,
		},
		{
			name:        "no access to owner",
			groupName:   gofakeit.Name(),
			displayName: gofakeit.LetterN(50),
			description: gofakeit.HipsterSentence(10),
			owner:       owner2.ID,
			allowed:     false,
			check:       true,
			errorMsg:    "not authorized",
		},
		{
			name:      "happy path group, minimum fields",
			groupName: gofakeit.Name(),
			owner:     owner1.ID,
			allowed:   true,
			check:     true,
		},
		{
			name:     "missing name",
			owner:    owner1.ID,
			errorMsg: "validator failed",
			allowed:  true,
			check:    true,
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			// clear mocks at end of each test
			defer mock_fga.ClearMocks(suite.client.fga)

			tc := tc
			input := datumclient.CreateGroupInput{
				Name:        tc.groupName,
				Description: &tc.description,
				DisplayName: &tc.displayName,
				OwnerID:     &tc.owner,
			}

			if tc.displayName != "" {
				input.DisplayName = &tc.displayName
			}

			if tc.settings != nil {
				input.CreateGroupSettings = tc.settings
			}

			if tc.check {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			// When calls are expected to fail, we won't ever write tuples
			if tc.errorMsg == "" {
				mock_fga.WriteAny(t, suite.client.fga)
				mock_fga.ListAny(t, suite.client.fga, listObjects)
			}

			resp, err := suite.client.datum.CreateGroup(reqCtx, input)

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
			assert.Equal(t, tc.owner, resp.CreateGroup.Group.Owner.ID)

			if tc.displayName != "" {
				assert.Equal(t, tc.displayName, resp.CreateGroup.Group.DisplayName)
			} else {
				// display name defaults to the name if not set
				assert.Equal(t, tc.groupName, resp.CreateGroup.Group.DisplayName)
			}

			if tc.settings != nil {
				assert.Equal(t, resp.CreateGroup.Group.Setting.JoinPolicy, enums.JoinPolicyInviteOnly)
			}

			// cleanup group
			(&GroupCleanup{client: suite.client, GroupID: resp.CreateGroup.Group.ID}).MustDelete(reqCtx, t)
		})
	}

	(&OrganizationCleanup{client: suite.client, OrgID: owner1.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestMutationUpdateGroup() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	nameUpdate := gofakeit.Name()
	displayNameUpdate := gofakeit.LetterN(40)
	descriptionUpdate := gofakeit.HipsterSentence(10)

	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org.ID)
	require.NoError(t, err)

	group := (&GroupBuilder{client: suite.client, Owner: org.ID}).MustNew(reqCtx, t)

	om := (&OrgMemberBuilder{client: suite.client, OrgID: org.ID}).MustNew(reqCtx, t)

	// setup auth for the tests
	listObjects := []string{fmt.Sprintf("group:%s", group.ID)}

	testCases := []struct {
		name        string
		allowed     bool
		updateInput datumclient.UpdateGroupInput
		expectedRes datumclient.UpdateGroup_UpdateGroup_Group
		list        bool
		errorMsg    string
	}{
		{
			name:    "update name, happy path",
			allowed: true,
			updateInput: datumclient.UpdateGroupInput{
				Name:        &nameUpdate,
				DisplayName: &displayNameUpdate,
				Description: &descriptionUpdate,
			},
			list: true,
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate,
				DisplayName: displayNameUpdate,
				Description: &descriptionUpdate,
			},
		},
		{
			name:    "add user as admin",
			allowed: true,
			updateInput: datumclient.UpdateGroupInput{
				AddGroupMembers: []*datumclient.CreateGroupMembershipInput{
					{
						UserID: om.UserID,
						Role:   &enums.RoleAdmin,
					},
				},
			},
			list: true,
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate,
				DisplayName: displayNameUpdate,
				Description: &descriptionUpdate,
				Members: []*datumclient.UpdateGroup_UpdateGroup_Group_Members{
					{
						Role: enums.RoleAdmin,
						User: datumclient.UpdateGroup_UpdateGroup_Group_Members_User{
							ID: om.UserID,
						},
					},
				},
			},
		},
		{
			name:    "update settings, happy path",
			allowed: true,
			updateInput: datumclient.UpdateGroupInput{
				UpdateGroupSettings: &datumclient.UpdateGroupSettingInput{
					JoinPolicy: &enums.JoinPolicyOpen,
				},
			},
			list: true,
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate,
				DisplayName: displayNameUpdate,
				Description: &descriptionUpdate,
				Setting: datumclient.UpdateGroup_UpdateGroup_Group_Setting{
					JoinPolicy: enums.JoinPolicyOpen,
				},
			},
		},
		{
			name:    "no access",
			allowed: false,
			updateInput: datumclient.UpdateGroupInput{
				Name:        &nameUpdate,
				DisplayName: &displayNameUpdate,
				Description: &descriptionUpdate,
			},
			list:     false,
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			if tc.list {
				mock_fga.ListAny(t, suite.client.fga, listObjects)
			}

			if tc.updateInput.AddGroupMembers != nil && tc.errorMsg == "" {
				mock_fga.WriteAny(t, suite.client.fga)
			}

			// update group
			resp, err := suite.client.datum.UpdateGroup(reqCtx, group.ID, tc.updateInput)

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
			assert.Equal(t, tc.expectedRes.Name, updatedGroup.Name)
			assert.Equal(t, tc.expectedRes.DisplayName, updatedGroup.DisplayName)
			assert.Equal(t, tc.expectedRes.Description, updatedGroup.Description)

			if tc.updateInput.AddGroupMembers != nil {
				// Adding a member to an group will make it 2 users, there is an admin
				// assigned to the group automatically
				assert.Len(t, updatedGroup.Members, 2)
				assert.Equal(t, tc.expectedRes.Members[0].Role, updatedGroup.Members[1].Role)
				assert.Equal(t, tc.expectedRes.Members[0].User.ID, updatedGroup.Members[1].User.ID)
			}

			if tc.updateInput.UpdateGroupSettings != nil {
				assert.Equal(t, updatedGroup.GetSetting().JoinPolicy, enums.JoinPolicyOpen)
			}
		})
	}

	(&GroupCleanup{client: suite.client, GroupID: group.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestMutationDeleteGroup() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group := (&GroupBuilder{client: suite.client}).MustNew(reqCtx, t)

	listObjects := []string{fmt.Sprintf("group:%s", group.ID)}

	testCases := []struct {
		name     string
		groupID  string
		allowed  bool
		errorMsg string
	}{
		{
			name:    "delete group, happy path",
			allowed: true,
			groupID: group.ID,
		},
		{
			name:     "delete group, no access",
			allowed:  false,
			groupID:  group.ID,
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// mock read of tuple
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			if tc.allowed {
				mock_fga.ReadAny(t, suite.client.fga)
				mock_fga.ListAny(t, suite.client.fga, listObjects)
				mock_fga.WriteAny(t, suite.client.fga)
			}

			// delete group
			resp, err := suite.client.datum.DeleteGroup(reqCtx, tc.groupID)

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

			o, err := suite.client.datum.GetGroupByID(reqCtx, tc.groupID)

			require.Nil(t, o)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")
		})
	}
}
