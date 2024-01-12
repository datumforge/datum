package graphapi_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

func TestQuery_Group(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	group1 := (&GroupBuilder{Owner: org1.ID}).MustNew(reqCtx)

	listOrgs := []string{fmt.Sprintf("organization:%s", org1.ID)}
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
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.allowed)

			// second check won't happen if org does not exist
			if tc.errorMsg == "" {
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listGroups)
				// we need to check list objects even on a get to check the user
				// has access to the owner (organization of the group)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
			}

			resp, err := authClient.gc.GetGroupByID(reqCtx, tc.queryID)

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
	(&GroupCleanup{GroupID: group1.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
}

func TestQuery_GroupsNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(t, EntClient)

	reqCtx := echoContext()

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

func TestQuery_GroupsByOwner(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	org2 := (&OrganizationBuilder{}).MustNew(reqCtx)

	group1 := (&GroupBuilder{Owner: org1.ID}).MustNew(reqCtx)
	group2 := (&GroupBuilder{Owner: org2.ID}).MustNew(reqCtx)

	t.Run("Get Groups By Owner", func(t *testing.T) {
		// check tuple per org
		listOrgs := []string{fmt.Sprintf("organization:%s", org1.ID)}
		listGroups := []string{fmt.Sprintf("group:%s", group1.ID)}

		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listGroups)

		whereInput := &datumclient.GroupWhereInput{
			HasOwnerWith: []*datumclient.OrganizationWhereInput{
				{
					ID: &org1.ID,
				},
			},
		}

		resp, err := authClient.gc.GroupsWhere(reqCtx, whereInput)

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

		// Try to get groups for org not authorized to access
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listGroups)

		whereInput = &datumclient.GroupWhereInput{
			HasOwnerWith: []*datumclient.OrganizationWhereInput{
				{
					ID: &org2.ID,
				},
			},
		}

		resp, err = authClient.gc.GroupsWhere(reqCtx, whereInput)

		require.NoError(t, err)
		require.Empty(t, resp.Groups.Edges)
	})

	// delete created orgs and groups
	(&GroupCleanup{GroupID: group1.ID}).MustDelete(reqCtx)
	(&GroupCleanup{GroupID: group2.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org2.ID}).MustDelete(reqCtx)
}

func TestQuery_GroupsByOwnerNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(t, EntClient)

	reqCtx := echoContext()

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

func TestQuery_Groups(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	org2 := (&OrganizationBuilder{}).MustNew(reqCtx)

	group1 := (&GroupBuilder{Owner: org1.ID}).MustNew(reqCtx)
	group2 := (&GroupBuilder{Owner: org2.ID}).MustNew(reqCtx)
	group3 := (&GroupBuilder{Owner: org2.ID}).MustNew(reqCtx)

	t.Run("Get Groups", func(t *testing.T) {
		// check org tuples
		listOrgs := []string{fmt.Sprintf("organization:%s", org2.ID)}
		listGroups := []string{fmt.Sprintf("group:%s", group2.ID), fmt.Sprintf("group:%s", group3.ID)}

		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs) // org check comes before group check
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listGroups)

		resp, err := authClient.gc.GetAllGroups(reqCtx)

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
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, []string{}) // list orgs
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, []string{}) // list group

		resp, err = authClient.gc.GetAllGroups(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)

		// make sure no organizations are returned
		assert.Equal(t, 0, len(resp.Groups.Edges))
	})

	// delete created orgs and groups
	(&GroupCleanup{GroupID: group1.ID}).MustDelete(reqCtx)
	(&GroupCleanup{GroupID: group2.ID}).MustDelete(reqCtx)
	(&GroupCleanup{GroupID: group3.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org2.ID}).MustDelete(reqCtx)
}

func TestQuery_GroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(t, EntClient)

	reqCtx := echoContext()

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

func TestMutation_CreateGroup(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	owner1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	owner2 := (&OrganizationBuilder{}).MustNew(reqCtx)

	listObjects := []string{fmt.Sprintf("organization:%s", owner1.ID)}

	testCases := []struct {
		name        string
		groupName   string
		description string
		displayName string
		owner       string
		allowed     bool
		errorMsg    string
	}{
		{
			name:        "happy path group",
			groupName:   gofakeit.Name(),
			displayName: gofakeit.LetterN(50),
			description: gofakeit.HipsterSentence(10),
			owner:       owner1.ID,
			allowed:     true,
		},
		{
			name:        "no access to owner",
			groupName:   gofakeit.Name(),
			displayName: gofakeit.LetterN(50),
			description: gofakeit.HipsterSentence(10),
			owner:       owner2.ID,
			allowed:     false,
			errorMsg:    "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			tc := tc
			input := datumclient.CreateGroupInput{
				Name:        tc.groupName,
				Description: &tc.description,
				DisplayName: &tc.displayName,
				OwnerID:     tc.owner,
			}

			if tc.displayName != "" {
				input.DisplayName = &tc.displayName
			}

			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.allowed)

			// When calls are expected to fail, we won't ever write tuples
			if tc.errorMsg == "" {
				mockWriteTuplesAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			resp, err := authClient.gc.CreateGroup(reqCtx, input)

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

			// cleanup group
			(&GroupCleanup{GroupID: resp.CreateGroup.Group.ID}).MustDelete(reqCtx)
		})
	}

	(&OrganizationCleanup{OrgID: owner1.ID}).MustDelete(reqCtx)
}

func TestMutation_CreateGroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(t, EntClient)

	reqCtx := echoContext()

	org := (&OrganizationBuilder{}).MustNew(reqCtx)

	testCases := []struct {
		name        string
		groupName   string
		description string
		displayName string
		owner       string
		errorMsg    string
	}{
		{
			name:        "happy path group",
			groupName:   gofakeit.Name(),
			displayName: gofakeit.LetterN(50),
			description: gofakeit.HipsterSentence(10),
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

func TestMutation_UpdateGroup(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	nameUpdate := gofakeit.Name()
	displayNameUpdate := gofakeit.LetterN(40)
	descriptionUpdate := gofakeit.HipsterSentence(10)

	org := (&OrganizationBuilder{}).MustNew(reqCtx)
	group := (&GroupBuilder{Owner: org.ID}).MustNew(reqCtx)

	listObjects := []string{fmt.Sprintf("group:%s", group.ID)}

	testCases := []struct {
		name        string
		allowed     bool
		updateInput datumclient.UpdateGroupInput
		expectedRes datumclient.UpdateGroup_UpdateGroup_Group
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
			expectedRes: datumclient.UpdateGroup_UpdateGroup_Group{
				ID:          group.ID,
				Name:        nameUpdate,
				DisplayName: displayNameUpdate,
				Description: &descriptionUpdate,
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
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			// get group
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.allowed)

			if tc.errorMsg == "" {
				// update group
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.allowed)
				// check access
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			// update group
			resp, err := authClient.gc.UpdateGroup(reqCtx, group.ID, tc.updateInput)

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
		})
	}

	(&GroupCleanup{GroupID: group.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org.ID}).MustDelete(reqCtx)
}

func TestMutation_UpdateGroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(t, EntClient)

	reqCtx := echoContext()

	reqCtx = privacy.DecisionContext(reqCtx, privacy.Allow)

	group := (&GroupBuilder{}).MustNew(reqCtx)

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
			assert.Equal(t, tc.expectedRes.Name, updatedGroup.Name)
			assert.Equal(t, tc.expectedRes.DisplayName, updatedGroup.DisplayName)
			assert.Equal(t, tc.expectedRes.Description, updatedGroup.Description)
		})
	}

	owner, _ := group.Owner(reqCtx)

	(&GroupCleanup{GroupID: group.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: owner.ID}).MustDelete(reqCtx)
}

func TestMutation_DeleteGroup(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group := (&GroupBuilder{}).MustNew(reqCtx)

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
			// mock read of tuple
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.allowed)

			if tc.allowed {
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.allowed)
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.allowed)

				mockReadAny(authClient.mockCtrl, authClient.mc, reqCtx)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			// delete group
			resp, err := authClient.gc.DeleteGroup(reqCtx, tc.groupID)

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

			o, err := authClient.gc.GetGroupByID(reqCtx, tc.groupID)

			require.Nil(t, o)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")
		})
	}
}

func TestMutation_DeleteGroupNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(t, EntClient)

	reqCtx := echoContext()

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
