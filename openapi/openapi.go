package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ogen-go/ogen"
	"github.com/samber/lo"
)

func getErrorResponsesResponses() []*ogen.NamedResponse {
	return []*ogen.NamedResponse{
		ogen.NewNamedResponse("400", ogen.NewResponse().SetRef("#/components/responses/400")),
		ogen.NewNamedResponse("403", ogen.NewResponse().SetRef("#/components/responses/403")),
		ogen.NewNamedResponse("404", ogen.NewResponse().SetRef("#/components/responses/404")),
		ogen.NewNamedResponse("409", ogen.NewResponse().SetRef("#/components/responses/409")),
		ogen.NewNamedResponse("429", ogen.NewResponse().SetRef("#/components/responses/429")),
		ogen.NewNamedResponse("500", ogen.NewResponse().SetRef("#/components/responses/500")),
	}
}

func getErrorResponses() map[string]*ogen.Response {
	return map[string]*ogen.Response{
		"400": ogen.NewResponse().SetDescription("Bad Request"),
		"403": ogen.NewResponse().SetDescription("Forbidden"),
		"404": ogen.NewResponse().SetDescription("Not Found"),
		"409": ogen.NewResponse().SetDescription("Conflict"),
		"429": ogen.NewResponse().SetDescription("Back the fuck off"),
		"500": ogen.NewResponse().SetDescription("Internal Server Error"),
	}
}

var (
	errorResponse = openapi3.ResponseRef{
		Value: &openapi3.Response{
			Description: lo.ToPtr("Error response"),
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: openapi3.TypeArray,
							Items: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: openapi3.TypeObject,
									Properties: openapi3.Schemas{
										"code": &openapi3.SchemaRef{
											Value: &openapi3.Schema{
												Type: openapi3.TypeInteger,
											},
										},
										"message": &openapi3.SchemaRef{
											Value: &openapi3.Schema{
												Type: openapi3.TypeString,
											},
										},
									},
									Required: []string{
										"code",
										"status",
										"message",
									},
								},
							},
						},
					},
				},
			},
		},
	}
)
