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
	conflict []sql.ConflictOption
}

// SetTenantID sets the "tenant_id" field.
func (dac *DbAppCreate) SetTenantID(i int) *DbAppCreate {
	dac.mutation.SetTenantID(i)
	return dac
}

// SetAppData sets the "AppData" field.
func (dac *DbAppCreate) SetAppData(id interfaces.AppData) *DbAppCreate {
	dac.mutation.SetAppData(id)
	return dac
}

// SetNillableAppData sets the "AppData" field if the given value is not nil.
func (dac *DbAppCreate) SetNillableAppData(id *interfaces.AppData) *DbAppCreate {
	if id != nil {
		dac.SetAppData(*id)
	}
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
	_spec.OnConflict = dac.conflict
	if value, ok := dac.mutation.AppData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbapp.FieldAppData,
		})
		_node.AppData = value
	}
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DbApp.Create().
//		SetTenantID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbAppUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
//
func (dac *DbAppCreate) OnConflict(opts ...sql.ConflictOption) *DbAppUpsertOne {
	dac.conflict = opts
	return &DbAppUpsertOne{
		create: dac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbApp.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dac *DbAppCreate) OnConflictColumns(columns ...string) *DbAppUpsertOne {
	dac.conflict = append(dac.conflict, sql.ConflictColumns(columns...))
	return &DbAppUpsertOne{
		create: dac,
	}
}

type (
	// DbAppUpsertOne is the builder for "upsert"-ing
	//  one DbApp node.
	DbAppUpsertOne struct {
		create *DbAppCreate
	}

	// DbAppUpsert is the "OnConflict" setter.
	DbAppUpsert struct {
		*sql.UpdateSet
	}
)

// SetTenantID sets the "tenant_id" field.
func (u *DbAppUpsert) SetTenantID(v int) *DbAppUpsert {
	u.Set(dbapp.FieldTenantID, v)
	return u
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbAppUpsert) UpdateTenantID() *DbAppUpsert {
	u.SetExcluded(dbapp.FieldTenantID)
	return u
}

// SetAppData sets the "AppData" field.
func (u *DbAppUpsert) SetAppData(v interfaces.AppData) *DbAppUpsert {
	u.Set(dbapp.FieldAppData, v)
	return u
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbAppUpsert) UpdateAppData() *DbAppUpsert {
	u.SetExcluded(dbapp.FieldAppData)
	return u
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbAppUpsert) ClearAppData() *DbAppUpsert {
	u.SetNull(dbapp.FieldAppData)
	return u
}

// SetName sets the "Name" field.
func (u *DbAppUpsert) SetName(v string) *DbAppUpsert {
	u.Set(dbapp.FieldName, v)
	return u
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbAppUpsert) UpdateName() *DbAppUpsert {
	u.SetExcluded(dbapp.FieldName)
	return u
}

// SetStatus sets the "Status" field.
func (u *DbAppUpsert) SetStatus(v interfaces.AppStatus) *DbAppUpsert {
	u.Set(dbapp.FieldStatus, v)
	return u
}

// UpdateStatus sets the "Status" field to the value that was provided on create.
func (u *DbAppUpsert) UpdateStatus() *DbAppUpsert {
	u.SetExcluded(dbapp.FieldStatus)
	return u
}

// SetDescription sets the "Description" field.
func (u *DbAppUpsert) SetDescription(v string) *DbAppUpsert {
	u.Set(dbapp.FieldDescription, v)
	return u
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbAppUpsert) UpdateDescription() *DbAppUpsert {
	u.SetExcluded(dbapp.FieldDescription)
	return u
}

// SetIcon sets the "icon" field.
func (u *DbAppUpsert) SetIcon(v string) *DbAppUpsert {
	u.Set(dbapp.FieldIcon, v)
	return u
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *DbAppUpsert) UpdateIcon() *DbAppUpsert {
	u.SetExcluded(dbapp.FieldIcon)
	return u
}

// ClearIcon clears the value of the "icon" field.
func (u *DbAppUpsert) ClearIcon() *DbAppUpsert {
	u.SetNull(dbapp.FieldIcon)
	return u
}

// SetURL sets the "url" field.
func (u *DbAppUpsert) SetURL(v string) *DbAppUpsert {
	u.Set(dbapp.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *DbAppUpsert) UpdateURL() *DbAppUpsert {
	u.SetExcluded(dbapp.FieldURL)
	return u
}

// ClearURL clears the value of the "url" field.
func (u *DbAppUpsert) ClearURL() *DbAppUpsert {
	u.SetNull(dbapp.FieldURL)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DbApp.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DbAppUpsertOne) UpdateNewValues() *DbAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.DbApp.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DbAppUpsertOne) Ignore() *DbAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbAppUpsertOne) DoNothing() *DbAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbAppCreate.OnConflict
// documentation for more info.
func (u *DbAppUpsertOne) Update(set func(*DbAppUpsert)) *DbAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbAppUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbAppUpsertOne) SetTenantID(v int) *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbAppUpsertOne) UpdateTenantID() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbAppUpsertOne) SetAppData(v interfaces.AppData) *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbAppUpsertOne) UpdateAppData() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbAppUpsertOne) ClearAppData() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.ClearAppData()
	})
}

// SetName sets the "Name" field.
func (u *DbAppUpsertOne) SetName(v string) *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbAppUpsertOne) UpdateName() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateName()
	})
}

// SetStatus sets the "Status" field.
func (u *DbAppUpsertOne) SetStatus(v interfaces.AppStatus) *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "Status" field to the value that was provided on create.
func (u *DbAppUpsertOne) UpdateStatus() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateStatus()
	})
}

// SetDescription sets the "Description" field.
func (u *DbAppUpsertOne) SetDescription(v string) *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbAppUpsertOne) UpdateDescription() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateDescription()
	})
}

// SetIcon sets the "icon" field.
func (u *DbAppUpsertOne) SetIcon(v string) *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *DbAppUpsertOne) UpdateIcon() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *DbAppUpsertOne) ClearIcon() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.ClearIcon()
	})
}

// SetURL sets the "url" field.
func (u *DbAppUpsertOne) SetURL(v string) *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *DbAppUpsertOne) UpdateURL() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *DbAppUpsertOne) ClearURL() *DbAppUpsertOne {
	return u.Update(func(s *DbAppUpsert) {
		s.ClearURL()
	})
}

// Exec executes the query.
func (u *DbAppUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbAppCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbAppUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DbAppUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DbAppUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (dac *DbAppCreate) SetDbAppFromStruct(input *DbApp) *DbAppCreate {

	dac.SetTenantID(input.TenantID)

	dac.SetAppData(input.AppData)

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
	conflict []sql.ConflictOption
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
					spec.OnConflict = dacb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DbApp.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DbAppUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
//
func (dacb *DbAppCreateBulk) OnConflict(opts ...sql.ConflictOption) *DbAppUpsertBulk {
	dacb.conflict = opts
	return &DbAppUpsertBulk{
		create: dacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DbApp.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dacb *DbAppCreateBulk) OnConflictColumns(columns ...string) *DbAppUpsertBulk {
	dacb.conflict = append(dacb.conflict, sql.ConflictColumns(columns...))
	return &DbAppUpsertBulk{
		create: dacb,
	}
}

// DbAppUpsertBulk is the builder for "upsert"-ing
// a bulk of DbApp nodes.
type DbAppUpsertBulk struct {
	create *DbAppCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DbApp.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DbAppUpsertBulk) UpdateNewValues() *DbAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DbApp.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DbAppUpsertBulk) Ignore() *DbAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DbAppUpsertBulk) DoNothing() *DbAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DbAppCreateBulk.OnConflict
// documentation for more info.
func (u *DbAppUpsertBulk) Update(set func(*DbAppUpsert)) *DbAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DbAppUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *DbAppUpsertBulk) SetTenantID(v int) *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *DbAppUpsertBulk) UpdateTenantID() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateTenantID()
	})
}

// SetAppData sets the "AppData" field.
func (u *DbAppUpsertBulk) SetAppData(v interfaces.AppData) *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.SetAppData(v)
	})
}

// UpdateAppData sets the "AppData" field to the value that was provided on create.
func (u *DbAppUpsertBulk) UpdateAppData() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateAppData()
	})
}

// ClearAppData clears the value of the "AppData" field.
func (u *DbAppUpsertBulk) ClearAppData() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.ClearAppData()
	})
}

// SetName sets the "Name" field.
func (u *DbAppUpsertBulk) SetName(v string) *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "Name" field to the value that was provided on create.
func (u *DbAppUpsertBulk) UpdateName() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateName()
	})
}

// SetStatus sets the "Status" field.
func (u *DbAppUpsertBulk) SetStatus(v interfaces.AppStatus) *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "Status" field to the value that was provided on create.
func (u *DbAppUpsertBulk) UpdateStatus() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateStatus()
	})
}

// SetDescription sets the "Description" field.
func (u *DbAppUpsertBulk) SetDescription(v string) *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "Description" field to the value that was provided on create.
func (u *DbAppUpsertBulk) UpdateDescription() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateDescription()
	})
}

// SetIcon sets the "icon" field.
func (u *DbAppUpsertBulk) SetIcon(v string) *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *DbAppUpsertBulk) UpdateIcon() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *DbAppUpsertBulk) ClearIcon() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.ClearIcon()
	})
}

// SetURL sets the "url" field.
func (u *DbAppUpsertBulk) SetURL(v string) *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *DbAppUpsertBulk) UpdateURL() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *DbAppUpsertBulk) ClearURL() *DbAppUpsertBulk {
	return u.Update(func(s *DbAppUpsert) {
		s.ClearURL()
	})
}

// Exec executes the query.
func (u *DbAppUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DbAppCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DbAppCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DbAppUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
