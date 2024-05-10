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
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/hush"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootoken"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/webhook"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// IntegrationQuery is the builder for querying Integration entities.
type IntegrationQuery struct {
	config
	ctx                   *QueryContext
	order                 []integration.OrderOption
	inters                []Interceptor
	predicates            []predicate.Integration
	withOwner             *OrganizationQuery
	withSecrets           *HushQuery
	withOauth2tokens      *OhAuthTooTokenQuery
	withEvents            *EventQuery
	withWebhooks          *WebhookQuery
	withFKs               bool
	modifiers             []func(*sql.Selector)
	loadTotal             []func(context.Context, []*Integration) error
	withNamedSecrets      map[string]*HushQuery
	withNamedOauth2tokens map[string]*OhAuthTooTokenQuery
	withNamedEvents       map[string]*EventQuery
	withNamedWebhooks     map[string]*WebhookQuery
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

// QuerySecrets chains the current query on the "secrets" edge.
func (iq *IntegrationQuery) QuerySecrets() *HushQuery {
	query := (&HushClient{config: iq.config}).Query()
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
			sqlgraph.To(hush.Table, hush.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, integration.SecretsTable, integration.SecretsPrimaryKey...),
		)
		schemaConfig := iq.schemaConfig
		step.To.Schema = schemaConfig.Hush
		step.Edge.Schema = schemaConfig.IntegrationSecrets
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOauth2tokens chains the current query on the "oauth2tokens" edge.
func (iq *IntegrationQuery) QueryOauth2tokens() *OhAuthTooTokenQuery {
	query := (&OhAuthTooTokenClient{config: iq.config}).Query()
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
			sqlgraph.To(ohauthtootoken.Table, ohauthtootoken.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, integration.Oauth2tokensTable, integration.Oauth2tokensPrimaryKey...),
		)
		schemaConfig := iq.schemaConfig
		step.To.Schema = schemaConfig.OhAuthTooToken
		step.Edge.Schema = schemaConfig.IntegrationOauth2tokens
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (iq *IntegrationQuery) QueryEvents() *EventQuery {
	query := (&EventClient{config: iq.config}).Query()
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
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, integration.EventsTable, integration.EventsPrimaryKey...),
		)
		schemaConfig := iq.schemaConfig
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.IntegrationEvents
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWebhooks chains the current query on the "webhooks" edge.
func (iq *IntegrationQuery) QueryWebhooks() *WebhookQuery {
	query := (&WebhookClient{config: iq.config}).Query()
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
			sqlgraph.To(webhook.Table, webhook.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, integration.WebhooksTable, integration.WebhooksPrimaryKey...),
		)
		schemaConfig := iq.schemaConfig
		step.To.Schema = schemaConfig.Webhook
		step.Edge.Schema = schemaConfig.IntegrationWebhooks
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
		config:           iq.config,
		ctx:              iq.ctx.Clone(),
		order:            append([]integration.OrderOption{}, iq.order...),
		inters:           append([]Interceptor{}, iq.inters...),
		predicates:       append([]predicate.Integration{}, iq.predicates...),
		withOwner:        iq.withOwner.Clone(),
		withSecrets:      iq.withSecrets.Clone(),
		withOauth2tokens: iq.withOauth2tokens.Clone(),
		withEvents:       iq.withEvents.Clone(),
		withWebhooks:     iq.withWebhooks.Clone(),
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

// WithSecrets tells the query-builder to eager-load the nodes that are connected to
// the "secrets" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithSecrets(opts ...func(*HushQuery)) *IntegrationQuery {
	query := (&HushClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withSecrets = query
	return iq
}

// WithOauth2tokens tells the query-builder to eager-load the nodes that are connected to
// the "oauth2tokens" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithOauth2tokens(opts ...func(*OhAuthTooTokenQuery)) *IntegrationQuery {
	query := (&OhAuthTooTokenClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withOauth2tokens = query
	return iq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithEvents(opts ...func(*EventQuery)) *IntegrationQuery {
	query := (&EventClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withEvents = query
	return iq
}

// WithWebhooks tells the query-builder to eager-load the nodes that are connected to
// the "webhooks" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithWebhooks(opts ...func(*WebhookQuery)) *IntegrationQuery {
	query := (&WebhookClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withWebhooks = query
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
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [5]bool{
			iq.withOwner != nil,
			iq.withSecrets != nil,
			iq.withOauth2tokens != nil,
			iq.withEvents != nil,
			iq.withWebhooks != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, integration.ForeignKeys...)
	}
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
	if query := iq.withSecrets; query != nil {
		if err := iq.loadSecrets(ctx, query, nodes,
			func(n *Integration) { n.Edges.Secrets = []*Hush{} },
			func(n *Integration, e *Hush) { n.Edges.Secrets = append(n.Edges.Secrets, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withOauth2tokens; query != nil {
		if err := iq.loadOauth2tokens(ctx, query, nodes,
			func(n *Integration) { n.Edges.Oauth2tokens = []*OhAuthTooToken{} },
			func(n *Integration, e *OhAuthTooToken) { n.Edges.Oauth2tokens = append(n.Edges.Oauth2tokens, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withEvents; query != nil {
		if err := iq.loadEvents(ctx, query, nodes,
			func(n *Integration) { n.Edges.Events = []*Event{} },
			func(n *Integration, e *Event) { n.Edges.Events = append(n.Edges.Events, e) }); err != nil {
			return nil, err
		}
	}
	if query := iq.withWebhooks; query != nil {
		if err := iq.loadWebhooks(ctx, query, nodes,
			func(n *Integration) { n.Edges.Webhooks = []*Webhook{} },
			func(n *Integration, e *Webhook) { n.Edges.Webhooks = append(n.Edges.Webhooks, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range iq.withNamedSecrets {
		if err := iq.loadSecrets(ctx, query, nodes,
			func(n *Integration) { n.appendNamedSecrets(name) },
			func(n *Integration, e *Hush) { n.appendNamedSecrets(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range iq.withNamedOauth2tokens {
		if err := iq.loadOauth2tokens(ctx, query, nodes,
			func(n *Integration) { n.appendNamedOauth2tokens(name) },
			func(n *Integration, e *OhAuthTooToken) { n.appendNamedOauth2tokens(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range iq.withNamedEvents {
		if err := iq.loadEvents(ctx, query, nodes,
			func(n *Integration) { n.appendNamedEvents(name) },
			func(n *Integration, e *Event) { n.appendNamedEvents(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range iq.withNamedWebhooks {
		if err := iq.loadWebhooks(ctx, query, nodes,
			func(n *Integration) { n.appendNamedWebhooks(name) },
			func(n *Integration, e *Webhook) { n.appendNamedWebhooks(name, e) }); err != nil {
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
func (iq *IntegrationQuery) loadSecrets(ctx context.Context, query *HushQuery, nodes []*Integration, init func(*Integration), assign func(*Integration, *Hush)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Integration)
	nids := make(map[string]map[*Integration]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(integration.SecretsTable)
		joinT.Schema(iq.schemaConfig.IntegrationSecrets)
		s.Join(joinT).On(s.C(hush.FieldID), joinT.C(integration.SecretsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(integration.SecretsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(integration.SecretsPrimaryKey[0]))
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
					nids[inValue] = map[*Integration]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Hush](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "secrets" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (iq *IntegrationQuery) loadOauth2tokens(ctx context.Context, query *OhAuthTooTokenQuery, nodes []*Integration, init func(*Integration), assign func(*Integration, *OhAuthTooToken)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Integration)
	nids := make(map[string]map[*Integration]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(integration.Oauth2tokensTable)
		joinT.Schema(iq.schemaConfig.IntegrationOauth2tokens)
		s.Join(joinT).On(s.C(ohauthtootoken.FieldID), joinT.C(integration.Oauth2tokensPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(integration.Oauth2tokensPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(integration.Oauth2tokensPrimaryKey[0]))
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
					nids[inValue] = map[*Integration]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*OhAuthTooToken](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "oauth2tokens" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (iq *IntegrationQuery) loadEvents(ctx context.Context, query *EventQuery, nodes []*Integration, init func(*Integration), assign func(*Integration, *Event)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Integration)
	nids := make(map[string]map[*Integration]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(integration.EventsTable)
		joinT.Schema(iq.schemaConfig.IntegrationEvents)
		s.Join(joinT).On(s.C(event.FieldID), joinT.C(integration.EventsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(integration.EventsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(integration.EventsPrimaryKey[0]))
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
					nids[inValue] = map[*Integration]struct{}{byID[outValue]: {}}
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
func (iq *IntegrationQuery) loadWebhooks(ctx context.Context, query *WebhookQuery, nodes []*Integration, init func(*Integration), assign func(*Integration, *Webhook)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Integration)
	nids := make(map[string]map[*Integration]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(integration.WebhooksTable)
		joinT.Schema(iq.schemaConfig.IntegrationWebhooks)
		s.Join(joinT).On(s.C(webhook.FieldID), joinT.C(integration.WebhooksPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(integration.WebhooksPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(integration.WebhooksPrimaryKey[0]))
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
					nids[inValue] = map[*Integration]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Webhook](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "webhooks" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
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

// WithNamedSecrets tells the query-builder to eager-load the nodes that are connected to the "secrets"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithNamedSecrets(name string, opts ...func(*HushQuery)) *IntegrationQuery {
	query := (&HushClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if iq.withNamedSecrets == nil {
		iq.withNamedSecrets = make(map[string]*HushQuery)
	}
	iq.withNamedSecrets[name] = query
	return iq
}

// WithNamedOauth2tokens tells the query-builder to eager-load the nodes that are connected to the "oauth2tokens"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithNamedOauth2tokens(name string, opts ...func(*OhAuthTooTokenQuery)) *IntegrationQuery {
	query := (&OhAuthTooTokenClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if iq.withNamedOauth2tokens == nil {
		iq.withNamedOauth2tokens = make(map[string]*OhAuthTooTokenQuery)
	}
	iq.withNamedOauth2tokens[name] = query
	return iq
}

// WithNamedEvents tells the query-builder to eager-load the nodes that are connected to the "events"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithNamedEvents(name string, opts ...func(*EventQuery)) *IntegrationQuery {
	query := (&EventClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if iq.withNamedEvents == nil {
		iq.withNamedEvents = make(map[string]*EventQuery)
	}
	iq.withNamedEvents[name] = query
	return iq
}

// WithNamedWebhooks tells the query-builder to eager-load the nodes that are connected to the "webhooks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (iq *IntegrationQuery) WithNamedWebhooks(name string, opts ...func(*WebhookQuery)) *IntegrationQuery {
	query := (&WebhookClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if iq.withNamedWebhooks == nil {
		iq.withNamedWebhooks = make(map[string]*WebhookQuery)
	}
	iq.withNamedWebhooks[name] = query
	return iq
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
