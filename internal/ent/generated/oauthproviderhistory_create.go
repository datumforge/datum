// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/generated/oauthproviderhistory"
	"github.com/datumforge/enthistory"
)

// OauthProviderHistoryCreate is the builder for creating a OauthProviderHistory entity.
type OauthProviderHistoryCreate struct {
	config
	mutation *OauthProviderHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (ophc *OauthProviderHistoryCreate) SetHistoryTime(t time.Time) *OauthProviderHistoryCreate {
	ophc.mutation.SetHistoryTime(t)
	return ophc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableHistoryTime(t *time.Time) *OauthProviderHistoryCreate {
	if t != nil {
		ophc.SetHistoryTime(*t)
	}
	return ophc
}

// SetOperation sets the "operation" field.
func (ophc *OauthProviderHistoryCreate) SetOperation(et enthistory.OpType) *OauthProviderHistoryCreate {
	ophc.mutation.SetOperation(et)
	return ophc
}

// SetRef sets the "ref" field.
func (ophc *OauthProviderHistoryCreate) SetRef(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetRef(s)
	return ophc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableRef(s *string) *OauthProviderHistoryCreate {
	if s != nil {
		ophc.SetRef(*s)
	}
	return ophc
}

// SetCreatedAt sets the "created_at" field.
func (ophc *OauthProviderHistoryCreate) SetCreatedAt(t time.Time) *OauthProviderHistoryCreate {
	ophc.mutation.SetCreatedAt(t)
	return ophc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableCreatedAt(t *time.Time) *OauthProviderHistoryCreate {
	if t != nil {
		ophc.SetCreatedAt(*t)
	}
	return ophc
}

// SetUpdatedAt sets the "updated_at" field.
func (ophc *OauthProviderHistoryCreate) SetUpdatedAt(t time.Time) *OauthProviderHistoryCreate {
	ophc.mutation.SetUpdatedAt(t)
	return ophc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableUpdatedAt(t *time.Time) *OauthProviderHistoryCreate {
	if t != nil {
		ophc.SetUpdatedAt(*t)
	}
	return ophc
}

// SetCreatedBy sets the "created_by" field.
func (ophc *OauthProviderHistoryCreate) SetCreatedBy(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetCreatedBy(s)
	return ophc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableCreatedBy(s *string) *OauthProviderHistoryCreate {
	if s != nil {
		ophc.SetCreatedBy(*s)
	}
	return ophc
}

// SetUpdatedBy sets the "updated_by" field.
func (ophc *OauthProviderHistoryCreate) SetUpdatedBy(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetUpdatedBy(s)
	return ophc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableUpdatedBy(s *string) *OauthProviderHistoryCreate {
	if s != nil {
		ophc.SetUpdatedBy(*s)
	}
	return ophc
}

// SetMappingID sets the "mapping_id" field.
func (ophc *OauthProviderHistoryCreate) SetMappingID(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetMappingID(s)
	return ophc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableMappingID(s *string) *OauthProviderHistoryCreate {
	if s != nil {
		ophc.SetMappingID(*s)
	}
	return ophc
}

// SetDeletedAt sets the "deleted_at" field.
func (ophc *OauthProviderHistoryCreate) SetDeletedAt(t time.Time) *OauthProviderHistoryCreate {
	ophc.mutation.SetDeletedAt(t)
	return ophc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableDeletedAt(t *time.Time) *OauthProviderHistoryCreate {
	if t != nil {
		ophc.SetDeletedAt(*t)
	}
	return ophc
}

// SetDeletedBy sets the "deleted_by" field.
func (ophc *OauthProviderHistoryCreate) SetDeletedBy(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetDeletedBy(s)
	return ophc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableDeletedBy(s *string) *OauthProviderHistoryCreate {
	if s != nil {
		ophc.SetDeletedBy(*s)
	}
	return ophc
}

// SetName sets the "name" field.
func (ophc *OauthProviderHistoryCreate) SetName(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetName(s)
	return ophc
}

// SetClientID sets the "client_id" field.
func (ophc *OauthProviderHistoryCreate) SetClientID(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetClientID(s)
	return ophc
}

// SetClientSecret sets the "client_secret" field.
func (ophc *OauthProviderHistoryCreate) SetClientSecret(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetClientSecret(s)
	return ophc
}

// SetRedirectURL sets the "redirect_url" field.
func (ophc *OauthProviderHistoryCreate) SetRedirectURL(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetRedirectURL(s)
	return ophc
}

// SetScopes sets the "scopes" field.
func (ophc *OauthProviderHistoryCreate) SetScopes(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetScopes(s)
	return ophc
}

// SetAuthURL sets the "auth_url" field.
func (ophc *OauthProviderHistoryCreate) SetAuthURL(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetAuthURL(s)
	return ophc
}

// SetTokenURL sets the "token_url" field.
func (ophc *OauthProviderHistoryCreate) SetTokenURL(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetTokenURL(s)
	return ophc
}

// SetAuthStyle sets the "auth_style" field.
func (ophc *OauthProviderHistoryCreate) SetAuthStyle(c customtypes.Uint8) *OauthProviderHistoryCreate {
	ophc.mutation.SetAuthStyle(c)
	return ophc
}

// SetInfoURL sets the "info_url" field.
func (ophc *OauthProviderHistoryCreate) SetInfoURL(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetInfoURL(s)
	return ophc
}

// SetID sets the "id" field.
func (ophc *OauthProviderHistoryCreate) SetID(s string) *OauthProviderHistoryCreate {
	ophc.mutation.SetID(s)
	return ophc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ophc *OauthProviderHistoryCreate) SetNillableID(s *string) *OauthProviderHistoryCreate {
	if s != nil {
		ophc.SetID(*s)
	}
	return ophc
}

// Mutation returns the OauthProviderHistoryMutation object of the builder.
func (ophc *OauthProviderHistoryCreate) Mutation() *OauthProviderHistoryMutation {
	return ophc.mutation
}

// Save creates the OauthProviderHistory in the database.
func (ophc *OauthProviderHistoryCreate) Save(ctx context.Context) (*OauthProviderHistory, error) {
	ophc.defaults()
	return withHooks(ctx, ophc.sqlSave, ophc.mutation, ophc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ophc *OauthProviderHistoryCreate) SaveX(ctx context.Context) *OauthProviderHistory {
	v, err := ophc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ophc *OauthProviderHistoryCreate) Exec(ctx context.Context) error {
	_, err := ophc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ophc *OauthProviderHistoryCreate) ExecX(ctx context.Context) {
	if err := ophc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ophc *OauthProviderHistoryCreate) defaults() {
	if _, ok := ophc.mutation.HistoryTime(); !ok {
		v := oauthproviderhistory.DefaultHistoryTime()
		ophc.mutation.SetHistoryTime(v)
	}
	if _, ok := ophc.mutation.CreatedAt(); !ok {
		v := oauthproviderhistory.DefaultCreatedAt()
		ophc.mutation.SetCreatedAt(v)
	}
	if _, ok := ophc.mutation.UpdatedAt(); !ok {
		v := oauthproviderhistory.DefaultUpdatedAt()
		ophc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ophc.mutation.MappingID(); !ok {
		v := oauthproviderhistory.DefaultMappingID()
		ophc.mutation.SetMappingID(v)
	}
	if _, ok := ophc.mutation.ID(); !ok {
		v := oauthproviderhistory.DefaultID()
		ophc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ophc *OauthProviderHistoryCreate) check() error {
	if _, ok := ophc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "OauthProviderHistory.history_time"`)}
	}
	if _, ok := ophc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "OauthProviderHistory.operation"`)}
	}
	if v, ok := ophc.mutation.Operation(); ok {
		if err := oauthproviderhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "OauthProviderHistory.operation": %w`, err)}
		}
	}
	if _, ok := ophc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "OauthProviderHistory.mapping_id"`)}
	}
	if _, ok := ophc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "OauthProviderHistory.name"`)}
	}
	if _, ok := ophc.mutation.ClientID(); !ok {
		return &ValidationError{Name: "client_id", err: errors.New(`generated: missing required field "OauthProviderHistory.client_id"`)}
	}
	if _, ok := ophc.mutation.ClientSecret(); !ok {
		return &ValidationError{Name: "client_secret", err: errors.New(`generated: missing required field "OauthProviderHistory.client_secret"`)}
	}
	if _, ok := ophc.mutation.RedirectURL(); !ok {
		return &ValidationError{Name: "redirect_url", err: errors.New(`generated: missing required field "OauthProviderHistory.redirect_url"`)}
	}
	if _, ok := ophc.mutation.Scopes(); !ok {
		return &ValidationError{Name: "scopes", err: errors.New(`generated: missing required field "OauthProviderHistory.scopes"`)}
	}
	if _, ok := ophc.mutation.AuthURL(); !ok {
		return &ValidationError{Name: "auth_url", err: errors.New(`generated: missing required field "OauthProviderHistory.auth_url"`)}
	}
	if _, ok := ophc.mutation.TokenURL(); !ok {
		return &ValidationError{Name: "token_url", err: errors.New(`generated: missing required field "OauthProviderHistory.token_url"`)}
	}
	if _, ok := ophc.mutation.AuthStyle(); !ok {
		return &ValidationError{Name: "auth_style", err: errors.New(`generated: missing required field "OauthProviderHistory.auth_style"`)}
	}
	if _, ok := ophc.mutation.InfoURL(); !ok {
		return &ValidationError{Name: "info_url", err: errors.New(`generated: missing required field "OauthProviderHistory.info_url"`)}
	}
	return nil
}

func (ophc *OauthProviderHistoryCreate) sqlSave(ctx context.Context) (*OauthProviderHistory, error) {
	if err := ophc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ophc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ophc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OauthProviderHistory.ID type: %T", _spec.ID.Value)
		}
	}
	ophc.mutation.id = &_node.ID
	ophc.mutation.done = true
	return _node, nil
}

func (ophc *OauthProviderHistoryCreate) createSpec() (*OauthProviderHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &OauthProviderHistory{config: ophc.config}
		_spec = sqlgraph.NewCreateSpec(oauthproviderhistory.Table, sqlgraph.NewFieldSpec(oauthproviderhistory.FieldID, field.TypeString))
	)
	_spec.Schema = ophc.schemaConfig.OauthProviderHistory
	if id, ok := ophc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ophc.mutation.HistoryTime(); ok {
		_spec.SetField(oauthproviderhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ophc.mutation.Operation(); ok {
		_spec.SetField(oauthproviderhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ophc.mutation.Ref(); ok {
		_spec.SetField(oauthproviderhistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := ophc.mutation.CreatedAt(); ok {
		_spec.SetField(oauthproviderhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ophc.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthproviderhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ophc.mutation.CreatedBy(); ok {
		_spec.SetField(oauthproviderhistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ophc.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthproviderhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ophc.mutation.MappingID(); ok {
		_spec.SetField(oauthproviderhistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := ophc.mutation.DeletedAt(); ok {
		_spec.SetField(oauthproviderhistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ophc.mutation.DeletedBy(); ok {
		_spec.SetField(oauthproviderhistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ophc.mutation.Name(); ok {
		_spec.SetField(oauthproviderhistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ophc.mutation.ClientID(); ok {
		_spec.SetField(oauthproviderhistory.FieldClientID, field.TypeString, value)
		_node.ClientID = value
	}
	if value, ok := ophc.mutation.ClientSecret(); ok {
		_spec.SetField(oauthproviderhistory.FieldClientSecret, field.TypeString, value)
		_node.ClientSecret = value
	}
	if value, ok := ophc.mutation.RedirectURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldRedirectURL, field.TypeString, value)
		_node.RedirectURL = value
	}
	if value, ok := ophc.mutation.Scopes(); ok {
		_spec.SetField(oauthproviderhistory.FieldScopes, field.TypeString, value)
		_node.Scopes = value
	}
	if value, ok := ophc.mutation.AuthURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldAuthURL, field.TypeString, value)
		_node.AuthURL = value
	}
	if value, ok := ophc.mutation.TokenURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldTokenURL, field.TypeString, value)
		_node.TokenURL = value
	}
	if value, ok := ophc.mutation.AuthStyle(); ok {
		_spec.SetField(oauthproviderhistory.FieldAuthStyle, field.TypeUint8, value)
		_node.AuthStyle = value
	}
	if value, ok := ophc.mutation.InfoURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldInfoURL, field.TypeString, value)
		_node.InfoURL = value
	}
	return _node, _spec
}

// OauthProviderHistoryCreateBulk is the builder for creating many OauthProviderHistory entities in bulk.
type OauthProviderHistoryCreateBulk struct {
	config
	err      error
	builders []*OauthProviderHistoryCreate
}

// Save creates the OauthProviderHistory entities in the database.
func (ophcb *OauthProviderHistoryCreateBulk) Save(ctx context.Context) ([]*OauthProviderHistory, error) {
	if ophcb.err != nil {
		return nil, ophcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ophcb.builders))
	nodes := make([]*OauthProviderHistory, len(ophcb.builders))
	mutators := make([]Mutator, len(ophcb.builders))
	for i := range ophcb.builders {
		func(i int, root context.Context) {
			builder := ophcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OauthProviderHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ophcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ophcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ophcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ophcb *OauthProviderHistoryCreateBulk) SaveX(ctx context.Context) []*OauthProviderHistory {
	v, err := ophcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ophcb *OauthProviderHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ophcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ophcb *OauthProviderHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ophcb.Exec(ctx); err != nil {
		panic(err)
	}
}
