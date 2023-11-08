package entx

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// Extension is an implementation of entc.Extension
type Extension struct {
	entc.DefaultExtension

	templates []*gen.Template

	gqlSchemaHooks []entgql.SchemaHook
}

// ExtensionOption allow for control over the behavior of the generator
type ExtensionOption func(*Extension) error

// WithJSONScalar adds the JSON scalar definition
func WithJSONScalar() ExtensionOption {
	return func(ex *Extension) error {
		ex.gqlSchemaHooks = append(ex.gqlSchemaHooks, addJSONScalar)
		return nil
	}
}

// NewExtension returns an entc Extension that allows the entx package to generate
// the schema changes and templates needed to function
func NewExtension(opts ...ExtensionOption) (*Extension, error) {
	e := &Extension{
		templates:      []*gen.Template{},
		gqlSchemaHooks: []entgql.SchemaHook{},
	}

	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}

	return e, nil
}

// Templates of the extension
func (e *Extension) Templates() []*gen.Template {
	return e.templates
}

// GQLSchemaHooks of the extension to seamlessly edit the final gql interface
func (e *Extension) GQLSchemaHooks() []entgql.SchemaHook {
	return e.gqlSchemaHooks
}

var _ entc.Extension = (*Extension)(nil)
