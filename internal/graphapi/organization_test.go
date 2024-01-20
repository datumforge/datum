package graphapi_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/mixin"
)

const (
	organization = "organization"
)

func TestQuery_Organization(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)

	listObjects := []string{fmt.Sprintf("%s:%s", organization, org1.ID)}

	testCases := []struct {
		name     string
		queryID  string
		expected *ent.Organization
		errorMsg string
	}{
		{
			name:     "happy path, get organization",
			queryID:  org1.ID,
			expected: org1,
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			errorMsg: "organization not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
			mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

			// second check won't happen if org does not exist
			if tc.errorMsg == "" {
				// we need to check list objects even on a get
				// because a parent could be request and that access must always be
				// checked before being returned
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			resp, err := authClient.gc.GetOrganizationByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Organization)
		})
	}

	// delete created org
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
}

func TestQuery_Organizations(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	org2 := (&OrganizationBuilder{}).MustNew(reqCtx)

	t.Run("Get Organizations", func(t *testing.T) {
		// check tuple per org
		listObjects := []string{fmt.Sprintf("organization:%s", org1.ID), fmt.Sprintf("organization:%s", org2.ID)}

		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

		resp, err := authClient.gc.GetAllOrganizations(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Organizations.Edges)

		// make sure two organizations are returned
		assert.Equal(t, 2, len(resp.Organizations.Edges))

		org1Found := false
		org2Found := false
		for _, o := range resp.Organizations.Edges {
			if o.Node.ID == org1.ID {
				org1Found = true
			} else if o.Node.ID == org2.ID {
				org2Found = true
			}
		}

		// if one of the orgs isn't found, fail the test
		if !org1Found || !org2Found {
			t.Fail()
		}

		// Check user with no relations, gets no orgs back
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, []string{})

		resp, err = authClient.gc.GetAllOrganizations(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)

		// make sure no organizations are returned
		assert.Equal(t, 0, len(resp.Organizations.Edges))
	})
}

func TestMutation_CreateOrganization(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	parentOrg := (&OrganizationBuilder{}).MustNew(reqCtx)
	parentPersonalOrg := (&OrganizationBuilder{PersonalOrg: true}).MustNew(reqCtx)

	listObjects := []string{fmt.Sprintf("organization:%s", parentOrg.ID), fmt.Sprintf("organization:%s", parentPersonalOrg.ID)}

	// setup deleted org
	orgToDelete := (&OrganizationBuilder{}).MustNew(reqCtx)
	// delete said org
	(&OrganizationCleanup{OrgID: orgToDelete.ID}).MustDelete(reqCtx)

	testCases := []struct {
		name           string
		orgName        string
		displayName    string
		orgDescription string
		parentOrgID    string
		settings       *datumclient.CreateOrganizationSettingInput
		errorMsg       string
	}{
		{
			name:           "happy path organization",
			orgName:        gofakeit.Name(),
			displayName:    gofakeit.LetterN(50),
			orgDescription: gofakeit.HipsterSentence(10),
			parentOrgID:    "", // root org
		},
		{
			name:           "happy path organization with settings",
			orgName:        gofakeit.Name(),
			displayName:    gofakeit.LetterN(50),
			orgDescription: gofakeit.HipsterSentence(10),
			settings: &datumclient.CreateOrganizationSettingInput{
				Domains: []string{"meow.datum.net"},
			},
			parentOrgID: "", // root org
		},
		{
			name:           "happy path organization with parent org",
			orgName:        gofakeit.Name(),
			orgDescription: gofakeit.HipsterSentence(10),
			parentOrgID:    parentOrg.ID,
		},
		{
			name:           "happy path organization with parent personal org",
			orgName:        gofakeit.Name(),
			orgDescription: gofakeit.HipsterSentence(10),
			parentOrgID:    parentPersonalOrg.ID,
			errorMsg:       "personal organizations are not allowed to have child organizations",
		},
		{
			name:           "empty organization name",
			orgName:        "",
			orgDescription: gofakeit.HipsterSentence(10),
			errorMsg:       "value is less than the required length",
		},
		{
			name:           "long organization name",
			orgName:        gofakeit.LetterN(161),
			orgDescription: gofakeit.HipsterSentence(10),
			errorMsg:       "value is greater than the required length",
		},
		{
			name:           "organization with no description",
			orgName:        gofakeit.Name(),
			orgDescription: "",
			parentOrgID:    parentOrg.ID,
		},
		{
			name:           "duplicate organization name",
			orgName:        parentOrg.Name,
			orgDescription: gofakeit.HipsterSentence(10),
			errorMsg:       "constraint failed",
		},
		{
			name:           "duplicate organization name, but other was deleted, should pass",
			orgName:        orgToDelete.Name,
			orgDescription: gofakeit.HipsterSentence(10),
			errorMsg:       "",
		},
		{
			name:           "duplicate display name, should be allowed",
			orgName:        gofakeit.LetterN(80),
			displayName:    parentOrg.DisplayName,
			orgDescription: gofakeit.HipsterSentence(10),
		},
		{
			name:           "display name with spaces should pass",
			orgName:        gofakeit.Name(),
			displayName:    gofakeit.Sentence(3),
			orgDescription: gofakeit.HipsterSentence(10),
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			tc := tc
			input := datumclient.CreateOrganizationInput{
				Name:        tc.orgName,
				Description: &tc.orgDescription,
			}

			if tc.displayName != "" {
				input.DisplayName = &tc.displayName
			}

			if tc.parentOrgID != "" {
				input.ParentID = &tc.parentOrgID

				// There is a check to ensure user has write access to parent org
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
				// There is a check to ensure the parent org is not a parent org
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			if tc.settings != nil {
				input.CreateOrgSettings = tc.settings
			}

			// When calls are expected to fail, we won't ever write tuples
			if tc.errorMsg == "" {
				mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			resp, err := authClient.gc.CreateOrganization(reqCtx, input)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateOrganization.Organization)

			// Make sure provided values match
			assert.Equal(t, tc.orgName, resp.CreateOrganization.Organization.Name)
			assert.Equal(t, tc.orgDescription, *resp.CreateOrganization.Organization.Description)

			if tc.parentOrgID == "" {
				assert.Nil(t, resp.CreateOrganization.Organization.Parent)
			} else {
				parent := resp.CreateOrganization.Organization.GetParent()
				assert.Equal(t, tc.parentOrgID, parent.ID)
			}

			// Ensure org settings is not null
			assert.NotNil(t, resp.CreateOrganization.Organization.Setting.ID)

			// Ensure display name is not empty
			assert.NotEmpty(t, resp.CreateOrganization.Organization.DisplayName)

			if tc.settings != nil {
				assert.Len(t, resp.CreateOrganization.Organization.Setting.Domains, 1)
			}

			// cleanup org
			(&OrganizationCleanup{OrgID: resp.CreateOrganization.Organization.ID}).MustDelete(reqCtx)
		})
	}

	(&OrganizationCleanup{OrgID: parentOrg.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: parentPersonalOrg.ID}).MustDelete(reqCtx)
}

func TestMutation_UpdateOrganization(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	nameUpdate := gofakeit.Name()
	displayNameUpdate := gofakeit.LetterN(40)
	descriptionUpdate := gofakeit.HipsterSentence(10)
	nameUpdateLong := gofakeit.LetterN(200)

	org := (&OrganizationBuilder{}).MustNew(reqCtx)
	testUser1 := (&UserBuilder{}).MustNew(reqCtx)

	listObjects := []string{fmt.Sprintf("organization:%s", org.ID)}

	testCases := []struct {
		name        string
		updateInput datumclient.UpdateOrganizationInput
		expectedRes datumclient.UpdateOrganization_UpdateOrganization_Organization
		errorMsg    string
	}{
		{
			name: "update name, happy path",
			updateInput: datumclient.UpdateOrganizationInput{
				Name: &nameUpdate,
			},
			expectedRes: datumclient.UpdateOrganization_UpdateOrganization_Organization{
				ID:          org.ID,
				Name:        nameUpdate,
				DisplayName: org.DisplayName,
				Description: &org.Description,
			},
		},
		{
			name: "add member as admin",
			updateInput: datumclient.UpdateOrganizationInput{
				Name: &nameUpdate,
				AddOrgMembers: []*datumclient.CreateOrgMembershipInput{
					{
						UserID: testUser1.ID,
						Role:   &enums.RoleAdmin,
					},
				},
			},
			expectedRes: datumclient.UpdateOrganization_UpdateOrganization_Organization{
				ID:          org.ID,
				Name:        nameUpdate,
				DisplayName: org.DisplayName,
				Description: &org.Description,
				Members: []*datumclient.UpdateOrganization_UpdateOrganization_Organization_Members{
					{
						Role:   enums.RoleAdmin,
						UserID: testUser1.ID,
					},
				},
			},
		},
		{
			name: "update description, happy path",
			updateInput: datumclient.UpdateOrganizationInput{
				Description: &descriptionUpdate,
			},
			expectedRes: datumclient.UpdateOrganization_UpdateOrganization_Organization{
				ID:          org.ID,
				Name:        nameUpdate, // this would have been updated on the prior test
				DisplayName: org.DisplayName,
				Description: &descriptionUpdate,
			},
		},
		{
			name: "update display name, happy path",
			updateInput: datumclient.UpdateOrganizationInput{
				DisplayName: &displayNameUpdate,
			},
			expectedRes: datumclient.UpdateOrganization_UpdateOrganization_Organization{
				ID:          org.ID,
				Name:        nameUpdate, // this would have been updated on the prior test
				DisplayName: displayNameUpdate,
				Description: &descriptionUpdate,
			},
		},
		{
			name: "update settings, happy path",
			updateInput: datumclient.UpdateOrganizationInput{
				Description: &descriptionUpdate,
				UpdateOrgSettings: &datumclient.UpdateOrganizationSettingInput{
					Domains: []string{"meow.datum.net", "woof.datum.net"},
				},
			},
			expectedRes: datumclient.UpdateOrganization_UpdateOrganization_Organization{
				ID:          org.ID,
				Name:        nameUpdate,        // this would have been updated on the prior test
				DisplayName: displayNameUpdate, // this would have been updated on the prior test
				Description: &descriptionUpdate,
			},
		},
		{
			name: "update name, too long",
			updateInput: datumclient.UpdateOrganizationInput{
				Name: &nameUpdateLong,
			},
			errorMsg: "value is greater than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			// mock checks of tuple
			// get organization
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
			// update organization
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
			// check access
			mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

			if tc.updateInput.AddOrgMembers != nil {
				// checks for adding orgs to ensure not a personal org
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
				mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
			}

			if tc.updateInput.UpdateOrgSettings != nil {
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			// update org
			resp, err := authClient.gc.UpdateOrganization(reqCtx, org.ID, tc.updateInput)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateOrganization.Organization)

			// Make sure provided values match
			updatedOrg := resp.GetUpdateOrganization().Organization
			assert.Equal(t, tc.expectedRes.Name, updatedOrg.Name)
			assert.Equal(t, tc.expectedRes.DisplayName, updatedOrg.DisplayName)
			assert.Equal(t, tc.expectedRes.Description, updatedOrg.Description)

			if tc.updateInput.AddOrgMembers != nil {
				// Adding a member to an org will make it 2 users, there is an owner
				// assigned to the org automatically
				assert.Len(t, updatedOrg.Members, 2)
				assert.Equal(t, tc.expectedRes.Members[0].Role, updatedOrg.Members[1].Role)
				assert.Equal(t, tc.expectedRes.Members[0].UserID, updatedOrg.Members[1].UserID)
			}

			if tc.updateInput.UpdateOrgSettings != nil {
				assert.Len(t, updatedOrg.GetSetting().Domains, 2)
			}
		})
	}

	(&OrganizationCleanup{OrgID: org.ID}).MustDelete(reqCtx)
	(&UserCleanup{UserID: testUser1.ID}).MustDelete(reqCtx)
}

func TestMutation_DeleteOrganization(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org := (&OrganizationBuilder{}).MustNew(reqCtx)

	listObjects := []string{fmt.Sprintf("organization:%s", org.ID)}

	testCases := []struct {
		name          string
		orgID         string
		accessAllowed bool
		errorMsg      string
	}{
		{
			name:          "delete org, access denied",
			orgID:         org.ID,
			accessAllowed: false,
			errorMsg:      "you are not authorized to perform this action",
		},
		{
			name:          "delete org, happy path",
			orgID:         org.ID,
			accessAllowed: true,
		},
		{
			name:          "delete org, not found",
			orgID:         "tacos-tuesday",
			accessAllowed: true,
			errorMsg:      "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			// mock read of tuple
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.accessAllowed)

			// if access is allowed, another call to `read` happens
			if tc.accessAllowed {
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.accessAllowed)

				// additional check happens when the resource is found
				if tc.errorMsg == "" {
					mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
					mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, tc.accessAllowed)
					mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
				}
			}

			// delete org
			resp, err := authClient.gc.DeleteOrganization(reqCtx, tc.orgID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.DeleteOrganization.DeletedID)

			// make sure the deletedID matches the ID we wanted to delete
			assert.Equal(t, tc.orgID, resp.DeleteOrganization.DeletedID)

			// make sure the org isn't returned
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
			if tc.errorMsg == "" {
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			o, err := authClient.gc.GetOrganizationByID(reqCtx, tc.orgID)

			require.Nil(t, o)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")

			// check that the soft delete occurred
			mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)

			ctx := mixin.SkipSoftDelete(reqCtx)

			o, err = authClient.gc.GetOrganizationByID(ctx, tc.orgID)

			require.Equal(t, o.Organization.ID, tc.orgID)
			require.NoError(t, err)
		})
	}
}

func TestMutation_OrganizationCascadeDelete(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org := (&OrganizationBuilder{}).MustNew(reqCtx)

	listOrgs := []string{fmt.Sprintf("organization:%s", org.ID)}

	group1 := (&GroupBuilder{Owner: org.ID}).MustNew(reqCtx)

	listGroups := []string{fmt.Sprintf("group:%s", group1.ID)}

	// mocks checks for all calls
	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)

	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listGroups)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listOrgs)

	// mock writes to delete member of org
	mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)

	// delete org
	resp, err := authClient.gc.DeleteOrganization(reqCtx, org.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteOrganization.DeletedID)

	// make sure the deletedID matches the ID we wanted to delete
	assert.Equal(t, org.ID, resp.DeleteOrganization.DeletedID)

	o, err := authClient.gc.GetOrganizationByID(reqCtx, org.ID)

	require.Nil(t, o)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	g, err := authClient.gc.GetGroupByID(reqCtx, group1.ID)

	require.Nil(t, g)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	// allow after tuples have been deleted
	ctx := privacy.DecisionContext(reqCtx, privacy.Allow)

	ctx = mixin.SkipSoftDelete(ctx)

	o, err = authClient.gc.GetOrganizationByID(ctx, org.ID)

	require.NoError(t, err)
	require.Equal(t, o.Organization.ID, org.ID)

	// allow after tuples have been deleted
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	g, err = authClient.gc.GetGroupByID(ctx, group1.ID)

	require.NoError(t, err)
	require.Equal(t, g.Group.ID, group1.ID)
}

func TestMutation_CreateOrganizationTransaction(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	t.Run("Create should not write if FGA transaction fails", func(t *testing.T) {
		input := datumclient.CreateOrganizationInput{
			Name: gofakeit.Name(),
		}

		fgaErr := errors.New("unable to create relationship") //nolint:goerr113
		mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, fgaErr)

		resp, err := authClient.gc.CreateOrganization(reqCtx, input)

		require.Error(t, err)
		require.Empty(t, resp)

		// Make sure the org was not added to the database (check without auth)
		listObjects := []string{fmt.Sprintf("%s:%s", organization, "test")}
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
		mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

		ctx := privacy.DecisionContext(reqCtx, privacy.Allow)

		orgs, err := authClient.gc.GetAllOrganizations(ctx)
		require.NoError(t, err)

		for _, o := range orgs.Organizations.Edges {
			if o.Node.Name == input.Name {
				t.Errorf("org found that should not have been created due to FGA error")
			}
		}
	})
}
