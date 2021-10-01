// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/predicate"
	"github.com/responserms/response/internal/ent/vehicleregistration"
)

// VehicleRegistrationDelete is the builder for deleting a VehicleRegistration entity.
type VehicleRegistrationDelete struct {
	config
	hooks    []Hook
	mutation *VehicleRegistrationMutation
}

// Where appends a list predicates to the VehicleRegistrationDelete builder.
func (vrd *VehicleRegistrationDelete) Where(ps ...predicate.VehicleRegistration) *VehicleRegistrationDelete {
	vrd.mutation.Where(ps...)
	return vrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vrd *VehicleRegistrationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vrd.hooks) == 0 {
		affected, err = vrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleRegistrationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vrd.mutation = mutation
			affected, err = vrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vrd.hooks) - 1; i >= 0; i-- {
			if vrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vrd *VehicleRegistrationDelete) ExecX(ctx context.Context) int {
	n, err := vrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vrd *VehicleRegistrationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vehicleregistration.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehicleregistration.FieldID,
			},
		},
	}
	if ps := vrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vrd.driver, _spec)
}

// VehicleRegistrationDeleteOne is the builder for deleting a single VehicleRegistration entity.
type VehicleRegistrationDeleteOne struct {
	vrd *VehicleRegistrationDelete
}

// Exec executes the deletion query.
func (vrdo *VehicleRegistrationDeleteOne) Exec(ctx context.Context) error {
	n, err := vrdo.vrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vehicleregistration.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vrdo *VehicleRegistrationDeleteOne) ExecX(ctx context.Context) {
	vrdo.vrd.ExecX(ctx)
}
