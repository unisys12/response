// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/predicate"
	"github.com/responserms/response/internal/ent/vehiclecolor"
)

// VehicleColorDelete is the builder for deleting a VehicleColor entity.
type VehicleColorDelete struct {
	config
	hooks    []Hook
	mutation *VehicleColorMutation
}

// Where appends a list predicates to the VehicleColorDelete builder.
func (vcd *VehicleColorDelete) Where(ps ...predicate.VehicleColor) *VehicleColorDelete {
	vcd.mutation.Where(ps...)
	return vcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vcd *VehicleColorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vcd.hooks) == 0 {
		affected, err = vcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleColorMutation)
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
func (vcd *VehicleColorDelete) ExecX(ctx context.Context) int {
	n, err := vcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vcd *VehicleColorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vehiclecolor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehiclecolor.FieldID,
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

// VehicleColorDeleteOne is the builder for deleting a single VehicleColor entity.
type VehicleColorDeleteOne struct {
	vcd *VehicleColorDelete
}

// Exec executes the deletion query.
func (vcdo *VehicleColorDeleteOne) Exec(ctx context.Context) error {
	n, err := vcdo.vcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vehiclecolor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vcdo *VehicleColorDeleteOne) ExecX(ctx context.Context) {
	vcdo.vcd.ExecX(ctx)
}
