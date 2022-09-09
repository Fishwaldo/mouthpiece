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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/app"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/group"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/message"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportrecipient"
)

// AppQuery is the builder for querying App entities.
type AppQuery struct {
	config
	limit                   *int
	offset                  *int
	unique                  *bool
	order                   []OrderFunc
	fields                  []string
	predicates              []predicate.App
	withTenant              *TenantQuery
	withMessages            *MessageQuery
	withFilters             *FilterQuery
	withGroups              *GroupQuery
	withTransportRecipients *TransportRecipientQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppQuery builder.
func (aq *AppQuery) Where(ps ...predicate.App) *AppQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AppQuery) Limit(limit int) *AppQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AppQuery) Offset(offset int) *AppQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AppQuery) Unique(unique bool) *AppQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AppQuery) Order(o ...OrderFunc) *AppQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryTenant chains the current query on the "tenant" edge.
func (aq *AppQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, app.TenantTable, app.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMessages chains the current query on the "messages" edge.
func (aq *AppQuery) QueryMessages() *MessageQuery {
	query := &MessageQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, selector),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, app.MessagesTable, app.MessagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFilters chains the current query on the "filters" edge.
func (aq *AppQuery) QueryFilters() *FilterQuery {
	query := &FilterQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, selector),
			sqlgraph.To(filter.Table, filter.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, app.FiltersTable, app.FiltersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroups chains the current query on the "groups" edge.
func (aq *AppQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, app.GroupsTable, app.GroupsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransportRecipients chains the current query on the "TransportRecipients" edge.
func (aq *AppQuery) QueryTransportRecipients() *TransportRecipientQuery {
	query := &TransportRecipientQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, selector),
			sqlgraph.To(transportrecipient.Table, transportrecipient.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, app.TransportRecipientsTable, app.TransportRecipientsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first App entity from the query.
// Returns a *NotFoundError when no App was found.
func (aq *AppQuery) First(ctx context.Context) (*App, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{app.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AppQuery) FirstX(ctx context.Context) *App {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first App ID from the query.
// Returns a *NotFoundError when no App ID was found.
func (aq *AppQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{app.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AppQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single App entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one App entity is found.
// Returns a *NotFoundError when no App entities are found.
func (aq *AppQuery) Only(ctx context.Context) (*App, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{app.Label}
	default:
		return nil, &NotSingularError{app.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AppQuery) OnlyX(ctx context.Context) *App {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only App ID in the query.
// Returns a *NotSingularError when more than one App ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AppQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{app.Label}
	default:
		err = &NotSingularError{app.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AppQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Apps.
func (aq *AppQuery) All(ctx context.Context) ([]*App, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AppQuery) AllX(ctx context.Context) []*App {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of App IDs.
func (aq *AppQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(app.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AppQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AppQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AppQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AppQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AppQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AppQuery) Clone() *AppQuery {
	if aq == nil {
		return nil
	}
	return &AppQuery{
		config:                  aq.config,
		limit:                   aq.limit,
		offset:                  aq.offset,
		order:                   append([]OrderFunc{}, aq.order...),
		predicates:              append([]predicate.App{}, aq.predicates...),
		withTenant:              aq.withTenant.Clone(),
		withMessages:            aq.withMessages.Clone(),
		withFilters:             aq.withFilters.Clone(),
		withGroups:              aq.withGroups.Clone(),
		withTransportRecipients: aq.withTransportRecipients.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppQuery) WithTenant(opts ...func(*TenantQuery)) *AppQuery {
	query := &TenantQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withTenant = query
	return aq
}

// WithMessages tells the query-builder to eager-load the nodes that are connected to
// the "messages" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppQuery) WithMessages(opts ...func(*MessageQuery)) *AppQuery {
	query := &MessageQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withMessages = query
	return aq
}

// WithFilters tells the query-builder to eager-load the nodes that are connected to
// the "filters" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppQuery) WithFilters(opts ...func(*FilterQuery)) *AppQuery {
	query := &FilterQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withFilters = query
	return aq
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppQuery) WithGroups(opts ...func(*GroupQuery)) *AppQuery {
	query := &GroupQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withGroups = query
	return aq
}

// WithTransportRecipients tells the query-builder to eager-load the nodes that are connected to
// the "TransportRecipients" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppQuery) WithTransportRecipients(opts ...func(*TransportRecipientQuery)) *AppQuery {
	query := &TransportRecipientQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withTransportRecipients = query
	return aq
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
//	client.App.Query().
//		GroupBy(app.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AppQuery) GroupBy(field string, fields ...string) *AppGroupBy {
	grbuild := &AppGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = app.Label
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
//	client.App.Query().
//		Select(app.FieldTenantID).
//		Scan(ctx, &v)
//
func (aq *AppQuery) Select(fields ...string) *AppSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AppSelect{AppQuery: aq}
	selbuild.label = app.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AppQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !app.ValidColumn(f) {
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
	if app.Policy == nil {
		return errors.New("ent: uninitialized app.Policy (forgotten import ent/runtime?)")
	}
	if err := app.Policy.EvalQuery(ctx, aq); err != nil {
		return err
	}
	return nil
}

func (aq *AppQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*App, error) {
	var (
		nodes       = []*App{}
		_spec       = aq.querySpec()
		loadedTypes = [5]bool{
			aq.withTenant != nil,
			aq.withMessages != nil,
			aq.withFilters != nil,
			aq.withGroups != nil,
			aq.withTransportRecipients != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*App).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &App{config: aq.config}
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
	if query := aq.withTenant; query != nil {
		if err := aq.loadTenant(ctx, query, nodes, nil,
			func(n *App, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withMessages; query != nil {
		if err := aq.loadMessages(ctx, query, nodes,
			func(n *App) { n.Edges.Messages = []*Message{} },
			func(n *App, e *Message) { n.Edges.Messages = append(n.Edges.Messages, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withFilters; query != nil {
		if err := aq.loadFilters(ctx, query, nodes,
			func(n *App) { n.Edges.Filters = []*Filter{} },
			func(n *App, e *Filter) { n.Edges.Filters = append(n.Edges.Filters, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withGroups; query != nil {
		if err := aq.loadGroups(ctx, query, nodes,
			func(n *App) { n.Edges.Groups = []*Group{} },
			func(n *App, e *Group) { n.Edges.Groups = append(n.Edges.Groups, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withTransportRecipients; query != nil {
		if err := aq.loadTransportRecipients(ctx, query, nodes,
			func(n *App) { n.Edges.TransportRecipients = []*TransportRecipient{} },
			func(n *App, e *TransportRecipient) {
				n.Edges.TransportRecipients = append(n.Edges.TransportRecipients, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AppQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*App, init func(*App), assign func(*App, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*App)
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
func (aq *AppQuery) loadMessages(ctx context.Context, query *MessageQuery, nodes []*App, init func(*App), assign func(*App, *Message)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*App)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Message(func(s *sql.Selector) {
		s.Where(sql.InValues(app.MessagesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.app_messages
		if fk == nil {
			return fmt.Errorf(`foreign-key "app_messages" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "app_messages" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AppQuery) loadFilters(ctx context.Context, query *FilterQuery, nodes []*App, init func(*App), assign func(*App, *Filter)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*App)
	nids := make(map[int]map[*App]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(app.FiltersTable)
		s.Join(joinT).On(s.C(filter.FieldID), joinT.C(app.FiltersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(app.FiltersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(app.FiltersPrimaryKey[0]))
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
				nids[inValue] = map[*App]struct{}{byID[outValue]: struct{}{}}
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
func (aq *AppQuery) loadGroups(ctx context.Context, query *GroupQuery, nodes []*App, init func(*App), assign func(*App, *Group)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*App)
	nids := make(map[int]map[*App]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(app.GroupsTable)
		s.Join(joinT).On(s.C(group.FieldID), joinT.C(app.GroupsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(app.GroupsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(app.GroupsPrimaryKey[0]))
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
				nids[inValue] = map[*App]struct{}{byID[outValue]: struct{}{}}
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
func (aq *AppQuery) loadTransportRecipients(ctx context.Context, query *TransportRecipientQuery, nodes []*App, init func(*App), assign func(*App, *TransportRecipient)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*App)
	nids := make(map[int]map[*App]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(app.TransportRecipientsTable)
		s.Join(joinT).On(s.C(transportrecipient.FieldID), joinT.C(app.TransportRecipientsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(app.TransportRecipientsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(app.TransportRecipientsPrimaryKey[0]))
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
				nids[inValue] = map[*App]struct{}{byID[outValue]: struct{}{}}
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
			return fmt.Errorf(`unexpected "TransportRecipients" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aq *AppQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AppQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AppQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, app.FieldID)
		for i := range fields {
			if fields[i] != app.FieldID {
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
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
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

func (aq *AppQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(app.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = app.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppGroupBy is the group-by builder for App entities.
type AppGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AppGroupBy) Aggregate(fns ...AggregateFunc) *AppGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AppGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AppGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !app.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AppGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AppSelect is the builder for selecting fields of App entities.
type AppSelect struct {
	*AppQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AppSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AppQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AppSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
