package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/invopop/jsonschema"

	config "github.com/datumforge/datum/config"
)

func main() {
	if err := generateSchema("./config", "./jsonschema/datum.config.json", &config.Config{}); err != nil {
		panic(err)
	}
}

// generateSchema function is a helper function that generates a JSON schema file based on a Go
// structure. It takes three parameters: `codePath`, `filePath`, and `structure`
func generateSchema(codePath, filePath string, structure interface{}) error {
	r := new(jsonschema.Reflector)

	if err := r.AddGoComments("github.com/datumforge/datum/", codePath); err != nil {
		panic(err.Error())
	}

	s := r.Reflect(structure)

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile(filePath, data, 0600) // nolint: gomnd
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
