package analytics

import (
	"github.com/datumforge/datum/internal/analytics/posthog"
)

var _ Handler = (*posthog.PostHog)(nil)

var (
	handler Handler
)

func init() {
	p := posthog.Init()
	if p != nil {
		handler = p
	}
}

type Handler interface {
	Event(eventName string, properties map[string]string)
	AssociateUser(userID string, organizationID string)
	OrganizationEvent(organization string, eventName string)
	Cleanup()
}

// Event function is used to send an event to the analytics handler
func Event(eventName string, properties ...map[string]string) {
	if handler != nil {
		if len(properties) > 0 {
			handler.Event(eventName, properties[0])
		} else {
			handler.Event(eventName, nil)
		}
	}
}

// AssociateUser function is used to associate a user with an organization in the analytics handler
func AssociateUser(userID string, organizationID string) {
	if handler != nil {
		handler.AssociateUser(userID, organizationID)
	}
}

// Cleanup is responsible for cleanup
func Cleanup() {
	if handler != nil {
		handler.Cleanup()
		handler = nil
	}
}
