package graphapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	ph "github.com/posthog/posthog-go"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/events/soiree"
	"github.com/datumforge/datum/pkg/utils/slack"
)

// CreateEvent creates an event for the mutation with the properties
func CreateEvent(c *ent.Client, m ent.Mutation, v ent.Value) {
	pool := soiree.NewPondPool(100, 1000)
	e := soiree.NewEventPool(soiree.WithPool(pool))

	out, err := parseValue(v)
	if err != nil {
		return
	}

	obj := strings.ToLower(m.Type())
	action := getOp(m)

	// debug log the event
	c.Logger.Debugw("tracking event", "object", obj, "action", action)

	event := fmt.Sprintf("%s.%sd", obj, action)
	e.EnsureTopic(event)

	id, ok := out["id"]
	if !ok {
		// keep going
		return
	}

	i, ok := id.(string)
	if !ok {
		// keep going
		return
	}

	// Set properties for the event
	// all events will have the id
	props := ph.NewProperties().
		Set(fmt.Sprintf("%s_id", obj), i)

	payload := map[string]string{"key": "value"}
	sEvent := soiree.NewBaseEvent(event, payload)

	soireeProps := soiree.NewProperties().Set(fmt.Sprintf("%s_id", obj), i)
	// set the name if it exists
	name, ok := out["name"]
	if ok {
		props.Set(fmt.Sprintf("%s_name", obj), name)
		soireeProps.Set(fmt.Sprintf("%s_name", obj), name)
		payload["name"] = name.(string)
	}

	// set the first name if it exists
	fName, ok := out["first_name"]
	if ok {
		props.Set("first_name", fName)
		soireeProps.Set("first_name", fName)
		payload["first_name"] = fName.(string)
	}

	// set the last name if it exists
	lName, ok := out["last_name"]
	if ok {
		props.Set("last_name", lName)
		soireeProps.Set("last_name", lName)
		payload["last_name"] = lName.(string)
	}

	// set the email if it exists
	email, ok := out["email"]
	if ok {
		props.Set("email", email)
		soireeProps.Set("email", email)
		payload["email"] = email.(string)
	}

	authprovider, ok := out["auth_provider"]
	if ok {
		props.Set("auth_provider", authprovider)
		soireeProps.Set("auth_provider", authprovider)
		payload["auth_provider"] = authprovider.(string)
	}

	userCreatedListener := func(evt soiree.Event) error {
		webhookURL := ""
		retrieve := sEvent.Payload().(map[string]string)
		log.Printf("event: %s\n", retrieve["key"])

		payload := slack.Payload{
			Text: fmt.Sprintf("A user with the following details has been created:\nName: %s\nFirst Name: %s\nLast Name: %s\nEmail: %s\nAuth Provider: %s", retrieve["name"], retrieve["first_name"], retrieve["last_name"], retrieve["email"], retrieve["auth_provider"]),
		}

		slackMessage := slack.New(webhookURL)
		if err := slackMessage.Post(context.Background(), &payload); err != nil {
			log.Printf("error: %s\n", err)
		}
		return nil
	}

	e.On("user.created", userCreatedListener)
	e.Emit(event, soireeProps)

	c.Analytics.Event(event, props)
	c.Analytics.Properties(i, obj, props)

	// debug log the event
	c.Logger.Debugw("event tracked", "event", event, "props", props)
}

// trackedEvent returns true if the mutation should be a tracked event
// for now, lets just track high level create and delete events
// TODO: make these configurable by integration
func TrackedEvent(m ent.Mutation) bool {
	switch m.Type() {
	case "User", "Organization", "Group", "Subscriber":
		switch getOp(m) {
		case ActionCreate, ActionDelete:
			return true
		}
		return false
	}

	return false
}

// getOp returns the string action for the mutation
func getOp(m ent.Mutation) string {
	switch m.Op() {
	case ent.OpCreate:
		return ActionCreate
	case ent.OpUpdate, ent.OpUpdateOne:
		return ActionUpdate
	case ent.OpDelete, ent.OpDeleteOne:
		return ActionDelete
	default:
		return ""
	}
}

// parseValue returns a map of the ent.Value
func parseValue(v ent.Value) (map[string]interface{}, error) {
	out, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var valMap map[string]interface{}

	if err := json.Unmarshal(out, &valMap); err != nil {
		return nil, err
	}

	return valMap, nil
}
