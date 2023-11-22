package fga

import (
	"context"
	"errors"
	"testing"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_client "github.com/datumforge/datum/internal/fga/mocks"
)

func Test_EntityString(t *testing.T) {
	memberRelation := Relation("member")

	testCases := []struct {
		name        string
		entity      Entity
		expectedRes string
	}{
		{
			name: "relationship empty",
			entity: Entity{
				Kind:       "user",
				Identifier: "bz0yOLsL460V-6L9HauX4",
				Relation:   "",
			},
			expectedRes: "user:bz0yOLsL460V-6L9HauX4",
		},
		{
			name: "relationship member",
			entity: Entity{
				Kind:       "organization",
				Identifier: "yKreKfzq3-iG-rhj0N9o9",
				Relation:   memberRelation,
			},
			expectedRes: "organization:yKreKfzq3-iG-rhj0N9o9#member",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			res := tc.entity.String()

			// result should never be empty
			assert.NotEmpty(t, res)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func Test_ParseEntity(t *testing.T) {
	memberRelation := Relation("member")

	testCases := []struct {
		name        string
		entity      string
		expectedRes Entity
		errRes      string
	}{
		{
			name: "happy path, user",

			entity: "user:bz0yOLsL460V-6L9HauX4",
			expectedRes: Entity{
				Kind:       "user",
				Identifier: "bz0yOLsL460V-6L9HauX4",
				Relation:   "",
			},
			errRes: "",
		},
		{
			name:   "relationship member",
			entity: "organization:yKreKfzq3-iG-rhj0N9o9#member",
			expectedRes: Entity{
				Kind:       "organization",
				Identifier: "yKreKfzq3-iG-rhj0N9o9",
				Relation:   memberRelation,
			},
			errRes: "",
		},
		{
			name:        "missing parts",
			entity:      "organization",
			expectedRes: Entity{},
			errRes:      "invalid entity representation",
		},
		{
			name:        "too many parts",
			entity:      "organization:yKreKfzq3-iG-rhj0N9o9#member:user:bz0yOLsL460V-6L9HauX4",
			expectedRes: Entity{},
			errRes:      "invalid entity representation",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			res, err := ParseEntity(tc.entity)

			// if we expect an error, check that first
			if tc.errRes != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.errRes)
				assert.Empty(t, res)

				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, res)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func Test_CreateRelationshipTuple(t *testing.T) {
	// setup mock controller
	mockCtrl := gomock.NewController(t)
	c := mock_client.NewMockSdkClient(mockCtrl)

	fc, err := NewTestFGAClient(t, mockCtrl, c)
	if err != nil {
		t.Fatal()
	}

	testCases := []struct {
		name        string
		relation    string
		object      string
		expectedRes []ofgaclient.ClientWriteSingleResponse
		errRes      error
	}{
		{
			name:     "happy path with relation",
			object:   "organization:datum",
			relation: "member",
			expectedRes: []ofgaclient.ClientWriteSingleResponse{
				{
					TupleKey: ofgaclient.ClientTupleKey{
						User:     "user:nano-id-of-member",
						Relation: "member",
						Object:   "organization:datum",
					},
					Status: ofgaclient.SUCCESS,
				},
			},
			errRes: nil,
		},
		{
			name:     "error, missing relation",
			object:   "organization:datum",
			relation: "",
			expectedRes: []ofgaclient.ClientWriteSingleResponse{
				{
					Status: ofgaclient.FAILURE,
				},
			},
			errRes: ErrMissingRelation,
		},
		{
			name:     "error, missing object",
			object:   "",
			relation: "can_view",
			expectedRes: []ofgaclient.ClientWriteSingleResponse{
				{
					Status: ofgaclient.FAILURE,
				},
			},
			errRes: ErrMissingObject,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// mock response for input
			ctk := []ofgaclient.ClientTupleKey{
				{
					User:     "user:nano-id-of-member",
					Relation: tc.relation,
					Object:   tc.object,
				},
			}

			mockWriteTuples(mockCtrl, c, context.Background(), ctk, tc.errRes)

			cr, err := fc.CreateRelationshipTuple(context.Background(), ctk)

			if tc.errRes != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.errRes)

				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, cr)
			assert.Equal(t, tc.expectedRes, cr.Writes)
		})
	}
}

func Test_DeleteRelationshipTuple(t *testing.T) {
	// setup mock controller
	mockCtrl := gomock.NewController(t)
	c := mock_client.NewMockSdkClient(mockCtrl)

	fc, err := NewTestFGAClient(t, mockCtrl, c)
	if err != nil {
		t.Fatal()
	}

	testCases := []struct {
		name        string
		relation    string
		object      string
		expectedRes string
		errRes      string
	}{
		{
			name:        "happy path with relation",
			object:      "organization:datum",
			relation:    "member",
			expectedRes: "",
			errRes:      "",
		},
		{
			name:        "error, missing relation",
			object:      "organization:datum",
			relation:    "",
			expectedRes: "",
			errRes:      "Reason: the 'relation' field is malformed",
		},
		{
			name:        "error, missing object",
			object:      "",
			relation:    "member",
			expectedRes: "",
			errRes:      "Reason: invalid 'object' field format",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tuples := []ofgaclient.ClientTupleKey{
				{
					User:     "user:nano-id-of-member",
					Relation: tc.relation,
					Object:   tc.object,
				},
			}

			mockDeleteTuples(mockCtrl, c, context.Background(), tuples, tc.errRes)

			_, err = fc.deleteRelationshipTuple(context.Background(), tuples)

			if tc.errRes != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.errRes)

				return
			}

			assert.NoError(t, err)
		})
	}
}

// mockWriteTuples creates mock responses based on the mock FGA client
func mockWriteTuples(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, tuples []ofgaclient.ClientTupleKey, errMsg error) {
	mockExecute := mock_client.NewMockSdkClientWriteTuplesRequestInterface(mockCtrl)

	if errMsg == nil {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteSingleResponse{
				{
					TupleKey: tuples[0],
					Status:   ofgaclient.SUCCESS,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, nil)
	} else {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteSingleResponse{
				{
					TupleKey: tuples[0],
					Status:   ofgaclient.FAILURE,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, errMsg)
	}

	mockRequest := mock_client.NewMockSdkClientWriteTuplesRequestInterface(mockCtrl)

	options := ofgaclient.ClientWriteOptions{AuthorizationModelId: openfga.PtrString("test-model-id")}

	mockRequest.EXPECT().Options(options).Return(mockExecute)

	mockBody := mock_client.NewMockSdkClientWriteTuplesRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(tuples).Return(mockRequest)

	c.EXPECT().WriteTuples(ctx).Return(mockBody)
}

// mockDeleteTuples creates mock responses based on the mock FGA client
func mockDeleteTuples(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, tuples []ofgaclient.ClientTupleKey, errMsg string) {
	mockExecute := mock_client.NewMockSdkClientDeleteTuplesRequestInterface(mockCtrl)

	if errMsg == "" {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteSingleResponse{
				{
					TupleKey: tuples[0],
					Status:   ofgaclient.SUCCESS,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, nil)
	} else {
		var err error

		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteSingleResponse{
				{
					TupleKey: tuples[0],
					Status:   ofgaclient.FAILURE,
				},
			},
		}

		if errMsg != "" {
			err = errors.New(errMsg) // nolint:goerr113
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, err)
	}

	mockRequest := mock_client.NewMockSdkClientDeleteTuplesRequestInterface(mockCtrl)

	options := ofgaclient.ClientWriteOptions{AuthorizationModelId: openfga.PtrString("test-model-id")}

	mockRequest.EXPECT().Options(options).Return(mockExecute)

	mockBody := mock_client.NewMockSdkClientDeleteTuplesRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(tuples).Return(mockRequest)

	c.EXPECT().DeleteTuples(ctx).Return(mockBody)
}
