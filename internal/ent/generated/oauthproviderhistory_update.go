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
	"github.com/datumforge/datum/internal/ent/generated/oauthproviderhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OauthProviderHistoryUpdate is the builder for updating OauthProviderHistory entities.
type OauthProviderHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *OauthProviderHistoryMutation
}

// Where appends a list predicates to the OauthProviderHistoryUpdate builder.
func (ophu *OauthProviderHistoryUpdate) Where(ps ...predicate.OauthProviderHistory) *OauthProviderHistoryUpdate {
	ophu.mutation.Where(ps...)
	return ophu
}

// SetUpdatedAt sets the "updated_at" field.
func (ophu *OauthProviderHistoryUpdate) SetUpdatedAt(t time.Time) *OauthProviderHistoryUpdate {
	ophu.mutation.SetUpdatedAt(t)
	return ophu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *OauthProviderHistoryUpdate {
	if t != nil {
		ophu.SetUpdatedAt(*t)
	}
	return ophu
}

// SetUpdatedBy sets the "updated_by" field.
func (ophu *OauthProviderHistoryUpdate) SetUpdatedBy(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetUpdatedBy(s)
	return ophu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableUpdatedBy(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetUpdatedBy(*s)
	}
	return ophu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ophu *OauthProviderHistoryUpdate) ClearUpdatedBy() *OauthProviderHistoryUpdate {
	ophu.mutation.ClearUpdatedBy()
	return ophu
}

// SetDeletedAt sets the "deleted_at" field.
func (ophu *OauthProviderHistoryUpdate) SetDeletedAt(t time.Time) *OauthProviderHistoryUpdate {
	ophu.mutation.SetDeletedAt(t)
	return ophu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableDeletedAt(t *time.Time) *OauthProviderHistoryUpdate {
	if t != nil {
		ophu.SetDeletedAt(*t)
	}
	return ophu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ophu *OauthProviderHistoryUpdate) ClearDeletedAt() *OauthProviderHistoryUpdate {
	ophu.mutation.ClearDeletedAt()
	return ophu
}

// SetDeletedBy sets the "deleted_by" field.
func (ophu *OauthProviderHistoryUpdate) SetDeletedBy(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetDeletedBy(s)
	return ophu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableDeletedBy(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetDeletedBy(*s)
	}
	return ophu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ophu *OauthProviderHistoryUpdate) ClearDeletedBy() *OauthProviderHistoryUpdate {
	ophu.mutation.ClearDeletedBy()
	return ophu
}

// SetName sets the "name" field.
func (ophu *OauthProviderHistoryUpdate) SetName(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetName(s)
	return ophu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableName(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetName(*s)
	}
	return ophu
}

// SetClientID sets the "client_id" field.
func (ophu *OauthProviderHistoryUpdate) SetClientID(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetClientID(s)
	return ophu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableClientID(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetClientID(*s)
	}
	return ophu
}

// SetClientSecret sets the "client_secret" field.
func (ophu *OauthProviderHistoryUpdate) SetClientSecret(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetClientSecret(s)
	return ophu
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableClientSecret(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetClientSecret(*s)
	}
	return ophu
}

// SetRedirectURL sets the "redirect_url" field.
func (ophu *OauthProviderHistoryUpdate) SetRedirectURL(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetRedirectURL(s)
	return ophu
}

// SetNillableRedirectURL sets the "redirect_url" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableRedirectURL(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetRedirectURL(*s)
	}
	return ophu
}

// SetScopes sets the "scopes" field.
func (ophu *OauthProviderHistoryUpdate) SetScopes(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetScopes(s)
	return ophu
}

// SetNillableScopes sets the "scopes" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableScopes(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetScopes(*s)
	}
	return ophu
}

// SetAuthURL sets the "auth_url" field.
func (ophu *OauthProviderHistoryUpdate) SetAuthURL(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetAuthURL(s)
	return ophu
}

// SetNillableAuthURL sets the "auth_url" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableAuthURL(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetAuthURL(*s)
	}
	return ophu
}

// SetTokenURL sets the "token_url" field.
func (ophu *OauthProviderHistoryUpdate) SetTokenURL(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetTokenURL(s)
	return ophu
}

// SetNillableTokenURL sets the "token_url" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableTokenURL(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetTokenURL(*s)
	}
	return ophu
}

// SetAuthStyle sets the "auth_style" field.
func (ophu *OauthProviderHistoryUpdate) SetAuthStyle(u uint8) *OauthProviderHistoryUpdate {
	ophu.mutation.ResetAuthStyle()
	ophu.mutation.SetAuthStyle(u)
	return ophu
}

// SetNillableAuthStyle sets the "auth_style" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableAuthStyle(u *uint8) *OauthProviderHistoryUpdate {
	if u != nil {
		ophu.SetAuthStyle(*u)
	}
	return ophu
}

// AddAuthStyle adds u to the "auth_style" field.
func (ophu *OauthProviderHistoryUpdate) AddAuthStyle(u int8) *OauthProviderHistoryUpdate {
	ophu.mutation.AddAuthStyle(u)
	return ophu
}

// SetInfoURL sets the "info_url" field.
func (ophu *OauthProviderHistoryUpdate) SetInfoURL(s string) *OauthProviderHistoryUpdate {
	ophu.mutation.SetInfoURL(s)
	return ophu
}

// SetNillableInfoURL sets the "info_url" field if the given value is not nil.
func (ophu *OauthProviderHistoryUpdate) SetNillableInfoURL(s *string) *OauthProviderHistoryUpdate {
	if s != nil {
		ophu.SetInfoURL(*s)
	}
	return ophu
}

// Mutation returns the OauthProviderHistoryMutation object of the builder.
func (ophu *OauthProviderHistoryUpdate) Mutation() *OauthProviderHistoryMutation {
	return ophu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ophu *OauthProviderHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ophu.sqlSave, ophu.mutation, ophu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ophu *OauthProviderHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ophu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ophu *OauthProviderHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ophu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ophu *OauthProviderHistoryUpdate) ExecX(ctx context.Context) {
	if err := ophu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ophu *OauthProviderHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(oauthproviderhistory.Table, oauthproviderhistory.Columns, sqlgraph.NewFieldSpec(oauthproviderhistory.FieldID, field.TypeString))
	if ps := ophu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ophu.mutation.RefCleared() {
		_spec.ClearField(oauthproviderhistory.FieldRef, field.TypeString)
	}
	if value, ok := ophu.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthproviderhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ophu.mutation.CreatedByCleared() {
		_spec.ClearField(oauthproviderhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ophu.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthproviderhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ophu.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthproviderhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ophu.mutation.DeletedAt(); ok {
		_spec.SetField(oauthproviderhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ophu.mutation.DeletedAtCleared() {
		_spec.ClearField(oauthproviderhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ophu.mutation.DeletedBy(); ok {
		_spec.SetField(oauthproviderhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ophu.mutation.DeletedByCleared() {
		_spec.ClearField(oauthproviderhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ophu.mutation.Name(); ok {
		_spec.SetField(oauthproviderhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ophu.mutation.ClientID(); ok {
		_spec.SetField(oauthproviderhistory.FieldClientID, field.TypeString, value)
	}
	if value, ok := ophu.mutation.ClientSecret(); ok {
		_spec.SetField(oauthproviderhistory.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := ophu.mutation.RedirectURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldRedirectURL, field.TypeString, value)
	}
	if value, ok := ophu.mutation.Scopes(); ok {
		_spec.SetField(oauthproviderhistory.FieldScopes, field.TypeString, value)
	}
	if value, ok := ophu.mutation.AuthURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := ophu.mutation.TokenURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldTokenURL, field.TypeString, value)
	}
	if value, ok := ophu.mutation.AuthStyle(); ok {
		_spec.SetField(oauthproviderhistory.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := ophu.mutation.AddedAuthStyle(); ok {
		_spec.AddField(oauthproviderhistory.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := ophu.mutation.InfoURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldInfoURL, field.TypeString, value)
	}
	_spec.Node.Schema = ophu.schemaConfig.OauthProviderHistory
	ctx = internal.NewSchemaConfigContext(ctx, ophu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ophu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthproviderhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ophu.mutation.done = true
	return n, nil
}

// OauthProviderHistoryUpdateOne is the builder for updating a single OauthProviderHistory entity.
type OauthProviderHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OauthProviderHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetUpdatedAt(t time.Time) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetUpdatedAt(t)
	return ophuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *OauthProviderHistoryUpdateOne {
	if t != nil {
		ophuo.SetUpdatedAt(*t)
	}
	return ophuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetUpdatedBy(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetUpdatedBy(s)
	return ophuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableUpdatedBy(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetUpdatedBy(*s)
	}
	return ophuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ophuo *OauthProviderHistoryUpdateOne) ClearUpdatedBy() *OauthProviderHistoryUpdateOne {
	ophuo.mutation.ClearUpdatedBy()
	return ophuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetDeletedAt(t time.Time) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetDeletedAt(t)
	return ophuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *OauthProviderHistoryUpdateOne {
	if t != nil {
		ophuo.SetDeletedAt(*t)
	}
	return ophuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ophuo *OauthProviderHistoryUpdateOne) ClearDeletedAt() *OauthProviderHistoryUpdateOne {
	ophuo.mutation.ClearDeletedAt()
	return ophuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetDeletedBy(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetDeletedBy(s)
	return ophuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableDeletedBy(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetDeletedBy(*s)
	}
	return ophuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ophuo *OauthProviderHistoryUpdateOne) ClearDeletedBy() *OauthProviderHistoryUpdateOne {
	ophuo.mutation.ClearDeletedBy()
	return ophuo
}

// SetName sets the "name" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetName(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetName(s)
	return ophuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableName(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetName(*s)
	}
	return ophuo
}

// SetClientID sets the "client_id" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetClientID(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetClientID(s)
	return ophuo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableClientID(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetClientID(*s)
	}
	return ophuo
}

// SetClientSecret sets the "client_secret" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetClientSecret(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetClientSecret(s)
	return ophuo
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableClientSecret(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetClientSecret(*s)
	}
	return ophuo
}

// SetRedirectURL sets the "redirect_url" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetRedirectURL(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetRedirectURL(s)
	return ophuo
}

// SetNillableRedirectURL sets the "redirect_url" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableRedirectURL(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetRedirectURL(*s)
	}
	return ophuo
}

// SetScopes sets the "scopes" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetScopes(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetScopes(s)
	return ophuo
}

// SetNillableScopes sets the "scopes" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableScopes(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetScopes(*s)
	}
	return ophuo
}

// SetAuthURL sets the "auth_url" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetAuthURL(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetAuthURL(s)
	return ophuo
}

// SetNillableAuthURL sets the "auth_url" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableAuthURL(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetAuthURL(*s)
	}
	return ophuo
}

// SetTokenURL sets the "token_url" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetTokenURL(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetTokenURL(s)
	return ophuo
}

// SetNillableTokenURL sets the "token_url" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableTokenURL(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetTokenURL(*s)
	}
	return ophuo
}

// SetAuthStyle sets the "auth_style" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetAuthStyle(u uint8) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.ResetAuthStyle()
	ophuo.mutation.SetAuthStyle(u)
	return ophuo
}

// SetNillableAuthStyle sets the "auth_style" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableAuthStyle(u *uint8) *OauthProviderHistoryUpdateOne {
	if u != nil {
		ophuo.SetAuthStyle(*u)
	}
	return ophuo
}

// AddAuthStyle adds u to the "auth_style" field.
func (ophuo *OauthProviderHistoryUpdateOne) AddAuthStyle(u int8) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.AddAuthStyle(u)
	return ophuo
}

// SetInfoURL sets the "info_url" field.
func (ophuo *OauthProviderHistoryUpdateOne) SetInfoURL(s string) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.SetInfoURL(s)
	return ophuo
}

// SetNillableInfoURL sets the "info_url" field if the given value is not nil.
func (ophuo *OauthProviderHistoryUpdateOne) SetNillableInfoURL(s *string) *OauthProviderHistoryUpdateOne {
	if s != nil {
		ophuo.SetInfoURL(*s)
	}
	return ophuo
}

// Mutation returns the OauthProviderHistoryMutation object of the builder.
func (ophuo *OauthProviderHistoryUpdateOne) Mutation() *OauthProviderHistoryMutation {
	return ophuo.mutation
}

// Where appends a list predicates to the OauthProviderHistoryUpdate builder.
func (ophuo *OauthProviderHistoryUpdateOne) Where(ps ...predicate.OauthProviderHistory) *OauthProviderHistoryUpdateOne {
	ophuo.mutation.Where(ps...)
	return ophuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ophuo *OauthProviderHistoryUpdateOne) Select(field string, fields ...string) *OauthProviderHistoryUpdateOne {
	ophuo.fields = append([]string{field}, fields...)
	return ophuo
}

// Save executes the query and returns the updated OauthProviderHistory entity.
func (ophuo *OauthProviderHistoryUpdateOne) Save(ctx context.Context) (*OauthProviderHistory, error) {
	return withHooks(ctx, ophuo.sqlSave, ophuo.mutation, ophuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ophuo *OauthProviderHistoryUpdateOne) SaveX(ctx context.Context) *OauthProviderHistory {
	node, err := ophuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ophuo *OauthProviderHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ophuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ophuo *OauthProviderHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ophuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ophuo *OauthProviderHistoryUpdateOne) sqlSave(ctx context.Context) (_node *OauthProviderHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(oauthproviderhistory.Table, oauthproviderhistory.Columns, sqlgraph.NewFieldSpec(oauthproviderhistory.FieldID, field.TypeString))
	id, ok := ophuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OauthProviderHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ophuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthproviderhistory.FieldID)
		for _, f := range fields {
			if !oauthproviderhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != oauthproviderhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ophuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ophuo.mutation.RefCleared() {
		_spec.ClearField(oauthproviderhistory.FieldRef, field.TypeString)
	}
	if value, ok := ophuo.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthproviderhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ophuo.mutation.CreatedByCleared() {
		_spec.ClearField(oauthproviderhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ophuo.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthproviderhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ophuo.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthproviderhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ophuo.mutation.DeletedAt(); ok {
		_spec.SetField(oauthproviderhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ophuo.mutation.DeletedAtCleared() {
		_spec.ClearField(oauthproviderhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ophuo.mutation.DeletedBy(); ok {
		_spec.SetField(oauthproviderhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ophuo.mutation.DeletedByCleared() {
		_spec.ClearField(oauthproviderhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ophuo.mutation.Name(); ok {
		_spec.SetField(oauthproviderhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ophuo.mutation.ClientID(); ok {
		_spec.SetField(oauthproviderhistory.FieldClientID, field.TypeString, value)
	}
	if value, ok := ophuo.mutation.ClientSecret(); ok {
		_spec.SetField(oauthproviderhistory.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := ophuo.mutation.RedirectURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldRedirectURL, field.TypeString, value)
	}
	if value, ok := ophuo.mutation.Scopes(); ok {
		_spec.SetField(oauthproviderhistory.FieldScopes, field.TypeString, value)
	}
	if value, ok := ophuo.mutation.AuthURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := ophuo.mutation.TokenURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldTokenURL, field.TypeString, value)
	}
	if value, ok := ophuo.mutation.AuthStyle(); ok {
		_spec.SetField(oauthproviderhistory.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := ophuo.mutation.AddedAuthStyle(); ok {
		_spec.AddField(oauthproviderhistory.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := ophuo.mutation.InfoURL(); ok {
		_spec.SetField(oauthproviderhistory.FieldInfoURL, field.TypeString, value)
	}
	_spec.Node.Schema = ophuo.schemaConfig.OauthProviderHistory
	ctx = internal.NewSchemaConfigContext(ctx, ophuo.schemaConfig)
	_node = &OauthProviderHistory{config: ophuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ophuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthproviderhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ophuo.mutation.done = true
	return _node, nil
}
