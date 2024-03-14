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
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// IntegrationQuery is the builder for querying Integration entities.
type IntegrationQuery struct {
	config
	ctx        *QueryContext
	order      []integration.OrderOption
	inters     []Interceptor
	predicates []predicate.Integration
	withOwner  *OrganizationQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Integration) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IntegrationQuery builder.
func (iq *IntegrationQuery) Where(ps ...predicate.Integration) *IntegrationQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *IntegrationQuery) Limit(limit int) *IntegrationQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *IntegrationQuery) Offset(offset int) *IntegrationQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IntegrationQuery) Unique(unique bool) *IntegrationQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *IntegrationQuery) Order(o ...integration.OrderOption) *IntegrationQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryOwner chains the current query on the "owner" edge.
func (iq *IntegrationQuery) QueryOwner() *OrganizationQuery {
	query := (&OrganizationClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(integration.Table, integration.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, integration.OwnerTable, integration.OwnerColumn),
		)
		schemaConfig := iq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Integration
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Integration entity from the query.
// Returns a *NotFoundError when no Integration was found.
func (iq *IntegrationQuery) First(ctx context.Context) (*Integration, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{integration.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IntegrationQuery) FirstX(ctx context.Context) *Integration {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Integration ID from the query.
// Returns a *NotFoundError when no Integration ID was found.
func (iq *IntegrationQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{integration.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IntegrationQuery) FirstIDX(ctx context.Context) string {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Integration entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Integration entity is found.
// Returns a *NotFoundError when no Integration entities are found.
func (iq *IntegrationQuery) Only(ctx context.Context) (*Integration, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{integration.Label}
	default:
		return nil, &NotSingularError{integration.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IntegrationQuery) OnlyX(ctx context.Context) *Integration {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Integration ID in the query.
// Returns a *NotSingularError when more than one Integration ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *IntegrationQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{integration.Label}
	default:
		err = &NotSingularError{integration.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IntegrationQuery) OnlyIDX(ctx context.Context) string {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Integrations.
func (iq *IntegrationQuery) All(ctx context.Context) ([]*Integration, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Integration, *IntegrationQuery]()
	return withInterceptors[[]*Integration](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *IntegrationQuery) AllX(ctx context.Context) []*Integration {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Integration IDs.
func (iq *IntegrationQuery) IDs(ctx context.Context) (ids []string, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(integration.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IntegrationQuery) IDsX(ctx context.Context) []string {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IntegrationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*IntegrationQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IntegrationQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IntegrationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IntegrationQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IntegrationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IntegrationQuery) Clone() *IntegrationQuery {
	if iq == nil {
		return nil
	}
	return &IntegrationQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]integration.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Integration{}, iq.predicates...),
		withOwner:  iq.withOwner.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithOwner(opts ...func(*OrganizationQuery)) *IntegrationQuery {
	query := (&OrganizationClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withOwner = query
	return iq
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
//	client.Integration.Query().
//		GroupBy(integration.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (iq *IntegrationQuery) GroupBy(field string, fields ...string) *IntegrationGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IntegrationGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = integration.Label
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
//	client.Integration.Query().
//		Select(integration.FieldCreatedAt).
//		Scan(ctx, &v)
func (iq *IntegrationQuery) Select(fields ...string) *IntegrationSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &IntegrationSelect{IntegrationQuery: iq}
	sbuild.label = integration.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IntegrationSelect configured with the given aggregations.
func (iq *IntegrationQuery) Aggregate(fns ...AggregateFunc) *IntegrationSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *IntegrationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !integration.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	if integration.Policy == nil {
		return errors.New("generated: uninitialized integration.Policy (forgotten import generated/runtime?)")
	}
	if err := integration.Policy.EvalQuery(ctx, iq); err != nil {
		return err
	}
	return nil
}

func (iq *IntegrationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Integration, error) {
	var (
		nodes       = []*Integration{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Integration).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Integration{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = iq.schemaConfig.Integration
	ctx = internal.NewSchemaConfigContext(ctx, iq.schemaConfig)
	if len(iq.modifiers) > 0 {
		_spec.Modifiers = iq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withOwner; query != nil {
		if err := iq.loadOwner(ctx, query, nodes, nil,
			func(n *Integration, e *Organization) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	for i := range iq.loadTotal {
		if err := iq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *IntegrationQuery) loadOwner(ctx context.Context, query *OrganizationQuery, nodes []*Integration, init func(*Integration), assign func(*Integration, *Organization)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Integration)
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

func (iq *IntegrationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Schema = iq.schemaConfig.Integration
	ctx = internal.NewSchemaConfigContext(ctx, iq.schemaConfig)
	if len(iq.modifiers) > 0 {
		_spec.Modifiers = iq.modifiers
	}
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IntegrationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(integration.Table, integration.Columns, sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, integration.FieldID)
		for i := range fields {
			if fields[i] != integration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if iq.withOwner != nil {
			_spec.Node.AddColumnOnce(integration.FieldOwnerID)
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IntegrationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(integration.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = integration.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(iq.schemaConfig.Integration)
	ctx = internal.NewSchemaConfigContext(ctx, iq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IntegrationGroupBy is the group-by builder for Integration entities.
type IntegrationGroupBy struct {
	selector
	build *IntegrationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IntegrationGroupBy) Aggregate(fns ...AggregateFunc) *IntegrationGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *IntegrationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IntegrationQuery, *IntegrationGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *IntegrationGroupBy) sqlScan(ctx context.Context, root *IntegrationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IntegrationSelect is the builder for selecting fields of Integration entities.
type IntegrationSelect struct {
	*IntegrationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *IntegrationSelect) Aggregate(fns ...AggregateFunc) *IntegrationSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *IntegrationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IntegrationQuery, *IntegrationSelect](ctx, is.IntegrationQuery, is, is.inters, v)
}

func (is *IntegrationSelect) sqlScan(ctx context.Context, root *IntegrationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
