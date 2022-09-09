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
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportinstance"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportrecipient"
)

// TransportInstanceQuery is the builder for querying TransportInstance entities.
type TransportInstanceQuery struct {
	config
	limit                   *int
	offset                  *int
	unique                  *bool
	order                   []OrderFunc
	fields                  []string
	predicates              []predicate.TransportInstance
	withTenant              *TenantQuery
	withTransportRecipients *TransportRecipientQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransportInstanceQuery builder.
func (tiq *TransportInstanceQuery) Where(ps ...predicate.TransportInstance) *TransportInstanceQuery {
	tiq.predicates = append(tiq.predicates, ps...)
	return tiq
}

// Limit adds a limit step to the query.
func (tiq *TransportInstanceQuery) Limit(limit int) *TransportInstanceQuery {
	tiq.limit = &limit
	return tiq
}

// Offset adds an offset step to the query.
func (tiq *TransportInstanceQuery) Offset(offset int) *TransportInstanceQuery {
	tiq.offset = &offset
	return tiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tiq *TransportInstanceQuery) Unique(unique bool) *TransportInstanceQuery {
	tiq.unique = &unique
	return tiq
}

// Order adds an order step to the query.
func (tiq *TransportInstanceQuery) Order(o ...OrderFunc) *TransportInstanceQuery {
	tiq.order = append(tiq.order, o...)
	return tiq
}

// QueryTenant chains the current query on the "tenant" edge.
func (tiq *TransportInstanceQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: tiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transportinstance.Table, transportinstance.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, transportinstance.TenantTable, transportinstance.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(tiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransportRecipients chains the current query on the "TransportRecipients" edge.
func (tiq *TransportInstanceQuery) QueryTransportRecipients() *TransportRecipientQuery {
	query := &TransportRecipientQuery{config: tiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transportinstance.Table, transportinstance.FieldID, selector),
			sqlgraph.To(transportrecipient.Table, transportrecipient.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transportinstance.TransportRecipientsTable, transportinstance.TransportRecipientsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TransportInstance entity from the query.
// Returns a *NotFoundError when no TransportInstance was found.
func (tiq *TransportInstanceQuery) First(ctx context.Context) (*TransportInstance, error) {
	nodes, err := tiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transportinstance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tiq *TransportInstanceQuery) FirstX(ctx context.Context) *TransportInstance {
	node, err := tiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TransportInstance ID from the query.
// Returns a *NotFoundError when no TransportInstance ID was found.
func (tiq *TransportInstanceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transportinstance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tiq *TransportInstanceQuery) FirstIDX(ctx context.Context) int {
	id, err := tiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TransportInstance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TransportInstance entity is found.
// Returns a *NotFoundError when no TransportInstance entities are found.
func (tiq *TransportInstanceQuery) Only(ctx context.Context) (*TransportInstance, error) {
	nodes, err := tiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transportinstance.Label}
	default:
		return nil, &NotSingularError{transportinstance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tiq *TransportInstanceQuery) OnlyX(ctx context.Context) *TransportInstance {
	node, err := tiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TransportInstance ID in the query.
// Returns a *NotSingularError when more than one TransportInstance ID is found.
// Returns a *NotFoundError when no entities are found.
func (tiq *TransportInstanceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transportinstance.Label}
	default:
		err = &NotSingularError{transportinstance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tiq *TransportInstanceQuery) OnlyIDX(ctx context.Context) int {
	id, err := tiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TransportInstances.
func (tiq *TransportInstanceQuery) All(ctx context.Context) ([]*TransportInstance, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tiq *TransportInstanceQuery) AllX(ctx context.Context) []*TransportInstance {
	nodes, err := tiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TransportInstance IDs.
func (tiq *TransportInstanceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tiq.Select(transportinstance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tiq *TransportInstanceQuery) IDsX(ctx context.Context) []int {
	ids, err := tiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tiq *TransportInstanceQuery) Count(ctx context.Context) (int, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tiq *TransportInstanceQuery) CountX(ctx context.Context) int {
	count, err := tiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tiq *TransportInstanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tiq *TransportInstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := tiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransportInstanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tiq *TransportInstanceQuery) Clone() *TransportInstanceQuery {
	if tiq == nil {
		return nil
	}
	return &TransportInstanceQuery{
		config:                  tiq.config,
		limit:                   tiq.limit,
		offset:                  tiq.offset,
		order:                   append([]OrderFunc{}, tiq.order...),
		predicates:              append([]predicate.TransportInstance{}, tiq.predicates...),
		withTenant:              tiq.withTenant.Clone(),
		withTransportRecipients: tiq.withTransportRecipients.Clone(),
		// clone intermediate query.
		sql:    tiq.sql.Clone(),
		path:   tiq.path,
		unique: tiq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (tiq *TransportInstanceQuery) WithTenant(opts ...func(*TenantQuery)) *TransportInstanceQuery {
	query := &TenantQuery{config: tiq.config}
	for _, opt := range opts {
		opt(query)
	}
	tiq.withTenant = query
	return tiq
}

// WithTransportRecipients tells the query-builder to eager-load the nodes that are connected to
// the "TransportRecipients" edge. The optional arguments are used to configure the query builder of the edge.
func (tiq *TransportInstanceQuery) WithTransportRecipients(opts ...func(*TransportRecipientQuery)) *TransportInstanceQuery {
	query := &TransportRecipientQuery{config: tiq.config}
	for _, opt := range opts {
		opt(query)
	}
	tiq.withTransportRecipients = query
	return tiq
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
//	client.TransportInstance.Query().
//		GroupBy(transportinstance.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tiq *TransportInstanceQuery) GroupBy(field string, fields ...string) *TransportInstanceGroupBy {
	grbuild := &TransportInstanceGroupBy{config: tiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tiq.sqlQuery(ctx), nil
	}
	grbuild.label = transportinstance.Label
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
//	client.TransportInstance.Query().
//		Select(transportinstance.FieldTenantID).
//		Scan(ctx, &v)
//
func (tiq *TransportInstanceQuery) Select(fields ...string) *TransportInstanceSelect {
	tiq.fields = append(tiq.fields, fields...)
	selbuild := &TransportInstanceSelect{TransportInstanceQuery: tiq}
	selbuild.label = transportinstance.Label
	selbuild.flds, selbuild.scan = &tiq.fields, selbuild.Scan
	return selbuild
}

func (tiq *TransportInstanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tiq.fields {
		if !transportinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tiq.path != nil {
		prev, err := tiq.path(ctx)
		if err != nil {
			return err
		}
		tiq.sql = prev
	}
	if transportinstance.Policy == nil {
		return errors.New("ent: uninitialized transportinstance.Policy (forgotten import ent/runtime?)")
	}
	if err := transportinstance.Policy.EvalQuery(ctx, tiq); err != nil {
		return err
	}
	return nil
}

func (tiq *TransportInstanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TransportInstance, error) {
	var (
		nodes       = []*TransportInstance{}
		_spec       = tiq.querySpec()
		loadedTypes = [2]bool{
			tiq.withTenant != nil,
			tiq.withTransportRecipients != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TransportInstance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TransportInstance{config: tiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tiq.withTenant; query != nil {
		if err := tiq.loadTenant(ctx, query, nodes, nil,
			func(n *TransportInstance, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := tiq.withTransportRecipients; query != nil {
		if err := tiq.loadTransportRecipients(ctx, query, nodes,
			func(n *TransportInstance) { n.Edges.TransportRecipients = []*TransportRecipient{} },
			func(n *TransportInstance, e *TransportRecipient) {
				n.Edges.TransportRecipients = append(n.Edges.TransportRecipients, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tiq *TransportInstanceQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*TransportInstance, init func(*TransportInstance), assign func(*TransportInstance, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TransportInstance)
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
func (tiq *TransportInstanceQuery) loadTransportRecipients(ctx context.Context, query *TransportRecipientQuery, nodes []*TransportInstance, init func(*TransportInstance), assign func(*TransportInstance, *TransportRecipient)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TransportInstance)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.InValues(transportinstance.TransportRecipientsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.transport_instance_transport_recipients
		if fk == nil {
			return fmt.Errorf(`foreign-key "transport_instance_transport_recipients" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "transport_instance_transport_recipients" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tiq *TransportInstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tiq.querySpec()
	_spec.Node.Columns = tiq.fields
	if len(tiq.fields) > 0 {
		_spec.Unique = tiq.unique != nil && *tiq.unique
	}
	return sqlgraph.CountNodes(ctx, tiq.driver, _spec)
}

func (tiq *TransportInstanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tiq *TransportInstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transportinstance.Table,
			Columns: transportinstance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transportinstance.FieldID,
			},
		},
		From:   tiq.sql,
		Unique: true,
	}
	if unique := tiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transportinstance.FieldID)
		for i := range fields {
			if fields[i] != transportinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tiq *TransportInstanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tiq.driver.Dialect())
	t1 := builder.Table(transportinstance.Table)
	columns := tiq.fields
	if len(columns) == 0 {
		columns = transportinstance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tiq.sql != nil {
		selector = tiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tiq.unique != nil && *tiq.unique {
		selector.Distinct()
	}
	for _, p := range tiq.predicates {
		p(selector)
	}
	for _, p := range tiq.order {
		p(selector)
	}
	if offset := tiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TransportInstanceGroupBy is the group-by builder for TransportInstance entities.
type TransportInstanceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tigb *TransportInstanceGroupBy) Aggregate(fns ...AggregateFunc) *TransportInstanceGroupBy {
	tigb.fns = append(tigb.fns, fns...)
	return tigb
}

// Scan applies the group-by query and scans the result into the given value.
func (tigb *TransportInstanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tigb.path(ctx)
	if err != nil {
		return err
	}
	tigb.sql = query
	return tigb.sqlScan(ctx, v)
}

func (tigb *TransportInstanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tigb.fields {
		if !transportinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tigb *TransportInstanceGroupBy) sqlQuery() *sql.Selector {
	selector := tigb.sql.Select()
	aggregation := make([]string, 0, len(tigb.fns))
	for _, fn := range tigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tigb.fields)+len(tigb.fns))
		for _, f := range tigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tigb.fields...)...)
}

// TransportInstanceSelect is the builder for selecting fields of TransportInstance entities.
type TransportInstanceSelect struct {
	*TransportInstanceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tis *TransportInstanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tis.prepareQuery(ctx); err != nil {
		return err
	}
	tis.sql = tis.TransportInstanceQuery.sqlQuery(ctx)
	return tis.sqlScan(ctx, v)
}

func (tis *TransportInstanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tis.sql.Query()
	if err := tis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}