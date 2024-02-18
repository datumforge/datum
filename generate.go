package main

//go:generate go run -mod=mod ./internal/ent/entc.go
//go:generate swagger generate spec -x github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options -o ./openapi
