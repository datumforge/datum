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
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OrganizationSettingQuery is the builder for querying OrganizationSetting entities.
type OrganizationSettingQuery struct {
	config
	ctx              *QueryContext
	order            []organizationsetting.OrderOption
	inters           []Interceptor
	predicates       []predicate.OrganizationSetting
	withOrganization *OrganizationQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*OrganizationSetting) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationSettingQuery builder.
func (osq *OrganizationSettingQuery) Where(ps ...predicate.OrganizationSetting) *OrganizationSettingQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit the number of records to be returned by this query.
func (osq *OrganizationSettingQuery) Limit(limit int) *OrganizationSettingQuery {
	osq.ctx.Limit = &limit
	return osq
}

// Offset to start from.
func (osq *OrganizationSettingQuery) Offset(offset int) *OrganizationSettingQuery {
	osq.ctx.Offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OrganizationSettingQuery) Unique(unique bool) *OrganizationSettingQuery {
	osq.ctx.Unique = &unique
	return osq
}

// Order specifies how the records should be ordered.
func (osq *OrganizationSettingQuery) Order(o ...organizationsetting.OrderOption) *OrganizationSettingQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// QueryOrganization chains the current query on the "organization" edge.
func (osq *OrganizationSettingQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationsetting.Table, organizationsetting.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, organizationsetting.OrganizationTable, organizationsetting.OrganizationColumn),
		)
		schemaConfig := osq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.OrganizationSetting
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrganizationSetting entity from the query.
// Returns a *NotFoundError when no OrganizationSetting was found.
func (osq *OrganizationSettingQuery) First(ctx context.Context) (*OrganizationSetting, error) {
	nodes, err := osq.Limit(1).All(setContextOp(ctx, osq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OrganizationSettingQuery) FirstX(ctx context.Context) *OrganizationSetting {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationSetting ID from the query.
// Returns a *NotFoundError when no OrganizationSetting ID was found.
func (osq *OrganizationSettingQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = osq.Limit(1).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OrganizationSettingQuery) FirstIDX(ctx context.Context) string {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationSetting entity is found.
// Returns a *NotFoundError when no OrganizationSetting entities are found.
func (osq *OrganizationSettingQuery) Only(ctx context.Context) (*OrganizationSetting, error) {
	nodes, err := osq.Limit(2).All(setContextOp(ctx, osq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationsetting.Label}
	default:
		return nil, &NotSingularError{organizationsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OrganizationSettingQuery) OnlyX(ctx context.Context) *OrganizationSetting {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationSetting ID in the query.
// Returns a *NotSingularError when more than one OrganizationSetting ID is found.
// Returns a *NotFoundError when no entities are found.
func (osq *OrganizationSettingQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = osq.Limit(2).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationsetting.Label}
	default:
		err = &NotSingularError{organizationsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OrganizationSettingQuery) OnlyIDX(ctx context.Context) string {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationSettings.
func (osq *OrganizationSettingQuery) All(ctx context.Context) ([]*OrganizationSetting, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryAll)
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrganizationSetting, *OrganizationSettingQuery]()
	return withInterceptors[[]*OrganizationSetting](ctx, osq, qr, osq.inters)
}

// AllX is like All, but panics if an error occurs.
func (osq *OrganizationSettingQuery) AllX(ctx context.Context) []*OrganizationSetting {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationSetting IDs.
func (osq *OrganizationSettingQuery) IDs(ctx context.Context) (ids []string, err error) {
	if osq.ctx.Unique == nil && osq.path != nil {
		osq.Unique(true)
	}
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryIDs)
	if err = osq.Select(organizationsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OrganizationSettingQuery) IDsX(ctx context.Context) []string {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OrganizationSettingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryCount)
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, osq, querierCount[*OrganizationSettingQuery](), osq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OrganizationSettingQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OrganizationSettingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryExist)
	switch _, err := osq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OrganizationSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OrganizationSettingQuery) Clone() *OrganizationSettingQuery {
	if osq == nil {
		return nil
	}
	return &OrganizationSettingQuery{
		config:           osq.config,
		ctx:              osq.ctx.Clone(),
		order:            append([]organizationsetting.OrderOption{}, osq.order...),
		inters:           append([]Interceptor{}, osq.inters...),
		predicates:       append([]predicate.OrganizationSetting{}, osq.predicates...),
		withOrganization: osq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  osq.sql.Clone(),
		path: osq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrganizationSettingQuery) WithOrganization(opts ...func(*OrganizationQuery)) *OrganizationSettingQuery {
	query := (&OrganizationClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withOrganization = query
	return osq
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
//	client.OrganizationSetting.Query().
//		GroupBy(organizationsetting.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (osq *OrganizationSettingQuery) GroupBy(field string, fields ...string) *OrganizationSettingGroupBy {
	osq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationSettingGroupBy{build: osq}
	grbuild.flds = &osq.ctx.Fields
	grbuild.label = organizationsetting.Label
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
//	client.OrganizationSetting.Query().
//		Select(organizationsetting.FieldCreatedAt).
//		Scan(ctx, &v)
func (osq *OrganizationSettingQuery) Select(fields ...string) *OrganizationSettingSelect {
	osq.ctx.Fields = append(osq.ctx.Fields, fields...)
	sbuild := &OrganizationSettingSelect{OrganizationSettingQuery: osq}
	sbuild.label = organizationsetting.Label
	sbuild.flds, sbuild.scan = &osq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationSettingSelect configured with the given aggregations.
func (osq *OrganizationSettingQuery) Aggregate(fns ...AggregateFunc) *OrganizationSettingSelect {
	return osq.Select().Aggregate(fns...)
}

func (osq *OrganizationSettingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range osq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, osq); err != nil {
				return err
			}
		}
	}
	for _, f := range osq.ctx.Fields {
		if !organizationsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	if organizationsetting.Policy == nil {
		return errors.New("generated: uninitialized organizationsetting.Policy (forgotten import generated/runtime?)")
	}
	if err := organizationsetting.Policy.EvalQuery(ctx, osq); err != nil {
		return err
	}
	return nil
}

func (osq *OrganizationSettingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationSetting, error) {
	var (
		nodes       = []*OrganizationSetting{}
		_spec       = osq.querySpec()
		loadedTypes = [1]bool{
			osq.withOrganization != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrganizationSetting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrganizationSetting{config: osq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = osq.schemaConfig.OrganizationSetting
	ctx = internal.NewSchemaConfigContext(ctx, osq.schemaConfig)
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := osq.withOrganization; query != nil {
		if err := osq.loadOrganization(ctx, query, nodes, nil,
			func(n *OrganizationSetting, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	for i := range osq.loadTotal {
		if err := osq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (osq *OrganizationSettingQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*OrganizationSetting, init func(*OrganizationSetting), assign func(*OrganizationSetting, *Organization)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrganizationSetting)
	for i := range nodes {
		fk := nodes[i].OrganizationID
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
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (osq *OrganizationSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	_spec.Node.Schema = osq.schemaConfig.OrganizationSetting
	ctx = internal.NewSchemaConfigContext(ctx, osq.schemaConfig)
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	_spec.Node.Columns = osq.ctx.Fields
	if len(osq.ctx.Fields) > 0 {
		_spec.Unique = osq.ctx.Unique != nil && *osq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OrganizationSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organizationsetting.Table, organizationsetting.Columns, sqlgraph.NewFieldSpec(organizationsetting.FieldID, field.TypeString))
	_spec.From = osq.sql
	if unique := osq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if osq.path != nil {
		_spec.Unique = true
	}
	if fields := osq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationsetting.FieldID)
		for i := range fields {
			if fields[i] != organizationsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if osq.withOrganization != nil {
			_spec.Node.AddColumnOnce(organizationsetting.FieldOrganizationID)
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OrganizationSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(organizationsetting.Table)
	columns := osq.ctx.Fields
	if len(columns) == 0 {
		columns = organizationsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osq.ctx.Unique != nil && *osq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(osq.schemaConfig.OrganizationSetting)
	ctx = internal.NewSchemaConfigContext(ctx, osq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrganizationSettingGroupBy is the group-by builder for OrganizationSetting entities.
type OrganizationSettingGroupBy struct {
	selector
	build *OrganizationSettingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OrganizationSettingGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationSettingGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the selector query and scans the result into the given value.
func (osgb *OrganizationSettingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, osgb.build.ctx, ent.OpQueryGroupBy)
	if err := osgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationSettingQuery, *OrganizationSettingGroupBy](ctx, osgb.build, osgb, osgb.build.inters, v)
}

func (osgb *OrganizationSettingGroupBy) sqlScan(ctx context.Context, root *OrganizationSettingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*osgb.flds)+len(osgb.fns))
		for _, f := range *osgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*osgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationSettingSelect is the builder for selecting fields of OrganizationSetting entities.
type OrganizationSettingSelect struct {
	*OrganizationSettingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oss *OrganizationSettingSelect) Aggregate(fns ...AggregateFunc) *OrganizationSettingSelect {
	oss.fns = append(oss.fns, fns...)
	return oss
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OrganizationSettingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oss.ctx, ent.OpQuerySelect)
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationSettingQuery, *OrganizationSettingSelect](ctx, oss.OrganizationSettingQuery, oss, oss.inters, v)
}

func (oss *OrganizationSettingSelect) sqlScan(ctx context.Context, root *OrganizationSettingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oss.fns))
	for _, fn := range oss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
