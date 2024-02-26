package main

import (
	"os"
	"path/filepath"

	"github.com/go-openapi/spec"
	"gopkg.in/yaml.v2"
)

func getOpenAPISpecFromFile() (*spec.Swagger, error) {
	path := filepath.Join("testdata", "swagger.json")
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	byteSpec, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	staticSpec := &spec.Swagger{}

	err = yaml.Unmarshal(byteSpec, staticSpec)
	if err != nil {
		return nil, err
	}

	return staticSpec, nil
}
