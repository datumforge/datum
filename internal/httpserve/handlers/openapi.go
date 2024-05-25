package handlers

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func badRequest() *openapi3.Response {
	return openapi3.NewResponse().
		WithDescription("Bad Request").
		WithContent(openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/responses/BadRequest"}))
}

func internalServerError() *openapi3.Response {
	return openapi3.NewResponse().
		WithDescription("Internal Server Error").
		WithContent(openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/responses/InternalServerError"}))
}

func notFound() *openapi3.Response {
	return openapi3.NewResponse().
		WithDescription("Not Found").
		WithContent(openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/responses/NotFound"}))
}

func created() *openapi3.Response {
	return openapi3.NewResponse().
		WithDescription("Created").
		WithContent(openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/responses/Created"}))
}

func conflict() *openapi3.Response {
	return openapi3.NewResponse().
		WithDescription("Conflict").
		WithContent(openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/responses/Conflict"}))
}

var appJSON = "application/json"

func (h *Handler) AddRequestBody(name string, body interface{}, op *openapi3.Operation) {
	request := openapi3.NewRequestBody().
		WithContent(openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/schemas/" + name}))
	op.RequestBody = &openapi3.RequestBodyRef{Value: request}
	request.Content.Get(appJSON).Examples = make(map[string]*openapi3.ExampleRef)
	request.Content.Get(appJSON).Examples["error"] = &openapi3.ExampleRef{Value: openapi3.NewExample(body)}
}

// AddResponse is used to add a response definition to the OpenAPI schema
func (h *Handler) AddResponse(name string, description string, body interface{}, op *openapi3.Operation, status int) {
	response := openapi3.NewResponse().
		WithDescription(description).
		WithContent(openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/schemas/" + name}))
	op.AddResponse(status, response)

	response.Content.Get(appJSON).Examples = make(map[string]*openapi3.ExampleRef)
	response.Content.Get(appJSON).Examples["error"] = &openapi3.ExampleRef{Value: openapi3.NewExample(body)}
}
