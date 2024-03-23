package oas

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/invopop/jsonschema"
	"github.com/rs/zerolog/log"
)

func getHandlers(handlers.LoginRequest) handlers.LoginRequest {
	return handlers.LoginRequest{}
}

func MagicParameters() error {
	reflector := new(jsonschema.Reflector)

	if err := reflector.AddGoComments("github.com/datumforge/datum/", "./internal/httpserve/handlers"); err != nil {
		log.Error().Err(err).Msg("failed to add go comments")
		return err
	}

	meow, _ := GetFunctionParametersJSONSchema(reflector, getHandlers)
	s, _ := json.MarshalIndent(meow, "", " ")
	fmt.Printf("getWeatherJsonSchema:\n%s\n\n", s)

	return nil
}

type Callable interface{}

func GetFunctionParametersJSONSchema(reflector *jsonschema.Reflector, f Callable) (*jsonschema.Schema, error) {
	// Get the type of the function
	funcVal := reflect.ValueOf(f)
	funcType := funcVal.Type()

	// Check if the function is indeed callable
	if funcType.Kind() != reflect.Func {
		return nil, fmt.Errorf("provided callable is not a function")
	}

	// Handle the case of a single parameter separately
	if funcType.NumIn() == 1 {
		singleParamType := funcType.In(0)
		singleParamInstance := reflect.New(singleParamType).Elem().Interface()
		return reflector.Reflect(singleParamInstance), nil
	}

	// Prepare a schema for multiple function parameters
	schema := &jsonschema.Schema{
		Type:  "array",
		Items: &jsonschema.Schema{},
	}

	// Create a slice to hold schemas for each parameter
	paramSchemas := make([]*jsonschema.Schema, 0, funcType.NumIn())

	// Loop over the function's input parameters
	for i := 0; i < funcType.NumIn(); i++ {
		paramType := funcType.In(i)
		paramInstance := reflect.New(paramType).Elem().Interface()
		paramSchema := reflector.Reflect(paramInstance)
		paramSchemas = append(paramSchemas, paramSchema)
	}

	// Use PrefixItems to define schemas for each parameter in the array
	schema.PrefixItems = paramSchemas

	return schema, nil
}
