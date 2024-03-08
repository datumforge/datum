// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupSettingQuery is the builder for querying GroupSetting entities.
type GroupSettingQuery struct {
	config
	ctx        *QueryContext
	order      []groupsetting.OrderOption
	inters     []Interceptor
	predicates []predicate.GroupSetting
	withGroup  *GroupQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*GroupSetting) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupSettingQuery builder.
func (gsq *GroupSettingQuery) Where(ps ...predicate.GroupSetting) *GroupSettingQuery {
	gsq.predicates = append(gsq.predicates, ps...)
	return gsq
}

// Limit the number of records to be returned by this query.
func (gsq *GroupSettingQuery) Limit(limit int) *GroupSettingQuery {
	gsq.ctx.Limit = &limit
	return gsq
}

// Offset to start from.
func (gsq *GroupSettingQuery) Offset(offset int) *GroupSettingQuery {
	gsq.ctx.Offset = &offset
	return gsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gsq *GroupSettingQuery) Unique(unique bool) *GroupSettingQuery {
	gsq.ctx.Unique = &unique
	return gsq
}

// Order specifies how the records should be ordered.
func (gsq *GroupSettingQuery) Order(o ...groupsetting.OrderOption) *GroupSettingQuery {
	gsq.order = append(gsq.order, o...)
	return gsq
}

// QueryGroup chains the current query on the "group" edge.
func (gsq *GroupSettingQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: gsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupsetting.Table, groupsetting.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, groupsetting.GroupTable, groupsetting.GroupColumn),
		)
		schemaConfig := gsq.schemaConfig
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.GroupSetting
		fromU = sqlgraph.SetNeighbors(gsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupSetting entity from the query.
// Returns a *NotFoundError when no GroupSetting was found.
func (gsq *GroupSettingQuery) First(ctx context.Context) (*GroupSetting, error) {
	nodes, err := gsq.Limit(1).All(setContextOp(ctx, gsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gsq *GroupSettingQuery) FirstX(ctx context.Context) *GroupSetting {
	node, err := gsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupSetting ID from the query.
// Returns a *NotFoundError when no GroupSetting ID was found.
func (gsq *GroupSettingQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gsq.Limit(1).IDs(setContextOp(ctx, gsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gsq *GroupSettingQuery) FirstIDX(ctx context.Context) string {
	id, err := gsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupSetting entity is found.
// Returns a *NotFoundError when no GroupSetting entities are found.
func (gsq *GroupSettingQuery) Only(ctx context.Context) (*GroupSetting, error) {
	nodes, err := gsq.Limit(2).All(setContextOp(ctx, gsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupsetting.Label}
	default:
		return nil, &NotSingularError{groupsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gsq *GroupSettingQuery) OnlyX(ctx context.Context) *GroupSetting {
	node, err := gsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupSetting ID in the query.
// Returns a *NotSingularError when more than one GroupSetting ID is found.
// Returns a *NotFoundError when no entities are found.
func (gsq *GroupSettingQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gsq.Limit(2).IDs(setContextOp(ctx, gsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupsetting.Label}
	default:
		err = &NotSingularError{groupsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gsq *GroupSettingQuery) OnlyIDX(ctx context.Context) string {
	id, err := gsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupSettings.
func (gsq *GroupSettingQuery) All(ctx context.Context) ([]*GroupSetting, error) {
	ctx = setContextOp(ctx, gsq.ctx, "All")
	if err := gsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupSetting, *GroupSettingQuery]()
	return withInterceptors[[]*GroupSetting](ctx, gsq, qr, gsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gsq *GroupSettingQuery) AllX(ctx context.Context) []*GroupSetting {
	nodes, err := gsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupSetting IDs.
func (gsq *GroupSettingQuery) IDs(ctx context.Context) (ids []string, err error) {
	if gsq.ctx.Unique == nil && gsq.path != nil {
		gsq.Unique(true)
	}
	ctx = setContextOp(ctx, gsq.ctx, "IDs")
	if err = gsq.Select(groupsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gsq *GroupSettingQuery) IDsX(ctx context.Context) []string {
	ids, err := gsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gsq *GroupSettingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gsq.ctx, "Count")
	if err := gsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gsq, querierCount[*GroupSettingQuery](), gsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gsq *GroupSettingQuery) CountX(ctx context.Context) int {
	count, err := gsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gsq *GroupSettingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gsq.ctx, "Exist")
	switch _, err := gsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gsq *GroupSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := gsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gsq *GroupSettingQuery) Clone() *GroupSettingQuery {
	if gsq == nil {
		return nil
	}
	return &GroupSettingQuery{
		config:     gsq.config,
		ctx:        gsq.ctx.Clone(),
		order:      append([]groupsetting.OrderOption{}, gsq.order...),
		inters:     append([]Interceptor{}, gsq.inters...),
		predicates: append([]predicate.GroupSetting{}, gsq.predicates...),
		withGroup:  gsq.withGroup.Clone(),
		// clone intermediate query.
		sql:  gsq.sql.Clone(),
		path: gsq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (gsq *GroupSettingQuery) WithGroup(opts ...func(*GroupQuery)) *GroupSettingQuery {
	query := (&GroupClient{config: gsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gsq.withGroup = query
	return gsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupSetting.Query().
//		GroupBy(groupsetting.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (gsq *GroupSettingQuery) GroupBy(field string, fields ...string) *GroupSettingGroupBy {
	gsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupSettingGroupBy{build: gsq}
	grbuild.flds = &gsq.ctx.Fields
	grbuild.label = groupsetting.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.GroupSetting.Query().
//		Select(groupsetting.FieldCreatedAt).
//		Scan(ctx, &v)
func (gsq *GroupSettingQuery) Select(fields ...string) *GroupSettingSelect {
	gsq.ctx.Fields = append(gsq.ctx.Fields, fields...)
	sbuild := &GroupSettingSelect{GroupSettingQuery: gsq}
	sbuild.label = groupsetting.Label
	sbuild.flds, sbuild.scan = &gsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupSettingSelect configured with the given aggregations.
func (gsq *GroupSettingQuery) Aggregate(fns ...AggregateFunc) *GroupSettingSelect {
	return gsq.Select().Aggregate(fns...)
}

func (gsq *GroupSettingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gsq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gsq); err != nil {
				return err
			}
		}
	}
	for _, f := range gsq.ctx.Fields {
		if !groupsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if gsq.path != nil {
		prev, err := gsq.path(ctx)
		if err != nil {
			return err
		}
		gsq.sql = prev
	}
	if groupsetting.Policy == nil {
		return errors.New("generated: uninitialized groupsetting.Policy (forgotten import generated/runtime?)")
	}
	if err := groupsetting.Policy.EvalQuery(ctx, gsq); err != nil {
		return err
	}
	return nil
}

func (gsq *GroupSettingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupSetting, error) {
	var (
		nodes       = []*GroupSetting{}
		_spec       = gsq.querySpec()
		loadedTypes = [1]bool{
			gsq.withGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupSetting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupSetting{config: gsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = gsq.schemaConfig.GroupSetting
	ctx = internal.NewSchemaConfigContext(ctx, gsq.schemaConfig)
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gsq.withGroup; query != nil {
		if err := gsq.loadGroup(ctx, query, nodes, nil,
			func(n *GroupSetting, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	for i := range gsq.loadTotal {
		if err := gsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gsq *GroupSettingQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GroupSetting, init func(*GroupSetting), assign func(*GroupSetting, *Group)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*GroupSetting)
	for i := range nodes {
		fk := nodes[i].GroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gsq *GroupSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gsq.querySpec()
	_spec.Node.Schema = gsq.schemaConfig.GroupSetting
	ctx = internal.NewSchemaConfigContext(ctx, gsq.schemaConfig)
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	_spec.Node.Columns = gsq.ctx.Fields
	if len(gsq.ctx.Fields) > 0 {
		_spec.Unique = gsq.ctx.Unique != nil && *gsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gsq.driver, _spec)
}

func (gsq *GroupSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupsetting.Table, groupsetting.Columns, sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString))
	_spec.From = gsq.sql
	if unique := gsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gsq.path != nil {
		_spec.Unique = true
	}
	if fields := gsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupsetting.FieldID)
		for i := range fields {
			if fields[i] != groupsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if gsq.withGroup != nil {
			_spec.Node.AddColumnOnce(groupsetting.FieldGroupID)
		}
	}
	if ps := gsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gsq *GroupSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gsq.driver.Dialect())
	t1 := builder.Table(groupsetting.Table)
	columns := gsq.ctx.Fields
	if len(columns) == 0 {
		columns = groupsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gsq.sql != nil {
		selector = gsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gsq.ctx.Unique != nil && *gsq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(gsq.schemaConfig.GroupSetting)
	ctx = internal.NewSchemaConfigContext(ctx, gsq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range gsq.predicates {
		p(selector)
	}
	for _, p := range gsq.order {
		p(selector)
	}
	if offset := gsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupSettingGroupBy is the group-by builder for GroupSetting entities.
type GroupSettingGroupBy struct {
	selector
	build *GroupSettingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gsgb *GroupSettingGroupBy) Aggregate(fns ...AggregateFunc) *GroupSettingGroupBy {
	gsgb.fns = append(gsgb.fns, fns...)
	return gsgb
}

// Scan applies the selector query and scans the result into the given value.
func (gsgb *GroupSettingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gsgb.build.ctx, "GroupBy")
	if err := gsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupSettingQuery, *GroupSettingGroupBy](ctx, gsgb.build, gsgb, gsgb.build.inters, v)
}

func (gsgb *GroupSettingGroupBy) sqlScan(ctx context.Context, root *GroupSettingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gsgb.fns))
	for _, fn := range gsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gsgb.flds)+len(gsgb.fns))
		for _, f := range *gsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupSettingSelect is the builder for selecting fields of GroupSetting entities.
type GroupSettingSelect struct {
	*GroupSettingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gss *GroupSettingSelect) Aggregate(fns ...AggregateFunc) *GroupSettingSelect {
	gss.fns = append(gss.fns, fns...)
	return gss
}

// Scan applies the selector query and scans the result into the given value.
func (gss *GroupSettingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gss.ctx, "Select")
	if err := gss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupSettingQuery, *GroupSettingSelect](ctx, gss.GroupSettingQuery, gss, gss.inters, v)
}

func (gss *GroupSettingSelect) sqlScan(ctx context.Context, root *GroupSettingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gss.fns))
	for _, fn := range gss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
