package fga

import (
	"context"
	"testing"

	openfga "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"

	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
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
	// setup mock client
	c := mock_fga.NewMockSdkClient(t)

	fc := NewMockFGAClient(t, c)

	testCases := []struct {
		name    string
		writes  []TupleKey
		deletes []TupleKey
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
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock_fga.WriteAny(t, c)

			_, err := fc.WriteTupleKeys(context.Background(), tc.writes, tc.deletes)
			assert.NoError(t, err)
		})
	}
}

func Test_DeleteRelationshipTuple(t *testing.T) {
	// setup mock client
	c := mock_fga.NewMockSdkClient(t)

	fc := NewMockFGAClient(t, c)

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
			defer mock_fga.ClearMocks(c)

			tuples := []openfga.TupleKeyWithoutCondition{
				{
					User:     "user:ulid-of-member",
					Relation: tc.relation,
					Object:   tc.object,
				},
			}

			mock_fga.DeleteAny(t, c, tc.errRes)

			_, err := fc.DeleteRelationshipTuple(context.Background(), tuples)

			if tc.errRes != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.errRes)

				return
			}

			assert.NoError(t, err)
		})
	}
}
