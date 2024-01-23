package fga

import (
	"testing"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	"go.uber.org/zap"

	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
)

// NewMockFGAClient is a mock client based on the mockery testing framework
func NewMockFGAClient(t *testing.T, c *mock_fga.MockSdkClient) *Client {
	client := Client{
		Config: ofgaclient.ClientConfiguration{
			// The api host is the only required field when setting up a new FGA client connection
			ApiHost:              "fga.datum.net",
			AuthorizationModelId: *openfga.PtrString("test-model-id"),
			StoreId:              *openfga.PtrString("test-store-id"),
		},
		Ofga:   c,
		Logger: zap.NewNop().Sugar(),
	}

	return &client
}
