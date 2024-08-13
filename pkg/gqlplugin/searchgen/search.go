package searchgen

import (
	"embed"
	_ "embed"
	"strings"
	"text/template"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/gertd/go-pluralize"
	"github.com/stoewer/go-strcase"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	relativeSchemaPath = "./internal/ent/schema"
)

var SearchDirective = entgql.NewDirective("search")

//go:embed templates/*
var _templates embed.FS

//go:embed templates/helpers.gotpl
var helperTemplate string

// SearchPlugin is a gqlgen plugin to generate search functions
type SearchPlugin struct{}

// Name returns the name of the plugin
// This name must match the upstream resolvergen to replace during code generation
func (r SearchPlugin) Name() string {
	return "searchgen"
}

// NewSearchPlugin returns a new search plugin
func New() *SearchPlugin {
	return &SearchPlugin{}
}

// SearchResolverBuild is a struct to hold the objects for the bulk resolver
type SearchResolverBuild struct {
	// Objects is a list of objects to generate bulk resolvers for
	Objects []Object
}

// Object is a struct to hold the object name for the bulk resolver
type Object struct {
	// Name of the object
	Name string
	// PluralName of the object
	Fields []string
}

// GenerateCode implements api.CodeGenerator
func (r SearchPlugin) GenerateCode(data *codegen.Data) error {
	inputData := SearchResolverBuild{
		Objects: []Object{},
	}

	graph, err := entc.LoadGraph(relativeSchemaPath, &gen.Config{})
	if err != nil {
		return err
	}

	for _, f := range data.Schema.Types {
		// Add the search fields
		if strings.Contains(f.Name, "Search") && !strings.Contains(f.Name, "GlobalSearch") {
			schemaName := strings.TrimSuffix(f.Name, "SearchResult")
			fields := getSearchableFields(schemaName, f, graph)

			if len(fields) > 0 {
				inputData.Objects = append(inputData.Objects, Object{
					Name:   schemaName,
					Fields: fields, // add the fields that are being searched
				})
			}
		}
	}

	// generate the search schema
	err = templates.Render(templates.Options{
		PackageName: data.Config.Resolver.Package,              // use the resolver package
		Filename:    data.Config.Resolver.Dir() + "/search.go", // write to the resolver directory
		FileNotice:  `// THIS CODE IS REGENERATED BY github.com/datumforge/datum/pkg/gqlplugin. DO NOT EDIT.`,
		Data:        inputData,
		Funcs: template.FuncMap{
			"toLower":  strings.ToLower,
			"toPlural": pluralize.NewClient().Plural,
		},
		Packages: data.Config.Packages,
		Template: helperTemplate,
	})

	// use the default resolver plugin to generate the code
	return err
}

func getEntSchema(graph *gen.Graph, name string) *load.Schema {
	for _, s := range graph.Schemas {
		if s.Name == name {
			return s
		}
	}

	return nil
}

func getSearchableFields(schemaName string, f *ast.Definition, graph *gen.Graph) (fields []string) {
	// add the object name that is being searched
	schema := getEntSchema(graph, schemaName)

	for _, field := range schema.Fields {
		if isFieldSearchable(field) {
			fields = append(fields, strcase.UpperCamelCase(field.Name))
		}
	}

	return
}

func isFieldSearchable(field *load.Field) bool {
	searchAnt := &SearchFieldAnnotation{}

	if ant, ok := field.Annotations[searchAnt.Name()]; ok {
		if err := searchAnt.Decode(ant); err != nil {
			return false
		}

		return searchAnt.Searchable
	}

	return false
}
