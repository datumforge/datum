package server

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// openAPISchemas is a mapping of types to auto generate schemas for
var openAPISchemas = map[string]any{
	"LoginRequest":  &handlers.LoginRequest{},
	"LoginResponse": &handlers.LoginReply{},
}

func NewOpenAPISpec() (*openapi3.T, error) {
	schemas := make(openapi3.Schemas)
	responses := make(openapi3.ResponseBodies)
	parameters := make(openapi3.ParametersMap)

	generator := openapi3gen.NewGenerator(openapi3gen.UseAllExportedFields())
	for key, val := range openAPISchemas {
		ref, err := generator.NewSchemaRefForValue(val, schemas)
		if err != nil {
			return nil, err
		}

		schemas[key] = ref
	}

	errorSchema := &openapi3.SchemaRef{
		Ref: "#/components/schemas/error",
	}

	errorResponse := openapi3.NewResponse().
		WithDescription("error").
		WithContent(openapi3.NewContentWithJSONSchemaRef(errorSchema))

	successResponse := openapi3.NewResponse().
		WithDescription("ok")

	// add common schemas, responses, and params so we can reference them
	schemas["document"] = &openapi3.SchemaRef{
		Value: openapi3.NewObjectSchema().WithAnyAdditionalProperties(),
	}
	responses["success"] = &openapi3.ResponseRef{
		Value: successResponse,
	}
	responses["error"] = &openapi3.ResponseRef{
		Value: errorResponse,
	}

	return &openapi3.T{
		OpenAPI: "3.1.0",
		Info: &openapi3.Info{
			Title:   "Datum OpenAPI 3.1.0 Specifications",
			Version: "v0.1.0",
		},
		Paths: openapi3.NewPaths(),
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Datum API Server",
				URL:         "https://api.datum.net/v1",
			},
		},
		ExternalDocs: &openapi3.ExternalDocs{
			Description: "Documentation for Datum's API services",
			URL:         "https://docs.datum.net",
		},
		Components: &openapi3.Components{
			Schemas:    schemas,
			Responses:  responses,
			Parameters: parameters,
		},
		Tags: openapi3.Tags{
			&openapi3.Tag{
				Name:        "schema",
				Description: "Add or update schema definitions",
			},
			&openapi3.Tag{
				Name:        "graphql",
				Description: "GraphQL query endpoints",
			},
		},
	}, nil
}
