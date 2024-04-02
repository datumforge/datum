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
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/tier"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// TierQuery is the builder for querying Tier entities.
type TierQuery struct {
	config
	ctx        *QueryContext
	order      []tier.OrderOption
	inters     []Interceptor
	predicates []predicate.Tier
	withOwner  *OrganizationQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Tier) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TierQuery builder.
func (tq *TierQuery) Where(ps ...predicate.Tier) *TierQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TierQuery) Limit(limit int) *TierQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TierQuery) Offset(offset int) *TierQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TierQuery) Unique(unique bool) *TierQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TierQuery) Order(o ...tier.OrderOption) *TierQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryOwner chains the current query on the "owner" edge.
func (tq *TierQuery) QueryOwner() *OrganizationQuery {
	query := (&OrganizationClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tier.Table, tier.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tier.OwnerTable, tier.OwnerColumn),
		)
		schemaConfig := tq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Tier
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tier entity from the query.
// Returns a *NotFoundError when no Tier was found.
func (tq *TierQuery) First(ctx context.Context) (*Tier, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tier.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TierQuery) FirstX(ctx context.Context) *Tier {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tier ID from the query.
// Returns a *NotFoundError when no Tier ID was found.
func (tq *TierQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tier.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TierQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tier entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tier entity is found.
// Returns a *NotFoundError when no Tier entities are found.
func (tq *TierQuery) Only(ctx context.Context) (*Tier, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tier.Label}
	default:
		return nil, &NotSingularError{tier.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TierQuery) OnlyX(ctx context.Context) *Tier {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tier ID in the query.
// Returns a *NotSingularError when more than one Tier ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TierQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tier.Label}
	default:
		err = &NotSingularError{tier.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TierQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tiers.
func (tq *TierQuery) All(ctx context.Context) ([]*Tier, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tier, *TierQuery]()
	return withInterceptors[[]*Tier](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TierQuery) AllX(ctx context.Context) []*Tier {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tier IDs.
func (tq *TierQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(tier.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TierQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TierQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TierQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TierQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TierQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TierQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TierQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TierQuery) Clone() *TierQuery {
	if tq == nil {
		return nil
	}
	return &TierQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]tier.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Tier{}, tq.predicates...),
		withOwner:  tq.withOwner.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TierQuery) WithOwner(opts ...func(*OrganizationQuery)) *TierQuery {
	query := (&OrganizationClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withOwner = query
	return tq
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
//	client.Tier.Query().
//		GroupBy(tier.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (tq *TierQuery) GroupBy(field string, fields ...string) *TierGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TierGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = tier.Label
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
//	client.Tier.Query().
//		Select(tier.FieldCreatedAt).
//		Scan(ctx, &v)
func (tq *TierQuery) Select(fields ...string) *TierSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TierSelect{TierQuery: tq}
	sbuild.label = tier.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TierSelect configured with the given aggregations.
func (tq *TierQuery) Aggregate(fns ...AggregateFunc) *TierSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TierQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !tier.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	if tier.Policy == nil {
		return errors.New("generated: uninitialized tier.Policy (forgotten import generated/runtime?)")
	}
	if err := tier.Policy.EvalQuery(ctx, tq); err != nil {
		return err
	}
	return nil
}

func (tq *TierQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tier, error) {
	var (
		nodes       = []*Tier{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tier).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tier{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tq.schemaConfig.Tier
	ctx = internal.NewSchemaConfigContext(ctx, tq.schemaConfig)
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withOwner; query != nil {
		if err := tq.loadOwner(ctx, query, nodes, nil,
			func(n *Tier, e *Organization) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TierQuery) loadOwner(ctx context.Context, query *OrganizationQuery, nodes []*Tier, init func(*Tier), assign func(*Tier, *Organization)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Tier)
	for i := range nodes {
		fk := nodes[i].OwnerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "owner_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TierQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Schema = tq.schemaConfig.Tier
	ctx = internal.NewSchemaConfigContext(ctx, tq.schemaConfig)
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TierQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tier.Table, tier.Columns, sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tier.FieldID)
		for i := range fields {
			if fields[i] != tier.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withOwner != nil {
			_spec.Node.AddColumnOnce(tier.FieldOwnerID)
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TierQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tier.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = tier.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tq.schemaConfig.Tier)
	ctx = internal.NewSchemaConfigContext(ctx, tq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TierGroupBy is the group-by builder for Tier entities.
type TierGroupBy struct {
	selector
	build *TierQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TierGroupBy) Aggregate(fns ...AggregateFunc) *TierGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TierGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TierQuery, *TierGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TierGroupBy) sqlScan(ctx context.Context, root *TierQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TierSelect is the builder for selecting fields of Tier entities.
type TierSelect struct {
	*TierQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TierSelect) Aggregate(fns ...AggregateFunc) *TierSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TierSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TierQuery, *TierSelect](ctx, ts.TierQuery, ts, ts.inters, v)
}

func (ts *TierSelect) sqlScan(ctx context.Context, root *TierQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
