// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    uuid.UUID `msgpack:"i"`
	Value Value     `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// IntegrationEdge is the edge representation of Integration.
type IntegrationEdge struct {
	Node   *Integration `json:"node"`
	Cursor Cursor       `json:"cursor"`
}

// IntegrationConnection is the connection containing edges to Integration.
type IntegrationConnection struct {
	Edges      []*IntegrationEdge `json:"edges"`
	PageInfo   PageInfo           `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

func (c *IntegrationConnection) build(nodes []*Integration, pager *integrationPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Integration
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Integration {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Integration {
			return nodes[i]
		}
	}
	c.Edges = make([]*IntegrationEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &IntegrationEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// IntegrationPaginateOption enables pagination customization.
type IntegrationPaginateOption func(*integrationPager) error

// WithIntegrationOrder configures pagination ordering.
func WithIntegrationOrder(order *IntegrationOrder) IntegrationPaginateOption {
	if order == nil {
		order = DefaultIntegrationOrder
	}
	o := *order
	return func(pager *integrationPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultIntegrationOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithIntegrationFilter configures pagination filter.
func WithIntegrationFilter(filter func(*IntegrationQuery) (*IntegrationQuery, error)) IntegrationPaginateOption {
	return func(pager *integrationPager) error {
		if filter == nil {
			return errors.New("IntegrationQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type integrationPager struct {
	order  *IntegrationOrder
	filter func(*IntegrationQuery) (*IntegrationQuery, error)
}

func newIntegrationPager(opts []IntegrationPaginateOption) (*integrationPager, error) {
	pager := &integrationPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultIntegrationOrder
	}
	return pager, nil
}

func (p *integrationPager) applyFilter(query *IntegrationQuery) (*IntegrationQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *integrationPager) toCursor(i *Integration) Cursor {
	return p.order.Field.toCursor(i)
}

func (p *integrationPager) applyCursors(query *IntegrationQuery, after, before *Cursor) *IntegrationQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultIntegrationOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *integrationPager) applyOrder(query *IntegrationQuery, reverse bool) *IntegrationQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultIntegrationOrder.Field {
		query = query.Order(direction.orderFunc(DefaultIntegrationOrder.Field.field))
	}
	return query
}

func (p *integrationPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultIntegrationOrder.Field {
			b.Comma().Ident(DefaultIntegrationOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Integration.
func (i *IntegrationQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...IntegrationPaginateOption,
) (*IntegrationConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newIntegrationPager(opts)
	if err != nil {
		return nil, err
	}
	if i, err = pager.applyFilter(i); err != nil {
		return nil, err
	}
	conn := &IntegrationConnection{Edges: []*IntegrationEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = i.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	i = pager.applyCursors(i, after, before)
	i = pager.applyOrder(i, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		i.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := i.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := i.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// IntegrationOrderField defines the ordering field of Integration.
type IntegrationOrderField struct {
	field    string
	toCursor func(*Integration) Cursor
}

// IntegrationOrder defines the ordering of Integration.
type IntegrationOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     *IntegrationOrderField `json:"field"`
}

// DefaultIntegrationOrder is the default ordering of Integration.
var DefaultIntegrationOrder = &IntegrationOrder{
	Direction: OrderDirectionAsc,
	Field: &IntegrationOrderField{
		field: integration.FieldID,
		toCursor: func(i *Integration) Cursor {
			return Cursor{ID: i.ID}
		},
	},
}

// ToEdge converts Integration into IntegrationEdge.
func (i *Integration) ToEdge(order *IntegrationOrder) *IntegrationEdge {
	if order == nil {
		order = DefaultIntegrationOrder
	}
	return &IntegrationEdge{
		Node:   i,
		Cursor: order.Field.toCursor(i),
	}
}

// MembershipEdge is the edge representation of Membership.
type MembershipEdge struct {
	Node   *Membership `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// MembershipConnection is the connection containing edges to Membership.
type MembershipConnection struct {
	Edges      []*MembershipEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *MembershipConnection) build(nodes []*Membership, pager *membershipPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Membership
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Membership {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Membership {
			return nodes[i]
		}
	}
	c.Edges = make([]*MembershipEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &MembershipEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// MembershipPaginateOption enables pagination customization.
type MembershipPaginateOption func(*membershipPager) error

// WithMembershipOrder configures pagination ordering.
func WithMembershipOrder(order *MembershipOrder) MembershipPaginateOption {
	if order == nil {
		order = DefaultMembershipOrder
	}
	o := *order
	return func(pager *membershipPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultMembershipOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithMembershipFilter configures pagination filter.
func WithMembershipFilter(filter func(*MembershipQuery) (*MembershipQuery, error)) MembershipPaginateOption {
	return func(pager *membershipPager) error {
		if filter == nil {
			return errors.New("MembershipQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type membershipPager struct {
	order  *MembershipOrder
	filter func(*MembershipQuery) (*MembershipQuery, error)
}

func newMembershipPager(opts []MembershipPaginateOption) (*membershipPager, error) {
	pager := &membershipPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMembershipOrder
	}
	return pager, nil
}

func (p *membershipPager) applyFilter(query *MembershipQuery) (*MembershipQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *membershipPager) toCursor(m *Membership) Cursor {
	return p.order.Field.toCursor(m)
}

func (p *membershipPager) applyCursors(query *MembershipQuery, after, before *Cursor) *MembershipQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultMembershipOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *membershipPager) applyOrder(query *MembershipQuery, reverse bool) *MembershipQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultMembershipOrder.Field {
		query = query.Order(direction.orderFunc(DefaultMembershipOrder.Field.field))
	}
	return query
}

func (p *membershipPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultMembershipOrder.Field {
			b.Comma().Ident(DefaultMembershipOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Membership.
func (m *MembershipQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...MembershipPaginateOption,
) (*MembershipConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newMembershipPager(opts)
	if err != nil {
		return nil, err
	}
	if m, err = pager.applyFilter(m); err != nil {
		return nil, err
	}
	conn := &MembershipConnection{Edges: []*MembershipEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = m.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	m = pager.applyCursors(m, after, before)
	m = pager.applyOrder(m, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		m.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := m.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// MembershipOrderField defines the ordering field of Membership.
type MembershipOrderField struct {
	field    string
	toCursor func(*Membership) Cursor
}

// MembershipOrder defines the ordering of Membership.
type MembershipOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *MembershipOrderField `json:"field"`
}

// DefaultMembershipOrder is the default ordering of Membership.
var DefaultMembershipOrder = &MembershipOrder{
	Direction: OrderDirectionAsc,
	Field: &MembershipOrderField{
		field: membership.FieldID,
		toCursor: func(m *Membership) Cursor {
			return Cursor{ID: m.ID}
		},
	},
}

// ToEdge converts Membership into MembershipEdge.
func (m *Membership) ToEdge(order *MembershipOrder) *MembershipEdge {
	if order == nil {
		order = DefaultMembershipOrder
	}
	return &MembershipEdge{
		Node:   m,
		Cursor: order.Field.toCursor(m),
	}
}

// OrganizationEdge is the edge representation of Organization.
type OrganizationEdge struct {
	Node   *Organization `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// OrganizationConnection is the connection containing edges to Organization.
type OrganizationConnection struct {
	Edges      []*OrganizationEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

func (c *OrganizationConnection) build(nodes []*Organization, pager *organizationPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Organization
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Organization {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Organization {
			return nodes[i]
		}
	}
	c.Edges = make([]*OrganizationEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &OrganizationEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// OrganizationPaginateOption enables pagination customization.
type OrganizationPaginateOption func(*organizationPager) error

// WithOrganizationOrder configures pagination ordering.
func WithOrganizationOrder(order *OrganizationOrder) OrganizationPaginateOption {
	if order == nil {
		order = DefaultOrganizationOrder
	}
	o := *order
	return func(pager *organizationPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOrganizationOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOrganizationFilter configures pagination filter.
func WithOrganizationFilter(filter func(*OrganizationQuery) (*OrganizationQuery, error)) OrganizationPaginateOption {
	return func(pager *organizationPager) error {
		if filter == nil {
			return errors.New("OrganizationQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type organizationPager struct {
	order  *OrganizationOrder
	filter func(*OrganizationQuery) (*OrganizationQuery, error)
}

func newOrganizationPager(opts []OrganizationPaginateOption) (*organizationPager, error) {
	pager := &organizationPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOrganizationOrder
	}
	return pager, nil
}

func (p *organizationPager) applyFilter(query *OrganizationQuery) (*OrganizationQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *organizationPager) toCursor(o *Organization) Cursor {
	return p.order.Field.toCursor(o)
}

func (p *organizationPager) applyCursors(query *OrganizationQuery, after, before *Cursor) *OrganizationQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultOrganizationOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *organizationPager) applyOrder(query *OrganizationQuery, reverse bool) *OrganizationQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultOrganizationOrder.Field {
		query = query.Order(direction.orderFunc(DefaultOrganizationOrder.Field.field))
	}
	return query
}

func (p *organizationPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultOrganizationOrder.Field {
			b.Comma().Ident(DefaultOrganizationOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Organization.
func (o *OrganizationQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OrganizationPaginateOption,
) (*OrganizationConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOrganizationPager(opts)
	if err != nil {
		return nil, err
	}
	if o, err = pager.applyFilter(o); err != nil {
		return nil, err
	}
	conn := &OrganizationConnection{Edges: []*OrganizationEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = o.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	o = pager.applyCursors(o, after, before)
	o = pager.applyOrder(o, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		o.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := o.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := o.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// OrganizationOrderField defines the ordering field of Organization.
type OrganizationOrderField struct {
	field    string
	toCursor func(*Organization) Cursor
}

// OrganizationOrder defines the ordering of Organization.
type OrganizationOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *OrganizationOrderField `json:"field"`
}

// DefaultOrganizationOrder is the default ordering of Organization.
var DefaultOrganizationOrder = &OrganizationOrder{
	Direction: OrderDirectionAsc,
	Field: &OrganizationOrderField{
		field: organization.FieldID,
		toCursor: func(o *Organization) Cursor {
			return Cursor{ID: o.ID}
		},
	},
}

// ToEdge converts Organization into OrganizationEdge.
func (o *Organization) ToEdge(order *OrganizationOrder) *OrganizationEdge {
	if order == nil {
		order = DefaultOrganizationOrder
	}
	return &OrganizationEdge{
		Node:   o,
		Cursor: order.Field.toCursor(o),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

func (p *userPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
