package main

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/datumforge/datum/config"
	"github.com/invopop/jsonschema"
	"github.com/invopop/yaml"
)

// const values used for the schema generator
const (
	jsonSchemaPath = "./jsonschema/datum.config.json"
	yamlConfigPath = "./config/config.example.yaml"
)

// includedPackages is a list of packages to include in the schema generation
// that contain Go comments to be added to the schema
// any external packages must use the jsonschema description tags to add comments
var includedPackages = []string{
	"./config",
	"./internal/cache",
	"./internal/entdb",
	"./internal/httpserve/handlers",
	"./internal/otelx",
	"./internal/sessions",
	"./internal/tokens",
	"./internal/utils/emails",
	"./internal/utils/sentry",
	"./internal/providers",
	"./pkg/analytics",
}

// schemaConfig represents the configuration for the schema generator
type schemaConfig struct {
	// jsonSchemaPath represents the file path of the JSON schema to be generated
	jsonSchemaPath string
	// yamlConfigPath is the file path to the YAML configuration to be generated
	yamlConfigPath string
}

func main() {
	c := schemaConfig{
		jsonSchemaPath: jsonSchemaPath,
		yamlConfigPath: yamlConfigPath,
	}

	if err := generateSchema(c, &config.Config{}); err != nil {
		panic(err)
	}
}

// generateSchema generates a JSON schema and a YAML schema based on the provided schemaConfig and structure
func generateSchema(c schemaConfig, structure interface{}) error {
	// override the default name to using the prefixed pkg name
	r := jsonschema.Reflector{Namer: namePkg}
	r.ExpandedStruct = true
	// set `jsonschema:required` tag to true to generate required fields
	r.RequiredFromJSONSchemaTags = true
	// set the tag name to `koanf` for the koanf struct tags
	r.FieldNameTag = "koanf"

	// add go comments to the schema
	for _, pkg := range includedPackages {
		if err := r.AddGoComments("github.com/datumforge/datum/", pkg); err != nil {
			panic(err.Error())
		}
	}

	s := r.Reflect(structure)

	// generate the json schema
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile(c.jsonSchemaPath, data, 0600) // nolint: gomnd
	if err != nil {
		panic(err.Error())
	}

	// generate yaml schema
	var yamlConfig config.Config

	// this uses the `json` tag to generate the yaml schema
	yamlSchema, err := yaml.Marshal(yamlConfig)

	err = os.WriteFile(c.yamlConfigPath, yamlSchema, 0600) // nolint: gomnd
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func namePkg(r reflect.Type) string {
	return r.String()
}
