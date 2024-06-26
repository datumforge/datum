// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/entitlementplanfeaturehistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntitlementPlanFeatureHistoryQuery is the builder for querying EntitlementPlanFeatureHistory entities.
type EntitlementPlanFeatureHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []entitlementplanfeaturehistory.OrderOption
	inters     []Interceptor
	predicates []predicate.EntitlementPlanFeatureHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*EntitlementPlanFeatureHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntitlementPlanFeatureHistoryQuery builder.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Where(ps ...predicate.EntitlementPlanFeatureHistory) *EntitlementPlanFeatureHistoryQuery {
	epfhq.predicates = append(epfhq.predicates, ps...)
	return epfhq
}

// Limit the number of records to be returned by this query.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Limit(limit int) *EntitlementPlanFeatureHistoryQuery {
	epfhq.ctx.Limit = &limit
	return epfhq
}

// Offset to start from.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Offset(offset int) *EntitlementPlanFeatureHistoryQuery {
	epfhq.ctx.Offset = &offset
	return epfhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Unique(unique bool) *EntitlementPlanFeatureHistoryQuery {
	epfhq.ctx.Unique = &unique
	return epfhq
}

// Order specifies how the records should be ordered.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Order(o ...entitlementplanfeaturehistory.OrderOption) *EntitlementPlanFeatureHistoryQuery {
	epfhq.order = append(epfhq.order, o...)
	return epfhq
}

// First returns the first EntitlementPlanFeatureHistory entity from the query.
// Returns a *NotFoundError when no EntitlementPlanFeatureHistory was found.
func (epfhq *EntitlementPlanFeatureHistoryQuery) First(ctx context.Context) (*EntitlementPlanFeatureHistory, error) {
	nodes, err := epfhq.Limit(1).All(setContextOp(ctx, epfhq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entitlementplanfeaturehistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) FirstX(ctx context.Context) *EntitlementPlanFeatureHistory {
	node, err := epfhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntitlementPlanFeatureHistory ID from the query.
// Returns a *NotFoundError when no EntitlementPlanFeatureHistory ID was found.
func (epfhq *EntitlementPlanFeatureHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = epfhq.Limit(1).IDs(setContextOp(ctx, epfhq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entitlementplanfeaturehistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := epfhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntitlementPlanFeatureHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntitlementPlanFeatureHistory entity is found.
// Returns a *NotFoundError when no EntitlementPlanFeatureHistory entities are found.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Only(ctx context.Context) (*EntitlementPlanFeatureHistory, error) {
	nodes, err := epfhq.Limit(2).All(setContextOp(ctx, epfhq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entitlementplanfeaturehistory.Label}
	default:
		return nil, &NotSingularError{entitlementplanfeaturehistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) OnlyX(ctx context.Context) *EntitlementPlanFeatureHistory {
	node, err := epfhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntitlementPlanFeatureHistory ID in the query.
// Returns a *NotSingularError when more than one EntitlementPlanFeatureHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (epfhq *EntitlementPlanFeatureHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = epfhq.Limit(2).IDs(setContextOp(ctx, epfhq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entitlementplanfeaturehistory.Label}
	default:
		err = &NotSingularError{entitlementplanfeaturehistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := epfhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntitlementPlanFeatureHistories.
func (epfhq *EntitlementPlanFeatureHistoryQuery) All(ctx context.Context) ([]*EntitlementPlanFeatureHistory, error) {
	ctx = setContextOp(ctx, epfhq.ctx, "All")
	if err := epfhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EntitlementPlanFeatureHistory, *EntitlementPlanFeatureHistoryQuery]()
	return withInterceptors[[]*EntitlementPlanFeatureHistory](ctx, epfhq, qr, epfhq.inters)
}

// AllX is like All, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) AllX(ctx context.Context) []*EntitlementPlanFeatureHistory {
	nodes, err := epfhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntitlementPlanFeatureHistory IDs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if epfhq.ctx.Unique == nil && epfhq.path != nil {
		epfhq.Unique(true)
	}
	ctx = setContextOp(ctx, epfhq.ctx, "IDs")
	if err = epfhq.Select(entitlementplanfeaturehistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := epfhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, epfhq.ctx, "Count")
	if err := epfhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, epfhq, querierCount[*EntitlementPlanFeatureHistoryQuery](), epfhq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) CountX(ctx context.Context) int {
	count, err := epfhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, epfhq.ctx, "Exist")
	switch _, err := epfhq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (epfhq *EntitlementPlanFeatureHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := epfhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntitlementPlanFeatureHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Clone() *EntitlementPlanFeatureHistoryQuery {
	if epfhq == nil {
		return nil
	}
	return &EntitlementPlanFeatureHistoryQuery{
		config:     epfhq.config,
		ctx:        epfhq.ctx.Clone(),
		order:      append([]entitlementplanfeaturehistory.OrderOption{}, epfhq.order...),
		inters:     append([]Interceptor{}, epfhq.inters...),
		predicates: append([]predicate.EntitlementPlanFeatureHistory{}, epfhq.predicates...),
		// clone intermediate query.
		sql:  epfhq.sql.Clone(),
		path: epfhq.path,
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
//	client.EntitlementPlanFeatureHistory.Query().
//		GroupBy(entitlementplanfeaturehistory.FieldHistoryTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (epfhq *EntitlementPlanFeatureHistoryQuery) GroupBy(field string, fields ...string) *EntitlementPlanFeatureHistoryGroupBy {
	epfhq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EntitlementPlanFeatureHistoryGroupBy{build: epfhq}
	grbuild.flds = &epfhq.ctx.Fields
	grbuild.label = entitlementplanfeaturehistory.Label
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
//	client.EntitlementPlanFeatureHistory.Query().
//		Select(entitlementplanfeaturehistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (epfhq *EntitlementPlanFeatureHistoryQuery) Select(fields ...string) *EntitlementPlanFeatureHistorySelect {
	epfhq.ctx.Fields = append(epfhq.ctx.Fields, fields...)
	sbuild := &EntitlementPlanFeatureHistorySelect{EntitlementPlanFeatureHistoryQuery: epfhq}
	sbuild.label = entitlementplanfeaturehistory.Label
	sbuild.flds, sbuild.scan = &epfhq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EntitlementPlanFeatureHistorySelect configured with the given aggregations.
func (epfhq *EntitlementPlanFeatureHistoryQuery) Aggregate(fns ...AggregateFunc) *EntitlementPlanFeatureHistorySelect {
	return epfhq.Select().Aggregate(fns...)
}

func (epfhq *EntitlementPlanFeatureHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range epfhq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, epfhq); err != nil {
				return err
			}
		}
	}
	for _, f := range epfhq.ctx.Fields {
		if !entitlementplanfeaturehistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if epfhq.path != nil {
		prev, err := epfhq.path(ctx)
		if err != nil {
			return err
		}
		epfhq.sql = prev
	}
	return nil
}

func (epfhq *EntitlementPlanFeatureHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntitlementPlanFeatureHistory, error) {
	var (
		nodes = []*EntitlementPlanFeatureHistory{}
		_spec = epfhq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EntitlementPlanFeatureHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EntitlementPlanFeatureHistory{config: epfhq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = epfhq.schemaConfig.EntitlementPlanFeatureHistory
	ctx = internal.NewSchemaConfigContext(ctx, epfhq.schemaConfig)
	if len(epfhq.modifiers) > 0 {
		_spec.Modifiers = epfhq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, epfhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range epfhq.loadTotal {
		if err := epfhq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (epfhq *EntitlementPlanFeatureHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := epfhq.querySpec()
	_spec.Node.Schema = epfhq.schemaConfig.EntitlementPlanFeatureHistory
	ctx = internal.NewSchemaConfigContext(ctx, epfhq.schemaConfig)
	if len(epfhq.modifiers) > 0 {
		_spec.Modifiers = epfhq.modifiers
	}
	_spec.Node.Columns = epfhq.ctx.Fields
	if len(epfhq.ctx.Fields) > 0 {
		_spec.Unique = epfhq.ctx.Unique != nil && *epfhq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, epfhq.driver, _spec)
}

func (epfhq *EntitlementPlanFeatureHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entitlementplanfeaturehistory.Table, entitlementplanfeaturehistory.Columns, sqlgraph.NewFieldSpec(entitlementplanfeaturehistory.FieldID, field.TypeString))
	_spec.From = epfhq.sql
	if unique := epfhq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if epfhq.path != nil {
		_spec.Unique = true
	}
	if fields := epfhq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitlementplanfeaturehistory.FieldID)
		for i := range fields {
			if fields[i] != entitlementplanfeaturehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := epfhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := epfhq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := epfhq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := epfhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (epfhq *EntitlementPlanFeatureHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(epfhq.driver.Dialect())
	t1 := builder.Table(entitlementplanfeaturehistory.Table)
	columns := epfhq.ctx.Fields
	if len(columns) == 0 {
		columns = entitlementplanfeaturehistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if epfhq.sql != nil {
		selector = epfhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if epfhq.ctx.Unique != nil && *epfhq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(epfhq.schemaConfig.EntitlementPlanFeatureHistory)
	ctx = internal.NewSchemaConfigContext(ctx, epfhq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range epfhq.predicates {
		p(selector)
	}
	for _, p := range epfhq.order {
		p(selector)
	}
	if offset := epfhq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := epfhq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntitlementPlanFeatureHistoryGroupBy is the group-by builder for EntitlementPlanFeatureHistory entities.
type EntitlementPlanFeatureHistoryGroupBy struct {
	selector
	build *EntitlementPlanFeatureHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (epfhgb *EntitlementPlanFeatureHistoryGroupBy) Aggregate(fns ...AggregateFunc) *EntitlementPlanFeatureHistoryGroupBy {
	epfhgb.fns = append(epfhgb.fns, fns...)
	return epfhgb
}

// Scan applies the selector query and scans the result into the given value.
func (epfhgb *EntitlementPlanFeatureHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, epfhgb.build.ctx, "GroupBy")
	if err := epfhgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntitlementPlanFeatureHistoryQuery, *EntitlementPlanFeatureHistoryGroupBy](ctx, epfhgb.build, epfhgb, epfhgb.build.inters, v)
}

func (epfhgb *EntitlementPlanFeatureHistoryGroupBy) sqlScan(ctx context.Context, root *EntitlementPlanFeatureHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(epfhgb.fns))
	for _, fn := range epfhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*epfhgb.flds)+len(epfhgb.fns))
		for _, f := range *epfhgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*epfhgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := epfhgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EntitlementPlanFeatureHistorySelect is the builder for selecting fields of EntitlementPlanFeatureHistory entities.
type EntitlementPlanFeatureHistorySelect struct {
	*EntitlementPlanFeatureHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (epfhs *EntitlementPlanFeatureHistorySelect) Aggregate(fns ...AggregateFunc) *EntitlementPlanFeatureHistorySelect {
	epfhs.fns = append(epfhs.fns, fns...)
	return epfhs
}

// Scan applies the selector query and scans the result into the given value.
func (epfhs *EntitlementPlanFeatureHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, epfhs.ctx, "Select")
	if err := epfhs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntitlementPlanFeatureHistoryQuery, *EntitlementPlanFeatureHistorySelect](ctx, epfhs.EntitlementPlanFeatureHistoryQuery, epfhs, epfhs.inters, v)
}

func (epfhs *EntitlementPlanFeatureHistorySelect) sqlScan(ctx context.Context, root *EntitlementPlanFeatureHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(epfhs.fns))
	for _, fn := range epfhs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*epfhs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := epfhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
