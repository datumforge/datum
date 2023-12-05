//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"

	"github.com/datumforge/datum/internal/graphapi"
)

// read in schema from internal package and save it to the schema file
func main() {
	execSchema := graphapi.NewExecutableSchema(graphapi.Config{})
	schema := execSchema.Schema()

	// Some of our federation fields get marked as "BuiltIn" by gengql and the formatter doesn't print builtin types, this adds them for us.
	if entities := schema.Types["_Entity"]; entities != nil {
		entities.BuiltIn = false
	}
	if service := schema.Types["_Service"]; service != nil {
		service.BuiltIn = false
	}

	// Add UUID Type to schema.graphql
	schema.Types["UUID"] = &ast.Definition{
		Kind:        ast.Scalar,
		Description: "A Universally Unique Identifier (UUID)",
		Name:        "UUID",
	}

	// Add JSON Type to schema.graphql
	//	schema.Types["JSON"] = &ast.Definition{
	//		Kind:        ast.Scalar,
	//		Description: "JSON",
	//		Name:        "JSON",
	//	}

	f, err := os.Create("schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmtr := formatter.NewFormatter(f)

	fmtr.FormatSchema(schema)

	f.Write(federationSchema)
}

var federationSchema = []byte(`
extend schema
  @link(
    url: "https://specs.apollo.dev/federation/v2.3"
    import: [
      "@key",
      "@interfaceObject",
      "@shareable",
      "@inaccessible",
      "@override",
      "@provides",
      "@requires",
      "@tag"
    ]
  )
`)
