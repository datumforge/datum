// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/entitlementplan"
	"github.com/datumforge/datum/internal/ent/generated/entitlementplanfeature"
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/feature"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntitlementPlanFeatureQuery is the builder for querying EntitlementPlanFeature entities.
type EntitlementPlanFeatureQuery struct {
	config
	ctx             *QueryContext
	order           []entitlementplanfeature.OrderOption
	inters          []Interceptor
	predicates      []predicate.EntitlementPlanFeature
	withOwner       *OrganizationQuery
	withPlan        *EntitlementPlanQuery
	withFeature     *FeatureQuery
	withEvents      *EventQuery
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*EntitlementPlanFeature) error
	withNamedEvents map[string]*EventQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntitlementPlanFeatureQuery builder.
func (epfq *EntitlementPlanFeatureQuery) Where(ps ...predicate.EntitlementPlanFeature) *EntitlementPlanFeatureQuery {
	epfq.predicates = append(epfq.predicates, ps...)
	return epfq
}

// Limit the number of records to be returned by this query.
func (epfq *EntitlementPlanFeatureQuery) Limit(limit int) *EntitlementPlanFeatureQuery {
	epfq.ctx.Limit = &limit
	return epfq
}

// Offset to start from.
func (epfq *EntitlementPlanFeatureQuery) Offset(offset int) *EntitlementPlanFeatureQuery {
	epfq.ctx.Offset = &offset
	return epfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (epfq *EntitlementPlanFeatureQuery) Unique(unique bool) *EntitlementPlanFeatureQuery {
	epfq.ctx.Unique = &unique
	return epfq
}

// Order specifies how the records should be ordered.
func (epfq *EntitlementPlanFeatureQuery) Order(o ...entitlementplanfeature.OrderOption) *EntitlementPlanFeatureQuery {
	epfq.order = append(epfq.order, o...)
	return epfq
}

// QueryOwner chains the current query on the "owner" edge.
func (epfq *EntitlementPlanFeatureQuery) QueryOwner() *OrganizationQuery {
	query := (&OrganizationClient{config: epfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := epfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := epfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entitlementplanfeature.Table, entitlementplanfeature.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entitlementplanfeature.OwnerTable, entitlementplanfeature.OwnerColumn),
		)
		schemaConfig := epfq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.EntitlementPlanFeature
		fromU = sqlgraph.SetNeighbors(epfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlan chains the current query on the "plan" edge.
func (epfq *EntitlementPlanFeatureQuery) QueryPlan() *EntitlementPlanQuery {
	query := (&EntitlementPlanClient{config: epfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := epfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := epfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entitlementplanfeature.Table, entitlementplanfeature.FieldID, selector),
			sqlgraph.To(entitlementplan.Table, entitlementplan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, entitlementplanfeature.PlanTable, entitlementplanfeature.PlanColumn),
		)
		schemaConfig := epfq.schemaConfig
		step.To.Schema = schemaConfig.EntitlementPlan
		step.Edge.Schema = schemaConfig.EntitlementPlanFeature
		fromU = sqlgraph.SetNeighbors(epfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeature chains the current query on the "feature" edge.
func (epfq *EntitlementPlanFeatureQuery) QueryFeature() *FeatureQuery {
	query := (&FeatureClient{config: epfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := epfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := epfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entitlementplanfeature.Table, entitlementplanfeature.FieldID, selector),
			sqlgraph.To(feature.Table, feature.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, entitlementplanfeature.FeatureTable, entitlementplanfeature.FeatureColumn),
		)
		schemaConfig := epfq.schemaConfig
		step.To.Schema = schemaConfig.Feature
		step.Edge.Schema = schemaConfig.EntitlementPlanFeature
		fromU = sqlgraph.SetNeighbors(epfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (epfq *EntitlementPlanFeatureQuery) QueryEvents() *EventQuery {
	query := (&EventClient{config: epfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := epfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := epfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entitlementplanfeature.Table, entitlementplanfeature.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, entitlementplanfeature.EventsTable, entitlementplanfeature.EventsPrimaryKey...),
		)
		schemaConfig := epfq.schemaConfig
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.EntitlementPlanFeatureEvents
		fromU = sqlgraph.SetNeighbors(epfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EntitlementPlanFeature entity from the query.
// Returns a *NotFoundError when no EntitlementPlanFeature was found.
func (epfq *EntitlementPlanFeatureQuery) First(ctx context.Context) (*EntitlementPlanFeature, error) {
	nodes, err := epfq.Limit(1).All(setContextOp(ctx, epfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entitlementplanfeature.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) FirstX(ctx context.Context) *EntitlementPlanFeature {
	node, err := epfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntitlementPlanFeature ID from the query.
// Returns a *NotFoundError when no EntitlementPlanFeature ID was found.
func (epfq *EntitlementPlanFeatureQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = epfq.Limit(1).IDs(setContextOp(ctx, epfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entitlementplanfeature.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) FirstIDX(ctx context.Context) string {
	id, err := epfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntitlementPlanFeature entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntitlementPlanFeature entity is found.
// Returns a *NotFoundError when no EntitlementPlanFeature entities are found.
func (epfq *EntitlementPlanFeatureQuery) Only(ctx context.Context) (*EntitlementPlanFeature, error) {
	nodes, err := epfq.Limit(2).All(setContextOp(ctx, epfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entitlementplanfeature.Label}
	default:
		return nil, &NotSingularError{entitlementplanfeature.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) OnlyX(ctx context.Context) *EntitlementPlanFeature {
	node, err := epfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntitlementPlanFeature ID in the query.
// Returns a *NotSingularError when more than one EntitlementPlanFeature ID is found.
// Returns a *NotFoundError when no entities are found.
func (epfq *EntitlementPlanFeatureQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = epfq.Limit(2).IDs(setContextOp(ctx, epfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entitlementplanfeature.Label}
	default:
		err = &NotSingularError{entitlementplanfeature.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) OnlyIDX(ctx context.Context) string {
	id, err := epfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntitlementPlanFeatures.
func (epfq *EntitlementPlanFeatureQuery) All(ctx context.Context) ([]*EntitlementPlanFeature, error) {
	ctx = setContextOp(ctx, epfq.ctx, "All")
	if err := epfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EntitlementPlanFeature, *EntitlementPlanFeatureQuery]()
	return withInterceptors[[]*EntitlementPlanFeature](ctx, epfq, qr, epfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) AllX(ctx context.Context) []*EntitlementPlanFeature {
	nodes, err := epfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntitlementPlanFeature IDs.
func (epfq *EntitlementPlanFeatureQuery) IDs(ctx context.Context) (ids []string, err error) {
	if epfq.ctx.Unique == nil && epfq.path != nil {
		epfq.Unique(true)
	}
	ctx = setContextOp(ctx, epfq.ctx, "IDs")
	if err = epfq.Select(entitlementplanfeature.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) IDsX(ctx context.Context) []string {
	ids, err := epfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (epfq *EntitlementPlanFeatureQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, epfq.ctx, "Count")
	if err := epfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, epfq, querierCount[*EntitlementPlanFeatureQuery](), epfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) CountX(ctx context.Context) int {
	count, err := epfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (epfq *EntitlementPlanFeatureQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, epfq.ctx, "Exist")
	switch _, err := epfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (epfq *EntitlementPlanFeatureQuery) ExistX(ctx context.Context) bool {
	exist, err := epfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntitlementPlanFeatureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (epfq *EntitlementPlanFeatureQuery) Clone() *EntitlementPlanFeatureQuery {
	if epfq == nil {
		return nil
	}
	return &EntitlementPlanFeatureQuery{
		config:      epfq.config,
		ctx:         epfq.ctx.Clone(),
		order:       append([]entitlementplanfeature.OrderOption{}, epfq.order...),
		inters:      append([]Interceptor{}, epfq.inters...),
		predicates:  append([]predicate.EntitlementPlanFeature{}, epfq.predicates...),
		withOwner:   epfq.withOwner.Clone(),
		withPlan:    epfq.withPlan.Clone(),
		withFeature: epfq.withFeature.Clone(),
		withEvents:  epfq.withEvents.Clone(),
		// clone intermediate query.
		sql:  epfq.sql.Clone(),
		path: epfq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (epfq *EntitlementPlanFeatureQuery) WithOwner(opts ...func(*OrganizationQuery)) *EntitlementPlanFeatureQuery {
	query := (&OrganizationClient{config: epfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	epfq.withOwner = query
	return epfq
}

// WithPlan tells the query-builder to eager-load the nodes that are connected to
// the "plan" edge. The optional arguments are used to configure the query builder of the edge.
func (epfq *EntitlementPlanFeatureQuery) WithPlan(opts ...func(*EntitlementPlanQuery)) *EntitlementPlanFeatureQuery {
	query := (&EntitlementPlanClient{config: epfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	epfq.withPlan = query
	return epfq
}

// WithFeature tells the query-builder to eager-load the nodes that are connected to
// the "feature" edge. The optional arguments are used to configure the query builder of the edge.
func (epfq *EntitlementPlanFeatureQuery) WithFeature(opts ...func(*FeatureQuery)) *EntitlementPlanFeatureQuery {
	query := (&FeatureClient{config: epfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	epfq.withFeature = query
	return epfq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (epfq *EntitlementPlanFeatureQuery) WithEvents(opts ...func(*EventQuery)) *EntitlementPlanFeatureQuery {
	query := (&EventClient{config: epfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	epfq.withEvents = query
	return epfq
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
//	client.EntitlementPlanFeature.Query().
//		GroupBy(entitlementplanfeature.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (epfq *EntitlementPlanFeatureQuery) GroupBy(field string, fields ...string) *EntitlementPlanFeatureGroupBy {
	epfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EntitlementPlanFeatureGroupBy{build: epfq}
	grbuild.flds = &epfq.ctx.Fields
	grbuild.label = entitlementplanfeature.Label
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
//	client.EntitlementPlanFeature.Query().
//		Select(entitlementplanfeature.FieldCreatedAt).
//		Scan(ctx, &v)
func (epfq *EntitlementPlanFeatureQuery) Select(fields ...string) *EntitlementPlanFeatureSelect {
	epfq.ctx.Fields = append(epfq.ctx.Fields, fields...)
	sbuild := &EntitlementPlanFeatureSelect{EntitlementPlanFeatureQuery: epfq}
	sbuild.label = entitlementplanfeature.Label
	sbuild.flds, sbuild.scan = &epfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EntitlementPlanFeatureSelect configured with the given aggregations.
func (epfq *EntitlementPlanFeatureQuery) Aggregate(fns ...AggregateFunc) *EntitlementPlanFeatureSelect {
	return epfq.Select().Aggregate(fns...)
}

func (epfq *EntitlementPlanFeatureQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range epfq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, epfq); err != nil {
				return err
			}
		}
	}
	for _, f := range epfq.ctx.Fields {
		if !entitlementplanfeature.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if epfq.path != nil {
		prev, err := epfq.path(ctx)
		if err != nil {
			return err
		}
		epfq.sql = prev
	}
	if entitlementplanfeature.Policy == nil {
		return errors.New("generated: uninitialized entitlementplanfeature.Policy (forgotten import generated/runtime?)")
	}
	if err := entitlementplanfeature.Policy.EvalQuery(ctx, epfq); err != nil {
		return err
	}
	return nil
}

func (epfq *EntitlementPlanFeatureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntitlementPlanFeature, error) {
	var (
		nodes       = []*EntitlementPlanFeature{}
		_spec       = epfq.querySpec()
		loadedTypes = [4]bool{
			epfq.withOwner != nil,
			epfq.withPlan != nil,
			epfq.withFeature != nil,
			epfq.withEvents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EntitlementPlanFeature).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EntitlementPlanFeature{config: epfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = epfq.schemaConfig.EntitlementPlanFeature
	ctx = internal.NewSchemaConfigContext(ctx, epfq.schemaConfig)
	if len(epfq.modifiers) > 0 {
		_spec.Modifiers = epfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, epfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := epfq.withOwner; query != nil {
		if err := epfq.loadOwner(ctx, query, nodes, nil,
			func(n *EntitlementPlanFeature, e *Organization) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := epfq.withPlan; query != nil {
		if err := epfq.loadPlan(ctx, query, nodes, nil,
			func(n *EntitlementPlanFeature, e *EntitlementPlan) { n.Edges.Plan = e }); err != nil {
			return nil, err
		}
	}
	if query := epfq.withFeature; query != nil {
		if err := epfq.loadFeature(ctx, query, nodes, nil,
			func(n *EntitlementPlanFeature, e *Feature) { n.Edges.Feature = e }); err != nil {
			return nil, err
		}
	}
	if query := epfq.withEvents; query != nil {
		if err := epfq.loadEvents(ctx, query, nodes,
			func(n *EntitlementPlanFeature) { n.Edges.Events = []*Event{} },
			func(n *EntitlementPlanFeature, e *Event) { n.Edges.Events = append(n.Edges.Events, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range epfq.withNamedEvents {
		if err := epfq.loadEvents(ctx, query, nodes,
			func(n *EntitlementPlanFeature) { n.appendNamedEvents(name) },
			func(n *EntitlementPlanFeature, e *Event) { n.appendNamedEvents(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range epfq.loadTotal {
		if err := epfq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (epfq *EntitlementPlanFeatureQuery) loadOwner(ctx context.Context, query *OrganizationQuery, nodes []*EntitlementPlanFeature, init func(*EntitlementPlanFeature), assign func(*EntitlementPlanFeature, *Organization)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*EntitlementPlanFeature)
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
func (epfq *EntitlementPlanFeatureQuery) loadPlan(ctx context.Context, query *EntitlementPlanQuery, nodes []*EntitlementPlanFeature, init func(*EntitlementPlanFeature), assign func(*EntitlementPlanFeature, *EntitlementPlan)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*EntitlementPlanFeature)
	for i := range nodes {
		fk := nodes[i].PlanID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(entitlementplan.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "plan_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (epfq *EntitlementPlanFeatureQuery) loadFeature(ctx context.Context, query *FeatureQuery, nodes []*EntitlementPlanFeature, init func(*EntitlementPlanFeature), assign func(*EntitlementPlanFeature, *Feature)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*EntitlementPlanFeature)
	for i := range nodes {
		fk := nodes[i].FeatureID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(feature.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "feature_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (epfq *EntitlementPlanFeatureQuery) loadEvents(ctx context.Context, query *EventQuery, nodes []*EntitlementPlanFeature, init func(*EntitlementPlanFeature), assign func(*EntitlementPlanFeature, *Event)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*EntitlementPlanFeature)
	nids := make(map[string]map[*EntitlementPlanFeature]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(entitlementplanfeature.EventsTable)
		joinT.Schema(epfq.schemaConfig.EntitlementPlanFeatureEvents)
		s.Join(joinT).On(s.C(event.FieldID), joinT.C(entitlementplanfeature.EventsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(entitlementplanfeature.EventsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(entitlementplanfeature.EventsPrimaryKey[0]))
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
					nids[inValue] = map[*EntitlementPlanFeature]struct{}{byID[outValue]: {}}
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

func (epfq *EntitlementPlanFeatureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := epfq.querySpec()
	_spec.Node.Schema = epfq.schemaConfig.EntitlementPlanFeature
	ctx = internal.NewSchemaConfigContext(ctx, epfq.schemaConfig)
	if len(epfq.modifiers) > 0 {
		_spec.Modifiers = epfq.modifiers
	}
	_spec.Node.Columns = epfq.ctx.Fields
	if len(epfq.ctx.Fields) > 0 {
		_spec.Unique = epfq.ctx.Unique != nil && *epfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, epfq.driver, _spec)
}

func (epfq *EntitlementPlanFeatureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entitlementplanfeature.Table, entitlementplanfeature.Columns, sqlgraph.NewFieldSpec(entitlementplanfeature.FieldID, field.TypeString))
	_spec.From = epfq.sql
	if unique := epfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if epfq.path != nil {
		_spec.Unique = true
	}
	if fields := epfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitlementplanfeature.FieldID)
		for i := range fields {
			if fields[i] != entitlementplanfeature.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if epfq.withOwner != nil {
			_spec.Node.AddColumnOnce(entitlementplanfeature.FieldOwnerID)
		}
		if epfq.withPlan != nil {
			_spec.Node.AddColumnOnce(entitlementplanfeature.FieldPlanID)
		}
		if epfq.withFeature != nil {
			_spec.Node.AddColumnOnce(entitlementplanfeature.FieldFeatureID)
		}
	}
	if ps := epfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := epfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := epfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := epfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (epfq *EntitlementPlanFeatureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(epfq.driver.Dialect())
	t1 := builder.Table(entitlementplanfeature.Table)
	columns := epfq.ctx.Fields
	if len(columns) == 0 {
		columns = entitlementplanfeature.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if epfq.sql != nil {
		selector = epfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if epfq.ctx.Unique != nil && *epfq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(epfq.schemaConfig.EntitlementPlanFeature)
	ctx = internal.NewSchemaConfigContext(ctx, epfq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range epfq.predicates {
		p(selector)
	}
	for _, p := range epfq.order {
		p(selector)
	}
	if offset := epfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := epfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedEvents tells the query-builder to eager-load the nodes that are connected to the "events"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (epfq *EntitlementPlanFeatureQuery) WithNamedEvents(name string, opts ...func(*EventQuery)) *EntitlementPlanFeatureQuery {
	query := (&EventClient{config: epfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if epfq.withNamedEvents == nil {
		epfq.withNamedEvents = make(map[string]*EventQuery)
	}
	epfq.withNamedEvents[name] = query
	return epfq
}

// EntitlementPlanFeatureGroupBy is the group-by builder for EntitlementPlanFeature entities.
type EntitlementPlanFeatureGroupBy struct {
	selector
	build *EntitlementPlanFeatureQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (epfgb *EntitlementPlanFeatureGroupBy) Aggregate(fns ...AggregateFunc) *EntitlementPlanFeatureGroupBy {
	epfgb.fns = append(epfgb.fns, fns...)
	return epfgb
}

// Scan applies the selector query and scans the result into the given value.
func (epfgb *EntitlementPlanFeatureGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, epfgb.build.ctx, "GroupBy")
	if err := epfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntitlementPlanFeatureQuery, *EntitlementPlanFeatureGroupBy](ctx, epfgb.build, epfgb, epfgb.build.inters, v)
}

func (epfgb *EntitlementPlanFeatureGroupBy) sqlScan(ctx context.Context, root *EntitlementPlanFeatureQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(epfgb.fns))
	for _, fn := range epfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*epfgb.flds)+len(epfgb.fns))
		for _, f := range *epfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*epfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := epfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EntitlementPlanFeatureSelect is the builder for selecting fields of EntitlementPlanFeature entities.
type EntitlementPlanFeatureSelect struct {
	*EntitlementPlanFeatureQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (epfs *EntitlementPlanFeatureSelect) Aggregate(fns ...AggregateFunc) *EntitlementPlanFeatureSelect {
	epfs.fns = append(epfs.fns, fns...)
	return epfs
}

// Scan applies the selector query and scans the result into the given value.
func (epfs *EntitlementPlanFeatureSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, epfs.ctx, "Select")
	if err := epfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntitlementPlanFeatureQuery, *EntitlementPlanFeatureSelect](ctx, epfs.EntitlementPlanFeatureQuery, epfs, epfs.inters, v)
}

func (epfs *EntitlementPlanFeatureSelect) sqlScan(ctx context.Context, root *EntitlementPlanFeatureQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(epfs.fns))
	for _, fn := range epfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*epfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := epfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
