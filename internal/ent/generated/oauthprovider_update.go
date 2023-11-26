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
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OauthProviderUpdate is the builder for updating OauthProvider entities.
type OauthProviderUpdate struct {
	config
	hooks    []Hook
	mutation *OauthProviderMutation
}

// Where appends a list predicates to the OauthProviderUpdate builder.
func (opu *OauthProviderUpdate) Where(ps ...predicate.OauthProvider) *OauthProviderUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetUpdatedAt sets the "updated_at" field.
func (opu *OauthProviderUpdate) SetUpdatedAt(t time.Time) *OauthProviderUpdate {
	opu.mutation.SetUpdatedAt(t)
	return opu
}

// SetUpdatedBy sets the "updated_by" field.
func (opu *OauthProviderUpdate) SetUpdatedBy(s string) *OauthProviderUpdate {
	opu.mutation.SetUpdatedBy(s)
	return opu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableUpdatedBy(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetUpdatedBy(*s)
	}
	return opu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (opu *OauthProviderUpdate) ClearUpdatedBy() *OauthProviderUpdate {
	opu.mutation.ClearUpdatedBy()
	return opu
}

// SetName sets the "name" field.
func (opu *OauthProviderUpdate) SetName(s string) *OauthProviderUpdate {
	opu.mutation.SetName(s)
	return opu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableName(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetName(*s)
	}
	return opu
}

// SetClientID sets the "client_id" field.
func (opu *OauthProviderUpdate) SetClientID(s string) *OauthProviderUpdate {
	opu.mutation.SetClientID(s)
	return opu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableClientID(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetClientID(*s)
	}
	return opu
}

// SetClientSecret sets the "client_secret" field.
func (opu *OauthProviderUpdate) SetClientSecret(s string) *OauthProviderUpdate {
	opu.mutation.SetClientSecret(s)
	return opu
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableClientSecret(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetClientSecret(*s)
	}
	return opu
}

// SetRedirectURL sets the "redirect_url" field.
func (opu *OauthProviderUpdate) SetRedirectURL(s string) *OauthProviderUpdate {
	opu.mutation.SetRedirectURL(s)
	return opu
}

// SetNillableRedirectURL sets the "redirect_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableRedirectURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetRedirectURL(*s)
	}
	return opu
}

// SetScopes sets the "scopes" field.
func (opu *OauthProviderUpdate) SetScopes(s string) *OauthProviderUpdate {
	opu.mutation.SetScopes(s)
	return opu
}

// SetNillableScopes sets the "scopes" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableScopes(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetScopes(*s)
	}
	return opu
}

// SetAuthURL sets the "auth_url" field.
func (opu *OauthProviderUpdate) SetAuthURL(s string) *OauthProviderUpdate {
	opu.mutation.SetAuthURL(s)
	return opu
}

// SetNillableAuthURL sets the "auth_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableAuthURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetAuthURL(*s)
	}
	return opu
}

// SetTokenURL sets the "token_url" field.
func (opu *OauthProviderUpdate) SetTokenURL(s string) *OauthProviderUpdate {
	opu.mutation.SetTokenURL(s)
	return opu
}

// SetNillableTokenURL sets the "token_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableTokenURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetTokenURL(*s)
	}
	return opu
}

// SetAuthStyle sets the "auth_style" field.
func (opu *OauthProviderUpdate) SetAuthStyle(u uint8) *OauthProviderUpdate {
	opu.mutation.ResetAuthStyle()
	opu.mutation.SetAuthStyle(u)
	return opu
}

// SetNillableAuthStyle sets the "auth_style" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableAuthStyle(u *uint8) *OauthProviderUpdate {
	if u != nil {
		opu.SetAuthStyle(*u)
	}
	return opu
}

// AddAuthStyle adds u to the "auth_style" field.
func (opu *OauthProviderUpdate) AddAuthStyle(u int8) *OauthProviderUpdate {
	opu.mutation.AddAuthStyle(u)
	return opu
}

// SetInfoURL sets the "info_url" field.
func (opu *OauthProviderUpdate) SetInfoURL(s string) *OauthProviderUpdate {
	opu.mutation.SetInfoURL(s)
	return opu
}

// SetNillableInfoURL sets the "info_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableInfoURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetInfoURL(*s)
	}
	return opu
}

// SetOwnerID sets the "owner" edge to the Organization entity by ID.
func (opu *OauthProviderUpdate) SetOwnerID(id string) *OauthProviderUpdate {
	opu.mutation.SetOwnerID(id)
	return opu
}

// SetNillableOwnerID sets the "owner" edge to the Organization entity by ID if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableOwnerID(id *string) *OauthProviderUpdate {
	if id != nil {
		opu = opu.SetOwnerID(*id)
	}
	return opu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (opu *OauthProviderUpdate) SetOwner(o *Organization) *OauthProviderUpdate {
	return opu.SetOwnerID(o.ID)
}

// Mutation returns the OauthProviderMutation object of the builder.
func (opu *OauthProviderUpdate) Mutation() *OauthProviderMutation {
	return opu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (opu *OauthProviderUpdate) ClearOwner() *OauthProviderUpdate {
	opu.mutation.ClearOwner()
	return opu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OauthProviderUpdate) Save(ctx context.Context) (int, error) {
	if err := opu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, opu.sqlSave, opu.mutation, opu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OauthProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OauthProviderUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OauthProviderUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opu *OauthProviderUpdate) defaults() error {
	if _, ok := opu.mutation.UpdatedAt(); !ok {
		if oauthprovider.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized oauthprovider.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := oauthprovider.UpdateDefaultUpdatedAt()
		opu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (opu *OauthProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(oauthprovider.Table, oauthprovider.Columns, sqlgraph.NewFieldSpec(oauthprovider.FieldID, field.TypeString))
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opu.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if opu.mutation.CreatedByCleared() {
		_spec.ClearField(oauthprovider.FieldCreatedBy, field.TypeString)
	}
	if value, ok := opu.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedBy, field.TypeString, value)
	}
	if opu.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthprovider.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := opu.mutation.Name(); ok {
		_spec.SetField(oauthprovider.FieldName, field.TypeString, value)
	}
	if value, ok := opu.mutation.ClientID(); ok {
		_spec.SetField(oauthprovider.FieldClientID, field.TypeString, value)
	}
	if value, ok := opu.mutation.ClientSecret(); ok {
		_spec.SetField(oauthprovider.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := opu.mutation.RedirectURL(); ok {
		_spec.SetField(oauthprovider.FieldRedirectURL, field.TypeString, value)
	}
	if value, ok := opu.mutation.Scopes(); ok {
		_spec.SetField(oauthprovider.FieldScopes, field.TypeString, value)
	}
	if value, ok := opu.mutation.AuthURL(); ok {
		_spec.SetField(oauthprovider.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := opu.mutation.TokenURL(); ok {
		_spec.SetField(oauthprovider.FieldTokenURL, field.TypeString, value)
	}
	if value, ok := opu.mutation.AuthStyle(); ok {
		_spec.SetField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opu.mutation.AddedAuthStyle(); ok {
		_spec.AddField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opu.mutation.InfoURL(); ok {
		_spec.SetField(oauthprovider.FieldInfoURL, field.TypeString, value)
	}
	if opu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opu.schemaConfig.OauthProvider
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opu.schemaConfig.OauthProvider
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = opu.schemaConfig.OauthProvider
	ctx = internal.NewSchemaConfigContext(ctx, opu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	opu.mutation.done = true
	return n, nil
}

// OauthProviderUpdateOne is the builder for updating a single OauthProvider entity.
type OauthProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OauthProviderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (opuo *OauthProviderUpdateOne) SetUpdatedAt(t time.Time) *OauthProviderUpdateOne {
	opuo.mutation.SetUpdatedAt(t)
	return opuo
}

// SetUpdatedBy sets the "updated_by" field.
func (opuo *OauthProviderUpdateOne) SetUpdatedBy(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetUpdatedBy(s)
	return opuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableUpdatedBy(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetUpdatedBy(*s)
	}
	return opuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (opuo *OauthProviderUpdateOne) ClearUpdatedBy() *OauthProviderUpdateOne {
	opuo.mutation.ClearUpdatedBy()
	return opuo
}

// SetName sets the "name" field.
func (opuo *OauthProviderUpdateOne) SetName(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetName(s)
	return opuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableName(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetName(*s)
	}
	return opuo
}

// SetClientID sets the "client_id" field.
func (opuo *OauthProviderUpdateOne) SetClientID(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetClientID(s)
	return opuo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableClientID(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetClientID(*s)
	}
	return opuo
}

// SetClientSecret sets the "client_secret" field.
func (opuo *OauthProviderUpdateOne) SetClientSecret(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetClientSecret(s)
	return opuo
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableClientSecret(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetClientSecret(*s)
	}
	return opuo
}

// SetRedirectURL sets the "redirect_url" field.
func (opuo *OauthProviderUpdateOne) SetRedirectURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetRedirectURL(s)
	return opuo
}

// SetNillableRedirectURL sets the "redirect_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableRedirectURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetRedirectURL(*s)
	}
	return opuo
}

// SetScopes sets the "scopes" field.
func (opuo *OauthProviderUpdateOne) SetScopes(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetScopes(s)
	return opuo
}

// SetNillableScopes sets the "scopes" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableScopes(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetScopes(*s)
	}
	return opuo
}

// SetAuthURL sets the "auth_url" field.
func (opuo *OauthProviderUpdateOne) SetAuthURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetAuthURL(s)
	return opuo
}

// SetNillableAuthURL sets the "auth_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableAuthURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetAuthURL(*s)
	}
	return opuo
}

// SetTokenURL sets the "token_url" field.
func (opuo *OauthProviderUpdateOne) SetTokenURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetTokenURL(s)
	return opuo
}

// SetNillableTokenURL sets the "token_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableTokenURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetTokenURL(*s)
	}
	return opuo
}

// SetAuthStyle sets the "auth_style" field.
func (opuo *OauthProviderUpdateOne) SetAuthStyle(u uint8) *OauthProviderUpdateOne {
	opuo.mutation.ResetAuthStyle()
	opuo.mutation.SetAuthStyle(u)
	return opuo
}

// SetNillableAuthStyle sets the "auth_style" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableAuthStyle(u *uint8) *OauthProviderUpdateOne {
	if u != nil {
		opuo.SetAuthStyle(*u)
	}
	return opuo
}

// AddAuthStyle adds u to the "auth_style" field.
func (opuo *OauthProviderUpdateOne) AddAuthStyle(u int8) *OauthProviderUpdateOne {
	opuo.mutation.AddAuthStyle(u)
	return opuo
}

// SetInfoURL sets the "info_url" field.
func (opuo *OauthProviderUpdateOne) SetInfoURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetInfoURL(s)
	return opuo
}

// SetNillableInfoURL sets the "info_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableInfoURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetInfoURL(*s)
	}
	return opuo
}

// SetOwnerID sets the "owner" edge to the Organization entity by ID.
func (opuo *OauthProviderUpdateOne) SetOwnerID(id string) *OauthProviderUpdateOne {
	opuo.mutation.SetOwnerID(id)
	return opuo
}

// SetNillableOwnerID sets the "owner" edge to the Organization entity by ID if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableOwnerID(id *string) *OauthProviderUpdateOne {
	if id != nil {
		opuo = opuo.SetOwnerID(*id)
	}
	return opuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (opuo *OauthProviderUpdateOne) SetOwner(o *Organization) *OauthProviderUpdateOne {
	return opuo.SetOwnerID(o.ID)
}

// Mutation returns the OauthProviderMutation object of the builder.
func (opuo *OauthProviderUpdateOne) Mutation() *OauthProviderMutation {
	return opuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (opuo *OauthProviderUpdateOne) ClearOwner() *OauthProviderUpdateOne {
	opuo.mutation.ClearOwner()
	return opuo
}

// Where appends a list predicates to the OauthProviderUpdate builder.
func (opuo *OauthProviderUpdateOne) Where(ps ...predicate.OauthProvider) *OauthProviderUpdateOne {
	opuo.mutation.Where(ps...)
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OauthProviderUpdateOne) Select(field string, fields ...string) *OauthProviderUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OauthProvider entity.
func (opuo *OauthProviderUpdateOne) Save(ctx context.Context) (*OauthProvider, error) {
	if err := opuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, opuo.sqlSave, opuo.mutation, opuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opuo *OauthProviderUpdateOne) SaveX(ctx context.Context) *OauthProvider {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OauthProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OauthProviderUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opuo *OauthProviderUpdateOne) defaults() error {
	if _, ok := opuo.mutation.UpdatedAt(); !ok {
		if oauthprovider.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized oauthprovider.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := oauthprovider.UpdateDefaultUpdatedAt()
		opuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (opuo *OauthProviderUpdateOne) sqlSave(ctx context.Context) (_node *OauthProvider, err error) {
	_spec := sqlgraph.NewUpdateSpec(oauthprovider.Table, oauthprovider.Columns, sqlgraph.NewFieldSpec(oauthprovider.FieldID, field.TypeString))
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OauthProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthprovider.FieldID)
		for _, f := range fields {
			if !oauthprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != oauthprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opuo.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if opuo.mutation.CreatedByCleared() {
		_spec.ClearField(oauthprovider.FieldCreatedBy, field.TypeString)
	}
	if value, ok := opuo.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedBy, field.TypeString, value)
	}
	if opuo.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthprovider.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := opuo.mutation.Name(); ok {
		_spec.SetField(oauthprovider.FieldName, field.TypeString, value)
	}
	if value, ok := opuo.mutation.ClientID(); ok {
		_spec.SetField(oauthprovider.FieldClientID, field.TypeString, value)
	}
	if value, ok := opuo.mutation.ClientSecret(); ok {
		_spec.SetField(oauthprovider.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := opuo.mutation.RedirectURL(); ok {
		_spec.SetField(oauthprovider.FieldRedirectURL, field.TypeString, value)
	}
	if value, ok := opuo.mutation.Scopes(); ok {
		_spec.SetField(oauthprovider.FieldScopes, field.TypeString, value)
	}
	if value, ok := opuo.mutation.AuthURL(); ok {
		_spec.SetField(oauthprovider.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := opuo.mutation.TokenURL(); ok {
		_spec.SetField(oauthprovider.FieldTokenURL, field.TypeString, value)
	}
	if value, ok := opuo.mutation.AuthStyle(); ok {
		_spec.SetField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opuo.mutation.AddedAuthStyle(); ok {
		_spec.AddField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opuo.mutation.InfoURL(); ok {
		_spec.SetField(oauthprovider.FieldInfoURL, field.TypeString, value)
	}
	if opuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opuo.schemaConfig.OauthProvider
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opuo.schemaConfig.OauthProvider
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = opuo.schemaConfig.OauthProvider
	ctx = internal.NewSchemaConfigContext(ctx, opuo.schemaConfig)
	_node = &OauthProvider{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	opuo.mutation.done = true
	return _node, nil
}
