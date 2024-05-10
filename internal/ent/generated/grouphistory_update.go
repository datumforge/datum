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
	"github.com/datumforge/datum/internal/ent/generated/grouphistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupHistoryUpdate is the builder for updating GroupHistory entities.
type GroupHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *GroupHistoryMutation
}

// Where appends a list predicates to the GroupHistoryUpdate builder.
func (ghu *GroupHistoryUpdate) Where(ps ...predicate.GroupHistory) *GroupHistoryUpdate {
	ghu.mutation.Where(ps...)
	return ghu
}

// SetUpdatedAt sets the "updated_at" field.
func (ghu *GroupHistoryUpdate) SetUpdatedAt(t time.Time) *GroupHistoryUpdate {
	ghu.mutation.SetUpdatedAt(t)
	return ghu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *GroupHistoryUpdate {
	if t != nil {
		ghu.SetUpdatedAt(*t)
	}
	return ghu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ghu *GroupHistoryUpdate) ClearUpdatedAt() *GroupHistoryUpdate {
	ghu.mutation.ClearUpdatedAt()
	return ghu
}

// SetUpdatedBy sets the "updated_by" field.
func (ghu *GroupHistoryUpdate) SetUpdatedBy(s string) *GroupHistoryUpdate {
	ghu.mutation.SetUpdatedBy(s)
	return ghu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableUpdatedBy(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetUpdatedBy(*s)
	}
	return ghu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ghu *GroupHistoryUpdate) ClearUpdatedBy() *GroupHistoryUpdate {
	ghu.mutation.ClearUpdatedBy()
	return ghu
}

// SetDeletedAt sets the "deleted_at" field.
func (ghu *GroupHistoryUpdate) SetDeletedAt(t time.Time) *GroupHistoryUpdate {
	ghu.mutation.SetDeletedAt(t)
	return ghu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableDeletedAt(t *time.Time) *GroupHistoryUpdate {
	if t != nil {
		ghu.SetDeletedAt(*t)
	}
	return ghu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ghu *GroupHistoryUpdate) ClearDeletedAt() *GroupHistoryUpdate {
	ghu.mutation.ClearDeletedAt()
	return ghu
}

// SetDeletedBy sets the "deleted_by" field.
func (ghu *GroupHistoryUpdate) SetDeletedBy(s string) *GroupHistoryUpdate {
	ghu.mutation.SetDeletedBy(s)
	return ghu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableDeletedBy(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetDeletedBy(*s)
	}
	return ghu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ghu *GroupHistoryUpdate) ClearDeletedBy() *GroupHistoryUpdate {
	ghu.mutation.ClearDeletedBy()
	return ghu
}

// SetOwnerID sets the "owner_id" field.
func (ghu *GroupHistoryUpdate) SetOwnerID(s string) *GroupHistoryUpdate {
	ghu.mutation.SetOwnerID(s)
	return ghu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableOwnerID(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetOwnerID(*s)
	}
	return ghu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (ghu *GroupHistoryUpdate) ClearOwnerID() *GroupHistoryUpdate {
	ghu.mutation.ClearOwnerID()
	return ghu
}

// SetName sets the "name" field.
func (ghu *GroupHistoryUpdate) SetName(s string) *GroupHistoryUpdate {
	ghu.mutation.SetName(s)
	return ghu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableName(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetName(*s)
	}
	return ghu
}

// SetDescription sets the "description" field.
func (ghu *GroupHistoryUpdate) SetDescription(s string) *GroupHistoryUpdate {
	ghu.mutation.SetDescription(s)
	return ghu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableDescription(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetDescription(*s)
	}
	return ghu
}

// ClearDescription clears the value of the "description" field.
func (ghu *GroupHistoryUpdate) ClearDescription() *GroupHistoryUpdate {
	ghu.mutation.ClearDescription()
	return ghu
}

// SetGravatarLogoURL sets the "gravatar_logo_url" field.
func (ghu *GroupHistoryUpdate) SetGravatarLogoURL(s string) *GroupHistoryUpdate {
	ghu.mutation.SetGravatarLogoURL(s)
	return ghu
}

// SetNillableGravatarLogoURL sets the "gravatar_logo_url" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableGravatarLogoURL(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetGravatarLogoURL(*s)
	}
	return ghu
}

// ClearGravatarLogoURL clears the value of the "gravatar_logo_url" field.
func (ghu *GroupHistoryUpdate) ClearGravatarLogoURL() *GroupHistoryUpdate {
	ghu.mutation.ClearGravatarLogoURL()
	return ghu
}

// SetLogoURL sets the "logo_url" field.
func (ghu *GroupHistoryUpdate) SetLogoURL(s string) *GroupHistoryUpdate {
	ghu.mutation.SetLogoURL(s)
	return ghu
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableLogoURL(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetLogoURL(*s)
	}
	return ghu
}

// ClearLogoURL clears the value of the "logo_url" field.
func (ghu *GroupHistoryUpdate) ClearLogoURL() *GroupHistoryUpdate {
	ghu.mutation.ClearLogoURL()
	return ghu
}

// SetDisplayName sets the "display_name" field.
func (ghu *GroupHistoryUpdate) SetDisplayName(s string) *GroupHistoryUpdate {
	ghu.mutation.SetDisplayName(s)
	return ghu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ghu *GroupHistoryUpdate) SetNillableDisplayName(s *string) *GroupHistoryUpdate {
	if s != nil {
		ghu.SetDisplayName(*s)
	}
	return ghu
}

// Mutation returns the GroupHistoryMutation object of the builder.
func (ghu *GroupHistoryUpdate) Mutation() *GroupHistoryMutation {
	return ghu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ghu *GroupHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ghu.sqlSave, ghu.mutation, ghu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ghu *GroupHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ghu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ghu *GroupHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ghu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghu *GroupHistoryUpdate) ExecX(ctx context.Context) {
	if err := ghu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ghu *GroupHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(grouphistory.Table, grouphistory.Columns, sqlgraph.NewFieldSpec(grouphistory.FieldID, field.TypeString))
	if ps := ghu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ghu.mutation.RefCleared() {
		_spec.ClearField(grouphistory.FieldRef, field.TypeString)
	}
	if ghu.mutation.CreatedAtCleared() {
		_spec.ClearField(grouphistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ghu.mutation.UpdatedAt(); ok {
		_spec.SetField(grouphistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ghu.mutation.UpdatedAtCleared() {
		_spec.ClearField(grouphistory.FieldUpdatedAt, field.TypeTime)
	}
	if ghu.mutation.CreatedByCleared() {
		_spec.ClearField(grouphistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ghu.mutation.UpdatedBy(); ok {
		_spec.SetField(grouphistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ghu.mutation.UpdatedByCleared() {
		_spec.ClearField(grouphistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ghu.mutation.DeletedAt(); ok {
		_spec.SetField(grouphistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ghu.mutation.DeletedAtCleared() {
		_spec.ClearField(grouphistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ghu.mutation.DeletedBy(); ok {
		_spec.SetField(grouphistory.FieldDeletedBy, field.TypeString, value)
	}
	if ghu.mutation.DeletedByCleared() {
		_spec.ClearField(grouphistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ghu.mutation.OwnerID(); ok {
		_spec.SetField(grouphistory.FieldOwnerID, field.TypeString, value)
	}
	if ghu.mutation.OwnerIDCleared() {
		_spec.ClearField(grouphistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := ghu.mutation.Name(); ok {
		_spec.SetField(grouphistory.FieldName, field.TypeString, value)
	}
	if value, ok := ghu.mutation.Description(); ok {
		_spec.SetField(grouphistory.FieldDescription, field.TypeString, value)
	}
	if ghu.mutation.DescriptionCleared() {
		_spec.ClearField(grouphistory.FieldDescription, field.TypeString)
	}
	if value, ok := ghu.mutation.GravatarLogoURL(); ok {
		_spec.SetField(grouphistory.FieldGravatarLogoURL, field.TypeString, value)
	}
	if ghu.mutation.GravatarLogoURLCleared() {
		_spec.ClearField(grouphistory.FieldGravatarLogoURL, field.TypeString)
	}
	if value, ok := ghu.mutation.LogoURL(); ok {
		_spec.SetField(grouphistory.FieldLogoURL, field.TypeString, value)
	}
	if ghu.mutation.LogoURLCleared() {
		_spec.ClearField(grouphistory.FieldLogoURL, field.TypeString)
	}
	if value, ok := ghu.mutation.DisplayName(); ok {
		_spec.SetField(grouphistory.FieldDisplayName, field.TypeString, value)
	}
	_spec.Node.Schema = ghu.schemaConfig.GroupHistory
	ctx = internal.NewSchemaConfigContext(ctx, ghu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ghu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grouphistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ghu.mutation.done = true
	return n, nil
}

// GroupHistoryUpdateOne is the builder for updating a single GroupHistory entity.
type GroupHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ghuo *GroupHistoryUpdateOne) SetUpdatedAt(t time.Time) *GroupHistoryUpdateOne {
	ghuo.mutation.SetUpdatedAt(t)
	return ghuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *GroupHistoryUpdateOne {
	if t != nil {
		ghuo.SetUpdatedAt(*t)
	}
	return ghuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ghuo *GroupHistoryUpdateOne) ClearUpdatedAt() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearUpdatedAt()
	return ghuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ghuo *GroupHistoryUpdateOne) SetUpdatedBy(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetUpdatedBy(s)
	return ghuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableUpdatedBy(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetUpdatedBy(*s)
	}
	return ghuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ghuo *GroupHistoryUpdateOne) ClearUpdatedBy() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearUpdatedBy()
	return ghuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ghuo *GroupHistoryUpdateOne) SetDeletedAt(t time.Time) *GroupHistoryUpdateOne {
	ghuo.mutation.SetDeletedAt(t)
	return ghuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupHistoryUpdateOne {
	if t != nil {
		ghuo.SetDeletedAt(*t)
	}
	return ghuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ghuo *GroupHistoryUpdateOne) ClearDeletedAt() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearDeletedAt()
	return ghuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ghuo *GroupHistoryUpdateOne) SetDeletedBy(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetDeletedBy(s)
	return ghuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableDeletedBy(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetDeletedBy(*s)
	}
	return ghuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ghuo *GroupHistoryUpdateOne) ClearDeletedBy() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearDeletedBy()
	return ghuo
}

// SetOwnerID sets the "owner_id" field.
func (ghuo *GroupHistoryUpdateOne) SetOwnerID(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetOwnerID(s)
	return ghuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableOwnerID(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetOwnerID(*s)
	}
	return ghuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (ghuo *GroupHistoryUpdateOne) ClearOwnerID() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearOwnerID()
	return ghuo
}

// SetName sets the "name" field.
func (ghuo *GroupHistoryUpdateOne) SetName(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetName(s)
	return ghuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableName(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetName(*s)
	}
	return ghuo
}

// SetDescription sets the "description" field.
func (ghuo *GroupHistoryUpdateOne) SetDescription(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetDescription(s)
	return ghuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableDescription(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetDescription(*s)
	}
	return ghuo
}

// ClearDescription clears the value of the "description" field.
func (ghuo *GroupHistoryUpdateOne) ClearDescription() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearDescription()
	return ghuo
}

// SetGravatarLogoURL sets the "gravatar_logo_url" field.
func (ghuo *GroupHistoryUpdateOne) SetGravatarLogoURL(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetGravatarLogoURL(s)
	return ghuo
}

// SetNillableGravatarLogoURL sets the "gravatar_logo_url" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableGravatarLogoURL(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetGravatarLogoURL(*s)
	}
	return ghuo
}

// ClearGravatarLogoURL clears the value of the "gravatar_logo_url" field.
func (ghuo *GroupHistoryUpdateOne) ClearGravatarLogoURL() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearGravatarLogoURL()
	return ghuo
}

// SetLogoURL sets the "logo_url" field.
func (ghuo *GroupHistoryUpdateOne) SetLogoURL(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetLogoURL(s)
	return ghuo
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableLogoURL(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetLogoURL(*s)
	}
	return ghuo
}

// ClearLogoURL clears the value of the "logo_url" field.
func (ghuo *GroupHistoryUpdateOne) ClearLogoURL() *GroupHistoryUpdateOne {
	ghuo.mutation.ClearLogoURL()
	return ghuo
}

// SetDisplayName sets the "display_name" field.
func (ghuo *GroupHistoryUpdateOne) SetDisplayName(s string) *GroupHistoryUpdateOne {
	ghuo.mutation.SetDisplayName(s)
	return ghuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ghuo *GroupHistoryUpdateOne) SetNillableDisplayName(s *string) *GroupHistoryUpdateOne {
	if s != nil {
		ghuo.SetDisplayName(*s)
	}
	return ghuo
}

// Mutation returns the GroupHistoryMutation object of the builder.
func (ghuo *GroupHistoryUpdateOne) Mutation() *GroupHistoryMutation {
	return ghuo.mutation
}

// Where appends a list predicates to the GroupHistoryUpdate builder.
func (ghuo *GroupHistoryUpdateOne) Where(ps ...predicate.GroupHistory) *GroupHistoryUpdateOne {
	ghuo.mutation.Where(ps...)
	return ghuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ghuo *GroupHistoryUpdateOne) Select(field string, fields ...string) *GroupHistoryUpdateOne {
	ghuo.fields = append([]string{field}, fields...)
	return ghuo
}

// Save executes the query and returns the updated GroupHistory entity.
func (ghuo *GroupHistoryUpdateOne) Save(ctx context.Context) (*GroupHistory, error) {
	return withHooks(ctx, ghuo.sqlSave, ghuo.mutation, ghuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ghuo *GroupHistoryUpdateOne) SaveX(ctx context.Context) *GroupHistory {
	node, err := ghuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ghuo *GroupHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ghuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghuo *GroupHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ghuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ghuo *GroupHistoryUpdateOne) sqlSave(ctx context.Context) (_node *GroupHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(grouphistory.Table, grouphistory.Columns, sqlgraph.NewFieldSpec(grouphistory.FieldID, field.TypeString))
	id, ok := ghuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "GroupHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ghuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grouphistory.FieldID)
		for _, f := range fields {
			if !grouphistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != grouphistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ghuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ghuo.mutation.RefCleared() {
		_spec.ClearField(grouphistory.FieldRef, field.TypeString)
	}
	if ghuo.mutation.CreatedAtCleared() {
		_spec.ClearField(grouphistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ghuo.mutation.UpdatedAt(); ok {
		_spec.SetField(grouphistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ghuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(grouphistory.FieldUpdatedAt, field.TypeTime)
	}
	if ghuo.mutation.CreatedByCleared() {
		_spec.ClearField(grouphistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ghuo.mutation.UpdatedBy(); ok {
		_spec.SetField(grouphistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ghuo.mutation.UpdatedByCleared() {
		_spec.ClearField(grouphistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ghuo.mutation.DeletedAt(); ok {
		_spec.SetField(grouphistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ghuo.mutation.DeletedAtCleared() {
		_spec.ClearField(grouphistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ghuo.mutation.DeletedBy(); ok {
		_spec.SetField(grouphistory.FieldDeletedBy, field.TypeString, value)
	}
	if ghuo.mutation.DeletedByCleared() {
		_spec.ClearField(grouphistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ghuo.mutation.OwnerID(); ok {
		_spec.SetField(grouphistory.FieldOwnerID, field.TypeString, value)
	}
	if ghuo.mutation.OwnerIDCleared() {
		_spec.ClearField(grouphistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := ghuo.mutation.Name(); ok {
		_spec.SetField(grouphistory.FieldName, field.TypeString, value)
	}
	if value, ok := ghuo.mutation.Description(); ok {
		_spec.SetField(grouphistory.FieldDescription, field.TypeString, value)
	}
	if ghuo.mutation.DescriptionCleared() {
		_spec.ClearField(grouphistory.FieldDescription, field.TypeString)
	}
	if value, ok := ghuo.mutation.GravatarLogoURL(); ok {
		_spec.SetField(grouphistory.FieldGravatarLogoURL, field.TypeString, value)
	}
	if ghuo.mutation.GravatarLogoURLCleared() {
		_spec.ClearField(grouphistory.FieldGravatarLogoURL, field.TypeString)
	}
	if value, ok := ghuo.mutation.LogoURL(); ok {
		_spec.SetField(grouphistory.FieldLogoURL, field.TypeString, value)
	}
	if ghuo.mutation.LogoURLCleared() {
		_spec.ClearField(grouphistory.FieldLogoURL, field.TypeString)
	}
	if value, ok := ghuo.mutation.DisplayName(); ok {
		_spec.SetField(grouphistory.FieldDisplayName, field.TypeString, value)
	}
	_spec.Node.Schema = ghuo.schemaConfig.GroupHistory
	ctx = internal.NewSchemaConfigContext(ctx, ghuo.schemaConfig)
	_node = &GroupHistory{config: ghuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ghuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grouphistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ghuo.mutation.done = true
	return _node, nil
}
