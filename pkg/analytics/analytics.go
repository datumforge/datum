package analytics

import (
	ph "github.com/posthog/posthog-go"
)

// EventManager isn't your normal party planner
type EventManager struct {
	Enabled bool
	Handler Handler
}

// Handler is an interface which can be used to call various event / event association parameters provided by the posthog API
type Handler interface {
	Event(eventName string, properties ph.Properties)
	AssociateUser(userID string, organizationID string)
	OrganizationEvent(organizationID, userID, eventName string, properties ph.Properties)
	NewOrganization(organizationID, userID string, properties ph.Properties)
	OrganizationProperties(organizationID string, properties ph.Properties)
	UserEvent(userID, eventName string, properties ph.Properties)
	NewUser(userID string, properties ph.Properties)
	UserProperties(userID string, properties ph.Properties)
	NewGroup(groupID string, properties ph.Properties)
	GroupEvent(groupID, userID, eventName string, properties ph.Properties)
	Cleanup()
}

// Event function is used to send an event to the analytics handler
func (e *EventManager) Event(eventName string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.Event(eventName, properties)
	}
}

func (e *EventManager) UserEvent(userID, eventName string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.UserEvent(userID, eventName, properties)
	}
}

// AssociateUser function is used to associate a user with an organization in the analytics handler
func (e *EventManager) AssociateUser(userID string, organizationID string) {
	if e.Enabled {
		e.Handler.AssociateUser(userID, organizationID)
	}
}

// NewOrganization is a wrapper for the new organization event
func (e *EventManager) NewOrganization(organizationID, userID string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.NewOrganization(organizationID, userID, properties)
	}
}

// OrganizationProperties is a wrapper to set organization properties
func (e *EventManager) OrganizationProperties(organizationID string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.OrganizationProperties(organizationID, properties)
	}
}

// OrganizationEvent is a generic wrapper for an event you can name which occurs within an organization (e.g. membership)
func (e *EventManager) OrganizationEvent(organizationID, userID, eventName string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.OrganizationEvent(organizationID, userID, eventName, properties)
	}
}

// NewUser is a wrapper for creation of a new user and associating the user with the user group type
func (e *EventManager) NewUser(userID string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.NewUser(userID, properties)
	}
}

// UserProperties is a wrapper to expand the metadata properties associated with a user
func (e *EventManager) UserProperties(userID string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.UserProperties(userID, properties)
	}
}

// NewGroup is a wrapper for creation of a new user and associating the user with the user group type
func (e *EventManager) NewGroup(groupID string, properties ph.Properties) {
	if e.Enabled {
		e.Handler.NewUser(groupID, properties)
	}
}

// Cleanup is responsible for cleanup
func (e *EventManager) Cleanup() {
	if e.Enabled {
		e.Handler.Cleanup()
		e.Handler = nil
	}
}
