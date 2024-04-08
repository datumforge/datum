// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/template"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// TemplateQuery is the builder for querying Template entities.
type TemplateQuery struct {
	config
	ctx        *QueryContext
	order      []template.OrderOption
	inters     []Interceptor
	predicates []predicate.Template
	withOwner  *OrganizationQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Template) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TemplateQuery builder.
func (tq *TemplateQuery) Where(ps ...predicate.Template) *TemplateQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TemplateQuery) Limit(limit int) *TemplateQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TemplateQuery) Offset(offset int) *TemplateQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TemplateQuery) Unique(unique bool) *TemplateQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TemplateQuery) Order(o ...template.OrderOption) *TemplateQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryOwner chains the current query on the "owner" edge.
func (tq *TemplateQuery) QueryOwner() *OrganizationQuery {
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
			sqlgraph.From(template.Table, template.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, template.OwnerTable, template.OwnerColumn),
		)
		schemaConfig := tq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Template
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Template entity from the query.
// Returns a *NotFoundError when no Template was found.
func (tq *TemplateQuery) First(ctx context.Context) (*Template, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{template.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TemplateQuery) FirstX(ctx context.Context) *Template {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Template ID from the query.
// Returns a *NotFoundError when no Template ID was found.
func (tq *TemplateQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{template.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TemplateQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Template entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Template entity is found.
// Returns a *NotFoundError when no Template entities are found.
func (tq *TemplateQuery) Only(ctx context.Context) (*Template, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{template.Label}
	default:
		return nil, &NotSingularError{template.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TemplateQuery) OnlyX(ctx context.Context) *Template {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Template ID in the query.
// Returns a *NotSingularError when more than one Template ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TemplateQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{template.Label}
	default:
		err = &NotSingularError{template.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TemplateQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Templates.
func (tq *TemplateQuery) All(ctx context.Context) ([]*Template, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Template, *TemplateQuery]()
	return withInterceptors[[]*Template](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TemplateQuery) AllX(ctx context.Context) []*Template {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Template IDs.
func (tq *TemplateQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(template.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TemplateQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TemplateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TemplateQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TemplateQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TemplateQuery) Exist(ctx context.Context) (bool, error) {
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
func (tq *TemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TemplateQuery) Clone() *TemplateQuery {
	if tq == nil {
		return nil
	}
	return &TemplateQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]template.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Template{}, tq.predicates...),
		withOwner:  tq.withOwner.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TemplateQuery) WithOwner(opts ...func(*OrganizationQuery)) *TemplateQuery {
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
//	client.Template.Query().
//		GroupBy(template.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (tq *TemplateQuery) GroupBy(field string, fields ...string) *TemplateGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TemplateGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = template.Label
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
//	client.Template.Query().
//		Select(template.FieldCreatedAt).
//		Scan(ctx, &v)
func (tq *TemplateQuery) Select(fields ...string) *TemplateSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TemplateSelect{TemplateQuery: tq}
	sbuild.label = template.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TemplateSelect configured with the given aggregations.
func (tq *TemplateQuery) Aggregate(fns ...AggregateFunc) *TemplateSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TemplateQuery) prepareQuery(ctx context.Context) error {
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
		if !template.ValidColumn(f) {
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
	return nil
}

func (tq *TemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Template, error) {
	var (
		nodes       = []*Template{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Template).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Template{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tq.schemaConfig.Template
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
			func(n *Template, e *Organization) { n.Edges.Owner = e }); err != nil {
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

func (tq *TemplateQuery) loadOwner(ctx context.Context, query *OrganizationQuery, nodes []*Template, init func(*Template), assign func(*Template, *Organization)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Template)
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

func (tq *TemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Schema = tq.schemaConfig.Template
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

func (tq *TemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(template.Table, template.Columns, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, template.FieldID)
		for i := range fields {
			if fields[i] != template.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withOwner != nil {
			_spec.Node.AddColumnOnce(template.FieldOwnerID)
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

func (tq *TemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(template.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = template.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tq.schemaConfig.Template)
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

// TemplateGroupBy is the group-by builder for Template entities.
type TemplateGroupBy struct {
	selector
	build *TemplateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TemplateGroupBy) Aggregate(fns ...AggregateFunc) *TemplateGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TemplateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TemplateQuery, *TemplateGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TemplateGroupBy) sqlScan(ctx context.Context, root *TemplateQuery, v any) error {
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

// TemplateSelect is the builder for selecting fields of Template entities.
type TemplateSelect struct {
	*TemplateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TemplateSelect) Aggregate(fns ...AggregateFunc) *TemplateSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TemplateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TemplateQuery, *TemplateSelect](ctx, ts.TemplateQuery, ts, ts.inters, v)
}

func (ts *TemplateSelect) sqlScan(ctx context.Context, root *TemplateQuery, v any) error {
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
