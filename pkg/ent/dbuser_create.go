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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

// DbUserCreate is the builder for creating a DbUser entity.
type DbUserCreate struct {
	config
	mutation *DbUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTenantID sets the "tenant_id" field.
func (duc *DbUserCreate) SetTenantID(i int) *DbUserCreate {
	duc.mutation.SetTenantID(i)
	return duc
}

// SetAppData sets the "AppData" field.
func (duc *DbUserCreate) SetAppData(id interfaces.AppData) *DbUserCreate {
	duc.mutation.SetAppData(id)
	return duc
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (duc *DbUserCreate) SetNillableAppData(id *interfaces.AppData) *DbUserCreate {
	if id != nil {
		duc.SetAppData(*id)
	}
	return duc
}

// SetEmail sets the "Email" field.
func (duc *DbUserCreate) SetEmail(s string) *DbUserCreate {
	duc.mutation.SetEmail(s)
	return duc
}

// SetName sets the "Name" field.
func (duc *DbUserCreate) SetName(s string) *DbUserCreate {
	duc.mutation.SetName(s)
	return duc
}

// SetDescription sets the "Description" field.
func (duc *DbUserCreate) SetDescription(s string) *DbUserCreate {
	duc.mutation.SetDescription(s)
	return duc
}

// SetNillableDescription sets the "Description" field if the given value is not nil.
func (duc *DbUserCreate) SetNillableDescription(s *string) *DbUserCreate {
	if s != nil {
		duc.SetDescription(*s)
	}
	return duc
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (duc *DbUserCreate) SetTenant(t *Tenant) *DbUserCreate {
	return duc.SetTenantID(t.ID)
}

// AddMetadatumIDs adds the "metadata" edge to the DbUserMetaData entity by IDs.
func (duc *DbUserCreate) AddMetadatumIDs(ids ...int) *DbUserCreate {
	duc.mutation.AddMetadatumIDs(ids...)
	return duc
}

// AddMetadata adds the "metadata" edges to the DbUserMetaData entity.
func (duc *DbUserCreate) AddMetadata(d ...*DbUserMetaData) *DbUserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duc.AddMetadatumIDs(ids...)
}

// AddFilterIDs adds the "filters" edge to the DbFilter entity by IDs.
func (duc *DbUserCreate) AddFilterIDs(ids ...int) *DbUserCreate {
	duc.mutation.AddFilterIDs(ids...)
	return duc
}

// AddFilters adds the "filters" edges to the DbFilter entity.
func (duc *DbUserCreate) AddFilters(d ...*DbFilter) *DbUserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duc.AddFilterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the DbGroup entity by IDs.
func (duc *DbUserCreate) AddGroupIDs(ids ...int) *DbUserCreate {
	duc.mutation.AddGroupIDs(ids...)
	return duc
}

// AddGroups adds the "groups" edges to the DbGroup entity.
func (duc *DbUserCreate) AddGroups(d ...*DbGroup) *DbUserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duc.AddGroupIDs(ids...)
}

// AddTransportRecipientIDs adds the "TransportRecipients" edge to the DbTransportRecipients entity by IDs.
func (duc *DbUserCreate) AddTransportRecipientIDs(ids ...int) *DbUserCreate {
	duc.mutation.AddTransportRecipientIDs(ids...)
	return duc
}

// AddTransportRecipients adds the "TransportRecipients" edges to the DbTransportRecipients entity.
func (duc *DbUserCreate) AddTransportRecipients(d ...*DbTransportRecipients) *DbUserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duc.AddTransportRecipientIDs(ids...)
}

// Mutation returns the DbUserMutation object of the builder.
func (duc *DbUserCreate) Mutation() *DbUserMutation {
	return duc.mutation
}

// Save creates the DbUser in the database.
func (duc *DbUserCreate) Save(ctx context.Context) (*DbUser, error) {
	var (
		err  error
		node *DbUser
	)
	if len(duc.hooks) == 0 {
		if err = duc.check(); err != nil {
			return nil, err
		}
		node, err = duc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duc.check(); err != nil {
				return nil, err
			}
			duc.mutation = mutation
			if node, err = duc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(duc.hooks) - 1; i >= 0; i-- {
			if duc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (duc *DbUserCreate) SaveX(ctx context.Context) *DbUser {
	v, err := duc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (duc *DbUserCreate) Exec(ctx context.Context) error {
	_, err := duc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duc *DbUserCreate) ExecX(ctx context.Context) {
	if err := duc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duc *DbUserCreate) check() error {
	if _, ok := duc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "DbUser.tenant_id"`)}
	}
	if _, ok := duc.mutation.Email(); !ok {
		return &ValidationError{Name: "Email", err: errors.New(`ent: missing required field "DbUser.Email"`)}
	}
	if v, ok := duc.mutation.Email(); ok {
		if err := dbuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`ent: validator failed for field "DbUser.Email": %w`, err)}
		}
	}
	if _, ok := duc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "DbUser.Name"`)}
	}
	if v, ok := duc.mutation.Name(); ok {
		if err := dbuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbUser.Name": %w`, err)}
		}
	}
	if _, ok := duc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`ent: missing required edge "DbUser.tenant"`)}
	}
	return nil
}

func (duc *DbUserCreate) sqlSave(ctx context.Context) (*DbUser, error) {
	_node, _spec := duc.createSpec()
	if err := sqlgraph.CreateNode(ctx, duc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (duc *DbUserCreate) createSpec() (*DbUser, *sqlgraph.CreateSpec) {
	var (
		_node = &DbUser{config: duc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dbuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbuser.FieldID,
			},
		}
	)
	_spec.OnConflict = duc.conflict
	if value, ok := duc.mutation.AppData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbuser.FieldAppData,
		})
		_node.AppData = value
	}
	if value, ok := duc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := duc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldName,
		})
		_node.Name = value
	}
	if value, ok := duc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbuser.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := duc.mutation.TenantIDs(); len(nodes) > 0 {
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
		_node.TenantID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := duc.mutation.MetadataIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := duc.mutation.FiltersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := duc.mutation.GroupsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := duc.mutation.TransportRecipientsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DbUser.Create().
//		SetTenantID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbUserUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (duc *DbUserCreate) OnConflict(opts ...sql.ConflictOption) *DbUserUpsertOne {
	duc.conflict = opts
	return &DbUserUpsertOne{
		create: duc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (duc *DbUserCreate) OnConflictColumns(columns ...string) *DbUserUpsertOne {
	duc.conflict = append(duc.conflict, sql.ConflictColumns(columns...))
	return &DbUserUpsertOne{
		create: duc,
	}
}

type (
	// DbUserUpsertOne is the builder for "upsert"-ing
	//  one DbUser node.
	DbUserUpsertOne struct {
		create *DbUserCreate
	}

	// DbUserUpsert is the "OnConflict" setter.
	DbUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetTenantID sets the "tenant_id" field.
func (u *DbUserUpsert) SetTenantID(v int) *DbUserUpsert {
	u.Set(dbuser.FieldTenantID, v)
	return u
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbUserUpsert) UpdateTenantID() *DbUserUpsert {
	u.SetExcluded(dbuser.FieldTenantID)
	return u
}

// SetAppData sets the "AppData" field.
func (u *DbUserUpsert) SetAppData(v interfaces.AppData) *DbUserUpsert {
	u.Set(dbuser.FieldAppData, v)
	return u
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbUserUpsert) UpdateAppData() *DbUserUpsert {
	u.SetExcluded(dbuser.FieldAppData)
	return u
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbUserUpsert) ClearAppData() *DbUserUpsert {
	u.SetNull(dbuser.FieldAppData)
	return u
}

// SetEmail sets the "Email" field.
func (u *DbUserUpsert) SetEmail(v string) *DbUserUpsert {
	u.Set(dbuser.FieldEmail, v)
	return u
}

// UpdateEmail sets the "Email" field to the value that was provided on create.
func (u *DbUserUpsert) UpdateEmail() *DbUserUpsert {
	u.SetExcluded(dbuser.FieldEmail)
	return u
}

// SetName sets the "Name" field.
func (u *DbUserUpsert) SetName(v string) *DbUserUpsert {
	u.Set(dbuser.FieldName, v)
	return u
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbUserUpsert) UpdateName() *DbUserUpsert {
	u.SetExcluded(dbuser.FieldName)
	return u
}

// SetDescription sets the "Description" field.
func (u *DbUserUpsert) SetDescription(v string) *DbUserUpsert {
	u.Set(dbuser.FieldDescription, v)
	return u
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbUserUpsert) UpdateDescription() *DbUserUpsert {
	u.SetExcluded(dbuser.FieldDescription)
	return u
}

// ClearDescription clears the value of the "Description" field.
func (u *DbUserUpsert) ClearDescription() *DbUserUpsert {
	u.SetNull(dbuser.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DbUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DbUserUpsertOne) UpdateNewValues() *DbUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DbUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DbUserUpsertOne) Ignore() *DbUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbUserUpsertOne) DoNothing() *DbUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbUserCreate.OnConflict
// documentation for more info.
func (u *DbUserUpsertOne) Update(set func(*DbUserUpsert)) *DbUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbUserUpsertOne) SetTenantID(v int) *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbUserUpsertOne) UpdateTenantID() *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbUserUpsertOne) SetAppData(v interfaces.AppData) *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbUserUpsertOne) UpdateAppData() *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbUserUpsertOne) ClearAppData() *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.ClearAppData()
	})
}

// SetEmail sets the "Email" field.
func (u *DbUserUpsertOne) SetEmail(v string) *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "Email" field to the value that was provided on create.
func (u *DbUserUpsertOne) UpdateEmail() *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateEmail()
	})
}

// SetName sets the "Name" field.
func (u *DbUserUpsertOne) SetName(v string) *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbUserUpsertOne) UpdateName() *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "Description" field.
func (u *DbUserUpsertOne) SetDescription(v string) *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbUserUpsertOne) UpdateDescription() *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "Description" field.
func (u *DbUserUpsertOne) ClearDescription() *DbUserUpsertOne {
	return u.Update(func(s *DbUserUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *DbUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DbUserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DbUserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (duc *DbUserCreate) SetDbUserFromStruct(input *DbUser) *DbUserCreate {

	duc.SetTenantID(input.TenantID)

	duc.SetAppData(input.AppData)

	duc.SetEmail(input.Email)

	duc.SetName(input.Name)

	duc.SetDescription(input.Description)

	return duc
}

// DbUserCreateBulk is the builder for creating many DbUser entities in bulk.
type DbUserCreateBulk struct {
	config
	builders []*DbUserCreate
	conflict []sql.ConflictOption
}

// Save creates the DbUser entities in the database.
func (ducb *DbUserCreateBulk) Save(ctx context.Context) ([]*DbUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ducb.builders))
	nodes := make([]*DbUser, len(ducb.builders))
	mutators := make([]Mutator, len(ducb.builders))
	for i := range ducb.builders {
		func(i int, root context.Context) {
			builder := ducb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DbUserMutation)
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
					_, err = mutators[i+1].Mutate(root, ducb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ducb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ducb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ducb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ducb *DbUserCreateBulk) SaveX(ctx context.Context) []*DbUser {
	v, err := ducb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ducb *DbUserCreateBulk) Exec(ctx context.Context) error {
	_, err := ducb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ducb *DbUserCreateBulk) ExecX(ctx context.Context) {
	if err := ducb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DbUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbUserUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (ducb *DbUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *DbUserUpsertBulk {
	ducb.conflict = opts
	return &DbUserUpsertBulk{
		create: ducb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ducb *DbUserCreateBulk) OnConflictColumns(columns ...string) *DbUserUpsertBulk {
	ducb.conflict = append(ducb.conflict, sql.ConflictColumns(columns...))
	return &DbUserUpsertBulk{
		create: ducb,
	}
}

// DbUserUpsertBulk is the builder for "upsert"-ing
// a bulk of DbUser nodes.
type DbUserUpsertBulk struct {
	create *DbUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DbUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DbUserUpsertBulk) UpdateNewValues() *DbUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DbUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DbUserUpsertBulk) Ignore() *DbUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbUserUpsertBulk) DoNothing() *DbUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbUserCreateBulk.OnConflict
// documentation for more info.
func (u *DbUserUpsertBulk) Update(set func(*DbUserUpsert)) *DbUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbUserUpsertBulk) SetTenantID(v int) *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbUserUpsertBulk) UpdateTenantID() *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbUserUpsertBulk) SetAppData(v interfaces.AppData) *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbUserUpsertBulk) UpdateAppData() *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbUserUpsertBulk) ClearAppData() *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.ClearAppData()
	})
}

// SetEmail sets the "Email" field.
func (u *DbUserUpsertBulk) SetEmail(v string) *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "Email" field to the value that was provided on create.
func (u *DbUserUpsertBulk) UpdateEmail() *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateEmail()
	})
}

// SetName sets the "Name" field.
func (u *DbUserUpsertBulk) SetName(v string) *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbUserUpsertBulk) UpdateName() *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "Description" field.
func (u *DbUserUpsertBulk) SetDescription(v string) *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbUserUpsertBulk) UpdateDescription() *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "Description" field.
func (u *DbUserUpsertBulk) ClearDescription() *DbUserUpsertBulk {
	return u.Update(func(s *DbUserUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *DbUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DbUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}