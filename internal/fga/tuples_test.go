package fga

import (
	"context"
	"os"
	"testing"

	"github.com/openfga/go-sdk/client"
	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/echox"
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

func Test_CreateCheckTupleWithUser(t *testing.T) {
	testCases := []struct {
		name        string
		relation    string
		object      string
		expectedRes *client.ClientCheckRequest
		errRes      error
	}{
		{
			name:     "happy path with relation",
			object:   "organization:datum",
			relation: "member",
			expectedRes: &client.ClientCheckRequest{
				User:     "user:foobar",
				Relation: "member",
				Object:   "organization:datum",
			},
			errRes: nil,
		},
		{
			name:        "error, missing relation",
			object:      "organization:datum",
			relation:    "",
			expectedRes: nil,
			errRes:      ErrMissingRelation,
		},
		{
			name:        "error, missing relation",
			object:      "",
			relation:    "can_view",
			expectedRes: nil,
			errRes:      ErrMissingObject,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			ec, err := echox.NewContextWithValidUser()
			if err != nil {
				t.Fatal()
			}

			echoContext := *ec

			ctx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

			echoContext.SetRequest(echoContext.Request().WithContext(ctx))

			url := os.Getenv("TEST_FGA_URL")
			if url == "" {
				url = defaultFGAURL
			}

			fc := newTestFGAClient(t, url)

			cr, err := fc.CreateCheckTupleWithUser(ctx, tc.relation, tc.object)

			if tc.errRes != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.errRes)

				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, cr)
			assert.Equal(t, tc.expectedRes, cr)
		})
	}
}
