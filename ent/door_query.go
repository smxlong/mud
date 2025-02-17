// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/smxlong/mud/ent/door"
	"github.com/smxlong/mud/ent/predicate"
	"github.com/smxlong/mud/ent/room"
)

// DoorQuery is the builder for querying Door entities.
type DoorQuery struct {
	config
	ctx        *QueryContext
	order      []door.OrderOption
	inters     []Interceptor
	predicates []predicate.Door
	withFrom   *RoomQuery
	withTo     *RoomQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DoorQuery builder.
func (dq *DoorQuery) Where(ps ...predicate.Door) *DoorQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DoorQuery) Limit(limit int) *DoorQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DoorQuery) Offset(offset int) *DoorQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DoorQuery) Unique(unique bool) *DoorQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DoorQuery) Order(o ...door.OrderOption) *DoorQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryFrom chains the current query on the "from" edge.
func (dq *DoorQuery) QueryFrom() *RoomQuery {
	query := (&RoomClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(door.Table, door.FieldID, selector),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, door.FromTable, door.FromColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTo chains the current query on the "to" edge.
func (dq *DoorQuery) QueryTo() *RoomQuery {
	query := (&RoomClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(door.Table, door.FieldID, selector),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, door.ToTable, door.ToColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Door entity from the query.
// Returns a *NotFoundError when no Door was found.
func (dq *DoorQuery) First(ctx context.Context) (*Door, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{door.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DoorQuery) FirstX(ctx context.Context) *Door {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Door ID from the query.
// Returns a *NotFoundError when no Door ID was found.
func (dq *DoorQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{door.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DoorQuery) FirstIDX(ctx context.Context) string {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Door entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Door entity is found.
// Returns a *NotFoundError when no Door entities are found.
func (dq *DoorQuery) Only(ctx context.Context) (*Door, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{door.Label}
	default:
		return nil, &NotSingularError{door.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DoorQuery) OnlyX(ctx context.Context) *Door {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Door ID in the query.
// Returns a *NotSingularError when more than one Door ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DoorQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{door.Label}
	default:
		err = &NotSingularError{door.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DoorQuery) OnlyIDX(ctx context.Context) string {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Doors.
func (dq *DoorQuery) All(ctx context.Context) ([]*Door, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Door, *DoorQuery]()
	return withInterceptors[[]*Door](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DoorQuery) AllX(ctx context.Context) []*Door {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Door IDs.
func (dq *DoorQuery) IDs(ctx context.Context) (ids []string, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(door.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DoorQuery) IDsX(ctx context.Context) []string {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DoorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DoorQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DoorQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DoorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DoorQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DoorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DoorQuery) Clone() *DoorQuery {
	if dq == nil {
		return nil
	}
	return &DoorQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]door.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Door{}, dq.predicates...),
		withFrom:   dq.withFrom.Clone(),
		withTo:     dq.withTo.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithFrom tells the query-builder to eager-load the nodes that are connected to
// the "from" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DoorQuery) WithFrom(opts ...func(*RoomQuery)) *DoorQuery {
	query := (&RoomClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withFrom = query
	return dq
}

// WithTo tells the query-builder to eager-load the nodes that are connected to
// the "to" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DoorQuery) WithTo(opts ...func(*RoomQuery)) *DoorQuery {
	query := (&RoomClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withTo = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Door.Query().
//		GroupBy(door.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DoorQuery) GroupBy(field string, fields ...string) *DoorGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DoorGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = door.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Door.Query().
//		Select(door.FieldName).
//		Scan(ctx, &v)
func (dq *DoorQuery) Select(fields ...string) *DoorSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DoorSelect{DoorQuery: dq}
	sbuild.label = door.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DoorSelect configured with the given aggregations.
func (dq *DoorQuery) Aggregate(fns ...AggregateFunc) *DoorSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DoorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !door.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DoorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Door, error) {
	var (
		nodes       = []*Door{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [2]bool{
			dq.withFrom != nil,
			dq.withTo != nil,
		}
	)
	if dq.withFrom != nil || dq.withTo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, door.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Door).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Door{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withFrom; query != nil {
		if err := dq.loadFrom(ctx, query, nodes, nil,
			func(n *Door, e *Room) { n.Edges.From = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withTo; query != nil {
		if err := dq.loadTo(ctx, query, nodes, nil,
			func(n *Door, e *Room) { n.Edges.To = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DoorQuery) loadFrom(ctx context.Context, query *RoomQuery, nodes []*Door, init func(*Door), assign func(*Door, *Room)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Door)
	for i := range nodes {
		if nodes[i].door_from == nil {
			continue
		}
		fk := *nodes[i].door_from
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(room.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "door_from" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DoorQuery) loadTo(ctx context.Context, query *RoomQuery, nodes []*Door, init func(*Door), assign func(*Door, *Room)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Door)
	for i := range nodes {
		if nodes[i].door_to == nil {
			continue
		}
		fk := *nodes[i].door_to
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(room.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "door_to" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DoorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DoorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(door.Table, door.Columns, sqlgraph.NewFieldSpec(door.FieldID, field.TypeString))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, door.FieldID)
		for i := range fields {
			if fields[i] != door.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DoorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(door.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = door.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DoorGroupBy is the group-by builder for Door entities.
type DoorGroupBy struct {
	selector
	build *DoorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DoorGroupBy) Aggregate(fns ...AggregateFunc) *DoorGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DoorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DoorQuery, *DoorGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DoorGroupBy) sqlScan(ctx context.Context, root *DoorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DoorSelect is the builder for selecting fields of Door entities.
type DoorSelect struct {
	*DoorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DoorSelect) Aggregate(fns ...AggregateFunc) *DoorSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DoorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DoorQuery, *DoorSelect](ctx, ds.DoorQuery, ds, ds.inters, v)
}

func (ds *DoorSelect) sqlScan(ctx context.Context, root *DoorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
