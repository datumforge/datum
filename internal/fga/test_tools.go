package fga

import (
	"context"
	"testing"
	"time"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	mock_client "github.com/datumforge/datum/internal/fga/mocks"
)

const (
	storeName = "datum_test"
)

func NewTestFGAClient(t testing.TB, mockCtrl *gomock.Controller, c *mock_client.MockSdkClient) (*Client, error) {
	// setup required mocks
	mockListStores(c, mockCtrl)
	mockCreateStore(c, mockCtrl)

	client := Client{
		Config: ofgaclient.ClientConfiguration{
			// The api host is the only required field when setting up a new FGA client connection
			ApiHost:              "fga.datum.net",
			AuthorizationModelId: openfga.PtrString("test-model-id"),
			StoreId:              *openfga.PtrString("test-store-id"),
		},
		Ofga:   c,
		Logger: zap.NewNop().Sugar(),
	}

	// Create new store
	_, err := client.CreateStore(context.Background(), storeName)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// mockListStores for testing functionality
func mockListStores(c *mock_client.MockSdkClient, mockCtrl *gomock.Controller) {
	mockExecute := mock_client.NewMockSdkClientListStoresRequestInterface(mockCtrl)

	// No stores exist, tests will create a store
	var stores *[]openfga.Store

	response := openfga.ListStoresResponse{
		ContinuationToken: openfga.PtrString(""),
		Stores:            stores,
	}

	mockExecute.EXPECT().Execute().Return(&response, nil)

	mockRequest := mock_client.NewMockSdkClientListStoresRequestInterface(mockCtrl)
	options := ofgaclient.ClientListStoresOptions{
		ContinuationToken: openfga.PtrString(""),
	}

	mockRequest.EXPECT().Options(options).Return(mockExecute)

	c.EXPECT().ListStores(context.Background()).Return(mockRequest)
}

// mockCreateStore for testing functionality
func mockCreateStore(c *mock_client.MockSdkClient, mockCtrl *gomock.Controller) {
	mockExecute := mock_client.NewMockSdkClientCreateStoreRequestInterface(mockCtrl)
	expectedTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	expectedResponse := ofgaclient.ClientCreateStoreResponse{
		Id:        openfga.PtrString("01HFJ0PR7XGSNP1H747YW0KZ6R"),
		Name:      openfga.PtrString(storeName),
		CreatedAt: openfga.PtrTime(expectedTime),
		UpdatedAt: openfga.PtrTime(expectedTime),
	}

	mockExecute.EXPECT().Execute().Return(&expectedResponse, nil)

	mockBody := mock_client.NewMockSdkClientCreateStoreRequestInterface(mockCtrl)

	body := ofgaclient.ClientCreateStoreRequest{
		Name: storeName,
	}
	mockBody.EXPECT().Body(body).Return(mockExecute)

	c.EXPECT().CreateStore(context.Background()).Return(mockBody)
}
