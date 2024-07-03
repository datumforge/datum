package graphapi_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

func (suite *GraphTestSuite) TestQueryContact() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	contact := (&ContactBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		allowed  bool
		expected *ent.Contact
		errorMsg string
	}{
		{
			name:     "happy path contact",
			allowed:  true,
			queryID:  contact.ID,
			expected: contact,
		},
		{
			name:     "no access",
			allowed:  false,
			queryID:  contact.ID,
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.GetContactByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Contact)
		})
	}

	// delete created org and contact
	(&ContactCleanup{client: suite.client, ID: contact.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryContacts() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	_ = (&ContactBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&ContactBuilder{client: suite.client}).MustNew(reqCtx, t)

	otherUser := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	otherCtx, err := userContextWithID(otherUser.ID)
	require.NoError(t, err)

	testCases := []struct {
		name            string
		ctx             context.Context
		expectedResults int
	}{
		{
			name:            "happy path",
			ctx:             reqCtx,
			expectedResults: 2,
		},
		{
			name:            "another user, no contacts should be returned",
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.GetAllContacts(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.Contacts.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateContact() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	testCases := []struct {
		name        string
		request     datumclient.CreateContactInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: datumclient.CreateContactInput{
				FullName: "Aemond Targaryen",
			},
			allowed: true,
		},
		{
			name: "happy path, all input",
			request: datumclient.CreateContactInput{
				FullName:    "Aemond Targaryen",
				Email:       lo.ToPtr("atargarygen@dragon.com"),
				PhoneNumber: lo.ToPtr(gofakeit.Phone()),
				Title:       lo.ToPtr("Prince of the Targaryen Dynasty"),
				Company:     lo.ToPtr("Targaryen Dynasty"),
				Status:      &enums.UserStatusOnboarding,
			},
			allowed: true,
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateContactInput{
				FullName: "Halaena Targaryen",
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on contact",
		},
		{
			name: "missing required field, name",
			request: datumclient.CreateContactInput{
				Email: lo.ToPtr("atargarygen@dragon.com"),
			},
			allowed:     true,
			expectedErr: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.CreateContact(reqCtx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.request.FullName, resp.CreateContact.Contact.FullName)

			if tc.request.Email == nil {
				assert.Empty(t, resp.CreateContact.Contact.Email)
			} else {
				assert.Equal(t, *tc.request.Email, *resp.CreateContact.Contact.Email)
			}

			if tc.request.PhoneNumber == nil {
				assert.Empty(t, resp.CreateContact.Contact.PhoneNumber)
			} else {
				assert.Equal(t, *tc.request.PhoneNumber, *resp.CreateContact.Contact.PhoneNumber)
			}

			if tc.request.Address == nil {
				assert.Empty(t, resp.CreateContact.Contact.Address)
			} else {
				assert.Equal(t, *tc.request.Address, *resp.CreateContact.Contact.Address)
			}

			if tc.request.Title == nil {
				assert.Empty(t, resp.CreateContact.Contact.Title)
			} else {
				assert.Equal(t, *tc.request.Title, *resp.CreateContact.Contact.Title)
			}

			if tc.request.Company == nil {
				assert.Empty(t, resp.CreateContact.Contact.Company)
			} else {
				assert.Equal(t, *tc.request.Company, *resp.CreateContact.Contact.Company)
			}

			// status should default to active
			if tc.request.Status == nil {
				assert.Equal(t, enums.UserStatusActive, resp.CreateContact.Contact.Status)
			} else {
				assert.Equal(t, *tc.request.Status, resp.CreateContact.Contact.Status)
			}

		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateContact() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	contact := (&ContactBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     datumclient.UpdateContactInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update name",
			request: datumclient.UpdateContactInput{
				FullName: lo.ToPtr("Alicent Hightower"),
			},
			allowed: true,
		},
		{
			name: "update phone number",
			request: datumclient.UpdateContactInput{
				PhoneNumber: lo.ToPtr(gofakeit.Phone()),
			},
			allowed: true,
		},
		{
			name: "update status",
			request: datumclient.UpdateContactInput{
				Status: &enums.UserStatusInactive,
			},
			allowed: true,
		},
		{
			name: "update email",
			request: datumclient.UpdateContactInput{
				Email: lo.ToPtr("a.hightower@dragon.net"),
			},
			allowed: true,
		},
		{
			name: "update phone number, invalid",
			request: datumclient.UpdateContactInput{
				PhoneNumber: lo.ToPtr("not a phone number"),
			},
			allowed:     true,
			expectedErr: rout.InvalidField("phone_number").Error(),
		},
		{
			name: "update email, invalid",
			request: datumclient.UpdateContactInput{
				Email: lo.ToPtr("a.hightower"),
			},
			allowed:     true,
			expectedErr: "validator failed for field",
		},
		{
			name: "update title",
			request: datumclient.UpdateContactInput{
				Title: lo.ToPtr("Queen of the Seven Kingdoms"),
			},
			allowed: true,
		},
		{
			name: "update company",
			request: datumclient.UpdateContactInput{
				Company: lo.ToPtr("House Targaryen"),
			},
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateContactInput{
				Company: lo.ToPtr("House Hightower"),
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on contact",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.UpdateContact(reqCtx, contact.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			if tc.request.PhoneNumber != nil {
				assert.Equal(t, *tc.request.PhoneNumber, *resp.UpdateContact.Contact.PhoneNumber)
			}

			if tc.request.Email != nil {
				assert.Equal(t, *tc.request.Email, *resp.UpdateContact.Contact.Email)
			}

			if tc.request.FullName != nil {
				assert.Equal(t, *tc.request.FullName, resp.UpdateContact.Contact.FullName)
			}

			if tc.request.Title != nil {
				assert.Equal(t, *tc.request.Title, *resp.UpdateContact.Contact.Title)
			}

			if tc.request.Company != nil {
				assert.Equal(t, *tc.request.Company, *resp.UpdateContact.Contact.Company)
			}

			if tc.request.Status != nil {
				assert.Equal(t, *tc.request.Status, resp.UpdateContact.Contact.Status)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteContact() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	contact := (&ContactBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		allowed     bool
		checkAccess bool
		expectedErr string
	}{
		{
			name:        "not allowed to delete",
			idToDelete:  contact.ID,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on contact",
		},
		{
			name:        "happy path, delete contact",
			idToDelete:  contact.ID,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "contact already deleted, not found",
			idToDelete:  contact.ID,
			checkAccess: false,
			allowed:     true,
			expectedErr: "contact not found",
		},
		{
			name:        "unknown contact, not found",
			idToDelete:  ulids.New().String(),
			checkAccess: false,
			allowed:     true,
			expectedErr: "contact not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization if contact exists
			if tc.checkAccess {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			resp, err := suite.client.datum.DeleteContact(reqCtx, contact.ID)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, contact.ID, resp.DeleteContact.DeletedID)
		})
	}
}
