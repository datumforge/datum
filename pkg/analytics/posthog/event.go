package posthog

import (
	"time"

	"github.com/posthog/posthog-go"

	"github.com/datumforge/datum/pkg/analytics/machine"
)

// Config is the configuration for PostHog
type Config struct {
	// Enabled is a flag to enable or disable PostHog
	Enabled bool `json:"enabled" koanf:"enabled" default:"false"`
	// APIKey is the PostHog API Key
	APIKey string `json:"apiKey" koanf:"apiKey"`
	// Host is the PostHog API Host
	Host string `json:"host" koanf:"host" default:"https://app.posthog.com"`
}

type PostHog struct {
	client     posthog.Client
	Identifier string
}

// Init returns a pointer to a PostHog object
func (c *Config) Init() *PostHog {
	if !c.Enabled || c.APIKey == "" || c.Host == "" || !machine.Available() {
		return nil
	}

	client, _ := posthog.NewWithConfig(c.APIKey, posthog.Config{
		Endpoint:  c.Host,
		BatchSize: 1,
		Logger:    new(noopLogger),
	})

	if client != nil {
		return &PostHog{
			client:     client,
			Identifier: machine.ID(),
		}
	}

	return nil
}

// Event is used to send an event to PostHog
func (p *PostHog) Event(eventName string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: p.Identifier,
		Event:      eventName,
		Timestamp:  time.Now(),
		Properties: properties,
	})
}

// UserEvent captures user properties
func (p *PostHog) UserEvent(userID, eventName string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      eventName,
		Timestamp:  time.Now(),
		Properties: properties,
	})
}

// AssociateUser function is used to associate a user with an organization in PostHog
func (p *PostHog) AssociateUser(userID string, organizationID string) {
	_ = p.client.Enqueue(posthog.GroupIdentify{
		DistinctId: organizationID,
		Type:       "user",
		Key:        userID,
		Timestamp:  time.Now(),
		Properties: map[string]interface{}{
			"$set": map[string]interface{}{
				"user":         userID,
				"organization": organizationID,
			},
		},
	})
}

// OrganizationEvent creates an event associated with the organization, where the eventName can be passed in generically and associated with the org ID if provided
func (p *PostHog) OrganizationEvent(organizationID, userID, eventName string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      eventName,
		Timestamp:  time.Now(),
		Properties: properties,
		Groups: posthog.NewGroups().
			Set("organization", organizationID),
	})
}

// GroupEvent creates an event associated with the group, where the eventName can be passed in generically and associated with the group ID if provided
func (p *PostHog) GroupEvent(groupID, userID, eventName string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      eventName,
		Timestamp:  time.Now(),
		Properties: properties,
		Groups: posthog.NewGroups().
			Set("group", groupID),
	})
}

// NewOrganization uses the NewGroups reference to create a new organization in the organization groups category, and also sets attributes for the organization
func (p *PostHog) NewOrganization(organizationID, userID string, properties posthog.Properties) {
	// this event is creating the organization and associating it with our internal organization ID
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      "organization_created",
		Timestamp:  time.Now(),
		Groups: posthog.NewGroups().
			Set("organization", organizationID),
	})

	// this is attempting to set
	_ = p.client.Enqueue(posthog.GroupIdentify{
		Type:       "organization",
		Key:        organizationID,
		Timestamp:  time.Now(),
		Properties: properties,
	})
}

// Properties sets generic properties
func (p *PostHog) Properties(id, obj string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.GroupIdentify{
		Type:       obj,
		Key:        id,
		Timestamp:  time.Now(),
		Properties: properties,
	})
}

// OrganizationProperties sets org properties
func (p *PostHog) OrganizationProperties(organizationID string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.GroupIdentify{
		Type:       "organization",
		Key:        organizationID,
		Timestamp:  time.Now(),
		Properties: properties,
	})
}

// UserProperties is to expand the properties of the user in the user group
func (p *PostHog) UserProperties(userID string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.GroupIdentify{
		Type:       "user",
		Key:        userID,
		Timestamp:  time.Now(),
		Properties: properties,
	})
}

// NewUser maps the userID to the user group
func (p *PostHog) NewUser(userID string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: userID,
		Event:      "user_created",
		Timestamp:  time.Now(),
		Groups: posthog.NewGroups().
			Set("user", userID),
	})
}

// NewGroup maps the groupID to the group group
func (p *PostHog) NewGroup(groupID string, properties posthog.Properties) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: groupID,
		Event:      "group_created",
		Timestamp:  time.Now(),
		Groups: posthog.NewGroups().
			Set("group", groupID),
	})
}

// Cleanup cleans up the cleanup
func (p *PostHog) Cleanup() {
	_ = p.client.Close()
}
