// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/oauthproviderhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OauthProviderHistoryQuery is the builder for querying OauthProviderHistory entities.
type OauthProviderHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []oauthproviderhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.OauthProviderHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*OauthProviderHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OauthProviderHistoryQuery builder.
func (ophq *OauthProviderHistoryQuery) Where(ps ...predicate.OauthProviderHistory) *OauthProviderHistoryQuery {
	ophq.predicates = append(ophq.predicates, ps...)
	return ophq
}

// Limit the number of records to be returned by this query.
func (ophq *OauthProviderHistoryQuery) Limit(limit int) *OauthProviderHistoryQuery {
	ophq.ctx.Limit = &limit
	return ophq
}

// Offset to start from.
func (ophq *OauthProviderHistoryQuery) Offset(offset int) *OauthProviderHistoryQuery {
	ophq.ctx.Offset = &offset
	return ophq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ophq *OauthProviderHistoryQuery) Unique(unique bool) *OauthProviderHistoryQuery {
	ophq.ctx.Unique = &unique
	return ophq
}

// Order specifies how the records should be ordered.
func (ophq *OauthProviderHistoryQuery) Order(o ...oauthproviderhistory.OrderOption) *OauthProviderHistoryQuery {
	ophq.order = append(ophq.order, o...)
	return ophq
}

// First returns the first OauthProviderHistory entity from the query.
// Returns a *NotFoundError when no OauthProviderHistory was found.
func (ophq *OauthProviderHistoryQuery) First(ctx context.Context) (*OauthProviderHistory, error) {
	nodes, err := ophq.Limit(1).All(setContextOp(ctx, ophq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oauthproviderhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) FirstX(ctx context.Context) *OauthProviderHistory {
	node, err := ophq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OauthProviderHistory ID from the query.
// Returns a *NotFoundError when no OauthProviderHistory ID was found.
func (ophq *OauthProviderHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ophq.Limit(1).IDs(setContextOp(ctx, ophq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauthproviderhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := ophq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OauthProviderHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OauthProviderHistory entity is found.
// Returns a *NotFoundError when no OauthProviderHistory entities are found.
func (ophq *OauthProviderHistoryQuery) Only(ctx context.Context) (*OauthProviderHistory, error) {
	nodes, err := ophq.Limit(2).All(setContextOp(ctx, ophq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oauthproviderhistory.Label}
	default:
		return nil, &NotSingularError{oauthproviderhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) OnlyX(ctx context.Context) *OauthProviderHistory {
	node, err := ophq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OauthProviderHistory ID in the query.
// Returns a *NotSingularError when more than one OauthProviderHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ophq *OauthProviderHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ophq.Limit(2).IDs(setContextOp(ctx, ophq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauthproviderhistory.Label}
	default:
		err = &NotSingularError{oauthproviderhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := ophq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OauthProviderHistories.
func (ophq *OauthProviderHistoryQuery) All(ctx context.Context) ([]*OauthProviderHistory, error) {
	ctx = setContextOp(ctx, ophq.ctx, ent.OpQueryAll)
	if err := ophq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OauthProviderHistory, *OauthProviderHistoryQuery]()
	return withInterceptors[[]*OauthProviderHistory](ctx, ophq, qr, ophq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) AllX(ctx context.Context) []*OauthProviderHistory {
	nodes, err := ophq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OauthProviderHistory IDs.
func (ophq *OauthProviderHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ophq.ctx.Unique == nil && ophq.path != nil {
		ophq.Unique(true)
	}
	ctx = setContextOp(ctx, ophq.ctx, ent.OpQueryIDs)
	if err = ophq.Select(oauthproviderhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := ophq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ophq *OauthProviderHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ophq.ctx, ent.OpQueryCount)
	if err := ophq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ophq, querierCount[*OauthProviderHistoryQuery](), ophq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) CountX(ctx context.Context) int {
	count, err := ophq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ophq *OauthProviderHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ophq.ctx, ent.OpQueryExist)
	switch _, err := ophq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ophq *OauthProviderHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ophq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OauthProviderHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ophq *OauthProviderHistoryQuery) Clone() *OauthProviderHistoryQuery {
	if ophq == nil {
		return nil
	}
	return &OauthProviderHistoryQuery{
		config:     ophq.config,
		ctx:        ophq.ctx.Clone(),
		order:      append([]oauthproviderhistory.OrderOption{}, ophq.order...),
		inters:     append([]Interceptor{}, ophq.inters...),
		predicates: append([]predicate.OauthProviderHistory{}, ophq.predicates...),
		// clone intermediate query.
		sql:  ophq.sql.Clone(),
		path: ophq.path,
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
//	client.OauthProviderHistory.Query().
//		GroupBy(oauthproviderhistory.FieldHistoryTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (ophq *OauthProviderHistoryQuery) GroupBy(field string, fields ...string) *OauthProviderHistoryGroupBy {
	ophq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OauthProviderHistoryGroupBy{build: ophq}
	grbuild.flds = &ophq.ctx.Fields
	grbuild.label = oauthproviderhistory.Label
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
//	client.OauthProviderHistory.Query().
//		Select(oauthproviderhistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (ophq *OauthProviderHistoryQuery) Select(fields ...string) *OauthProviderHistorySelect {
	ophq.ctx.Fields = append(ophq.ctx.Fields, fields...)
	sbuild := &OauthProviderHistorySelect{OauthProviderHistoryQuery: ophq}
	sbuild.label = oauthproviderhistory.Label
	sbuild.flds, sbuild.scan = &ophq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OauthProviderHistorySelect configured with the given aggregations.
func (ophq *OauthProviderHistoryQuery) Aggregate(fns ...AggregateFunc) *OauthProviderHistorySelect {
	return ophq.Select().Aggregate(fns...)
}

func (ophq *OauthProviderHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ophq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ophq); err != nil {
				return err
			}
		}
	}
	for _, f := range ophq.ctx.Fields {
		if !oauthproviderhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ophq.path != nil {
		prev, err := ophq.path(ctx)
		if err != nil {
			return err
		}
		ophq.sql = prev
	}
	if oauthproviderhistory.Policy == nil {
		return errors.New("generated: uninitialized oauthproviderhistory.Policy (forgotten import generated/runtime?)")
	}
	if err := oauthproviderhistory.Policy.EvalQuery(ctx, ophq); err != nil {
		return err
	}
	return nil
}

func (ophq *OauthProviderHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OauthProviderHistory, error) {
	var (
		nodes = []*OauthProviderHistory{}
		_spec = ophq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OauthProviderHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OauthProviderHistory{config: ophq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = ophq.schemaConfig.OauthProviderHistory
	ctx = internal.NewSchemaConfigContext(ctx, ophq.schemaConfig)
	if len(ophq.modifiers) > 0 {
		_spec.Modifiers = ophq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ophq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range ophq.loadTotal {
		if err := ophq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ophq *OauthProviderHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ophq.querySpec()
	_spec.Node.Schema = ophq.schemaConfig.OauthProviderHistory
	ctx = internal.NewSchemaConfigContext(ctx, ophq.schemaConfig)
	if len(ophq.modifiers) > 0 {
		_spec.Modifiers = ophq.modifiers
	}
	_spec.Node.Columns = ophq.ctx.Fields
	if len(ophq.ctx.Fields) > 0 {
		_spec.Unique = ophq.ctx.Unique != nil && *ophq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ophq.driver, _spec)
}

func (ophq *OauthProviderHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oauthproviderhistory.Table, oauthproviderhistory.Columns, sqlgraph.NewFieldSpec(oauthproviderhistory.FieldID, field.TypeString))
	_spec.From = ophq.sql
	if unique := ophq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ophq.path != nil {
		_spec.Unique = true
	}
	if fields := ophq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthproviderhistory.FieldID)
		for i := range fields {
			if fields[i] != oauthproviderhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ophq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ophq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ophq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ophq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ophq *OauthProviderHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ophq.driver.Dialect())
	t1 := builder.Table(oauthproviderhistory.Table)
	columns := ophq.ctx.Fields
	if len(columns) == 0 {
		columns = oauthproviderhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ophq.sql != nil {
		selector = ophq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ophq.ctx.Unique != nil && *ophq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(ophq.schemaConfig.OauthProviderHistory)
	ctx = internal.NewSchemaConfigContext(ctx, ophq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range ophq.predicates {
		p(selector)
	}
	for _, p := range ophq.order {
		p(selector)
	}
	if offset := ophq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ophq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OauthProviderHistoryGroupBy is the group-by builder for OauthProviderHistory entities.
type OauthProviderHistoryGroupBy struct {
	selector
	build *OauthProviderHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ophgb *OauthProviderHistoryGroupBy) Aggregate(fns ...AggregateFunc) *OauthProviderHistoryGroupBy {
	ophgb.fns = append(ophgb.fns, fns...)
	return ophgb
}

// Scan applies the selector query and scans the result into the given value.
func (ophgb *OauthProviderHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ophgb.build.ctx, ent.OpQueryGroupBy)
	if err := ophgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OauthProviderHistoryQuery, *OauthProviderHistoryGroupBy](ctx, ophgb.build, ophgb, ophgb.build.inters, v)
}

func (ophgb *OauthProviderHistoryGroupBy) sqlScan(ctx context.Context, root *OauthProviderHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ophgb.fns))
	for _, fn := range ophgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ophgb.flds)+len(ophgb.fns))
		for _, f := range *ophgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ophgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ophgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OauthProviderHistorySelect is the builder for selecting fields of OauthProviderHistory entities.
type OauthProviderHistorySelect struct {
	*OauthProviderHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ophs *OauthProviderHistorySelect) Aggregate(fns ...AggregateFunc) *OauthProviderHistorySelect {
	ophs.fns = append(ophs.fns, fns...)
	return ophs
}

// Scan applies the selector query and scans the result into the given value.
func (ophs *OauthProviderHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ophs.ctx, ent.OpQuerySelect)
	if err := ophs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OauthProviderHistoryQuery, *OauthProviderHistorySelect](ctx, ophs.OauthProviderHistoryQuery, ophs, ophs.inters, v)
}

func (ophs *OauthProviderHistorySelect) sqlScan(ctx context.Context, root *OauthProviderHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ophs.fns))
	for _, fn := range ophs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ophs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ophs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
