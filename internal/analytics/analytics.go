package analytics

import (
	ph "github.com/posthog/posthog-go"

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

type EventManager struct {
	EM Handler
}

type Handler interface {
	Event(eventName string, properties ph.Properties)
	AssociateUser(userID string, organizationID string)
	OrganizationEvent(organizationID, userID, eventName string, properties ph.Properties)
	NewOrganization(organizationID, userID string, properties ph.Properties)
	OrganizationProperties(organizationID string, properties ph.Properties)
	UserEvent(userID, eventName string, properties ph.Properties)
	NewUser(userID string, properties ph.Properties)
	UserProperties(userID string, properties ph.Properties)
	Cleanup()
}

// Event function is used to send an event to the analytics handler
func Event(eventName string, properties ph.Properties) {
	if handler != nil {
		handler.Event(eventName, properties)
	}
}

func UserEvent(userID, eventName string, properties ph.Properties) {
	if handler != nil {
		handler.UserEvent(userID, eventName, properties)
	}
}

// AssociateUser function is used to associate a user with an organization in the analytics handler
func AssociateUser(userID string, organizationID string) {
	if handler != nil {
		handler.AssociateUser(userID, organizationID)
	}
}

func NewOrganization(organizationID, userID string, properties ph.Properties) {
	if handler != nil {
		handler.NewOrganization(organizationID, userID, properties)
	}
}

func OrganizationProperties(organizationID string, properties ph.Properties) {
	if handler != nil {
		handler.OrganizationProperties(organizationID, properties)
	}
}

func OrganizationEvent(organizationID, userID, eventName string, properties ph.Properties) {
	if handler != nil {
		handler.OrganizationEvent(organizationID, userID, eventName, properties)
	}
}

func NewUser(userID string, properties ph.Properties) {
	if handler != nil {
		handler.NewUser(userID, properties)
	}
}

func UserProperties(userID string, properties ph.Properties) {
	if handler != nil {
		handler.UserProperties(userID, properties)
	}
}

// Cleanup is responsible for cleanup
func Cleanup() {
	if handler != nil {
		handler.Cleanup()
		handler = nil
	}
}
