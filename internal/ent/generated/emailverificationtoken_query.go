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
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EmailVerificationTokenQuery is the builder for querying EmailVerificationToken entities.
type EmailVerificationTokenQuery struct {
	config
	ctx        *QueryContext
	order      []emailverificationtoken.OrderOption
	inters     []Interceptor
	predicates []predicate.EmailVerificationToken
	withOwner  *UserQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*EmailVerificationToken) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailVerificationTokenQuery builder.
func (evtq *EmailVerificationTokenQuery) Where(ps ...predicate.EmailVerificationToken) *EmailVerificationTokenQuery {
	evtq.predicates = append(evtq.predicates, ps...)
	return evtq
}

// Limit the number of records to be returned by this query.
func (evtq *EmailVerificationTokenQuery) Limit(limit int) *EmailVerificationTokenQuery {
	evtq.ctx.Limit = &limit
	return evtq
}

// Offset to start from.
func (evtq *EmailVerificationTokenQuery) Offset(offset int) *EmailVerificationTokenQuery {
	evtq.ctx.Offset = &offset
	return evtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (evtq *EmailVerificationTokenQuery) Unique(unique bool) *EmailVerificationTokenQuery {
	evtq.ctx.Unique = &unique
	return evtq
}

// Order specifies how the records should be ordered.
func (evtq *EmailVerificationTokenQuery) Order(o ...emailverificationtoken.OrderOption) *EmailVerificationTokenQuery {
	evtq.order = append(evtq.order, o...)
	return evtq
}

// QueryOwner chains the current query on the "owner" edge.
func (evtq *EmailVerificationTokenQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: evtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := evtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := evtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(emailverificationtoken.Table, emailverificationtoken.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, emailverificationtoken.OwnerTable, emailverificationtoken.OwnerColumn),
		)
		schemaConfig := evtq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.EmailVerificationToken
		fromU = sqlgraph.SetNeighbors(evtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EmailVerificationToken entity from the query.
// Returns a *NotFoundError when no EmailVerificationToken was found.
func (evtq *EmailVerificationTokenQuery) First(ctx context.Context) (*EmailVerificationToken, error) {
	nodes, err := evtq.Limit(1).All(setContextOp(ctx, evtq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emailverificationtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) FirstX(ctx context.Context) *EmailVerificationToken {
	node, err := evtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailVerificationToken ID from the query.
// Returns a *NotFoundError when no EmailVerificationToken ID was found.
func (evtq *EmailVerificationTokenQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = evtq.Limit(1).IDs(setContextOp(ctx, evtq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emailverificationtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) FirstIDX(ctx context.Context) string {
	id, err := evtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailVerificationToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmailVerificationToken entity is found.
// Returns a *NotFoundError when no EmailVerificationToken entities are found.
func (evtq *EmailVerificationTokenQuery) Only(ctx context.Context) (*EmailVerificationToken, error) {
	nodes, err := evtq.Limit(2).All(setContextOp(ctx, evtq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emailverificationtoken.Label}
	default:
		return nil, &NotSingularError{emailverificationtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) OnlyX(ctx context.Context) *EmailVerificationToken {
	node, err := evtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailVerificationToken ID in the query.
// Returns a *NotSingularError when more than one EmailVerificationToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (evtq *EmailVerificationTokenQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = evtq.Limit(2).IDs(setContextOp(ctx, evtq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emailverificationtoken.Label}
	default:
		err = &NotSingularError{emailverificationtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) OnlyIDX(ctx context.Context) string {
	id, err := evtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailVerificationTokens.
func (evtq *EmailVerificationTokenQuery) All(ctx context.Context) ([]*EmailVerificationToken, error) {
	ctx = setContextOp(ctx, evtq.ctx, ent.OpQueryAll)
	if err := evtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmailVerificationToken, *EmailVerificationTokenQuery]()
	return withInterceptors[[]*EmailVerificationToken](ctx, evtq, qr, evtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) AllX(ctx context.Context) []*EmailVerificationToken {
	nodes, err := evtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailVerificationToken IDs.
func (evtq *EmailVerificationTokenQuery) IDs(ctx context.Context) (ids []string, err error) {
	if evtq.ctx.Unique == nil && evtq.path != nil {
		evtq.Unique(true)
	}
	ctx = setContextOp(ctx, evtq.ctx, ent.OpQueryIDs)
	if err = evtq.Select(emailverificationtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) IDsX(ctx context.Context) []string {
	ids, err := evtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (evtq *EmailVerificationTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, evtq.ctx, ent.OpQueryCount)
	if err := evtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, evtq, querierCount[*EmailVerificationTokenQuery](), evtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) CountX(ctx context.Context) int {
	count, err := evtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (evtq *EmailVerificationTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, evtq.ctx, ent.OpQueryExist)
	switch _, err := evtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (evtq *EmailVerificationTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := evtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailVerificationTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (evtq *EmailVerificationTokenQuery) Clone() *EmailVerificationTokenQuery {
	if evtq == nil {
		return nil
	}
	return &EmailVerificationTokenQuery{
		config:     evtq.config,
		ctx:        evtq.ctx.Clone(),
		order:      append([]emailverificationtoken.OrderOption{}, evtq.order...),
		inters:     append([]Interceptor{}, evtq.inters...),
		predicates: append([]predicate.EmailVerificationToken{}, evtq.predicates...),
		withOwner:  evtq.withOwner.Clone(),
		// clone intermediate query.
		sql:  evtq.sql.Clone(),
		path: evtq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (evtq *EmailVerificationTokenQuery) WithOwner(opts ...func(*UserQuery)) *EmailVerificationTokenQuery {
	query := (&UserClient{config: evtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	evtq.withOwner = query
	return evtq
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
//	client.EmailVerificationToken.Query().
//		GroupBy(emailverificationtoken.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (evtq *EmailVerificationTokenQuery) GroupBy(field string, fields ...string) *EmailVerificationTokenGroupBy {
	evtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmailVerificationTokenGroupBy{build: evtq}
	grbuild.flds = &evtq.ctx.Fields
	grbuild.label = emailverificationtoken.Label
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
//	client.EmailVerificationToken.Query().
//		Select(emailverificationtoken.FieldCreatedAt).
//		Scan(ctx, &v)
func (evtq *EmailVerificationTokenQuery) Select(fields ...string) *EmailVerificationTokenSelect {
	evtq.ctx.Fields = append(evtq.ctx.Fields, fields...)
	sbuild := &EmailVerificationTokenSelect{EmailVerificationTokenQuery: evtq}
	sbuild.label = emailverificationtoken.Label
	sbuild.flds, sbuild.scan = &evtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmailVerificationTokenSelect configured with the given aggregations.
func (evtq *EmailVerificationTokenQuery) Aggregate(fns ...AggregateFunc) *EmailVerificationTokenSelect {
	return evtq.Select().Aggregate(fns...)
}

func (evtq *EmailVerificationTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range evtq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, evtq); err != nil {
				return err
			}
		}
	}
	for _, f := range evtq.ctx.Fields {
		if !emailverificationtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if evtq.path != nil {
		prev, err := evtq.path(ctx)
		if err != nil {
			return err
		}
		evtq.sql = prev
	}
	if emailverificationtoken.Policy == nil {
		return errors.New("generated: uninitialized emailverificationtoken.Policy (forgotten import generated/runtime?)")
	}
	if err := emailverificationtoken.Policy.EvalQuery(ctx, evtq); err != nil {
		return err
	}
	return nil
}

func (evtq *EmailVerificationTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmailVerificationToken, error) {
	var (
		nodes       = []*EmailVerificationToken{}
		_spec       = evtq.querySpec()
		loadedTypes = [1]bool{
			evtq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmailVerificationToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmailVerificationToken{config: evtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = evtq.schemaConfig.EmailVerificationToken
	ctx = internal.NewSchemaConfigContext(ctx, evtq.schemaConfig)
	if len(evtq.modifiers) > 0 {
		_spec.Modifiers = evtq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, evtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := evtq.withOwner; query != nil {
		if err := evtq.loadOwner(ctx, query, nodes, nil,
			func(n *EmailVerificationToken, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	for i := range evtq.loadTotal {
		if err := evtq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (evtq *EmailVerificationTokenQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*EmailVerificationToken, init func(*EmailVerificationToken), assign func(*EmailVerificationToken, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*EmailVerificationToken)
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
	query.Where(user.IDIn(ids...))
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

func (evtq *EmailVerificationTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := evtq.querySpec()
	_spec.Node.Schema = evtq.schemaConfig.EmailVerificationToken
	ctx = internal.NewSchemaConfigContext(ctx, evtq.schemaConfig)
	if len(evtq.modifiers) > 0 {
		_spec.Modifiers = evtq.modifiers
	}
	_spec.Node.Columns = evtq.ctx.Fields
	if len(evtq.ctx.Fields) > 0 {
		_spec.Unique = evtq.ctx.Unique != nil && *evtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, evtq.driver, _spec)
}

func (evtq *EmailVerificationTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(emailverificationtoken.Table, emailverificationtoken.Columns, sqlgraph.NewFieldSpec(emailverificationtoken.FieldID, field.TypeString))
	_spec.From = evtq.sql
	if unique := evtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if evtq.path != nil {
		_spec.Unique = true
	}
	if fields := evtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailverificationtoken.FieldID)
		for i := range fields {
			if fields[i] != emailverificationtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if evtq.withOwner != nil {
			_spec.Node.AddColumnOnce(emailverificationtoken.FieldOwnerID)
		}
	}
	if ps := evtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := evtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := evtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := evtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (evtq *EmailVerificationTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(evtq.driver.Dialect())
	t1 := builder.Table(emailverificationtoken.Table)
	columns := evtq.ctx.Fields
	if len(columns) == 0 {
		columns = emailverificationtoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if evtq.sql != nil {
		selector = evtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if evtq.ctx.Unique != nil && *evtq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(evtq.schemaConfig.EmailVerificationToken)
	ctx = internal.NewSchemaConfigContext(ctx, evtq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range evtq.predicates {
		p(selector)
	}
	for _, p := range evtq.order {
		p(selector)
	}
	if offset := evtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := evtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmailVerificationTokenGroupBy is the group-by builder for EmailVerificationToken entities.
type EmailVerificationTokenGroupBy struct {
	selector
	build *EmailVerificationTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (evtgb *EmailVerificationTokenGroupBy) Aggregate(fns ...AggregateFunc) *EmailVerificationTokenGroupBy {
	evtgb.fns = append(evtgb.fns, fns...)
	return evtgb
}

// Scan applies the selector query and scans the result into the given value.
func (evtgb *EmailVerificationTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, evtgb.build.ctx, ent.OpQueryGroupBy)
	if err := evtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailVerificationTokenQuery, *EmailVerificationTokenGroupBy](ctx, evtgb.build, evtgb, evtgb.build.inters, v)
}

func (evtgb *EmailVerificationTokenGroupBy) sqlScan(ctx context.Context, root *EmailVerificationTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(evtgb.fns))
	for _, fn := range evtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*evtgb.flds)+len(evtgb.fns))
		for _, f := range *evtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*evtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := evtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmailVerificationTokenSelect is the builder for selecting fields of EmailVerificationToken entities.
type EmailVerificationTokenSelect struct {
	*EmailVerificationTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (evts *EmailVerificationTokenSelect) Aggregate(fns ...AggregateFunc) *EmailVerificationTokenSelect {
	evts.fns = append(evts.fns, fns...)
	return evts
}

// Scan applies the selector query and scans the result into the given value.
func (evts *EmailVerificationTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, evts.ctx, ent.OpQuerySelect)
	if err := evts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailVerificationTokenQuery, *EmailVerificationTokenSelect](ctx, evts.EmailVerificationTokenQuery, evts, evts.inters, v)
}

func (evts *EmailVerificationTokenSelect) sqlScan(ctx context.Context, root *EmailVerificationTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(evts.fns))
	for _, fn := range evts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*evts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := evts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
