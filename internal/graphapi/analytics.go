package graphapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	ent "github.com/datumforge/datum/internal/ent/generated"
	ph "github.com/posthog/posthog-go"
)

// CreateEvent creates an event for the mutation with the properties
func CreateEvent(ctx context.Context, c *ent.Client, m ent.Mutation, v ent.Value) {
	out, err := parseValue(v)
	if err != nil {
		return
	}

	obj := strings.ToLower(m.Type())
	action := getOp(m)

	// debug log the event
	c.Logger.Debugw("tracking event", "object", obj, "action", action)

	event := fmt.Sprintf("%s_%sd", obj, action)

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

	// set the name if it exists
	name, ok := out["name"]
	if ok {
		props.Set(fmt.Sprintf("%s_name", obj), name)
	}

	// set the first name if it exists
	fName, ok := out["first_name"]
	if ok {
		props.Set("first_name", fName)
	}

	// set the last name if it exists
	lName, ok := out["last_name"]
	if ok {
		props.Set("last_name", lName)
	}

	// set the email if it exists
	email, ok := out["email"]
	if ok {
		props.Set("email", email)
	}

	// check if the organization has posthog enabled
	// this should be a query to the `integration` by Organization and Integration name
	// for now, lets just assume the organization has posthog enabled
	// and send the event with _our_ token
	// this is a test of creating the client on the fly

	// Example of how we would query the integration for the db
	// viewer := viewer.FromContext(ctx)
	// int, err := c.Integration.Query().Where(
	// 	integration.OwnerIDEQ(viewer.GetOrganizationID()),
	// 	integration.NameEQ("posthog"),
	// ).Only(ctx)
	// if err != nil {
	// 	// keep going
	// 	return
	// }

	// phc := posthog.Config{
	// 	APIKey: int.Config // this field doesn't exist, but whatever we are storing the config for the integration
	// }

	// for now, lets use our config for the demo
	phc := c.Analytics.EventConfig

	// create the posthog client

	em := phc.Init()
	defer em.Cleanup()

	em.Event(event, props)
	em.Properties(i, obj, props)

	// debug log the event
	c.Logger.Debugw("event tracked", "event", event, "props", props)
}

// trackedEvent returns true if the mutation should be a tracked event
// for now, lets just track high level create and delete events
// TODO: make these configurable by integration
func TrackedEvent(m ent.Mutation) bool {
	switch m.Type() {
	case "User", "Organization", "Group":
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
