package oasrouter

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

// AddRoute add a route with json schema inferred by passed schema
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

// getSchemaFromInterface returns the openapi3 schema from an interface
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

// resolveRequestBodySchema adds the request body to the operation
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

// resolveResponsesSchema adds the responses to the operation
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

// resolveParameterSchema adds the parameters to the operation
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

// addContentToOASSchema adds the content to the openapi3 schema
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

// getPathParamsAutoComplete returns the path parameters from the path
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
