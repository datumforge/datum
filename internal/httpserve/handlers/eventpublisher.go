package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/pkg/rout"
)

// PublishRequest is the request payload for the event publisher
type PublishRequest struct {
	Tags    map[string]string `json:"tags"`
	Topic   string            `json:"topic"`
	Message string            `json:"message"`
}

// PublishReply holds the fields that are sent on a response to the `/event/publish` endpoint
type PublishReply struct {
	rout.Reply
	Message string `json:"message"`
}

// EventPublisher publishes an event to the configured topic in the message payload - today this can be anything but there is no event consumer on the other side yet
func (h *Handler) EventPublisher(ctx echo.Context) error {
	var in PublishRequest
	if err := ctx.Bind(&in); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	if err := h.EventManager.Publish(in.Topic, []byte(in.Message)); err != nil {
		return ctx.JSON(http.StatusConflict, rout.ErrorResponse(err))
	}

	out := &PublishReply{
		Reply:   rout.Reply{Success: true},
		Message: "success!",
	}

	return ctx.JSON(http.StatusOK, out)
}
