// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/predicate"
	"github.com/responserms/response/internal/ent/vehiclemodel"
)

// VehicleModelDelete is the builder for deleting a VehicleModel entity.
type VehicleModelDelete struct {
	config
	hooks    []Hook
	mutation *VehicleModelMutation
}

// Where appends a list predicates to the VehicleModelDelete builder.
func (vmd *VehicleModelDelete) Where(ps ...predicate.VehicleModel) *VehicleModelDelete {
	vmd.mutation.Where(ps...)
	return vmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vmd *VehicleModelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vmd.hooks) == 0 {
		affected, err = vmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vmd.mutation = mutation
			affected, err = vmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vmd.hooks) - 1; i >= 0; i-- {
			if vmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmd *VehicleModelDelete) ExecX(ctx context.Context) int {
	n, err := vmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vmd *VehicleModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vehiclemodel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehiclemodel.FieldID,
			},
		},
	}
	if ps := vmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vmd.driver, _spec)
}

// VehicleModelDeleteOne is the builder for deleting a single VehicleModel entity.
type VehicleModelDeleteOne struct {
	vmd *VehicleModelDelete
}

// Exec executes the deletion query.
func (vmdo *VehicleModelDeleteOne) Exec(ctx context.Context) error {
	n, err := vmdo.vmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vehiclemodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vmdo *VehicleModelDeleteOne) ExecX(ctx context.Context) {
	vmdo.vmd.ExecX(ctx)
}
