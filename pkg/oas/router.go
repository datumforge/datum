package oas

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"

	"github.com/datumforge/datum/pkg/oas/apirouter"
)

// Router handles calling the the apirouter and the OpenAPI schema to make that openAPI spec magic happen
type Router[HandlerFunc, Route any] struct {
	// router is the router to be used
	router apirouter.Router[HandlerFunc, Route]
	// openAPISchema is the openapi schema to be used by the router
	openAPISchema *openapi3.T
	// context is the context to be used by the router
	context context.Context
	// jsonDocumentationPath is the path exposed by json endpoint
	jsonDocumentationPath string
	// yamlDocumentationPath is the path exposed by yaml endpoint
	yamlDocumentationPath string
	// pathPrefix is the path prefix to be added to every route
	pathPrefix string
}

// Options to be passed to create the new router and openapi schema
type Options struct {
	// Context is the context to be used by the router
	Context context.Context
	// OpenAPI is the openapi schema to be used by the router
	OpenAPI *openapi3.T
	// JSONDocumentationPath is the path exposed by json endpoint. Default to /documentation/json
	JSONDocumentationPath string
	// YAMLDocumentationPath is the path exposed by yaml endpoint. Default to /documentation/yaml
	YAMLDocumentationPath string
	// Add path prefix to add to every router path
	PathPrefix string
}

var DefaultJSONDocumentationPath = "/documentation/json"
var DefaultYAMLDocumentationPath = "/documentation/yaml"

// NewRouter is a function that creates a new instance of the `Router` struct. It takes in a
// router, options, and returns a pointer to a `Router` instance along with an error. Inside the
// function, it validates the provided OpenAPI schema, sets default values for JSON and YAML
// documentation paths if not provided, and creates a new `Router` instance with the provided options
func NewRouter[HandlerFunc, Route any](router apirouter.Router[HandlerFunc, Route], options Options) (*Router[HandlerFunc, Route], error) {
	openAPI, err := generateNewValidOpenAPI(options.OpenAPI)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidatingOpenAPI, err)
	}

	var ctx = options.Context
	if options.Context == nil {
		ctx = context.Background()
	}

	yamlDocumentationPath := "/documentation/yaml"

	if options.YAMLDocumentationPath != "" {
		if err := isValidDocumentationPath(options.YAMLDocumentationPath); err != nil {
			return nil, err
		}

		yamlDocumentationPath = options.YAMLDocumentationPath
	}

	jsonDocumentationPath := "/documentation/json"

	if options.JSONDocumentationPath != "" {
		if err := isValidDocumentationPath(options.JSONDocumentationPath); err != nil {
			return nil, err
		}

		jsonDocumentationPath = options.JSONDocumentationPath
	}

	// Add path prefix to every route
	return &Router[HandlerFunc, Route]{
		router:                router,
		openAPISchema:         openAPI,
		context:               ctx,
		yamlDocumentationPath: yamlDocumentationPath,
		jsonDocumentationPath: jsonDocumentationPath,
		pathPrefix:            options.PathPrefix,
	}, nil
}

// SubRouterOptions to be passed to create the new sub router
type SubRouterOptions struct {
	PathPrefix string
}

// SubRouter is creating a new sub-router based on the existing main router - this is intended to allow you
// to create the same OpenAPI schema and context as the main router, but with a different path prefix
func (r Router[HandlerFunc, Route]) SubRouter(router apirouter.Router[HandlerFunc, Route], opts SubRouterOptions) (*Router[HandlerFunc, Route], error) {
	return &Router[HandlerFunc, Route]{
		router:                router,
		openAPISchema:         r.openAPISchema,
		context:               r.context,
		jsonDocumentationPath: r.jsonDocumentationPath,
		yamlDocumentationPath: r.yamlDocumentationPath,
		pathPrefix:            opts.PathPrefix,
	}, nil
}

// generateNewValidOpenAPI function is responsible for generating a new valid OpenAPI schema by
// checking and setting default values for certain core fields
func generateNewValidOpenAPI(openapi *openapi3.T) (*openapi3.T, error) {
	if openapi == nil {
		return nil, ErrOpenAPIRequired
	}

	if openapi.OpenAPI == "" {
		openapi.OpenAPI = "3.1.0"
	}

	if openapi.Paths == nil {
		openapi.Paths = &openapi3.Paths{}
	}

	if openapi.Info == nil {
		return nil, ErrOepnAPIInfoRequired
	}

	if openapi.Info.Title == "" {
		return nil, ErrOpenAPITitleRequired
	}

	if openapi.Info.Version == "" {
		return nil, ErrOpenAPIVersionRequired
	}

	return openapi, nil
}

// GenerateAndExposeOpenAPI creates a /documentation/json route on router and exposes the generated openAPI specifications
func (r Router[_, _]) GenerateAndExposeOpenAPI() error {
	if err := r.openAPISchema.Validate(r.context); err != nil {
		return fmt.Errorf("%w: %s", ErrValidatingOpenAPI, err)
	}

	// use the OASHandler to add the route and expose the specification embedded in the router / server
	jsonOAS, err := r.openAPISchema.MarshalJSON()
	if err != nil {
		return fmt.Errorf("%w json marshal: %s", ErrGenerateOpenAPI, err)
	}

	r.router.AddRoute(http.MethodGet, r.jsonDocumentationPath, r.router.OASHandler("application/json", jsonOAS))

	yamlOAS, err := yaml.JSONToYAML(jsonOAS)
	if err != nil {
		return fmt.Errorf("%w yaml marshal: %s", ErrGenerateOpenAPI, err)
	}

	r.router.AddRoute(http.MethodGet, r.yamlDocumentationPath, r.router.OASHandler("text/plain", yamlOAS))

	return nil
}

// isValidDocumentationPath checks if the path is valid
func isValidDocumentationPath(path string) error {
	if !strings.HasPrefix(path, "/") {
		return fmt.Errorf("invalid path %s. Path should start with '/'", path) //nolint: goerr113
	}

	return nil
}
