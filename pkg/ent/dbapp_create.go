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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbapp"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbgroup"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/google/uuid"
)

// DbAppCreate is the builder for creating a DbApp entity.
type DbAppCreate struct {
	config
	mutation *DbAppMutation
	hooks    []Hook
}

// SetTenantID sets the "tenant_id" field.
func (dac *DbAppCreate) SetTenantID(i int) *DbAppCreate {
	dac.mutation.SetTenantID(i)
	return dac
}

// SetName sets the "Name" field.
func (dac *DbAppCreate) SetName(s string) *DbAppCreate {
	dac.mutation.SetName(s)
	return dac
}

// SetStatus sets the "Status" field.
func (dac *DbAppCreate) SetStatus(is interfaces.AppStatus) *DbAppCreate {
	dac.mutation.SetStatus(is)
	return dac
}

// SetDescription sets the "Description" field.
func (dac *DbAppCreate) SetDescription(s string) *DbAppCreate {
	dac.mutation.SetDescription(s)
	return dac
}

// SetIcon sets the "icon" field.
func (dac *DbAppCreate) SetIcon(s string) *DbAppCreate {
	dac.mutation.SetIcon(s)
	return dac
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (dac *DbAppCreate) SetNillableIcon(s *string) *DbAppCreate {
	if s != nil {
		dac.SetIcon(*s)
	}
	return dac
}

// SetURL sets the "url" field.
func (dac *DbAppCreate) SetURL(s string) *DbAppCreate {
	dac.mutation.SetURL(s)
	return dac
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (dac *DbAppCreate) SetNillableURL(s *string) *DbAppCreate {
	if s != nil {
		dac.SetURL(*s)
	}
	return dac
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (dac *DbAppCreate) SetTenant(t *Tenant) *DbAppCreate {
	return dac.SetTenantID(t.ID)
}

// AddMessageIDs adds the "messages" edge to the DbMessage entity by IDs.
func (dac *DbAppCreate) AddMessageIDs(ids ...uuid.UUID) *DbAppCreate {
	dac.mutation.AddMessageIDs(ids...)
	return dac
}

// AddMessages adds the "messages" edges to the DbMessage entity.
func (dac *DbAppCreate) AddMessages(d ...*DbMessage) *DbAppCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dac.AddMessageIDs(ids...)
}

// AddFilterIDs adds the "filters" edge to the DbFilter entity by IDs.
func (dac *DbAppCreate) AddFilterIDs(ids ...int) *DbAppCreate {
	dac.mutation.AddFilterIDs(ids...)
	return dac
}

// AddFilters adds the "filters" edges to the DbFilter entity.
func (dac *DbAppCreate) AddFilters(d ...*DbFilter) *DbAppCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dac.AddFilterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the DbGroup entity by IDs.
func (dac *DbAppCreate) AddGroupIDs(ids ...int) *DbAppCreate {
	dac.mutation.AddGroupIDs(ids...)
	return dac
}

// AddGroups adds the "groups" edges to the DbGroup entity.
func (dac *DbAppCreate) AddGroups(d ...*DbGroup) *DbAppCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dac.AddGroupIDs(ids...)
}

// Mutation returns the DbAppMutation object of the builder.
func (dac *DbAppCreate) Mutation() *DbAppMutation {
	return dac.mutation
}

// Save creates the DbApp in the database.
func (dac *DbAppCreate) Save(ctx context.Context) (*DbApp, error) {
	var (
		err  error
		node *DbApp
	)
	if len(dac.hooks) == 0 {
		if err = dac.check(); err != nil {
			return nil, err
		}
		node, err = dac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbAppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dac.check(); err != nil {
				return nil, err
			}
			dac.mutation = mutation
			if node, err = dac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dac.hooks) - 1; i >= 0; i-- {
			if dac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DbApp)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DbAppMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dac *DbAppCreate) SaveX(ctx context.Context) *DbApp {
	v, err := dac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dac *DbAppCreate) Exec(ctx context.Context) error {
	_, err := dac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dac *DbAppCreate) ExecX(ctx context.Context) {
	if err := dac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dac *DbAppCreate) check() error {
	if _, ok := dac.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "DbApp.tenant_id"`)}
	}
	if _, ok := dac.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "DbApp.Name"`)}
	}
	if v, ok := dac.mutation.Name(); ok {
		if err := dbapp.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbApp.Name": %w`, err)}
		}
	}
	if _, ok := dac.mutation.Status(); !ok {
		return &ValidationError{Name: "Status", err: errors.New(`ent: missing required field "DbApp.Status"`)}
	}
	if v, ok := dac.mutation.Status(); ok {
		if err := dbapp.StatusValidator(v); err != nil {
			return &ValidationError{Name: "Status", err: fmt.Errorf(`ent: validator failed for field "DbApp.Status": %w`, err)}
		}
	}
	if _, ok := dac.mutation.Description(); !ok {
		return &ValidationError{Name: "Description", err: errors.New(`ent: missing required field "DbApp.Description"`)}
	}
	if v, ok := dac.mutation.Description(); ok {
		if err := dbapp.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "Description", err: fmt.Errorf(`ent: validator failed for field "DbApp.Description": %w`, err)}
		}
	}
	if v, ok := dac.mutation.Icon(); ok {
		if err := dbapp.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "DbApp.icon": %w`, err)}
		}
	}
	if v, ok := dac.mutation.URL(); ok {
		if err := dbapp.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "DbApp.url": %w`, err)}
		}
	}
	if _, ok := dac.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`ent: missing required edge "DbApp.tenant"`)}
	}
	return nil
}

func (dac *DbAppCreate) sqlSave(ctx context.Context) (*DbApp, error) {
	_node, _spec := dac.createSpec()
	if err := sqlgraph.CreateNode(ctx, dac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dac *DbAppCreate) createSpec() (*DbApp, *sqlgraph.CreateSpec) {
	var (
		_node = &DbApp{config: dac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dbapp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbapp.FieldID,
			},
		}
	)
	if value, ok := dac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbapp.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbapp.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := dac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbapp.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := dac.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbapp.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := dac.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbapp.FieldURL,
		})
		_node.URL = value
	}
	if nodes := dac.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbapp.TenantTable,
			Columns: []string{dbapp.TenantColumn},
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
	if nodes := dac.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbapp.MessagesTable,
			Columns: []string{dbapp.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dbmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dac.mutation.FiltersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbapp.FiltersTable,
			Columns: dbapp.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbfilter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dac.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbapp.GroupsTable,
			Columns: dbapp.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbgroup.FieldID,
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

func (dac *DbAppCreate) SetDbAppFromStruct(input *DbApp) *DbAppCreate {

	dac.SetTenantID(input.TenantID)

	dac.SetName(input.Name)

	dac.SetStatus(input.Status)

	dac.SetDescription(input.Description)

	dac.SetIcon(input.Icon)

	dac.SetURL(input.URL)

	return dac
}

// DbAppCreateBulk is the builder for creating many DbApp entities in bulk.
type DbAppCreateBulk struct {
	config
	builders []*DbAppCreate
}

// Save creates the DbApp entities in the database.
func (dacb *DbAppCreateBulk) Save(ctx context.Context) ([]*DbApp, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dacb.builders))
	nodes := make([]*DbApp, len(dacb.builders))
	mutators := make([]Mutator, len(dacb.builders))
	for i := range dacb.builders {
		func(i int, root context.Context) {
			builder := dacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DbAppMutation)
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
					_, err = mutators[i+1].Mutate(root, dacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dacb *DbAppCreateBulk) SaveX(ctx context.Context) []*DbApp {
	v, err := dacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dacb *DbAppCreateBulk) Exec(ctx context.Context) error {
	_, err := dacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dacb *DbAppCreateBulk) ExecX(ctx context.Context) {
	if err := dacb.Exec(ctx); err != nil {
		panic(err)
	}
}
