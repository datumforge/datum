package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"gopkg.in/yaml.v2"

	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/openapi/payloads"
)

// ResponseError is used as a payload for all errors. Use NewResponseError function
// to create new type to set some fields correctly.
type ResponseError struct {
	// HTTP status code
	HTTPStatusCode int `json:"-" yaml:"-"`

	// user facing error message
	Message string `json:"msg,omitempty" yaml:"msg,omitempty"`

	// full root cause
	Error string `json:"error" yaml:"error"`

	// build commit
	Version string `json:"version" yaml:"version"`

	// environment (prod or stage or ephemeral)
	Environment string `json:"environment,omitempty" yaml:"environment"`
}

// APISchemaGen is a helper struct to generate OpenAPI schema
type APISchemaGen struct {
	Components openapi3.Components `json:"components,omitempty" yaml:"components,omitempty"`
	Servers    openapi3.Servers    `json:"servers,omitempty" yaml:"servers,omitempty"`
}

// NewSchemaGenerator creates new APISchemaGen
func NewSchemaGenerator() *APISchemaGen {
	s := &APISchemaGen{}
	//	s.Servers = openapi3.Servers{
	//		&openapi3.Server{
	//			Description: "Local development",
	//			URL:         "http://0.0.0.0:{port}/api/{applicationName}",
	//			Variables: map[string]*openapi3.ServerVariable{
	//				"applicationName": {Default: "datum"},
	//				"port":            {Default: "8000"},
	//			},
	//		},
	//	}
	s.Components = openapi3.NewComponents()
	s.Components.Schemas = make(map[string]*openapi3.SchemaRef)
	s.Components.Responses = make(map[string]*openapi3.ResponseRef)
	s.Components.Examples = make(map[string]*openapi3.ExampleRef)
	s.Components.Parameters = make(map[string]*openapi3.ParameterRef)

	return s
}

func (s *APISchemaGen) addResponse(name string, description string, ref string, example interface{}) {
	response := openapi3.NewResponse().WithDescription(description).WithJSONSchemaRef(&openapi3.SchemaRef{Ref: ref})

	response.Content.Get("application/json").Examples = make(map[string]*openapi3.ExampleRef)

	response.Content.Get("application/json").Examples["error"] = &openapi3.ExampleRef{Value: openapi3.NewExample(example)}

	s.Components.Responses[name] = &openapi3.ResponseRef{Value: response}
}

// addErrorSchemas all generic errors, that can be returned.
func addErrorSchemas(gen *APISchemaGen) {
	gen.addSchema("ErrorResponse", &rout.StatusError{})
	gen.addResponse("BadRequest", "The request's parameters are not valid", "#/components/schemas/ErrorResponse", rout.BadRequest())
}

// Schema customizer allowing tagging with description and nullable to work
var enableNullableAndDescriptionOpts = openapi3gen.SchemaCustomizer(
	func(_name string, _t reflect.Type, tag reflect.StructTag, schema *openapi3.Schema) error {
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

func (s *APISchemaGen) addSchema(name string, model interface{}) {
	schema, err := openapi3gen.NewSchemaRefForValue(model, s.Components.Schemas, enableNullableAndDescriptionOpts)
	if err != nil {
		panic(err)
	}

	s.Components.Schemas[name] = schema
}

var ResponseErrorGenericExample = payloads.ResponseError{
	Error:   "error: this can be pretty long string",
	Version: "df8a489",
}

type Parameter struct {
	Name        string      `json:"name" yaml:"name"`
	Description string      `json:"description" yaml:"description"`
	Required    bool        `json:"required" yaml:"required"`
	Default     interface{} `json:"default" yaml:"default"`
	Type        string      `json:"type" yaml:"type"`
	In          string      `json:"in" yaml:"in"`
}

var LimitQueryParam = Parameter{
	Name:        "limit",
	Description: "The number of items to return.",
	Default:     100,
	Type:        "integer",
	Required:    false,
	In:          "query",
}

func (s *APISchemaGen) addExample(name string, value interface{}) {
	// verify all fields has both json and yaml struct flags
	rval := reflect.TypeOf(value)
	checkTags(rval)

	example := openapi3.NewExample(value)
	s.Components.Examples[name] = &openapi3.ExampleRef{Value: example}
}

func addParameters(gen *APISchemaGen) {
	gen.addQueryParameter("Limit", LimitQueryParam)
}

func addExamples(gen *APISchemaGen) {
	gen.addExample("v1.GenericExample", ResponseErrorGenericExample)
}

func (s *APISchemaGen) addQueryParameter(name string, value Parameter) {
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

//nolint:goerr113
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
				panic(fmt.Errorf("type %s does not have struct flag '%s'", rval.Name(), tagName))
			}
		}
	}
}

func addPayloads(gen *APISchemaGen) {
	gen.addSchema("v1.PubkeyRequest", &payloads.PubkeyRequest{})
	gen.addSchema("v1.PubkeyResponse", &payloads.PubkeyResponse{})
}

func main() {
	gen := NewSchemaGenerator()
	addErrorSchemas(gen)
	addPayloads(gen)
	addExamples(gen)
	addParameters(gen)

	// store schema part as buffer
	schemasYaml, err := yaml.Marshal(&gen)
	if err != nil {
		panic(err)
	}

	bufferYAML, err := os.ReadFile("./openapi/base.yaml")
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

	err = os.WriteFile("./openapi/mitb.gen.json", bufferJSON, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./openapi/mitb.gen.yaml", bufferYAML, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}
}
