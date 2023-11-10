package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
)

func TestQuery_Organization(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	org1 := (&OrganizationBuilder{}).MustNew(ctx)

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
			resp, err := client.GetOrganizationByID(ctx, tc.queryID)

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

func TestQuery_Organizations(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	org1 := (&OrganizationBuilder{}).MustNew(ctx)
	org2 := (&OrganizationBuilder{}).MustNew(ctx)

	t.Run("Get Organizations", func(t *testing.T) {
		resp, err := client.GetAllOrganizations(ctx)

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
	client := graphTestClient()
	ctx := context.Background()

	parentOrg := (&OrganizationBuilder{}).MustNew(ctx)

	testCases := []struct {
		name           string
		orgName        string
		orgDescription string
		parentOrgID    string
		errorMsg       string
	}{
		{
			name:           "happy path organization",
			orgName:        gofakeit.Name(),
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
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			tc := tc
			input := datumclient.CreateOrganizationInput{
				Name:        tc.orgName,
				Description: &tc.orgDescription,
			}

			if tc.parentOrgID != "" {
				input.ParentID = &tc.parentOrgID
			}

			resp, err := client.CreateOrganization(ctx, input)

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
		})
	}
}

func TestMutation_UpdateOrganization(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	nameUpdate := gofakeit.Name()
	descriptionUpdate := gofakeit.HipsterSentence(10)
	nameUpdateLong := gofakeit.LetterN(200)

	org := (&OrganizationBuilder{}).MustNew(ctx)

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
		t.Run("Get "+tc.name, func(t *testing.T) {
			// update org
			resp, err := client.UpdateOrganization(ctx, org.ID, tc.updateInput)

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
	client := graphTestClient()
	ctx := context.Background()

	org := (&OrganizationBuilder{}).MustNew(ctx)

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
		t.Run("Get "+tc.name, func(t *testing.T) {
			// delete org
			resp, err := client.DeleteOrganization(ctx, tc.orgID)

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
		})
	}
}
