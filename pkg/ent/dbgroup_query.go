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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbgroup"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
)

// DbGroupQuery is the builder for querying DbGroup entities.
type DbGroupQuery struct {
	config
	limit                   *int
	offset                  *int
	unique                  *bool
	order                   []OrderFunc
	fields                  []string
	predicates              []predicate.DbGroup
	withTenant              *TenantQuery
	withTransportRecipients *DbTransportRecipientsQuery
	withUsers               *DbUserQuery
	withFilters             *DbFilterQuery
	withApps                *DbAppQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DbGroupQuery builder.
func (dgq *DbGroupQuery) Where(ps ...predicate.DbGroup) *DbGroupQuery {
	dgq.predicates = append(dgq.predicates, ps...)
	return dgq
}

// Limit adds a limit step to the query.
func (dgq *DbGroupQuery) Limit(limit int) *DbGroupQuery {
	dgq.limit = &limit
	return dgq
}

// Offset adds an offset step to the query.
func (dgq *DbGroupQuery) Offset(offset int) *DbGroupQuery {
	dgq.offset = &offset
	return dgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dgq *DbGroupQuery) Unique(unique bool) *DbGroupQuery {
	dgq.unique = &unique
	return dgq
}

// Order adds an order step to the query.
func (dgq *DbGroupQuery) Order(o ...OrderFunc) *DbGroupQuery {
	dgq.order = append(dgq.order, o...)
	return dgq
}

// QueryTenant chains the current query on the "tenant" edge.
func (dgq *DbGroupQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: dgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbgroup.Table, dbgroup.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, dbgroup.TenantTable, dbgroup.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransportRecipients chains the current query on the "TransportRecipients" edge.
func (dgq *DbGroupQuery) QueryTransportRecipients() *DbTransportRecipientsQuery {
	query := &DbTransportRecipientsQuery{config: dgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbgroup.Table, dbgroup.FieldID, selector),
			sqlgraph.To(dbtransportrecipients.Table, dbtransportrecipients.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dbgroup.TransportRecipientsTable, dbgroup.TransportRecipientsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (dgq *DbGroupQuery) QueryUsers() *DbUserQuery {
	query := &DbUserQuery{config: dgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbgroup.Table, dbgroup.FieldID, selector),
			sqlgraph.To(dbuser.Table, dbuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, dbgroup.UsersTable, dbgroup.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFilters chains the current query on the "filters" edge.
func (dgq *DbGroupQuery) QueryFilters() *DbFilterQuery {
	query := &DbFilterQuery{config: dgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbgroup.Table, dbgroup.FieldID, selector),
			sqlgraph.To(dbfilter.Table, dbfilter.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, dbgroup.FiltersTable, dbgroup.FiltersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApps chains the current query on the "apps" edge.
func (dgq *DbGroupQuery) QueryApps() *DbAppQuery {
	query := &DbAppQuery{config: dgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbgroup.Table, dbgroup.FieldID, selector),
			sqlgraph.To(dbapp.Table, dbapp.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, dbgroup.AppsTable, dbgroup.AppsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DbGroup entity from the query.
// Returns a *NotFoundError when no DbGroup was found.
func (dgq *DbGroupQuery) First(ctx context.Context) (*DbGroup, error) {
	nodes, err := dgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dbgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dgq *DbGroupQuery) FirstX(ctx context.Context) *DbGroup {
	node, err := dgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DbGroup ID from the query.
// Returns a *NotFoundError when no DbGroup ID was found.
func (dgq *DbGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dbgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dgq *DbGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := dgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DbGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DbGroup entity is found.
// Returns a *NotFoundError when no DbGroup entities are found.
func (dgq *DbGroupQuery) Only(ctx context.Context) (*DbGroup, error) {
	nodes, err := dgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dbgroup.Label}
	default:
		return nil, &NotSingularError{dbgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dgq *DbGroupQuery) OnlyX(ctx context.Context) *DbGroup {
	node, err := dgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DbGroup ID in the query.
// Returns a *NotSingularError when more than one DbGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (dgq *DbGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dbgroup.Label}
	default:
		err = &NotSingularError{dbgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dgq *DbGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := dgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DbGroups.
func (dgq *DbGroupQuery) All(ctx context.Context) ([]*DbGroup, error) {
	if err := dgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dgq *DbGroupQuery) AllX(ctx context.Context) []*DbGroup {
	nodes, err := dgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DbGroup IDs.
func (dgq *DbGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dgq.Select(dbgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dgq *DbGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := dgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dgq *DbGroupQuery) Count(ctx context.Context) (int, error) {
	if err := dgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dgq *DbGroupQuery) CountX(ctx context.Context) int {
	count, err := dgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dgq *DbGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := dgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dgq *DbGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := dgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DbGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dgq *DbGroupQuery) Clone() *DbGroupQuery {
	if dgq == nil {
		return nil
	}
	return &DbGroupQuery{
		config:                  dgq.config,
		limit:                   dgq.limit,
		offset:                  dgq.offset,
		order:                   append([]OrderFunc{}, dgq.order...),
		predicates:              append([]predicate.DbGroup{}, dgq.predicates...),
		withTenant:              dgq.withTenant.Clone(),
		withTransportRecipients: dgq.withTransportRecipients.Clone(),
		withUsers:               dgq.withUsers.Clone(),
		withFilters:             dgq.withFilters.Clone(),
		withApps:                dgq.withApps.Clone(),
		// clone intermediate query.
		sql:    dgq.sql.Clone(),
		path:   dgq.path,
		unique: dgq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DbGroupQuery) WithTenant(opts ...func(*TenantQuery)) *DbGroupQuery {
	query := &TenantQuery{config: dgq.config}
	for _, opt := range opts {
		opt(query)
	}
	dgq.withTenant = query
	return dgq
}

// WithTransportRecipients tells the query-builder to eager-load the nodes that are connected to
// the "TransportRecipients" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DbGroupQuery) WithTransportRecipients(opts ...func(*DbTransportRecipientsQuery)) *DbGroupQuery {
	query := &DbTransportRecipientsQuery{config: dgq.config}
	for _, opt := range opts {
		opt(query)
	}
	dgq.withTransportRecipients = query
	return dgq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DbGroupQuery) WithUsers(opts ...func(*DbUserQuery)) *DbGroupQuery {
	query := &DbUserQuery{config: dgq.config}
	for _, opt := range opts {
		opt(query)
	}
	dgq.withUsers = query
	return dgq
}

// WithFilters tells the query-builder to eager-load the nodes that are connected to
// the "filters" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DbGroupQuery) WithFilters(opts ...func(*DbFilterQuery)) *DbGroupQuery {
	query := &DbFilterQuery{config: dgq.config}
	for _, opt := range opts {
		opt(query)
	}
	dgq.withFilters = query
	return dgq
}

// WithApps tells the query-builder to eager-load the nodes that are connected to
// the "apps" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DbGroupQuery) WithApps(opts ...func(*DbAppQuery)) *DbGroupQuery {
	query := &DbAppQuery{config: dgq.config}
	for _, opt := range opts {
		opt(query)
	}
	dgq.withApps = query
	return dgq
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
//	client.DbGroup.Query().
//		GroupBy(dbgroup.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dgq *DbGroupQuery) GroupBy(field string, fields ...string) *DbGroupGroupBy {
	grbuild := &DbGroupGroupBy{config: dgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dgq.sqlQuery(ctx), nil
	}
	grbuild.label = dbgroup.Label
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
//	client.DbGroup.Query().
//		Select(dbgroup.FieldTenantID).
//		Scan(ctx, &v)
func (dgq *DbGroupQuery) Select(fields ...string) *DbGroupSelect {
	dgq.fields = append(dgq.fields, fields...)
	selbuild := &DbGroupSelect{DbGroupQuery: dgq}
	selbuild.label = dbgroup.Label
	selbuild.flds, selbuild.scan = &dgq.fields, selbuild.Scan
	return selbuild
}

func (dgq *DbGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dgq.fields {
		if !dbgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dgq.path != nil {
		prev, err := dgq.path(ctx)
		if err != nil {
			return err
		}
		dgq.sql = prev
	}
	if dbgroup.Policy == nil {
		return errors.New("ent: uninitialized dbgroup.Policy (forgotten import ent/runtime?)")
	}
	if err := dbgroup.Policy.EvalQuery(ctx, dgq); err != nil {
		return err
	}
	return nil
}

func (dgq *DbGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DbGroup, error) {
	var (
		nodes       = []*DbGroup{}
		_spec       = dgq.querySpec()
		loadedTypes = [5]bool{
			dgq.withTenant != nil,
			dgq.withTransportRecipients != nil,
			dgq.withUsers != nil,
			dgq.withFilters != nil,
			dgq.withApps != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DbGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DbGroup{config: dgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dgq.withTenant; query != nil {
		if err := dgq.loadTenant(ctx, query, nodes, nil,
			func(n *DbGroup, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := dgq.withTransportRecipients; query != nil {
		if err := dgq.loadTransportRecipients(ctx, query, nodes,
			func(n *DbGroup) { n.Edges.TransportRecipients = []*DbTransportRecipients{} },
			func(n *DbGroup, e *DbTransportRecipients) {
				n.Edges.TransportRecipients = append(n.Edges.TransportRecipients, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := dgq.withUsers; query != nil {
		if err := dgq.loadUsers(ctx, query, nodes,
			func(n *DbGroup) { n.Edges.Users = []*DbUser{} },
			func(n *DbGroup, e *DbUser) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := dgq.withFilters; query != nil {
		if err := dgq.loadFilters(ctx, query, nodes,
			func(n *DbGroup) { n.Edges.Filters = []*DbFilter{} },
			func(n *DbGroup, e *DbFilter) { n.Edges.Filters = append(n.Edges.Filters, e) }); err != nil {
			return nil, err
		}
	}
	if query := dgq.withApps; query != nil {
		if err := dgq.loadApps(ctx, query, nodes,
			func(n *DbGroup) { n.Edges.Apps = []*DbApp{} },
			func(n *DbGroup, e *DbApp) { n.Edges.Apps = append(n.Edges.Apps, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dgq *DbGroupQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*DbGroup, init func(*DbGroup), assign func(*DbGroup, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DbGroup)
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
func (dgq *DbGroupQuery) loadTransportRecipients(ctx context.Context, query *DbTransportRecipientsQuery, nodes []*DbGroup, init func(*DbGroup), assign func(*DbGroup, *DbTransportRecipients)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DbGroup)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DbTransportRecipients(func(s *sql.Selector) {
		s.Where(sql.InValues(dbgroup.TransportRecipientsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.db_group_transport_recipients
		if fk == nil {
			return fmt.Errorf(`foreign-key "db_group_transport_recipients" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "db_group_transport_recipients" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dgq *DbGroupQuery) loadUsers(ctx context.Context, query *DbUserQuery, nodes []*DbGroup, init func(*DbGroup), assign func(*DbGroup, *DbUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DbGroup)
	nids := make(map[int]map[*DbGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(dbgroup.UsersTable)
		s.Join(joinT).On(s.C(dbuser.FieldID), joinT.C(dbgroup.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(dbgroup.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(dbgroup.UsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*DbGroup]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dgq *DbGroupQuery) loadFilters(ctx context.Context, query *DbFilterQuery, nodes []*DbGroup, init func(*DbGroup), assign func(*DbGroup, *DbFilter)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DbGroup)
	nids := make(map[int]map[*DbGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(dbgroup.FiltersTable)
		s.Join(joinT).On(s.C(dbfilter.FieldID), joinT.C(dbgroup.FiltersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(dbgroup.FiltersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(dbgroup.FiltersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*DbGroup]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "filters" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dgq *DbGroupQuery) loadApps(ctx context.Context, query *DbAppQuery, nodes []*DbGroup, init func(*DbGroup), assign func(*DbGroup, *DbApp)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DbGroup)
	nids := make(map[int]map[*DbGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(dbgroup.AppsTable)
		s.Join(joinT).On(s.C(dbapp.FieldID), joinT.C(dbgroup.AppsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(dbgroup.AppsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(dbgroup.AppsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*DbGroup]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "apps" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (dgq *DbGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dgq.querySpec()
	_spec.Node.Columns = dgq.fields
	if len(dgq.fields) > 0 {
		_spec.Unique = dgq.unique != nil && *dgq.unique
	}
	return sqlgraph.CountNodes(ctx, dgq.driver, _spec)
}

func (dgq *DbGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dgq *DbGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbgroup.Table,
			Columns: dbgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbgroup.FieldID,
			},
		},
		From:   dgq.sql,
		Unique: true,
	}
	if unique := dgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbgroup.FieldID)
		for i := range fields {
			if fields[i] != dbgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dgq *DbGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dgq.driver.Dialect())
	t1 := builder.Table(dbgroup.Table)
	columns := dgq.fields
	if len(columns) == 0 {
		columns = dbgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dgq.sql != nil {
		selector = dgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dgq.unique != nil && *dgq.unique {
		selector.Distinct()
	}
	for _, p := range dgq.predicates {
		p(selector)
	}
	for _, p := range dgq.order {
		p(selector)
	}
	if offset := dgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DbGroupGroupBy is the group-by builder for DbGroup entities.
type DbGroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dggb *DbGroupGroupBy) Aggregate(fns ...AggregateFunc) *DbGroupGroupBy {
	dggb.fns = append(dggb.fns, fns...)
	return dggb
}

// Scan applies the group-by query and scans the result into the given value.
func (dggb *DbGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dggb.path(ctx)
	if err != nil {
		return err
	}
	dggb.sql = query
	return dggb.sqlScan(ctx, v)
}

func (dggb *DbGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dggb.fields {
		if !dbgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dggb *DbGroupGroupBy) sqlQuery() *sql.Selector {
	selector := dggb.sql.Select()
	aggregation := make([]string, 0, len(dggb.fns))
	for _, fn := range dggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dggb.fields)+len(dggb.fns))
		for _, f := range dggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dggb.fields...)...)
}

// DbGroupSelect is the builder for selecting fields of DbGroup entities.
type DbGroupSelect struct {
	*DbGroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dgs *DbGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dgs.prepareQuery(ctx); err != nil {
		return err
	}
	dgs.sql = dgs.DbGroupQuery.sqlQuery(ctx)
	return dgs.sqlScan(ctx, v)
}

func (dgs *DbGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgs.sql.Query()
	if err := dgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}