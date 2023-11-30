package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/echox"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/mixin"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
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

	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx, mockCtrl, mc)

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
}

func TestQuery_OrganizationsNoAuth(t *testing.T) {
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	// Setup Test Graph Client Without Auth
	client := graphTestClientNoAuth(entClient)

	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx, mockCtrl, mc)
	org2 := (&OrganizationBuilder{}).MustNew(reqCtx, mockCtrl, mc)

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
	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	parentOrg := (&OrganizationBuilder{}).MustNew(reqCtx, mockCtrl, mc)

	// setup deleted org
	orgToDelete := (&OrganizationBuilder{}).MustNew(reqCtx, mockCtrl, mc)

	mockCheckAny(mockCtrl, mc, reqCtx, true)
	mockCheckAny(mockCtrl, mc, reqCtx, true)

	if _, err = client.DeleteOrganization(reqCtx, orgToDelete.ID); err != nil {
		t.Errorf("error deleting test org")
		t.Fail()
	}

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
			name:           "display name with spaces should fail",
			orgName:        gofakeit.Name(),
			displayName:    gofakeit.Sentence(3),
			orgDescription: gofakeit.HipsterSentence(10),
			errorMsg:       "field should not contain spaces",
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

			// When calls are expected to fail, we won't ever write tuples
			if tc.errorMsg == "" {
				mockWriteTuplesAny(mockCtrl, mc, reqCtx, nil)
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
			// TODO - come back to parent orgs

			// if tc.parentOrgID == "" {
			// 	// assert.Nil(t, resp.CreateOrganization.Organization.Parent)
			// } else {
			// 	// parent := resp.CreateOrganization.Organization.GetParent()
			// 	assert.Equal(t, tc.parentOrgID, parent.ID)
			// }
		})
	}
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
	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	nameUpdate := gofakeit.Name()
	displayNameUpdate := gofakeit.LetterN(40)
	descriptionUpdate := gofakeit.HipsterSentence(10)
	nameUpdateLong := gofakeit.LetterN(200)

	org := (&OrganizationBuilder{}).MustNew(reqCtx, mockCtrl, mc)

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
				DisplayName: "unknown", // this is the default if not set
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
				DisplayName: "unknown",  // this is the default if not set
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
	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	org := (&OrganizationBuilder{}).MustNew(reqCtx, mockCtrl, mc)

	testCases := []struct {
		name     string
		orgID    string
		errorMsg string
	}{
		{
			name:  "delete org, happy path",
			orgID: org.ID,
		},
		{
			name:     "delete org, not found",
			orgID:    "tacos-tuesday",
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			// mock read of tuple
			mockCheckAny(mockCtrl, mc, reqCtx, true)
			mockCheckAny(mockCtrl, mc, reqCtx, true)

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
