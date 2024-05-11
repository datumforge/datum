// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/hush"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootoken"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// IntegrationUpdate is the builder for updating Integration entities.
type IntegrationUpdate struct {
	config
	hooks    []Hook
	mutation *IntegrationMutation
}

// Where appends a list predicates to the IntegrationUpdate builder.
func (iu *IntegrationUpdate) Where(ps ...predicate.Integration) *IntegrationUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *IntegrationUpdate) SetUpdatedAt(t time.Time) *IntegrationUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iu *IntegrationUpdate) ClearUpdatedAt() *IntegrationUpdate {
	iu.mutation.ClearUpdatedAt()
	return iu
}

// SetUpdatedBy sets the "updated_by" field.
func (iu *IntegrationUpdate) SetUpdatedBy(s string) *IntegrationUpdate {
	iu.mutation.SetUpdatedBy(s)
	return iu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableUpdatedBy(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetUpdatedBy(*s)
	}
	return iu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iu *IntegrationUpdate) ClearUpdatedBy() *IntegrationUpdate {
	iu.mutation.ClearUpdatedBy()
	return iu
}

// SetDeletedAt sets the "deleted_at" field.
func (iu *IntegrationUpdate) SetDeletedAt(t time.Time) *IntegrationUpdate {
	iu.mutation.SetDeletedAt(t)
	return iu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableDeletedAt(t *time.Time) *IntegrationUpdate {
	if t != nil {
		iu.SetDeletedAt(*t)
	}
	return iu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iu *IntegrationUpdate) ClearDeletedAt() *IntegrationUpdate {
	iu.mutation.ClearDeletedAt()
	return iu
}

// SetDeletedBy sets the "deleted_by" field.
func (iu *IntegrationUpdate) SetDeletedBy(s string) *IntegrationUpdate {
	iu.mutation.SetDeletedBy(s)
	return iu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableDeletedBy(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetDeletedBy(*s)
	}
	return iu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (iu *IntegrationUpdate) ClearDeletedBy() *IntegrationUpdate {
	iu.mutation.ClearDeletedBy()
	return iu
}

// SetOwnerID sets the "owner_id" field.
func (iu *IntegrationUpdate) SetOwnerID(s string) *IntegrationUpdate {
	iu.mutation.SetOwnerID(s)
	return iu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableOwnerID(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetOwnerID(*s)
	}
	return iu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (iu *IntegrationUpdate) ClearOwnerID() *IntegrationUpdate {
	iu.mutation.ClearOwnerID()
	return iu
}

// SetName sets the "name" field.
func (iu *IntegrationUpdate) SetName(s string) *IntegrationUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableName(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// SetDescription sets the "description" field.
func (iu *IntegrationUpdate) SetDescription(s string) *IntegrationUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableDescription(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// ClearDescription clears the value of the "description" field.
func (iu *IntegrationUpdate) ClearDescription() *IntegrationUpdate {
	iu.mutation.ClearDescription()
	return iu
}

// SetKind sets the "kind" field.
func (iu *IntegrationUpdate) SetKind(s string) *IntegrationUpdate {
	iu.mutation.SetKind(s)
	return iu
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableKind(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetKind(*s)
	}
	return iu
}

// ClearKind clears the value of the "kind" field.
func (iu *IntegrationUpdate) ClearKind() *IntegrationUpdate {
	iu.mutation.ClearKind()
	return iu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (iu *IntegrationUpdate) SetOwner(o *Organization) *IntegrationUpdate {
	return iu.SetOwnerID(o.ID)
}

// AddSecretIDs adds the "secrets" edge to the Hush entity by IDs.
func (iu *IntegrationUpdate) AddSecretIDs(ids ...string) *IntegrationUpdate {
	iu.mutation.AddSecretIDs(ids...)
	return iu
}

// AddSecrets adds the "secrets" edges to the Hush entity.
func (iu *IntegrationUpdate) AddSecrets(h ...*Hush) *IntegrationUpdate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return iu.AddSecretIDs(ids...)
}

// AddOauth2tokenIDs adds the "oauth2tokens" edge to the OhAuthTooToken entity by IDs.
func (iu *IntegrationUpdate) AddOauth2tokenIDs(ids ...string) *IntegrationUpdate {
	iu.mutation.AddOauth2tokenIDs(ids...)
	return iu
}

// AddOauth2tokens adds the "oauth2tokens" edges to the OhAuthTooToken entity.
func (iu *IntegrationUpdate) AddOauth2tokens(o ...*OhAuthTooToken) *IntegrationUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return iu.AddOauth2tokenIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (iu *IntegrationUpdate) AddEventIDs(ids ...string) *IntegrationUpdate {
	iu.mutation.AddEventIDs(ids...)
	return iu
}

// AddEvents adds the "events" edges to the Event entity.
func (iu *IntegrationUpdate) AddEvents(e ...*Event) *IntegrationUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iu.AddEventIDs(ids...)
}

// Mutation returns the IntegrationMutation object of the builder.
func (iu *IntegrationUpdate) Mutation() *IntegrationMutation {
	return iu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (iu *IntegrationUpdate) ClearOwner() *IntegrationUpdate {
	iu.mutation.ClearOwner()
	return iu
}

// ClearSecrets clears all "secrets" edges to the Hush entity.
func (iu *IntegrationUpdate) ClearSecrets() *IntegrationUpdate {
	iu.mutation.ClearSecrets()
	return iu
}

// RemoveSecretIDs removes the "secrets" edge to Hush entities by IDs.
func (iu *IntegrationUpdate) RemoveSecretIDs(ids ...string) *IntegrationUpdate {
	iu.mutation.RemoveSecretIDs(ids...)
	return iu
}

// RemoveSecrets removes "secrets" edges to Hush entities.
func (iu *IntegrationUpdate) RemoveSecrets(h ...*Hush) *IntegrationUpdate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return iu.RemoveSecretIDs(ids...)
}

// ClearOauth2tokens clears all "oauth2tokens" edges to the OhAuthTooToken entity.
func (iu *IntegrationUpdate) ClearOauth2tokens() *IntegrationUpdate {
	iu.mutation.ClearOauth2tokens()
	return iu
}

// RemoveOauth2tokenIDs removes the "oauth2tokens" edge to OhAuthTooToken entities by IDs.
func (iu *IntegrationUpdate) RemoveOauth2tokenIDs(ids ...string) *IntegrationUpdate {
	iu.mutation.RemoveOauth2tokenIDs(ids...)
	return iu
}

// RemoveOauth2tokens removes "oauth2tokens" edges to OhAuthTooToken entities.
func (iu *IntegrationUpdate) RemoveOauth2tokens(o ...*OhAuthTooToken) *IntegrationUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return iu.RemoveOauth2tokenIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (iu *IntegrationUpdate) ClearEvents() *IntegrationUpdate {
	iu.mutation.ClearEvents()
	return iu
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (iu *IntegrationUpdate) RemoveEventIDs(ids ...string) *IntegrationUpdate {
	iu.mutation.RemoveEventIDs(ids...)
	return iu
}

// RemoveEvents removes "events" edges to Event entities.
func (iu *IntegrationUpdate) RemoveEvents(e ...*Event) *IntegrationUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iu.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IntegrationUpdate) Save(ctx context.Context) (int, error) {
	if err := iu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IntegrationUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IntegrationUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IntegrationUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *IntegrationUpdate) defaults() error {
	if _, ok := iu.mutation.UpdatedAt(); !ok && !iu.mutation.UpdatedAtCleared() {
		if integration.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized integration.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := integration.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (iu *IntegrationUpdate) check() error {
	if v, ok := iu.mutation.OwnerID(); ok {
		if err := integration.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Integration.owner_id": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Name(); ok {
		if err := integration.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Integration.name": %w`, err)}
		}
	}
	return nil
}

func (iu *IntegrationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(integration.Table, integration.Columns, sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if iu.mutation.CreatedAtCleared() {
		_spec.ClearField(integration.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(integration.FieldUpdatedAt, field.TypeTime, value)
	}
	if iu.mutation.UpdatedAtCleared() {
		_spec.ClearField(integration.FieldUpdatedAt, field.TypeTime)
	}
	if iu.mutation.CreatedByCleared() {
		_spec.ClearField(integration.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iu.mutation.UpdatedBy(); ok {
		_spec.SetField(integration.FieldUpdatedBy, field.TypeString, value)
	}
	if iu.mutation.UpdatedByCleared() {
		_spec.ClearField(integration.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iu.mutation.DeletedAt(); ok {
		_spec.SetField(integration.FieldDeletedAt, field.TypeTime, value)
	}
	if iu.mutation.DeletedAtCleared() {
		_spec.ClearField(integration.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.DeletedBy(); ok {
		_spec.SetField(integration.FieldDeletedBy, field.TypeString, value)
	}
	if iu.mutation.DeletedByCleared() {
		_spec.ClearField(integration.FieldDeletedBy, field.TypeString)
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(integration.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.SetField(integration.FieldDescription, field.TypeString, value)
	}
	if iu.mutation.DescriptionCleared() {
		_spec.ClearField(integration.FieldDescription, field.TypeString)
	}
	if value, ok := iu.mutation.Kind(); ok {
		_spec.SetField(integration.FieldKind, field.TypeString, value)
	}
	if iu.mutation.KindCleared() {
		_spec.ClearField(integration.FieldKind, field.TypeString)
	}
	if iu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.Integration
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.Integration
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.SecretsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.SecretsTable,
			Columns: integration.SecretsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationSecrets
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedSecretsIDs(); len(nodes) > 0 && !iu.mutation.SecretsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.SecretsTable,
			Columns: integration.SecretsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.SecretsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.SecretsTable,
			Columns: integration.SecretsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.Oauth2tokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.Oauth2tokensTable,
			Columns: integration.Oauth2tokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationOauth2tokens
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedOauth2tokensIDs(); len(nodes) > 0 && !iu.mutation.Oauth2tokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.Oauth2tokensTable,
			Columns: integration.Oauth2tokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationOauth2tokens
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.Oauth2tokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.Oauth2tokensTable,
			Columns: integration.Oauth2tokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationOauth2tokens
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.EventsTable,
			Columns: integration.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationEvents
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedEventsIDs(); len(nodes) > 0 && !iu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.EventsTable,
			Columns: integration.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.EventsTable,
			Columns: integration.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.IntegrationEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = iu.schemaConfig.Integration
	ctx = internal.NewSchemaConfigContext(ctx, iu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IntegrationUpdateOne is the builder for updating a single Integration entity.
type IntegrationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IntegrationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *IntegrationUpdateOne) SetUpdatedAt(t time.Time) *IntegrationUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iuo *IntegrationUpdateOne) ClearUpdatedAt() *IntegrationUpdateOne {
	iuo.mutation.ClearUpdatedAt()
	return iuo
}

// SetUpdatedBy sets the "updated_by" field.
func (iuo *IntegrationUpdateOne) SetUpdatedBy(s string) *IntegrationUpdateOne {
	iuo.mutation.SetUpdatedBy(s)
	return iuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableUpdatedBy(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetUpdatedBy(*s)
	}
	return iuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iuo *IntegrationUpdateOne) ClearUpdatedBy() *IntegrationUpdateOne {
	iuo.mutation.ClearUpdatedBy()
	return iuo
}

// SetDeletedAt sets the "deleted_at" field.
func (iuo *IntegrationUpdateOne) SetDeletedAt(t time.Time) *IntegrationUpdateOne {
	iuo.mutation.SetDeletedAt(t)
	return iuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableDeletedAt(t *time.Time) *IntegrationUpdateOne {
	if t != nil {
		iuo.SetDeletedAt(*t)
	}
	return iuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iuo *IntegrationUpdateOne) ClearDeletedAt() *IntegrationUpdateOne {
	iuo.mutation.ClearDeletedAt()
	return iuo
}

// SetDeletedBy sets the "deleted_by" field.
func (iuo *IntegrationUpdateOne) SetDeletedBy(s string) *IntegrationUpdateOne {
	iuo.mutation.SetDeletedBy(s)
	return iuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableDeletedBy(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetDeletedBy(*s)
	}
	return iuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (iuo *IntegrationUpdateOne) ClearDeletedBy() *IntegrationUpdateOne {
	iuo.mutation.ClearDeletedBy()
	return iuo
}

// SetOwnerID sets the "owner_id" field.
func (iuo *IntegrationUpdateOne) SetOwnerID(s string) *IntegrationUpdateOne {
	iuo.mutation.SetOwnerID(s)
	return iuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableOwnerID(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetOwnerID(*s)
	}
	return iuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (iuo *IntegrationUpdateOne) ClearOwnerID() *IntegrationUpdateOne {
	iuo.mutation.ClearOwnerID()
	return iuo
}

// SetName sets the "name" field.
func (iuo *IntegrationUpdateOne) SetName(s string) *IntegrationUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableName(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *IntegrationUpdateOne) SetDescription(s string) *IntegrationUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableDescription(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// ClearDescription clears the value of the "description" field.
func (iuo *IntegrationUpdateOne) ClearDescription() *IntegrationUpdateOne {
	iuo.mutation.ClearDescription()
	return iuo
}

// SetKind sets the "kind" field.
func (iuo *IntegrationUpdateOne) SetKind(s string) *IntegrationUpdateOne {
	iuo.mutation.SetKind(s)
	return iuo
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableKind(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetKind(*s)
	}
	return iuo
}

// ClearKind clears the value of the "kind" field.
func (iuo *IntegrationUpdateOne) ClearKind() *IntegrationUpdateOne {
	iuo.mutation.ClearKind()
	return iuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (iuo *IntegrationUpdateOne) SetOwner(o *Organization) *IntegrationUpdateOne {
	return iuo.SetOwnerID(o.ID)
}

// AddSecretIDs adds the "secrets" edge to the Hush entity by IDs.
func (iuo *IntegrationUpdateOne) AddSecretIDs(ids ...string) *IntegrationUpdateOne {
	iuo.mutation.AddSecretIDs(ids...)
	return iuo
}

// AddSecrets adds the "secrets" edges to the Hush entity.
func (iuo *IntegrationUpdateOne) AddSecrets(h ...*Hush) *IntegrationUpdateOne {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return iuo.AddSecretIDs(ids...)
}

// AddOauth2tokenIDs adds the "oauth2tokens" edge to the OhAuthTooToken entity by IDs.
func (iuo *IntegrationUpdateOne) AddOauth2tokenIDs(ids ...string) *IntegrationUpdateOne {
	iuo.mutation.AddOauth2tokenIDs(ids...)
	return iuo
}

// AddOauth2tokens adds the "oauth2tokens" edges to the OhAuthTooToken entity.
func (iuo *IntegrationUpdateOne) AddOauth2tokens(o ...*OhAuthTooToken) *IntegrationUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return iuo.AddOauth2tokenIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (iuo *IntegrationUpdateOne) AddEventIDs(ids ...string) *IntegrationUpdateOne {
	iuo.mutation.AddEventIDs(ids...)
	return iuo
}

// AddEvents adds the "events" edges to the Event entity.
func (iuo *IntegrationUpdateOne) AddEvents(e ...*Event) *IntegrationUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iuo.AddEventIDs(ids...)
}

// Mutation returns the IntegrationMutation object of the builder.
func (iuo *IntegrationUpdateOne) Mutation() *IntegrationMutation {
	return iuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (iuo *IntegrationUpdateOne) ClearOwner() *IntegrationUpdateOne {
	iuo.mutation.ClearOwner()
	return iuo
}

// ClearSecrets clears all "secrets" edges to the Hush entity.
func (iuo *IntegrationUpdateOne) ClearSecrets() *IntegrationUpdateOne {
	iuo.mutation.ClearSecrets()
	return iuo
}

// RemoveSecretIDs removes the "secrets" edge to Hush entities by IDs.
func (iuo *IntegrationUpdateOne) RemoveSecretIDs(ids ...string) *IntegrationUpdateOne {
	iuo.mutation.RemoveSecretIDs(ids...)
	return iuo
}

// RemoveSecrets removes "secrets" edges to Hush entities.
func (iuo *IntegrationUpdateOne) RemoveSecrets(h ...*Hush) *IntegrationUpdateOne {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return iuo.RemoveSecretIDs(ids...)
}

// ClearOauth2tokens clears all "oauth2tokens" edges to the OhAuthTooToken entity.
func (iuo *IntegrationUpdateOne) ClearOauth2tokens() *IntegrationUpdateOne {
	iuo.mutation.ClearOauth2tokens()
	return iuo
}

// RemoveOauth2tokenIDs removes the "oauth2tokens" edge to OhAuthTooToken entities by IDs.
func (iuo *IntegrationUpdateOne) RemoveOauth2tokenIDs(ids ...string) *IntegrationUpdateOne {
	iuo.mutation.RemoveOauth2tokenIDs(ids...)
	return iuo
}

// RemoveOauth2tokens removes "oauth2tokens" edges to OhAuthTooToken entities.
func (iuo *IntegrationUpdateOne) RemoveOauth2tokens(o ...*OhAuthTooToken) *IntegrationUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return iuo.RemoveOauth2tokenIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (iuo *IntegrationUpdateOne) ClearEvents() *IntegrationUpdateOne {
	iuo.mutation.ClearEvents()
	return iuo
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (iuo *IntegrationUpdateOne) RemoveEventIDs(ids ...string) *IntegrationUpdateOne {
	iuo.mutation.RemoveEventIDs(ids...)
	return iuo
}

// RemoveEvents removes "events" edges to Event entities.
func (iuo *IntegrationUpdateOne) RemoveEvents(e ...*Event) *IntegrationUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iuo.RemoveEventIDs(ids...)
}

// Where appends a list predicates to the IntegrationUpdate builder.
func (iuo *IntegrationUpdateOne) Where(ps ...predicate.Integration) *IntegrationUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IntegrationUpdateOne) Select(field string, fields ...string) *IntegrationUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Integration entity.
func (iuo *IntegrationUpdateOne) Save(ctx context.Context) (*Integration, error) {
	if err := iuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IntegrationUpdateOne) SaveX(ctx context.Context) *Integration {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IntegrationUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IntegrationUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *IntegrationUpdateOne) defaults() error {
	if _, ok := iuo.mutation.UpdatedAt(); !ok && !iuo.mutation.UpdatedAtCleared() {
		if integration.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized integration.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := integration.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IntegrationUpdateOne) check() error {
	if v, ok := iuo.mutation.OwnerID(); ok {
		if err := integration.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Integration.owner_id": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Name(); ok {
		if err := integration.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Integration.name": %w`, err)}
		}
	}
	return nil
}

func (iuo *IntegrationUpdateOne) sqlSave(ctx context.Context) (_node *Integration, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(integration.Table, integration.Columns, sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Integration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, integration.FieldID)
		for _, f := range fields {
			if !integration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != integration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if iuo.mutation.CreatedAtCleared() {
		_spec.ClearField(integration.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(integration.FieldUpdatedAt, field.TypeTime, value)
	}
	if iuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(integration.FieldUpdatedAt, field.TypeTime)
	}
	if iuo.mutation.CreatedByCleared() {
		_spec.ClearField(integration.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.UpdatedBy(); ok {
		_spec.SetField(integration.FieldUpdatedBy, field.TypeString, value)
	}
	if iuo.mutation.UpdatedByCleared() {
		_spec.ClearField(integration.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.DeletedAt(); ok {
		_spec.SetField(integration.FieldDeletedAt, field.TypeTime, value)
	}
	if iuo.mutation.DeletedAtCleared() {
		_spec.ClearField(integration.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.DeletedBy(); ok {
		_spec.SetField(integration.FieldDeletedBy, field.TypeString, value)
	}
	if iuo.mutation.DeletedByCleared() {
		_spec.ClearField(integration.FieldDeletedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(integration.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.SetField(integration.FieldDescription, field.TypeString, value)
	}
	if iuo.mutation.DescriptionCleared() {
		_spec.ClearField(integration.FieldDescription, field.TypeString)
	}
	if value, ok := iuo.mutation.Kind(); ok {
		_spec.SetField(integration.FieldKind, field.TypeString, value)
	}
	if iuo.mutation.KindCleared() {
		_spec.ClearField(integration.FieldKind, field.TypeString)
	}
	if iuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.Integration
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.Integration
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.SecretsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.SecretsTable,
			Columns: integration.SecretsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationSecrets
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedSecretsIDs(); len(nodes) > 0 && !iuo.mutation.SecretsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.SecretsTable,
			Columns: integration.SecretsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.SecretsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.SecretsTable,
			Columns: integration.SecretsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.Oauth2tokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.Oauth2tokensTable,
			Columns: integration.Oauth2tokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationOauth2tokens
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedOauth2tokensIDs(); len(nodes) > 0 && !iuo.mutation.Oauth2tokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.Oauth2tokensTable,
			Columns: integration.Oauth2tokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationOauth2tokens
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.Oauth2tokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.Oauth2tokensTable,
			Columns: integration.Oauth2tokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationOauth2tokens
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.EventsTable,
			Columns: integration.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationEvents
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !iuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.EventsTable,
			Columns: integration.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.EventsTable,
			Columns: integration.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.IntegrationEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = iuo.schemaConfig.Integration
	ctx = internal.NewSchemaConfigContext(ctx, iuo.schemaConfig)
	_node = &Integration{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
