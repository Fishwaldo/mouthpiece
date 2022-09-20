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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbgroup"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportinstances"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
)

// DbTransportRecipientsCreate is the builder for creating a DbTransportRecipients entity.
type DbTransportRecipientsCreate struct {
	config
	mutation *DbTransportRecipientsMutation
	hooks    []Hook
}

// SetTenantID sets the "tenant_id" field.
func (dtrc *DbTransportRecipientsCreate) SetTenantID(i int) *DbTransportRecipientsCreate {
	dtrc.mutation.SetTenantID(i)
	return dtrc
}

// SetName sets the "Name" field.
func (dtrc *DbTransportRecipientsCreate) SetName(s string) *DbTransportRecipientsCreate {
	dtrc.mutation.SetName(s)
	return dtrc
}

// SetDescription sets the "Description" field.
func (dtrc *DbTransportRecipientsCreate) SetDescription(s string) *DbTransportRecipientsCreate {
	dtrc.mutation.SetDescription(s)
	return dtrc
}

// SetNillableDescription sets the "Description" field if the given value is not nil.
func (dtrc *DbTransportRecipientsCreate) SetNillableDescription(s *string) *DbTransportRecipientsCreate {
	if s != nil {
		dtrc.SetDescription(*s)
	}
	return dtrc
}

// SetConfig sets the "config" field.
func (dtrc *DbTransportRecipientsCreate) SetConfig(s string) *DbTransportRecipientsCreate {
	dtrc.mutation.SetConfig(s)
	return dtrc
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (dtrc *DbTransportRecipientsCreate) SetTenant(t *Tenant) *DbTransportRecipientsCreate {
	return dtrc.SetTenantID(t.ID)
}

// SetTransportInstanceID sets the "TransportInstance" edge to the DbTransportInstances entity by ID.
func (dtrc *DbTransportRecipientsCreate) SetTransportInstanceID(id int) *DbTransportRecipientsCreate {
	dtrc.mutation.SetTransportInstanceID(id)
	return dtrc
}

// SetTransportInstance sets the "TransportInstance" edge to the DbTransportInstances entity.
func (dtrc *DbTransportRecipientsCreate) SetTransportInstance(d *DbTransportInstances) *DbTransportRecipientsCreate {
	return dtrc.SetTransportInstanceID(d.ID)
}

// SetGroupRecipientID sets the "GroupRecipient" edge to the DbGroup entity by ID.
func (dtrc *DbTransportRecipientsCreate) SetGroupRecipientID(id int) *DbTransportRecipientsCreate {
	dtrc.mutation.SetGroupRecipientID(id)
	return dtrc
}

// SetNillableGroupRecipientID sets the "GroupRecipient" edge to the DbGroup entity by ID if the given value is not nil.
func (dtrc *DbTransportRecipientsCreate) SetNillableGroupRecipientID(id *int) *DbTransportRecipientsCreate {
	if id != nil {
		dtrc = dtrc.SetGroupRecipientID(*id)
	}
	return dtrc
}

// SetGroupRecipient sets the "GroupRecipient" edge to the DbGroup entity.
func (dtrc *DbTransportRecipientsCreate) SetGroupRecipient(d *DbGroup) *DbTransportRecipientsCreate {
	return dtrc.SetGroupRecipientID(d.ID)
}

// SetUserRecipientID sets the "UserRecipient" edge to the DbUser entity by ID.
func (dtrc *DbTransportRecipientsCreate) SetUserRecipientID(id int) *DbTransportRecipientsCreate {
	dtrc.mutation.SetUserRecipientID(id)
	return dtrc
}

// SetNillableUserRecipientID sets the "UserRecipient" edge to the DbUser entity by ID if the given value is not nil.
func (dtrc *DbTransportRecipientsCreate) SetNillableUserRecipientID(id *int) *DbTransportRecipientsCreate {
	if id != nil {
		dtrc = dtrc.SetUserRecipientID(*id)
	}
	return dtrc
}

// SetUserRecipient sets the "UserRecipient" edge to the DbUser entity.
func (dtrc *DbTransportRecipientsCreate) SetUserRecipient(d *DbUser) *DbTransportRecipientsCreate {
	return dtrc.SetUserRecipientID(d.ID)
}

// Mutation returns the DbTransportRecipientsMutation object of the builder.
func (dtrc *DbTransportRecipientsCreate) Mutation() *DbTransportRecipientsMutation {
	return dtrc.mutation
}

// Save creates the DbTransportRecipients in the database.
func (dtrc *DbTransportRecipientsCreate) Save(ctx context.Context) (*DbTransportRecipients, error) {
	var (
		err  error
		node *DbTransportRecipients
	)
	if len(dtrc.hooks) == 0 {
		if err = dtrc.check(); err != nil {
			return nil, err
		}
		node, err = dtrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbTransportRecipientsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dtrc.check(); err != nil {
				return nil, err
			}
			dtrc.mutation = mutation
			if node, err = dtrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dtrc.hooks) - 1; i >= 0; i-- {
			if dtrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dtrc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dtrc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DbTransportRecipients)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DbTransportRecipientsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dtrc *DbTransportRecipientsCreate) SaveX(ctx context.Context) *DbTransportRecipients {
	v, err := dtrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtrc *DbTransportRecipientsCreate) Exec(ctx context.Context) error {
	_, err := dtrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtrc *DbTransportRecipientsCreate) ExecX(ctx context.Context) {
	if err := dtrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtrc *DbTransportRecipientsCreate) check() error {
	if _, ok := dtrc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "DbTransportRecipients.tenant_id"`)}
	}
	if _, ok := dtrc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "DbTransportRecipients.Name"`)}
	}
	if v, ok := dtrc.mutation.Name(); ok {
		if err := dbtransportrecipients.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "DbTransportRecipients.Name": %w`, err)}
		}
	}
	if _, ok := dtrc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required field "DbTransportRecipients.config"`)}
	}
	if _, ok := dtrc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`ent: missing required edge "DbTransportRecipients.tenant"`)}
	}
	if _, ok := dtrc.mutation.TransportInstanceID(); !ok {
		return &ValidationError{Name: "TransportInstance", err: errors.New(`ent: missing required edge "DbTransportRecipients.TransportInstance"`)}
	}
	return nil
}

func (dtrc *DbTransportRecipientsCreate) sqlSave(ctx context.Context) (*DbTransportRecipients, error) {
	_node, _spec := dtrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dtrc *DbTransportRecipientsCreate) createSpec() (*DbTransportRecipients, *sqlgraph.CreateSpec) {
	var (
		_node = &DbTransportRecipients{config: dtrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dbtransportrecipients.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbtransportrecipients.FieldID,
			},
		}
	)
	if value, ok := dtrc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbtransportrecipients.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dtrc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbtransportrecipients.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := dtrc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbtransportrecipients.FieldConfig,
		})
		_node.Config = value
	}
	if nodes := dtrc.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dbtransportrecipients.TenantTable,
			Columns: []string{dbtransportrecipients.TenantColumn},
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
	if nodes := dtrc.mutation.TransportInstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbtransportrecipients.TransportInstanceTable,
			Columns: []string{dbtransportrecipients.TransportInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dbtransportinstances.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.db_transport_instances_transport_recipients = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dtrc.mutation.GroupRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbtransportrecipients.GroupRecipientTable,
			Columns: []string{dbtransportrecipients.GroupRecipientColumn},
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
		_node.db_group_transport_recipients = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dtrc.mutation.UserRecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dbtransportrecipients.UserRecipientTable,
			Columns: []string{dbtransportrecipients.UserRecipientColumn},
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
		_node.db_user_transport_recipients = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

func (dtrc *DbTransportRecipientsCreate) SetDbTransportRecipientsFromStruct(input *DbTransportRecipients) *DbTransportRecipientsCreate {

	dtrc.SetTenantID(input.TenantID)

	dtrc.SetName(input.Name)

	dtrc.SetDescription(input.Description)

	dtrc.SetConfig(input.Config)

	return dtrc
}

// DbTransportRecipientsCreateBulk is the builder for creating many DbTransportRecipients entities in bulk.
type DbTransportRecipientsCreateBulk struct {
	config
	builders []*DbTransportRecipientsCreate
}

// Save creates the DbTransportRecipients entities in the database.
func (dtrcb *DbTransportRecipientsCreateBulk) Save(ctx context.Context) ([]*DbTransportRecipients, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dtrcb.builders))
	nodes := make([]*DbTransportRecipients, len(dtrcb.builders))
	mutators := make([]Mutator, len(dtrcb.builders))
	for i := range dtrcb.builders {
		func(i int, root context.Context) {
			builder := dtrcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DbTransportRecipientsMutation)
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
					_, err = mutators[i+1].Mutate(root, dtrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dtrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dtrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dtrcb *DbTransportRecipientsCreateBulk) SaveX(ctx context.Context) []*DbTransportRecipients {
	v, err := dtrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtrcb *DbTransportRecipientsCreateBulk) Exec(ctx context.Context) error {
	_, err := dtrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtrcb *DbTransportRecipientsCreateBulk) ExecX(ctx context.Context) {
	if err := dtrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
