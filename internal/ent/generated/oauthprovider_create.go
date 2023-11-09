// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
)

// OauthProviderCreate is the builder for creating a OauthProvider entity.
type OauthProviderCreate struct {
	config
	mutation *OauthProviderMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (opc *OauthProviderCreate) SetCreatedAt(t time.Time) *OauthProviderCreate {
	opc.mutation.SetCreatedAt(t)
	return opc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (opc *OauthProviderCreate) SetNillableCreatedAt(t *time.Time) *OauthProviderCreate {
	if t != nil {
		opc.SetCreatedAt(*t)
	}
	return opc
}

// SetUpdatedAt sets the "updated_at" field.
func (opc *OauthProviderCreate) SetUpdatedAt(t time.Time) *OauthProviderCreate {
	opc.mutation.SetUpdatedAt(t)
	return opc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (opc *OauthProviderCreate) SetNillableUpdatedAt(t *time.Time) *OauthProviderCreate {
	if t != nil {
		opc.SetUpdatedAt(*t)
	}
	return opc
}

// SetCreatedBy sets the "created_by" field.
func (opc *OauthProviderCreate) SetCreatedBy(s string) *OauthProviderCreate {
	opc.mutation.SetCreatedBy(s)
	return opc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (opc *OauthProviderCreate) SetNillableCreatedBy(s *string) *OauthProviderCreate {
	if s != nil {
		opc.SetCreatedBy(*s)
	}
	return opc
}

// SetUpdatedBy sets the "updated_by" field.
func (opc *OauthProviderCreate) SetUpdatedBy(s string) *OauthProviderCreate {
	opc.mutation.SetUpdatedBy(s)
	return opc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (opc *OauthProviderCreate) SetNillableUpdatedBy(s *string) *OauthProviderCreate {
	if s != nil {
		opc.SetUpdatedBy(*s)
	}
	return opc
}

// SetName sets the "name" field.
func (opc *OauthProviderCreate) SetName(s string) *OauthProviderCreate {
	opc.mutation.SetName(s)
	return opc
}

// SetClientID sets the "client_id" field.
func (opc *OauthProviderCreate) SetClientID(s string) *OauthProviderCreate {
	opc.mutation.SetClientID(s)
	return opc
}

// SetClientSecret sets the "client_secret" field.
func (opc *OauthProviderCreate) SetClientSecret(s string) *OauthProviderCreate {
	opc.mutation.SetClientSecret(s)
	return opc
}

// SetRedirectURL sets the "redirect_url" field.
func (opc *OauthProviderCreate) SetRedirectURL(s string) *OauthProviderCreate {
	opc.mutation.SetRedirectURL(s)
	return opc
}

// SetScopes sets the "scopes" field.
func (opc *OauthProviderCreate) SetScopes(s string) *OauthProviderCreate {
	opc.mutation.SetScopes(s)
	return opc
}

// SetAuthURL sets the "auth_url" field.
func (opc *OauthProviderCreate) SetAuthURL(s string) *OauthProviderCreate {
	opc.mutation.SetAuthURL(s)
	return opc
}

// SetTokenURL sets the "token_url" field.
func (opc *OauthProviderCreate) SetTokenURL(s string) *OauthProviderCreate {
	opc.mutation.SetTokenURL(s)
	return opc
}

// SetAuthStyle sets the "auth_style" field.
func (opc *OauthProviderCreate) SetAuthStyle(u uint8) *OauthProviderCreate {
	opc.mutation.SetAuthStyle(u)
	return opc
}

// SetInfoURL sets the "info_url" field.
func (opc *OauthProviderCreate) SetInfoURL(s string) *OauthProviderCreate {
	opc.mutation.SetInfoURL(s)
	return opc
}

// SetID sets the "id" field.
func (opc *OauthProviderCreate) SetID(s string) *OauthProviderCreate {
	opc.mutation.SetID(s)
	return opc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (opc *OauthProviderCreate) SetNillableID(s *string) *OauthProviderCreate {
	if s != nil {
		opc.SetID(*s)
	}
	return opc
}

// Mutation returns the OauthProviderMutation object of the builder.
func (opc *OauthProviderCreate) Mutation() *OauthProviderMutation {
	return opc.mutation
}

// Save creates the OauthProvider in the database.
func (opc *OauthProviderCreate) Save(ctx context.Context) (*OauthProvider, error) {
	if err := opc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, opc.sqlSave, opc.mutation, opc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (opc *OauthProviderCreate) SaveX(ctx context.Context) *OauthProvider {
	v, err := opc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opc *OauthProviderCreate) Exec(ctx context.Context) error {
	_, err := opc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opc *OauthProviderCreate) ExecX(ctx context.Context) {
	if err := opc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opc *OauthProviderCreate) defaults() error {
	if _, ok := opc.mutation.CreatedAt(); !ok {
		if oauthprovider.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized oauthprovider.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := oauthprovider.DefaultCreatedAt()
		opc.mutation.SetCreatedAt(v)
	}
	if _, ok := opc.mutation.UpdatedAt(); !ok {
		if oauthprovider.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized oauthprovider.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := oauthprovider.DefaultUpdatedAt()
		opc.mutation.SetUpdatedAt(v)
	}
	if _, ok := opc.mutation.ID(); !ok {
		if oauthprovider.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized oauthprovider.DefaultID (forgotten import generated/runtime?)")
		}
		v := oauthprovider.DefaultID()
		opc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (opc *OauthProviderCreate) check() error {
	if _, ok := opc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "OauthProvider.created_at"`)}
	}
	if _, ok := opc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "OauthProvider.updated_at"`)}
	}
	if _, ok := opc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "OauthProvider.name"`)}
	}
	if _, ok := opc.mutation.ClientID(); !ok {
		return &ValidationError{Name: "client_id", err: errors.New(`generated: missing required field "OauthProvider.client_id"`)}
	}
	if _, ok := opc.mutation.ClientSecret(); !ok {
		return &ValidationError{Name: "client_secret", err: errors.New(`generated: missing required field "OauthProvider.client_secret"`)}
	}
	if _, ok := opc.mutation.RedirectURL(); !ok {
		return &ValidationError{Name: "redirect_url", err: errors.New(`generated: missing required field "OauthProvider.redirect_url"`)}
	}
	if _, ok := opc.mutation.Scopes(); !ok {
		return &ValidationError{Name: "scopes", err: errors.New(`generated: missing required field "OauthProvider.scopes"`)}
	}
	if _, ok := opc.mutation.AuthURL(); !ok {
		return &ValidationError{Name: "auth_url", err: errors.New(`generated: missing required field "OauthProvider.auth_url"`)}
	}
	if _, ok := opc.mutation.TokenURL(); !ok {
		return &ValidationError{Name: "token_url", err: errors.New(`generated: missing required field "OauthProvider.token_url"`)}
	}
	if _, ok := opc.mutation.AuthStyle(); !ok {
		return &ValidationError{Name: "auth_style", err: errors.New(`generated: missing required field "OauthProvider.auth_style"`)}
	}
	if _, ok := opc.mutation.InfoURL(); !ok {
		return &ValidationError{Name: "info_url", err: errors.New(`generated: missing required field "OauthProvider.info_url"`)}
	}
	return nil
}

func (opc *OauthProviderCreate) sqlSave(ctx context.Context) (*OauthProvider, error) {
	if err := opc.check(); err != nil {
		return nil, err
	}
	_node, _spec := opc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OauthProvider.ID type: %T", _spec.ID.Value)
		}
	}
	opc.mutation.id = &_node.ID
	opc.mutation.done = true
	return _node, nil
}

func (opc *OauthProviderCreate) createSpec() (*OauthProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &OauthProvider{config: opc.config}
		_spec = sqlgraph.NewCreateSpec(oauthprovider.Table, sqlgraph.NewFieldSpec(oauthprovider.FieldID, field.TypeString))
	)
	_spec.Schema = opc.schemaConfig.OauthProvider
	if id, ok := opc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := opc.mutation.CreatedAt(); ok {
		_spec.SetField(oauthprovider.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := opc.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := opc.mutation.CreatedBy(); ok {
		_spec.SetField(oauthprovider.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := opc.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := opc.mutation.Name(); ok {
		_spec.SetField(oauthprovider.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := opc.mutation.ClientID(); ok {
		_spec.SetField(oauthprovider.FieldClientID, field.TypeString, value)
		_node.ClientID = value
	}
	if value, ok := opc.mutation.ClientSecret(); ok {
		_spec.SetField(oauthprovider.FieldClientSecret, field.TypeString, value)
		_node.ClientSecret = value
	}
	if value, ok := opc.mutation.RedirectURL(); ok {
		_spec.SetField(oauthprovider.FieldRedirectURL, field.TypeString, value)
		_node.RedirectURL = value
	}
	if value, ok := opc.mutation.Scopes(); ok {
		_spec.SetField(oauthprovider.FieldScopes, field.TypeString, value)
		_node.Scopes = value
	}
	if value, ok := opc.mutation.AuthURL(); ok {
		_spec.SetField(oauthprovider.FieldAuthURL, field.TypeString, value)
		_node.AuthURL = value
	}
	if value, ok := opc.mutation.TokenURL(); ok {
		_spec.SetField(oauthprovider.FieldTokenURL, field.TypeString, value)
		_node.TokenURL = value
	}
	if value, ok := opc.mutation.AuthStyle(); ok {
		_spec.SetField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
		_node.AuthStyle = value
	}
	if value, ok := opc.mutation.InfoURL(); ok {
		_spec.SetField(oauthprovider.FieldInfoURL, field.TypeString, value)
		_node.InfoURL = value
	}
	return _node, _spec
}

// OauthProviderCreateBulk is the builder for creating many OauthProvider entities in bulk.
type OauthProviderCreateBulk struct {
	config
	err      error
	builders []*OauthProviderCreate
}

// Save creates the OauthProvider entities in the database.
func (opcb *OauthProviderCreateBulk) Save(ctx context.Context) ([]*OauthProvider, error) {
	if opcb.err != nil {
		return nil, opcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(opcb.builders))
	nodes := make([]*OauthProvider, len(opcb.builders))
	mutators := make([]Mutator, len(opcb.builders))
	for i := range opcb.builders {
		func(i int, root context.Context) {
			builder := opcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OauthProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, opcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, opcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opcb *OauthProviderCreateBulk) SaveX(ctx context.Context) []*OauthProvider {
	v, err := opcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opcb *OauthProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := opcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opcb *OauthProviderCreateBulk) ExecX(ctx context.Context) {
	if err := opcb.Exec(ctx); err != nil {
		panic(err)
	}
}
