// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/webhook"
)

// WebhookCreate is the builder for creating a Webhook entity.
type WebhookCreate struct {
	config
	mutation *WebhookMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (wc *WebhookCreate) SetCreatedAt(t time.Time) *WebhookCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableCreatedAt(t *time.Time) *WebhookCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WebhookCreate) SetUpdatedAt(t time.Time) *WebhookCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableUpdatedAt(t *time.Time) *WebhookCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetCreatedBy sets the "created_by" field.
func (wc *WebhookCreate) SetCreatedBy(s string) *WebhookCreate {
	wc.mutation.SetCreatedBy(s)
	return wc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableCreatedBy(s *string) *WebhookCreate {
	if s != nil {
		wc.SetCreatedBy(*s)
	}
	return wc
}

// SetUpdatedBy sets the "updated_by" field.
func (wc *WebhookCreate) SetUpdatedBy(s string) *WebhookCreate {
	wc.mutation.SetUpdatedBy(s)
	return wc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableUpdatedBy(s *string) *WebhookCreate {
	if s != nil {
		wc.SetUpdatedBy(*s)
	}
	return wc
}

// SetMappingID sets the "mapping_id" field.
func (wc *WebhookCreate) SetMappingID(s string) *WebhookCreate {
	wc.mutation.SetMappingID(s)
	return wc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableMappingID(s *string) *WebhookCreate {
	if s != nil {
		wc.SetMappingID(*s)
	}
	return wc
}

// SetTags sets the "tags" field.
func (wc *WebhookCreate) SetTags(s []string) *WebhookCreate {
	wc.mutation.SetTags(s)
	return wc
}

// SetDeletedAt sets the "deleted_at" field.
func (wc *WebhookCreate) SetDeletedAt(t time.Time) *WebhookCreate {
	wc.mutation.SetDeletedAt(t)
	return wc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableDeletedAt(t *time.Time) *WebhookCreate {
	if t != nil {
		wc.SetDeletedAt(*t)
	}
	return wc
}

// SetDeletedBy sets the "deleted_by" field.
func (wc *WebhookCreate) SetDeletedBy(s string) *WebhookCreate {
	wc.mutation.SetDeletedBy(s)
	return wc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableDeletedBy(s *string) *WebhookCreate {
	if s != nil {
		wc.SetDeletedBy(*s)
	}
	return wc
}

// SetOwnerID sets the "owner_id" field.
func (wc *WebhookCreate) SetOwnerID(s string) *WebhookCreate {
	wc.mutation.SetOwnerID(s)
	return wc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableOwnerID(s *string) *WebhookCreate {
	if s != nil {
		wc.SetOwnerID(*s)
	}
	return wc
}

// SetName sets the "name" field.
func (wc *WebhookCreate) SetName(s string) *WebhookCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WebhookCreate) SetDescription(s string) *WebhookCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableDescription(s *string) *WebhookCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetDestinationURL sets the "destination_url" field.
func (wc *WebhookCreate) SetDestinationURL(s string) *WebhookCreate {
	wc.mutation.SetDestinationURL(s)
	return wc
}

// SetEnabled sets the "enabled" field.
func (wc *WebhookCreate) SetEnabled(b bool) *WebhookCreate {
	wc.mutation.SetEnabled(b)
	return wc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableEnabled(b *bool) *WebhookCreate {
	if b != nil {
		wc.SetEnabled(*b)
	}
	return wc
}

// SetCallback sets the "callback" field.
func (wc *WebhookCreate) SetCallback(s string) *WebhookCreate {
	wc.mutation.SetCallback(s)
	return wc
}

// SetNillableCallback sets the "callback" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableCallback(s *string) *WebhookCreate {
	if s != nil {
		wc.SetCallback(*s)
	}
	return wc
}

// SetExpiresAt sets the "expires_at" field.
func (wc *WebhookCreate) SetExpiresAt(t time.Time) *WebhookCreate {
	wc.mutation.SetExpiresAt(t)
	return wc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableExpiresAt(t *time.Time) *WebhookCreate {
	if t != nil {
		wc.SetExpiresAt(*t)
	}
	return wc
}

// SetSecret sets the "secret" field.
func (wc *WebhookCreate) SetSecret(b []byte) *WebhookCreate {
	wc.mutation.SetSecret(b)
	return wc
}

// SetFailures sets the "failures" field.
func (wc *WebhookCreate) SetFailures(i int) *WebhookCreate {
	wc.mutation.SetFailures(i)
	return wc
}

// SetNillableFailures sets the "failures" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableFailures(i *int) *WebhookCreate {
	if i != nil {
		wc.SetFailures(*i)
	}
	return wc
}

// SetLastError sets the "last_error" field.
func (wc *WebhookCreate) SetLastError(s string) *WebhookCreate {
	wc.mutation.SetLastError(s)
	return wc
}

// SetNillableLastError sets the "last_error" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableLastError(s *string) *WebhookCreate {
	if s != nil {
		wc.SetLastError(*s)
	}
	return wc
}

// SetLastResponse sets the "last_response" field.
func (wc *WebhookCreate) SetLastResponse(s string) *WebhookCreate {
	wc.mutation.SetLastResponse(s)
	return wc
}

// SetNillableLastResponse sets the "last_response" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableLastResponse(s *string) *WebhookCreate {
	if s != nil {
		wc.SetLastResponse(*s)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WebhookCreate) SetID(s string) *WebhookCreate {
	wc.mutation.SetID(s)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableID(s *string) *WebhookCreate {
	if s != nil {
		wc.SetID(*s)
	}
	return wc
}

// SetOwner sets the "owner" edge to the Organization entity.
func (wc *WebhookCreate) SetOwner(o *Organization) *WebhookCreate {
	return wc.SetOwnerID(o.ID)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (wc *WebhookCreate) AddEventIDs(ids ...string) *WebhookCreate {
	wc.mutation.AddEventIDs(ids...)
	return wc
}

// AddEvents adds the "events" edges to the Event entity.
func (wc *WebhookCreate) AddEvents(e ...*Event) *WebhookCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return wc.AddEventIDs(ids...)
}

// AddIntegrationIDs adds the "integrations" edge to the Integration entity by IDs.
func (wc *WebhookCreate) AddIntegrationIDs(ids ...string) *WebhookCreate {
	wc.mutation.AddIntegrationIDs(ids...)
	return wc
}

// AddIntegrations adds the "integrations" edges to the Integration entity.
func (wc *WebhookCreate) AddIntegrations(i ...*Integration) *WebhookCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wc.AddIntegrationIDs(ids...)
}

// Mutation returns the WebhookMutation object of the builder.
func (wc *WebhookCreate) Mutation() *WebhookMutation {
	return wc.mutation
}

// Save creates the Webhook in the database.
func (wc *WebhookCreate) Save(ctx context.Context) (*Webhook, error) {
	if err := wc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WebhookCreate) SaveX(ctx context.Context) *Webhook {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WebhookCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WebhookCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WebhookCreate) defaults() error {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		if webhook.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized webhook.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := webhook.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		if webhook.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized webhook.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := webhook.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.MappingID(); !ok {
		if webhook.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized webhook.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := webhook.DefaultMappingID()
		wc.mutation.SetMappingID(v)
	}
	if _, ok := wc.mutation.Tags(); !ok {
		v := webhook.DefaultTags
		wc.mutation.SetTags(v)
	}
	if _, ok := wc.mutation.Enabled(); !ok {
		v := webhook.DefaultEnabled
		wc.mutation.SetEnabled(v)
	}
	if _, ok := wc.mutation.Failures(); !ok {
		v := webhook.DefaultFailures
		wc.mutation.SetFailures(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		if webhook.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized webhook.DefaultID (forgotten import generated/runtime?)")
		}
		v := webhook.DefaultID()
		wc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wc *WebhookCreate) check() error {
	if _, ok := wc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "Webhook.mapping_id"`)}
	}
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Webhook.name"`)}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := webhook.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Webhook.name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.DestinationURL(); !ok {
		return &ValidationError{Name: "destination_url", err: errors.New(`generated: missing required field "Webhook.destination_url"`)}
	}
	if v, ok := wc.mutation.DestinationURL(); ok {
		if err := webhook.DestinationURLValidator(v); err != nil {
			return &ValidationError{Name: "destination_url", err: fmt.Errorf(`generated: validator failed for field "Webhook.destination_url": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`generated: missing required field "Webhook.enabled"`)}
	}
	return nil
}

func (wc *WebhookCreate) sqlSave(ctx context.Context) (*Webhook, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Webhook.ID type: %T", _spec.ID.Value)
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WebhookCreate) createSpec() (*Webhook, *sqlgraph.CreateSpec) {
	var (
		_node = &Webhook{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(webhook.Table, sqlgraph.NewFieldSpec(webhook.FieldID, field.TypeString))
	)
	_spec.Schema = wc.schemaConfig.Webhook
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(webhook.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.SetField(webhook.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wc.mutation.CreatedBy(); ok {
		_spec.SetField(webhook.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := wc.mutation.UpdatedBy(); ok {
		_spec.SetField(webhook.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := wc.mutation.MappingID(); ok {
		_spec.SetField(webhook.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := wc.mutation.Tags(); ok {
		_spec.SetField(webhook.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := wc.mutation.DeletedAt(); ok {
		_spec.SetField(webhook.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := wc.mutation.DeletedBy(); ok {
		_spec.SetField(webhook.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(webhook.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(webhook.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.DestinationURL(); ok {
		_spec.SetField(webhook.FieldDestinationURL, field.TypeString, value)
		_node.DestinationURL = value
	}
	if value, ok := wc.mutation.Enabled(); ok {
		_spec.SetField(webhook.FieldEnabled, field.TypeBool, value)
		_node.Enabled = value
	}
	if value, ok := wc.mutation.Callback(); ok {
		_spec.SetField(webhook.FieldCallback, field.TypeString, value)
		_node.Callback = value
	}
	if value, ok := wc.mutation.ExpiresAt(); ok {
		_spec.SetField(webhook.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := wc.mutation.Secret(); ok {
		_spec.SetField(webhook.FieldSecret, field.TypeBytes, value)
		_node.Secret = value
	}
	if value, ok := wc.mutation.Failures(); ok {
		_spec.SetField(webhook.FieldFailures, field.TypeInt, value)
		_node.Failures = value
	}
	if value, ok := wc.mutation.LastError(); ok {
		_spec.SetField(webhook.FieldLastError, field.TypeString, value)
		_node.LastError = value
	}
	if value, ok := wc.mutation.LastResponse(); ok {
		_spec.SetField(webhook.FieldLastResponse, field.TypeString, value)
		_node.LastResponse = value
	}
	if nodes := wc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webhook.OwnerTable,
			Columns: []string{webhook.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = wc.schemaConfig.Webhook
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   webhook.EventsTable,
			Columns: webhook.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = wc.schemaConfig.WebhookEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.IntegrationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   webhook.IntegrationsTable,
			Columns: webhook.IntegrationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = wc.schemaConfig.IntegrationWebhooks
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebhookCreateBulk is the builder for creating many Webhook entities in bulk.
type WebhookCreateBulk struct {
	config
	err      error
	builders []*WebhookCreate
}

// Save creates the Webhook entities in the database.
func (wcb *WebhookCreateBulk) Save(ctx context.Context) ([]*Webhook, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Webhook, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebhookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WebhookCreateBulk) SaveX(ctx context.Context) []*Webhook {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WebhookCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WebhookCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
