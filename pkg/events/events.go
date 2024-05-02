package events

import (
	"context"
)

type EventConsumer interface {
	Subscribe(ctx context.Context) error
	Close(context.Context) error
}

type EventPublisher interface {
	StartPublisher(context.Context) error
	Publish(ctx context.Context, topic string, payload interface{}) error
	Close(context.Context) error
}

type Properties map[string]interface{}

func NewProperties() Properties {
	return make(Properties, 10)
}

func (p Properties) Set(name string, value interface{}) Properties {
	p[name] = value

	return p
}
