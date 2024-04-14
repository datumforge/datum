package main

import (
	"encoding/json"
	"os"

	"github.com/invopop/jsonschema"
	"github.com/invopop/yaml"
	"github.com/mcuadros/go-defaults"

	"github.com/datumforge/datum/jsonschema/models/marketplace"
)

// const values used for the schema generator
const (
	tagName        = "json"
	skipper        = "-"
	defaultTag     = "default"
	jsonSchemaPath = "./jsonschema/models/marketplace/generate/datum.marketplace.json"
	yamlConfigPath = "./jsonschema/models/marketplace/generate/datum.marketplace.yaml"
	ownerReadWrite = 0600
)

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

	if err := generateSchema(c, &marketplace.Config{}); err != nil {
		panic(err)
	}
}

// generateSchema generates a JSON schema and a YAML schema based on the provided schemaConfig and structure
func generateSchema(c schemaConfig, structure interface{}) error {
	// override the default name to using the prefixed pkg name
	r := new(jsonschema.Reflector)
	r.ExpandedStruct = true
	// set `jsonschema:required` tag to true to generate required fields
	r.RequiredFromJSONSchemaTags = true
	// set the tag name to `koanf` for the koanf struct tags
	r.FieldNameTag = tagName

	if err := r.AddGoComments("github.com/datumforge/datum/", "./"); err != nil {
		panic(err.Error())
	}

	s := r.Reflect(structure)

	// generate the json schema
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	if err = os.WriteFile(c.jsonSchemaPath, data, ownerReadWrite); err != nil {
		panic(err.Error())
	}

	// generate yaml schema with default
	yamlConfig := &marketplace.Config{}
	defaults.SetDefaults(yamlConfig)

	// this uses the `json` tag to generate the yaml schema
	yamlSchema, err := yaml.Marshal(yamlConfig)
	if err != nil {
		panic(err.Error())
	}

	if err = os.WriteFile(c.yamlConfigPath, yamlSchema, ownerReadWrite); err != nil {
		panic(err.Error())
	}

	return nil
}
