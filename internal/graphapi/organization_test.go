package graphapi_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/mixin"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
	auth "github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func TestQuery_Organization(t *testing.T) {
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	// Setup Test Graph Client
	client := graphTestClient(entClient)

	sub := ulids.New().String()

	ec, err := auth.NewTestContextWithValidUser(sub)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echocontext.EchoContextKey, echoContext)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, entClient)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	listObjects := []string{fmt.Sprintf("organization:%s", org1.ID)}

	testCases := []struct {
		name     string
		queryID  string
		expected *ent.Organization
		errorMsg string
	}{
		{
			name:     "happy path organization",
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
			mockCheckAny(mockCtrl, mc, reqCtx, true)
			mockListAny(mockCtrl, mc, reqCtx, listObjects)

			// second check won't happen if org does not exist
			if tc.errorMsg == "" {
				// we need to check list objects even on a get
				// because a parent could be request and that access must always be
				// checked before being returned
				mockListAny(mockCtrl, mc, reqCtx, listObjects)
				mockListAny(mockCtrl, mc, reqCtx, listObjects)
			}

			resp, err := client.GetOrganizationByID(reqCtx, tc.queryID)

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

func TestQuery_OrganizationsNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, EntClient)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	org2 := (&OrganizationBuilder{ParentOrgID: org1.ParentOrganizationID}).MustNew(reqCtx)

	t.Run("Get Organizations", func(t *testing.T) {
		resp, err := client.GetAllOrganizations(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Organizations.Edges)

		// make sure at least two organizations are returned
		assert.GreaterOrEqual(t, len(resp.Organizations.Edges), 2)

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
	})

	// delete created orgs
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
	(&OrganizationCleanup{OrgID: org2.ID}).MustDelete(reqCtx)
}

func TestQuery_OrganizationsAuth(t *testing.T) {
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	// Setup Test Graph Client
	client := graphTestClient(entClient)

	sub := ulids.New().String()

	ec, err := auth.NewTestContextWithValidUser(sub)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echocontext.EchoContextKey, echoContext)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, entClient)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	org2 := (&OrganizationBuilder{}).MustNew(reqCtx)

	t.Run("Get Organizations", func(t *testing.T) {
		// check tuple per org
		listObjects := []string{fmt.Sprintf("organization:%s", org1.ID), fmt.Sprintf("organization:%s", org2.ID)}

		mockListAny(mockCtrl, mc, reqCtx, listObjects)
		mockListAny(mockCtrl, mc, reqCtx, listObjects)
		mockListAny(mockCtrl, mc, reqCtx, listObjects)
		mockListAny(mockCtrl, mc, reqCtx, listObjects)
		mockListAny(mockCtrl, mc, reqCtx, listObjects)

		resp, err := client.GetAllOrganizations(reqCtx)

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
		mockListAny(mockCtrl, mc, reqCtx, []string{})

		resp, err = client.GetAllOrganizations(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)

		// make sure no organizations are returned
		assert.Equal(t, 0, len(resp.Organizations.Edges))
	})
}

func TestMutation_CreateOrganization(t *testing.T) {
	// Add Authz Client Mock
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	// Setup Test Graph Client
	client := graphTestClient(entClient)

	// Setup echo context
	sub := ulids.New().String()

	ec, err := auth.NewTestContextWithValidUser(sub)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echocontext.EchoContextKey, echoContext)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, entClient)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	parentOrg := (&OrganizationBuilder{}).MustNew(reqCtx)

	listObjects := []string{fmt.Sprintf("organization:%s", parentOrg.ID)}

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
			name:           "happy path organization with parent org",
			orgName:        gofakeit.Name(),
			orgDescription: gofakeit.HipsterSentence(10),
			parentOrgID:    parentOrg.ID,
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
			errorMsg:       "UNIQUE constraint failed",
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
				mockCheckAny(mockCtrl, mc, reqCtx, true)
			}

			// When calls are expected to fail, we won't ever write tuples
			if tc.errorMsg == "" {
				mockWriteTuplesAny(mockCtrl, mc, reqCtx, nil)
				mockListAny(mockCtrl, mc, reqCtx, listObjects)
				mockListAny(mockCtrl, mc, reqCtx, listObjects)
			}

			resp, err := client.CreateOrganization(reqCtx, input)

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

			// cleanup org
			(&OrganizationCleanup{OrgID: resp.CreateOrganization.Organization.ID}).MustDelete(reqCtx)
		})
	}

	(&OrganizationCleanup{OrgID: parentOrg.ID}).MustDelete(reqCtx)
}

func TestMutation_CreateOrganizationNoAuth(t *testing.T) {
	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, EntClient)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	parentOrg := (&OrganizationBuilder{}).MustNew(reqCtx)

	testCases := []struct {
		name           string
		orgName        string
		displayName    string
		orgDescription string
		parentOrgID    string
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
			name:           "happy path organization with parent org",
			orgName:        gofakeit.Name(),
			orgDescription: gofakeit.HipsterSentence(10),
			parentOrgID:    parentOrg.ID,
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
			}

			resp, err := client.CreateOrganization(reqCtx, input)

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

			// cleanup org
			(&OrganizationCleanup{OrgID: resp.CreateOrganization.Organization.ID}).MustDelete(reqCtx)
		})
	}

	(&OrganizationCleanup{OrgID: parentOrg.ID}).MustDelete(reqCtx)
}

func TestMutation_UpdateOrganization(t *testing.T) {
	// Add Authz Client Mock
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	// Setup Test Graph Client
	client := graphTestClient(entClient)

	// Setup echo context
	sub := ulids.New().String()

	ec, err := auth.NewTestContextWithValidUser(sub)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echocontext.EchoContextKey, echoContext)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, entClient)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	nameUpdate := gofakeit.Name()
	displayNameUpdate := gofakeit.LetterN(40)
	descriptionUpdate := gofakeit.HipsterSentence(10)
	nameUpdateLong := gofakeit.LetterN(200)

	org := (&OrganizationBuilder{}).MustNew(reqCtx)
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
			mockCheckAny(mockCtrl, mc, reqCtx, true)
			// update organization
			mockCheckAny(mockCtrl, mc, reqCtx, true)
			// check access
			mockListAny(mockCtrl, mc, reqCtx, listObjects)

			// update org
			resp, err := client.UpdateOrganization(reqCtx, org.ID, tc.updateInput)

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
			assert.Equal(t, tc.expectedRes, updatedOrg)
		})
	}

	(&OrganizationCleanup{OrgID: org.ID}).MustDelete(reqCtx)
}

func TestMutation_DeleteOrganization(t *testing.T) {
	// Add Authz Client Mock
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	// Setup Test Graph Client
	client := graphTestClient(entClient)

	// Setup echo context
	sub := ulids.New().String()

	ec, err := auth.NewTestContextWithValidUser(sub)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echocontext.EchoContextKey, echoContext)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, entClient)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

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
			mockCheckAny(mockCtrl, mc, reqCtx, tc.accessAllowed)

			// if access is allowed, another call to `read` happens
			if tc.accessAllowed {
				mockCheckAny(mockCtrl, mc, reqCtx, tc.accessAllowed)
				mockReadAny(mockCtrl, mc, reqCtx)

				// additional check happens when the resource is found
				if tc.errorMsg == "" {
					mockListAny(mockCtrl, mc, reqCtx, listObjects)
					mockCheckAny(mockCtrl, mc, reqCtx, tc.accessAllowed)
				}
			}

			// delete org
			resp, err := client.DeleteOrganization(reqCtx, tc.orgID)

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
			mockCheckAny(mockCtrl, mc, reqCtx, true)
			if tc.errorMsg == "" {
				mockListAny(mockCtrl, mc, reqCtx, listObjects)
				mockListAny(mockCtrl, mc, reqCtx, listObjects)
				mockListAny(mockCtrl, mc, reqCtx, listObjects)
			}

			o, err := client.GetOrganizationByID(reqCtx, tc.orgID)

			require.Nil(t, o)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")

			// check that the soft delete occurred
			mockCheckAny(mockCtrl, mc, reqCtx, true)

			ctx := mixin.SkipSoftDelete(reqCtx)

			o, err = client.GetOrganizationByID(ctx, tc.orgID)

			require.Equal(t, o.Organization.ID, tc.orgID)
			require.NoError(t, err)
		})
	}
}

func TestMutation_CascadeDelete(t *testing.T) {
	client := graphTestClientNoAuth(EntClient)

	ec := echocontext.NewTestEchoContext()

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, EntClient)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	org := (&OrganizationBuilder{}).MustNew(reqCtx)

	group1 := (&GroupBuilder{Owner: org.ID}).MustNew(reqCtx)

	// delete org
	resp, err := client.DeleteOrganization(reqCtx, org.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteOrganization.DeletedID)

	// make sure the deletedID matches the ID we wanted to delete
	assert.Equal(t, org.ID, resp.DeleteOrganization.DeletedID)

	o, err := client.GetOrganizationByID(reqCtx, org.ID)

	require.Nil(t, o)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	g, err := client.GetGroupByID(reqCtx, group1.ID)

	require.Nil(t, g)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	ctx := mixin.SkipSoftDelete(reqCtx)

	o, err = client.GetOrganizationByID(ctx, org.ID)

	require.Equal(t, o.Organization.ID, org.ID)
	require.NoError(t, err)

	g, err = client.GetGroupByID(ctx, group1.ID)

	require.Equal(t, g.Group.ID, group1.ID)
	require.NoError(t, err)
}

func TestMutation_CreateOrganizationTransaction(t *testing.T) {
	// Add Authz Client Mock
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	// Setup Test Graph Client
	client := graphTestClient(entClient)

	// Setup echo context
	sub := ulids.New().String()

	ec, err := auth.NewTestContextWithValidUser(sub)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echocontext.EchoContextKey, echoContext)

	// add client to context for transactional client
	reqCtx = ent.NewContext(reqCtx, entClient)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	t.Run("Create should not write if FGA transaction fails", func(t *testing.T) {
		input := datumclient.CreateOrganizationInput{
			Name: gofakeit.Name(),
		}

		fgaErr := errors.New("unable to create relationship") //nolint:goerr113
		mockWriteTuplesAny(mockCtrl, mc, reqCtx, fgaErr)

		resp, err := client.CreateOrganization(reqCtx, input)

		require.Error(t, err)
		require.Empty(t, resp)

		// Make sure the org was not added to the database (check without auth)
		clientNoAuth := graphTestClientNoAuth(EntClient)

		ec := echocontext.NewTestEchoContext()

		reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

		orgs, err := clientNoAuth.GetAllOrganizations(reqCtx)
		require.NoError(t, err)

		for _, o := range orgs.Organizations.Edges {
			if o.Node.Name == input.Name {
				t.Errorf("org found that should not have been created due to FGA error")
			}
		}
	})
}
