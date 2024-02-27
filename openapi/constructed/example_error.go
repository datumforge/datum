package main

import "github.com/datumforge/datum/openapi/payloads"

var ResponseErrorGenericExample = payloads.ResponseError{
	Error:     "error: this can be pretty long string",
	Version:   "df8a489",
}

var ResponseNotFoundErrorExample = payloads.ResponseError{
	Error:     "error: resource not found: details can be long",
	Version:   "df8a489",
}

var ResponseBadRequestErrorExample = payloads.ResponseError{
	Error:     "error: bad request: details can be long",
	Version:   "df8a489",
}

var ResponseErrorUserFriendlyExample = payloads.ResponseError{
	Message:   "vCPU limit reached, contact AWS support",
	Error:     "meow",
	Version:   "df8a489"
}
