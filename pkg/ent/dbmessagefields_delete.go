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
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessagefields"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
)

// DbMessageFieldsDelete is the builder for deleting a DbMessageFields entity.
type DbMessageFieldsDelete struct {
	config
	hooks    []Hook
	mutation *DbMessageFieldsMutation
}

// Where appends a list predicates to the DbMessageFieldsDelete builder.
func (dmfd *DbMessageFieldsDelete) Where(ps ...predicate.DbMessageFields) *DbMessageFieldsDelete {
	dmfd.mutation.Where(ps...)
	return dmfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dmfd *DbMessageFieldsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dmfd.hooks) == 0 {
		affected, err = dmfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbMessageFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dmfd.mutation = mutation
			affected, err = dmfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dmfd.hooks) - 1; i >= 0; i-- {
			if dmfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dmfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dmfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmfd *DbMessageFieldsDelete) ExecX(ctx context.Context) int {
	n, err := dmfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dmfd *DbMessageFieldsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dbmessagefields.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbmessagefields.FieldID,
			},
		},
	}
	if ps := dmfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dmfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DbMessageFieldsDeleteOne is the builder for deleting a single DbMessageFields entity.
type DbMessageFieldsDeleteOne struct {
	dmfd *DbMessageFieldsDelete
}

// Exec executes the deletion query.
func (dmfdo *DbMessageFieldsDeleteOne) Exec(ctx context.Context) error {
	n, err := dmfdo.dmfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dbmessagefields.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dmfdo *DbMessageFieldsDeleteOne) ExecX(ctx context.Context) {
	dmfdo.dmfd.ExecX(ctx)
}