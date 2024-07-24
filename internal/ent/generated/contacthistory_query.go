// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/contacthistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ContactHistoryQuery is the builder for querying ContactHistory entities.
type ContactHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []contacthistory.OrderOption
	inters     []Interceptor
	predicates []predicate.ContactHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*ContactHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContactHistoryQuery builder.
func (chq *ContactHistoryQuery) Where(ps ...predicate.ContactHistory) *ContactHistoryQuery {
	chq.predicates = append(chq.predicates, ps...)
	return chq
}

// Limit the number of records to be returned by this query.
func (chq *ContactHistoryQuery) Limit(limit int) *ContactHistoryQuery {
	chq.ctx.Limit = &limit
	return chq
}

// Offset to start from.
func (chq *ContactHistoryQuery) Offset(offset int) *ContactHistoryQuery {
	chq.ctx.Offset = &offset
	return chq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (chq *ContactHistoryQuery) Unique(unique bool) *ContactHistoryQuery {
	chq.ctx.Unique = &unique
	return chq
}

// Order specifies how the records should be ordered.
func (chq *ContactHistoryQuery) Order(o ...contacthistory.OrderOption) *ContactHistoryQuery {
	chq.order = append(chq.order, o...)
	return chq
}

// First returns the first ContactHistory entity from the query.
// Returns a *NotFoundError when no ContactHistory was found.
func (chq *ContactHistoryQuery) First(ctx context.Context) (*ContactHistory, error) {
	nodes, err := chq.Limit(1).All(setContextOp(ctx, chq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contacthistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (chq *ContactHistoryQuery) FirstX(ctx context.Context) *ContactHistory {
	node, err := chq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ContactHistory ID from the query.
// Returns a *NotFoundError when no ContactHistory ID was found.
func (chq *ContactHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = chq.Limit(1).IDs(setContextOp(ctx, chq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contacthistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (chq *ContactHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := chq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ContactHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ContactHistory entity is found.
// Returns a *NotFoundError when no ContactHistory entities are found.
func (chq *ContactHistoryQuery) Only(ctx context.Context) (*ContactHistory, error) {
	nodes, err := chq.Limit(2).All(setContextOp(ctx, chq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contacthistory.Label}
	default:
		return nil, &NotSingularError{contacthistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (chq *ContactHistoryQuery) OnlyX(ctx context.Context) *ContactHistory {
	node, err := chq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ContactHistory ID in the query.
// Returns a *NotSingularError when more than one ContactHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (chq *ContactHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = chq.Limit(2).IDs(setContextOp(ctx, chq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contacthistory.Label}
	default:
		err = &NotSingularError{contacthistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (chq *ContactHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := chq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ContactHistories.
func (chq *ContactHistoryQuery) All(ctx context.Context) ([]*ContactHistory, error) {
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryAll)
	if err := chq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ContactHistory, *ContactHistoryQuery]()
	return withInterceptors[[]*ContactHistory](ctx, chq, qr, chq.inters)
}

// AllX is like All, but panics if an error occurs.
func (chq *ContactHistoryQuery) AllX(ctx context.Context) []*ContactHistory {
	nodes, err := chq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ContactHistory IDs.
func (chq *ContactHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if chq.ctx.Unique == nil && chq.path != nil {
		chq.Unique(true)
	}
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryIDs)
	if err = chq.Select(contacthistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (chq *ContactHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := chq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (chq *ContactHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryCount)
	if err := chq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, chq, querierCount[*ContactHistoryQuery](), chq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (chq *ContactHistoryQuery) CountX(ctx context.Context) int {
	count, err := chq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (chq *ContactHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryExist)
	switch _, err := chq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (chq *ContactHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := chq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContactHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (chq *ContactHistoryQuery) Clone() *ContactHistoryQuery {
	if chq == nil {
		return nil
	}
	return &ContactHistoryQuery{
		config:     chq.config,
		ctx:        chq.ctx.Clone(),
		order:      append([]contacthistory.OrderOption{}, chq.order...),
		inters:     append([]Interceptor{}, chq.inters...),
		predicates: append([]predicate.ContactHistory{}, chq.predicates...),
		// clone intermediate query.
		sql:  chq.sql.Clone(),
		path: chq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ContactHistory.Query().
//		GroupBy(contacthistory.FieldHistoryTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (chq *ContactHistoryQuery) GroupBy(field string, fields ...string) *ContactHistoryGroupBy {
	chq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ContactHistoryGroupBy{build: chq}
	grbuild.flds = &chq.ctx.Fields
	grbuild.label = contacthistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.ContactHistory.Query().
//		Select(contacthistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (chq *ContactHistoryQuery) Select(fields ...string) *ContactHistorySelect {
	chq.ctx.Fields = append(chq.ctx.Fields, fields...)
	sbuild := &ContactHistorySelect{ContactHistoryQuery: chq}
	sbuild.label = contacthistory.Label
	sbuild.flds, sbuild.scan = &chq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContactHistorySelect configured with the given aggregations.
func (chq *ContactHistoryQuery) Aggregate(fns ...AggregateFunc) *ContactHistorySelect {
	return chq.Select().Aggregate(fns...)
}

func (chq *ContactHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range chq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, chq); err != nil {
				return err
			}
		}
	}
	for _, f := range chq.ctx.Fields {
		if !contacthistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if chq.path != nil {
		prev, err := chq.path(ctx)
		if err != nil {
			return err
		}
		chq.sql = prev
	}
	return nil
}

func (chq *ContactHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ContactHistory, error) {
	var (
		nodes = []*ContactHistory{}
		_spec = chq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ContactHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ContactHistory{config: chq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = chq.schemaConfig.ContactHistory
	ctx = internal.NewSchemaConfigContext(ctx, chq.schemaConfig)
	if len(chq.modifiers) > 0 {
		_spec.Modifiers = chq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, chq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range chq.loadTotal {
		if err := chq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (chq *ContactHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := chq.querySpec()
	_spec.Node.Schema = chq.schemaConfig.ContactHistory
	ctx = internal.NewSchemaConfigContext(ctx, chq.schemaConfig)
	if len(chq.modifiers) > 0 {
		_spec.Modifiers = chq.modifiers
	}
	_spec.Node.Columns = chq.ctx.Fields
	if len(chq.ctx.Fields) > 0 {
		_spec.Unique = chq.ctx.Unique != nil && *chq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, chq.driver, _spec)
}

func (chq *ContactHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(contacthistory.Table, contacthistory.Columns, sqlgraph.NewFieldSpec(contacthistory.FieldID, field.TypeString))
	_spec.From = chq.sql
	if unique := chq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if chq.path != nil {
		_spec.Unique = true
	}
	if fields := chq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contacthistory.FieldID)
		for i := range fields {
			if fields[i] != contacthistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := chq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := chq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := chq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := chq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (chq *ContactHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(chq.driver.Dialect())
	t1 := builder.Table(contacthistory.Table)
	columns := chq.ctx.Fields
	if len(columns) == 0 {
		columns = contacthistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if chq.sql != nil {
		selector = chq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if chq.ctx.Unique != nil && *chq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(chq.schemaConfig.ContactHistory)
	ctx = internal.NewSchemaConfigContext(ctx, chq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range chq.predicates {
		p(selector)
	}
	for _, p := range chq.order {
		p(selector)
	}
	if offset := chq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := chq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ContactHistoryGroupBy is the group-by builder for ContactHistory entities.
type ContactHistoryGroupBy struct {
	selector
	build *ContactHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (chgb *ContactHistoryGroupBy) Aggregate(fns ...AggregateFunc) *ContactHistoryGroupBy {
	chgb.fns = append(chgb.fns, fns...)
	return chgb
}

// Scan applies the selector query and scans the result into the given value.
func (chgb *ContactHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chgb.build.ctx, ent.OpQueryGroupBy)
	if err := chgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContactHistoryQuery, *ContactHistoryGroupBy](ctx, chgb.build, chgb, chgb.build.inters, v)
}

func (chgb *ContactHistoryGroupBy) sqlScan(ctx context.Context, root *ContactHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(chgb.fns))
	for _, fn := range chgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*chgb.flds)+len(chgb.fns))
		for _, f := range *chgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*chgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ContactHistorySelect is the builder for selecting fields of ContactHistory entities.
type ContactHistorySelect struct {
	*ContactHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (chs *ContactHistorySelect) Aggregate(fns ...AggregateFunc) *ContactHistorySelect {
	chs.fns = append(chs.fns, fns...)
	return chs
}

// Scan applies the selector query and scans the result into the given value.
func (chs *ContactHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chs.ctx, ent.OpQuerySelect)
	if err := chs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContactHistoryQuery, *ContactHistorySelect](ctx, chs.ContactHistoryQuery, chs, chs.inters, v)
}

func (chs *ContactHistorySelect) sqlScan(ctx context.Context, root *ContactHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(chs.fns))
	for _, fn := range chs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*chs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
