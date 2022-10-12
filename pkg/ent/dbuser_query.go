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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbgroup"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbusermetadata"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
)

// DbUserQuery is the builder for querying DbUser entities.
type DbUserQuery struct {
	config
	limit                   *int
	offset                  *int
	unique                  *bool
	order                   []OrderFunc
	fields                  []string
	predicates              []predicate.DbUser
	withTenant              *TenantQuery
	withMetadata            *DbUserMetaDataQuery
	withFilters             *DbFilterQuery
	withGroups              *DbGroupQuery
	withTransportRecipients *DbTransportRecipientsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DbUserQuery builder.
func (duq *DbUserQuery) Where(ps ...predicate.DbUser) *DbUserQuery {
	duq.predicates = append(duq.predicates, ps...)
	return duq
}

// Limit adds a limit step to the query.
func (duq *DbUserQuery) Limit(limit int) *DbUserQuery {
	duq.limit = &limit
	return duq
}

// Offset adds an offset step to the query.
func (duq *DbUserQuery) Offset(offset int) *DbUserQuery {
	duq.offset = &offset
	return duq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (duq *DbUserQuery) Unique(unique bool) *DbUserQuery {
	duq.unique = &unique
	return duq
}

// Order adds an order step to the query.
func (duq *DbUserQuery) Order(o ...OrderFunc) *DbUserQuery {
	duq.order = append(duq.order, o...)
	return duq
}

// QueryTenant chains the current query on the "tenant" edge.
func (duq *DbUserQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: duq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := duq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := duq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbuser.Table, dbuser.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, dbuser.TenantTable, dbuser.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(duq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMetadata chains the current query on the "metadata" edge.
func (duq *DbUserQuery) QueryMetadata() *DbUserMetaDataQuery {
	query := &DbUserMetaDataQuery{config: duq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := duq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := duq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbuser.Table, dbuser.FieldID, selector),
			sqlgraph.To(dbusermetadata.Table, dbusermetadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dbuser.MetadataTable, dbuser.MetadataColumn),
		)
		fromU = sqlgraph.SetNeighbors(duq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFilters chains the current query on the "filters" edge.
func (duq *DbUserQuery) QueryFilters() *DbFilterQuery {
	query := &DbFilterQuery{config: duq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := duq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := duq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbuser.Table, dbuser.FieldID, selector),
			sqlgraph.To(dbfilter.Table, dbfilter.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, dbuser.FiltersTable, dbuser.FiltersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(duq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroups chains the current query on the "groups" edge.
func (duq *DbUserQuery) QueryGroups() *DbGroupQuery {
	query := &DbGroupQuery{config: duq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := duq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := duq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbuser.Table, dbuser.FieldID, selector),
			sqlgraph.To(dbgroup.Table, dbgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, dbuser.GroupsTable, dbuser.GroupsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(duq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransportRecipients chains the current query on the "TransportRecipients" edge.
func (duq *DbUserQuery) QueryTransportRecipients() *DbTransportRecipientsQuery {
	query := &DbTransportRecipientsQuery{config: duq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := duq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := duq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dbuser.Table, dbuser.FieldID, selector),
			sqlgraph.To(dbtransportrecipients.Table, dbtransportrecipients.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dbuser.TransportRecipientsTable, dbuser.TransportRecipientsColumn),
		)
		fromU = sqlgraph.SetNeighbors(duq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DbUser entity from the query.
// Returns a *NotFoundError when no DbUser was found.
func (duq *DbUserQuery) First(ctx context.Context) (*DbUser, error) {
	nodes, err := duq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dbuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (duq *DbUserQuery) FirstX(ctx context.Context) *DbUser {
	node, err := duq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DbUser ID from the query.
// Returns a *NotFoundError when no DbUser ID was found.
func (duq *DbUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = duq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dbuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (duq *DbUserQuery) FirstIDX(ctx context.Context) int {
	id, err := duq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DbUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DbUser entity is found.
// Returns a *NotFoundError when no DbUser entities are found.
func (duq *DbUserQuery) Only(ctx context.Context) (*DbUser, error) {
	nodes, err := duq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dbuser.Label}
	default:
		return nil, &NotSingularError{dbuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (duq *DbUserQuery) OnlyX(ctx context.Context) *DbUser {
	node, err := duq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DbUser ID in the query.
// Returns a *NotSingularError when more than one DbUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (duq *DbUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = duq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dbuser.Label}
	default:
		err = &NotSingularError{dbuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (duq *DbUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := duq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DbUsers.
func (duq *DbUserQuery) All(ctx context.Context) ([]*DbUser, error) {
	if err := duq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return duq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (duq *DbUserQuery) AllX(ctx context.Context) []*DbUser {
	nodes, err := duq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DbUser IDs.
func (duq *DbUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := duq.Select(dbuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (duq *DbUserQuery) IDsX(ctx context.Context) []int {
	ids, err := duq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (duq *DbUserQuery) Count(ctx context.Context) (int, error) {
	if err := duq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return duq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (duq *DbUserQuery) CountX(ctx context.Context) int {
	count, err := duq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (duq *DbUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := duq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return duq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (duq *DbUserQuery) ExistX(ctx context.Context) bool {
	exist, err := duq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DbUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (duq *DbUserQuery) Clone() *DbUserQuery {
	if duq == nil {
		return nil
	}
	return &DbUserQuery{
		config:                  duq.config,
		limit:                   duq.limit,
		offset:                  duq.offset,
		order:                   append([]OrderFunc{}, duq.order...),
		predicates:              append([]predicate.DbUser{}, duq.predicates...),
		withTenant:              duq.withTenant.Clone(),
		withMetadata:            duq.withMetadata.Clone(),
		withFilters:             duq.withFilters.Clone(),
		withGroups:              duq.withGroups.Clone(),
		withTransportRecipients: duq.withTransportRecipients.Clone(),
		// clone intermediate query.
		sql:    duq.sql.Clone(),
		path:   duq.path,
		unique: duq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (duq *DbUserQuery) WithTenant(opts ...func(*TenantQuery)) *DbUserQuery {
	query := &TenantQuery{config: duq.config}
	for _, opt := range opts {
		opt(query)
	}
	duq.withTenant = query
	return duq
}

// WithMetadata tells the query-builder to eager-load the nodes that are connected to
// the "metadata" edge. The optional arguments are used to configure the query builder of the edge.
func (duq *DbUserQuery) WithMetadata(opts ...func(*DbUserMetaDataQuery)) *DbUserQuery {
	query := &DbUserMetaDataQuery{config: duq.config}
	for _, opt := range opts {
		opt(query)
	}
	duq.withMetadata = query
	return duq
}

// WithFilters tells the query-builder to eager-load the nodes that are connected to
// the "filters" edge. The optional arguments are used to configure the query builder of the edge.
func (duq *DbUserQuery) WithFilters(opts ...func(*DbFilterQuery)) *DbUserQuery {
	query := &DbFilterQuery{config: duq.config}
	for _, opt := range opts {
		opt(query)
	}
	duq.withFilters = query
	return duq
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (duq *DbUserQuery) WithGroups(opts ...func(*DbGroupQuery)) *DbUserQuery {
	query := &DbGroupQuery{config: duq.config}
	for _, opt := range opts {
		opt(query)
	}
	duq.withGroups = query
	return duq
}

// WithTransportRecipients tells the query-builder to eager-load the nodes that are connected to
// the "TransportRecipients" edge. The optional arguments are used to configure the query builder of the edge.
func (duq *DbUserQuery) WithTransportRecipients(opts ...func(*DbTransportRecipientsQuery)) *DbUserQuery {
	query := &DbTransportRecipientsQuery{config: duq.config}
	for _, opt := range opts {
		opt(query)
	}
	duq.withTransportRecipients = query
	return duq
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
//	client.DbUser.Query().
//		GroupBy(dbuser.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (duq *DbUserQuery) GroupBy(field string, fields ...string) *DbUserGroupBy {
	grbuild := &DbUserGroupBy{config: duq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := duq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return duq.sqlQuery(ctx), nil
	}
	grbuild.label = dbuser.Label
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
//	client.DbUser.Query().
//		Select(dbuser.FieldTenantID).
//		Scan(ctx, &v)
func (duq *DbUserQuery) Select(fields ...string) *DbUserSelect {
	duq.fields = append(duq.fields, fields...)
	selbuild := &DbUserSelect{DbUserQuery: duq}
	selbuild.label = dbuser.Label
	selbuild.flds, selbuild.scan = &duq.fields, selbuild.Scan
	return selbuild
}

func (duq *DbUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range duq.fields {
		if !dbuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if duq.path != nil {
		prev, err := duq.path(ctx)
		if err != nil {
			return err
		}
		duq.sql = prev
	}
	if dbuser.Policy == nil {
		return errors.New("ent: uninitialized dbuser.Policy (forgotten import ent/runtime?)")
	}
	if err := dbuser.Policy.EvalQuery(ctx, duq); err != nil {
		return err
	}
	return nil
}

func (duq *DbUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DbUser, error) {
	var (
		nodes       = []*DbUser{}
		_spec       = duq.querySpec()
		loadedTypes = [5]bool{
			duq.withTenant != nil,
			duq.withMetadata != nil,
			duq.withFilters != nil,
			duq.withGroups != nil,
			duq.withTransportRecipients != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DbUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DbUser{config: duq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, duq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := duq.withTenant; query != nil {
		if err := duq.loadTenant(ctx, query, nodes, nil,
			func(n *DbUser, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := duq.withMetadata; query != nil {
		if err := duq.loadMetadata(ctx, query, nodes,
			func(n *DbUser) { n.Edges.Metadata = []*DbUserMetaData{} },
			func(n *DbUser, e *DbUserMetaData) { n.Edges.Metadata = append(n.Edges.Metadata, e) }); err != nil {
			return nil, err
		}
	}
	if query := duq.withFilters; query != nil {
		if err := duq.loadFilters(ctx, query, nodes,
			func(n *DbUser) { n.Edges.Filters = []*DbFilter{} },
			func(n *DbUser, e *DbFilter) { n.Edges.Filters = append(n.Edges.Filters, e) }); err != nil {
			return nil, err
		}
	}
	if query := duq.withGroups; query != nil {
		if err := duq.loadGroups(ctx, query, nodes,
			func(n *DbUser) { n.Edges.Groups = []*DbGroup{} },
			func(n *DbUser, e *DbGroup) { n.Edges.Groups = append(n.Edges.Groups, e) }); err != nil {
			return nil, err
		}
	}
	if query := duq.withTransportRecipients; query != nil {
		if err := duq.loadTransportRecipients(ctx, query, nodes,
			func(n *DbUser) { n.Edges.TransportRecipients = []*DbTransportRecipients{} },
			func(n *DbUser, e *DbTransportRecipients) {
				n.Edges.TransportRecipients = append(n.Edges.TransportRecipients, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (duq *DbUserQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*DbUser, init func(*DbUser), assign func(*DbUser, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DbUser)
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
func (duq *DbUserQuery) loadMetadata(ctx context.Context, query *DbUserMetaDataQuery, nodes []*DbUser, init func(*DbUser), assign func(*DbUser, *DbUserMetaData)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DbUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DbUserMetaData(func(s *sql.Selector) {
		s.Where(sql.InValues(dbuser.MetadataColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.db_user_metadata
		if fk == nil {
			return fmt.Errorf(`foreign-key "db_user_metadata" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "db_user_metadata" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (duq *DbUserQuery) loadFilters(ctx context.Context, query *DbFilterQuery, nodes []*DbUser, init func(*DbUser), assign func(*DbUser, *DbFilter)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DbUser)
	nids := make(map[int]map[*DbUser]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(dbuser.FiltersTable)
		s.Join(joinT).On(s.C(dbfilter.FieldID), joinT.C(dbuser.FiltersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(dbuser.FiltersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(dbuser.FiltersPrimaryKey[0]))
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
				nids[inValue] = map[*DbUser]struct{}{byID[outValue]: struct{}{}}
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
func (duq *DbUserQuery) loadGroups(ctx context.Context, query *DbGroupQuery, nodes []*DbUser, init func(*DbUser), assign func(*DbUser, *DbGroup)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DbUser)
	nids := make(map[int]map[*DbUser]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(dbuser.GroupsTable)
		s.Join(joinT).On(s.C(dbgroup.FieldID), joinT.C(dbuser.GroupsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(dbuser.GroupsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(dbuser.GroupsPrimaryKey[0]))
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
				nids[inValue] = map[*DbUser]struct{}{byID[outValue]: struct{}{}}
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
			return fmt.Errorf(`unexpected "groups" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (duq *DbUserQuery) loadTransportRecipients(ctx context.Context, query *DbTransportRecipientsQuery, nodes []*DbUser, init func(*DbUser), assign func(*DbUser, *DbTransportRecipients)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DbUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DbTransportRecipients(func(s *sql.Selector) {
		s.Where(sql.InValues(dbuser.TransportRecipientsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.db_user_transport_recipients
		if fk == nil {
			return fmt.Errorf(`foreign-key "db_user_transport_recipients" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "db_user_transport_recipients" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (duq *DbUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := duq.querySpec()
	_spec.Node.Columns = duq.fields
	if len(duq.fields) > 0 {
		_spec.Unique = duq.unique != nil && *duq.unique
	}
	return sqlgraph.CountNodes(ctx, duq.driver, _spec)
}

func (duq *DbUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := duq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (duq *DbUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbuser.Table,
			Columns: dbuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbuser.FieldID,
			},
		},
		From:   duq.sql,
		Unique: true,
	}
	if unique := duq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := duq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbuser.FieldID)
		for i := range fields {
			if fields[i] != dbuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := duq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := duq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := duq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := duq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (duq *DbUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(duq.driver.Dialect())
	t1 := builder.Table(dbuser.Table)
	columns := duq.fields
	if len(columns) == 0 {
		columns = dbuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if duq.sql != nil {
		selector = duq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if duq.unique != nil && *duq.unique {
		selector.Distinct()
	}
	for _, p := range duq.predicates {
		p(selector)
	}
	for _, p := range duq.order {
		p(selector)
	}
	if offset := duq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := duq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DbUserGroupBy is the group-by builder for DbUser entities.
type DbUserGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dugb *DbUserGroupBy) Aggregate(fns ...AggregateFunc) *DbUserGroupBy {
	dugb.fns = append(dugb.fns, fns...)
	return dugb
}

// Scan applies the group-by query and scans the result into the given value.
func (dugb *DbUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dugb.path(ctx)
	if err != nil {
		return err
	}
	dugb.sql = query
	return dugb.sqlScan(ctx, v)
}

func (dugb *DbUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dugb.fields {
		if !dbuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dugb *DbUserGroupBy) sqlQuery() *sql.Selector {
	selector := dugb.sql.Select()
	aggregation := make([]string, 0, len(dugb.fns))
	for _, fn := range dugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dugb.fields)+len(dugb.fns))
		for _, f := range dugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dugb.fields...)...)
}

// DbUserSelect is the builder for selecting fields of DbUser entities.
type DbUserSelect struct {
	*DbUserQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dus *DbUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dus.prepareQuery(ctx); err != nil {
		return err
	}
	dus.sql = dus.DbUserQuery.sqlQuery(ctx)
	return dus.sqlScan(ctx, v)
}

func (dus *DbUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dus.sql.Query()
	if err := dus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}