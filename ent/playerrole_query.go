// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/smxlong/mud/ent/player"
	"github.com/smxlong/mud/ent/playerrole"
	"github.com/smxlong/mud/ent/predicate"
)

// PlayerRoleQuery is the builder for querying PlayerRole entities.
type PlayerRoleQuery struct {
	config
	ctx         *QueryContext
	order       []playerrole.OrderOption
	inters      []Interceptor
	predicates  []predicate.PlayerRole
	withPlayers *PlayerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlayerRoleQuery builder.
func (prq *PlayerRoleQuery) Where(ps ...predicate.PlayerRole) *PlayerRoleQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit the number of records to be returned by this query.
func (prq *PlayerRoleQuery) Limit(limit int) *PlayerRoleQuery {
	prq.ctx.Limit = &limit
	return prq
}

// Offset to start from.
func (prq *PlayerRoleQuery) Offset(offset int) *PlayerRoleQuery {
	prq.ctx.Offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *PlayerRoleQuery) Unique(unique bool) *PlayerRoleQuery {
	prq.ctx.Unique = &unique
	return prq
}

// Order specifies how the records should be ordered.
func (prq *PlayerRoleQuery) Order(o ...playerrole.OrderOption) *PlayerRoleQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// QueryPlayers chains the current query on the "players" edge.
func (prq *PlayerRoleQuery) QueryPlayers() *PlayerQuery {
	query := (&PlayerClient{config: prq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playerrole.Table, playerrole.FieldID, selector),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, playerrole.PlayersTable, playerrole.PlayersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlayerRole entity from the query.
// Returns a *NotFoundError when no PlayerRole was found.
func (prq *PlayerRoleQuery) First(ctx context.Context) (*PlayerRole, error) {
	nodes, err := prq.Limit(1).All(setContextOp(ctx, prq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{playerrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *PlayerRoleQuery) FirstX(ctx context.Context) *PlayerRole {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlayerRole ID from the query.
// Returns a *NotFoundError when no PlayerRole ID was found.
func (prq *PlayerRoleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(1).IDs(setContextOp(ctx, prq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{playerrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *PlayerRoleQuery) FirstIDX(ctx context.Context) int {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlayerRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlayerRole entity is found.
// Returns a *NotFoundError when no PlayerRole entities are found.
func (prq *PlayerRoleQuery) Only(ctx context.Context) (*PlayerRole, error) {
	nodes, err := prq.Limit(2).All(setContextOp(ctx, prq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{playerrole.Label}
	default:
		return nil, &NotSingularError{playerrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *PlayerRoleQuery) OnlyX(ctx context.Context) *PlayerRole {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlayerRole ID in the query.
// Returns a *NotSingularError when more than one PlayerRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (prq *PlayerRoleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(2).IDs(setContextOp(ctx, prq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{playerrole.Label}
	default:
		err = &NotSingularError{playerrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *PlayerRoleQuery) OnlyIDX(ctx context.Context) int {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlayerRoles.
func (prq *PlayerRoleQuery) All(ctx context.Context) ([]*PlayerRole, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryAll)
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PlayerRole, *PlayerRoleQuery]()
	return withInterceptors[[]*PlayerRole](ctx, prq, qr, prq.inters)
}

// AllX is like All, but panics if an error occurs.
func (prq *PlayerRoleQuery) AllX(ctx context.Context) []*PlayerRole {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlayerRole IDs.
func (prq *PlayerRoleQuery) IDs(ctx context.Context) (ids []int, err error) {
	if prq.ctx.Unique == nil && prq.path != nil {
		prq.Unique(true)
	}
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryIDs)
	if err = prq.Select(playerrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *PlayerRoleQuery) IDsX(ctx context.Context) []int {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *PlayerRoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryCount)
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, prq, querierCount[*PlayerRoleQuery](), prq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (prq *PlayerRoleQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *PlayerRoleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryExist)
	switch _, err := prq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *PlayerRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlayerRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *PlayerRoleQuery) Clone() *PlayerRoleQuery {
	if prq == nil {
		return nil
	}
	return &PlayerRoleQuery{
		config:      prq.config,
		ctx:         prq.ctx.Clone(),
		order:       append([]playerrole.OrderOption{}, prq.order...),
		inters:      append([]Interceptor{}, prq.inters...),
		predicates:  append([]predicate.PlayerRole{}, prq.predicates...),
		withPlayers: prq.withPlayers.Clone(),
		// clone intermediate query.
		sql:  prq.sql.Clone(),
		path: prq.path,
	}
}

// WithPlayers tells the query-builder to eager-load the nodes that are connected to
// the "players" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *PlayerRoleQuery) WithPlayers(opts ...func(*PlayerQuery)) *PlayerRoleQuery {
	query := (&PlayerClient{config: prq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prq.withPlayers = query
	return prq
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
//	client.PlayerRole.Query().
//		GroupBy(playerrole.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (prq *PlayerRoleQuery) GroupBy(field string, fields ...string) *PlayerRoleGroupBy {
	prq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlayerRoleGroupBy{build: prq}
	grbuild.flds = &prq.ctx.Fields
	grbuild.label = playerrole.Label
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
//	client.PlayerRole.Query().
//		Select(playerrole.FieldName).
//		Scan(ctx, &v)
func (prq *PlayerRoleQuery) Select(fields ...string) *PlayerRoleSelect {
	prq.ctx.Fields = append(prq.ctx.Fields, fields...)
	sbuild := &PlayerRoleSelect{PlayerRoleQuery: prq}
	sbuild.label = playerrole.Label
	sbuild.flds, sbuild.scan = &prq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlayerRoleSelect configured with the given aggregations.
func (prq *PlayerRoleQuery) Aggregate(fns ...AggregateFunc) *PlayerRoleSelect {
	return prq.Select().Aggregate(fns...)
}

func (prq *PlayerRoleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range prq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, prq); err != nil {
				return err
			}
		}
	}
	for _, f := range prq.ctx.Fields {
		if !playerrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	return nil
}

func (prq *PlayerRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlayerRole, error) {
	var (
		nodes       = []*PlayerRole{}
		_spec       = prq.querySpec()
		loadedTypes = [1]bool{
			prq.withPlayers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PlayerRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PlayerRole{config: prq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := prq.withPlayers; query != nil {
		if err := prq.loadPlayers(ctx, query, nodes,
			func(n *PlayerRole) { n.Edges.Players = []*Player{} },
			func(n *PlayerRole, e *Player) { n.Edges.Players = append(n.Edges.Players, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (prq *PlayerRoleQuery) loadPlayers(ctx context.Context, query *PlayerQuery, nodes []*PlayerRole, init func(*PlayerRole), assign func(*PlayerRole, *Player)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*PlayerRole)
	nids := make(map[string]map[*PlayerRole]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(playerrole.PlayersTable)
		s.Join(joinT).On(s.C(player.FieldID), joinT.C(playerrole.PlayersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(playerrole.PlayersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(playerrole.PlayersPrimaryKey[1]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*PlayerRole]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Player](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "players" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (prq *PlayerRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	_spec.Node.Columns = prq.ctx.Fields
	if len(prq.ctx.Fields) > 0 {
		_spec.Unique = prq.ctx.Unique != nil && *prq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *PlayerRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(playerrole.Table, playerrole.Columns, sqlgraph.NewFieldSpec(playerrole.FieldID, field.TypeInt))
	_spec.From = prq.sql
	if unique := prq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if prq.path != nil {
		_spec.Unique = true
	}
	if fields := prq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playerrole.FieldID)
		for i := range fields {
			if fields[i] != playerrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *PlayerRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(playerrole.Table)
	columns := prq.ctx.Fields
	if len(columns) == 0 {
		columns = playerrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prq.ctx.Unique != nil && *prq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlayerRoleGroupBy is the group-by builder for PlayerRole entities.
type PlayerRoleGroupBy struct {
	selector
	build *PlayerRoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *PlayerRoleGroupBy) Aggregate(fns ...AggregateFunc) *PlayerRoleGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the selector query and scans the result into the given value.
func (prgb *PlayerRoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prgb.build.ctx, ent.OpQueryGroupBy)
	if err := prgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlayerRoleQuery, *PlayerRoleGroupBy](ctx, prgb.build, prgb, prgb.build.inters, v)
}

func (prgb *PlayerRoleGroupBy) sqlScan(ctx context.Context, root *PlayerRoleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*prgb.flds)+len(prgb.fns))
		for _, f := range *prgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*prgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PlayerRoleSelect is the builder for selecting fields of PlayerRole entities.
type PlayerRoleSelect struct {
	*PlayerRoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (prs *PlayerRoleSelect) Aggregate(fns ...AggregateFunc) *PlayerRoleSelect {
	prs.fns = append(prs.fns, fns...)
	return prs
}

// Scan applies the selector query and scans the result into the given value.
func (prs *PlayerRoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prs.ctx, ent.OpQuerySelect)
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlayerRoleQuery, *PlayerRoleSelect](ctx, prs.PlayerRoleQuery, prs, prs.inters, v)
}

func (prs *PlayerRoleSelect) sqlScan(ctx context.Context, root *PlayerRoleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(prs.fns))
	for _, fn := range prs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*prs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
