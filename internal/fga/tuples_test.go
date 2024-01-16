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

func Test_tupleKeyToWriteRequest(t *testing.T) {
	testCases := []struct {
		name             string
		writes           []TupleKey
		expectedUser     string
		expectedRelation string
		expectedObject   string
		expectedCount    int
	}{
		{
			name: "happy path, user",
			writes: []TupleKey{
				{
					Subject: Entity{
						Kind:       "user",
						Identifier: "THEBESTUSER",
					},
					Relation: "member",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
					},
				},
			},
			expectedUser:     "user:THEBESTUSER",
			expectedRelation: "member",
			expectedObject:   "organization:IDOFTHEORG",
			expectedCount:    1,
		},
		{
			name: "happy path, group",
			writes: []TupleKey{
				{
					Subject: Entity{
						Kind:       "group",
						Identifier: "ADATUMGROUP",
					},
					Relation: "parent",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
						Relation:   "member",
					},
				},
			},
			expectedUser:     "group:ADATUMGROUP",
			expectedRelation: "parent",
			expectedObject:   "organization:IDOFTHEORG#member",
			expectedCount:    1,
		},
		{
			name: "happy path, multiple",
			writes: []TupleKey{
				{
					Subject: Entity{
						Kind:       "user",
						Identifier: "SITB",
					},
					Relation: "member",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
					},
				},
				{
					Subject: Entity{
						Kind:       "user",
						Identifier: "MITB",
					},
					Relation: "admin",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
					},
				},
			},
			expectedCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			ctk := tupleKeyToWriteRequest(tc.writes)
			assert.NotEmpty(t, ctk)
			if tc.expectedCount == 1 {
				assert.Equal(t, tc.expectedUser, ctk[0].User)
				assert.Equal(t, tc.expectedRelation, ctk[0].Relation)
				assert.Equal(t, tc.expectedObject, ctk[0].Object)
			} else {
				assert.Len(t, ctk, tc.expectedCount)
			}
		})
	}
}

func Test_tupleKeyToDeleteRequest(t *testing.T) {
	testCases := []struct {
		name             string
		writes           []TupleKey
		expectedUser     string
		expectedRelation string
		expectedObject   string
		expectedCount    int
	}{
		{
			name: "happy path, user",
			writes: []TupleKey{
				{
					Subject: Entity{
						Kind:       "user",
						Identifier: "THEBESTUSER",
					},
					Relation: "member",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
					},
				},
			},
			expectedUser:     "user:THEBESTUSER",
			expectedRelation: "member",
			expectedObject:   "organization:IDOFTHEORG",
			expectedCount:    1,
		},
		{
			name: "happy path, group",
			writes: []TupleKey{
				{
					Subject: Entity{
						Kind:       "group",
						Identifier: "ADATUMGROUP",
					},
					Relation: "parent",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
						Relation:   "member",
					},
				},
			},
			expectedUser:     "group:ADATUMGROUP",
			expectedRelation: "parent",
			expectedObject:   "organization:IDOFTHEORG#member",
			expectedCount:    1,
		},
		{
			name: "happy path, multiple",
			writes: []TupleKey{
				{
					Subject: Entity{
						Kind:       "user",
						Identifier: "SITB",
					},
					Relation: "member",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
					},
				},
				{
					Subject: Entity{
						Kind:       "user",
						Identifier: "MITB",
					},
					Relation: "admin",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
					},
				},
			},
			expectedCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			ctk := tupleKeyToDeleteRequest(tc.writes)
			assert.NotEmpty(t, ctk)
			if tc.expectedCount == 1 {
				assert.Equal(t, tc.expectedUser, ctk[0].User)
				assert.Equal(t, tc.expectedRelation, ctk[0].Relation)
				assert.Equal(t, tc.expectedObject, ctk[0].Object)
			} else {
				assert.Len(t, ctk, tc.expectedCount)
			}
		})
	}
}

func Test_WriteTupleKeys(t *testing.T) {
	// setup mock controller
	mockCtrl := gomock.NewController(t)
	c := mock_client.NewMockSdkClient(mockCtrl)

	fc, err := NewTestFGAClient(t, mockCtrl, c)
	if err != nil {
		t.Fatal()
	}

	testCases := []struct {
		name    string
		writes  []TupleKey
		deletes []TupleKey
		errRes  error
	}{
		{
			name: "happy path with relation",
			writes: []TupleKey{
				{
					Subject: Entity{
						Kind:       "user",
						Identifier: "THEBESTUSER",
					},
					Relation: "member",
					Object: Entity{
						Kind:       "organization",
						Identifier: "IDOFTHEORG",
					},
				},
			},
			errRes: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockWriteAny(mockCtrl, c, context.Background(), tc.errRes)

			_, err := fc.WriteTupleKeys(context.Background(), tc.writes, tc.deletes)
			assert.NoError(t, err)
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
			tuples := []openfga.TupleKeyWithoutCondition{
				{
					User:     "user:ulid-of-member",
					Relation: tc.relation,
					Object:   tc.object,
				},
			}

			mockDeleteTuples(mockCtrl, c, context.Background(), tuples, tc.errRes)

			_, err = fc.DeleteRelationshipTuple(context.Background(), tuples)

			if tc.errRes != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.errRes)

				return
			}

			assert.NoError(t, err)
		})
	}
}

// mockDeleteTuples creates mock responses based on the mock FGA client
func mockDeleteTuples(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, tuples []openfga.TupleKeyWithoutCondition, errMsg string) {
	mockExecute := mock_client.NewMockSdkClientDeleteTuplesRequestInterface(mockCtrl)

	if errMsg == "" {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Deletes: []ofgaclient.ClientWriteRequestDeleteResponse{
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
			Deletes: []ofgaclient.ClientWriteRequestDeleteResponse{
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

// mockWriteAny creates mock responses based on the mock FGA client
func mockWriteAny(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, errMsg error) {
	mockExecute := mock_client.NewMockSdkClientWriteRequestInterface(mockCtrl)

	if errMsg == nil {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteRequestWriteResponse{
				{
					Status: ofgaclient.SUCCESS,
				},
			},
			Deletes: []ofgaclient.ClientWriteRequestDeleteResponse{
				{
					Status: ofgaclient.SUCCESS,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, nil)
	} else {
		expectedResponse := ofgaclient.ClientWriteResponse{
			Writes: []ofgaclient.ClientWriteRequestWriteResponse{
				{
					Status: ofgaclient.FAILURE,
				},
			},
			Deletes: []ofgaclient.ClientWriteRequestDeleteResponse{
				{
					Status: ofgaclient.FAILURE,
				},
			},
		}

		mockExecute.EXPECT().Execute().Return(&expectedResponse, errMsg)
	}

	mockRequest := mock_client.NewMockSdkClientWriteRequestInterface(mockCtrl)

	mockRequest.EXPECT().Options(gomock.Any()).Return(mockExecute)

	mockBody := mock_client.NewMockSdkClientWriteRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(gomock.Any()).Return(mockRequest)

	c.EXPECT().Write(gomock.Any()).Return(mockBody)
}
