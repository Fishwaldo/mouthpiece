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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessagefields"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/google/uuid"
)

// DbMessageFieldsQuery is the builder for querying DbMessageFields entities.
type DbMessageFieldsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DbMessageFields
	withTenant *TenantQuery
	withOwner  *DbMessageQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DbMessageFieldsQuery builder.
func (dmfq *DbMessageFieldsQuery) Where(ps ...predicate.DbMessageFields) *DbMessageFieldsQuery {
	dmfq.predicates = append(dmfq.predicates, ps...)
	return dmfq
}

// Limit adds a limit step to the query.
func (dmfq *DbMessageFieldsQuery) Limit(limit int) *DbMessageFieldsQuery {
	dmfq.limit = &limit
	return dmfq
}

// Offset adds an offset step to the query.
func (dmfq *DbMessageFieldsQuery) Offset(offset int) *DbMessageFieldsQuery {
	dmfq.offset = &offset
	return dmfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dmfq *DbMessageFieldsQuery) Unique(unique bool) *DbMessageFieldsQuery {
	dmfq.unique = &unique
	return dmfq
}

// Order adds an order step to the query.
func (dmfq *DbMessageFieldsQuery) Order(o ...OrderFunc) *DbMessageFieldsQuery {
	dmfq.order = append(dmfq.order, o...)
	return dmfq
}

// QueryTenant chains the current query on the "tenant" edge.
func (dmfq *DbMessageFieldsQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: dmfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbmessagefields.Table, dbmessagefields.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, dbmessagefields.TenantTable, dbmessagefields.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (dmfq *DbMessageFieldsQuery) QueryOwner() *DbMessageQuery {
	query := &DbMessageQuery{config: dmfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbmessagefields.Table, dbmessagefields.FieldID, selector),
			sqlgraph.To(dbmessage.Table, dbmessage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dbmessagefields.OwnerTable, dbmessagefields.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DbMessageFields entity from the query.
// Returns a *NotFoundError when no DbMessageFields was found.
func (dmfq *DbMessageFieldsQuery) First(ctx context.Context) (*DbMessageFields, error) {
	nodes, err := dmfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dbmessagefields.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) FirstX(ctx context.Context) *DbMessageFields {
	node, err := dmfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DbMessageFields ID from the query.
// Returns a *NotFoundError when no DbMessageFields ID was found.
func (dmfq *DbMessageFieldsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dmfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dbmessagefields.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) FirstIDX(ctx context.Context) int {
	id, err := dmfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DbMessageFields entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DbMessageFields entity is found.
// Returns a *NotFoundError when no DbMessageFields entities are found.
func (dmfq *DbMessageFieldsQuery) Only(ctx context.Context) (*DbMessageFields, error) {
	nodes, err := dmfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dbmessagefields.Label}
	default:
		return nil, &NotSingularError{dbmessagefields.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) OnlyX(ctx context.Context) *DbMessageFields {
	node, err := dmfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DbMessageFields ID in the query.
// Returns a *NotSingularError when more than one DbMessageFields ID is found.
// Returns a *NotFoundError when no entities are found.
func (dmfq *DbMessageFieldsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dmfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dbmessagefields.Label}
	default:
		err = &NotSingularError{dbmessagefields.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) OnlyIDX(ctx context.Context) int {
	id, err := dmfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DbMessageFieldsSlice.
func (dmfq *DbMessageFieldsQuery) All(ctx context.Context) ([]*DbMessageFields, error) {
	if err := dmfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dmfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) AllX(ctx context.Context) []*DbMessageFields {
	nodes, err := dmfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DbMessageFields IDs.
func (dmfq *DbMessageFieldsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dmfq.Select(dbmessagefields.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) IDsX(ctx context.Context) []int {
	ids, err := dmfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dmfq *DbMessageFieldsQuery) Count(ctx context.Context) (int, error) {
	if err := dmfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dmfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) CountX(ctx context.Context) int {
	count, err := dmfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dmfq *DbMessageFieldsQuery) Exist(ctx context.Context) (bool, error) {
	if err := dmfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dmfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dmfq *DbMessageFieldsQuery) ExistX(ctx context.Context) bool {
	exist, err := dmfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DbMessageFieldsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dmfq *DbMessageFieldsQuery) Clone() *DbMessageFieldsQuery {
	if dmfq == nil {
		return nil
	}
	return &DbMessageFieldsQuery{
		config:     dmfq.config,
		limit:      dmfq.limit,
		offset:     dmfq.offset,
		order:      append([]OrderFunc{}, dmfq.order...),
		predicates: append([]predicate.DbMessageFields{}, dmfq.predicates...),
		withTenant: dmfq.withTenant.Clone(),
		withOwner:  dmfq.withOwner.Clone(),
		// clone intermediate query.
		sql:    dmfq.sql.Clone(),
		path:   dmfq.path,
		unique: dmfq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (dmfq *DbMessageFieldsQuery) WithTenant(opts ...func(*TenantQuery)) *DbMessageFieldsQuery {
	query := &TenantQuery{config: dmfq.config}
	for _, opt := range opts {
		opt(query)
	}
	dmfq.withTenant = query
	return dmfq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (dmfq *DbMessageFieldsQuery) WithOwner(opts ...func(*DbMessageQuery)) *DbMessageFieldsQuery {
	query := &DbMessageQuery{config: dmfq.config}
	for _, opt := range opts {
		opt(query)
	}
	dmfq.withOwner = query
	return dmfq
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
//	client.DbMessageFields.Query().
//		GroupBy(dbmessagefields.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dmfq *DbMessageFieldsQuery) GroupBy(field string, fields ...string) *DbMessageFieldsGroupBy {
	grbuild := &DbMessageFieldsGroupBy{config: dmfq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dmfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dmfq.sqlQuery(ctx), nil
	}
	grbuild.label = dbmessagefields.Label
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
//	client.DbMessageFields.Query().
//		Select(dbmessagefields.FieldTenantID).
//		Scan(ctx, &v)
func (dmfq *DbMessageFieldsQuery) Select(fields ...string) *DbMessageFieldsSelect {
	dmfq.fields = append(dmfq.fields, fields...)
	selbuild := &DbMessageFieldsSelect{DbMessageFieldsQuery: dmfq}
	selbuild.label = dbmessagefields.Label
	selbuild.flds, selbuild.scan = &dmfq.fields, selbuild.Scan
	return selbuild
}

func (dmfq *DbMessageFieldsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dmfq.fields {
		if !dbmessagefields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dmfq.path != nil {
		prev, err := dmfq.path(ctx)
		if err != nil {
			return err
		}
		dmfq.sql = prev
	}
	if dbmessagefields.Policy == nil {
		return errors.New("ent: uninitialized dbmessagefields.Policy (forgotten import ent/runtime?)")
	}
	if err := dbmessagefields.Policy.EvalQuery(ctx, dmfq); err != nil {
		return err
	}
	return nil
}

func (dmfq *DbMessageFieldsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DbMessageFields, error) {
	var (
		nodes       = []*DbMessageFields{}
		withFKs     = dmfq.withFKs
		_spec       = dmfq.querySpec()
		loadedTypes = [2]bool{
			dmfq.withTenant != nil,
			dmfq.withOwner != nil,
		}
	)
	if dmfq.withTenant != nil || dmfq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dbmessagefields.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DbMessageFields).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DbMessageFields{config: dmfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dmfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dmfq.withTenant; query != nil {
		if err := dmfq.loadTenant(ctx, query, nodes, nil,
			func(n *DbMessageFields, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := dmfq.withOwner; query != nil {
		if err := dmfq.loadOwner(ctx, query, nodes, nil,
			func(n *DbMessageFields, e *DbMessage) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dmfq *DbMessageFieldsQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*DbMessageFields, init func(*DbMessageFields), assign func(*DbMessageFields, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DbMessageFields)
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
func (dmfq *DbMessageFieldsQuery) loadOwner(ctx context.Context, query *DbMessageQuery, nodes []*DbMessageFields, init func(*DbMessageFields), assign func(*DbMessageFields, *DbMessage)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DbMessageFields)
	for i := range nodes {
		if nodes[i].db_message_fields == nil {
			continue
		}
		fk := *nodes[i].db_message_fields
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(dbmessage.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "db_message_fields" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dmfq *DbMessageFieldsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dmfq.querySpec()
	_spec.Node.Columns = dmfq.fields
	if len(dmfq.fields) > 0 {
		_spec.Unique = dmfq.unique != nil && *dmfq.unique
	}
	return sqlgraph.CountNodes(ctx, dmfq.driver, _spec)
}

func (dmfq *DbMessageFieldsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dmfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dmfq *DbMessageFieldsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbmessagefields.Table,
			Columns: dbmessagefields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbmessagefields.FieldID,
			},
		},
		From:   dmfq.sql,
		Unique: true,
	}
	if unique := dmfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dmfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbmessagefields.FieldID)
		for i := range fields {
			if fields[i] != dbmessagefields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dmfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dmfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dmfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dmfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dmfq *DbMessageFieldsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dmfq.driver.Dialect())
	t1 := builder.Table(dbmessagefields.Table)
	columns := dmfq.fields
	if len(columns) == 0 {
		columns = dbmessagefields.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dmfq.sql != nil {
		selector = dmfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dmfq.unique != nil && *dmfq.unique {
		selector.Distinct()
	}
	for _, p := range dmfq.predicates {
		p(selector)
	}
	for _, p := range dmfq.order {
		p(selector)
	}
	if offset := dmfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dmfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DbMessageFieldsGroupBy is the group-by builder for DbMessageFields entities.
type DbMessageFieldsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dmfgb *DbMessageFieldsGroupBy) Aggregate(fns ...AggregateFunc) *DbMessageFieldsGroupBy {
	dmfgb.fns = append(dmfgb.fns, fns...)
	return dmfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dmfgb *DbMessageFieldsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dmfgb.path(ctx)
	if err != nil {
		return err
	}
	dmfgb.sql = query
	return dmfgb.sqlScan(ctx, v)
}

func (dmfgb *DbMessageFieldsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dmfgb.fields {
		if !dbmessagefields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dmfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dmfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dmfgb *DbMessageFieldsGroupBy) sqlQuery() *sql.Selector {
	selector := dmfgb.sql.Select()
	aggregation := make([]string, 0, len(dmfgb.fns))
	for _, fn := range dmfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dmfgb.fields)+len(dmfgb.fns))
		for _, f := range dmfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dmfgb.fields...)...)
}

// DbMessageFieldsSelect is the builder for selecting fields of DbMessageFields entities.
type DbMessageFieldsSelect struct {
	*DbMessageFieldsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dmfs *DbMessageFieldsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dmfs.prepareQuery(ctx); err != nil {
		return err
	}
	dmfs.sql = dmfs.DbMessageFieldsQuery.sqlQuery(ctx)
	return dmfs.sqlScan(ctx, v)
}

func (dmfs *DbMessageFieldsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dmfs.sql.Query()
	if err := dmfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}