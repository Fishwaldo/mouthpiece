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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbusermetadata"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
)

// DbUserMetaDataQuery is the builder for querying DbUserMetaData entities.
type DbUserMetaDataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DbUserMetaData
	withTenant *TenantQuery
	withUser   *DbUserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DbUserMetaDataQuery builder.
func (dumdq *DbUserMetaDataQuery) Where(ps ...predicate.DbUserMetaData) *DbUserMetaDataQuery {
	dumdq.predicates = append(dumdq.predicates, ps...)
	return dumdq
}

// Limit adds a limit step to the query.
func (dumdq *DbUserMetaDataQuery) Limit(limit int) *DbUserMetaDataQuery {
	dumdq.limit = &limit
	return dumdq
}

// Offset adds an offset step to the query.
func (dumdq *DbUserMetaDataQuery) Offset(offset int) *DbUserMetaDataQuery {
	dumdq.offset = &offset
	return dumdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dumdq *DbUserMetaDataQuery) Unique(unique bool) *DbUserMetaDataQuery {
	dumdq.unique = &unique
	return dumdq
}

// Order adds an order step to the query.
func (dumdq *DbUserMetaDataQuery) Order(o ...OrderFunc) *DbUserMetaDataQuery {
	dumdq.order = append(dumdq.order, o...)
	return dumdq
}

// QueryTenant chains the current query on the "tenant" edge.
func (dumdq *DbUserMetaDataQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: dumdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dumdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dumdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbusermetadata.Table, dbusermetadata.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, dbusermetadata.TenantTable, dbusermetadata.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(dumdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (dumdq *DbUserMetaDataQuery) QueryUser() *DbUserQuery {
	query := &DbUserQuery{config: dumdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dumdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dumdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbusermetadata.Table, dbusermetadata.FieldID, selector),
			sqlgraph.To(dbuser.Table, dbuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dbusermetadata.UserTable, dbusermetadata.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(dumdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DbUserMetaData entity from the query.
// Returns a *NotFoundError when no DbUserMetaData was found.
func (dumdq *DbUserMetaDataQuery) First(ctx context.Context) (*DbUserMetaData, error) {
	nodes, err := dumdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dbusermetadata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) FirstX(ctx context.Context) *DbUserMetaData {
	node, err := dumdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DbUserMetaData ID from the query.
// Returns a *NotFoundError when no DbUserMetaData ID was found.
func (dumdq *DbUserMetaDataQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dumdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dbusermetadata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) FirstIDX(ctx context.Context) int {
	id, err := dumdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DbUserMetaData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DbUserMetaData entity is found.
// Returns a *NotFoundError when no DbUserMetaData entities are found.
func (dumdq *DbUserMetaDataQuery) Only(ctx context.Context) (*DbUserMetaData, error) {
	nodes, err := dumdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dbusermetadata.Label}
	default:
		return nil, &NotSingularError{dbusermetadata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) OnlyX(ctx context.Context) *DbUserMetaData {
	node, err := dumdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DbUserMetaData ID in the query.
// Returns a *NotSingularError when more than one DbUserMetaData ID is found.
// Returns a *NotFoundError when no entities are found.
func (dumdq *DbUserMetaDataQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dumdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dbusermetadata.Label}
	default:
		err = &NotSingularError{dbusermetadata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) OnlyIDX(ctx context.Context) int {
	id, err := dumdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DbUserMetaDataSlice.
func (dumdq *DbUserMetaDataQuery) All(ctx context.Context) ([]*DbUserMetaData, error) {
	if err := dumdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dumdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) AllX(ctx context.Context) []*DbUserMetaData {
	nodes, err := dumdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DbUserMetaData IDs.
func (dumdq *DbUserMetaDataQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dumdq.Select(dbusermetadata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) IDsX(ctx context.Context) []int {
	ids, err := dumdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dumdq *DbUserMetaDataQuery) Count(ctx context.Context) (int, error) {
	if err := dumdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dumdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) CountX(ctx context.Context) int {
	count, err := dumdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dumdq *DbUserMetaDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := dumdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dumdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dumdq *DbUserMetaDataQuery) ExistX(ctx context.Context) bool {
	exist, err := dumdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DbUserMetaDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dumdq *DbUserMetaDataQuery) Clone() *DbUserMetaDataQuery {
	if dumdq == nil {
		return nil
	}
	return &DbUserMetaDataQuery{
		config:     dumdq.config,
		limit:      dumdq.limit,
		offset:     dumdq.offset,
		order:      append([]OrderFunc{}, dumdq.order...),
		predicates: append([]predicate.DbUserMetaData{}, dumdq.predicates...),
		withTenant: dumdq.withTenant.Clone(),
		withUser:   dumdq.withUser.Clone(),
		// clone intermediate query.
		sql:    dumdq.sql.Clone(),
		path:   dumdq.path,
		unique: dumdq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (dumdq *DbUserMetaDataQuery) WithTenant(opts ...func(*TenantQuery)) *DbUserMetaDataQuery {
	query := &TenantQuery{config: dumdq.config}
	for _, opt := range opts {
		opt(query)
	}
	dumdq.withTenant = query
	return dumdq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (dumdq *DbUserMetaDataQuery) WithUser(opts ...func(*DbUserQuery)) *DbUserMetaDataQuery {
	query := &DbUserQuery{config: dumdq.config}
	for _, opt := range opts {
		opt(query)
	}
	dumdq.withUser = query
	return dumdq
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
//	client.DbUserMetaData.Query().
//		GroupBy(dbusermetadata.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dumdq *DbUserMetaDataQuery) GroupBy(field string, fields ...string) *DbUserMetaDataGroupBy {
	grbuild := &DbUserMetaDataGroupBy{config: dumdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dumdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dumdq.sqlQuery(ctx), nil
	}
	grbuild.label = dbusermetadata.Label
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
//	client.DbUserMetaData.Query().
//		Select(dbusermetadata.FieldTenantID).
//		Scan(ctx, &v)
func (dumdq *DbUserMetaDataQuery) Select(fields ...string) *DbUserMetaDataSelect {
	dumdq.fields = append(dumdq.fields, fields...)
	selbuild := &DbUserMetaDataSelect{DbUserMetaDataQuery: dumdq}
	selbuild.label = dbusermetadata.Label
	selbuild.flds, selbuild.scan = &dumdq.fields, selbuild.Scan
	return selbuild
}

func (dumdq *DbUserMetaDataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dumdq.fields {
		if !dbusermetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dumdq.path != nil {
		prev, err := dumdq.path(ctx)
		if err != nil {
			return err
		}
		dumdq.sql = prev
	}
	if dbusermetadata.Policy == nil {
		return errors.New("ent: uninitialized dbusermetadata.Policy (forgotten import ent/runtime?)")
	}
	if err := dbusermetadata.Policy.EvalQuery(ctx, dumdq); err != nil {
		return err
	}
	return nil
}

func (dumdq *DbUserMetaDataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DbUserMetaData, error) {
	var (
		nodes       = []*DbUserMetaData{}
		withFKs     = dumdq.withFKs
		_spec       = dumdq.querySpec()
		loadedTypes = [2]bool{
			dumdq.withTenant != nil,
			dumdq.withUser != nil,
		}
	)
	if dumdq.withTenant != nil || dumdq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dbusermetadata.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DbUserMetaData).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DbUserMetaData{config: dumdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dumdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dumdq.withTenant; query != nil {
		if err := dumdq.loadTenant(ctx, query, nodes, nil,
			func(n *DbUserMetaData, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := dumdq.withUser; query != nil {
		if err := dumdq.loadUser(ctx, query, nodes, nil,
			func(n *DbUserMetaData, e *DbUser) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dumdq *DbUserMetaDataQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*DbUserMetaData, init func(*DbUserMetaData), assign func(*DbUserMetaData, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DbUserMetaData)
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
func (dumdq *DbUserMetaDataQuery) loadUser(ctx context.Context, query *DbUserQuery, nodes []*DbUserMetaData, init func(*DbUserMetaData), assign func(*DbUserMetaData, *DbUser)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DbUserMetaData)
	for i := range nodes {
		if nodes[i].db_user_metadata == nil {
			continue
		}
		fk := *nodes[i].db_user_metadata
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(dbuser.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "db_user_metadata" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dumdq *DbUserMetaDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dumdq.querySpec()
	_spec.Node.Columns = dumdq.fields
	if len(dumdq.fields) > 0 {
		_spec.Unique = dumdq.unique != nil && *dumdq.unique
	}
	return sqlgraph.CountNodes(ctx, dumdq.driver, _spec)
}

func (dumdq *DbUserMetaDataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dumdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dumdq *DbUserMetaDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbusermetadata.Table,
			Columns: dbusermetadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbusermetadata.FieldID,
			},
		},
		From:   dumdq.sql,
		Unique: true,
	}
	if unique := dumdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dumdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbusermetadata.FieldID)
		for i := range fields {
			if fields[i] != dbusermetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dumdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dumdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dumdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dumdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dumdq *DbUserMetaDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dumdq.driver.Dialect())
	t1 := builder.Table(dbusermetadata.Table)
	columns := dumdq.fields
	if len(columns) == 0 {
		columns = dbusermetadata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dumdq.sql != nil {
		selector = dumdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dumdq.unique != nil && *dumdq.unique {
		selector.Distinct()
	}
	for _, p := range dumdq.predicates {
		p(selector)
	}
	for _, p := range dumdq.order {
		p(selector)
	}
	if offset := dumdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dumdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DbUserMetaDataGroupBy is the group-by builder for DbUserMetaData entities.
type DbUserMetaDataGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dumdgb *DbUserMetaDataGroupBy) Aggregate(fns ...AggregateFunc) *DbUserMetaDataGroupBy {
	dumdgb.fns = append(dumdgb.fns, fns...)
	return dumdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dumdgb *DbUserMetaDataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dumdgb.path(ctx)
	if err != nil {
		return err
	}
	dumdgb.sql = query
	return dumdgb.sqlScan(ctx, v)
}

func (dumdgb *DbUserMetaDataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dumdgb.fields {
		if !dbusermetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dumdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dumdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dumdgb *DbUserMetaDataGroupBy) sqlQuery() *sql.Selector {
	selector := dumdgb.sql.Select()
	aggregation := make([]string, 0, len(dumdgb.fns))
	for _, fn := range dumdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dumdgb.fields)+len(dumdgb.fns))
		for _, f := range dumdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dumdgb.fields...)...)
}

// DbUserMetaDataSelect is the builder for selecting fields of DbUserMetaData entities.
type DbUserMetaDataSelect struct {
	*DbUserMetaDataQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dumds *DbUserMetaDataSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dumds.prepareQuery(ctx); err != nil {
		return err
	}
	dumds.sql = dumds.DbUserMetaDataQuery.sqlQuery(ctx)
	return dumds.sqlScan(ctx, v)
}

func (dumds *DbUserMetaDataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dumds.sql.Query()
	if err := dumds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
