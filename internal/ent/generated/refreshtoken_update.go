// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/refreshtoken"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// RefreshTokenUpdate is the builder for updating RefreshToken entities.
type RefreshTokenUpdate struct {
	config
	hooks    []Hook
	mutation *RefreshTokenMutation
}

// Where appends a list predicates to the RefreshTokenUpdate builder.
func (rtu *RefreshTokenUpdate) Where(ps ...predicate.RefreshToken) *RefreshTokenUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetClientID sets the "client_id" field.
func (rtu *RefreshTokenUpdate) SetClientID(s string) *RefreshTokenUpdate {
	rtu.mutation.SetClientID(s)
	return rtu
}

// SetScopes sets the "scopes" field.
func (rtu *RefreshTokenUpdate) SetScopes(s []string) *RefreshTokenUpdate {
	rtu.mutation.SetScopes(s)
	return rtu
}

// AppendScopes appends s to the "scopes" field.
func (rtu *RefreshTokenUpdate) AppendScopes(s []string) *RefreshTokenUpdate {
	rtu.mutation.AppendScopes(s)
	return rtu
}

// ClearScopes clears the value of the "scopes" field.
func (rtu *RefreshTokenUpdate) ClearScopes() *RefreshTokenUpdate {
	rtu.mutation.ClearScopes()
	return rtu
}

// SetNonce sets the "nonce" field.
func (rtu *RefreshTokenUpdate) SetNonce(s string) *RefreshTokenUpdate {
	rtu.mutation.SetNonce(s)
	return rtu
}

// SetClaimsUserID sets the "claims_user_id" field.
func (rtu *RefreshTokenUpdate) SetClaimsUserID(s string) *RefreshTokenUpdate {
	rtu.mutation.SetClaimsUserID(s)
	return rtu
}

// SetClaimsUsername sets the "claims_username" field.
func (rtu *RefreshTokenUpdate) SetClaimsUsername(s string) *RefreshTokenUpdate {
	rtu.mutation.SetClaimsUsername(s)
	return rtu
}

// SetClaimsEmail sets the "claims_email" field.
func (rtu *RefreshTokenUpdate) SetClaimsEmail(s string) *RefreshTokenUpdate {
	rtu.mutation.SetClaimsEmail(s)
	return rtu
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (rtu *RefreshTokenUpdate) SetClaimsEmailVerified(b bool) *RefreshTokenUpdate {
	rtu.mutation.SetClaimsEmailVerified(b)
	return rtu
}

// SetClaimsGroups sets the "claims_groups" field.
func (rtu *RefreshTokenUpdate) SetClaimsGroups(s []string) *RefreshTokenUpdate {
	rtu.mutation.SetClaimsGroups(s)
	return rtu
}

// AppendClaimsGroups appends s to the "claims_groups" field.
func (rtu *RefreshTokenUpdate) AppendClaimsGroups(s []string) *RefreshTokenUpdate {
	rtu.mutation.AppendClaimsGroups(s)
	return rtu
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (rtu *RefreshTokenUpdate) ClearClaimsGroups() *RefreshTokenUpdate {
	rtu.mutation.ClearClaimsGroups()
	return rtu
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (rtu *RefreshTokenUpdate) SetClaimsPreferredUsername(s string) *RefreshTokenUpdate {
	rtu.mutation.SetClaimsPreferredUsername(s)
	return rtu
}

// SetConnectorID sets the "connector_id" field.
func (rtu *RefreshTokenUpdate) SetConnectorID(s string) *RefreshTokenUpdate {
	rtu.mutation.SetConnectorID(s)
	return rtu
}

// SetConnectorData sets the "connector_data" field.
func (rtu *RefreshTokenUpdate) SetConnectorData(s []string) *RefreshTokenUpdate {
	rtu.mutation.SetConnectorData(s)
	return rtu
}

// AppendConnectorData appends s to the "connector_data" field.
func (rtu *RefreshTokenUpdate) AppendConnectorData(s []string) *RefreshTokenUpdate {
	rtu.mutation.AppendConnectorData(s)
	return rtu
}

// ClearConnectorData clears the value of the "connector_data" field.
func (rtu *RefreshTokenUpdate) ClearConnectorData() *RefreshTokenUpdate {
	rtu.mutation.ClearConnectorData()
	return rtu
}

// SetToken sets the "token" field.
func (rtu *RefreshTokenUpdate) SetToken(s string) *RefreshTokenUpdate {
	rtu.mutation.SetToken(s)
	return rtu
}

// SetObsoleteToken sets the "obsolete_token" field.
func (rtu *RefreshTokenUpdate) SetObsoleteToken(s string) *RefreshTokenUpdate {
	rtu.mutation.SetObsoleteToken(s)
	return rtu
}

// SetLastUsed sets the "last_used" field.
func (rtu *RefreshTokenUpdate) SetLastUsed(t time.Time) *RefreshTokenUpdate {
	rtu.mutation.SetLastUsed(t)
	return rtu
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (rtu *RefreshTokenUpdate) SetNillableLastUsed(t *time.Time) *RefreshTokenUpdate {
	if t != nil {
		rtu.SetLastUsed(*t)
	}
	return rtu
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtu *RefreshTokenUpdate) Mutation() *RefreshTokenMutation {
	return rtu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *RefreshTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rtu.sqlSave, rtu.mutation, rtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RefreshTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RefreshTokenUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RefreshTokenUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtu *RefreshTokenUpdate) check() error {
	if v, ok := rtu.mutation.ClientID(); ok {
		if err := refreshtoken.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.client_id": %w`, err)}
		}
	}
	if v, ok := rtu.mutation.Nonce(); ok {
		if err := refreshtoken.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.nonce": %w`, err)}
		}
	}
	if v, ok := rtu.mutation.ClaimsUserID(); ok {
		if err := refreshtoken.ClaimsUserIDValidator(v); err != nil {
			return &ValidationError{Name: "claims_user_id", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.claims_user_id": %w`, err)}
		}
	}
	if v, ok := rtu.mutation.ClaimsUsername(); ok {
		if err := refreshtoken.ClaimsUsernameValidator(v); err != nil {
			return &ValidationError{Name: "claims_username", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.claims_username": %w`, err)}
		}
	}
	if v, ok := rtu.mutation.ClaimsEmail(); ok {
		if err := refreshtoken.ClaimsEmailValidator(v); err != nil {
			return &ValidationError{Name: "claims_email", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.claims_email": %w`, err)}
		}
	}
	if v, ok := rtu.mutation.ConnectorID(); ok {
		if err := refreshtoken.ConnectorIDValidator(v); err != nil {
			return &ValidationError{Name: "connector_id", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.connector_id": %w`, err)}
		}
	}
	return nil
}

func (rtu *RefreshTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(refreshtoken.Table, refreshtoken.Columns, sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeString))
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.ClientID(); ok {
		_spec.SetField(refreshtoken.FieldClientID, field.TypeString, value)
	}
	if value, ok := rtu.mutation.Scopes(); ok {
		_spec.SetField(refreshtoken.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := rtu.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, refreshtoken.FieldScopes, value)
		})
	}
	if rtu.mutation.ScopesCleared() {
		_spec.ClearField(refreshtoken.FieldScopes, field.TypeJSON)
	}
	if value, ok := rtu.mutation.Nonce(); ok {
		_spec.SetField(refreshtoken.FieldNonce, field.TypeString, value)
	}
	if value, ok := rtu.mutation.ClaimsUserID(); ok {
		_spec.SetField(refreshtoken.FieldClaimsUserID, field.TypeString, value)
	}
	if value, ok := rtu.mutation.ClaimsUsername(); ok {
		_spec.SetField(refreshtoken.FieldClaimsUsername, field.TypeString, value)
	}
	if value, ok := rtu.mutation.ClaimsEmail(); ok {
		_spec.SetField(refreshtoken.FieldClaimsEmail, field.TypeString, value)
	}
	if value, ok := rtu.mutation.ClaimsEmailVerified(); ok {
		_spec.SetField(refreshtoken.FieldClaimsEmailVerified, field.TypeBool, value)
	}
	if value, ok := rtu.mutation.ClaimsGroups(); ok {
		_spec.SetField(refreshtoken.FieldClaimsGroups, field.TypeJSON, value)
	}
	if value, ok := rtu.mutation.AppendedClaimsGroups(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, refreshtoken.FieldClaimsGroups, value)
		})
	}
	if rtu.mutation.ClaimsGroupsCleared() {
		_spec.ClearField(refreshtoken.FieldClaimsGroups, field.TypeJSON)
	}
	if value, ok := rtu.mutation.ClaimsPreferredUsername(); ok {
		_spec.SetField(refreshtoken.FieldClaimsPreferredUsername, field.TypeString, value)
	}
	if value, ok := rtu.mutation.ConnectorID(); ok {
		_spec.SetField(refreshtoken.FieldConnectorID, field.TypeString, value)
	}
	if value, ok := rtu.mutation.ConnectorData(); ok {
		_spec.SetField(refreshtoken.FieldConnectorData, field.TypeJSON, value)
	}
	if value, ok := rtu.mutation.AppendedConnectorData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, refreshtoken.FieldConnectorData, value)
		})
	}
	if rtu.mutation.ConnectorDataCleared() {
		_spec.ClearField(refreshtoken.FieldConnectorData, field.TypeJSON)
	}
	if value, ok := rtu.mutation.Token(); ok {
		_spec.SetField(refreshtoken.FieldToken, field.TypeString, value)
	}
	if value, ok := rtu.mutation.ObsoleteToken(); ok {
		_spec.SetField(refreshtoken.FieldObsoleteToken, field.TypeString, value)
	}
	if value, ok := rtu.mutation.LastUsed(); ok {
		_spec.SetField(refreshtoken.FieldLastUsed, field.TypeTime, value)
	}
	_spec.Node.Schema = rtu.schemaConfig.RefreshToken
	ctx = internal.NewSchemaConfigContext(ctx, rtu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refreshtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rtu.mutation.done = true
	return n, nil
}

// RefreshTokenUpdateOne is the builder for updating a single RefreshToken entity.
type RefreshTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RefreshTokenMutation
}

// SetClientID sets the "client_id" field.
func (rtuo *RefreshTokenUpdateOne) SetClientID(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetClientID(s)
	return rtuo
}

// SetScopes sets the "scopes" field.
func (rtuo *RefreshTokenUpdateOne) SetScopes(s []string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetScopes(s)
	return rtuo
}

// AppendScopes appends s to the "scopes" field.
func (rtuo *RefreshTokenUpdateOne) AppendScopes(s []string) *RefreshTokenUpdateOne {
	rtuo.mutation.AppendScopes(s)
	return rtuo
}

// ClearScopes clears the value of the "scopes" field.
func (rtuo *RefreshTokenUpdateOne) ClearScopes() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearScopes()
	return rtuo
}

// SetNonce sets the "nonce" field.
func (rtuo *RefreshTokenUpdateOne) SetNonce(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetNonce(s)
	return rtuo
}

// SetClaimsUserID sets the "claims_user_id" field.
func (rtuo *RefreshTokenUpdateOne) SetClaimsUserID(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetClaimsUserID(s)
	return rtuo
}

// SetClaimsUsername sets the "claims_username" field.
func (rtuo *RefreshTokenUpdateOne) SetClaimsUsername(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetClaimsUsername(s)
	return rtuo
}

// SetClaimsEmail sets the "claims_email" field.
func (rtuo *RefreshTokenUpdateOne) SetClaimsEmail(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetClaimsEmail(s)
	return rtuo
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (rtuo *RefreshTokenUpdateOne) SetClaimsEmailVerified(b bool) *RefreshTokenUpdateOne {
	rtuo.mutation.SetClaimsEmailVerified(b)
	return rtuo
}

// SetClaimsGroups sets the "claims_groups" field.
func (rtuo *RefreshTokenUpdateOne) SetClaimsGroups(s []string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetClaimsGroups(s)
	return rtuo
}

// AppendClaimsGroups appends s to the "claims_groups" field.
func (rtuo *RefreshTokenUpdateOne) AppendClaimsGroups(s []string) *RefreshTokenUpdateOne {
	rtuo.mutation.AppendClaimsGroups(s)
	return rtuo
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (rtuo *RefreshTokenUpdateOne) ClearClaimsGroups() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearClaimsGroups()
	return rtuo
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (rtuo *RefreshTokenUpdateOne) SetClaimsPreferredUsername(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetClaimsPreferredUsername(s)
	return rtuo
}

// SetConnectorID sets the "connector_id" field.
func (rtuo *RefreshTokenUpdateOne) SetConnectorID(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetConnectorID(s)
	return rtuo
}

// SetConnectorData sets the "connector_data" field.
func (rtuo *RefreshTokenUpdateOne) SetConnectorData(s []string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetConnectorData(s)
	return rtuo
}

// AppendConnectorData appends s to the "connector_data" field.
func (rtuo *RefreshTokenUpdateOne) AppendConnectorData(s []string) *RefreshTokenUpdateOne {
	rtuo.mutation.AppendConnectorData(s)
	return rtuo
}

// ClearConnectorData clears the value of the "connector_data" field.
func (rtuo *RefreshTokenUpdateOne) ClearConnectorData() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearConnectorData()
	return rtuo
}

// SetToken sets the "token" field.
func (rtuo *RefreshTokenUpdateOne) SetToken(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetToken(s)
	return rtuo
}

// SetObsoleteToken sets the "obsolete_token" field.
func (rtuo *RefreshTokenUpdateOne) SetObsoleteToken(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetObsoleteToken(s)
	return rtuo
}

// SetLastUsed sets the "last_used" field.
func (rtuo *RefreshTokenUpdateOne) SetLastUsed(t time.Time) *RefreshTokenUpdateOne {
	rtuo.mutation.SetLastUsed(t)
	return rtuo
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (rtuo *RefreshTokenUpdateOne) SetNillableLastUsed(t *time.Time) *RefreshTokenUpdateOne {
	if t != nil {
		rtuo.SetLastUsed(*t)
	}
	return rtuo
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtuo *RefreshTokenUpdateOne) Mutation() *RefreshTokenMutation {
	return rtuo.mutation
}

// Where appends a list predicates to the RefreshTokenUpdate builder.
func (rtuo *RefreshTokenUpdateOne) Where(ps ...predicate.RefreshToken) *RefreshTokenUpdateOne {
	rtuo.mutation.Where(ps...)
	return rtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *RefreshTokenUpdateOne) Select(field string, fields ...string) *RefreshTokenUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated RefreshToken entity.
func (rtuo *RefreshTokenUpdateOne) Save(ctx context.Context) (*RefreshToken, error) {
	return withHooks(ctx, rtuo.sqlSave, rtuo.mutation, rtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RefreshTokenUpdateOne) SaveX(ctx context.Context) *RefreshToken {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *RefreshTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RefreshTokenUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtuo *RefreshTokenUpdateOne) check() error {
	if v, ok := rtuo.mutation.ClientID(); ok {
		if err := refreshtoken.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.client_id": %w`, err)}
		}
	}
	if v, ok := rtuo.mutation.Nonce(); ok {
		if err := refreshtoken.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.nonce": %w`, err)}
		}
	}
	if v, ok := rtuo.mutation.ClaimsUserID(); ok {
		if err := refreshtoken.ClaimsUserIDValidator(v); err != nil {
			return &ValidationError{Name: "claims_user_id", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.claims_user_id": %w`, err)}
		}
	}
	if v, ok := rtuo.mutation.ClaimsUsername(); ok {
		if err := refreshtoken.ClaimsUsernameValidator(v); err != nil {
			return &ValidationError{Name: "claims_username", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.claims_username": %w`, err)}
		}
	}
	if v, ok := rtuo.mutation.ClaimsEmail(); ok {
		if err := refreshtoken.ClaimsEmailValidator(v); err != nil {
			return &ValidationError{Name: "claims_email", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.claims_email": %w`, err)}
		}
	}
	if v, ok := rtuo.mutation.ConnectorID(); ok {
		if err := refreshtoken.ConnectorIDValidator(v); err != nil {
			return &ValidationError{Name: "connector_id", err: fmt.Errorf(`generated: validator failed for field "RefreshToken.connector_id": %w`, err)}
		}
	}
	return nil
}

func (rtuo *RefreshTokenUpdateOne) sqlSave(ctx context.Context) (_node *RefreshToken, err error) {
	if err := rtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(refreshtoken.Table, refreshtoken.Columns, sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeString))
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "RefreshToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, refreshtoken.FieldID)
		for _, f := range fields {
			if !refreshtoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != refreshtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.ClientID(); ok {
		_spec.SetField(refreshtoken.FieldClientID, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.Scopes(); ok {
		_spec.SetField(refreshtoken.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := rtuo.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, refreshtoken.FieldScopes, value)
		})
	}
	if rtuo.mutation.ScopesCleared() {
		_spec.ClearField(refreshtoken.FieldScopes, field.TypeJSON)
	}
	if value, ok := rtuo.mutation.Nonce(); ok {
		_spec.SetField(refreshtoken.FieldNonce, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.ClaimsUserID(); ok {
		_spec.SetField(refreshtoken.FieldClaimsUserID, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.ClaimsUsername(); ok {
		_spec.SetField(refreshtoken.FieldClaimsUsername, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.ClaimsEmail(); ok {
		_spec.SetField(refreshtoken.FieldClaimsEmail, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.ClaimsEmailVerified(); ok {
		_spec.SetField(refreshtoken.FieldClaimsEmailVerified, field.TypeBool, value)
	}
	if value, ok := rtuo.mutation.ClaimsGroups(); ok {
		_spec.SetField(refreshtoken.FieldClaimsGroups, field.TypeJSON, value)
	}
	if value, ok := rtuo.mutation.AppendedClaimsGroups(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, refreshtoken.FieldClaimsGroups, value)
		})
	}
	if rtuo.mutation.ClaimsGroupsCleared() {
		_spec.ClearField(refreshtoken.FieldClaimsGroups, field.TypeJSON)
	}
	if value, ok := rtuo.mutation.ClaimsPreferredUsername(); ok {
		_spec.SetField(refreshtoken.FieldClaimsPreferredUsername, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.ConnectorID(); ok {
		_spec.SetField(refreshtoken.FieldConnectorID, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.ConnectorData(); ok {
		_spec.SetField(refreshtoken.FieldConnectorData, field.TypeJSON, value)
	}
	if value, ok := rtuo.mutation.AppendedConnectorData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, refreshtoken.FieldConnectorData, value)
		})
	}
	if rtuo.mutation.ConnectorDataCleared() {
		_spec.ClearField(refreshtoken.FieldConnectorData, field.TypeJSON)
	}
	if value, ok := rtuo.mutation.Token(); ok {
		_spec.SetField(refreshtoken.FieldToken, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.ObsoleteToken(); ok {
		_spec.SetField(refreshtoken.FieldObsoleteToken, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.LastUsed(); ok {
		_spec.SetField(refreshtoken.FieldLastUsed, field.TypeTime, value)
	}
	_spec.Node.Schema = rtuo.schemaConfig.RefreshToken
	ctx = internal.NewSchemaConfigContext(ctx, rtuo.schemaConfig)
	_node = &RefreshToken{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refreshtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rtuo.mutation.done = true
	return _node, nil
}
