package emitter

import (
	"testing"
)

func TestNewBaseEvent(t *testing.T) {
	payload := map[string]string{"key": "value"} // Payload is a map
	event := NewBaseEvent("test_topic", payload)

	if event.Topic() != "test_topic" {
		t.Errorf("NewBaseEvent() topic = %s; want test_topic", event.Topic())
	}

	retrievedPayload, ok := event.Payload().(map[string]string)
	if !ok {
		t.Fatalf("Payload is not of type map[string]string")
	}

	if retrievedPayload["key"] != "value" {
		t.Errorf("NewBaseEvent() payload = %v; want %v", event.Payload(), payload)
	}
}

func TestBaseEventSetAbortedAndIsAborted(t *testing.T) {
	type Payload struct {
		Data string
	}

	event := NewBaseEvent("test_topic", Payload{Data: "some data"}) // Payload is a struct

	if event.IsAborted() {
		t.Errorf("Newly created event should not be aborted")
	}

	event.SetAborted(true)

	if !event.IsAborted() {
		t.Errorf("BaseEvent.Abort(true) did not abort the event")
	}

	event.SetAborted(false)

	if event.IsAborted() {
		t.Errorf("BaseEvent.Abort(false) did not unabort the event")
	}
}
