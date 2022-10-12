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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbapp"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessagefields"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/google/uuid"
)

// DbMessageQuery is the builder for querying DbMessage entities.
type DbMessageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DbMessage
	withTenant *TenantQuery
	withFields *DbMessageFieldsQuery
	withApp    *DbAppQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DbMessageQuery builder.
func (dmq *DbMessageQuery) Where(ps ...predicate.DbMessage) *DbMessageQuery {
	dmq.predicates = append(dmq.predicates, ps...)
	return dmq
}

// Limit adds a limit step to the query.
func (dmq *DbMessageQuery) Limit(limit int) *DbMessageQuery {
	dmq.limit = &limit
	return dmq
}

// Offset adds an offset step to the query.
func (dmq *DbMessageQuery) Offset(offset int) *DbMessageQuery {
	dmq.offset = &offset
	return dmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dmq *DbMessageQuery) Unique(unique bool) *DbMessageQuery {
	dmq.unique = &unique
	return dmq
}

// Order adds an order step to the query.
func (dmq *DbMessageQuery) Order(o ...OrderFunc) *DbMessageQuery {
	dmq.order = append(dmq.order, o...)
	return dmq
}

// QueryTenant chains the current query on the "tenant" edge.
func (dmq *DbMessageQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: dmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbmessage.Table, dbmessage.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, dbmessage.TenantTable, dbmessage.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFields chains the current query on the "fields" edge.
func (dmq *DbMessageQuery) QueryFields() *DbMessageFieldsQuery {
	query := &DbMessageFieldsQuery{config: dmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbmessage.Table, dbmessage.FieldID, selector),
			sqlgraph.To(dbmessagefields.Table, dbmessagefields.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dbmessage.FieldsTable, dbmessage.FieldsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApp chains the current query on the "app" edge.
func (dmq *DbMessageQuery) QueryApp() *DbAppQuery {
	query := &DbAppQuery{config: dmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbmessage.Table, dbmessage.FieldID, selector),
			sqlgraph.To(dbapp.Table, dbapp.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dbmessage.AppTable, dbmessage.AppColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DbMessage entity from the query.
// Returns a *NotFoundError when no DbMessage was found.
func (dmq *DbMessageQuery) First(ctx context.Context) (*DbMessage, error) {
	nodes, err := dmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dbmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dmq *DbMessageQuery) FirstX(ctx context.Context) *DbMessage {
	node, err := dmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DbMessage ID from the query.
// Returns a *NotFoundError when no DbMessage ID was found.
func (dmq *DbMessageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dbmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dmq *DbMessageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DbMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DbMessage entity is found.
// Returns a *NotFoundError when no DbMessage entities are found.
func (dmq *DbMessageQuery) Only(ctx context.Context) (*DbMessage, error) {
	nodes, err := dmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dbmessage.Label}
	default:
		return nil, &NotSingularError{dbmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dmq *DbMessageQuery) OnlyX(ctx context.Context) *DbMessage {
	node, err := dmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DbMessage ID in the query.
// Returns a *NotSingularError when more than one DbMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (dmq *DbMessageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dbmessage.Label}
	default:
		err = &NotSingularError{dbmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dmq *DbMessageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DbMessages.
func (dmq *DbMessageQuery) All(ctx context.Context) ([]*DbMessage, error) {
	if err := dmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dmq *DbMessageQuery) AllX(ctx context.Context) []*DbMessage {
	nodes, err := dmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DbMessage IDs.
func (dmq *DbMessageQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dmq.Select(dbmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dmq *DbMessageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dmq *DbMessageQuery) Count(ctx context.Context) (int, error) {
	if err := dmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dmq *DbMessageQuery) CountX(ctx context.Context) int {
	count, err := dmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dmq *DbMessageQuery) Exist(ctx context.Context) (bool, error) {
	if err := dmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dmq *DbMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := dmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DbMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dmq *DbMessageQuery) Clone() *DbMessageQuery {
	if dmq == nil {
		return nil
	}
	return &DbMessageQuery{
		config:     dmq.config,
		limit:      dmq.limit,
		offset:     dmq.offset,
		order:      append([]OrderFunc{}, dmq.order...),
		predicates: append([]predicate.DbMessage{}, dmq.predicates...),
		withTenant: dmq.withTenant.Clone(),
		withFields: dmq.withFields.Clone(),
		withApp:    dmq.withApp.Clone(),
		// clone intermediate query.
		sql:    dmq.sql.Clone(),
		path:   dmq.path,
		unique: dmq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (dmq *DbMessageQuery) WithTenant(opts ...func(*TenantQuery)) *DbMessageQuery {
	query := &TenantQuery{config: dmq.config}
	for _, opt := range opts {
		opt(query)
	}
	dmq.withTenant = query
	return dmq
}

// WithFields tells the query-builder to eager-load the nodes that are connected to
// the "fields" edge. The optional arguments are used to configure the query builder of the edge.
func (dmq *DbMessageQuery) WithFields(opts ...func(*DbMessageFieldsQuery)) *DbMessageQuery {
	query := &DbMessageFieldsQuery{config: dmq.config}
	for _, opt := range opts {
		opt(query)
	}
	dmq.withFields = query
	return dmq
}

// WithApp tells the query-builder to eager-load the nodes that are connected to
// the "app" edge. The optional arguments are used to configure the query builder of the edge.
func (dmq *DbMessageQuery) WithApp(opts ...func(*DbAppQuery)) *DbMessageQuery {
	query := &DbAppQuery{config: dmq.config}
	for _, opt := range opts {
		opt(query)
	}
	dmq.withApp = query
	return dmq
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
//	client.DbMessage.Query().
//		GroupBy(dbmessage.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dmq *DbMessageQuery) GroupBy(field string, fields ...string) *DbMessageGroupBy {
	grbuild := &DbMessageGroupBy{config: dmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dmq.sqlQuery(ctx), nil
	}
	grbuild.label = dbmessage.Label
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
//	client.DbMessage.Query().
//		Select(dbmessage.FieldTenantID).
//		Scan(ctx, &v)
func (dmq *DbMessageQuery) Select(fields ...string) *DbMessageSelect {
	dmq.fields = append(dmq.fields, fields...)
	selbuild := &DbMessageSelect{DbMessageQuery: dmq}
	selbuild.label = dbmessage.Label
	selbuild.flds, selbuild.scan = &dmq.fields, selbuild.Scan
	return selbuild
}

func (dmq *DbMessageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dmq.fields {
		if !dbmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dmq.path != nil {
		prev, err := dmq.path(ctx)
		if err != nil {
			return err
		}
		dmq.sql = prev
	}
	if dbmessage.Policy == nil {
		return errors.New("ent: uninitialized dbmessage.Policy (forgotten import ent/runtime?)")
	}
	if err := dbmessage.Policy.EvalQuery(ctx, dmq); err != nil {
		return err
	}
	return nil
}

func (dmq *DbMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DbMessage, error) {
	var (
		nodes       = []*DbMessage{}
		withFKs     = dmq.withFKs
		_spec       = dmq.querySpec()
		loadedTypes = [3]bool{
			dmq.withTenant != nil,
			dmq.withFields != nil,
			dmq.withApp != nil,
		}
	)
	if dmq.withTenant != nil || dmq.withApp != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dbmessage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DbMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DbMessage{config: dmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dmq.withTenant; query != nil {
		if err := dmq.loadTenant(ctx, query, nodes, nil,
			func(n *DbMessage, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := dmq.withFields; query != nil {
		if err := dmq.loadFields(ctx, query, nodes,
			func(n *DbMessage) { n.Edges.Fields = []*DbMessageFields{} },
			func(n *DbMessage, e *DbMessageFields) { n.Edges.Fields = append(n.Edges.Fields, e) }); err != nil {
			return nil, err
		}
	}
	if query := dmq.withApp; query != nil {
		if err := dmq.loadApp(ctx, query, nodes, nil,
			func(n *DbMessage, e *DbApp) { n.Edges.App = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dmq *DbMessageQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*DbMessage, init func(*DbMessage), assign func(*DbMessage, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DbMessage)
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
func (dmq *DbMessageQuery) loadFields(ctx context.Context, query *DbMessageFieldsQuery, nodes []*DbMessage, init func(*DbMessage), assign func(*DbMessage, *DbMessageFields)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*DbMessage)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DbMessageFields(func(s *sql.Selector) {
		s.Where(sql.InValues(dbmessage.FieldsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.db_message_fields
		if fk == nil {
			return fmt.Errorf(`foreign-key "db_message_fields" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "db_message_fields" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dmq *DbMessageQuery) loadApp(ctx context.Context, query *DbAppQuery, nodes []*DbMessage, init func(*DbMessage), assign func(*DbMessage, *DbApp)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DbMessage)
	for i := range nodes {
		if nodes[i].db_app_messages == nil {
			continue
		}
		fk := *nodes[i].db_app_messages
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(dbapp.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "db_app_messages" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dmq *DbMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dmq.querySpec()
	_spec.Node.Columns = dmq.fields
	if len(dmq.fields) > 0 {
		_spec.Unique = dmq.unique != nil && *dmq.unique
	}
	return sqlgraph.CountNodes(ctx, dmq.driver, _spec)
}

func (dmq *DbMessageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dmq *DbMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbmessage.Table,
			Columns: dbmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dbmessage.FieldID,
			},
		},
		From:   dmq.sql,
		Unique: true,
	}
	if unique := dmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbmessage.FieldID)
		for i := range fields {
			if fields[i] != dbmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dmq *DbMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dmq.driver.Dialect())
	t1 := builder.Table(dbmessage.Table)
	columns := dmq.fields
	if len(columns) == 0 {
		columns = dbmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dmq.sql != nil {
		selector = dmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dmq.unique != nil && *dmq.unique {
		selector.Distinct()
	}
	for _, p := range dmq.predicates {
		p(selector)
	}
	for _, p := range dmq.order {
		p(selector)
	}
	if offset := dmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DbMessageGroupBy is the group-by builder for DbMessage entities.
type DbMessageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dmgb *DbMessageGroupBy) Aggregate(fns ...AggregateFunc) *DbMessageGroupBy {
	dmgb.fns = append(dmgb.fns, fns...)
	return dmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dmgb *DbMessageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dmgb.path(ctx)
	if err != nil {
		return err
	}
	dmgb.sql = query
	return dmgb.sqlScan(ctx, v)
}

func (dmgb *DbMessageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dmgb.fields {
		if !dbmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dmgb *DbMessageGroupBy) sqlQuery() *sql.Selector {
	selector := dmgb.sql.Select()
	aggregation := make([]string, 0, len(dmgb.fns))
	for _, fn := range dmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dmgb.fields)+len(dmgb.fns))
		for _, f := range dmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dmgb.fields...)...)
}

// DbMessageSelect is the builder for selecting fields of DbMessage entities.
type DbMessageSelect struct {
	*DbMessageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dms *DbMessageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dms.prepareQuery(ctx); err != nil {
		return err
	}
	dms.sql = dms.DbMessageQuery.sqlQuery(ctx)
	return dms.sqlScan(ctx, v)
}

func (dms *DbMessageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dms.sql.Query()
	if err := dms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
