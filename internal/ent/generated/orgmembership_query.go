// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OrgMembershipQuery is the builder for querying OrgMembership entities.
type OrgMembershipQuery struct {
	config
	ctx              *QueryContext
	order            []orgmembership.OrderOption
	inters           []Interceptor
	predicates       []predicate.OrgMembership
	withOrganization *OrganizationQuery
	withUser         *UserQuery
	withEvents       *EventQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*OrgMembership) error
	withNamedEvents  map[string]*EventQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgMembershipQuery builder.
func (omq *OrgMembershipQuery) Where(ps ...predicate.OrgMembership) *OrgMembershipQuery {
	omq.predicates = append(omq.predicates, ps...)
	return omq
}

// Limit the number of records to be returned by this query.
func (omq *OrgMembershipQuery) Limit(limit int) *OrgMembershipQuery {
	omq.ctx.Limit = &limit
	return omq
}

// Offset to start from.
func (omq *OrgMembershipQuery) Offset(offset int) *OrgMembershipQuery {
	omq.ctx.Offset = &offset
	return omq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (omq *OrgMembershipQuery) Unique(unique bool) *OrgMembershipQuery {
	omq.ctx.Unique = &unique
	return omq
}

// Order specifies how the records should be ordered.
func (omq *OrgMembershipQuery) Order(o ...orgmembership.OrderOption) *OrgMembershipQuery {
	omq.order = append(omq.order, o...)
	return omq
}

// QueryOrganization chains the current query on the "organization" edge.
func (omq *OrgMembershipQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: omq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgmembership.Table, orgmembership.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgmembership.OrganizationTable, orgmembership.OrganizationColumn),
		)
		schemaConfig := omq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.OrgMembership
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (omq *OrgMembershipQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: omq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgmembership.Table, orgmembership.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgmembership.UserTable, orgmembership.UserColumn),
		)
		schemaConfig := omq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.OrgMembership
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (omq *OrgMembershipQuery) QueryEvents() *EventQuery {
	query := (&EventClient{config: omq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := omq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := omq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgmembership.Table, orgmembership.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, orgmembership.EventsTable, orgmembership.EventsPrimaryKey...),
		)
		schemaConfig := omq.schemaConfig
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.OrgMembershipEvents
		fromU = sqlgraph.SetNeighbors(omq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgMembership entity from the query.
// Returns a *NotFoundError when no OrgMembership was found.
func (omq *OrgMembershipQuery) First(ctx context.Context) (*OrgMembership, error) {
	nodes, err := omq.Limit(1).All(setContextOp(ctx, omq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orgmembership.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (omq *OrgMembershipQuery) FirstX(ctx context.Context) *OrgMembership {
	node, err := omq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgMembership ID from the query.
// Returns a *NotFoundError when no OrgMembership ID was found.
func (omq *OrgMembershipQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = omq.Limit(1).IDs(setContextOp(ctx, omq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orgmembership.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (omq *OrgMembershipQuery) FirstIDX(ctx context.Context) string {
	id, err := omq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgMembership entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrgMembership entity is found.
// Returns a *NotFoundError when no OrgMembership entities are found.
func (omq *OrgMembershipQuery) Only(ctx context.Context) (*OrgMembership, error) {
	nodes, err := omq.Limit(2).All(setContextOp(ctx, omq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orgmembership.Label}
	default:
		return nil, &NotSingularError{orgmembership.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (omq *OrgMembershipQuery) OnlyX(ctx context.Context) *OrgMembership {
	node, err := omq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgMembership ID in the query.
// Returns a *NotSingularError when more than one OrgMembership ID is found.
// Returns a *NotFoundError when no entities are found.
func (omq *OrgMembershipQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = omq.Limit(2).IDs(setContextOp(ctx, omq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orgmembership.Label}
	default:
		err = &NotSingularError{orgmembership.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (omq *OrgMembershipQuery) OnlyIDX(ctx context.Context) string {
	id, err := omq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgMemberships.
func (omq *OrgMembershipQuery) All(ctx context.Context) ([]*OrgMembership, error) {
	ctx = setContextOp(ctx, omq.ctx, ent.OpQueryAll)
	if err := omq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrgMembership, *OrgMembershipQuery]()
	return withInterceptors[[]*OrgMembership](ctx, omq, qr, omq.inters)
}

// AllX is like All, but panics if an error occurs.
func (omq *OrgMembershipQuery) AllX(ctx context.Context) []*OrgMembership {
	nodes, err := omq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgMembership IDs.
func (omq *OrgMembershipQuery) IDs(ctx context.Context) (ids []string, err error) {
	if omq.ctx.Unique == nil && omq.path != nil {
		omq.Unique(true)
	}
	ctx = setContextOp(ctx, omq.ctx, ent.OpQueryIDs)
	if err = omq.Select(orgmembership.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (omq *OrgMembershipQuery) IDsX(ctx context.Context) []string {
	ids, err := omq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (omq *OrgMembershipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, omq.ctx, ent.OpQueryCount)
	if err := omq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, omq, querierCount[*OrgMembershipQuery](), omq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (omq *OrgMembershipQuery) CountX(ctx context.Context) int {
	count, err := omq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (omq *OrgMembershipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, omq.ctx, ent.OpQueryExist)
	switch _, err := omq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (omq *OrgMembershipQuery) ExistX(ctx context.Context) bool {
	exist, err := omq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgMembershipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (omq *OrgMembershipQuery) Clone() *OrgMembershipQuery {
	if omq == nil {
		return nil
	}
	return &OrgMembershipQuery{
		config:           omq.config,
		ctx:              omq.ctx.Clone(),
		order:            append([]orgmembership.OrderOption{}, omq.order...),
		inters:           append([]Interceptor{}, omq.inters...),
		predicates:       append([]predicate.OrgMembership{}, omq.predicates...),
		withOrganization: omq.withOrganization.Clone(),
		withUser:         omq.withUser.Clone(),
		withEvents:       omq.withEvents.Clone(),
		// clone intermediate query.
		sql:  omq.sql.Clone(),
		path: omq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OrgMembershipQuery) WithOrganization(opts ...func(*OrganizationQuery)) *OrgMembershipQuery {
	query := (&OrganizationClient{config: omq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	omq.withOrganization = query
	return omq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OrgMembershipQuery) WithUser(opts ...func(*UserQuery)) *OrgMembershipQuery {
	query := (&UserClient{config: omq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	omq.withUser = query
	return omq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (omq *OrgMembershipQuery) WithEvents(opts ...func(*EventQuery)) *OrgMembershipQuery {
	query := (&EventClient{config: omq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	omq.withEvents = query
	return omq
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
//	client.OrgMembership.Query().
//		GroupBy(orgmembership.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (omq *OrgMembershipQuery) GroupBy(field string, fields ...string) *OrgMembershipGroupBy {
	omq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrgMembershipGroupBy{build: omq}
	grbuild.flds = &omq.ctx.Fields
	grbuild.label = orgmembership.Label
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
//	client.OrgMembership.Query().
//		Select(orgmembership.FieldCreatedAt).
//		Scan(ctx, &v)
func (omq *OrgMembershipQuery) Select(fields ...string) *OrgMembershipSelect {
	omq.ctx.Fields = append(omq.ctx.Fields, fields...)
	sbuild := &OrgMembershipSelect{OrgMembershipQuery: omq}
	sbuild.label = orgmembership.Label
	sbuild.flds, sbuild.scan = &omq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrgMembershipSelect configured with the given aggregations.
func (omq *OrgMembershipQuery) Aggregate(fns ...AggregateFunc) *OrgMembershipSelect {
	return omq.Select().Aggregate(fns...)
}

func (omq *OrgMembershipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range omq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, omq); err != nil {
				return err
			}
		}
	}
	for _, f := range omq.ctx.Fields {
		if !orgmembership.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if omq.path != nil {
		prev, err := omq.path(ctx)
		if err != nil {
			return err
		}
		omq.sql = prev
	}
	if orgmembership.Policy == nil {
		return errors.New("generated: uninitialized orgmembership.Policy (forgotten import generated/runtime?)")
	}
	if err := orgmembership.Policy.EvalQuery(ctx, omq); err != nil {
		return err
	}
	return nil
}

func (omq *OrgMembershipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrgMembership, error) {
	var (
		nodes       = []*OrgMembership{}
		_spec       = omq.querySpec()
		loadedTypes = [3]bool{
			omq.withOrganization != nil,
			omq.withUser != nil,
			omq.withEvents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrgMembership).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrgMembership{config: omq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = omq.schemaConfig.OrgMembership
	ctx = internal.NewSchemaConfigContext(ctx, omq.schemaConfig)
	if len(omq.modifiers) > 0 {
		_spec.Modifiers = omq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, omq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := omq.withOrganization; query != nil {
		if err := omq.loadOrganization(ctx, query, nodes, nil,
			func(n *OrgMembership, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := omq.withUser; query != nil {
		if err := omq.loadUser(ctx, query, nodes, nil,
			func(n *OrgMembership, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := omq.withEvents; query != nil {
		if err := omq.loadEvents(ctx, query, nodes,
			func(n *OrgMembership) { n.Edges.Events = []*Event{} },
			func(n *OrgMembership, e *Event) { n.Edges.Events = append(n.Edges.Events, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range omq.withNamedEvents {
		if err := omq.loadEvents(ctx, query, nodes,
			func(n *OrgMembership) { n.appendNamedEvents(name) },
			func(n *OrgMembership, e *Event) { n.appendNamedEvents(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range omq.loadTotal {
		if err := omq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (omq *OrgMembershipQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*OrgMembership, init func(*OrgMembership), assign func(*OrgMembership, *Organization)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrgMembership)
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
func (omq *OrgMembershipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*OrgMembership, init func(*OrgMembership), assign func(*OrgMembership, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrgMembership)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (omq *OrgMembershipQuery) loadEvents(ctx context.Context, query *EventQuery, nodes []*OrgMembership, init func(*OrgMembership), assign func(*OrgMembership, *Event)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*OrgMembership)
	nids := make(map[string]map[*OrgMembership]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(orgmembership.EventsTable)
		joinT.Schema(omq.schemaConfig.OrgMembershipEvents)
		s.Join(joinT).On(s.C(event.FieldID), joinT.C(orgmembership.EventsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(orgmembership.EventsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(orgmembership.EventsPrimaryKey[0]))
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
					nids[inValue] = map[*OrgMembership]struct{}{byID[outValue]: {}}
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

func (omq *OrgMembershipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := omq.querySpec()
	_spec.Node.Schema = omq.schemaConfig.OrgMembership
	ctx = internal.NewSchemaConfigContext(ctx, omq.schemaConfig)
	if len(omq.modifiers) > 0 {
		_spec.Modifiers = omq.modifiers
	}
	_spec.Node.Columns = omq.ctx.Fields
	if len(omq.ctx.Fields) > 0 {
		_spec.Unique = omq.ctx.Unique != nil && *omq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, omq.driver, _spec)
}

func (omq *OrgMembershipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orgmembership.Table, orgmembership.Columns, sqlgraph.NewFieldSpec(orgmembership.FieldID, field.TypeString))
	_spec.From = omq.sql
	if unique := omq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if omq.path != nil {
		_spec.Unique = true
	}
	if fields := omq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgmembership.FieldID)
		for i := range fields {
			if fields[i] != orgmembership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if omq.withOrganization != nil {
			_spec.Node.AddColumnOnce(orgmembership.FieldOrganizationID)
		}
		if omq.withUser != nil {
			_spec.Node.AddColumnOnce(orgmembership.FieldUserID)
		}
	}
	if ps := omq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := omq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := omq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := omq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (omq *OrgMembershipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(omq.driver.Dialect())
	t1 := builder.Table(orgmembership.Table)
	columns := omq.ctx.Fields
	if len(columns) == 0 {
		columns = orgmembership.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if omq.sql != nil {
		selector = omq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if omq.ctx.Unique != nil && *omq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(omq.schemaConfig.OrgMembership)
	ctx = internal.NewSchemaConfigContext(ctx, omq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range omq.predicates {
		p(selector)
	}
	for _, p := range omq.order {
		p(selector)
	}
	if offset := omq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := omq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedEvents tells the query-builder to eager-load the nodes that are connected to the "events"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (omq *OrgMembershipQuery) WithNamedEvents(name string, opts ...func(*EventQuery)) *OrgMembershipQuery {
	query := (&EventClient{config: omq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if omq.withNamedEvents == nil {
		omq.withNamedEvents = make(map[string]*EventQuery)
	}
	omq.withNamedEvents[name] = query
	return omq
}

// OrgMembershipGroupBy is the group-by builder for OrgMembership entities.
type OrgMembershipGroupBy struct {
	selector
	build *OrgMembershipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (omgb *OrgMembershipGroupBy) Aggregate(fns ...AggregateFunc) *OrgMembershipGroupBy {
	omgb.fns = append(omgb.fns, fns...)
	return omgb
}

// Scan applies the selector query and scans the result into the given value.
func (omgb *OrgMembershipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, omgb.build.ctx, ent.OpQueryGroupBy)
	if err := omgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgMembershipQuery, *OrgMembershipGroupBy](ctx, omgb.build, omgb, omgb.build.inters, v)
}

func (omgb *OrgMembershipGroupBy) sqlScan(ctx context.Context, root *OrgMembershipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(omgb.fns))
	for _, fn := range omgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*omgb.flds)+len(omgb.fns))
		for _, f := range *omgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*omgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := omgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrgMembershipSelect is the builder for selecting fields of OrgMembership entities.
type OrgMembershipSelect struct {
	*OrgMembershipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oms *OrgMembershipSelect) Aggregate(fns ...AggregateFunc) *OrgMembershipSelect {
	oms.fns = append(oms.fns, fns...)
	return oms
}

// Scan applies the selector query and scans the result into the given value.
func (oms *OrgMembershipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oms.ctx, ent.OpQuerySelect)
	if err := oms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgMembershipQuery, *OrgMembershipSelect](ctx, oms.OrgMembershipQuery, oms, oms.inters, v)
}

func (oms *OrgMembershipSelect) sqlScan(ctx context.Context, root *OrgMembershipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oms.fns))
	for _, fn := range oms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
