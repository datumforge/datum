//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/formatter"

	"github.com/datumforge/datum/internal/api"
)

// read in schema from internal package and save it to the schema file
func main() {
	execSchema := api.NewExecutableSchema(api.Config{})
	schema := execSchema.Schema()

	// Some of our federation fields get marked as "BuiltIn" by gengql and the formatter doesn't print builtin types, this adds them for us.
	if entities := schema.Types["_Entity"]; entities != nil {
		entities.BuiltIn = false
	}
	if service := schema.Types["_Service"]; service != nil {
		service.BuiltIn = false
	}
	// entities.Position.Src.BuiltIn = false

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
