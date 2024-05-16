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
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupSettingUpdate is the builder for updating GroupSetting entities.
type GroupSettingUpdate struct {
	config
	hooks    []Hook
	mutation *GroupSettingMutation
}

// Where appends a list predicates to the GroupSettingUpdate builder.
func (gsu *GroupSettingUpdate) Where(ps ...predicate.GroupSetting) *GroupSettingUpdate {
	gsu.mutation.Where(ps...)
	return gsu
}

// SetUpdatedAt sets the "updated_at" field.
func (gsu *GroupSettingUpdate) SetUpdatedAt(t time.Time) *GroupSettingUpdate {
	gsu.mutation.SetUpdatedAt(t)
	return gsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gsu *GroupSettingUpdate) ClearUpdatedAt() *GroupSettingUpdate {
	gsu.mutation.ClearUpdatedAt()
	return gsu
}

// SetUpdatedBy sets the "updated_by" field.
func (gsu *GroupSettingUpdate) SetUpdatedBy(s string) *GroupSettingUpdate {
	gsu.mutation.SetUpdatedBy(s)
	return gsu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableUpdatedBy(s *string) *GroupSettingUpdate {
	if s != nil {
		gsu.SetUpdatedBy(*s)
	}
	return gsu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gsu *GroupSettingUpdate) ClearUpdatedBy() *GroupSettingUpdate {
	gsu.mutation.ClearUpdatedBy()
	return gsu
}

// SetTags sets the "tags" field.
func (gsu *GroupSettingUpdate) SetTags(s []string) *GroupSettingUpdate {
	gsu.mutation.SetTags(s)
	return gsu
}

// AppendTags appends s to the "tags" field.
func (gsu *GroupSettingUpdate) AppendTags(s []string) *GroupSettingUpdate {
	gsu.mutation.AppendTags(s)
	return gsu
}

// ClearTags clears the value of the "tags" field.
func (gsu *GroupSettingUpdate) ClearTags() *GroupSettingUpdate {
	gsu.mutation.ClearTags()
	return gsu
}

// SetDeletedAt sets the "deleted_at" field.
func (gsu *GroupSettingUpdate) SetDeletedAt(t time.Time) *GroupSettingUpdate {
	gsu.mutation.SetDeletedAt(t)
	return gsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableDeletedAt(t *time.Time) *GroupSettingUpdate {
	if t != nil {
		gsu.SetDeletedAt(*t)
	}
	return gsu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gsu *GroupSettingUpdate) ClearDeletedAt() *GroupSettingUpdate {
	gsu.mutation.ClearDeletedAt()
	return gsu
}

// SetDeletedBy sets the "deleted_by" field.
func (gsu *GroupSettingUpdate) SetDeletedBy(s string) *GroupSettingUpdate {
	gsu.mutation.SetDeletedBy(s)
	return gsu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableDeletedBy(s *string) *GroupSettingUpdate {
	if s != nil {
		gsu.SetDeletedBy(*s)
	}
	return gsu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (gsu *GroupSettingUpdate) ClearDeletedBy() *GroupSettingUpdate {
	gsu.mutation.ClearDeletedBy()
	return gsu
}

// SetVisibility sets the "visibility" field.
func (gsu *GroupSettingUpdate) SetVisibility(e enums.Visibility) *GroupSettingUpdate {
	gsu.mutation.SetVisibility(e)
	return gsu
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableVisibility(e *enums.Visibility) *GroupSettingUpdate {
	if e != nil {
		gsu.SetVisibility(*e)
	}
	return gsu
}

// SetJoinPolicy sets the "join_policy" field.
func (gsu *GroupSettingUpdate) SetJoinPolicy(ep enums.JoinPolicy) *GroupSettingUpdate {
	gsu.mutation.SetJoinPolicy(ep)
	return gsu
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableJoinPolicy(ep *enums.JoinPolicy) *GroupSettingUpdate {
	if ep != nil {
		gsu.SetJoinPolicy(*ep)
	}
	return gsu
}

// SetSyncToSlack sets the "sync_to_slack" field.
func (gsu *GroupSettingUpdate) SetSyncToSlack(b bool) *GroupSettingUpdate {
	gsu.mutation.SetSyncToSlack(b)
	return gsu
}

// SetNillableSyncToSlack sets the "sync_to_slack" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableSyncToSlack(b *bool) *GroupSettingUpdate {
	if b != nil {
		gsu.SetSyncToSlack(*b)
	}
	return gsu
}

// ClearSyncToSlack clears the value of the "sync_to_slack" field.
func (gsu *GroupSettingUpdate) ClearSyncToSlack() *GroupSettingUpdate {
	gsu.mutation.ClearSyncToSlack()
	return gsu
}

// SetSyncToGithub sets the "sync_to_github" field.
func (gsu *GroupSettingUpdate) SetSyncToGithub(b bool) *GroupSettingUpdate {
	gsu.mutation.SetSyncToGithub(b)
	return gsu
}

// SetNillableSyncToGithub sets the "sync_to_github" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableSyncToGithub(b *bool) *GroupSettingUpdate {
	if b != nil {
		gsu.SetSyncToGithub(*b)
	}
	return gsu
}

// ClearSyncToGithub clears the value of the "sync_to_github" field.
func (gsu *GroupSettingUpdate) ClearSyncToGithub() *GroupSettingUpdate {
	gsu.mutation.ClearSyncToGithub()
	return gsu
}

// SetGroupID sets the "group_id" field.
func (gsu *GroupSettingUpdate) SetGroupID(s string) *GroupSettingUpdate {
	gsu.mutation.SetGroupID(s)
	return gsu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (gsu *GroupSettingUpdate) SetNillableGroupID(s *string) *GroupSettingUpdate {
	if s != nil {
		gsu.SetGroupID(*s)
	}
	return gsu
}

// ClearGroupID clears the value of the "group_id" field.
func (gsu *GroupSettingUpdate) ClearGroupID() *GroupSettingUpdate {
	gsu.mutation.ClearGroupID()
	return gsu
}

// SetGroup sets the "group" edge to the Group entity.
func (gsu *GroupSettingUpdate) SetGroup(g *Group) *GroupSettingUpdate {
	return gsu.SetGroupID(g.ID)
}

// Mutation returns the GroupSettingMutation object of the builder.
func (gsu *GroupSettingUpdate) Mutation() *GroupSettingMutation {
	return gsu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gsu *GroupSettingUpdate) ClearGroup() *GroupSettingUpdate {
	gsu.mutation.ClearGroup()
	return gsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gsu *GroupSettingUpdate) Save(ctx context.Context) (int, error) {
	if err := gsu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, gsu.sqlSave, gsu.mutation, gsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gsu *GroupSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := gsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gsu *GroupSettingUpdate) Exec(ctx context.Context) error {
	_, err := gsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsu *GroupSettingUpdate) ExecX(ctx context.Context) {
	if err := gsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsu *GroupSettingUpdate) defaults() error {
	if _, ok := gsu.mutation.UpdatedAt(); !ok && !gsu.mutation.UpdatedAtCleared() {
		if groupsetting.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsetting.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupsetting.UpdateDefaultUpdatedAt()
		gsu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gsu *GroupSettingUpdate) check() error {
	if v, ok := gsu.mutation.Visibility(); ok {
		if err := groupsetting.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSetting.visibility": %w`, err)}
		}
	}
	if v, ok := gsu.mutation.JoinPolicy(); ok {
		if err := groupsetting.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSetting.join_policy": %w`, err)}
		}
	}
	return nil
}

func (gsu *GroupSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupsetting.Table, groupsetting.Columns, sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString))
	if ps := gsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gsu.mutation.CreatedAtCleared() {
		_spec.ClearField(groupsetting.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := gsu.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if gsu.mutation.UpdatedAtCleared() {
		_spec.ClearField(groupsetting.FieldUpdatedAt, field.TypeTime)
	}
	if gsu.mutation.CreatedByCleared() {
		_spec.ClearField(groupsetting.FieldCreatedBy, field.TypeString)
	}
	if value, ok := gsu.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsetting.FieldUpdatedBy, field.TypeString, value)
	}
	if gsu.mutation.UpdatedByCleared() {
		_spec.ClearField(groupsetting.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := gsu.mutation.Tags(); ok {
		_spec.SetField(groupsetting.FieldTags, field.TypeJSON, value)
	}
	if value, ok := gsu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, groupsetting.FieldTags, value)
		})
	}
	if gsu.mutation.TagsCleared() {
		_spec.ClearField(groupsetting.FieldTags, field.TypeJSON)
	}
	if value, ok := gsu.mutation.DeletedAt(); ok {
		_spec.SetField(groupsetting.FieldDeletedAt, field.TypeTime, value)
	}
	if gsu.mutation.DeletedAtCleared() {
		_spec.ClearField(groupsetting.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gsu.mutation.DeletedBy(); ok {
		_spec.SetField(groupsetting.FieldDeletedBy, field.TypeString, value)
	}
	if gsu.mutation.DeletedByCleared() {
		_spec.ClearField(groupsetting.FieldDeletedBy, field.TypeString)
	}
	if value, ok := gsu.mutation.Visibility(); ok {
		_spec.SetField(groupsetting.FieldVisibility, field.TypeEnum, value)
	}
	if value, ok := gsu.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsetting.FieldJoinPolicy, field.TypeEnum, value)
	}
	if value, ok := gsu.mutation.SyncToSlack(); ok {
		_spec.SetField(groupsetting.FieldSyncToSlack, field.TypeBool, value)
	}
	if gsu.mutation.SyncToSlackCleared() {
		_spec.ClearField(groupsetting.FieldSyncToSlack, field.TypeBool)
	}
	if value, ok := gsu.mutation.SyncToGithub(); ok {
		_spec.SetField(groupsetting.FieldSyncToGithub, field.TypeBool, value)
	}
	if gsu.mutation.SyncToGithubCleared() {
		_spec.ClearField(groupsetting.FieldSyncToGithub, field.TypeBool)
	}
	if gsu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsetting.GroupTable,
			Columns: []string{groupsetting.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = gsu.schemaConfig.GroupSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsetting.GroupTable,
			Columns: []string{groupsetting.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = gsu.schemaConfig.GroupSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = gsu.schemaConfig.GroupSetting
	ctx = internal.NewSchemaConfigContext(ctx, gsu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, gsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gsu.mutation.done = true
	return n, nil
}

// GroupSettingUpdateOne is the builder for updating a single GroupSetting entity.
type GroupSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupSettingMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gsuo *GroupSettingUpdateOne) SetUpdatedAt(t time.Time) *GroupSettingUpdateOne {
	gsuo.mutation.SetUpdatedAt(t)
	return gsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gsuo *GroupSettingUpdateOne) ClearUpdatedAt() *GroupSettingUpdateOne {
	gsuo.mutation.ClearUpdatedAt()
	return gsuo
}

// SetUpdatedBy sets the "updated_by" field.
func (gsuo *GroupSettingUpdateOne) SetUpdatedBy(s string) *GroupSettingUpdateOne {
	gsuo.mutation.SetUpdatedBy(s)
	return gsuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableUpdatedBy(s *string) *GroupSettingUpdateOne {
	if s != nil {
		gsuo.SetUpdatedBy(*s)
	}
	return gsuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gsuo *GroupSettingUpdateOne) ClearUpdatedBy() *GroupSettingUpdateOne {
	gsuo.mutation.ClearUpdatedBy()
	return gsuo
}

// SetTags sets the "tags" field.
func (gsuo *GroupSettingUpdateOne) SetTags(s []string) *GroupSettingUpdateOne {
	gsuo.mutation.SetTags(s)
	return gsuo
}

// AppendTags appends s to the "tags" field.
func (gsuo *GroupSettingUpdateOne) AppendTags(s []string) *GroupSettingUpdateOne {
	gsuo.mutation.AppendTags(s)
	return gsuo
}

// ClearTags clears the value of the "tags" field.
func (gsuo *GroupSettingUpdateOne) ClearTags() *GroupSettingUpdateOne {
	gsuo.mutation.ClearTags()
	return gsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gsuo *GroupSettingUpdateOne) SetDeletedAt(t time.Time) *GroupSettingUpdateOne {
	gsuo.mutation.SetDeletedAt(t)
	return gsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupSettingUpdateOne {
	if t != nil {
		gsuo.SetDeletedAt(*t)
	}
	return gsuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gsuo *GroupSettingUpdateOne) ClearDeletedAt() *GroupSettingUpdateOne {
	gsuo.mutation.ClearDeletedAt()
	return gsuo
}

// SetDeletedBy sets the "deleted_by" field.
func (gsuo *GroupSettingUpdateOne) SetDeletedBy(s string) *GroupSettingUpdateOne {
	gsuo.mutation.SetDeletedBy(s)
	return gsuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableDeletedBy(s *string) *GroupSettingUpdateOne {
	if s != nil {
		gsuo.SetDeletedBy(*s)
	}
	return gsuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (gsuo *GroupSettingUpdateOne) ClearDeletedBy() *GroupSettingUpdateOne {
	gsuo.mutation.ClearDeletedBy()
	return gsuo
}

// SetVisibility sets the "visibility" field.
func (gsuo *GroupSettingUpdateOne) SetVisibility(e enums.Visibility) *GroupSettingUpdateOne {
	gsuo.mutation.SetVisibility(e)
	return gsuo
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableVisibility(e *enums.Visibility) *GroupSettingUpdateOne {
	if e != nil {
		gsuo.SetVisibility(*e)
	}
	return gsuo
}

// SetJoinPolicy sets the "join_policy" field.
func (gsuo *GroupSettingUpdateOne) SetJoinPolicy(ep enums.JoinPolicy) *GroupSettingUpdateOne {
	gsuo.mutation.SetJoinPolicy(ep)
	return gsuo
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableJoinPolicy(ep *enums.JoinPolicy) *GroupSettingUpdateOne {
	if ep != nil {
		gsuo.SetJoinPolicy(*ep)
	}
	return gsuo
}

// SetSyncToSlack sets the "sync_to_slack" field.
func (gsuo *GroupSettingUpdateOne) SetSyncToSlack(b bool) *GroupSettingUpdateOne {
	gsuo.mutation.SetSyncToSlack(b)
	return gsuo
}

// SetNillableSyncToSlack sets the "sync_to_slack" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableSyncToSlack(b *bool) *GroupSettingUpdateOne {
	if b != nil {
		gsuo.SetSyncToSlack(*b)
	}
	return gsuo
}

// ClearSyncToSlack clears the value of the "sync_to_slack" field.
func (gsuo *GroupSettingUpdateOne) ClearSyncToSlack() *GroupSettingUpdateOne {
	gsuo.mutation.ClearSyncToSlack()
	return gsuo
}

// SetSyncToGithub sets the "sync_to_github" field.
func (gsuo *GroupSettingUpdateOne) SetSyncToGithub(b bool) *GroupSettingUpdateOne {
	gsuo.mutation.SetSyncToGithub(b)
	return gsuo
}

// SetNillableSyncToGithub sets the "sync_to_github" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableSyncToGithub(b *bool) *GroupSettingUpdateOne {
	if b != nil {
		gsuo.SetSyncToGithub(*b)
	}
	return gsuo
}

// ClearSyncToGithub clears the value of the "sync_to_github" field.
func (gsuo *GroupSettingUpdateOne) ClearSyncToGithub() *GroupSettingUpdateOne {
	gsuo.mutation.ClearSyncToGithub()
	return gsuo
}

// SetGroupID sets the "group_id" field.
func (gsuo *GroupSettingUpdateOne) SetGroupID(s string) *GroupSettingUpdateOne {
	gsuo.mutation.SetGroupID(s)
	return gsuo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (gsuo *GroupSettingUpdateOne) SetNillableGroupID(s *string) *GroupSettingUpdateOne {
	if s != nil {
		gsuo.SetGroupID(*s)
	}
	return gsuo
}

// ClearGroupID clears the value of the "group_id" field.
func (gsuo *GroupSettingUpdateOne) ClearGroupID() *GroupSettingUpdateOne {
	gsuo.mutation.ClearGroupID()
	return gsuo
}

// SetGroup sets the "group" edge to the Group entity.
func (gsuo *GroupSettingUpdateOne) SetGroup(g *Group) *GroupSettingUpdateOne {
	return gsuo.SetGroupID(g.ID)
}

// Mutation returns the GroupSettingMutation object of the builder.
func (gsuo *GroupSettingUpdateOne) Mutation() *GroupSettingMutation {
	return gsuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gsuo *GroupSettingUpdateOne) ClearGroup() *GroupSettingUpdateOne {
	gsuo.mutation.ClearGroup()
	return gsuo
}

// Where appends a list predicates to the GroupSettingUpdate builder.
func (gsuo *GroupSettingUpdateOne) Where(ps ...predicate.GroupSetting) *GroupSettingUpdateOne {
	gsuo.mutation.Where(ps...)
	return gsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gsuo *GroupSettingUpdateOne) Select(field string, fields ...string) *GroupSettingUpdateOne {
	gsuo.fields = append([]string{field}, fields...)
	return gsuo
}

// Save executes the query and returns the updated GroupSetting entity.
func (gsuo *GroupSettingUpdateOne) Save(ctx context.Context) (*GroupSetting, error) {
	if err := gsuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gsuo.sqlSave, gsuo.mutation, gsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gsuo *GroupSettingUpdateOne) SaveX(ctx context.Context) *GroupSetting {
	node, err := gsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gsuo *GroupSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := gsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsuo *GroupSettingUpdateOne) ExecX(ctx context.Context) {
	if err := gsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsuo *GroupSettingUpdateOne) defaults() error {
	if _, ok := gsuo.mutation.UpdatedAt(); !ok && !gsuo.mutation.UpdatedAtCleared() {
		if groupsetting.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsetting.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupsetting.UpdateDefaultUpdatedAt()
		gsuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gsuo *GroupSettingUpdateOne) check() error {
	if v, ok := gsuo.mutation.Visibility(); ok {
		if err := groupsetting.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSetting.visibility": %w`, err)}
		}
	}
	if v, ok := gsuo.mutation.JoinPolicy(); ok {
		if err := groupsetting.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSetting.join_policy": %w`, err)}
		}
	}
	return nil
}

func (gsuo *GroupSettingUpdateOne) sqlSave(ctx context.Context) (_node *GroupSetting, err error) {
	if err := gsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupsetting.Table, groupsetting.Columns, sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString))
	id, ok := gsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "GroupSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupsetting.FieldID)
		for _, f := range fields {
			if !groupsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != groupsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gsuo.mutation.CreatedAtCleared() {
		_spec.ClearField(groupsetting.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := gsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if gsuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(groupsetting.FieldUpdatedAt, field.TypeTime)
	}
	if gsuo.mutation.CreatedByCleared() {
		_spec.ClearField(groupsetting.FieldCreatedBy, field.TypeString)
	}
	if value, ok := gsuo.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsetting.FieldUpdatedBy, field.TypeString, value)
	}
	if gsuo.mutation.UpdatedByCleared() {
		_spec.ClearField(groupsetting.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := gsuo.mutation.Tags(); ok {
		_spec.SetField(groupsetting.FieldTags, field.TypeJSON, value)
	}
	if value, ok := gsuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, groupsetting.FieldTags, value)
		})
	}
	if gsuo.mutation.TagsCleared() {
		_spec.ClearField(groupsetting.FieldTags, field.TypeJSON)
	}
	if value, ok := gsuo.mutation.DeletedAt(); ok {
		_spec.SetField(groupsetting.FieldDeletedAt, field.TypeTime, value)
	}
	if gsuo.mutation.DeletedAtCleared() {
		_spec.ClearField(groupsetting.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gsuo.mutation.DeletedBy(); ok {
		_spec.SetField(groupsetting.FieldDeletedBy, field.TypeString, value)
	}
	if gsuo.mutation.DeletedByCleared() {
		_spec.ClearField(groupsetting.FieldDeletedBy, field.TypeString)
	}
	if value, ok := gsuo.mutation.Visibility(); ok {
		_spec.SetField(groupsetting.FieldVisibility, field.TypeEnum, value)
	}
	if value, ok := gsuo.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsetting.FieldJoinPolicy, field.TypeEnum, value)
	}
	if value, ok := gsuo.mutation.SyncToSlack(); ok {
		_spec.SetField(groupsetting.FieldSyncToSlack, field.TypeBool, value)
	}
	if gsuo.mutation.SyncToSlackCleared() {
		_spec.ClearField(groupsetting.FieldSyncToSlack, field.TypeBool)
	}
	if value, ok := gsuo.mutation.SyncToGithub(); ok {
		_spec.SetField(groupsetting.FieldSyncToGithub, field.TypeBool, value)
	}
	if gsuo.mutation.SyncToGithubCleared() {
		_spec.ClearField(groupsetting.FieldSyncToGithub, field.TypeBool)
	}
	if gsuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsetting.GroupTable,
			Columns: []string{groupsetting.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = gsuo.schemaConfig.GroupSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsetting.GroupTable,
			Columns: []string{groupsetting.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = gsuo.schemaConfig.GroupSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = gsuo.schemaConfig.GroupSetting
	ctx = internal.NewSchemaConfigContext(ctx, gsuo.schemaConfig)
	_node = &GroupSetting{config: gsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gsuo.mutation.done = true
	return _node, nil
}
