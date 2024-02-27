package main

import (
	"encoding/json"

	"github.com/ogen-go/ogen"
)

type ContentType string

func setupBase() *ogen.Spec {
	ex := &ogen.Spec{
		OpenAPI: "3.1.0",
		Info: ogen.Info{
			Title:          "title",
			Description:    "description",
			TermsOfService: "terms of service",
			Contact: &ogen.Contact{
				Name:  "Name",
				URL:   "url",
				Email: "email",
			},
			License: &ogen.License{
				Name: "Name",
				URL:  "url",
			},
			Version: "0.1.0",
		},
		Servers: []ogen.Server{
			{
				URL:         "https://api.example.com",
				Description: "Main server",
				Variables: map[string]ogen.ServerVariable{
					"version": {
						Enum:        []string{"v1", "v2"},
						Default:     "v1",
						Description: "API version",
					},

					"env": {
						Enum:        []string{"dev", "prod"},
						Default:     "dev",
						Description: "Environment",
					},

					"port": {
						Enum:        []string{"80", "443"},
						Default:     "80",
						Description: "Port",
					},
				},
			},
		},
		Paths: map[string]*ogen.PathItem{
			pathWithID: {
				Description: "This is my first path",
				Parameters: []*ogen.Parameter{
					{Ref: "#/components/parameters/authInQuery"},
					{Ref: "#/components/parameters/authInHeader"},
					{Ref: "#/components/parameters/csrf"},
				},
				Get: &ogen.Operation{
					Tags:        []string{"default"},
					Description: "Description for my path",
					OperationID: "path-with-id",
					Parameters: []*ogen.Parameter{
						{
							Name:        "id",
							In:          "path",
							Description: "ID Parameter in path",
							Required:    true,
							Schema:      &ogen.Schema{Type: "integer", Format: "int32"},
						},
					},
					Responses: ogen.Responses{
						"error": {Ref: "#/components/responses/error"},
						"ok": {
							Description: "Success",
							Content: map[string]ogen.Media{
								string(ContentTypeJSON): {Schema: &ogen.Schema{
									Type:        "object",
									Description: "Success",
									Properties: []ogen.Property{
										{Name: "prop1", Schema: &ogen.Schema{Type: "integer", Format: "int32"}},
										{Name: "prop2", Schema: &ogen.Schema{Type: "string"}},
									},
								}},
							},
						},
					},
				},
			},
			refPathWithID: {
				Ref: "#/paths/~1path~1with~1{id}",
			},
			pathWithBody: {
				Post: &ogen.Operation{
					Tags:        []string{"post"},
					Description: "Description for my path with body",
					OperationID: "path-with-body",
					Parameters: []*ogen.Parameter{
						{Ref: "#/components/parameters/authInQuery"},
						{Ref: "#/components/parameters/authInHeader"},
						{Ref: "#/components/parameters/csrf"},
					},
					Responses:   ogen.Responses{"error": {Ref: "#/components/responses/error"}},
					RequestBody: &ogen.RequestBody{Ref: "#/components/requestBodies/~1path~1with~1body"},
				},
			},
		},
		Components: &ogen.Components{
			Schemas: map[string]*ogen.Schema{
				_petSchema.Name: _petSchema.Schema,
				_toySchema.Name: _toySchema.Schema,
			},
			Responses: ogen.Responses{
				"error": {
					Description: "An Error Response",
					Content: map[string]ogen.Media{
						string(ContentTypeJSON): {Schema: &ogen.Schema{
							Type:        "object",
							Description: "Error Response Schema",
							Properties: []ogen.Property{
								{Name: "code", Schema: &ogen.Schema{Type: "integer", Format: "int32"}},
								{Name: "status", Schema: &ogen.Schema{Type: "string"}},
							},
						}},
					},
				},
			},
			Parameters: map[string]*ogen.Parameter{
				"authInQuery": {
					Name:        "auth",
					In:          "query",
					Description: "Optional bearer token",
				},
				"authInHeader": {
					Name:        "Authorization",
					In:          "header",
					Description: "Optional bearer token",
				},
				"csrf": {
					Name:        "csrf",
					In:          "cookie",
					Description: "CSRF token",
				},
			},
			RequestBodies: map[string]*ogen.RequestBody{
				pathWithBody: {
					Description: "Referenced RequestBody",
					Content: map[string]ogen.Media{
						string(ContentTypeJSON): {
							Schema: &ogen.Schema{Ref: "#/components/schemas/" + _toySchema.Name},
						},
					},
					Required: true,
				},
			},
		},
	}
	return ex
}

const (
	pathWithID    = "/path/with/{id}"
	refPathWithID = "/ref/path/with/id"
	pathWithBody  = "/path/with/body"

	ContentTypeJSON ContentType = "application/json"
)

var (
	// reusable query param
	_queryParam = ogen.NewParameter().
			InQuery().
			SetName("auth").
			SetDescription("Optional bearer token").
			ToNamed("authInQuery")
	// reusable header param
	_headerParam = ogen.NewNamedParameter(
		"authInHeader",
		ogen.NewParameter().
			SetIn("header").
			SetName("Authorization").
			SetDescription("Optional bearer token"),
	)
	// reusable cookie param
	_cookieParam = ogen.NewParameter().
			InCookie().
			SetName("csrf").
			SetDescription("CSRF token").
			ToNamed("csrf")
	// reusable pet schema
	_petSchema = ogen.NewNamedSchema(
		"Pet",
		ogen.NewSchema().
			SetDescription("A Pet").
			AddRequiredProperties(
				ogen.Int().ToProperty("required_Int"),
				ogen.Int32().ToProperty("required_Int32"),
				ogen.Int64().ToProperty("required_Int64"),
				ogen.Float().ToProperty("required_Float"),
				ogen.Double().ToProperty("required_Double"),
				ogen.String().ToProperty("required_String"),
				ogen.Bytes().ToProperty("required_Bytes"),
				ogen.Binary().ToProperty("required_Binary"),
				ogen.Bool().ToProperty("required_Bool"),
				ogen.Date().ToProperty("required_Date"),
				ogen.DateTime().ToProperty("required_DateTime"),
				ogen.Password().ToProperty("required_Password"),
				ogen.Int32().AsArray().ToProperty("required_array_Int32"),
				ogen.Int32().AsEnum(json.RawMessage("0"), json.RawMessage("0"), json.RawMessage("1")).
					ToProperty("required_enum_Int32"),
				ogen.Int32().AsEnum(json.RawMessage(`"off"`), json.RawMessage(`"0"`), json.RawMessage(`"1"`)).
					ToProperty("required_enum_String"),
			).
			AddOptionalProperties(
				ogen.UUID().ToProperty("optional_UUID"),
				ogen.Int32().ToProperty("optional_Int32"),
				ogen.Int64().ToProperty("optional_Int64"),
				ogen.Float().ToProperty("optional_Float"),
				ogen.Double().ToProperty("optional_Double"),
				ogen.String().ToProperty("optional_String"),
				ogen.Bytes().ToProperty("optional_Bytes"),
				ogen.Binary().ToProperty("optional_Binary"),
				ogen.Bool().ToProperty("optional_Bool"),
				ogen.Date().ToProperty("optional_Date"),
				ogen.DateTime().ToProperty("optional_DateTime"),
				ogen.Password().ToProperty("optional_Password"),
			),
	)
	// reusable toy schema
	_toySchema = ogen.NewSchema().
			SetDescription("A toy of a Pet").
			ToNamed("User")
)
