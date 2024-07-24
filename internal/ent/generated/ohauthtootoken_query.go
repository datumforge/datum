// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootoken"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OhAuthTooTokenQuery is the builder for querying OhAuthTooToken entities.
type OhAuthTooTokenQuery struct {
	config
	ctx                  *QueryContext
	order                []ohauthtootoken.OrderOption
	inters               []Interceptor
	predicates           []predicate.OhAuthTooToken
	withIntegration      *IntegrationQuery
	withEvents           *EventQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*OhAuthTooToken) error
	withNamedIntegration map[string]*IntegrationQuery
	withNamedEvents      map[string]*EventQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OhAuthTooTokenQuery builder.
func (oattq *OhAuthTooTokenQuery) Where(ps ...predicate.OhAuthTooToken) *OhAuthTooTokenQuery {
	oattq.predicates = append(oattq.predicates, ps...)
	return oattq
}

// Limit the number of records to be returned by this query.
func (oattq *OhAuthTooTokenQuery) Limit(limit int) *OhAuthTooTokenQuery {
	oattq.ctx.Limit = &limit
	return oattq
}

// Offset to start from.
func (oattq *OhAuthTooTokenQuery) Offset(offset int) *OhAuthTooTokenQuery {
	oattq.ctx.Offset = &offset
	return oattq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oattq *OhAuthTooTokenQuery) Unique(unique bool) *OhAuthTooTokenQuery {
	oattq.ctx.Unique = &unique
	return oattq
}

// Order specifies how the records should be ordered.
func (oattq *OhAuthTooTokenQuery) Order(o ...ohauthtootoken.OrderOption) *OhAuthTooTokenQuery {
	oattq.order = append(oattq.order, o...)
	return oattq
}

// QueryIntegration chains the current query on the "integration" edge.
func (oattq *OhAuthTooTokenQuery) QueryIntegration() *IntegrationQuery {
	query := (&IntegrationClient{config: oattq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oattq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oattq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ohauthtootoken.Table, ohauthtootoken.FieldID, selector),
			sqlgraph.To(integration.Table, integration.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ohauthtootoken.IntegrationTable, ohauthtootoken.IntegrationPrimaryKey...),
		)
		schemaConfig := oattq.schemaConfig
		step.To.Schema = schemaConfig.Integration
		step.Edge.Schema = schemaConfig.IntegrationOauth2tokens
		fromU = sqlgraph.SetNeighbors(oattq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (oattq *OhAuthTooTokenQuery) QueryEvents() *EventQuery {
	query := (&EventClient{config: oattq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oattq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oattq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ohauthtootoken.Table, ohauthtootoken.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ohauthtootoken.EventsTable, ohauthtootoken.EventsPrimaryKey...),
		)
		schemaConfig := oattq.schemaConfig
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.OhAuthTooTokenEvents
		fromU = sqlgraph.SetNeighbors(oattq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OhAuthTooToken entity from the query.
// Returns a *NotFoundError when no OhAuthTooToken was found.
func (oattq *OhAuthTooTokenQuery) First(ctx context.Context) (*OhAuthTooToken, error) {
	nodes, err := oattq.Limit(1).All(setContextOp(ctx, oattq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ohauthtootoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) FirstX(ctx context.Context) *OhAuthTooToken {
	node, err := oattq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OhAuthTooToken ID from the query.
// Returns a *NotFoundError when no OhAuthTooToken ID was found.
func (oattq *OhAuthTooTokenQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oattq.Limit(1).IDs(setContextOp(ctx, oattq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ohauthtootoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) FirstIDX(ctx context.Context) string {
	id, err := oattq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OhAuthTooToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OhAuthTooToken entity is found.
// Returns a *NotFoundError when no OhAuthTooToken entities are found.
func (oattq *OhAuthTooTokenQuery) Only(ctx context.Context) (*OhAuthTooToken, error) {
	nodes, err := oattq.Limit(2).All(setContextOp(ctx, oattq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ohauthtootoken.Label}
	default:
		return nil, &NotSingularError{ohauthtootoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) OnlyX(ctx context.Context) *OhAuthTooToken {
	node, err := oattq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OhAuthTooToken ID in the query.
// Returns a *NotSingularError when more than one OhAuthTooToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (oattq *OhAuthTooTokenQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oattq.Limit(2).IDs(setContextOp(ctx, oattq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ohauthtootoken.Label}
	default:
		err = &NotSingularError{ohauthtootoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) OnlyIDX(ctx context.Context) string {
	id, err := oattq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OhAuthTooTokens.
func (oattq *OhAuthTooTokenQuery) All(ctx context.Context) ([]*OhAuthTooToken, error) {
	ctx = setContextOp(ctx, oattq.ctx, ent.OpQueryAll)
	if err := oattq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OhAuthTooToken, *OhAuthTooTokenQuery]()
	return withInterceptors[[]*OhAuthTooToken](ctx, oattq, qr, oattq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) AllX(ctx context.Context) []*OhAuthTooToken {
	nodes, err := oattq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OhAuthTooToken IDs.
func (oattq *OhAuthTooTokenQuery) IDs(ctx context.Context) (ids []string, err error) {
	if oattq.ctx.Unique == nil && oattq.path != nil {
		oattq.Unique(true)
	}
	ctx = setContextOp(ctx, oattq.ctx, ent.OpQueryIDs)
	if err = oattq.Select(ohauthtootoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) IDsX(ctx context.Context) []string {
	ids, err := oattq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oattq *OhAuthTooTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oattq.ctx, ent.OpQueryCount)
	if err := oattq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oattq, querierCount[*OhAuthTooTokenQuery](), oattq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) CountX(ctx context.Context) int {
	count, err := oattq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oattq *OhAuthTooTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oattq.ctx, ent.OpQueryExist)
	switch _, err := oattq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oattq *OhAuthTooTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := oattq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OhAuthTooTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oattq *OhAuthTooTokenQuery) Clone() *OhAuthTooTokenQuery {
	if oattq == nil {
		return nil
	}
	return &OhAuthTooTokenQuery{
		config:          oattq.config,
		ctx:             oattq.ctx.Clone(),
		order:           append([]ohauthtootoken.OrderOption{}, oattq.order...),
		inters:          append([]Interceptor{}, oattq.inters...),
		predicates:      append([]predicate.OhAuthTooToken{}, oattq.predicates...),
		withIntegration: oattq.withIntegration.Clone(),
		withEvents:      oattq.withEvents.Clone(),
		// clone intermediate query.
		sql:  oattq.sql.Clone(),
		path: oattq.path,
	}
}

// WithIntegration tells the query-builder to eager-load the nodes that are connected to
// the "integration" edge. The optional arguments are used to configure the query builder of the edge.
func (oattq *OhAuthTooTokenQuery) WithIntegration(opts ...func(*IntegrationQuery)) *OhAuthTooTokenQuery {
	query := (&IntegrationClient{config: oattq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oattq.withIntegration = query
	return oattq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (oattq *OhAuthTooTokenQuery) WithEvents(opts ...func(*EventQuery)) *OhAuthTooTokenQuery {
	query := (&EventClient{config: oattq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oattq.withEvents = query
	return oattq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MappingID string `json:"mapping_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OhAuthTooToken.Query().
//		GroupBy(ohauthtootoken.FieldMappingID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (oattq *OhAuthTooTokenQuery) GroupBy(field string, fields ...string) *OhAuthTooTokenGroupBy {
	oattq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OhAuthTooTokenGroupBy{build: oattq}
	grbuild.flds = &oattq.ctx.Fields
	grbuild.label = ohauthtootoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MappingID string `json:"mapping_id,omitempty"`
//	}
//
//	client.OhAuthTooToken.Query().
//		Select(ohauthtootoken.FieldMappingID).
//		Scan(ctx, &v)
func (oattq *OhAuthTooTokenQuery) Select(fields ...string) *OhAuthTooTokenSelect {
	oattq.ctx.Fields = append(oattq.ctx.Fields, fields...)
	sbuild := &OhAuthTooTokenSelect{OhAuthTooTokenQuery: oattq}
	sbuild.label = ohauthtootoken.Label
	sbuild.flds, sbuild.scan = &oattq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OhAuthTooTokenSelect configured with the given aggregations.
func (oattq *OhAuthTooTokenQuery) Aggregate(fns ...AggregateFunc) *OhAuthTooTokenSelect {
	return oattq.Select().Aggregate(fns...)
}

func (oattq *OhAuthTooTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oattq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oattq); err != nil {
				return err
			}
		}
	}
	for _, f := range oattq.ctx.Fields {
		if !ohauthtootoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if oattq.path != nil {
		prev, err := oattq.path(ctx)
		if err != nil {
			return err
		}
		oattq.sql = prev
	}
	return nil
}

func (oattq *OhAuthTooTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OhAuthTooToken, error) {
	var (
		nodes       = []*OhAuthTooToken{}
		_spec       = oattq.querySpec()
		loadedTypes = [2]bool{
			oattq.withIntegration != nil,
			oattq.withEvents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OhAuthTooToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OhAuthTooToken{config: oattq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = oattq.schemaConfig.OhAuthTooToken
	ctx = internal.NewSchemaConfigContext(ctx, oattq.schemaConfig)
	if len(oattq.modifiers) > 0 {
		_spec.Modifiers = oattq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oattq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oattq.withIntegration; query != nil {
		if err := oattq.loadIntegration(ctx, query, nodes,
			func(n *OhAuthTooToken) { n.Edges.Integration = []*Integration{} },
			func(n *OhAuthTooToken, e *Integration) { n.Edges.Integration = append(n.Edges.Integration, e) }); err != nil {
			return nil, err
		}
	}
	if query := oattq.withEvents; query != nil {
		if err := oattq.loadEvents(ctx, query, nodes,
			func(n *OhAuthTooToken) { n.Edges.Events = []*Event{} },
			func(n *OhAuthTooToken, e *Event) { n.Edges.Events = append(n.Edges.Events, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range oattq.withNamedIntegration {
		if err := oattq.loadIntegration(ctx, query, nodes,
			func(n *OhAuthTooToken) { n.appendNamedIntegration(name) },
			func(n *OhAuthTooToken, e *Integration) { n.appendNamedIntegration(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range oattq.withNamedEvents {
		if err := oattq.loadEvents(ctx, query, nodes,
			func(n *OhAuthTooToken) { n.appendNamedEvents(name) },
			func(n *OhAuthTooToken, e *Event) { n.appendNamedEvents(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range oattq.loadTotal {
		if err := oattq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oattq *OhAuthTooTokenQuery) loadIntegration(ctx context.Context, query *IntegrationQuery, nodes []*OhAuthTooToken, init func(*OhAuthTooToken), assign func(*OhAuthTooToken, *Integration)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*OhAuthTooToken)
	nids := make(map[string]map[*OhAuthTooToken]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(ohauthtootoken.IntegrationTable)
		joinT.Schema(oattq.schemaConfig.IntegrationOauth2tokens)
		s.Join(joinT).On(s.C(integration.FieldID), joinT.C(ohauthtootoken.IntegrationPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(ohauthtootoken.IntegrationPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(ohauthtootoken.IntegrationPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*OhAuthTooToken]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Integration](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "integration" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (oattq *OhAuthTooTokenQuery) loadEvents(ctx context.Context, query *EventQuery, nodes []*OhAuthTooToken, init func(*OhAuthTooToken), assign func(*OhAuthTooToken, *Event)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*OhAuthTooToken)
	nids := make(map[string]map[*OhAuthTooToken]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(ohauthtootoken.EventsTable)
		joinT.Schema(oattq.schemaConfig.OhAuthTooTokenEvents)
		s.Join(joinT).On(s.C(event.FieldID), joinT.C(ohauthtootoken.EventsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(ohauthtootoken.EventsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(ohauthtootoken.EventsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*OhAuthTooToken]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Event](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "events" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (oattq *OhAuthTooTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oattq.querySpec()
	_spec.Node.Schema = oattq.schemaConfig.OhAuthTooToken
	ctx = internal.NewSchemaConfigContext(ctx, oattq.schemaConfig)
	if len(oattq.modifiers) > 0 {
		_spec.Modifiers = oattq.modifiers
	}
	_spec.Node.Columns = oattq.ctx.Fields
	if len(oattq.ctx.Fields) > 0 {
		_spec.Unique = oattq.ctx.Unique != nil && *oattq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oattq.driver, _spec)
}

func (oattq *OhAuthTooTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ohauthtootoken.Table, ohauthtootoken.Columns, sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString))
	_spec.From = oattq.sql
	if unique := oattq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oattq.path != nil {
		_spec.Unique = true
	}
	if fields := oattq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ohauthtootoken.FieldID)
		for i := range fields {
			if fields[i] != ohauthtootoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oattq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oattq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oattq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oattq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oattq *OhAuthTooTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oattq.driver.Dialect())
	t1 := builder.Table(ohauthtootoken.Table)
	columns := oattq.ctx.Fields
	if len(columns) == 0 {
		columns = ohauthtootoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oattq.sql != nil {
		selector = oattq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oattq.ctx.Unique != nil && *oattq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(oattq.schemaConfig.OhAuthTooToken)
	ctx = internal.NewSchemaConfigContext(ctx, oattq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range oattq.predicates {
		p(selector)
	}
	for _, p := range oattq.order {
		p(selector)
	}
	if offset := oattq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oattq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedIntegration tells the query-builder to eager-load the nodes that are connected to the "integration"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (oattq *OhAuthTooTokenQuery) WithNamedIntegration(name string, opts ...func(*IntegrationQuery)) *OhAuthTooTokenQuery {
	query := (&IntegrationClient{config: oattq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if oattq.withNamedIntegration == nil {
		oattq.withNamedIntegration = make(map[string]*IntegrationQuery)
	}
	oattq.withNamedIntegration[name] = query
	return oattq
}

// WithNamedEvents tells the query-builder to eager-load the nodes that are connected to the "events"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (oattq *OhAuthTooTokenQuery) WithNamedEvents(name string, opts ...func(*EventQuery)) *OhAuthTooTokenQuery {
	query := (&EventClient{config: oattq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if oattq.withNamedEvents == nil {
		oattq.withNamedEvents = make(map[string]*EventQuery)
	}
	oattq.withNamedEvents[name] = query
	return oattq
}

// OhAuthTooTokenGroupBy is the group-by builder for OhAuthTooToken entities.
type OhAuthTooTokenGroupBy struct {
	selector
	build *OhAuthTooTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oattgb *OhAuthTooTokenGroupBy) Aggregate(fns ...AggregateFunc) *OhAuthTooTokenGroupBy {
	oattgb.fns = append(oattgb.fns, fns...)
	return oattgb
}

// Scan applies the selector query and scans the result into the given value.
func (oattgb *OhAuthTooTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oattgb.build.ctx, ent.OpQueryGroupBy)
	if err := oattgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OhAuthTooTokenQuery, *OhAuthTooTokenGroupBy](ctx, oattgb.build, oattgb, oattgb.build.inters, v)
}

func (oattgb *OhAuthTooTokenGroupBy) sqlScan(ctx context.Context, root *OhAuthTooTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oattgb.fns))
	for _, fn := range oattgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oattgb.flds)+len(oattgb.fns))
		for _, f := range *oattgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oattgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oattgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OhAuthTooTokenSelect is the builder for selecting fields of OhAuthTooToken entities.
type OhAuthTooTokenSelect struct {
	*OhAuthTooTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oatts *OhAuthTooTokenSelect) Aggregate(fns ...AggregateFunc) *OhAuthTooTokenSelect {
	oatts.fns = append(oatts.fns, fns...)
	return oatts
}

// Scan applies the selector query and scans the result into the given value.
func (oatts *OhAuthTooTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oatts.ctx, ent.OpQuerySelect)
	if err := oatts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OhAuthTooTokenQuery, *OhAuthTooTokenSelect](ctx, oatts.OhAuthTooTokenQuery, oatts, oatts.inters, v)
}

func (oatts *OhAuthTooTokenSelect) sqlScan(ctx context.Context, root *OhAuthTooTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oatts.fns))
	for _, fn := range oatts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oatts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oatts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
