package model

import (
	"github.com/invopop/jsonschema"
)

// OpenAPI is used to store an entire openapi spec
type OpenAPI struct {
	OpenAPI           string                `json:"openapi,omitempty"`
	Info              Info                  `json:"info,omitempty"`
	JSONSchemaDialect string                `json:"jsonSchemaDialect,omitempty"`
	Servers           []Server              `json:"servers,omitempty"`
	Paths             map[string]Path       `json:"paths,omitempty"`
	Webhooks          map[string]Path       `json:"webhooks,omitempty"`
	Components        *Components           `json:"components,omitempty"`
	Security          []map[string][]string `json:"security,omitempty"`
	Tags              []Tag                 `json:"tags,omitempty"`
	ExternalDocs      *ExternalDocs         `json:"externalDocs,omitempty"`
}

// Components describes all openapi components
type Components struct {
	Schemas         map[string]jsonschema.Schema    `json:"schemas,omitempty"`
	Responses       Responses                       `json:"responses,omitempty"`
	Parameters      map[string]Parameter            `json:"parameters,omitempty"`
	Examples        map[string]Example              `json:"examples,omitempty"`
	RequestBodies   map[string]RequestBody          `json:"requestBodies,omitempty"`
	Headers         map[string]map[string]Parameter `json:"headers,omitempty"`
	SecuritySchemes map[string]map[string][]string  `json:"securitySchemes,omitempty"`
	Links           map[string]Link                 `json:"links,omitempty"`
	Callbacks       map[string]map[string]Path      `json:"callbacks,omitempty"`
	PathItems       map[string]Path                 `json:"pathItems,omitempty"`
}

// Tag is used to tag parts of an API
type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ExternalDocs is a link to external API documentation
type ExternalDocs struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Info holds info about an API
type Info struct {
	Title          string   `json:"title"`
	Summary        string   `json:"summary,omitempty"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version"`
}

// License describes the license of an API
type License struct {
	Name       string `json:"name,omitempty"`
	URL        string `json:"url,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

// Contact stores basic contanct info
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// Server describes a server
type Server struct {
	URL         string                    `json:"url,omitempty"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

// ServerVariable is a variable used in servers
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
}

// Path stores all operations allowed on a particular path
type Path struct {
	// Ref         string      `json:"$ref,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Description string      `json:"description,omitempty"`
	Get         *Operation  `json:"get,omitempty"`
	Put         *Operation  `json:"put,omitempty"`
	Post        *Operation  `json:"post,omitempty"`
	Delete      *Operation  `json:"delete,omitempty"`
	Options     *Operation  `json:"options,omitempty"`
	Head        *Operation  `json:"head,omitempty"`
	Patch       *Operation  `json:"patch,omitempty"`
	Trace       *Operation  `json:"trace,omitempty"`
	Servers     []Server    `json:"servers,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty"`
}

// RequestSpec is the description of an openapi request used in an Operation
type RequestSpec struct {
	Parameters  []Parameter  `json:"parameters,omitempty"`
	RequestBody *RequestBody `json:"requestBody,omitempty"`
}

// Operation describes an openapi Operation
type Operation struct {
	*RequestSpec
	Tags         []string                   `json:"tags,omitempty"`
	Summary      string                     `json:"summary,omitempty"`
	Description  string                     `json:"description,omitempty"`
	ExternalDocs *ExternalDocs              `json:"externalDocs,omitempty"`
	OperationID  string                     `json:"operationId,omitempty"`
	Callbacks    map[string]map[string]Path `json:"callbacks,omitempty"`
	Deprecated   bool                       `json:"deprecated,omitempty"`
	Security     []map[string][]string      `json:"security,omitempty"`
	Servers      []Server                   `json:"servers,omitempty"`
	Responses    Responses                  `json:"responses,omitempty"`
}

// ResponseSpec is an openapi Response description
type ResponseSpec struct {
	Description string               `json:"description"`
	Headers     map[string]Parameter `json:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Links       map[string]Link      `json:"links,omitempty"`
}

// Link descript a link to parts of a spec
type Link struct {
	OperationRef string         `json:"operationRef,omitempty"`
	OperationId  string         `json:"operationId,omitempty"` //nolint: stylecheck
	Parameters   map[string]any `json:"parameters,omitempty"`
	RequestBody  any            `json:"requestBody,omitempty"`
	Description  string         `json:"description,omitempty"`
	Server       *Server        `json:"server,omitempty"`
}

// RequestBody is the spec of a request body
type RequestBody struct {
	Description string               `json:"description"`
	Required    bool                 `json:"required,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
}

// MediaType describes a media type in openapi
type MediaType struct {
	Schema   *jsonschema.Schema  `json:"schema,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
	Example  any                 `json:"example,omitempty"`
	Examples *map[string]Example `json:"examples,omitempty"`
}

// Encoding is used to describe a content encoding in an API
type Encoding struct {
	ContentType   string               `json:"contentType,omitempty"`
	Headers       map[string]Parameter `json:"headers,omitempty"`
	Style         string               `json:"style,omitempty"`
	Explode       bool                 `json:"explode,omitempty"`
	AllowReserved bool                 `json:"allowReserved,omitempty"`
}

// Parameter describes a paramater used in requests/responses
type Parameter struct {
	Name            string              `json:"name,omitempty"`
	In              string              `json:"in,omitempty"`
	Description     string              `json:"description,omitempty"`
	Required        bool                `json:"required,omitempty"`
	Deprecated      bool                `json:"deprecated,omitempty"`
	AllowEmptyValue bool                `json:"allowEmptyValue,omitempty"`
	Style           string              `json:"style,omitempty"`
	Explode         bool                `json:"explode,omitempty"`
	AllowReserved   bool                `json:"allowReserved,omitempty"`
	Schema          *jsonschema.Schema  `json:"schema,omitempty"`
	Example         any                 `json:"example,omitempty"`
	Examples        *map[string]Example `json:"examples,omitempty"`
}

// Example is an example of any type
type Example struct {
	Summary       string `json:"summary,omitempty"`
	Description   string `json:"description,omitempty"`
	Value         any    `json:"value,omitempty"`
	ExternalValue string `json:"externalValue,omitempty"`
	Example       any    `json:"example,omitempty"`
}

// Responses is a map of status code string to ResponseSpec
type Responses map[string]ResponseSpec

// type Reference struct {
// 	Ref         string `json:"$ref"`
// 	Summary     string `json:"summary,omitempty"`
// 	Description string `json:"description,omitempty"`
// }
