package oas

import (
	"errors"
)

var (
	// ErrResponses is thrown if an error occurs generating the responses schemas
	ErrResponses = errors.New("errors generating responses schema")
	// ErrRequestBody is thrown if an error occurs generating the request body schemas
	ErrRequestBody = errors.New("errors generating request body schema")
	// ErrPathParams is thrown if an error occurs generating path params schemas
	ErrPathParams = errors.New("errors generating path parameters schema")
	// ErrQuerystring is thrown if error occurs generating querystring params schemas
	ErrQuerystring = errors.New("errors generating querystring schema")
	// ErrInvalidParamType is thrown if the parameter type is invalid
	ErrInvalidParamType = errors.New("invalid parameter type")
	// ErrGenerateSwagger throws when fails the marshalling of the swagger struct
	ErrGenerateSwagger = errors.New("fail to generate swagger")
	// ErrValidatingSwagger throws when given swagger params are not correct
	ErrValidatingSwagger = errors.New("fails to validate swagger")
	// ErrOpenAPIRequired throws when openapi is required
	ErrOpenAPIRequired = errors.New("openapi is required")
	// ErrOpenAPITitleRequired throws when openapi title is required
	ErrOpenAPITitleRequired = errors.New("openapi title is required")
	// ErrOpenAPIVersionRequired throws when openapi version is required
	ErrOpenAPIVersionRequired = errors.New("openapi version is required")
	// ErrOepnAPIInfoRequired throws when openapi info is required
	ErrOepnAPIInfoRequired = errors.New("openapi info is required")
)
