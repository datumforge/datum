package model

import (
	"github.com/invopop/jsonschema"
)

// OpenAPI is used to store an entire openapi spec
type OpenAPI struct {
	OpenAPI           string                `json:"openapi,omitempty"`
	Info              *Info                 `json:"info,omitempty"`
	JSONSchemaDialect string                `json:"jsonSchemaDialect,omitempty"`
	Servers           []*Server             `json:"servers,omitempty"`
	Paths             map[string]Path       `json:"paths,omitempty"`
	Webhooks          map[string]Path       `json:"webhooks,omitempty"`
	Components        *Components           `json:"components,omitempty"`
	Security          []map[string][]string `json:"security,omitempty"`
	Tags              []*Tag                `json:"tags,omitempty"`
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
	Title          string         `json:"title"`
	Summary        string         `json:"summary,omitempty"`
	Description    string         `json:"description,omitempty"`
	TermsOfService string         `json:"termsOfService,omitempty"`
	Contact        *Contact       `json:"contact,omitempty"`
	License        *License       `json:"license,omitempty"`
	Version        string         `json:"version"`
	Extensions     map[string]any `json:"-"` // The field name is used as the key in the map
}

// License describes the license of an API
type License struct {
	Name       string         `json:"name,omitempty"`
	URL        string         `json:"url,omitempty"`
	Identifier string         `json:"identifier,omitempty"`
	Extensions map[string]any `json:"-"` // The field name is used as the key in the map
}

// Contact stores basic contanct info
type Contact struct {
	// Name of the contact person / organization
	Name string `json:"name,omitempty"`
	// The URL pointing to the contact information
	URL string `json:"url,omitempty"`
	// The email address of the contact person/organization
	Email string `json:"email,omitempty"`
	// Extensions (user-defined properties) if any
	Extensions map[string]any `json:"-"` // The field name is used as the key in the map
}

// Server describes a server
type Server struct {
	URL         string                     `json:"url,omitempty"`
	Description string                     `json:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
	Extensions  map[string]any             `json:"-"` // The field name is used as the key in the map
}

// ServerVariable is a variable used in servers
type ServerVariable struct {
	Enum        []string       `json:"enum,omitempty"`
	Default     string         `json:"default,omitempty"`
	Description string         `json:"description,omitempty"`
	Extensions  map[string]any `json:"-"` // The field name is used as the key in the map
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
	Servers      []*Server                  `json:"servers,omitempty"`
	Responses    *Responses                 `json:"responses,omitempty"`
}

// Link descript a link to parts of a spec
type Link struct {
	Ref          string         `json:"$ref,omitempty"`
	OperationRef string         `json:"operationRef,omitempty"`
	OperationId  string         `json:"operationId,omitempty"` //nolint: stylecheck
	Parameters   map[string]any `json:"parameters,omitempty"`
	RequestBody  any            `json:"requestBody,omitempty"`
	Description  string         `json:"description,omitempty"`
	Server       *Server        `json:"server,omitempty"`
	Extensions   map[string]any `json:"-"` // The field name is used as the key in the map
}

// RequestBody is the spec of a request body
type RequestBody struct {
	Ref         string                `json:"$ref,omitempty"`
	Description string                `json:"description"`
	Required    bool                  `json:"required,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty"`
	Extensions  map[string]any        `json:"-"` // The field name is used as the key in the map
}

// MediaType describes a media type in openapi
type MediaType struct {
	Schema     *jsonschema.Schema   `json:"schema,omitempty"`
	Encoding   map[string]*Encoding `json:"encoding,omitempty"`
	Example    any                  `json:"example,omitempty"`
	Examples   map[string]*Example  `json:"examples,omitempty"`
	Extensions map[string]any       `json:"-"` // The field name is used as the key in the map
}

// Encoding is used to describe a content encoding in an API
type Encoding struct {
	ContentType   string               `json:"contentType,omitempty"`
	Headers       map[string]Parameter `json:"headers,omitempty"`
	Style         string               `json:"style,omitempty"`
	Explode       bool                 `json:"explode,omitempty"`
	AllowReserved bool                 `json:"allowReserved,omitempty"`
	Extensions    map[string]any       `json:"-"` // The field name is used as the key in the map
}

// Parameter describes a paramater used in requests/responses
type Parameter struct {
	Ref             string              `json:"$ref,omitempty"`
	Name            string              `json:"name,omitempty"`
	In              string              `json:"in,omitempty"`
	Description     string              `json:"description,omitempty"`
	Required        bool                `json:"required,omitempty"`
	Deprecated      bool                `json:"deprecated,omitempty"`
	AllowEmptyValue bool                `json:"allowEmptyValue,omitempty"`
	Style           string              `json:"style,omitempty"`
	Explode         *bool               `json:"explode,omitempty"`
	AllowReserved   bool                `json:"allowReserved,omitempty"`
	Schema          *jsonschema.Schema  `json:"schema,omitempty"`
	Example         any                 `json:"example,omitempty"`
	Examples        map[string]*Example `json:"examples,omitempty"`
	Extensions      map[string]any      `json:"-"` // The field name is used as the key in the map
}

// Example is an example of any type
type Example struct {
	Ref           string         `json:"$ref,omitempty"`
	Summary       string         `json:"summary,omitempty"`
	Description   string         `json:"description,omitempty"`
	Value         any            `json:"value,omitempty"`
	ExternalValue string         `json:"externalValue,omitempty"`
	Extensions    map[string]any `json:"-"` // The field name is used as the key in the map
}

// Responses is a map of status code string to ResponseSpec
type Responses map[string]ResponseSpec

// ResponseSpec is an openapi Response description
type ResponseSpec struct {
	Ref         string                `json:"$ref,omitempty"`
	Description string                `json:"description"`
	Headers     map[string]*Parameter `json:"headers,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty"`
	Links       map[string]*Link      `json:"links,omitempty"`
	Extensions  map[string]any        `json:"-"` // The field name is used as the key in the map
}

// OAuthFlow stores configuration details for a supported OAuth Flow.
//
//	type: oauth2
//	flows:
//	  implicit:
//	    authorizationUrl: https://example.com/api/oauth/dialog
//	    scopes:
//	      write:pets: modify pets in your account
//	      read:pets: read your pets
//	  authorizationCode:
//	    authorizationUrl: https://example.com/api/oauth/dialog
//	    tokenUrl: https://example.com/api/oauth/token
//	    scopes:
//	      write:pets: modify pets in your account
//	      read:pets: read your pets
type OAuthFlow struct {
	// AuthorizationURL is REQUIRED for `implicit` and `authorizationCode` flows.
	// The authorization URL to be used for this flow. This MUST be in the form of
	// a URL. The OAuth2 standard requires the use of TLS.
	AuthorizationURL string `yaml:"authorizationUrl,omitempty"`

	// TokenURL is REQUIRED. The token URL to be used for this flow. This MUST be
	// in the form of a URL. The OAuth2 standard requires the use of TLS.
	TokenURL string `yaml:"tokenUrl"`

	// RefreshURL is the URL to be used for obtaining refresh tokens. This MUST be
	// in the form of a URL. The OAuth2 standard requires the use of TLS.
	RefreshURL string `yaml:"refreshUrl,omitempty"`

	// Scopes are REQUIRED. The available scopes for the OAuth2 security scheme. A
	// map between the scope name and a short description for it. The map MAY be
	// empty.
	Scopes map[string]string `yaml:"scopes"`

	// Extensions (user-defined properties), if any. Values in this map will
	// be marshalled as siblings of the other properties above.
	Extensions map[string]any `yaml:",inline"`
}

// OAuthFlows allows configuration of the supported OAuth Flows.
type OAuthFlows struct {
	// Implicit is the configuration for the OAuth Implicit flow.
	Implicit *OAuthFlow `yaml:"implicit,omitempty"`

	// Password is the configuration for the OAuth Resource Owner Password flow.
	Password *OAuthFlow `yaml:"password,omitempty"`

	// ClientCredentials is the configuration for the OAuth Client Credentials
	// flow. Previously called application in OpenAPI 2.0.
	ClientCredentials *OAuthFlow `yaml:"clientCredentials,omitempty"`

	// AuthorizationCode is the configuration for the OAuth Authorization Code
	// flow. Previously called accessCode in OpenAPI 2.0.
	AuthorizationCode *OAuthFlow `yaml:"authorizationCode,omitempty"`

	// Extensions (user-defined properties), if any. Values in this map will
	// be marshalled as siblings of the other properties above.
	Extensions map[string]any `yaml:",inline"`
}
