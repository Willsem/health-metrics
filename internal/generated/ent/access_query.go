// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Willsem/health-metrics/internal/generated/ent/access"
	"github.com/Willsem/health-metrics/internal/generated/ent/predicate"
	"github.com/Willsem/health-metrics/internal/generated/ent/user"
	"github.com/google/uuid"
)

// AccessQuery is the builder for querying Access entities.
type AccessQuery struct {
	config
	ctx        *QueryContext
	order      []access.OrderOption
	inters     []Interceptor
	predicates []predicate.Access
	withOwner  *UserQuery
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccessQuery builder.
func (aq *AccessQuery) Where(ps ...predicate.Access) *AccessQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AccessQuery) Limit(limit int) *AccessQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AccessQuery) Offset(offset int) *AccessQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AccessQuery) Unique(unique bool) *AccessQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AccessQuery) Order(o ...access.OrderOption) *AccessQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryOwner chains the current query on the "owner" edge.
func (aq *AccessQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(access.Table, access.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, access.OwnerTable, access.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (aq *AccessQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(access.Table, access.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, access.UserTable, access.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Access entity from the query.
// Returns a *NotFoundError when no Access was found.
func (aq *AccessQuery) First(ctx context.Context) (*Access, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{access.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AccessQuery) FirstX(ctx context.Context) *Access {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Access ID from the query.
// Returns a *NotFoundError when no Access ID was found.
func (aq *AccessQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{access.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AccessQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Access entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Access entity is found.
// Returns a *NotFoundError when no Access entities are found.
func (aq *AccessQuery) Only(ctx context.Context) (*Access, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{access.Label}
	default:
		return nil, &NotSingularError{access.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AccessQuery) OnlyX(ctx context.Context) *Access {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Access ID in the query.
// Returns a *NotSingularError when more than one Access ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AccessQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{access.Label}
	default:
		err = &NotSingularError{access.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AccessQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Accesses.
func (aq *AccessQuery) All(ctx context.Context) ([]*Access, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Access, *AccessQuery]()
	return withInterceptors[[]*Access](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AccessQuery) AllX(ctx context.Context) []*Access {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Access IDs.
func (aq *AccessQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(access.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AccessQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AccessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AccessQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AccessQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AccessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AccessQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AccessQuery) Clone() *AccessQuery {
	if aq == nil {
		return nil
	}
	return &AccessQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]access.OrderOption{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Access{}, aq.predicates...),
		withOwner:  aq.withOwner.Clone(),
		withUser:   aq.withUser.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccessQuery) WithOwner(opts ...func(*UserQuery)) *AccessQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withOwner = query
	return aq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccessQuery) WithUser(opts ...func(*UserQuery)) *AccessQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (aq *AccessQuery) GroupBy(field string, fields ...string) *AccessGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccessGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = access.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (aq *AccessQuery) Select(fields ...string) *AccessSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AccessSelect{AccessQuery: aq}
	sbuild.label = access.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccessSelect configured with the given aggregations.
func (aq *AccessQuery) Aggregate(fns ...AggregateFunc) *AccessSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AccessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !access.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AccessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Access, error) {
	var (
		nodes       = []*Access{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withOwner != nil,
			aq.withUser != nil,
		}
	)
	if aq.withOwner != nil || aq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, access.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Access).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Access{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withOwner; query != nil {
		if err := aq.loadOwner(ctx, query, nodes, nil,
			func(n *Access, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withUser; query != nil {
		if err := aq.loadUser(ctx, query, nodes, nil,
			func(n *Access, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AccessQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Access, init func(*Access), assign func(*Access, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Access)
	for i := range nodes {
		if nodes[i].user_given_accesses == nil {
			continue
		}
		fk := *nodes[i].user_given_accesses
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
			return fmt.Errorf(`unexpected foreign-key "user_given_accesses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AccessQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Access, init func(*Access), assign func(*Access, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Access)
	for i := range nodes {
		if nodes[i].user_recieved_accesses == nil {
			continue
		}
		fk := *nodes[i].user_recieved_accesses
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
			return fmt.Errorf(`unexpected foreign-key "user_recieved_accesses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AccessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AccessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, access.FieldID)
		for i := range fields {
			if fields[i] != access.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AccessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(access.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = access.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccessGroupBy is the group-by builder for Access entities.
type AccessGroupBy struct {
	selector
	build *AccessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AccessGroupBy) Aggregate(fns ...AggregateFunc) *AccessGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AccessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessQuery, *AccessGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AccessGroupBy) sqlScan(ctx context.Context, root *AccessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccessSelect is the builder for selecting fields of Access entities.
type AccessSelect struct {
	*AccessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AccessSelect) Aggregate(fns ...AggregateFunc) *AccessSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AccessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessQuery, *AccessSelect](ctx, as.AccessQuery, as, as.inters, v)
}

func (as *AccessSelect) sqlScan(ctx context.Context, root *AccessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
