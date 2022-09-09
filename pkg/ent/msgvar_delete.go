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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/msgvar"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
)

// MsgVarDelete is the builder for deleting a MsgVar entity.
type MsgVarDelete struct {
	config
	hooks    []Hook
	mutation *MsgVarMutation
}

// Where appends a list predicates to the MsgVarDelete builder.
func (mvd *MsgVarDelete) Where(ps ...predicate.MsgVar) *MsgVarDelete {
	mvd.mutation.Where(ps...)
	return mvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mvd *MsgVarDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mvd.hooks) == 0 {
		affected, err = mvd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MsgVarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mvd.mutation = mutation
			affected, err = mvd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mvd.hooks) - 1; i >= 0; i-- {
			if mvd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mvd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mvd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mvd *MsgVarDelete) ExecX(ctx context.Context) int {
	n, err := mvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mvd *MsgVarDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: msgvar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: msgvar.FieldID,
			},
		},
	}
	if ps := mvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// MsgVarDeleteOne is the builder for deleting a single MsgVar entity.
type MsgVarDeleteOne struct {
	mvd *MsgVarDelete
}

// Exec executes the deletion query.
func (mvdo *MsgVarDeleteOne) Exec(ctx context.Context) error {
	n, err := mvdo.mvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{msgvar.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mvdo *MsgVarDeleteOne) ExecX(ctx context.Context) {
	mvdo.mvd.ExecX(ctx)
}
