// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/deathcertifier"
	"github.com/responserms/response/internal/ent/predicate"
)

// DeathCertifierDelete is the builder for deleting a DeathCertifier entity.
type DeathCertifierDelete struct {
	config
	hooks    []Hook
	mutation *DeathCertifierMutation
}

// Where appends a list predicates to the DeathCertifierDelete builder.
func (dcd *DeathCertifierDelete) Where(ps ...predicate.DeathCertifier) *DeathCertifierDelete {
	dcd.mutation.Where(ps...)
	return dcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dcd *DeathCertifierDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dcd.hooks) == 0 {
		affected, err = dcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeathCertifierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dcd.mutation = mutation
			affected, err = dcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dcd.hooks) - 1; i >= 0; i-- {
			if dcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcd *DeathCertifierDelete) ExecX(ctx context.Context) int {
	n, err := dcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dcd *DeathCertifierDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: deathcertifier.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deathcertifier.FieldID,
			},
		},
	}
	if ps := dcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dcd.driver, _spec)
}

// DeathCertifierDeleteOne is the builder for deleting a single DeathCertifier entity.
type DeathCertifierDeleteOne struct {
	dcd *DeathCertifierDelete
}

// Exec executes the deletion query.
func (dcdo *DeathCertifierDeleteOne) Exec(ctx context.Context) error {
	n, err := dcdo.dcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deathcertifier.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dcdo *DeathCertifierDeleteOne) ExecX(ctx context.Context) {
	dcdo.dcd.ExecX(ctx)
}
