package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
)

// EventPublisher publishes an event to the configured topic in the message payload - today this can be anything but there is no event consumer on the other side yet
func (h *Handler) EventPublisher(ctx echo.Context) error {
	var in models.PublishRequest
	if err := ctx.Bind(&in); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	if err := h.EventManager.Publish(in.Topic, []byte(in.Message)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	out := &models.PublishReply{
		Reply:   rout.Reply{Success: true},
		Message: "success!",
	}

	return ctx.JSON(http.StatusOK, out)
}

// BindEventPublisher is used to bind the event publisher endpoint to the OpenAPI schema
func (h *Handler) BindEventPublisher() *openapi3.Operation {
	eventCreate := openapi3.NewOperation()
	eventCreate.Description = "Publish and Correleate Events"
	eventCreate.OperationID = "EventPublisher"

	h.AddRequestBody("PublishRequest", models.PublishRequest{}, eventCreate)
	h.AddResponse("PublishReply", "success", models.PublishReply{}, eventCreate, http.StatusOK)
	eventCreate.AddResponse(http.StatusInternalServerError, internalServerError())
	eventCreate.AddResponse(http.StatusBadRequest, badRequest())

	return eventCreate
}
