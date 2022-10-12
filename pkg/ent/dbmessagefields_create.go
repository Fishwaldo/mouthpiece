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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessagefields"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/google/uuid"
)

// DbMessageFieldsCreate is the builder for creating a DbMessageFields entity.
type DbMessageFieldsCreate struct {
	config
	mutation *DbMessageFieldsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTenantID sets the "tenant_id" field.
func (dmfc *DbMessageFieldsCreate) SetTenantID(i int) *DbMessageFieldsCreate {
	dmfc.mutation.SetTenantID(i)
	return dmfc
}

// SetAppData sets the "AppData" field.
func (dmfc *DbMessageFieldsCreate) SetAppData(id interfaces.AppData) *DbMessageFieldsCreate {
	dmfc.mutation.SetAppData(id)
	return dmfc
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (dmfc *DbMessageFieldsCreate) SetNillableAppData(id *interfaces.AppData) *DbMessageFieldsCreate {
	if id != nil {
		dmfc.SetAppData(*id)
	}
	return dmfc
}

// SetName sets the "Name" field.
func (dmfc *DbMessageFieldsCreate) SetName(s string) *DbMessageFieldsCreate {
	dmfc.mutation.SetName(s)
	return dmfc
}

// SetValue sets the "Value" field.
func (dmfc *DbMessageFieldsCreate) SetValue(s string) *DbMessageFieldsCreate {
	dmfc.mutation.SetValue(s)
	return dmfc
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (dmfc *DbMessageFieldsCreate) SetTenant(t *Tenant) *DbMessageFieldsCreate {
	return dmfc.SetTenantID(t.ID)
}

// SetOwnerID sets the "owner" edge to the DbMessage entity by ID.
func (dmfc *DbMessageFieldsCreate) SetOwnerID(id uuid.UUID) *DbMessageFieldsCreate {
	dmfc.mutation.SetOwnerID(id)
	return dmfc
}

// SetOwner sets the "owner" edge to the DbMessage entity.
func (dmfc *DbMessageFieldsCreate) SetOwner(d *DbMessage) *DbMessageFieldsCreate {
	return dmfc.SetOwnerID(d.ID)
}

// Mutation returns the DbMessageFieldsMutation object of the builder.
func (dmfc *DbMessageFieldsCreate) Mutation() *DbMessageFieldsMutation {
	return dmfc.mutation
}

// Save creates the DbMessageFields in the database.
func (dmfc *DbMessageFieldsCreate) Save(ctx context.Context) (*DbMessageFields, error) {
	var (
		err  error
		node *DbMessageFields
	)
	if len(dmfc.hooks) == 0 {
		if err = dmfc.check(); err != nil {
			return nil, err
		}
		node, err = dmfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbMessageFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dmfc.check(); err != nil {
				return nil, err
			}
			dmfc.mutation = mutation
			if node, err = dmfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dmfc.hooks) - 1; i >= 0; i-- {
			if dmfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dmfc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dmfc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DbMessageFields)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DbMessageFieldsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dmfc *DbMessageFieldsCreate) SaveX(ctx context.Context) *DbMessageFields {
	v, err := dmfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmfc *DbMessageFieldsCreate) Exec(ctx context.Context) error {
	_, err := dmfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmfc *DbMessageFieldsCreate) ExecX(ctx context.Context) {
	if err := dmfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dmfc *DbMessageFieldsCreate) check() error {
	if _, ok := dmfc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "DbMessageFields.tenant_id"`)}
	}
	if _, ok := dmfc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "DbMessageFields.Name"`)}
	}
	if v, ok := dmfc.mutation.Name(); ok {
		if err := dbmessagefields.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbMessageFields.Name": %w`, err)}
		}
	}
	if _, ok := dmfc.mutation.Value(); !ok {
		return &ValidationError{Name: "Value", err: errors.New(`ent: missing required field "DbMessageFields.Value"`)}
	}
	if v, ok := dmfc.mutation.Value(); ok {
		if err := dbmessagefields.ValueValidator(v); err != nil {
			return &ValidationError{Name: "Value", err: fmt.Errorf(`ent: validator failed for field "DbMessageFields.Value": %w`, err)}
		}
	}
	if _, ok := dmfc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`ent: missing required edge "DbMessageFields.tenant"`)}
	}
	if _, ok := dmfc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "DbMessageFields.owner"`)}
	}
	return nil
}

func (dmfc *DbMessageFieldsCreate) sqlSave(ctx context.Context) (*DbMessageFields, error) {
	_node, _spec := dmfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dmfc *DbMessageFieldsCreate) createSpec() (*DbMessageFields, *sqlgraph.CreateSpec) {
	var (
		_node = &DbMessageFields{config: dmfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dbmessagefields.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbmessagefields.FieldID,
			},
		}
	)
	_spec.OnConflict = dmfc.conflict
	if value, ok := dmfc.mutation.AppData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbmessagefields.FieldAppData,
		})
		_node.AppData = value
	}
	if value, ok := dmfc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbmessagefields.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dmfc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbmessagefields.FieldValue,
		})
		_node.Value = value
	}
	if nodes := dmfc.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbmessagefields.TenantTable,
			Columns: []string{dbmessagefields.TenantColumn},
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
	if nodes := dmfc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbmessagefields.OwnerTable,
			Columns: []string{dbmessagefields.OwnerColumn},
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
		_node.db_message_fields = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DbMessageFields.Create().
//		SetTenantID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbMessageFieldsUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (dmfc *DbMessageFieldsCreate) OnConflict(opts ...sql.ConflictOption) *DbMessageFieldsUpsertOne {
	dmfc.conflict = opts
	return &DbMessageFieldsUpsertOne{
		create: dmfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbMessageFields.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmfc *DbMessageFieldsCreate) OnConflictColumns(columns ...string) *DbMessageFieldsUpsertOne {
	dmfc.conflict = append(dmfc.conflict, sql.ConflictColumns(columns...))
	return &DbMessageFieldsUpsertOne{
		create: dmfc,
	}
}

type (
	// DbMessageFieldsUpsertOne is the builder for "upsert"-ing
	//  one DbMessageFields node.
	DbMessageFieldsUpsertOne struct {
		create *DbMessageFieldsCreate
	}

	// DbMessageFieldsUpsert is the "OnConflict" setter.
	DbMessageFieldsUpsert struct {
		*sql.UpdateSet
	}
)

// SetTenantID sets the "tenant_id" field.
func (u *DbMessageFieldsUpsert) SetTenantID(v int) *DbMessageFieldsUpsert {
	u.Set(dbmessagefields.FieldTenantID, v)
	return u
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbMessageFieldsUpsert) UpdateTenantID() *DbMessageFieldsUpsert {
	u.SetExcluded(dbmessagefields.FieldTenantID)
	return u
}

// SetAppData sets the "AppData" field.
func (u *DbMessageFieldsUpsert) SetAppData(v interfaces.AppData) *DbMessageFieldsUpsert {
	u.Set(dbmessagefields.FieldAppData, v)
	return u
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbMessageFieldsUpsert) UpdateAppData() *DbMessageFieldsUpsert {
	u.SetExcluded(dbmessagefields.FieldAppData)
	return u
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbMessageFieldsUpsert) ClearAppData() *DbMessageFieldsUpsert {
	u.SetNull(dbmessagefields.FieldAppData)
	return u
}

// SetName sets the "Name" field.
func (u *DbMessageFieldsUpsert) SetName(v string) *DbMessageFieldsUpsert {
	u.Set(dbmessagefields.FieldName, v)
	return u
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbMessageFieldsUpsert) UpdateName() *DbMessageFieldsUpsert {
	u.SetExcluded(dbmessagefields.FieldName)
	return u
}

// SetValue sets the "Value" field.
func (u *DbMessageFieldsUpsert) SetValue(v string) *DbMessageFieldsUpsert {
	u.Set(dbmessagefields.FieldValue, v)
	return u
}

// UpdateValue sets the "Value" field to the value that was provided on create.
func (u *DbMessageFieldsUpsert) UpdateValue() *DbMessageFieldsUpsert {
	u.SetExcluded(dbmessagefields.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DbMessageFields.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DbMessageFieldsUpsertOne) UpdateNewValues() *DbMessageFieldsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DbMessageFields.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DbMessageFieldsUpsertOne) Ignore() *DbMessageFieldsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbMessageFieldsUpsertOne) DoNothing() *DbMessageFieldsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbMessageFieldsCreate.OnConflict
// documentation for more info.
func (u *DbMessageFieldsUpsertOne) Update(set func(*DbMessageFieldsUpsert)) *DbMessageFieldsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbMessageFieldsUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbMessageFieldsUpsertOne) SetTenantID(v int) *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertOne) UpdateTenantID() *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbMessageFieldsUpsertOne) SetAppData(v interfaces.AppData) *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertOne) UpdateAppData() *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbMessageFieldsUpsertOne) ClearAppData() *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.ClearAppData()
	})
}

// SetName sets the "Name" field.
func (u *DbMessageFieldsUpsertOne) SetName(v string) *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertOne) UpdateName() *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateName()
	})
}

// SetValue sets the "Value" field.
func (u *DbMessageFieldsUpsertOne) SetValue(v string) *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "Value" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertOne) UpdateValue() *DbMessageFieldsUpsertOne {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *DbMessageFieldsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbMessageFieldsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbMessageFieldsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DbMessageFieldsUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DbMessageFieldsUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (dmfc *DbMessageFieldsCreate) SetDbMessageFieldsFromStruct(input *DbMessageFields) *DbMessageFieldsCreate {

	dmfc.SetTenantID(input.TenantID)

	dmfc.SetAppData(input.AppData)

	dmfc.SetName(input.Name)

	dmfc.SetValue(input.Value)

	return dmfc
}

// DbMessageFieldsCreateBulk is the builder for creating many DbMessageFields entities in bulk.
type DbMessageFieldsCreateBulk struct {
	config
	builders []*DbMessageFieldsCreate
	conflict []sql.ConflictOption
}

// Save creates the DbMessageFields entities in the database.
func (dmfcb *DbMessageFieldsCreateBulk) Save(ctx context.Context) ([]*DbMessageFields, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dmfcb.builders))
	nodes := make([]*DbMessageFields, len(dmfcb.builders))
	mutators := make([]Mutator, len(dmfcb.builders))
	for i := range dmfcb.builders {
		func(i int, root context.Context) {
			builder := dmfcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DbMessageFieldsMutation)
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
					_, err = mutators[i+1].Mutate(root, dmfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dmfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dmfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dmfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dmfcb *DbMessageFieldsCreateBulk) SaveX(ctx context.Context) []*DbMessageFields {
	v, err := dmfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmfcb *DbMessageFieldsCreateBulk) Exec(ctx context.Context) error {
	_, err := dmfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmfcb *DbMessageFieldsCreateBulk) ExecX(ctx context.Context) {
	if err := dmfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DbMessageFields.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbMessageFieldsUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (dmfcb *DbMessageFieldsCreateBulk) OnConflict(opts ...sql.ConflictOption) *DbMessageFieldsUpsertBulk {
	dmfcb.conflict = opts
	return &DbMessageFieldsUpsertBulk{
		create: dmfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbMessageFields.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmfcb *DbMessageFieldsCreateBulk) OnConflictColumns(columns ...string) *DbMessageFieldsUpsertBulk {
	dmfcb.conflict = append(dmfcb.conflict, sql.ConflictColumns(columns...))
	return &DbMessageFieldsUpsertBulk{
		create: dmfcb,
	}
}

// DbMessageFieldsUpsertBulk is the builder for "upsert"-ing
// a bulk of DbMessageFields nodes.
type DbMessageFieldsUpsertBulk struct {
	create *DbMessageFieldsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DbMessageFields.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DbMessageFieldsUpsertBulk) UpdateNewValues() *DbMessageFieldsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DbMessageFields.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DbMessageFieldsUpsertBulk) Ignore() *DbMessageFieldsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbMessageFieldsUpsertBulk) DoNothing() *DbMessageFieldsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbMessageFieldsCreateBulk.OnConflict
// documentation for more info.
func (u *DbMessageFieldsUpsertBulk) Update(set func(*DbMessageFieldsUpsert)) *DbMessageFieldsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbMessageFieldsUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbMessageFieldsUpsertBulk) SetTenantID(v int) *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertBulk) UpdateTenantID() *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbMessageFieldsUpsertBulk) SetAppData(v interfaces.AppData) *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertBulk) UpdateAppData() *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbMessageFieldsUpsertBulk) ClearAppData() *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.ClearAppData()
	})
}

// SetName sets the "Name" field.
func (u *DbMessageFieldsUpsertBulk) SetName(v string) *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertBulk) UpdateName() *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateName()
	})
}

// SetValue sets the "Value" field.
func (u *DbMessageFieldsUpsertBulk) SetValue(v string) *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "Value" field to the value that was provided on create.
func (u *DbMessageFieldsUpsertBulk) UpdateValue() *DbMessageFieldsUpsertBulk {
	return u.Update(func(s *DbMessageFieldsUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *DbMessageFieldsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DbMessageFieldsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbMessageFieldsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbMessageFieldsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
