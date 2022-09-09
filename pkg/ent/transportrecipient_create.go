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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/app"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/group"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportinstance"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportrecipient"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/user"
)

// TransportRecipientCreate is the builder for creating a TransportRecipient entity.
type TransportRecipientCreate struct {
	config
	mutation *TransportRecipientMutation
	hooks    []Hook
}

// SetTenantID sets the "tenant_id" field.
func (trc *TransportRecipientCreate) SetTenantID(i int) *TransportRecipientCreate {
	trc.mutation.SetTenantID(i)
	return trc
}

// SetName sets the "Name" field.
func (trc *TransportRecipientCreate) SetName(s string) *TransportRecipientCreate {
	trc.mutation.SetName(s)
	return trc
}

// SetDescription sets the "Description" field.
func (trc *TransportRecipientCreate) SetDescription(s string) *TransportRecipientCreate {
	trc.mutation.SetDescription(s)
	return trc
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (trc *TransportRecipientCreate) SetTenant(t *Tenant) *TransportRecipientCreate {
	return trc.SetTenantID(t.ID)
}

// SetTransportInstanceID sets the "TransportInstance" edge to the TransportInstance entity by ID.
func (trc *TransportRecipientCreate) SetTransportInstanceID(id int) *TransportRecipientCreate {
	trc.mutation.SetTransportInstanceID(id)
	return trc
}

// SetTransportInstance sets the "TransportInstance" edge to the TransportInstance entity.
func (trc *TransportRecipientCreate) SetTransportInstance(t *TransportInstance) *TransportRecipientCreate {
	return trc.SetTransportInstanceID(t.ID)
}

// AddAppRecipientIDs adds the "AppRecipient" edge to the App entity by IDs.
func (trc *TransportRecipientCreate) AddAppRecipientIDs(ids ...int) *TransportRecipientCreate {
	trc.mutation.AddAppRecipientIDs(ids...)
	return trc
}

// AddAppRecipient adds the "AppRecipient" edges to the App entity.
func (trc *TransportRecipientCreate) AddAppRecipient(a ...*App) *TransportRecipientCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return trc.AddAppRecipientIDs(ids...)
}

// AddGroupRecipientIDs adds the "GroupRecipient" edge to the Group entity by IDs.
func (trc *TransportRecipientCreate) AddGroupRecipientIDs(ids ...int) *TransportRecipientCreate {
	trc.mutation.AddGroupRecipientIDs(ids...)
	return trc
}

// AddGroupRecipient adds the "GroupRecipient" edges to the Group entity.
func (trc *TransportRecipientCreate) AddGroupRecipient(g ...*Group) *TransportRecipientCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return trc.AddGroupRecipientIDs(ids...)
}

// AddUserRecipientIDs adds the "UserRecipient" edge to the User entity by IDs.
func (trc *TransportRecipientCreate) AddUserRecipientIDs(ids ...int) *TransportRecipientCreate {
	trc.mutation.AddUserRecipientIDs(ids...)
	return trc
}

// AddUserRecipient adds the "UserRecipient" edges to the User entity.
func (trc *TransportRecipientCreate) AddUserRecipient(u ...*User) *TransportRecipientCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return trc.AddUserRecipientIDs(ids...)
}

// Mutation returns the TransportRecipientMutation object of the builder.
func (trc *TransportRecipientCreate) Mutation() *TransportRecipientMutation {
	return trc.mutation
}

// Save creates the TransportRecipient in the database.
func (trc *TransportRecipientCreate) Save(ctx context.Context) (*TransportRecipient, error) {
	var (
		err  error
		node *TransportRecipient
	)
	if len(trc.hooks) == 0 {
		if err = trc.check(); err != nil {
			return nil, err
		}
		node, err = trc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransportRecipientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = trc.check(); err != nil {
				return nil, err
			}
			trc.mutation = mutation
			if node, err = trc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(trc.hooks) - 1; i >= 0; i-- {
			if trc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = trc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, trc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (trc *TransportRecipientCreate) SaveX(ctx context.Context) *TransportRecipient {
	v, err := trc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trc *TransportRecipientCreate) Exec(ctx context.Context) error {
	_, err := trc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trc *TransportRecipientCreate) ExecX(ctx context.Context) {
	if err := trc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trc *TransportRecipientCreate) check() error {
	if _, ok := trc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "TransportRecipient.tenant_id"`)}
	}
	if _, ok := trc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "TransportRecipient.Name"`)}
	}
	if v, ok := trc.mutation.Name(); ok {
		if err := transportrecipient.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "TransportRecipient.Name": %w`, err)}
		}
	}
	if _, ok := trc.mutation.Description(); !ok {
		return &ValidationError{Name: "Description", err: errors.New(`ent: missing required field "TransportRecipient.Description"`)}
	}
	if v, ok := trc.mutation.Description(); ok {
		if err := transportrecipient.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "Description", err: fmt.Errorf(`ent: validator failed for field "TransportRecipient.Description": %w`, err)}
		}
	}
	if _, ok := trc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`ent: missing required edge "TransportRecipient.tenant"`)}
	}
	if _, ok := trc.mutation.TransportInstanceID(); !ok {
		return &ValidationError{Name: "TransportInstance", err: errors.New(`ent: missing required edge "TransportRecipient.TransportInstance"`)}
	}
	return nil
}

func (trc *TransportRecipientCreate) sqlSave(ctx context.Context) (*TransportRecipient, error) {
	_node, _spec := trc.createSpec()
	if err := sqlgraph.CreateNode(ctx, trc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (trc *TransportRecipientCreate) createSpec() (*TransportRecipient, *sqlgraph.CreateSpec) {
	var (
		_node = &TransportRecipient{config: trc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transportrecipient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transportrecipient.FieldID,
			},
		}
	)
	if value, ok := trc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transportrecipient.FieldName,
		})
		_node.Name = value
	}
	if value, ok := trc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transportrecipient.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := trc.mutation.TenantIDs(); len(nodes) > 0 {
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
		_node.TenantID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.TransportInstanceIDs(); len(nodes) > 0 {
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
		_node.transport_instance_transport_recipients = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.AppRecipientIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.GroupRecipientIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.UserRecipientIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

func (trc *TransportRecipientCreate) SetTransportRecipientFromStruct(input *TransportRecipient) *TransportRecipientCreate {

	trc.SetTenantID(input.TenantID)

	trc.SetName(input.Name)

	trc.SetDescription(input.Description)

	return trc
}

// TransportRecipientCreateBulk is the builder for creating many TransportRecipient entities in bulk.
type TransportRecipientCreateBulk struct {
	config
	builders []*TransportRecipientCreate
}

// Save creates the TransportRecipient entities in the database.
func (trcb *TransportRecipientCreateBulk) Save(ctx context.Context) ([]*TransportRecipient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(trcb.builders))
	nodes := make([]*TransportRecipient, len(trcb.builders))
	mutators := make([]Mutator, len(trcb.builders))
	for i := range trcb.builders {
		func(i int, root context.Context) {
			builder := trcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransportRecipientMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, trcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, trcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, trcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (trcb *TransportRecipientCreateBulk) SaveX(ctx context.Context) []*TransportRecipient {
	v, err := trcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trcb *TransportRecipientCreateBulk) Exec(ctx context.Context) error {
	_, err := trcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trcb *TransportRecipientCreateBulk) ExecX(ctx context.Context) {
	if err := trcb.Exec(ctx); err != nil {
		panic(err)
	}
}