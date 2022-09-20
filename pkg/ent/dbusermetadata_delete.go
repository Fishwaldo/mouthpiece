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
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbusermetadata"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
)

// DbUserMetaDataDelete is the builder for deleting a DbUserMetaData entity.
type DbUserMetaDataDelete struct {
	config
	hooks    []Hook
	mutation *DbUserMetaDataMutation
}

// Where appends a list predicates to the DbUserMetaDataDelete builder.
func (dumdd *DbUserMetaDataDelete) Where(ps ...predicate.DbUserMetaData) *DbUserMetaDataDelete {
	dumdd.mutation.Where(ps...)
	return dumdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dumdd *DbUserMetaDataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dumdd.hooks) == 0 {
		affected, err = dumdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbUserMetaDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dumdd.mutation = mutation
			affected, err = dumdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dumdd.hooks) - 1; i >= 0; i-- {
			if dumdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dumdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dumdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dumdd *DbUserMetaDataDelete) ExecX(ctx context.Context) int {
	n, err := dumdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dumdd *DbUserMetaDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dbusermetadata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbusermetadata.FieldID,
			},
		},
	}
	if ps := dumdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dumdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DbUserMetaDataDeleteOne is the builder for deleting a single DbUserMetaData entity.
type DbUserMetaDataDeleteOne struct {
	dumdd *DbUserMetaDataDelete
}

// Exec executes the deletion query.
func (dumddo *DbUserMetaDataDeleteOne) Exec(ctx context.Context) error {
	n, err := dumddo.dumdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dbusermetadata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dumddo *DbUserMetaDataDeleteOne) ExecX(ctx context.Context) {
	dumddo.dumdd.ExecX(ctx)
}
