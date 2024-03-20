package oas

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/invopop/jsonschema"
)

// AddRawRoute adds route to router with specific method, path and handler as well as adding it to the openAPI schema
func (r Router[HandlerFunc, Route]) AddRawRoute(method string, routePath string, handler HandlerFunc, operation Operation) (Route, error) {
	op := operation.Operation
	if op != nil {
		err := operation.Validate(r.context)
		if err != nil {
			return getZero[Route](), err
		}
	} else {
		op = openapi3.NewOperation()
		if op.Responses == nil {
			op.Responses = openapi3.NewResponses()
		}
	}

	pathWithPrefix := path.Join(r.pathPrefix, routePath)

	oasPath := r.router.TransformPathToOASPath(pathWithPrefix)
	r.openAPISchema.AddOperation(oasPath, method, op)

	return r.router.AddRoute(method, pathWithPrefix, handler), nil
}

// Content is the type of a content in a map
type Content map[string]Schema

// Schema contains the value and if properties allow additional properties
type Schema struct {
	Value                     interface{}
	AllowAdditionalProperties bool
}

// Parameter is the struct containing the schema or the content information
type Parameter struct {
	Content     Content
	Schema      *Schema
	Description string
}

// ParameterValue is the map containing the parameter information
type ParameterValue map[string]Parameter

// ContentValue is the struct containing the content and the description
type ContentValue struct {
	Content     Content
	Description string
}

// SecurityRequirements is the array of security requirements
type SecurityRequirements []SecurityRequirement

// SecurityRequirement is the map containing the security requirements
type SecurityRequirement map[string][]string

// Definitions of the route https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#definitions
type Definitions struct {
	// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#specification-extensions
	Extensions map[string]interface{}
	// Optional field for documentation
	Tags        []string
	Summary     string
	Description string
	Deprecated  bool

	// PathParams contains the path parameters
	PathParams  ParameterValue
	Querystring ParameterValue
	Headers     ParameterValue
	Cookies     ParameterValue
	RequestBody *ContentValue
	Responses   map[int]ContentValue

	Security SecurityRequirements
}

// newOperationFromDefinition is a function that creates and returns a new `Operation` object
// based on the provided `Definitions` schema. It initializes the `Operation` object with values
// extracted from the `Definitions` - it sets up the responses, request body, and
// parameters for the operation based on the schema provided
func newOperationFromDefinition(schema Definitions) Operation {
	operation := NewOperation()
	operation.Responses = &openapi3.Responses{}
	operation.Tags = schema.Tags
	operation.Extensions = schema.Extensions
	operation.addSecurityRequirements(schema.Security)
	operation.Description = schema.Description
	operation.Summary = schema.Summary
	operation.Deprecated = schema.Deprecated

	return operation
}

const (
	pathParamsType  = "path"
	queryParamType  = "query"
	headerParamType = "header"
	cookieParamType = "cookie"
	errorFormat     = "%w: %s"
)

// AddRoute is a method of the `Router` struct that adds a route with a specified method, path, handler function, and schema
// definitions. It creates a new `Operation` object based on the provided schema definitions, sets up the request body,
// responses, and parameters for the operation, and then adds the route to the router with the specified details
func (r Router[HandlerFunc, Route]) AddRoute(method string, path string, handler HandlerFunc, schema Definitions) (Route, error) {
	operation := newOperationFromDefinition(schema)

	err := r.resolveRequestBodySchema(schema.RequestBody, operation)
	if err != nil {
		return getZero[Route](), fmt.Errorf(errorFormat, ErrRequestBody, err)
	}

	err = r.resolveResponsesSchema(schema.Responses, operation)
	if err != nil {
		return getZero[Route](), fmt.Errorf(errorFormat, ErrResponses, err)
	}

	oasPath := r.router.TransformPathToOASPath(path)

	err = r.resolveParameterSchema(pathParamsType, getPathParamsAutoComplete(schema, oasPath), operation)
	if err != nil {
		return getZero[Route](), fmt.Errorf(errorFormat, ErrPathParams, err)
	}

	err = r.resolveParameterSchema(queryParamType, schema.Querystring, operation)
	if err != nil {
		return getZero[Route](), fmt.Errorf(errorFormat, ErrPathParams, err)
	}

	err = r.resolveParameterSchema(headerParamType, schema.Headers, operation)
	if err != nil {
		return getZero[Route](), fmt.Errorf(errorFormat, ErrPathParams, err)
	}

	err = r.resolveParameterSchema(cookieParamType, schema.Cookies, operation)
	if err != nil {
		return getZero[Route](), fmt.Errorf(errorFormat, ErrPathParams, err)
	}

	return r.AddRawRoute(method, path, handler, operation)
}

// `getSchemaFromInterface` is a method defined in the `Router` struct that takes an interface `v` and a boolean `allowAdditionalProperties` as input
// parameters. This method is responsible for converting the input interface `v` into an OpenAPI 3 schema. It uses the `jsonschema.Reflector` to reflect the
// input interface and generate a JSON schema representation. The generated JSON schema is then converted into an OpenAPI 3 schema by creating a new
// `openapi3.Schema` object and unmarshaling the JSON data into it
func (r Router[_, _]) getSchemaFromInterface(v interface{}, allowAdditionalProperties bool) (*openapi3.Schema, error) {
	if v == nil {
		return &openapi3.Schema{}, nil
	}

	reflector := &jsonschema.Reflector{
		DoNotReference:            true,
		AllowAdditionalProperties: allowAdditionalProperties,
		Anonymous:                 true,
	}

	jsonSchema := reflector.Reflect(v)
	jsonSchema.Version = ""
	// Definitions are not valid in openapi3, which use components
	jsonSchema.Definitions = nil

	data, err := jsonSchema.MarshalJSON()
	if err != nil {
		return nil, err
	}

	schema := openapi3.NewSchema()
	err = schema.UnmarshalJSON(data)

	if err != nil {
		return nil, err
	}

	return schema, nil
}

// resolveRequestBodySchema function is responsible for adding the request body to the operation
// in the OpenAPI schema. It takes in the body schema information, which includes the content and
// description of the request body
func (r Router[_, _]) resolveRequestBodySchema(bodySchema *ContentValue, operation Operation) error {
	if bodySchema == nil {
		return nil
	}

	content, err := r.addContentToOASSchema(bodySchema.Content)

	if err != nil {
		return err
	}

	requestBody := openapi3.NewRequestBody().WithContent(content)

	if bodySchema.Description != "" {
		requestBody.WithDescription(bodySchema.Description)
	}

	operation.AddRequestBody(requestBody)

	return nil
}

// resolveResponsesSchema is a function that adds the responses to the operation in the OpenAPI schema - it takes in a map of response status codes
// and corresponding content values. For each status code and content value pair in the map, it creates a new OpenAPI 3 response object, adds the content to
// the response, sets the description of the response, and then adds this response to the operation
func (r Router[_, _]) resolveResponsesSchema(responses map[int]ContentValue, operation Operation) error {
	if responses == nil {
		operation.Responses = openapi3.NewResponses()
	}

	for statusCode, v := range responses {
		response := openapi3.NewResponse()

		content, err := r.addContentToOASSchema(v.Content)
		if err != nil {
			return err
		}

		response = response.WithContent(content)
		response = response.WithDescription(v.Description)

		operation.AddResponse(statusCode, response)
	}

	return nil
}

// resolveParameterSchema function is responsible for adding parameters to the operation in the
// OpenAPI schema - it takes in the parameter type (such as path, query, header, or cookie), the
// parameter configuration (ParameterValue), and the operation to which the parameters need to be added
func (r Router[_, _]) resolveParameterSchema(paramType string, paramConfig ParameterValue, operation Operation) error {
	var keys = make([]string, 0, len(paramConfig))
	for k := range paramConfig {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		v := paramConfig[key]

		var param *openapi3.Parameter

		switch paramType {
		case pathParamsType:
			param = openapi3.NewPathParameter(key)
		case queryParamType:
			param = openapi3.NewQueryParameter(key)
		case headerParamType:
			param = openapi3.NewHeaderParameter(key)
		case cookieParamType:
			param = openapi3.NewCookieParameter(key)
		default:
			return ErrInvalidParamType
		}

		if v.Description != "" {
			param = param.WithDescription(v.Description)
		}

		if v.Content != nil {
			content, err := r.addContentToOASSchema(v.Content)
			if err != nil {
				return err
			}

			param.Content = content
		} else {
			schema := openapi3.NewSchema()

			if v.Schema != nil {
				var err error
				schema, err = r.getSchemaFromInterface(v.Schema.Value, v.Schema.AllowAdditionalProperties)

				if err != nil {
					return err
				}
			}

			param.WithSchema(schema)
		}

		operation.AddParameter(param)
	}

	return nil
}

// addContentToOASSchema function is responsible for converting content information into an OpenAPI 3 schema format
func (r Router[_, _]) addContentToOASSchema(content Content) (openapi3.Content, error) {
	oasContent := openapi3.NewContent()

	for k, v := range content {
		var err error

		schema, err := r.getSchemaFromInterface(v.Value, v.AllowAdditionalProperties)
		if err != nil {
			return nil, err
		}

		oasContent[k] = openapi3.NewMediaType().WithSchema(schema)
	}

	return oasContent, nil
}

// getPathParamsAutoComplete function is used to extract path parameters from the provided path string - it checks the path segments for curly braces `{}`
// which indicate a path parameter; if it finds any path segment enclosed in curly braces, it extracts the parameter name and adds it to the
// `PathParams` field in the `Definitions` schema - if the `PathParams` field is not already initialized, it initializes it as a map of parameter names to `Parameter` objects with an empty
// schema value
func getPathParamsAutoComplete(schema Definitions, path string) ParameterValue {
	if schema.PathParams == nil {
		pathParams := strings.Split(path, "/")
		for _, param := range pathParams {
			if strings.HasPrefix(param, "{") && strings.HasSuffix(param, "}") {
				if schema.PathParams == nil {
					schema.PathParams = make(ParameterValue)
				}

				param = strings.Replace(param, "{", "", 1)
				param = strings.Replace(param, "}", "", 1)
				schema.PathParams[param] = Parameter{
					Schema: &Schema{Value: ""},
				}
			}
		}
	}

	return schema.PathParams
}

// getZero returns the zero value of a type
func getZero[T any]() T {
	var result T
	return result
}
