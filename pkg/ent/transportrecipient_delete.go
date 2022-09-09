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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/transportrecipient"
)

// TransportRecipientDelete is the builder for deleting a TransportRecipient entity.
type TransportRecipientDelete struct {
	config
	hooks    []Hook
	mutation *TransportRecipientMutation
}

// Where appends a list predicates to the TransportRecipientDelete builder.
func (trd *TransportRecipientDelete) Where(ps ...predicate.TransportRecipient) *TransportRecipientDelete {
	trd.mutation.Where(ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TransportRecipientDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(trd.hooks) == 0 {
		affected, err = trd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransportRecipientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			trd.mutation = mutation
			affected, err = trd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(trd.hooks) - 1; i >= 0; i-- {
			if trd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = trd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, trd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TransportRecipientDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TransportRecipientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: transportrecipient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transportrecipient.FieldID,
			},
		},
	}
	if ps := trd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TransportRecipientDeleteOne is the builder for deleting a single TransportRecipient entity.
type TransportRecipientDeleteOne struct {
	trd *TransportRecipientDelete
}

// Exec executes the deletion query.
func (trdo *TransportRecipientDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{transportrecipient.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TransportRecipientDeleteOne) ExecX(ctx context.Context) {
	trdo.trd.ExecX(ctx)
}