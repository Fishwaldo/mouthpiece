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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/app"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/group"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportinstance"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportrecipient"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/user"
)

// TransportRecipientUpdate is the builder for updating TransportRecipient entities.
type TransportRecipientUpdate struct {
	config
	hooks    []Hook
	mutation *TransportRecipientMutation
}

// Where appends a list predicates to the TransportRecipientUpdate builder.
func (tru *TransportRecipientUpdate) Where(ps ...predicate.TransportRecipient) *TransportRecipientUpdate {
	tru.mutation.Where(ps...)
	return tru
}

// SetTenantID sets the "tenant_id" field.
func (tru *TransportRecipientUpdate) SetTenantID(i int) *TransportRecipientUpdate {
	tru.mutation.SetTenantID(i)
	return tru
}

// SetName sets the "Name" field.
func (tru *TransportRecipientUpdate) SetName(s string) *TransportRecipientUpdate {
	tru.mutation.SetName(s)
	return tru
}

// SetDescription sets the "Description" field.
func (tru *TransportRecipientUpdate) SetDescription(s string) *TransportRecipientUpdate {
	tru.mutation.SetDescription(s)
	return tru
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (tru *TransportRecipientUpdate) SetTenant(t *Tenant) *TransportRecipientUpdate {
	return tru.SetTenantID(t.ID)
}

// SetTransportInstanceID sets the "TransportInstance" edge to the TransportInstance entity by ID.
func (tru *TransportRecipientUpdate) SetTransportInstanceID(id int) *TransportRecipientUpdate {
	tru.mutation.SetTransportInstanceID(id)
	return tru
}

// SetTransportInstance sets the "TransportInstance" edge to the TransportInstance entity.
func (tru *TransportRecipientUpdate) SetTransportInstance(t *TransportInstance) *TransportRecipientUpdate {
	return tru.SetTransportInstanceID(t.ID)
}

// AddAppRecipientIDs adds the "AppRecipient" edge to the App entity by IDs.
func (tru *TransportRecipientUpdate) AddAppRecipientIDs(ids ...int) *TransportRecipientUpdate {
	tru.mutation.AddAppRecipientIDs(ids...)
	return tru
}

// AddAppRecipient adds the "AppRecipient" edges to the App entity.
func (tru *TransportRecipientUpdate) AddAppRecipient(a ...*App) *TransportRecipientUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tru.AddAppRecipientIDs(ids...)
}

// AddGroupRecipientIDs adds the "GroupRecipient" edge to the Group entity by IDs.
func (tru *TransportRecipientUpdate) AddGroupRecipientIDs(ids ...int) *TransportRecipientUpdate {
	tru.mutation.AddGroupRecipientIDs(ids...)
	return tru
}

// AddGroupRecipient adds the "GroupRecipient" edges to the Group entity.
func (tru *TransportRecipientUpdate) AddGroupRecipient(g ...*Group) *TransportRecipientUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tru.AddGroupRecipientIDs(ids...)
}

// AddUserRecipientIDs adds the "UserRecipient" edge to the User entity by IDs.
func (tru *TransportRecipientUpdate) AddUserRecipientIDs(ids ...int) *TransportRecipientUpdate {
	tru.mutation.AddUserRecipientIDs(ids...)
	return tru
}

// AddUserRecipient adds the "UserRecipient" edges to the User entity.
func (tru *TransportRecipientUpdate) AddUserRecipient(u ...*User) *TransportRecipientUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tru.AddUserRecipientIDs(ids...)
}

// Mutation returns the TransportRecipientMutation object of the builder.
func (tru *TransportRecipientUpdate) Mutation() *TransportRecipientMutation {
	return tru.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (tru *TransportRecipientUpdate) ClearTenant() *TransportRecipientUpdate {
	tru.mutation.ClearTenant()
	return tru
}

// ClearTransportInstance clears the "TransportInstance" edge to the TransportInstance entity.
func (tru *TransportRecipientUpdate) ClearTransportInstance() *TransportRecipientUpdate {
	tru.mutation.ClearTransportInstance()
	return tru
}

// ClearAppRecipient clears all "AppRecipient" edges to the App entity.
func (tru *TransportRecipientUpdate) ClearAppRecipient() *TransportRecipientUpdate {
	tru.mutation.ClearAppRecipient()
	return tru
}

// RemoveAppRecipientIDs removes the "AppRecipient" edge to App entities by IDs.
func (tru *TransportRecipientUpdate) RemoveAppRecipientIDs(ids ...int) *TransportRecipientUpdate {
	tru.mutation.RemoveAppRecipientIDs(ids...)
	return tru
}

// RemoveAppRecipient removes "AppRecipient" edges to App entities.
func (tru *TransportRecipientUpdate) RemoveAppRecipient(a ...*App) *TransportRecipientUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tru.RemoveAppRecipientIDs(ids...)
}

// ClearGroupRecipient clears all "GroupRecipient" edges to the Group entity.
func (tru *TransportRecipientUpdate) ClearGroupRecipient() *TransportRecipientUpdate {
	tru.mutation.ClearGroupRecipient()
	return tru
}

// RemoveGroupRecipientIDs removes the "GroupRecipient" edge to Group entities by IDs.
func (tru *TransportRecipientUpdate) RemoveGroupRecipientIDs(ids ...int) *TransportRecipientUpdate {
	tru.mutation.RemoveGroupRecipientIDs(ids...)
	return tru
}

// RemoveGroupRecipient removes "GroupRecipient" edges to Group entities.
func (tru *TransportRecipientUpdate) RemoveGroupRecipient(g ...*Group) *TransportRecipientUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tru.RemoveGroupRecipientIDs(ids...)
}

// ClearUserRecipient clears all "UserRecipient" edges to the User entity.
func (tru *TransportRecipientUpdate) ClearUserRecipient() *TransportRecipientUpdate {
	tru.mutation.ClearUserRecipient()
	return tru
}

// RemoveUserRecipientIDs removes the "UserRecipient" edge to User entities by IDs.
func (tru *TransportRecipientUpdate) RemoveUserRecipientIDs(ids ...int) *TransportRecipientUpdate {
	tru.mutation.RemoveUserRecipientIDs(ids...)
	return tru
}

// RemoveUserRecipient removes "UserRecipient" edges to User entities.
func (tru *TransportRecipientUpdate) RemoveUserRecipient(u ...*User) *TransportRecipientUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tru.RemoveUserRecipientIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TransportRecipientUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tru.hooks) == 0 {
		if err = tru.check(); err != nil {
			return 0, err
		}
		affected, err = tru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransportRecipientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tru.check(); err != nil {
				return 0, err
			}
			tru.mutation = mutation
			affected, err = tru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tru.hooks) - 1; i >= 0; i-- {
			if tru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TransportRecipientUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TransportRecipientUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TransportRecipientUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tru *TransportRecipientUpdate) check() error {
	if v, ok := tru.mutation.Name(); ok {
		if err := transportrecipient.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "TransportRecipient.Name": %w`, err)}
		}
	}
	if v, ok := tru.mutation.Description(); ok {
		if err := transportrecipient.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "Description", err: fmt.Errorf(`ent: validator failed for field "TransportRecipient.Description": %w`, err)}
		}
	}
	if _, ok := tru.mutation.TenantID(); tru.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TransportRecipient.tenant"`)
	}
	if _, ok := tru.mutation.TransportInstanceID(); tru.mutation.TransportInstanceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TransportRecipient.TransportInstance"`)
	}
	return nil
}

func (tru *TransportRecipientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transportrecipient.Table,
			Columns: transportrecipient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transportrecipient.FieldID,
			},
		},
	}
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transportrecipient.FieldName,
		})
	}
	if value, ok := tru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transportrecipient.FieldDescription,
		})
	}
	if tru.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transportrecipient.TenantTable,
			Columns: []string{transportrecipient.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transportrecipient.TenantTable,
			Columns: []string{transportrecipient.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.TransportInstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transportrecipient.TransportInstanceTable,
			Columns: []string{transportrecipient.TransportInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportinstance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TransportInstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transportrecipient.TransportInstanceTable,
			Columns: []string{transportrecipient.TransportInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.AppRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.AppRecipientTable,
			Columns: transportrecipient.AppRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.RemovedAppRecipientIDs(); len(nodes) > 0 && !tru.mutation.AppRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.AppRecipientTable,
			Columns: transportrecipient.AppRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.AppRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.AppRecipientTable,
			Columns: transportrecipient.AppRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.GroupRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.GroupRecipientTable,
			Columns: transportrecipient.GroupRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.RemovedGroupRecipientIDs(); len(nodes) > 0 && !tru.mutation.GroupRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.GroupRecipientTable,
			Columns: transportrecipient.GroupRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.GroupRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.GroupRecipientTable,
			Columns: transportrecipient.GroupRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.UserRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.UserRecipientTable,
			Columns: transportrecipient.UserRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.RemovedUserRecipientIDs(); len(nodes) > 0 && !tru.mutation.UserRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.UserRecipientTable,
			Columns: transportrecipient.UserRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.UserRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.UserRecipientTable,
			Columns: transportrecipient.UserRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transportrecipient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TransportRecipientUpdateOne is the builder for updating a single TransportRecipient entity.
type TransportRecipientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransportRecipientMutation
}

// SetTenantID sets the "tenant_id" field.
func (truo *TransportRecipientUpdateOne) SetTenantID(i int) *TransportRecipientUpdateOne {
	truo.mutation.SetTenantID(i)
	return truo
}

// SetName sets the "Name" field.
func (truo *TransportRecipientUpdateOne) SetName(s string) *TransportRecipientUpdateOne {
	truo.mutation.SetName(s)
	return truo
}

// SetDescription sets the "Description" field.
func (truo *TransportRecipientUpdateOne) SetDescription(s string) *TransportRecipientUpdateOne {
	truo.mutation.SetDescription(s)
	return truo
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (truo *TransportRecipientUpdateOne) SetTenant(t *Tenant) *TransportRecipientUpdateOne {
	return truo.SetTenantID(t.ID)
}

// SetTransportInstanceID sets the "TransportInstance" edge to the TransportInstance entity by ID.
func (truo *TransportRecipientUpdateOne) SetTransportInstanceID(id int) *TransportRecipientUpdateOne {
	truo.mutation.SetTransportInstanceID(id)
	return truo
}

// SetTransportInstance sets the "TransportInstance" edge to the TransportInstance entity.
func (truo *TransportRecipientUpdateOne) SetTransportInstance(t *TransportInstance) *TransportRecipientUpdateOne {
	return truo.SetTransportInstanceID(t.ID)
}

// AddAppRecipientIDs adds the "AppRecipient" edge to the App entity by IDs.
func (truo *TransportRecipientUpdateOne) AddAppRecipientIDs(ids ...int) *TransportRecipientUpdateOne {
	truo.mutation.AddAppRecipientIDs(ids...)
	return truo
}

// AddAppRecipient adds the "AppRecipient" edges to the App entity.
func (truo *TransportRecipientUpdateOne) AddAppRecipient(a ...*App) *TransportRecipientUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return truo.AddAppRecipientIDs(ids...)
}

// AddGroupRecipientIDs adds the "GroupRecipient" edge to the Group entity by IDs.
func (truo *TransportRecipientUpdateOne) AddGroupRecipientIDs(ids ...int) *TransportRecipientUpdateOne {
	truo.mutation.AddGroupRecipientIDs(ids...)
	return truo
}

// AddGroupRecipient adds the "GroupRecipient" edges to the Group entity.
func (truo *TransportRecipientUpdateOne) AddGroupRecipient(g ...*Group) *TransportRecipientUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return truo.AddGroupRecipientIDs(ids...)
}

// AddUserRecipientIDs adds the "UserRecipient" edge to the User entity by IDs.
func (truo *TransportRecipientUpdateOne) AddUserRecipientIDs(ids ...int) *TransportRecipientUpdateOne {
	truo.mutation.AddUserRecipientIDs(ids...)
	return truo
}

// AddUserRecipient adds the "UserRecipient" edges to the User entity.
func (truo *TransportRecipientUpdateOne) AddUserRecipient(u ...*User) *TransportRecipientUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return truo.AddUserRecipientIDs(ids...)
}

// Mutation returns the TransportRecipientMutation object of the builder.
func (truo *TransportRecipientUpdateOne) Mutation() *TransportRecipientMutation {
	return truo.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (truo *TransportRecipientUpdateOne) ClearTenant() *TransportRecipientUpdateOne {
	truo.mutation.ClearTenant()
	return truo
}

// ClearTransportInstance clears the "TransportInstance" edge to the TransportInstance entity.
func (truo *TransportRecipientUpdateOne) ClearTransportInstance() *TransportRecipientUpdateOne {
	truo.mutation.ClearTransportInstance()
	return truo
}

// ClearAppRecipient clears all "AppRecipient" edges to the App entity.
func (truo *TransportRecipientUpdateOne) ClearAppRecipient() *TransportRecipientUpdateOne {
	truo.mutation.ClearAppRecipient()
	return truo
}

// RemoveAppRecipientIDs removes the "AppRecipient" edge to App entities by IDs.
func (truo *TransportRecipientUpdateOne) RemoveAppRecipientIDs(ids ...int) *TransportRecipientUpdateOne {
	truo.mutation.RemoveAppRecipientIDs(ids...)
	return truo
}

// RemoveAppRecipient removes "AppRecipient" edges to App entities.
func (truo *TransportRecipientUpdateOne) RemoveAppRecipient(a ...*App) *TransportRecipientUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return truo.RemoveAppRecipientIDs(ids...)
}

// ClearGroupRecipient clears all "GroupRecipient" edges to the Group entity.
func (truo *TransportRecipientUpdateOne) ClearGroupRecipient() *TransportRecipientUpdateOne {
	truo.mutation.ClearGroupRecipient()
	return truo
}

// RemoveGroupRecipientIDs removes the "GroupRecipient" edge to Group entities by IDs.
func (truo *TransportRecipientUpdateOne) RemoveGroupRecipientIDs(ids ...int) *TransportRecipientUpdateOne {
	truo.mutation.RemoveGroupRecipientIDs(ids...)
	return truo
}

// RemoveGroupRecipient removes "GroupRecipient" edges to Group entities.
func (truo *TransportRecipientUpdateOne) RemoveGroupRecipient(g ...*Group) *TransportRecipientUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return truo.RemoveGroupRecipientIDs(ids...)
}

// ClearUserRecipient clears all "UserRecipient" edges to the User entity.
func (truo *TransportRecipientUpdateOne) ClearUserRecipient() *TransportRecipientUpdateOne {
	truo.mutation.ClearUserRecipient()
	return truo
}

// RemoveUserRecipientIDs removes the "UserRecipient" edge to User entities by IDs.
func (truo *TransportRecipientUpdateOne) RemoveUserRecipientIDs(ids ...int) *TransportRecipientUpdateOne {
	truo.mutation.RemoveUserRecipientIDs(ids...)
	return truo
}

// RemoveUserRecipient removes "UserRecipient" edges to User entities.
func (truo *TransportRecipientUpdateOne) RemoveUserRecipient(u ...*User) *TransportRecipientUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return truo.RemoveUserRecipientIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (truo *TransportRecipientUpdateOne) Select(field string, fields ...string) *TransportRecipientUpdateOne {
	truo.fields = append([]string{field}, fields...)
	return truo
}

// Save executes the query and returns the updated TransportRecipient entity.
func (truo *TransportRecipientUpdateOne) Save(ctx context.Context) (*TransportRecipient, error) {
	var (
		err  error
		node *TransportRecipient
	)
	if len(truo.hooks) == 0 {
		if err = truo.check(); err != nil {
			return nil, err
		}
		node, err = truo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransportRecipientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = truo.check(); err != nil {
				return nil, err
			}
			truo.mutation = mutation
			node, err = truo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(truo.hooks) - 1; i >= 0; i-- {
			if truo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = truo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, truo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TransportRecipient)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TransportRecipientMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TransportRecipientUpdateOne) SaveX(ctx context.Context) *TransportRecipient {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TransportRecipientUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TransportRecipientUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (truo *TransportRecipientUpdateOne) check() error {
	if v, ok := truo.mutation.Name(); ok {
		if err := transportrecipient.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "TransportRecipient.Name": %w`, err)}
		}
	}
	if v, ok := truo.mutation.Description(); ok {
		if err := transportrecipient.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "Description", err: fmt.Errorf(`ent: validator failed for field "TransportRecipient.Description": %w`, err)}
		}
	}
	if _, ok := truo.mutation.TenantID(); truo.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TransportRecipient.tenant"`)
	}
	if _, ok := truo.mutation.TransportInstanceID(); truo.mutation.TransportInstanceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TransportRecipient.TransportInstance"`)
	}
	return nil
}

func (truo *TransportRecipientUpdateOne) sqlSave(ctx context.Context) (_node *TransportRecipient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transportrecipient.Table,
			Columns: transportrecipient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transportrecipient.FieldID,
			},
		},
	}
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TransportRecipient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := truo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transportrecipient.FieldID)
		for _, f := range fields {
			if !transportrecipient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transportrecipient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := truo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := truo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transportrecipient.FieldName,
		})
	}
	if value, ok := truo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transportrecipient.FieldDescription,
		})
	}
	if truo.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transportrecipient.TenantTable,
			Columns: []string{transportrecipient.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transportrecipient.TenantTable,
			Columns: []string{transportrecipient.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.TransportInstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transportrecipient.TransportInstanceTable,
			Columns: []string{transportrecipient.TransportInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportinstance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TransportInstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transportrecipient.TransportInstanceTable,
			Columns: []string{transportrecipient.TransportInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.AppRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.AppRecipientTable,
			Columns: transportrecipient.AppRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.RemovedAppRecipientIDs(); len(nodes) > 0 && !truo.mutation.AppRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.AppRecipientTable,
			Columns: transportrecipient.AppRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.AppRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.AppRecipientTable,
			Columns: transportrecipient.AppRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.GroupRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.GroupRecipientTable,
			Columns: transportrecipient.GroupRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.RemovedGroupRecipientIDs(); len(nodes) > 0 && !truo.mutation.GroupRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.GroupRecipientTable,
			Columns: transportrecipient.GroupRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.GroupRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.GroupRecipientTable,
			Columns: transportrecipient.GroupRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.UserRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.UserRecipientTable,
			Columns: transportrecipient.UserRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.RemovedUserRecipientIDs(); len(nodes) > 0 && !truo.mutation.UserRecipientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.UserRecipientTable,
			Columns: transportrecipient.UserRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.UserRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   transportrecipient.UserRecipientTable,
			Columns: transportrecipient.UserRecipientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TransportRecipient{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transportrecipient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}