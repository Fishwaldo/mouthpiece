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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/group"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportrecipient"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/user"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/usermetadata"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetTenantID sets the "tenant_id" field.
func (uu *UserUpdate) SetTenantID(i int) *UserUpdate {
	uu.mutation.SetTenantID(i)
	return uu
}

// SetEmail sets the "Email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetName sets the "Name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetDescription sets the "Description" field.
func (uu *UserUpdate) SetDescription(s string) *UserUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (uu *UserUpdate) SetTenant(t *Tenant) *UserUpdate {
	return uu.SetTenantID(t.ID)
}

// AddMetadatumIDs adds the "metadata" edge to the UserMetaData entity by IDs.
func (uu *UserUpdate) AddMetadatumIDs(ids ...int) *UserUpdate {
	uu.mutation.AddMetadatumIDs(ids...)
	return uu
}

// AddMetadata adds the "metadata" edges to the UserMetaData entity.
func (uu *UserUpdate) AddMetadata(u ...*UserMetaData) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddMetadatumIDs(ids...)
}

// AddFilterIDs adds the "filters" edge to the Filter entity by IDs.
func (uu *UserUpdate) AddFilterIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFilterIDs(ids...)
	return uu
}

// AddFilters adds the "filters" edges to the Filter entity.
func (uu *UserUpdate) AddFilters(f ...*Filter) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.AddFilterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the Group entity.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// AddTransportRecipientIDs adds the "TransportRecipients" edge to the TransportRecipient entity by IDs.
func (uu *UserUpdate) AddTransportRecipientIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTransportRecipientIDs(ids...)
	return uu
}

// AddTransportRecipients adds the "TransportRecipients" edges to the TransportRecipient entity.
func (uu *UserUpdate) AddTransportRecipients(t ...*TransportRecipient) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTransportRecipientIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (uu *UserUpdate) ClearTenant() *UserUpdate {
	uu.mutation.ClearTenant()
	return uu
}

// ClearMetadata clears all "metadata" edges to the UserMetaData entity.
func (uu *UserUpdate) ClearMetadata() *UserUpdate {
	uu.mutation.ClearMetadata()
	return uu
}

// RemoveMetadatumIDs removes the "metadata" edge to UserMetaData entities by IDs.
func (uu *UserUpdate) RemoveMetadatumIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveMetadatumIDs(ids...)
	return uu
}

// RemoveMetadata removes "metadata" edges to UserMetaData entities.
func (uu *UserUpdate) RemoveMetadata(u ...*UserMetaData) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveMetadatumIDs(ids...)
}

// ClearFilters clears all "filters" edges to the Filter entity.
func (uu *UserUpdate) ClearFilters() *UserUpdate {
	uu.mutation.ClearFilters()
	return uu
}

// RemoveFilterIDs removes the "filters" edge to Filter entities by IDs.
func (uu *UserUpdate) RemoveFilterIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFilterIDs(ids...)
	return uu
}

// RemoveFilters removes "filters" edges to Filter entities.
func (uu *UserUpdate) RemoveFilters(f ...*Filter) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.RemoveFilterIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to Group entities.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// ClearTransportRecipients clears all "TransportRecipients" edges to the TransportRecipient entity.
func (uu *UserUpdate) ClearTransportRecipients() *UserUpdate {
	uu.mutation.ClearTransportRecipients()
	return uu
}

// RemoveTransportRecipientIDs removes the "TransportRecipients" edge to TransportRecipient entities by IDs.
func (uu *UserUpdate) RemoveTransportRecipientIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTransportRecipientIDs(ids...)
	return uu
}

// RemoveTransportRecipients removes "TransportRecipients" edges to TransportRecipient entities.
func (uu *UserUpdate) RemoveTransportRecipients(t ...*TransportRecipient) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTransportRecipientIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`ent: validator failed for field "User.Email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "User.Name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Description(); ok {
		if err := user.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "Description", err: fmt.Errorf(`ent: validator failed for field "User.Description": %w`, err)}
		}
	}
	if _, ok := uu.mutation.TenantID(); uu.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.tenant"`)
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDescription,
		})
	}
	if uu.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.TenantTable,
			Columns: []string{user.TenantColumn},
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
	if nodes := uu.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.TenantTable,
			Columns: []string{user.TenantColumn},
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
	if uu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetadataTable,
			Columns: []string{user.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usermetadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !uu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetadataTable,
			Columns: []string{user.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetadataTable,
			Columns: []string{user.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FiltersTable,
			Columns: user.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filter.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFiltersIDs(); len(nodes) > 0 && !uu.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FiltersTable,
			Columns: user.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FiltersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FiltersTable,
			Columns: user.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	if uu.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TransportRecipientsTable,
			Columns: user.TransportRecipientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportrecipient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTransportRecipientsIDs(); len(nodes) > 0 && !uu.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TransportRecipientsTable,
			Columns: user.TransportRecipientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportrecipient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TransportRecipientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TransportRecipientsTable,
			Columns: user.TransportRecipientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportrecipient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetTenantID sets the "tenant_id" field.
func (uuo *UserUpdateOne) SetTenantID(i int) *UserUpdateOne {
	uuo.mutation.SetTenantID(i)
	return uuo
}

// SetEmail sets the "Email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetName sets the "Name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetDescription sets the "Description" field.
func (uuo *UserUpdateOne) SetDescription(s string) *UserUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (uuo *UserUpdateOne) SetTenant(t *Tenant) *UserUpdateOne {
	return uuo.SetTenantID(t.ID)
}

// AddMetadatumIDs adds the "metadata" edge to the UserMetaData entity by IDs.
func (uuo *UserUpdateOne) AddMetadatumIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddMetadatumIDs(ids...)
	return uuo
}

// AddMetadata adds the "metadata" edges to the UserMetaData entity.
func (uuo *UserUpdateOne) AddMetadata(u ...*UserMetaData) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddMetadatumIDs(ids...)
}

// AddFilterIDs adds the "filters" edge to the Filter entity by IDs.
func (uuo *UserUpdateOne) AddFilterIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFilterIDs(ids...)
	return uuo
}

// AddFilters adds the "filters" edges to the Filter entity.
func (uuo *UserUpdateOne) AddFilters(f ...*Filter) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.AddFilterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// AddTransportRecipientIDs adds the "TransportRecipients" edge to the TransportRecipient entity by IDs.
func (uuo *UserUpdateOne) AddTransportRecipientIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTransportRecipientIDs(ids...)
	return uuo
}

// AddTransportRecipients adds the "TransportRecipients" edges to the TransportRecipient entity.
func (uuo *UserUpdateOne) AddTransportRecipients(t ...*TransportRecipient) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTransportRecipientIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (uuo *UserUpdateOne) ClearTenant() *UserUpdateOne {
	uuo.mutation.ClearTenant()
	return uuo
}

// ClearMetadata clears all "metadata" edges to the UserMetaData entity.
func (uuo *UserUpdateOne) ClearMetadata() *UserUpdateOne {
	uuo.mutation.ClearMetadata()
	return uuo
}

// RemoveMetadatumIDs removes the "metadata" edge to UserMetaData entities by IDs.
func (uuo *UserUpdateOne) RemoveMetadatumIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveMetadatumIDs(ids...)
	return uuo
}

// RemoveMetadata removes "metadata" edges to UserMetaData entities.
func (uuo *UserUpdateOne) RemoveMetadata(u ...*UserMetaData) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveMetadatumIDs(ids...)
}

// ClearFilters clears all "filters" edges to the Filter entity.
func (uuo *UserUpdateOne) ClearFilters() *UserUpdateOne {
	uuo.mutation.ClearFilters()
	return uuo
}

// RemoveFilterIDs removes the "filters" edge to Filter entities by IDs.
func (uuo *UserUpdateOne) RemoveFilterIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFilterIDs(ids...)
	return uuo
}

// RemoveFilters removes "filters" edges to Filter entities.
func (uuo *UserUpdateOne) RemoveFilters(f ...*Filter) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.RemoveFilterIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// ClearTransportRecipients clears all "TransportRecipients" edges to the TransportRecipient entity.
func (uuo *UserUpdateOne) ClearTransportRecipients() *UserUpdateOne {
	uuo.mutation.ClearTransportRecipients()
	return uuo
}

// RemoveTransportRecipientIDs removes the "TransportRecipients" edge to TransportRecipient entities by IDs.
func (uuo *UserUpdateOne) RemoveTransportRecipientIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTransportRecipientIDs(ids...)
	return uuo
}

// RemoveTransportRecipients removes "TransportRecipients" edges to TransportRecipient entities.
func (uuo *UserUpdateOne) RemoveTransportRecipients(t ...*TransportRecipient) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTransportRecipientIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`ent: validator failed for field "User.Email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "User.Name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Description(); ok {
		if err := user.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "Description", err: fmt.Errorf(`ent: validator failed for field "User.Description": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.TenantID(); uuo.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.tenant"`)
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDescription,
		})
	}
	if uuo.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.TenantTable,
			Columns: []string{user.TenantColumn},
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
	if nodes := uuo.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.TenantTable,
			Columns: []string{user.TenantColumn},
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
	if uuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetadataTable,
			Columns: []string{user.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usermetadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !uuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetadataTable,
			Columns: []string{user.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetadataTable,
			Columns: []string{user.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usermetadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FiltersTable,
			Columns: user.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filter.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFiltersIDs(); len(nodes) > 0 && !uuo.mutation.FiltersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FiltersTable,
			Columns: user.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FiltersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FiltersTable,
			Columns: user.FiltersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	if uuo.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TransportRecipientsTable,
			Columns: user.TransportRecipientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportrecipient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTransportRecipientsIDs(); len(nodes) > 0 && !uuo.mutation.TransportRecipientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TransportRecipientsTable,
			Columns: user.TransportRecipientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportrecipient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TransportRecipientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TransportRecipientsTable,
			Columns: user.TransportRecipientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transportrecipient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}