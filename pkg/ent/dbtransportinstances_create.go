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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportinstances"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

// DbTransportInstancesCreate is the builder for creating a DbTransportInstances entity.
type DbTransportInstancesCreate struct {
	config
	mutation *DbTransportInstancesMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTenantID sets the "tenant_id" field.
func (dtic *DbTransportInstancesCreate) SetTenantID(i int) *DbTransportInstancesCreate {
	dtic.mutation.SetTenantID(i)
	return dtic
}

// SetAppData sets the "AppData" field.
func (dtic *DbTransportInstancesCreate) SetAppData(id interfaces.AppData) *DbTransportInstancesCreate {
	dtic.mutation.SetAppData(id)
	return dtic
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (dtic *DbTransportInstancesCreate) SetNillableAppData(id *interfaces.AppData) *DbTransportInstancesCreate {
	if id != nil {
		dtic.SetAppData(*id)
	}
	return dtic
}

// SetName sets the "Name" field.
func (dtic *DbTransportInstancesCreate) SetName(s string) *DbTransportInstancesCreate {
	dtic.mutation.SetName(s)
	return dtic
}

// SetDescription sets the "Description" field.
func (dtic *DbTransportInstancesCreate) SetDescription(s string) *DbTransportInstancesCreate {
	dtic.mutation.SetDescription(s)
	return dtic
}

// SetNillableDescription sets the "Description" field if the given value is not nil.
func (dtic *DbTransportInstancesCreate) SetNillableDescription(s *string) *DbTransportInstancesCreate {
	if s != nil {
		dtic.SetDescription(*s)
	}
	return dtic
}

// SetConfig sets the "Config" field.
func (dtic *DbTransportInstancesCreate) SetConfig(s string) *DbTransportInstancesCreate {
	dtic.mutation.SetConfig(s)
	return dtic
}

// SetTransportProvider sets the "TransportProvider" field.
func (dtic *DbTransportInstancesCreate) SetTransportProvider(s string) *DbTransportInstancesCreate {
	dtic.mutation.SetTransportProvider(s)
	return dtic
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (dtic *DbTransportInstancesCreate) SetTenant(t *Tenant) *DbTransportInstancesCreate {
	return dtic.SetTenantID(t.ID)
}

// AddTransportRecipientIDs adds the "TransportRecipients" edge to the DbTransportRecipients entity by IDs.
func (dtic *DbTransportInstancesCreate) AddTransportRecipientIDs(ids ...int) *DbTransportInstancesCreate {
	dtic.mutation.AddTransportRecipientIDs(ids...)
	return dtic
}

// AddTransportRecipients adds the "TransportRecipients" edges to the DbTransportRecipients entity.
func (dtic *DbTransportInstancesCreate) AddTransportRecipients(d ...*DbTransportRecipients) *DbTransportInstancesCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtic.AddTransportRecipientIDs(ids...)
}

// Mutation returns the DbTransportInstancesMutation object of the builder.
func (dtic *DbTransportInstancesCreate) Mutation() *DbTransportInstancesMutation {
	return dtic.mutation
}

// Save creates the DbTransportInstances in the database.
func (dtic *DbTransportInstancesCreate) Save(ctx context.Context) (*DbTransportInstances, error) {
	var (
		err  error
		node *DbTransportInstances
	)
	if len(dtic.hooks) == 0 {
		if err = dtic.check(); err != nil {
			return nil, err
		}
		node, err = dtic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbTransportInstancesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dtic.check(); err != nil {
				return nil, err
			}
			dtic.mutation = mutation
			if node, err = dtic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dtic.hooks) - 1; i >= 0; i-- {
			if dtic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dtic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dtic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DbTransportInstances)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DbTransportInstancesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dtic *DbTransportInstancesCreate) SaveX(ctx context.Context) *DbTransportInstances {
	v, err := dtic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtic *DbTransportInstancesCreate) Exec(ctx context.Context) error {
	_, err := dtic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtic *DbTransportInstancesCreate) ExecX(ctx context.Context) {
	if err := dtic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtic *DbTransportInstancesCreate) check() error {
	if _, ok := dtic.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "DbTransportInstances.tenant_id"`)}
	}
	if _, ok := dtic.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "DbTransportInstances.Name"`)}
	}
	if v, ok := dtic.mutation.Name(); ok {
		if err := dbtransportinstances.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbTransportInstances.Name": %w`, err)}
		}
	}
	if _, ok := dtic.mutation.Config(); !ok {
		return &ValidationError{Name: "Config", err: errors.New(`ent: missing required field "DbTransportInstances.Config"`)}
	}
	if v, ok := dtic.mutation.Config(); ok {
		if err := dbtransportinstances.ConfigValidator(v); err != nil {
			return &ValidationError{Name: "Config", err: fmt.Errorf(`ent: validator failed for field "DbTransportInstances.Config": %w`, err)}
		}
	}
	if _, ok := dtic.mutation.TransportProvider(); !ok {
		return &ValidationError{Name: "TransportProvider", err: errors.New(`ent: missing required field "DbTransportInstances.TransportProvider"`)}
	}
	if v, ok := dtic.mutation.TransportProvider(); ok {
		if err := dbtransportinstances.TransportProviderValidator(v); err != nil {
			return &ValidationError{Name: "TransportProvider", err: fmt.Errorf(`ent: validator failed for field "DbTransportInstances.TransportProvider": %w`, err)}
		}
	}
	if _, ok := dtic.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`ent: missing required edge "DbTransportInstances.tenant"`)}
	}
	return nil
}

func (dtic *DbTransportInstancesCreate) sqlSave(ctx context.Context) (*DbTransportInstances, error) {
	_node, _spec := dtic.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dtic *DbTransportInstancesCreate) createSpec() (*DbTransportInstances, *sqlgraph.CreateSpec) {
	var (
		_node = &DbTransportInstances{config: dtic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dbtransportinstances.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbtransportinstances.FieldID,
			},
		}
	)
	_spec.OnConflict = dtic.conflict
	if value, ok := dtic.mutation.AppData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbtransportinstances.FieldAppData,
		})
		_node.AppData = value
	}
	if value, ok := dtic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbtransportinstances.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dtic.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbtransportinstances.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := dtic.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbtransportinstances.FieldConfig,
		})
		_node.Config = value
	}
	if value, ok := dtic.mutation.TransportProvider(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbtransportinstances.FieldTransportProvider,
		})
		_node.TransportProvider = value
	}
	if nodes := dtic.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbtransportinstances.TenantTable,
			Columns: []string{dbtransportinstances.TenantColumn},
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
	if nodes := dtic.mutation.TransportRecipientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dbtransportinstances.TransportRecipientsTable,
			Columns: []string{dbtransportinstances.TransportRecipientsColumn},
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
//	client.DbTransportInstances.Create().
//		SetTenantID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbTransportInstancesUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
//
func (dtic *DbTransportInstancesCreate) OnConflict(opts ...sql.ConflictOption) *DbTransportInstancesUpsertOne {
	dtic.conflict = opts
	return &DbTransportInstancesUpsertOne{
		create: dtic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbTransportInstances.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dtic *DbTransportInstancesCreate) OnConflictColumns(columns ...string) *DbTransportInstancesUpsertOne {
	dtic.conflict = append(dtic.conflict, sql.ConflictColumns(columns...))
	return &DbTransportInstancesUpsertOne{
		create: dtic,
	}
}

type (
	// DbTransportInstancesUpsertOne is the builder for "upsert"-ing
	//  one DbTransportInstances node.
	DbTransportInstancesUpsertOne struct {
		create *DbTransportInstancesCreate
	}

	// DbTransportInstancesUpsert is the "OnConflict" setter.
	DbTransportInstancesUpsert struct {
		*sql.UpdateSet
	}
)

// SetTenantID sets the "tenant_id" field.
func (u *DbTransportInstancesUpsert) SetTenantID(v int) *DbTransportInstancesUpsert {
	u.Set(dbtransportinstances.FieldTenantID, v)
	return u
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbTransportInstancesUpsert) UpdateTenantID() *DbTransportInstancesUpsert {
	u.SetExcluded(dbtransportinstances.FieldTenantID)
	return u
}

// SetAppData sets the "AppData" field.
func (u *DbTransportInstancesUpsert) SetAppData(v interfaces.AppData) *DbTransportInstancesUpsert {
	u.Set(dbtransportinstances.FieldAppData, v)
	return u
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbTransportInstancesUpsert) UpdateAppData() *DbTransportInstancesUpsert {
	u.SetExcluded(dbtransportinstances.FieldAppData)
	return u
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbTransportInstancesUpsert) ClearAppData() *DbTransportInstancesUpsert {
	u.SetNull(dbtransportinstances.FieldAppData)
	return u
}

// SetName sets the "Name" field.
func (u *DbTransportInstancesUpsert) SetName(v string) *DbTransportInstancesUpsert {
	u.Set(dbtransportinstances.FieldName, v)
	return u
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbTransportInstancesUpsert) UpdateName() *DbTransportInstancesUpsert {
	u.SetExcluded(dbtransportinstances.FieldName)
	return u
}

// SetDescription sets the "Description" field.
func (u *DbTransportInstancesUpsert) SetDescription(v string) *DbTransportInstancesUpsert {
	u.Set(dbtransportinstances.FieldDescription, v)
	return u
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbTransportInstancesUpsert) UpdateDescription() *DbTransportInstancesUpsert {
	u.SetExcluded(dbtransportinstances.FieldDescription)
	return u
}

// ClearDescription clears the value of the "Description" field.
func (u *DbTransportInstancesUpsert) ClearDescription() *DbTransportInstancesUpsert {
	u.SetNull(dbtransportinstances.FieldDescription)
	return u
}

// SetConfig sets the "Config" field.
func (u *DbTransportInstancesUpsert) SetConfig(v string) *DbTransportInstancesUpsert {
	u.Set(dbtransportinstances.FieldConfig, v)
	return u
}

// UpdateConfig sets the "Config" field to the value that was provided on create.
func (u *DbTransportInstancesUpsert) UpdateConfig() *DbTransportInstancesUpsert {
	u.SetExcluded(dbtransportinstances.FieldConfig)
	return u
}

// SetTransportProvider sets the "TransportProvider" field.
func (u *DbTransportInstancesUpsert) SetTransportProvider(v string) *DbTransportInstancesUpsert {
	u.Set(dbtransportinstances.FieldTransportProvider, v)
	return u
}

// UpdateTransportProvider sets the "TransportProvider" field to the value that was provided on create.
func (u *DbTransportInstancesUpsert) UpdateTransportProvider() *DbTransportInstancesUpsert {
	u.SetExcluded(dbtransportinstances.FieldTransportProvider)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DbTransportInstances.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DbTransportInstancesUpsertOne) UpdateNewValues() *DbTransportInstancesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.DbTransportInstances.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DbTransportInstancesUpsertOne) Ignore() *DbTransportInstancesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbTransportInstancesUpsertOne) DoNothing() *DbTransportInstancesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbTransportInstancesCreate.OnConflict
// documentation for more info.
func (u *DbTransportInstancesUpsertOne) Update(set func(*DbTransportInstancesUpsert)) *DbTransportInstancesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbTransportInstancesUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbTransportInstancesUpsertOne) SetTenantID(v int) *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertOne) UpdateTenantID() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbTransportInstancesUpsertOne) SetAppData(v interfaces.AppData) *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertOne) UpdateAppData() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbTransportInstancesUpsertOne) ClearAppData() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.ClearAppData()
	})
}

// SetName sets the "Name" field.
func (u *DbTransportInstancesUpsertOne) SetName(v string) *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertOne) UpdateName() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "Description" field.
func (u *DbTransportInstancesUpsertOne) SetDescription(v string) *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertOne) UpdateDescription() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "Description" field.
func (u *DbTransportInstancesUpsertOne) ClearDescription() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.ClearDescription()
	})
}

// SetConfig sets the "Config" field.
func (u *DbTransportInstancesUpsertOne) SetConfig(v string) *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "Config" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertOne) UpdateConfig() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateConfig()
	})
}

// SetTransportProvider sets the "TransportProvider" field.
func (u *DbTransportInstancesUpsertOne) SetTransportProvider(v string) *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetTransportProvider(v)
	})
}

// UpdateTransportProvider sets the "TransportProvider" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertOne) UpdateTransportProvider() *DbTransportInstancesUpsertOne {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateTransportProvider()
	})
}

// Exec executes the query.
func (u *DbTransportInstancesUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbTransportInstancesCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbTransportInstancesUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DbTransportInstancesUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DbTransportInstancesUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (dtic *DbTransportInstancesCreate) SetDbTransportInstancesFromStruct(input *DbTransportInstances) *DbTransportInstancesCreate {

	dtic.SetTenantID(input.TenantID)

	dtic.SetAppData(input.AppData)

	dtic.SetName(input.Name)

	dtic.SetDescription(input.Description)

	dtic.SetConfig(input.Config)

	dtic.SetTransportProvider(input.TransportProvider)

	return dtic
}

// DbTransportInstancesCreateBulk is the builder for creating many DbTransportInstances entities in bulk.
type DbTransportInstancesCreateBulk struct {
	config
	builders []*DbTransportInstancesCreate
	conflict []sql.ConflictOption
}

// Save creates the DbTransportInstances entities in the database.
func (dticb *DbTransportInstancesCreateBulk) Save(ctx context.Context) ([]*DbTransportInstances, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dticb.builders))
	nodes := make([]*DbTransportInstances, len(dticb.builders))
	mutators := make([]Mutator, len(dticb.builders))
	for i := range dticb.builders {
		func(i int, root context.Context) {
			builder := dticb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DbTransportInstancesMutation)
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
					_, err = mutators[i+1].Mutate(root, dticb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dticb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dticb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dticb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dticb *DbTransportInstancesCreateBulk) SaveX(ctx context.Context) []*DbTransportInstances {
	v, err := dticb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dticb *DbTransportInstancesCreateBulk) Exec(ctx context.Context) error {
	_, err := dticb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dticb *DbTransportInstancesCreateBulk) ExecX(ctx context.Context) {
	if err := dticb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DbTransportInstances.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbTransportInstancesUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
//
func (dticb *DbTransportInstancesCreateBulk) OnConflict(opts ...sql.ConflictOption) *DbTransportInstancesUpsertBulk {
	dticb.conflict = opts
	return &DbTransportInstancesUpsertBulk{
		create: dticb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbTransportInstances.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dticb *DbTransportInstancesCreateBulk) OnConflictColumns(columns ...string) *DbTransportInstancesUpsertBulk {
	dticb.conflict = append(dticb.conflict, sql.ConflictColumns(columns...))
	return &DbTransportInstancesUpsertBulk{
		create: dticb,
	}
}

// DbTransportInstancesUpsertBulk is the builder for "upsert"-ing
// a bulk of DbTransportInstances nodes.
type DbTransportInstancesUpsertBulk struct {
	create *DbTransportInstancesCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DbTransportInstances.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DbTransportInstancesUpsertBulk) UpdateNewValues() *DbTransportInstancesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DbTransportInstances.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DbTransportInstancesUpsertBulk) Ignore() *DbTransportInstancesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbTransportInstancesUpsertBulk) DoNothing() *DbTransportInstancesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbTransportInstancesCreateBulk.OnConflict
// documentation for more info.
func (u *DbTransportInstancesUpsertBulk) Update(set func(*DbTransportInstancesUpsert)) *DbTransportInstancesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbTransportInstancesUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbTransportInstancesUpsertBulk) SetTenantID(v int) *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertBulk) UpdateTenantID() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbTransportInstancesUpsertBulk) SetAppData(v interfaces.AppData) *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertBulk) UpdateAppData() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbTransportInstancesUpsertBulk) ClearAppData() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.ClearAppData()
	})
}

// SetName sets the "Name" field.
func (u *DbTransportInstancesUpsertBulk) SetName(v string) *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertBulk) UpdateName() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "Description" field.
func (u *DbTransportInstancesUpsertBulk) SetDescription(v string) *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertBulk) UpdateDescription() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "Description" field.
func (u *DbTransportInstancesUpsertBulk) ClearDescription() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.ClearDescription()
	})
}

// SetConfig sets the "Config" field.
func (u *DbTransportInstancesUpsertBulk) SetConfig(v string) *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "Config" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertBulk) UpdateConfig() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateConfig()
	})
}

// SetTransportProvider sets the "TransportProvider" field.
func (u *DbTransportInstancesUpsertBulk) SetTransportProvider(v string) *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.SetTransportProvider(v)
	})
}

// UpdateTransportProvider sets the "TransportProvider" field to the value that was provided on create.
func (u *DbTransportInstancesUpsertBulk) UpdateTransportProvider() *DbTransportInstancesUpsertBulk {
	return u.Update(func(s *DbTransportInstancesUpsert) {
		s.UpdateTransportProvider()
	})
}

// Exec executes the query.
func (u *DbTransportInstancesUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DbTransportInstancesCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbTransportInstancesCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbTransportInstancesUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
