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
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootoken"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OhAuthTooTokenUpdate is the builder for updating OhAuthTooToken entities.
type OhAuthTooTokenUpdate struct {
	config
	hooks    []Hook
	mutation *OhAuthTooTokenMutation
}

// Where appends a list predicates to the OhAuthTooTokenUpdate builder.
func (oattu *OhAuthTooTokenUpdate) Where(ps ...predicate.OhAuthTooToken) *OhAuthTooTokenUpdate {
	oattu.mutation.Where(ps...)
	return oattu
}

// SetClientID sets the "client_id" field.
func (oattu *OhAuthTooTokenUpdate) SetClientID(s string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetClientID(s)
	return oattu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableClientID(s *string) *OhAuthTooTokenUpdate {
	if s != nil {
		oattu.SetClientID(*s)
	}
	return oattu
}

// SetScopes sets the "scopes" field.
func (oattu *OhAuthTooTokenUpdate) SetScopes(s []string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetScopes(s)
	return oattu
}

// AppendScopes appends s to the "scopes" field.
func (oattu *OhAuthTooTokenUpdate) AppendScopes(s []string) *OhAuthTooTokenUpdate {
	oattu.mutation.AppendScopes(s)
	return oattu
}

// ClearScopes clears the value of the "scopes" field.
func (oattu *OhAuthTooTokenUpdate) ClearScopes() *OhAuthTooTokenUpdate {
	oattu.mutation.ClearScopes()
	return oattu
}

// SetNonce sets the "nonce" field.
func (oattu *OhAuthTooTokenUpdate) SetNonce(s string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetNonce(s)
	return oattu
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableNonce(s *string) *OhAuthTooTokenUpdate {
	if s != nil {
		oattu.SetNonce(*s)
	}
	return oattu
}

// SetClaimsUserID sets the "claims_user_id" field.
func (oattu *OhAuthTooTokenUpdate) SetClaimsUserID(s string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetClaimsUserID(s)
	return oattu
}

// SetNillableClaimsUserID sets the "claims_user_id" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableClaimsUserID(s *string) *OhAuthTooTokenUpdate {
	if s != nil {
		oattu.SetClaimsUserID(*s)
	}
	return oattu
}

// SetClaimsUsername sets the "claims_username" field.
func (oattu *OhAuthTooTokenUpdate) SetClaimsUsername(s string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetClaimsUsername(s)
	return oattu
}

// SetNillableClaimsUsername sets the "claims_username" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableClaimsUsername(s *string) *OhAuthTooTokenUpdate {
	if s != nil {
		oattu.SetClaimsUsername(*s)
	}
	return oattu
}

// SetClaimsEmail sets the "claims_email" field.
func (oattu *OhAuthTooTokenUpdate) SetClaimsEmail(s string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetClaimsEmail(s)
	return oattu
}

// SetNillableClaimsEmail sets the "claims_email" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableClaimsEmail(s *string) *OhAuthTooTokenUpdate {
	if s != nil {
		oattu.SetClaimsEmail(*s)
	}
	return oattu
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (oattu *OhAuthTooTokenUpdate) SetClaimsEmailVerified(b bool) *OhAuthTooTokenUpdate {
	oattu.mutation.SetClaimsEmailVerified(b)
	return oattu
}

// SetNillableClaimsEmailVerified sets the "claims_email_verified" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableClaimsEmailVerified(b *bool) *OhAuthTooTokenUpdate {
	if b != nil {
		oattu.SetClaimsEmailVerified(*b)
	}
	return oattu
}

// SetClaimsGroups sets the "claims_groups" field.
func (oattu *OhAuthTooTokenUpdate) SetClaimsGroups(s []string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetClaimsGroups(s)
	return oattu
}

// AppendClaimsGroups appends s to the "claims_groups" field.
func (oattu *OhAuthTooTokenUpdate) AppendClaimsGroups(s []string) *OhAuthTooTokenUpdate {
	oattu.mutation.AppendClaimsGroups(s)
	return oattu
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (oattu *OhAuthTooTokenUpdate) ClearClaimsGroups() *OhAuthTooTokenUpdate {
	oattu.mutation.ClearClaimsGroups()
	return oattu
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (oattu *OhAuthTooTokenUpdate) SetClaimsPreferredUsername(s string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetClaimsPreferredUsername(s)
	return oattu
}

// SetNillableClaimsPreferredUsername sets the "claims_preferred_username" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableClaimsPreferredUsername(s *string) *OhAuthTooTokenUpdate {
	if s != nil {
		oattu.SetClaimsPreferredUsername(*s)
	}
	return oattu
}

// SetConnectorID sets the "connector_id" field.
func (oattu *OhAuthTooTokenUpdate) SetConnectorID(s string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetConnectorID(s)
	return oattu
}

// SetNillableConnectorID sets the "connector_id" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableConnectorID(s *string) *OhAuthTooTokenUpdate {
	if s != nil {
		oattu.SetConnectorID(*s)
	}
	return oattu
}

// SetConnectorData sets the "connector_data" field.
func (oattu *OhAuthTooTokenUpdate) SetConnectorData(s []string) *OhAuthTooTokenUpdate {
	oattu.mutation.SetConnectorData(s)
	return oattu
}

// AppendConnectorData appends s to the "connector_data" field.
func (oattu *OhAuthTooTokenUpdate) AppendConnectorData(s []string) *OhAuthTooTokenUpdate {
	oattu.mutation.AppendConnectorData(s)
	return oattu
}

// ClearConnectorData clears the value of the "connector_data" field.
func (oattu *OhAuthTooTokenUpdate) ClearConnectorData() *OhAuthTooTokenUpdate {
	oattu.mutation.ClearConnectorData()
	return oattu
}

// SetLastUsed sets the "last_used" field.
func (oattu *OhAuthTooTokenUpdate) SetLastUsed(t time.Time) *OhAuthTooTokenUpdate {
	oattu.mutation.SetLastUsed(t)
	return oattu
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (oattu *OhAuthTooTokenUpdate) SetNillableLastUsed(t *time.Time) *OhAuthTooTokenUpdate {
	if t != nil {
		oattu.SetLastUsed(*t)
	}
	return oattu
}

// Mutation returns the OhAuthTooTokenMutation object of the builder.
func (oattu *OhAuthTooTokenUpdate) Mutation() *OhAuthTooTokenMutation {
	return oattu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oattu *OhAuthTooTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oattu.sqlSave, oattu.mutation, oattu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oattu *OhAuthTooTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := oattu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oattu *OhAuthTooTokenUpdate) Exec(ctx context.Context) error {
	_, err := oattu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oattu *OhAuthTooTokenUpdate) ExecX(ctx context.Context) {
	if err := oattu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oattu *OhAuthTooTokenUpdate) check() error {
	if v, ok := oattu.mutation.ClientID(); ok {
		if err := ohauthtootoken.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.client_id": %w`, err)}
		}
	}
	if v, ok := oattu.mutation.Nonce(); ok {
		if err := ohauthtootoken.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.nonce": %w`, err)}
		}
	}
	if v, ok := oattu.mutation.ClaimsUserID(); ok {
		if err := ohauthtootoken.ClaimsUserIDValidator(v); err != nil {
			return &ValidationError{Name: "claims_user_id", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.claims_user_id": %w`, err)}
		}
	}
	if v, ok := oattu.mutation.ClaimsUsername(); ok {
		if err := ohauthtootoken.ClaimsUsernameValidator(v); err != nil {
			return &ValidationError{Name: "claims_username", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.claims_username": %w`, err)}
		}
	}
	if v, ok := oattu.mutation.ClaimsEmail(); ok {
		if err := ohauthtootoken.ClaimsEmailValidator(v); err != nil {
			return &ValidationError{Name: "claims_email", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.claims_email": %w`, err)}
		}
	}
	if v, ok := oattu.mutation.ConnectorID(); ok {
		if err := ohauthtootoken.ConnectorIDValidator(v); err != nil {
			return &ValidationError{Name: "connector_id", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.connector_id": %w`, err)}
		}
	}
	return nil
}

func (oattu *OhAuthTooTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oattu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ohauthtootoken.Table, ohauthtootoken.Columns, sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString))
	if ps := oattu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oattu.mutation.ClientID(); ok {
		_spec.SetField(ohauthtootoken.FieldClientID, field.TypeString, value)
	}
	if value, ok := oattu.mutation.Scopes(); ok {
		_spec.SetField(ohauthtootoken.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := oattu.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ohauthtootoken.FieldScopes, value)
		})
	}
	if oattu.mutation.ScopesCleared() {
		_spec.ClearField(ohauthtootoken.FieldScopes, field.TypeJSON)
	}
	if value, ok := oattu.mutation.Nonce(); ok {
		_spec.SetField(ohauthtootoken.FieldNonce, field.TypeString, value)
	}
	if value, ok := oattu.mutation.ClaimsUserID(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsUserID, field.TypeString, value)
	}
	if value, ok := oattu.mutation.ClaimsUsername(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsUsername, field.TypeString, value)
	}
	if value, ok := oattu.mutation.ClaimsEmail(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsEmail, field.TypeString, value)
	}
	if value, ok := oattu.mutation.ClaimsEmailVerified(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsEmailVerified, field.TypeBool, value)
	}
	if value, ok := oattu.mutation.ClaimsGroups(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsGroups, field.TypeJSON, value)
	}
	if value, ok := oattu.mutation.AppendedClaimsGroups(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ohauthtootoken.FieldClaimsGroups, value)
		})
	}
	if oattu.mutation.ClaimsGroupsCleared() {
		_spec.ClearField(ohauthtootoken.FieldClaimsGroups, field.TypeJSON)
	}
	if value, ok := oattu.mutation.ClaimsPreferredUsername(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsPreferredUsername, field.TypeString, value)
	}
	if value, ok := oattu.mutation.ConnectorID(); ok {
		_spec.SetField(ohauthtootoken.FieldConnectorID, field.TypeString, value)
	}
	if value, ok := oattu.mutation.ConnectorData(); ok {
		_spec.SetField(ohauthtootoken.FieldConnectorData, field.TypeJSON, value)
	}
	if value, ok := oattu.mutation.AppendedConnectorData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ohauthtootoken.FieldConnectorData, value)
		})
	}
	if oattu.mutation.ConnectorDataCleared() {
		_spec.ClearField(ohauthtootoken.FieldConnectorData, field.TypeJSON)
	}
	if value, ok := oattu.mutation.LastUsed(); ok {
		_spec.SetField(ohauthtootoken.FieldLastUsed, field.TypeTime, value)
	}
	_spec.Node.Schema = oattu.schemaConfig.OhAuthTooToken
	ctx = internal.NewSchemaConfigContext(ctx, oattu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, oattu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ohauthtootoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oattu.mutation.done = true
	return n, nil
}

// OhAuthTooTokenUpdateOne is the builder for updating a single OhAuthTooToken entity.
type OhAuthTooTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OhAuthTooTokenMutation
}

// SetClientID sets the "client_id" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetClientID(s string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetClientID(s)
	return oattuo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableClientID(s *string) *OhAuthTooTokenUpdateOne {
	if s != nil {
		oattuo.SetClientID(*s)
	}
	return oattuo
}

// SetScopes sets the "scopes" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetScopes(s []string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetScopes(s)
	return oattuo
}

// AppendScopes appends s to the "scopes" field.
func (oattuo *OhAuthTooTokenUpdateOne) AppendScopes(s []string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.AppendScopes(s)
	return oattuo
}

// ClearScopes clears the value of the "scopes" field.
func (oattuo *OhAuthTooTokenUpdateOne) ClearScopes() *OhAuthTooTokenUpdateOne {
	oattuo.mutation.ClearScopes()
	return oattuo
}

// SetNonce sets the "nonce" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetNonce(s string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetNonce(s)
	return oattuo
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableNonce(s *string) *OhAuthTooTokenUpdateOne {
	if s != nil {
		oattuo.SetNonce(*s)
	}
	return oattuo
}

// SetClaimsUserID sets the "claims_user_id" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetClaimsUserID(s string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetClaimsUserID(s)
	return oattuo
}

// SetNillableClaimsUserID sets the "claims_user_id" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableClaimsUserID(s *string) *OhAuthTooTokenUpdateOne {
	if s != nil {
		oattuo.SetClaimsUserID(*s)
	}
	return oattuo
}

// SetClaimsUsername sets the "claims_username" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetClaimsUsername(s string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetClaimsUsername(s)
	return oattuo
}

// SetNillableClaimsUsername sets the "claims_username" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableClaimsUsername(s *string) *OhAuthTooTokenUpdateOne {
	if s != nil {
		oattuo.SetClaimsUsername(*s)
	}
	return oattuo
}

// SetClaimsEmail sets the "claims_email" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetClaimsEmail(s string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetClaimsEmail(s)
	return oattuo
}

// SetNillableClaimsEmail sets the "claims_email" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableClaimsEmail(s *string) *OhAuthTooTokenUpdateOne {
	if s != nil {
		oattuo.SetClaimsEmail(*s)
	}
	return oattuo
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetClaimsEmailVerified(b bool) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetClaimsEmailVerified(b)
	return oattuo
}

// SetNillableClaimsEmailVerified sets the "claims_email_verified" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableClaimsEmailVerified(b *bool) *OhAuthTooTokenUpdateOne {
	if b != nil {
		oattuo.SetClaimsEmailVerified(*b)
	}
	return oattuo
}

// SetClaimsGroups sets the "claims_groups" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetClaimsGroups(s []string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetClaimsGroups(s)
	return oattuo
}

// AppendClaimsGroups appends s to the "claims_groups" field.
func (oattuo *OhAuthTooTokenUpdateOne) AppendClaimsGroups(s []string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.AppendClaimsGroups(s)
	return oattuo
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (oattuo *OhAuthTooTokenUpdateOne) ClearClaimsGroups() *OhAuthTooTokenUpdateOne {
	oattuo.mutation.ClearClaimsGroups()
	return oattuo
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetClaimsPreferredUsername(s string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetClaimsPreferredUsername(s)
	return oattuo
}

// SetNillableClaimsPreferredUsername sets the "claims_preferred_username" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableClaimsPreferredUsername(s *string) *OhAuthTooTokenUpdateOne {
	if s != nil {
		oattuo.SetClaimsPreferredUsername(*s)
	}
	return oattuo
}

// SetConnectorID sets the "connector_id" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetConnectorID(s string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetConnectorID(s)
	return oattuo
}

// SetNillableConnectorID sets the "connector_id" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableConnectorID(s *string) *OhAuthTooTokenUpdateOne {
	if s != nil {
		oattuo.SetConnectorID(*s)
	}
	return oattuo
}

// SetConnectorData sets the "connector_data" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetConnectorData(s []string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetConnectorData(s)
	return oattuo
}

// AppendConnectorData appends s to the "connector_data" field.
func (oattuo *OhAuthTooTokenUpdateOne) AppendConnectorData(s []string) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.AppendConnectorData(s)
	return oattuo
}

// ClearConnectorData clears the value of the "connector_data" field.
func (oattuo *OhAuthTooTokenUpdateOne) ClearConnectorData() *OhAuthTooTokenUpdateOne {
	oattuo.mutation.ClearConnectorData()
	return oattuo
}

// SetLastUsed sets the "last_used" field.
func (oattuo *OhAuthTooTokenUpdateOne) SetLastUsed(t time.Time) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.SetLastUsed(t)
	return oattuo
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (oattuo *OhAuthTooTokenUpdateOne) SetNillableLastUsed(t *time.Time) *OhAuthTooTokenUpdateOne {
	if t != nil {
		oattuo.SetLastUsed(*t)
	}
	return oattuo
}

// Mutation returns the OhAuthTooTokenMutation object of the builder.
func (oattuo *OhAuthTooTokenUpdateOne) Mutation() *OhAuthTooTokenMutation {
	return oattuo.mutation
}

// Where appends a list predicates to the OhAuthTooTokenUpdate builder.
func (oattuo *OhAuthTooTokenUpdateOne) Where(ps ...predicate.OhAuthTooToken) *OhAuthTooTokenUpdateOne {
	oattuo.mutation.Where(ps...)
	return oattuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oattuo *OhAuthTooTokenUpdateOne) Select(field string, fields ...string) *OhAuthTooTokenUpdateOne {
	oattuo.fields = append([]string{field}, fields...)
	return oattuo
}

// Save executes the query and returns the updated OhAuthTooToken entity.
func (oattuo *OhAuthTooTokenUpdateOne) Save(ctx context.Context) (*OhAuthTooToken, error) {
	return withHooks(ctx, oattuo.sqlSave, oattuo.mutation, oattuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oattuo *OhAuthTooTokenUpdateOne) SaveX(ctx context.Context) *OhAuthTooToken {
	node, err := oattuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oattuo *OhAuthTooTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := oattuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oattuo *OhAuthTooTokenUpdateOne) ExecX(ctx context.Context) {
	if err := oattuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oattuo *OhAuthTooTokenUpdateOne) check() error {
	if v, ok := oattuo.mutation.ClientID(); ok {
		if err := ohauthtootoken.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.client_id": %w`, err)}
		}
	}
	if v, ok := oattuo.mutation.Nonce(); ok {
		if err := ohauthtootoken.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.nonce": %w`, err)}
		}
	}
	if v, ok := oattuo.mutation.ClaimsUserID(); ok {
		if err := ohauthtootoken.ClaimsUserIDValidator(v); err != nil {
			return &ValidationError{Name: "claims_user_id", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.claims_user_id": %w`, err)}
		}
	}
	if v, ok := oattuo.mutation.ClaimsUsername(); ok {
		if err := ohauthtootoken.ClaimsUsernameValidator(v); err != nil {
			return &ValidationError{Name: "claims_username", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.claims_username": %w`, err)}
		}
	}
	if v, ok := oattuo.mutation.ClaimsEmail(); ok {
		if err := ohauthtootoken.ClaimsEmailValidator(v); err != nil {
			return &ValidationError{Name: "claims_email", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.claims_email": %w`, err)}
		}
	}
	if v, ok := oattuo.mutation.ConnectorID(); ok {
		if err := ohauthtootoken.ConnectorIDValidator(v); err != nil {
			return &ValidationError{Name: "connector_id", err: fmt.Errorf(`generated: validator failed for field "OhAuthTooToken.connector_id": %w`, err)}
		}
	}
	return nil
}

func (oattuo *OhAuthTooTokenUpdateOne) sqlSave(ctx context.Context) (_node *OhAuthTooToken, err error) {
	if err := oattuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ohauthtootoken.Table, ohauthtootoken.Columns, sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString))
	id, ok := oattuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OhAuthTooToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oattuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ohauthtootoken.FieldID)
		for _, f := range fields {
			if !ohauthtootoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != ohauthtootoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oattuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oattuo.mutation.ClientID(); ok {
		_spec.SetField(ohauthtootoken.FieldClientID, field.TypeString, value)
	}
	if value, ok := oattuo.mutation.Scopes(); ok {
		_spec.SetField(ohauthtootoken.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := oattuo.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ohauthtootoken.FieldScopes, value)
		})
	}
	if oattuo.mutation.ScopesCleared() {
		_spec.ClearField(ohauthtootoken.FieldScopes, field.TypeJSON)
	}
	if value, ok := oattuo.mutation.Nonce(); ok {
		_spec.SetField(ohauthtootoken.FieldNonce, field.TypeString, value)
	}
	if value, ok := oattuo.mutation.ClaimsUserID(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsUserID, field.TypeString, value)
	}
	if value, ok := oattuo.mutation.ClaimsUsername(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsUsername, field.TypeString, value)
	}
	if value, ok := oattuo.mutation.ClaimsEmail(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsEmail, field.TypeString, value)
	}
	if value, ok := oattuo.mutation.ClaimsEmailVerified(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsEmailVerified, field.TypeBool, value)
	}
	if value, ok := oattuo.mutation.ClaimsGroups(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsGroups, field.TypeJSON, value)
	}
	if value, ok := oattuo.mutation.AppendedClaimsGroups(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ohauthtootoken.FieldClaimsGroups, value)
		})
	}
	if oattuo.mutation.ClaimsGroupsCleared() {
		_spec.ClearField(ohauthtootoken.FieldClaimsGroups, field.TypeJSON)
	}
	if value, ok := oattuo.mutation.ClaimsPreferredUsername(); ok {
		_spec.SetField(ohauthtootoken.FieldClaimsPreferredUsername, field.TypeString, value)
	}
	if value, ok := oattuo.mutation.ConnectorID(); ok {
		_spec.SetField(ohauthtootoken.FieldConnectorID, field.TypeString, value)
	}
	if value, ok := oattuo.mutation.ConnectorData(); ok {
		_spec.SetField(ohauthtootoken.FieldConnectorData, field.TypeJSON, value)
	}
	if value, ok := oattuo.mutation.AppendedConnectorData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ohauthtootoken.FieldConnectorData, value)
		})
	}
	if oattuo.mutation.ConnectorDataCleared() {
		_spec.ClearField(ohauthtootoken.FieldConnectorData, field.TypeJSON)
	}
	if value, ok := oattuo.mutation.LastUsed(); ok {
		_spec.SetField(ohauthtootoken.FieldLastUsed, field.TypeTime, value)
	}
	_spec.Node.Schema = oattuo.schemaConfig.OhAuthTooToken
	ctx = internal.NewSchemaConfigContext(ctx, oattuo.schemaConfig)
	_node = &OhAuthTooToken{config: oattuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oattuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ohauthtootoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oattuo.mutation.done = true
	return _node, nil
}
