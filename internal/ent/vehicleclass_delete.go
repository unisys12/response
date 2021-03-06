// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/predicate"
	"github.com/responserms/response/internal/ent/vehicleclass"
)

// VehicleClassDelete is the builder for deleting a VehicleClass entity.
type VehicleClassDelete struct {
	config
	hooks    []Hook
	mutation *VehicleClassMutation
}

// Where appends a list predicates to the VehicleClassDelete builder.
func (vcd *VehicleClassDelete) Where(ps ...predicate.VehicleClass) *VehicleClassDelete {
	vcd.mutation.Where(ps...)
	return vcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vcd *VehicleClassDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vcd.hooks) == 0 {
		affected, err = vcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vcd.mutation = mutation
			affected, err = vcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vcd.hooks) - 1; i >= 0; i-- {
			if vcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcd *VehicleClassDelete) ExecX(ctx context.Context) int {
	n, err := vcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vcd *VehicleClassDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vehicleclass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehicleclass.FieldID,
			},
		},
	}
	if ps := vcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vcd.driver, _spec)
}

// VehicleClassDeleteOne is the builder for deleting a single VehicleClass entity.
type VehicleClassDeleteOne struct {
	vcd *VehicleClassDelete
}

// Exec executes the deletion query.
func (vcdo *VehicleClassDeleteOne) Exec(ctx context.Context) error {
	n, err := vcdo.vcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vehicleclass.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vcdo *VehicleClassDeleteOne) ExecX(ctx context.Context) {
	vcdo.vcd.ExecX(ctx)
}
