// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/file"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// FileQuery is the builder for querying File entities.
type FileQuery struct {
	config
	ctx                   *QueryContext
	order                 []file.OrderOption
	inters                []Interceptor
	predicates            []predicate.File
	withUser              *UserQuery
	withOrganization      *OrganizationQuery
	withGroup             *GroupQuery
	withFKs               bool
	modifiers             []func(*sql.Selector)
	loadTotal             []func(context.Context, []*File) error
	withNamedOrganization map[string]*OrganizationQuery
	withNamedGroup        map[string]*GroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileQuery builder.
func (fq *FileQuery) Where(ps ...predicate.File) *FileQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FileQuery) Limit(limit int) *FileQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FileQuery) Offset(offset int) *FileQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FileQuery) Unique(unique bool) *FileQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FileQuery) Order(o ...file.OrderOption) *FileQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryUser chains the current query on the "user" edge.
func (fq *FileQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, file.UserTable, file.UserColumn),
		)
		schemaConfig := fq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.File
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (fq *FileQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, file.OrganizationTable, file.OrganizationPrimaryKey...),
		)
		schemaConfig := fq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.OrganizationFiles
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (fq *FileQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, file.GroupTable, file.GroupPrimaryKey...),
		)
		schemaConfig := fq.schemaConfig
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.GroupFiles
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first File entity from the query.
// Returns a *NotFoundError when no File was found.
func (fq *FileQuery) First(ctx context.Context) (*File, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{file.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FileQuery) FirstX(ctx context.Context) *File {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first File ID from the query.
// Returns a *NotFoundError when no File ID was found.
func (fq *FileQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{file.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FileQuery) FirstIDX(ctx context.Context) string {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single File entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one File entity is found.
// Returns a *NotFoundError when no File entities are found.
func (fq *FileQuery) Only(ctx context.Context) (*File, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{file.Label}
	default:
		return nil, &NotSingularError{file.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FileQuery) OnlyX(ctx context.Context) *File {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only File ID in the query.
// Returns a *NotSingularError when more than one File ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FileQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{file.Label}
	default:
		err = &NotSingularError{file.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FileQuery) OnlyIDX(ctx context.Context) string {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Files.
func (fq *FileQuery) All(ctx context.Context) ([]*File, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*File, *FileQuery]()
	return withInterceptors[[]*File](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FileQuery) AllX(ctx context.Context) []*File {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of File IDs.
func (fq *FileQuery) IDs(ctx context.Context) (ids []string, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err = fq.Select(file.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FileQuery) IDsX(ctx context.Context) []string {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FileQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FileQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FileQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FileQuery) Clone() *FileQuery {
	if fq == nil {
		return nil
	}
	return &FileQuery{
		config:           fq.config,
		ctx:              fq.ctx.Clone(),
		order:            append([]file.OrderOption{}, fq.order...),
		inters:           append([]Interceptor{}, fq.inters...),
		predicates:       append([]predicate.File{}, fq.predicates...),
		withUser:         fq.withUser.Clone(),
		withOrganization: fq.withOrganization.Clone(),
		withGroup:        fq.withGroup.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithUser(opts ...func(*UserQuery)) *FileQuery {
	query := (&UserClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withUser = query
	return fq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithOrganization(opts ...func(*OrganizationQuery)) *FileQuery {
	query := (&OrganizationClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withOrganization = query
	return fq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithGroup(opts ...func(*GroupQuery)) *FileQuery {
	query := (&GroupClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withGroup = query
	return fq
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
//	client.File.Query().
//		GroupBy(file.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (fq *FileQuery) GroupBy(field string, fields ...string) *FileGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FileGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = file.Label
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
//	client.File.Query().
//		Select(file.FieldCreatedAt).
//		Scan(ctx, &v)
func (fq *FileQuery) Select(fields ...string) *FileSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FileSelect{FileQuery: fq}
	sbuild.label = file.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FileSelect configured with the given aggregations.
func (fq *FileQuery) Aggregate(fns ...AggregateFunc) *FileSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !file.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*File, error) {
	var (
		nodes       = []*File{}
		withFKs     = fq.withFKs
		_spec       = fq.querySpec()
		loadedTypes = [3]bool{
			fq.withUser != nil,
			fq.withOrganization != nil,
			fq.withGroup != nil,
		}
	)
	if fq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, file.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*File).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &File{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = fq.schemaConfig.File
	ctx = internal.NewSchemaConfigContext(ctx, fq.schemaConfig)
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withUser; query != nil {
		if err := fq.loadUser(ctx, query, nodes, nil,
			func(n *File, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withOrganization; query != nil {
		if err := fq.loadOrganization(ctx, query, nodes,
			func(n *File) { n.Edges.Organization = []*Organization{} },
			func(n *File, e *Organization) { n.Edges.Organization = append(n.Edges.Organization, e) }); err != nil {
			return nil, err
		}
	}
	if query := fq.withGroup; query != nil {
		if err := fq.loadGroup(ctx, query, nodes,
			func(n *File) { n.Edges.Group = []*Group{} },
			func(n *File, e *Group) { n.Edges.Group = append(n.Edges.Group, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range fq.withNamedOrganization {
		if err := fq.loadOrganization(ctx, query, nodes,
			func(n *File) { n.appendNamedOrganization(name) },
			func(n *File, e *Organization) { n.appendNamedOrganization(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range fq.withNamedGroup {
		if err := fq.loadGroup(ctx, query, nodes,
			func(n *File) { n.appendNamedGroup(name) },
			func(n *File, e *Group) { n.appendNamedGroup(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range fq.loadTotal {
		if err := fq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FileQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*File, init func(*File), assign func(*File, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*File)
	for i := range nodes {
		if nodes[i].user_files == nil {
			continue
		}
		fk := *nodes[i].user_files
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
			return fmt.Errorf(`unexpected foreign-key "user_files" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FileQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*File, init func(*File), assign func(*File, *Organization)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*File)
	nids := make(map[string]map[*File]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(file.OrganizationTable)
		joinT.Schema(fq.schemaConfig.OrganizationFiles)
		s.Join(joinT).On(s.C(organization.FieldID), joinT.C(file.OrganizationPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(file.OrganizationPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(file.OrganizationPrimaryKey[1]))
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
					nids[inValue] = map[*File]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Organization](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "organization" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (fq *FileQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*File, init func(*File), assign func(*File, *Group)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*File)
	nids := make(map[string]map[*File]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(file.GroupTable)
		joinT.Schema(fq.schemaConfig.GroupFiles)
		s.Join(joinT).On(s.C(group.FieldID), joinT.C(file.GroupPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(file.GroupPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(file.GroupPrimaryKey[1]))
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
					nids[inValue] = map[*File]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Group](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "group" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (fq *FileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Schema = fq.schemaConfig.File
	ctx = internal.NewSchemaConfigContext(ctx, fq.schemaConfig)
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeString))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for i := range fields {
			if fields[i] != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(file.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = file.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(fq.schemaConfig.File)
	ctx = internal.NewSchemaConfigContext(ctx, fq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOrganization tells the query-builder to eager-load the nodes that are connected to the "organization"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithNamedOrganization(name string, opts ...func(*OrganizationQuery)) *FileQuery {
	query := (&OrganizationClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if fq.withNamedOrganization == nil {
		fq.withNamedOrganization = make(map[string]*OrganizationQuery)
	}
	fq.withNamedOrganization[name] = query
	return fq
}

// WithNamedGroup tells the query-builder to eager-load the nodes that are connected to the "group"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithNamedGroup(name string, opts ...func(*GroupQuery)) *FileQuery {
	query := (&GroupClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if fq.withNamedGroup == nil {
		fq.withNamedGroup = make(map[string]*GroupQuery)
	}
	fq.withNamedGroup[name] = query
	return fq
}

// FileGroupBy is the group-by builder for File entities.
type FileGroupBy struct {
	selector
	build *FileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FileGroupBy) Aggregate(fns ...AggregateFunc) *FileGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileQuery, *FileGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FileGroupBy) sqlScan(ctx context.Context, root *FileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FileSelect is the builder for selecting fields of File entities.
type FileSelect struct {
	*FileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FileSelect) Aggregate(fns ...AggregateFunc) *FileSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileQuery, *FileSelect](ctx, fs.FileQuery, fs, fs.inters, v)
}

func (fs *FileSelect) sqlScan(ctx context.Context, root *FileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
