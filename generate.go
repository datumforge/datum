package main

//go:generate go run -mod=mod ./internal/ent/entc.go

//go:generate mockgen -destination=internal/fga/mocks/client.go -package=mock_client --build_flags=--mod=mod github.com/openfga/go-sdk/client SdkClient,SdkClientListStoresRequestInterface,SdkClientCreateStoreRequestInterface,SdkClientGetStoreRequestInterface,SdkClientDeleteStoreRequestInterface,SdkClientReadAuthorizationModelsRequestInterface,SdkClientWriteAuthorizationModelRequestInterface,SdkClientReadAuthorizationModelRequestInterface,SdkClientReadLatestAuthorizationModelRequestInterface,SdkClientReadChangesRequestInterface,SdkClientReadRequestInterface,SdkClientWriteRequestInterface,SdkClientWriteTuplesRequestInterface,SdkClientDeleteTuplesRequestInterface,SdkClientCheckRequestInterface,SdkClientBatchCheckRequestInterface,SdkClientExpandRequestInterface,SdkClientListObjectsRequestInterface,SdkClientListRelationsRequestInterface,SdkClientReadAssertionsRequestInterface,SdkClientWriteAssertionsRequestInterface
