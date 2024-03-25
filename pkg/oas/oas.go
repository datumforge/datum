package oas

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/invopop/jsonschema"
	"github.com/ogen-go/ogen"
	"github.com/rs/zerolog/log"
)

func MagicParameters() error {
	reflector := new(jsonschema.Reflector)

	if err := reflector.AddGoComments("github.com/datumforge/datum/", "./internal/httpserve/handlers"); err != nil {
		log.Error().Err(err).Msg("failed to add go comments")
		return err
	}

	getLoginRequestSchema, _ := GetFunctionParametersJSONSchema(reflector, getLoginRequest)
	ogenSpec := ogen.NewSpec()
	matt, _ := getLoginRequestSchema.MarshalJSON()
	ogenSchema := ogen.NewSchema()
	getLoginRequestSchema.MarshalJSON()

	ogenSpec.AddSchema("LoginRequest", matt)
	s, _ := json.MarshalIndent(getLoginRequestSchema, "", " ")
	fmt.Printf("loginRequest:\n%s\n\n", s)

	getLoginReplySchema, _ := GetFunctionParametersJSONSchema(reflector, getLoginReply)
	s, _ = json.MarshalIndent(getLoginReplySchema, "", " ")
	fmt.Printf("loginReply:\n%s\n\n", s)

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

func getLoginRequest(handlers.LoginRequest) handlers.LoginRequest {
	return handlers.LoginRequest{}
}

func getLoginReply(handlers.LoginReply) handlers.LoginReply {
	return handlers.LoginReply{}
}

func getRefreshRequest(handlers.RefreshRequest) handlers.RefreshRequest {
	return handlers.RefreshRequest{}
}

func getRefreshReply(handlers.RefreshReply) handlers.RefreshReply {
	return handlers.RefreshReply{}
}

func getRegisterRequest(handlers.RegisterRequest) handlers.RegisterRequest {
	return handlers.RegisterRequest{}
}

func getRegisterReply(handlers.RegisterReply) handlers.RegisterReply {
	return handlers.RegisterReply{}
}

func getForgotPasswordRequest(handlers.ForgotPasswordRequest) handlers.ForgotPasswordRequest {
	return handlers.ForgotPasswordRequest{}
}

func getForgotPasswordReply(handlers.ForgotPasswordReply) handlers.ForgotPasswordReply {
	return handlers.ForgotPasswordReply{}
}

func getVerifyRequest(handlers.VerifyRequest) handlers.VerifyRequest {
	return handlers.VerifyRequest{}
}

func getVerifyReply(handlers.VerifyReply) handlers.VerifyReply {
	return handlers.VerifyReply{}
}

func getResendEmailRequest(handlers.ResendRequest) handlers.ResendRequest {
	return handlers.ResendRequest{}
}

func getResendEmailReply(handlers.ResendReply) handlers.ResendReply {
	return handlers.ResendReply{}
}

func getSubscribeReply(handlers.SubscribeReply) handlers.SubscribeReply {
	return handlers.SubscribeReply{}
}

func getVerifySubscribeReply(handlers.VerifySubscribeReply) handlers.VerifySubscribeReply {
	return handlers.VerifySubscribeReply{}
}

func getUnsubscribeReply(handlers.UnsubscribeReply) handlers.UnsubscribeReply {
	return handlers.UnsubscribeReply{}
}

func getInviteReply(handlers.InviteReply) handlers.InviteReply {
	return handlers.InviteReply{}
}

func getWebauthnRegistrationRequest(handlers.WebauthnRegistrationRequest) handlers.WebauthnRegistrationRequest {
	return handlers.WebauthnRegistrationRequest{}
}
