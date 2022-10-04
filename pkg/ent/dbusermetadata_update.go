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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbusermetadata"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

// DbUserMetaDataUpdate is the builder for updating DbUserMetaData entities.
type DbUserMetaDataUpdate struct {
	config
	hooks    []Hook
	mutation *DbUserMetaDataMutation
}

// Where appends a list predicates to the DbUserMetaDataUpdate builder.
func (dumdu *DbUserMetaDataUpdate) Where(ps ...predicate.DbUserMetaData) *DbUserMetaDataUpdate {
	dumdu.mutation.Where(ps...)
	return dumdu
}

// SetTenantID sets the "tenant_id" field.
func (dumdu *DbUserMetaDataUpdate) SetTenantID(i int) *DbUserMetaDataUpdate {
	dumdu.mutation.SetTenantID(i)
	return dumdu
}

// SetAppData sets the "AppData" field.
func (dumdu *DbUserMetaDataUpdate) SetAppData(id interfaces.AppData) *DbUserMetaDataUpdate {
	dumdu.mutation.SetAppData(id)
	return dumdu
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (dumdu *DbUserMetaDataUpdate) SetNillableAppData(id *interfaces.AppData) *DbUserMetaDataUpdate {
	if id != nil {
		dumdu.SetAppData(*id)
	}
	return dumdu
}

// ClearAppData clears the value of the "AppData" field.
func (dumdu *DbUserMetaDataUpdate) ClearAppData() *DbUserMetaDataUpdate {
	dumdu.mutation.ClearAppData()
	return dumdu
}

// SetName sets the "Name" field.
func (dumdu *DbUserMetaDataUpdate) SetName(s string) *DbUserMetaDataUpdate {
	dumdu.mutation.SetName(s)
	return dumdu
}

// SetValue sets the "Value" field.
func (dumdu *DbUserMetaDataUpdate) SetValue(s string) *DbUserMetaDataUpdate {
	dumdu.mutation.SetValue(s)
	return dumdu
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (dumdu *DbUserMetaDataUpdate) SetTenant(t *Tenant) *DbUserMetaDataUpdate {
	return dumdu.SetTenantID(t.ID)
}

// SetUserID sets the "user" edge to the DbUser entity by ID.
func (dumdu *DbUserMetaDataUpdate) SetUserID(id int) *DbUserMetaDataUpdate {
	dumdu.mutation.SetUserID(id)
	return dumdu
}

// SetUser sets the "user" edge to the DbUser entity.
func (dumdu *DbUserMetaDataUpdate) SetUser(d *DbUser) *DbUserMetaDataUpdate {
	return dumdu.SetUserID(d.ID)
}

// Mutation returns the DbUserMetaDataMutation object of the builder.
func (dumdu *DbUserMetaDataUpdate) Mutation() *DbUserMetaDataMutation {
	return dumdu.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (dumdu *DbUserMetaDataUpdate) ClearTenant() *DbUserMetaDataUpdate {
	dumdu.mutation.ClearTenant()
	return dumdu
}

// ClearUser clears the "user" edge to the DbUser entity.
func (dumdu *DbUserMetaDataUpdate) ClearUser() *DbUserMetaDataUpdate {
	dumdu.mutation.ClearUser()
	return dumdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dumdu *DbUserMetaDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dumdu.hooks) == 0 {
		if err = dumdu.check(); err != nil {
			return 0, err
		}
		affected, err = dumdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbUserMetaDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dumdu.check(); err != nil {
				return 0, err
			}
			dumdu.mutation = mutation
			affected, err = dumdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dumdu.hooks) - 1; i >= 0; i-- {
			if dumdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dumdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dumdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dumdu *DbUserMetaDataUpdate) SaveX(ctx context.Context) int {
	affected, err := dumdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dumdu *DbUserMetaDataUpdate) Exec(ctx context.Context) error {
	_, err := dumdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dumdu *DbUserMetaDataUpdate) ExecX(ctx context.Context) {
	if err := dumdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dumdu *DbUserMetaDataUpdate) check() error {
	if v, ok := dumdu.mutation.Name(); ok {
		if err := dbusermetadata.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbUserMetaData.Name": %w`, err)}
		}
	}
	if v, ok := dumdu.mutation.Value(); ok {
		if err := dbusermetadata.ValueValidator(v); err != nil {
			return &ValidationError{Name: "Value", err: fmt.Errorf(`ent: validator failed for field "DbUserMetaData.Value": %w`, err)}
		}
	}
	if _, ok := dumdu.mutation.TenantID(); dumdu.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DbUserMetaData.tenant"`)
	}
	if _, ok := dumdu.mutation.UserID(); dumdu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DbUserMetaData.user"`)
	}
	return nil
}

func (dumdu *DbUserMetaDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbusermetadata.Table,
			Columns: dbusermetadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbusermetadata.FieldID,
			},
		},
	}
	if ps := dumdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dumdu.mutation.AppData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbusermetadata.FieldAppData,
		})
	}
	if dumdu.mutation.AppDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dbusermetadata.FieldAppData,
		})
	}
	if value, ok := dumdu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbusermetadata.FieldName,
		})
	}
	if value, ok := dumdu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbusermetadata.FieldValue,
		})
	}
	if dumdu.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbusermetadata.TenantTable,
			Columns: []string{dbusermetadata.TenantColumn},
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
	if nodes := dumdu.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbusermetadata.TenantTable,
			Columns: []string{dbusermetadata.TenantColumn},
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
	if dumdu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbusermetadata.UserTable,
			Columns: []string{dbusermetadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dumdu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbusermetadata.UserTable,
			Columns: []string{dbusermetadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dumdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbusermetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DbUserMetaDataUpdateOne is the builder for updating a single DbUserMetaData entity.
type DbUserMetaDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DbUserMetaDataMutation
}

// SetTenantID sets the "tenant_id" field.
func (dumduo *DbUserMetaDataUpdateOne) SetTenantID(i int) *DbUserMetaDataUpdateOne {
	dumduo.mutation.SetTenantID(i)
	return dumduo
}

// SetAppData sets the "AppData" field.
func (dumduo *DbUserMetaDataUpdateOne) SetAppData(id interfaces.AppData) *DbUserMetaDataUpdateOne {
	dumduo.mutation.SetAppData(id)
	return dumduo
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (dumduo *DbUserMetaDataUpdateOne) SetNillableAppData(id *interfaces.AppData) *DbUserMetaDataUpdateOne {
	if id != nil {
		dumduo.SetAppData(*id)
	}
	return dumduo
}

// ClearAppData clears the value of the "AppData" field.
func (dumduo *DbUserMetaDataUpdateOne) ClearAppData() *DbUserMetaDataUpdateOne {
	dumduo.mutation.ClearAppData()
	return dumduo
}

// SetName sets the "Name" field.
func (dumduo *DbUserMetaDataUpdateOne) SetName(s string) *DbUserMetaDataUpdateOne {
	dumduo.mutation.SetName(s)
	return dumduo
}

// SetValue sets the "Value" field.
func (dumduo *DbUserMetaDataUpdateOne) SetValue(s string) *DbUserMetaDataUpdateOne {
	dumduo.mutation.SetValue(s)
	return dumduo
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (dumduo *DbUserMetaDataUpdateOne) SetTenant(t *Tenant) *DbUserMetaDataUpdateOne {
	return dumduo.SetTenantID(t.ID)
}

// SetUserID sets the "user" edge to the DbUser entity by ID.
func (dumduo *DbUserMetaDataUpdateOne) SetUserID(id int) *DbUserMetaDataUpdateOne {
	dumduo.mutation.SetUserID(id)
	return dumduo
}

// SetUser sets the "user" edge to the DbUser entity.
func (dumduo *DbUserMetaDataUpdateOne) SetUser(d *DbUser) *DbUserMetaDataUpdateOne {
	return dumduo.SetUserID(d.ID)
}

// Mutation returns the DbUserMetaDataMutation object of the builder.
func (dumduo *DbUserMetaDataUpdateOne) Mutation() *DbUserMetaDataMutation {
	return dumduo.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (dumduo *DbUserMetaDataUpdateOne) ClearTenant() *DbUserMetaDataUpdateOne {
	dumduo.mutation.ClearTenant()
	return dumduo
}

// ClearUser clears the "user" edge to the DbUser entity.
func (dumduo *DbUserMetaDataUpdateOne) ClearUser() *DbUserMetaDataUpdateOne {
	dumduo.mutation.ClearUser()
	return dumduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dumduo *DbUserMetaDataUpdateOne) Select(field string, fields ...string) *DbUserMetaDataUpdateOne {
	dumduo.fields = append([]string{field}, fields...)
	return dumduo
}

// Save executes the query and returns the updated DbUserMetaData entity.
func (dumduo *DbUserMetaDataUpdateOne) Save(ctx context.Context) (*DbUserMetaData, error) {
	var (
		err  error
		node *DbUserMetaData
	)
	if len(dumduo.hooks) == 0 {
		if err = dumduo.check(); err != nil {
			return nil, err
		}
		node, err = dumduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbUserMetaDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dumduo.check(); err != nil {
				return nil, err
			}
			dumduo.mutation = mutation
			node, err = dumduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dumduo.hooks) - 1; i >= 0; i-- {
			if dumduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dumduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dumduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DbUserMetaData)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DbUserMetaDataMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dumduo *DbUserMetaDataUpdateOne) SaveX(ctx context.Context) *DbUserMetaData {
	node, err := dumduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dumduo *DbUserMetaDataUpdateOne) Exec(ctx context.Context) error {
	_, err := dumduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dumduo *DbUserMetaDataUpdateOne) ExecX(ctx context.Context) {
	if err := dumduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dumduo *DbUserMetaDataUpdateOne) check() error {
	if v, ok := dumduo.mutation.Name(); ok {
		if err := dbusermetadata.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbUserMetaData.Name": %w`, err)}
		}
	}
	if v, ok := dumduo.mutation.Value(); ok {
		if err := dbusermetadata.ValueValidator(v); err != nil {
			return &ValidationError{Name: "Value", err: fmt.Errorf(`ent: validator failed for field "DbUserMetaData.Value": %w`, err)}
		}
	}
	if _, ok := dumduo.mutation.TenantID(); dumduo.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DbUserMetaData.tenant"`)
	}
	if _, ok := dumduo.mutation.UserID(); dumduo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DbUserMetaData.user"`)
	}
	return nil
}

func (dumduo *DbUserMetaDataUpdateOne) sqlSave(ctx context.Context) (_node *DbUserMetaData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbusermetadata.Table,
			Columns: dbusermetadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbusermetadata.FieldID,
			},
		},
	}
	id, ok := dumduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DbUserMetaData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dumduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbusermetadata.FieldID)
		for _, f := range fields {
			if !dbusermetadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dbusermetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dumduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dumduo.mutation.AppData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbusermetadata.FieldAppData,
		})
	}
	if dumduo.mutation.AppDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dbusermetadata.FieldAppData,
		})
	}
	if value, ok := dumduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbusermetadata.FieldName,
		})
	}
	if value, ok := dumduo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbusermetadata.FieldValue,
		})
	}
	if dumduo.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbusermetadata.TenantTable,
			Columns: []string{dbusermetadata.TenantColumn},
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
	if nodes := dumduo.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbusermetadata.TenantTable,
			Columns: []string{dbusermetadata.TenantColumn},
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
	if dumduo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbusermetadata.UserTable,
			Columns: []string{dbusermetadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dumduo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbusermetadata.UserTable,
			Columns: []string{dbusermetadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DbUserMetaData{config: dumduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dumduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbusermetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
