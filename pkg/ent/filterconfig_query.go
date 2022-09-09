/*
	MIT License

	Copyright (c) 2021 Justin Hammond

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/filterconfig"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
)

// FilterConfigQuery is the builder for querying FilterConfig entities.
type FilterConfigQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FilterConfig
	withTenant *TenantQuery
	withFilter *FilterQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FilterConfigQuery builder.
func (fcq *FilterConfigQuery) Where(ps ...predicate.FilterConfig) *FilterConfigQuery {
	fcq.predicates = append(fcq.predicates, ps...)
	return fcq
}

// Limit adds a limit step to the query.
func (fcq *FilterConfigQuery) Limit(limit int) *FilterConfigQuery {
	fcq.limit = &limit
	return fcq
}

// Offset adds an offset step to the query.
func (fcq *FilterConfigQuery) Offset(offset int) *FilterConfigQuery {
	fcq.offset = &offset
	return fcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcq *FilterConfigQuery) Unique(unique bool) *FilterConfigQuery {
	fcq.unique = &unique
	return fcq
}

// Order adds an order step to the query.
func (fcq *FilterConfigQuery) Order(o ...OrderFunc) *FilterConfigQuery {
	fcq.order = append(fcq.order, o...)
	return fcq
}

// QueryTenant chains the current query on the "tenant" edge.
func (fcq *FilterConfigQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: fcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filterconfig.Table, filterconfig.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, filterconfig.TenantTable, filterconfig.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFilter chains the current query on the "filter" edge.
func (fcq *FilterConfigQuery) QueryFilter() *FilterQuery {
	query := &FilterQuery{config: fcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filterconfig.Table, filterconfig.FieldID, selector),
			sqlgraph.To(filter.Table, filter.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, filterconfig.FilterTable, filterconfig.FilterColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FilterConfig entity from the query.
// Returns a *NotFoundError when no FilterConfig was found.
func (fcq *FilterConfigQuery) First(ctx context.Context) (*FilterConfig, error) {
	nodes, err := fcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filterconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fcq *FilterConfigQuery) FirstX(ctx context.Context) *FilterConfig {
	node, err := fcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FilterConfig ID from the query.
// Returns a *NotFoundError when no FilterConfig ID was found.
func (fcq *FilterConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filterconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fcq *FilterConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := fcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FilterConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FilterConfig entity is found.
// Returns a *NotFoundError when no FilterConfig entities are found.
func (fcq *FilterConfigQuery) Only(ctx context.Context) (*FilterConfig, error) {
	nodes, err := fcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filterconfig.Label}
	default:
		return nil, &NotSingularError{filterconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fcq *FilterConfigQuery) OnlyX(ctx context.Context) *FilterConfig {
	node, err := fcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FilterConfig ID in the query.
// Returns a *NotSingularError when more than one FilterConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (fcq *FilterConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filterconfig.Label}
	default:
		err = &NotSingularError{filterconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fcq *FilterConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := fcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FilterConfigs.
func (fcq *FilterConfigQuery) All(ctx context.Context) ([]*FilterConfig, error) {
	if err := fcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fcq *FilterConfigQuery) AllX(ctx context.Context) []*FilterConfig {
	nodes, err := fcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FilterConfig IDs.
func (fcq *FilterConfigQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fcq.Select(filterconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fcq *FilterConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := fcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fcq *FilterConfigQuery) Count(ctx context.Context) (int, error) {
	if err := fcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fcq *FilterConfigQuery) CountX(ctx context.Context) int {
	count, err := fcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fcq *FilterConfigQuery) Exist(ctx context.Context) (bool, error) {
	if err := fcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fcq *FilterConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := fcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FilterConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fcq *FilterConfigQuery) Clone() *FilterConfigQuery {
	if fcq == nil {
		return nil
	}
	return &FilterConfigQuery{
		config:     fcq.config,
		limit:      fcq.limit,
		offset:     fcq.offset,
		order:      append([]OrderFunc{}, fcq.order...),
		predicates: append([]predicate.FilterConfig{}, fcq.predicates...),
		withTenant: fcq.withTenant.Clone(),
		withFilter: fcq.withFilter.Clone(),
		// clone intermediate query.
		sql:    fcq.sql.Clone(),
		path:   fcq.path,
		unique: fcq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FilterConfigQuery) WithTenant(opts ...func(*TenantQuery)) *FilterConfigQuery {
	query := &TenantQuery{config: fcq.config}
	for _, opt := range opts {
		opt(query)
	}
	fcq.withTenant = query
	return fcq
}

// WithFilter tells the query-builder to eager-load the nodes that are connected to
// the "filter" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FilterConfigQuery) WithFilter(opts ...func(*FilterQuery)) *FilterConfigQuery {
	query := &FilterQuery{config: fcq.config}
	for _, opt := range opts {
		opt(query)
	}
	fcq.withFilter = query
	return fcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TenantID int `json:"tenant_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FilterConfig.Query().
//		GroupBy(filterconfig.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fcq *FilterConfigQuery) GroupBy(field string, fields ...string) *FilterConfigGroupBy {
	grbuild := &FilterConfigGroupBy{config: fcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fcq.sqlQuery(ctx), nil
	}
	grbuild.label = filterconfig.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TenantID int `json:"tenant_id,omitempty"`
//	}
//
//	client.FilterConfig.Query().
//		Select(filterconfig.FieldTenantID).
//		Scan(ctx, &v)
//
func (fcq *FilterConfigQuery) Select(fields ...string) *FilterConfigSelect {
	fcq.fields = append(fcq.fields, fields...)
	selbuild := &FilterConfigSelect{FilterConfigQuery: fcq}
	selbuild.label = filterconfig.Label
	selbuild.flds, selbuild.scan = &fcq.fields, selbuild.Scan
	return selbuild
}

func (fcq *FilterConfigQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fcq.fields {
		if !filterconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fcq.path != nil {
		prev, err := fcq.path(ctx)
		if err != nil {
			return err
		}
		fcq.sql = prev
	}
	if filterconfig.Policy == nil {
		return errors.New("ent: uninitialized filterconfig.Policy (forgotten import ent/runtime?)")
	}
	if err := filterconfig.Policy.EvalQuery(ctx, fcq); err != nil {
		return err
	}
	return nil
}

func (fcq *FilterConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FilterConfig, error) {
	var (
		nodes       = []*FilterConfig{}
		withFKs     = fcq.withFKs
		_spec       = fcq.querySpec()
		loadedTypes = [2]bool{
			fcq.withTenant != nil,
			fcq.withFilter != nil,
		}
	)
	if fcq.withTenant != nil || fcq.withFilter != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, filterconfig.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FilterConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FilterConfig{config: fcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fcq.withTenant; query != nil {
		if err := fcq.loadTenant(ctx, query, nodes, nil,
			func(n *FilterConfig, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := fcq.withFilter; query != nil {
		if err := fcq.loadFilter(ctx, query, nodes, nil,
			func(n *FilterConfig, e *Filter) { n.Edges.Filter = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fcq *FilterConfigQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*FilterConfig, init func(*FilterConfig), assign func(*FilterConfig, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*FilterConfig)
	for i := range nodes {
		fk := nodes[i].TenantID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tenant.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tenant_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fcq *FilterConfigQuery) loadFilter(ctx context.Context, query *FilterQuery, nodes []*FilterConfig, init func(*FilterConfig), assign func(*FilterConfig, *Filter)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*FilterConfig)
	for i := range nodes {
		if nodes[i].filter_config == nil {
			continue
		}
		fk := *nodes[i].filter_config
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(filter.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "filter_config" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fcq *FilterConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcq.querySpec()
	_spec.Node.Columns = fcq.fields
	if len(fcq.fields) > 0 {
		_spec.Unique = fcq.unique != nil && *fcq.unique
	}
	return sqlgraph.CountNodes(ctx, fcq.driver, _spec)
}

func (fcq *FilterConfigQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fcq *FilterConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filterconfig.Table,
			Columns: filterconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filterconfig.FieldID,
			},
		},
		From:   fcq.sql,
		Unique: true,
	}
	if unique := fcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filterconfig.FieldID)
		for i := range fields {
			if fields[i] != filterconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fcq *FilterConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fcq.driver.Dialect())
	t1 := builder.Table(filterconfig.Table)
	columns := fcq.fields
	if len(columns) == 0 {
		columns = filterconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcq.sql != nil {
		selector = fcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcq.unique != nil && *fcq.unique {
		selector.Distinct()
	}
	for _, p := range fcq.predicates {
		p(selector)
	}
	for _, p := range fcq.order {
		p(selector)
	}
	if offset := fcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FilterConfigGroupBy is the group-by builder for FilterConfig entities.
type FilterConfigGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcgb *FilterConfigGroupBy) Aggregate(fns ...AggregateFunc) *FilterConfigGroupBy {
	fcgb.fns = append(fcgb.fns, fns...)
	return fcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fcgb *FilterConfigGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fcgb.path(ctx)
	if err != nil {
		return err
	}
	fcgb.sql = query
	return fcgb.sqlScan(ctx, v)
}

func (fcgb *FilterConfigGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fcgb.fields {
		if !filterconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fcgb *FilterConfigGroupBy) sqlQuery() *sql.Selector {
	selector := fcgb.sql.Select()
	aggregation := make([]string, 0, len(fcgb.fns))
	for _, fn := range fcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fcgb.fields)+len(fcgb.fns))
		for _, f := range fcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fcgb.fields...)...)
}

// FilterConfigSelect is the builder for selecting fields of FilterConfig entities.
type FilterConfigSelect struct {
	*FilterConfigQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fcs *FilterConfigSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fcs.prepareQuery(ctx); err != nil {
		return err
	}
	fcs.sql = fcs.FilterConfigQuery.sqlQuery(ctx)
	return fcs.sqlScan(ctx, v)
}

func (fcs *FilterConfigSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fcs.sql.Query()
	if err := fcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
