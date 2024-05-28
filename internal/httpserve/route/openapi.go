package route

import (
	"fmt"
	"reflect"

	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// Router is a struct that holds the echo router, the OpenAPI schema, and the handler - it's a way to group these components together
type Router struct {
	Echo    *echo.Echo
	OAS     *openapi3.T
	Handler *handlers.Handler
}

// AddRoute is used to add a route to the echo router and OpenAPI schema at the same time ensuring consistency between the spec and the server
func (r *Router) AddRoute(pattern, method string, op *openapi3.Operation, route echo.Routable) error {
	_, err := r.Echo.AddRoute(route)
	if err != nil {
		return err
	}

	r.OAS.AddOperation(pattern, method, op)

	return nil
}

func (r *Router) AddEchoOnlyRoute(pattern, method string, route echo.Routable) error {
	_, err := r.Echo.AddRoute(route)
	if err != nil {
		return err
	}

	return nil
}

// OpenAPI returns the OpenAPI specification.
func (r *Router) OpenAPI() *openapi3.T {
	return r.OAS
}

func (r *Router) AddQueryParameter(name string, value Parameter) {
	checkTags(reflect.TypeOf(value))

	param := &openapi3.Parameter{
		Name:        value.Name,
		In:          value.In,
		Description: value.Description,
		Required:    value.Required,
		Schema: &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Default: value.Default,
				Type:    (*openapi3.Types)(&value.Types),
			},
		},
	}
	r.OAS.Components.Parameters[name] = &openapi3.ParameterRef{Value: param}
}

// Parameter is a struct that represents a parameter in the OpenAPI schema
type Parameter struct {
	Name        string      `json:"name" yaml:"name"`
	Description string      `json:"description" yaml:"description"`
	Required    bool        `json:"required" yaml:"required"`
	Default     interface{} `json:"default" yaml:"default"`
	Types       []string    `json:"type" yaml:"type"`
	In          string      `json:"in" yaml:"in"`
}

// TokenQuery is a parameter for the token query
var TokenQuery = Parameter{
	Name:        "token",
	Description: "the token to parse out of a URL",
	Types:       []string{"integer"},
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
				return
			}
		}
	}
}

// AddExample adds an example to the OpenAPI schema
func (r *Router) AddExample(name string, value interface{}) {
	rval := reflect.TypeOf(value)
	checkTags(rval)

	example := openapi3.NewExample(value)
	r.OAS.Components.Examples[name] = &openapi3.ExampleRef{Value: example}
}

// AddSchema adds a schema to the OpenAPI schema
func (r *Router) AddSchema(name string, model interface{}) error {
	schema, err := openapi3gen.NewSchemaRefForValue(model, r.OAS.Components.Schemas)
	if err != nil {
		return err
	}

	r.OAS.Components.Schemas[name] = schema

	return nil
}

// SchemaGenerator is a helper function that generates a schema from a map
// example: var openapiSchemas = map[string]any{LoginRequest: LoginRequest{}}
func SchemaGenerator(openAPISchemas map[string]any) (openapi3.Schemas, error) {
	schemas := make(openapi3.Schemas)
	generator := openapi3gen.NewGenerator(openapi3gen.UseAllExportedFields())

	for key, val := range openAPISchemas {
		ref, err := generator.NewSchemaRefForValue(val, schemas)
		if err != nil {
			return nil, err
		}

		schemas[key] = ref
	}

	return schemas, nil
}

func (r *Router) AddSecurityScheme(name string, scheme *openapi3.SecurityScheme) {
	r.OAS.Components.SecuritySchemes[name] = &openapi3.SecuritySchemeRef{Value: scheme}
}
