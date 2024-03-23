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
	// ErrGenerateOpenAPI throws when fails the marshalling of the openapi struct
	ErrGenerateOpenAPI = errors.New("fail to generate openapi")
	// ErrValidatingOpenAPI throws when given openapi params are not correct
	ErrValidatingOpenAPI = errors.New("fails to validate openapi")
	// ErrOpenAPIRequired throws when openapi is required
	ErrOpenAPIRequired = errors.New("openapi is required")
	// ErrOpenAPITitleRequired throws when openapi title is required
	ErrOpenAPITitleRequired = errors.New("openapi title is required")
	// ErrOpenAPIVersionRequired throws when openapi version is required
	ErrOpenAPIVersionRequired = errors.New("openapi version is required")
	// ErrOepnAPIInfoRequired throws when openapi info is required
	ErrOepnAPIInfoRequired = errors.New("openapi info is required")
)
