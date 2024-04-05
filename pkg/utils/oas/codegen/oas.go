package oas

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"gopkg.in/yaml.v3"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/rout"
)

// OAS is a helper struct to generate OpenAPI schema
type OAS struct {
	// Info provides metadata about the API
	Info openapi3.Info `json:"info,omitempty" yaml:"info,omitempty"`
	// Servers provides the base URL for the API
	Servers openapi3.Servers `json:"servers,omitempty" yaml:"servers,omitempty"`
	// Components provides reusable objects for the API
	Components openapi3.Components `json:"components,omitempty" yaml:"components,omitempty"`
	// Paths provides the available paths for the API
	Paths openapi3.Paths `json:"paths,omitempty" yaml:"paths,omitempty"`
	// Schemas provides reusable references to object schemas for the API
	Schemas openapi3.Schemas `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	// Responses provides the available responses for the API
	Responses openapi3.Responses `json:"responses,omitempty" yaml:"responses,omitempty"`
	// Operation provides the available operations for the API
	Operation openapi3.Operation `json:"operation,omitempty" yaml:"operation,omitempty"`
	// Parameters provides the available parameters for the API
	Parameters openapi3.Parameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	// Headers provides the available headers for the API
	Headers openapi3.Headers `json:"headers,omitempty" yaml:"headers,omitempty"`
	// SecuritySchemes provides the available security schemes for the API
	SecuritySchemes openapi3.SecuritySchemes `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
}

// NewSchemaGenerator is a constructor function that creates a new instance of the `OAS` struct
// and initializes the `Paths`, `Components`, `Schemas`, `Responses`, `Operation`, `Parameters`, `Headers`,
// and `SecuritySchemes` fields of the struct. It then returns the newly created `OAS` struct
func NewSchemaGenerator() *OAS {
	s := &OAS{}
	s.Paths = openapi3.Paths{}
	s.Components = openapi3.NewComponents()
	s.Components.Schemas = make(map[string]*openapi3.SchemaRef)
	s.Components.Responses = make(map[string]*openapi3.ResponseRef)
	s.Components.Examples = make(map[string]*openapi3.ExampleRef)
	s.Components.Parameters = make(map[string]*openapi3.ParameterRef)
	s.Components.RequestBodies = make(map[string]*openapi3.RequestBodyRef)
	s.Components.Headers = make(map[string]*openapi3.HeaderRef)
	s.Components.SecuritySchemes = make(map[string]*openapi3.SecuritySchemeRef)
	Operation := openapi3.NewOperation()
	s.Operation = *Operation

	return s
}

var appJSON = "application/json"

// AddSecurityScheme adds a security scheme to the OAS
func (s *OAS) AddSecurityScheme(name string, scheme *openapi3.SecurityScheme) {
	s.Components.SecuritySchemes[name] = &openapi3.SecuritySchemeRef{Value: scheme}
}

// AddRequestBody is used to add a request body to the OpenAPI schema - it takes a name for the request body
// and the body interface as parameters; it then creates a new request body object with a JSON schema reference
// and adds an example for the request body in the `application/json` content type
func (s *OAS) AddRequestBody(name string, body interface{}) {
	request := openapi3.NewRequestBody().WithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/schemas/" + name})
	request.Content.Get(appJSON).Examples = make(map[string]*openapi3.ExampleRef)
	request.Content.Get(appJSON).Examples["error"] = &openapi3.ExampleRef{Value: openapi3.NewExample(body)}
	s.Components.RequestBodies[name] = &openapi3.RequestBodyRef{Value: request}
}

// AddHeader is a method of the `OAS` struct that is used to add a header
// to the OpenAPI schema. It takes a name for the header and a pointer to an `openapi3.Header` object
// as parameters. Inside the function, it adds the provided header to the `Components.Headers` field of
// the `OAS` struct, making it available for use in the OpenAPI schema generation process. This
// function allows you to define and include custom headers in your API documentation
func (s *OAS) AddHeader(name string, header *openapi3.Header) {
	s.Components.Headers[name] = &openapi3.HeaderRef{Value: header}
}

// AddResponse is used to add a response definition to the OpenAPI schema
// It takes parameters such as the name of the response, a description of the response, a
// reference to a schema, and an example of the response body
func (s *OAS) AddResponse(name string, description string, ref string, example interface{}) {
	response := openapi3.NewResponse().WithDescription(description).WithJSONSchemaRef(&openapi3.SchemaRef{Ref: ref})
	response.Content.Get(appJSON).Examples = make(map[string]*openapi3.ExampleRef)
	response.Content.Get(appJSON).Examples["error"] = &openapi3.ExampleRef{Value: openapi3.NewExample(example)}
	s.Components.Responses[name] = &openapi3.ResponseRef{Value: response}
}

// AddErrorSchemas function is adding error schemas to the OpenAPI schema. It is defining and
// including error schemas such as "StatusError", "MissingRequiredFieldError", and corresponding
// responses like "BadRequest", "Unauthorized", "InternalServerError", and "Conflict" in the OpenAPI
// schema generation process. These error schemas and responses are used to document and handle
// different error scenarios that may occur in the API
func AddErrorSchemas(gen *OAS) {
	gen.AddSchema("StatusError", &rout.StatusError{})
	gen.AddSchema("MissingRequiredFieldError", &rout.MissingRequiredFieldError{})

	var path = "#/components/schemas/StatusError"

	gen.AddResponse("BadRequest", "The request's parameters are not valid", path, rout.BadRequest())
	gen.AddResponse("Unauthorized", "The request is not authorized", path, rout.Unauthorized())
	gen.AddResponse("InternalServerError", "The server encountered an error", path, rout.InternalServerError())
	gen.AddResponse("Conflict", "The request conflicts with the current state of the server", path, rout.Conflict())
}

// AddHandlers is a function that adds handlers to the OpenAPI schema. It takes a pointer to an `OAS` struct
// as a parameter and adds the following handlers to the `Components.Schemas` field of the `OAS` struct:
// `LoginRequest`, `LoginReply`, `ForgotPasswordRequest`, `ForgotPasswordReply`, `ResetPasswordRequest`,
// `ResetPasswordReply`, `RefreshRequest`, `RefreshReply`, `RegisterRequest`, `RegisterReply`, `ResendEmailRequest`,
// `ResendEmailReply`, `VerifyRequest`, `VerifyReply`, `SubscribeRequest`, and `SubscribeReply`. These handlers
// are used to define the request and response schemas for the API endpoints in the OpenAPI schema generation process
func AddHandlers(gen *OAS) {
	gen.AddSchema("LoginRequest", &handlers.LoginRequest{})
	gen.AddRequestBody("LoginRequest", &handlers.LoginRequest{})
	gen.AddSchema("LoginReply", &handlers.LoginReply{})
	gen.AddResponse("LoginReply", "The login was successful", "#/components/schemas/LoginReply", handlers.LoginReply{})
	gen.AddSchema("ForgotPasswordRequest", &handlers.ForgotPasswordRequest{})
	gen.AddRequestBody("ForgotPasswordRequest", &handlers.ForgotPasswordRequest{})
	gen.AddSchema("ForgotPasswordReply", &handlers.ForgotPasswordReply{})
	gen.AddResponse("ForgotPasswordReply", "The password reset email was sent", "#/components/schemas/ForgotPasswordReply", handlers.ForgotPasswordReply{})
	gen.AddSchema("ResetPasswordRequest", &handlers.ResetPasswordRequest{})
	gen.AddRequestBody("ResetPasswordRequest", &handlers.ResetPasswordRequest{})
	gen.AddSchema("ResetPasswordReply", &handlers.ResetPasswordReply{})
	gen.AddResponse("ResetPasswordReply", "The password was reset", "#/components/schemas/ResetPasswordReply", handlers.ResetPasswordReply{})
	gen.AddSchema("RefreshRequest", &handlers.RefreshRequest{})
	gen.AddRequestBody("RefreshRequest", &handlers.RefreshRequest{})
	gen.AddSchema("RefreshReply", &handlers.RefreshReply{})
	gen.AddResponse("RefreshReply", "The token was refreshed", "#/components/schemas/RefreshReply", handlers.RefreshReply{})
	gen.AddSchema("RegisterRequest", &handlers.RegisterRequest{})
	gen.AddRequestBody("RegisterRequest", &handlers.RegisterRequest{})
	gen.AddSchema("RegisterReply", &handlers.RegisterReply{})
	gen.AddResponse("RegisterReply", "The user was registered", "#/components/schemas/RegisterReply", handlers.RegisterReply{})
	gen.AddSchema("ResendEmailRequest", &handlers.ResendRequest{})
	gen.AddRequestBody("ResendEmailRequest", &handlers.ResendRequest{})
	gen.AddSchema("ResendEmailReply", &handlers.ResendReply{})
	gen.AddResponse("ResendEmailReply", "The email was resent", "#/components/schemas/ResendEmailReply", handlers.ResendReply{})
	gen.AddSchema("VerifyRequest", &handlers.VerifyRequest{})
	gen.AddRequestBody("VerifyRequest", &handlers.VerifyRequest{})
	gen.AddSchema("VerifyReply", &handlers.VerifyReply{})
	gen.AddResponse("VerifyReply", "The email was verified", "#/components/schemas/VerifyReply", handlers.VerifyReply{})
}

// AddSchema is a method of the `OAS` struct that is used to add a schema to the OpenAPI schema
func (s *OAS) AddSchema(name string, model interface{}) {
	schema, err := openapi3gen.NewSchemaRefForValue(model, s.Components.Schemas, enableNullableAndDescriptionOpts)
	if err != nil {
		panic(err)
	}

	s.Components.Schemas[name] = schema
}

// AddExample is a method of the `OAS` struct that is used to add an example to the OpenAPI schema
func (s *OAS) AddExample(name string, value interface{}) {
	rval := reflect.TypeOf(value)
	checkTags(rval)

	example := openapi3.NewExample(value)
	s.Components.Examples[name] = &openapi3.ExampleRef{Value: example}
}

// AddParameters is a function that adds parameters to the OpenAPI schema. It takes a pointer to an `OAS` struct
// as a parameter and adds the `TokenQuery` parameter to the `Components.Parameters` field of the `OAS` struct
func AddParameters(gen *OAS) {
	gen.AddQueryParameter("Token", TokenQuery)
}

// AddExample is a function that adds an example to the OpenAPI schema. It takes a name and a value as
// parameters and creates a new example object with the provided value. The example is then added to the
// `Components.Examples` field of the `OAS` struct, making it available for use in the OpenAPI schema generation process
func AddExamples(gen *OAS) {
	gen.AddExample("BadRequest", rout.BadRequest())
}

// AddQueryParameter is a method of the `OAS` struct that is used to add a query parameter
func (s *OAS) AddQueryParameter(name string, value Parameter) {
	checkTags(reflect.TypeOf(value))

	param := &openapi3.Parameter{
		Name:        value.Name,
		In:          value.In,
		Description: value.Description,
		Required:    value.Required,
		Schema: &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Default: value.Default,
				Type:    value.Type,
			},
		},
	}
	s.Components.Parameters[name] = &openapi3.ParameterRef{Value: param}
}

// Parameter is a struct that represents a parameter in the OpenAPI schema
type Parameter struct {
	Name        string      `json:"name" yaml:"name"`
	Description string      `json:"description" yaml:"description"`
	Required    bool        `json:"required" yaml:"required"`
	Default     interface{} `json:"default" yaml:"default"`
	Type        string      `json:"type" yaml:"type"`
	In          string      `json:"in" yaml:"in"`
}

// TokenQuery is a parameter for the token query
var TokenQuery = Parameter{
	Name:        "token",
	Description: "the token to parse out of a URL",
	Type:        "integer",
	Required:    true,
	In:          "query",
}

// checkTags is a helper function that checks if a struct has the required JSON or YAML tags
func checkTags(rval reflect.Type) {
	if rval.Kind() == reflect.Array || rval.Kind() == reflect.Slice {
		checkTags(rval.Elem())
		return
	}

	if rval.Kind() != reflect.Struct {
		fmt.Printf("unable to check type %s of kind %s for struct tags, skipped\n", rval.Name(), rval.Kind().String())
		return
	}

	for i := 0; i < rval.NumField(); i++ {
		for _, tagName := range []string{"json", "yaml"} {
			if _, ok := rval.Field(i).Tag.Lookup(tagName); !ok {
				panic(fmt.Errorf("type %s does not have struct flag '%s'", rval.Name(), tagName)) // nolint: goerr113
			}
		}
	}
}

// oas is the entry point for the schemagen tool. It generates an OpenAPI schema for the Datum API
func oas() { // nolint: unused
	gen := NewSchemaGenerator()
	AddErrorSchemas(gen)
	AddExamples(gen)
	AddParameters(gen)
	AddHandlers(gen)

	// store schema part as buffer
	schemasYaml, err := yaml.Marshal(&gen)
	if err != nil {
		panic(err)
	}

	bufferYAML, err := os.ReadFile("./schemagen/base.yaml")
	if err != nil {
		panic(err)
	}

	// append both into single schema
	bufferYAML = append(bufferYAML, schemasYaml...)

	// load full schema
	loadedSchema, err := openapi3.NewLoader().LoadFromData(bufferYAML)
	if err != nil {
		panic(err)
	}

	// update version in the full schema and store it again
	if len(os.Args) >= 2 { // nolint: gomnd
		loadedSchema.Info.Version = os.Args[1]
		bufferYAML, err = yaml.Marshal(&loadedSchema)

		if err != nil {
			panic(err)
		}
	}

	// validate it
	err = loadedSchema.Validate(context.Background())
	if err != nil {
		panic(err)
	}

	// and store the full schema as JSON and YAML
	bufferJSON, err := json.MarshalIndent(loadedSchema, "", "  ")

	if err != nil {
		panic(err)
	}

	tmp := make([]byte, len(bufferJSON), len(bufferJSON)+1)
	copy(tmp, bufferJSON)
	tmp = append(tmp, '\n')
	bufferJSON = tmp

	err = os.WriteFile("./oas/outputschema.gen.json", bufferJSON, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./oas/outputschema.gen.yaml", bufferYAML, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}
}

// enableNullableAndDescriptionOpts is allowing tagging with description and nullable to work
var enableNullableAndDescriptionOpts = openapi3gen.SchemaCustomizer(
	func(name string, t reflect.Type, tag reflect.StructTag, schema *openapi3.Schema) error {
		if tag.Get("nullable") == "true" {
			schema.Nullable = true
		}
		if tag.Get("deprecated") == "true" {
			schema.Deprecated = true
		}
		if desc, ok := tag.Lookup("description"); ok && desc != "-" {
			schema.Description = desc
		}
		return nil
	},
)
