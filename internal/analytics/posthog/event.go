package posthog

import (
	"os"
	"time"

	"github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/analytics/machine"
)

var (
	// PosthogAPIKey is the PostHog API Key
	PosthogAPIKey = os.Getenv("POSTHOG_API_KEY")

	// PosthogAPIHost is the PostHog API Host
	PosthogAPIHost = "https://app.posthog.com"
)

type PostHog struct {
	client     posthog.Client
	Identifier string
}

// Init returns a pointer to a PostHog object
func Init() *PostHog {
	if PosthogAPIKey == "" || PosthogAPIHost == "" || !machine.Available() {
		return nil
	}

	client, _ := posthog.NewWithConfig(PosthogAPIKey, posthog.Config{
		Endpoint:  PosthogAPIHost,
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
func (p *PostHog) Event(eventName string, properties map[string]string) {
	if PosthogAPIKey == "" {
		return
	}

	c := posthog.Capture{
		DistinctId: p.Identifier,
		Event:      eventName,
		Timestamp:  time.Now(),
	}

	if len(properties) > 0 {
		props := posthog.NewProperties()
		for k, v := range properties {
			props.Set(k, v)
		}

		c.Properties = props
	}

	_ = p.client.Enqueue(c)
}

// AssociateUser function is used to associate a user with an organization in PostHog
func (p *PostHog) AssociateUser(userID string, organizationID string) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: p.Identifier,
		Event:      "authentication",
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
func (p *PostHog) OrganizationEvent(organizationID, eventName string) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: p.Identifier,
		Event:      eventName,
		Timestamp:  time.Now(),
		Groups: posthog.NewGroups().
			Set("organization", organizationID),
	})
}

// NewOrganization uses the NewGroups reference to create a new organization in the organization groups category, and also sets attributes for the organization
func (p *PostHog) NewOrganization(organizationID string, properties map[string]string) {
	_ = p.client.Enqueue(posthog.Capture{
		DistinctId: p.Identifier,
		Event:      "organization_created",
		Timestamp:  time.Now(),
		Groups: posthog.NewGroups().
			Set("organization", organizationID),
	})

	c := posthog.GroupIdentify{
		Type:      "organization",
		Key:       organizationID,
		Timestamp: time.Now(),
	}

	if len(properties) > 0 {
		props := posthog.NewProperties()
		for k, v := range properties {
			props.Set(k, v)
		}

		c.Properties = props
	}

	_ = p.client.Enqueue(c)

}

func (p *PostHog) Cleanup() {
	_ = p.client.Close()
}
