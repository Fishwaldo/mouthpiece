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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbgroup"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbusermetadata"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

// DbUserUpdate is the builder for updating DbUser entities.
type DbUserUpdate struct {
	config
	hooks    []Hook
	mutation *DbUserMutation
}

// Where appends a list predicates to the DbUserUpdate builder.
func (duu *DbUserUpdate) Where(ps ...predicate.DbUser) *DbUserUpdate {
	duu.mutation.Where(ps...)
	return duu
}

// SetTenantID sets the "tenant_id" field.
func (duu *DbUserUpdate) SetTenantID(i int) *DbUserUpdate {
	duu.mutation.SetTenantID(i)
	return duu
}

// SetAppData sets the "AppData" field.
func (duu *DbUserUpdate) SetAppData(id interfaces.AppData) *DbUserUpdate {
	duu.mutation.SetAppData(id)
	return duu
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (duu *DbUserUpdate) SetNillableAppData(id *interfaces.AppData) *DbUserUpdate {
	if id != nil {
		duu.SetAppData(*id)
	}
	return duu
}

// ClearAppData clears the value of the "AppData" field.
func (duu *DbUserUpdate) ClearAppData() *DbUserUpdate {
	duu.mutation.ClearAppData()
	return duu
}

// SetEmail sets the "Email" field.
func (duu *DbUserUpdate) SetEmail(s string) *DbUserUpdate {
	duu.mutation.SetEmail(s)
	return duu
}

// SetName sets the "Name" field.
func (duu *DbUserUpdate) SetName(s string) *DbUserUpdate {
	duu.mutation.SetName(s)
	return duu
}

// SetDescription sets the "Description" field.
func (duu *DbUserUpdate) SetDescription(s string) *DbUserUpdate {
	duu.mutation.SetDescription(s)
	return duu
}

// SetNillableDescription sets the "Description" field if the given value is not nil.
func (duu *DbUserUpdate) SetNillableDescription(s *string) *DbUserUpdate {
	if s != nil {
		duu.SetDescription(*s)
	}
	return duu
}

// ClearDescription clears the value of the "Description" field.
func (duu *DbUserUpdate) ClearDescription() *DbUserUpdate {
	duu.mutation.ClearDescription()
	return duu
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (duu *DbUserUpdate) SetTenant(t *Tenant) *DbUserUpdate {
	return duu.SetTenantID(t.ID)
}

// AddMetadatumIDs adds the "metadata" edge to the DbUserMetaData entity by IDs.
func (duu *DbUserUpdate) AddMetadatumIDs(ids ...int) *DbUserUpdate {
	duu.mutation.AddMetadatumIDs(ids...)
	return duu
}

// AddMetadata adds the "metadata" edges to the DbUserMetaData entity.
func (duu *DbUserUpdate) AddMetadata(d ...*DbUserMetaData) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.AddMetadatumIDs(ids...)
}

// AddFilterIDs adds the "filters" edge to the DbFilter entity by IDs.
func (duu *DbUserUpdate) AddFilterIDs(ids ...int) *DbUserUpdate {
	duu.mutation.AddFilterIDs(ids...)
	return duu
}

// AddFilters adds the "filters" edges to the DbFilter entity.
func (duu *DbUserUpdate) AddFilters(d ...*DbFilter) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.AddFilterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the DbGroup entity by IDs.
func (duu *DbUserUpdate) AddGroupIDs(ids ...int) *DbUserUpdate {
	duu.mutation.AddGroupIDs(ids...)
	return duu
}

// AddGroups adds the "groups" edges to the DbGroup entity.
func (duu *DbUserUpdate) AddGroups(d ...*DbGroup) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.AddGroupIDs(ids...)
}

// AddTransportRecipientIDs adds the "TransportRecipients" edge to the DbTransportRecipients entity by IDs.
func (duu *DbUserUpdate) AddTransportRecipientIDs(ids ...int) *DbUserUpdate {
	duu.mutation.AddTransportRecipientIDs(ids...)
	return duu
}

// AddTransportRecipients adds the "TransportRecipients" edges to the DbTransportRecipients entity.
func (duu *DbUserUpdate) AddTransportRecipients(d ...*DbTransportRecipients) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.AddTransportRecipientIDs(ids...)
}

// Mutation returns the DbUserMutation object of the builder.
func (duu *DbUserUpdate) Mutation() *DbUserMutation {
	return duu.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (duu *DbUserUpdate) ClearTenant() *DbUserUpdate {
	duu.mutation.ClearTenant()
	return duu
}

// ClearMetadata clears all "metadata" edges to the DbUserMetaData entity.
func (duu *DbUserUpdate) ClearMetadata() *DbUserUpdate {
	duu.mutation.ClearMetadata()
	return duu
}

// RemoveMetadatumIDs removes the "metadata" edge to DbUserMetaData entities by IDs.
func (duu *DbUserUpdate) RemoveMetadatumIDs(ids ...int) *DbUserUpdate {
	duu.mutation.RemoveMetadatumIDs(ids...)
	return duu
}

// RemoveMetadata removes "metadata" edges to DbUserMetaData entities.
func (duu *DbUserUpdate) RemoveMetadata(d ...*DbUserMetaData) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.RemoveMetadatumIDs(ids...)
}

// ClearFilters clears all "filters" edges to the DbFilter entity.
func (duu *DbUserUpdate) ClearFilters() *DbUserUpdate {
	duu.mutation.ClearFilters()
	return duu
}

// RemoveFilterIDs removes the "filters" edge to DbFilter entities by IDs.
func (duu *DbUserUpdate) RemoveFilterIDs(ids ...int) *DbUserUpdate {
	duu.mutation.RemoveFilterIDs(ids...)
	return duu
}

// RemoveFilters removes "filters" edges to DbFilter entities.
func (duu *DbUserUpdate) RemoveFilters(d ...*DbFilter) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.RemoveFilterIDs(ids...)
}

// ClearGroups clears all "groups" edges to the DbGroup entity.
func (duu *DbUserUpdate) ClearGroups() *DbUserUpdate {
	duu.mutation.ClearGroups()
	return duu
}

// RemoveGroupIDs removes the "groups" edge to DbGroup entities by IDs.
func (duu *DbUserUpdate) RemoveGroupIDs(ids ...int) *DbUserUpdate {
	duu.mutation.RemoveGroupIDs(ids...)
	return duu
}

// RemoveGroups removes "groups" edges to DbGroup entities.
func (duu *DbUserUpdate) RemoveGroups(d ...*DbGroup) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.RemoveGroupIDs(ids...)
}

// ClearTransportRecipients clears all "TransportRecipients" edges to the DbTransportRecipients entity.
func (duu *DbUserUpdate) ClearTransportRecipients() *DbUserUpdate {
	duu.mutation.ClearTransportRecipients()
	return duu
}

// RemoveTransportRecipientIDs removes the "TransportRecipients" edge to DbTransportRecipients entities by IDs.
func (duu *DbUserUpdate) RemoveTransportRecipientIDs(ids ...int) *DbUserUpdate {
	duu.mutation.RemoveTransportRecipientIDs(ids...)
	return duu
}

// RemoveTransportRecipients removes "TransportRecipients" edges to DbTransportRecipients entities.
func (duu *DbUserUpdate) RemoveTransportRecipients(d ...*DbTransportRecipients) *DbUserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.RemoveTransportRecipientIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (duu *DbUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(duu.hooks) == 0 {
		if err = duu.check(); err != nil {
			return 0, err
		}
		affected, err = duu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duu.check(); err != nil {
				return 0, err
			}
			duu.mutation = mutation
			affected, err = duu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(duu.hooks) - 1; i >= 0; i-- {
			if duu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (duu *DbUserUpdate) SaveX(ctx context.Context) int {
	affected, err := duu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (duu *DbUserUpdate) Exec(ctx context.Context) error {
	_, err := duu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duu *DbUserUpdate) ExecX(ctx context.Context) {
	if err := duu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duu *DbUserUpdate) check() error {
	if v, ok := duu.mutation.Email(); ok {
		if err := dbuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`ent: validator failed for field "DbUser.Email": %w`, err)}
		}
	}
	if v, ok := duu.mutation.Name(); ok {
		if err := dbuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbUser.Name": %w`, err)}
		}
	}
	if _, ok := duu.mutation.TenantID(); duu.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DbUser.tenant"`)
	}
	return nil
}

func (duu *DbUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbuser.Table,
			Columns: dbuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbuser.FieldID,
			},
		},
	}
	if ps := duu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duu.mutation.AppData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbuser.FieldAppData,
		})
	}
	if duu.mutation.AppDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dbuser.FieldAppData,
		})
	}
	if value, ok := duu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldEmail,
		})
	}
	if value, ok := duu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldName,
		})
	}
	if value, ok := duu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldDescription,
		})
	}
	if duu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbuser.FieldDescription,
		})
	}
	if duu.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbuser.TenantTable,
			Columns: []string{dbuser.TenantColumn},
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
	if nodes := duu.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbuser.TenantTable,
			Columns: []string{dbuser.TenantColumn},
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
	if duu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.MetadataTable,
			Columns: []string{dbuser.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbusermetadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !duu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.MetadataTable,
			Columns: []string{dbuser.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbusermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.MetadataTable,
			Columns: []string{dbuser.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbusermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duu.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.FiltersTable,
			Columns: dbuser.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbfilter.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.RemovedFiltersIDs(); len(nodes) > 0 && !duu.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.FiltersTable,
			Columns: dbuser.FiltersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.FiltersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.FiltersTable,
			Columns: dbuser.FiltersPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.GroupsTable,
			Columns: dbuser.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !duu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.GroupsTable,
			Columns: dbuser.GroupsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.GroupsTable,
			Columns: dbuser.GroupsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duu.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.TransportRecipientsTable,
			Columns: []string{dbuser.TransportRecipientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbtransportrecipients.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.RemovedTransportRecipientsIDs(); len(nodes) > 0 && !duu.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.TransportRecipientsTable,
			Columns: []string{dbuser.TransportRecipientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbtransportrecipients.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.TransportRecipientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.TransportRecipientsTable,
			Columns: []string{dbuser.TransportRecipientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbtransportrecipients.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, duu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DbUserUpdateOne is the builder for updating a single DbUser entity.
type DbUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DbUserMutation
}

// SetTenantID sets the "tenant_id" field.
func (duuo *DbUserUpdateOne) SetTenantID(i int) *DbUserUpdateOne {
	duuo.mutation.SetTenantID(i)
	return duuo
}

// SetAppData sets the "AppData" field.
func (duuo *DbUserUpdateOne) SetAppData(id interfaces.AppData) *DbUserUpdateOne {
	duuo.mutation.SetAppData(id)
	return duuo
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (duuo *DbUserUpdateOne) SetNillableAppData(id *interfaces.AppData) *DbUserUpdateOne {
	if id != nil {
		duuo.SetAppData(*id)
	}
	return duuo
}

// ClearAppData clears the value of the "AppData" field.
func (duuo *DbUserUpdateOne) ClearAppData() *DbUserUpdateOne {
	duuo.mutation.ClearAppData()
	return duuo
}

// SetEmail sets the "Email" field.
func (duuo *DbUserUpdateOne) SetEmail(s string) *DbUserUpdateOne {
	duuo.mutation.SetEmail(s)
	return duuo
}

// SetName sets the "Name" field.
func (duuo *DbUserUpdateOne) SetName(s string) *DbUserUpdateOne {
	duuo.mutation.SetName(s)
	return duuo
}

// SetDescription sets the "Description" field.
func (duuo *DbUserUpdateOne) SetDescription(s string) *DbUserUpdateOne {
	duuo.mutation.SetDescription(s)
	return duuo
}

// SetNillableDescription sets the "Description" field if the given value is not nil.
func (duuo *DbUserUpdateOne) SetNillableDescription(s *string) *DbUserUpdateOne {
	if s != nil {
		duuo.SetDescription(*s)
	}
	return duuo
}

// ClearDescription clears the value of the "Description" field.
func (duuo *DbUserUpdateOne) ClearDescription() *DbUserUpdateOne {
	duuo.mutation.ClearDescription()
	return duuo
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (duuo *DbUserUpdateOne) SetTenant(t *Tenant) *DbUserUpdateOne {
	return duuo.SetTenantID(t.ID)
}

// AddMetadatumIDs adds the "metadata" edge to the DbUserMetaData entity by IDs.
func (duuo *DbUserUpdateOne) AddMetadatumIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.AddMetadatumIDs(ids...)
	return duuo
}

// AddMetadata adds the "metadata" edges to the DbUserMetaData entity.
func (duuo *DbUserUpdateOne) AddMetadata(d ...*DbUserMetaData) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.AddMetadatumIDs(ids...)
}

// AddFilterIDs adds the "filters" edge to the DbFilter entity by IDs.
func (duuo *DbUserUpdateOne) AddFilterIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.AddFilterIDs(ids...)
	return duuo
}

// AddFilters adds the "filters" edges to the DbFilter entity.
func (duuo *DbUserUpdateOne) AddFilters(d ...*DbFilter) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.AddFilterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the DbGroup entity by IDs.
func (duuo *DbUserUpdateOne) AddGroupIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.AddGroupIDs(ids...)
	return duuo
}

// AddGroups adds the "groups" edges to the DbGroup entity.
func (duuo *DbUserUpdateOne) AddGroups(d ...*DbGroup) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.AddGroupIDs(ids...)
}

// AddTransportRecipientIDs adds the "TransportRecipients" edge to the DbTransportRecipients entity by IDs.
func (duuo *DbUserUpdateOne) AddTransportRecipientIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.AddTransportRecipientIDs(ids...)
	return duuo
}

// AddTransportRecipients adds the "TransportRecipients" edges to the DbTransportRecipients entity.
func (duuo *DbUserUpdateOne) AddTransportRecipients(d ...*DbTransportRecipients) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.AddTransportRecipientIDs(ids...)
}

// Mutation returns the DbUserMutation object of the builder.
func (duuo *DbUserUpdateOne) Mutation() *DbUserMutation {
	return duuo.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (duuo *DbUserUpdateOne) ClearTenant() *DbUserUpdateOne {
	duuo.mutation.ClearTenant()
	return duuo
}

// ClearMetadata clears all "metadata" edges to the DbUserMetaData entity.
func (duuo *DbUserUpdateOne) ClearMetadata() *DbUserUpdateOne {
	duuo.mutation.ClearMetadata()
	return duuo
}

// RemoveMetadatumIDs removes the "metadata" edge to DbUserMetaData entities by IDs.
func (duuo *DbUserUpdateOne) RemoveMetadatumIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.RemoveMetadatumIDs(ids...)
	return duuo
}

// RemoveMetadata removes "metadata" edges to DbUserMetaData entities.
func (duuo *DbUserUpdateOne) RemoveMetadata(d ...*DbUserMetaData) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.RemoveMetadatumIDs(ids...)
}

// ClearFilters clears all "filters" edges to the DbFilter entity.
func (duuo *DbUserUpdateOne) ClearFilters() *DbUserUpdateOne {
	duuo.mutation.ClearFilters()
	return duuo
}

// RemoveFilterIDs removes the "filters" edge to DbFilter entities by IDs.
func (duuo *DbUserUpdateOne) RemoveFilterIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.RemoveFilterIDs(ids...)
	return duuo
}

// RemoveFilters removes "filters" edges to DbFilter entities.
func (duuo *DbUserUpdateOne) RemoveFilters(d ...*DbFilter) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.RemoveFilterIDs(ids...)
}

// ClearGroups clears all "groups" edges to the DbGroup entity.
func (duuo *DbUserUpdateOne) ClearGroups() *DbUserUpdateOne {
	duuo.mutation.ClearGroups()
	return duuo
}

// RemoveGroupIDs removes the "groups" edge to DbGroup entities by IDs.
func (duuo *DbUserUpdateOne) RemoveGroupIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.RemoveGroupIDs(ids...)
	return duuo
}

// RemoveGroups removes "groups" edges to DbGroup entities.
func (duuo *DbUserUpdateOne) RemoveGroups(d ...*DbGroup) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.RemoveGroupIDs(ids...)
}

// ClearTransportRecipients clears all "TransportRecipients" edges to the DbTransportRecipients entity.
func (duuo *DbUserUpdateOne) ClearTransportRecipients() *DbUserUpdateOne {
	duuo.mutation.ClearTransportRecipients()
	return duuo
}

// RemoveTransportRecipientIDs removes the "TransportRecipients" edge to DbTransportRecipients entities by IDs.
func (duuo *DbUserUpdateOne) RemoveTransportRecipientIDs(ids ...int) *DbUserUpdateOne {
	duuo.mutation.RemoveTransportRecipientIDs(ids...)
	return duuo
}

// RemoveTransportRecipients removes "TransportRecipients" edges to DbTransportRecipients entities.
func (duuo *DbUserUpdateOne) RemoveTransportRecipients(d ...*DbTransportRecipients) *DbUserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.RemoveTransportRecipientIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duuo *DbUserUpdateOne) Select(field string, fields ...string) *DbUserUpdateOne {
	duuo.fields = append([]string{field}, fields...)
	return duuo
}

// Save executes the query and returns the updated DbUser entity.
func (duuo *DbUserUpdateOne) Save(ctx context.Context) (*DbUser, error) {
	var (
		err  error
		node *DbUser
	)
	if len(duuo.hooks) == 0 {
		if err = duuo.check(); err != nil {
			return nil, err
		}
		node, err = duuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duuo.check(); err != nil {
				return nil, err
			}
			duuo.mutation = mutation
			node, err = duuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duuo.hooks) - 1; i >= 0; i-- {
			if duuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DbUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DbUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duuo *DbUserUpdateOne) SaveX(ctx context.Context) *DbUser {
	node, err := duuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duuo *DbUserUpdateOne) Exec(ctx context.Context) error {
	_, err := duuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duuo *DbUserUpdateOne) ExecX(ctx context.Context) {
	if err := duuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duuo *DbUserUpdateOne) check() error {
	if v, ok := duuo.mutation.Email(); ok {
		if err := dbuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`ent: validator failed for field "DbUser.Email": %w`, err)}
		}
	}
	if v, ok := duuo.mutation.Name(); ok {
		if err := dbuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbUser.Name": %w`, err)}
		}
	}
	if _, ok := duuo.mutation.TenantID(); duuo.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DbUser.tenant"`)
	}
	return nil
}

func (duuo *DbUserUpdateOne) sqlSave(ctx context.Context) (_node *DbUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbuser.Table,
			Columns: dbuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbuser.FieldID,
			},
		},
	}
	id, ok := duuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DbUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbuser.FieldID)
		for _, f := range fields {
			if !dbuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dbuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duuo.mutation.AppData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbuser.FieldAppData,
		})
	}
	if duuo.mutation.AppDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dbuser.FieldAppData,
		})
	}
	if value, ok := duuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldEmail,
		})
	}
	if value, ok := duuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldName,
		})
	}
	if value, ok := duuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldDescription,
		})
	}
	if duuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbuser.FieldDescription,
		})
	}
	if duuo.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbuser.TenantTable,
			Columns: []string{dbuser.TenantColumn},
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
	if nodes := duuo.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbuser.TenantTable,
			Columns: []string{dbuser.TenantColumn},
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
	if duuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.MetadataTable,
			Columns: []string{dbuser.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbusermetadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !duuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.MetadataTable,
			Columns: []string{dbuser.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbusermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.MetadataTable,
			Columns: []string{dbuser.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbusermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duuo.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.FiltersTable,
			Columns: dbuser.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbfilter.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.RemovedFiltersIDs(); len(nodes) > 0 && !duuo.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.FiltersTable,
			Columns: dbuser.FiltersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.FiltersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.FiltersTable,
			Columns: dbuser.FiltersPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.GroupsTable,
			Columns: dbuser.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !duuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.GroupsTable,
			Columns: dbuser.GroupsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dbuser.GroupsTable,
			Columns: dbuser.GroupsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duuo.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.TransportRecipientsTable,
			Columns: []string{dbuser.TransportRecipientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbtransportrecipients.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.RemovedTransportRecipientsIDs(); len(nodes) > 0 && !duuo.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.TransportRecipientsTable,
			Columns: []string{dbuser.TransportRecipientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbtransportrecipients.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.TransportRecipientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbuser.TransportRecipientsTable,
			Columns: []string{dbuser.TransportRecipientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbtransportrecipients.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DbUser{config: duuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}