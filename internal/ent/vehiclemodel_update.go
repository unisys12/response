// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/predicate"
	"github.com/responserms/response/internal/ent/vehicle"
	"github.com/responserms/response/internal/ent/vehiclemodel"
)

// VehicleModelUpdate is the builder for updating VehicleModel entities.
type VehicleModelUpdate struct {
	config
	hooks    []Hook
	mutation *VehicleModelMutation
}

// Where appends a list predicates to the VehicleModelUpdate builder.
func (vmu *VehicleModelUpdate) Where(ps ...predicate.VehicleModel) *VehicleModelUpdate {
	vmu.mutation.Where(ps...)
	return vmu
}

// SetCreatedBy sets the "created_by" field.
func (vmu *VehicleModelUpdate) SetCreatedBy(i int) *VehicleModelUpdate {
	vmu.mutation.ResetCreatedBy()
	vmu.mutation.SetCreatedBy(i)
	return vmu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vmu *VehicleModelUpdate) SetNillableCreatedBy(i *int) *VehicleModelUpdate {
	if i != nil {
		vmu.SetCreatedBy(*i)
	}
	return vmu
}

// AddCreatedBy adds i to the "created_by" field.
func (vmu *VehicleModelUpdate) AddCreatedBy(i int) *VehicleModelUpdate {
	vmu.mutation.AddCreatedBy(i)
	return vmu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (vmu *VehicleModelUpdate) ClearCreatedBy() *VehicleModelUpdate {
	vmu.mutation.ClearCreatedBy()
	return vmu
}

// SetCreatedWith sets the "created_with" field.
func (vmu *VehicleModelUpdate) SetCreatedWith(s string) *VehicleModelUpdate {
	vmu.mutation.SetCreatedWith(s)
	return vmu
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (vmu *VehicleModelUpdate) SetNillableCreatedWith(s *string) *VehicleModelUpdate {
	if s != nil {
		vmu.SetCreatedWith(*s)
	}
	return vmu
}

// ClearCreatedWith clears the value of the "created_with" field.
func (vmu *VehicleModelUpdate) ClearCreatedWith() *VehicleModelUpdate {
	vmu.mutation.ClearCreatedWith()
	return vmu
}

// SetUpdatedAt sets the "updated_at" field.
func (vmu *VehicleModelUpdate) SetUpdatedAt(t time.Time) *VehicleModelUpdate {
	vmu.mutation.SetUpdatedAt(t)
	return vmu
}

// SetUpdatedBy sets the "updated_by" field.
func (vmu *VehicleModelUpdate) SetUpdatedBy(i int) *VehicleModelUpdate {
	vmu.mutation.ResetUpdatedBy()
	vmu.mutation.SetUpdatedBy(i)
	return vmu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vmu *VehicleModelUpdate) SetNillableUpdatedBy(i *int) *VehicleModelUpdate {
	if i != nil {
		vmu.SetUpdatedBy(*i)
	}
	return vmu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (vmu *VehicleModelUpdate) AddUpdatedBy(i int) *VehicleModelUpdate {
	vmu.mutation.AddUpdatedBy(i)
	return vmu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (vmu *VehicleModelUpdate) ClearUpdatedBy() *VehicleModelUpdate {
	vmu.mutation.ClearUpdatedBy()
	return vmu
}

// SetUpdatedWith sets the "updated_with" field.
func (vmu *VehicleModelUpdate) SetUpdatedWith(s string) *VehicleModelUpdate {
	vmu.mutation.SetUpdatedWith(s)
	return vmu
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (vmu *VehicleModelUpdate) SetNillableUpdatedWith(s *string) *VehicleModelUpdate {
	if s != nil {
		vmu.SetUpdatedWith(*s)
	}
	return vmu
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (vmu *VehicleModelUpdate) ClearUpdatedWith() *VehicleModelUpdate {
	vmu.mutation.ClearUpdatedWith()
	return vmu
}

// SetShort sets the "short" field.
func (vmu *VehicleModelUpdate) SetShort(s string) *VehicleModelUpdate {
	vmu.mutation.SetShort(s)
	return vmu
}

// SetTitle sets the "title" field.
func (vmu *VehicleModelUpdate) SetTitle(s string) *VehicleModelUpdate {
	vmu.mutation.SetTitle(s)
	return vmu
}

// SetDescription sets the "description" field.
func (vmu *VehicleModelUpdate) SetDescription(s string) *VehicleModelUpdate {
	vmu.mutation.SetDescription(s)
	return vmu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vmu *VehicleModelUpdate) SetNillableDescription(s *string) *VehicleModelUpdate {
	if s != nil {
		vmu.SetDescription(*s)
	}
	return vmu
}

// ClearDescription clears the value of the "description" field.
func (vmu *VehicleModelUpdate) ClearDescription() *VehicleModelUpdate {
	vmu.mutation.ClearDescription()
	return vmu
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (vmu *VehicleModelUpdate) SetMetadataID(id int) *VehicleModelUpdate {
	vmu.mutation.SetMetadataID(id)
	return vmu
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (vmu *VehicleModelUpdate) SetNillableMetadataID(id *int) *VehicleModelUpdate {
	if id != nil {
		vmu = vmu.SetMetadataID(*id)
	}
	return vmu
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (vmu *VehicleModelUpdate) SetMetadata(m *Metadata) *VehicleModelUpdate {
	return vmu.SetMetadataID(m.ID)
}

// AddVehicleIDs adds the "vehicles" edge to the Vehicle entity by IDs.
func (vmu *VehicleModelUpdate) AddVehicleIDs(ids ...int) *VehicleModelUpdate {
	vmu.mutation.AddVehicleIDs(ids...)
	return vmu
}

// AddVehicles adds the "vehicles" edges to the Vehicle entity.
func (vmu *VehicleModelUpdate) AddVehicles(v ...*Vehicle) *VehicleModelUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vmu.AddVehicleIDs(ids...)
}

// Mutation returns the VehicleModelMutation object of the builder.
func (vmu *VehicleModelUpdate) Mutation() *VehicleModelMutation {
	return vmu.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (vmu *VehicleModelUpdate) ClearMetadata() *VehicleModelUpdate {
	vmu.mutation.ClearMetadata()
	return vmu
}

// ClearVehicles clears all "vehicles" edges to the Vehicle entity.
func (vmu *VehicleModelUpdate) ClearVehicles() *VehicleModelUpdate {
	vmu.mutation.ClearVehicles()
	return vmu
}

// RemoveVehicleIDs removes the "vehicles" edge to Vehicle entities by IDs.
func (vmu *VehicleModelUpdate) RemoveVehicleIDs(ids ...int) *VehicleModelUpdate {
	vmu.mutation.RemoveVehicleIDs(ids...)
	return vmu
}

// RemoveVehicles removes "vehicles" edges to Vehicle entities.
func (vmu *VehicleModelUpdate) RemoveVehicles(v ...*Vehicle) *VehicleModelUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vmu.RemoveVehicleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vmu *VehicleModelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := vmu.defaults(); err != nil {
		return 0, err
	}
	if len(vmu.hooks) == 0 {
		if err = vmu.check(); err != nil {
			return 0, err
		}
		affected, err = vmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vmu.check(); err != nil {
				return 0, err
			}
			vmu.mutation = mutation
			affected, err = vmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vmu.hooks) - 1; i >= 0; i-- {
			if vmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vmu *VehicleModelUpdate) SaveX(ctx context.Context) int {
	affected, err := vmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vmu *VehicleModelUpdate) Exec(ctx context.Context) error {
	_, err := vmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmu *VehicleModelUpdate) ExecX(ctx context.Context) {
	if err := vmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vmu *VehicleModelUpdate) defaults() error {
	if _, ok := vmu.mutation.UpdatedAt(); !ok {
		if vehiclemodel.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized vehiclemodel.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := vehiclemodel.UpdateDefaultUpdatedAt()
		vmu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vmu *VehicleModelUpdate) check() error {
	if v, ok := vmu.mutation.Short(); ok {
		if err := vehiclemodel.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf("ent: validator failed for field \"short\": %w", err)}
		}
	}
	if v, ok := vmu.mutation.Title(); ok {
		if err := vehiclemodel.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := vmu.mutation.Description(); ok {
		if err := vehiclemodel.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (vmu *VehicleModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vehiclemodel.Table,
			Columns: vehiclemodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehiclemodel.FieldID,
			},
		},
	}
	if ps := vmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vmu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldCreatedBy,
		})
	}
	if value, ok := vmu.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldCreatedBy,
		})
	}
	if vmu.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: vehiclemodel.FieldCreatedBy,
		})
	}
	if value, ok := vmu.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldCreatedWith,
		})
	}
	if vmu.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehiclemodel.FieldCreatedWith,
		})
	}
	if value, ok := vmu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedAt,
		})
	}
	if value, ok := vmu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedBy,
		})
	}
	if value, ok := vmu.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedBy,
		})
	}
	if vmu.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: vehiclemodel.FieldUpdatedBy,
		})
	}
	if value, ok := vmu.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedWith,
		})
	}
	if vmu.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehiclemodel.FieldUpdatedWith,
		})
	}
	if value, ok := vmu.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldShort,
		})
	}
	if value, ok := vmu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldTitle,
		})
	}
	if value, ok := vmu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldDescription,
		})
	}
	if vmu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehiclemodel.FieldDescription,
		})
	}
	if vmu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vehiclemodel.MetadataTable,
			Columns: []string{vehiclemodel.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vmu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vehiclemodel.MetadataTable,
			Columns: []string{vehiclemodel.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vmu.mutation.VehiclesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vehiclemodel.VehiclesTable,
			Columns: []string{vehiclemodel.VehiclesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vehicle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vmu.mutation.RemovedVehiclesIDs(); len(nodes) > 0 && !vmu.mutation.VehiclesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vehiclemodel.VehiclesTable,
			Columns: []string{vehiclemodel.VehiclesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vehicle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vmu.mutation.VehiclesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vehiclemodel.VehiclesTable,
			Columns: []string{vehiclemodel.VehiclesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vehicle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehiclemodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VehicleModelUpdateOne is the builder for updating a single VehicleModel entity.
type VehicleModelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VehicleModelMutation
}

// SetCreatedBy sets the "created_by" field.
func (vmuo *VehicleModelUpdateOne) SetCreatedBy(i int) *VehicleModelUpdateOne {
	vmuo.mutation.ResetCreatedBy()
	vmuo.mutation.SetCreatedBy(i)
	return vmuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vmuo *VehicleModelUpdateOne) SetNillableCreatedBy(i *int) *VehicleModelUpdateOne {
	if i != nil {
		vmuo.SetCreatedBy(*i)
	}
	return vmuo
}

// AddCreatedBy adds i to the "created_by" field.
func (vmuo *VehicleModelUpdateOne) AddCreatedBy(i int) *VehicleModelUpdateOne {
	vmuo.mutation.AddCreatedBy(i)
	return vmuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (vmuo *VehicleModelUpdateOne) ClearCreatedBy() *VehicleModelUpdateOne {
	vmuo.mutation.ClearCreatedBy()
	return vmuo
}

// SetCreatedWith sets the "created_with" field.
func (vmuo *VehicleModelUpdateOne) SetCreatedWith(s string) *VehicleModelUpdateOne {
	vmuo.mutation.SetCreatedWith(s)
	return vmuo
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (vmuo *VehicleModelUpdateOne) SetNillableCreatedWith(s *string) *VehicleModelUpdateOne {
	if s != nil {
		vmuo.SetCreatedWith(*s)
	}
	return vmuo
}

// ClearCreatedWith clears the value of the "created_with" field.
func (vmuo *VehicleModelUpdateOne) ClearCreatedWith() *VehicleModelUpdateOne {
	vmuo.mutation.ClearCreatedWith()
	return vmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vmuo *VehicleModelUpdateOne) SetUpdatedAt(t time.Time) *VehicleModelUpdateOne {
	vmuo.mutation.SetUpdatedAt(t)
	return vmuo
}

// SetUpdatedBy sets the "updated_by" field.
func (vmuo *VehicleModelUpdateOne) SetUpdatedBy(i int) *VehicleModelUpdateOne {
	vmuo.mutation.ResetUpdatedBy()
	vmuo.mutation.SetUpdatedBy(i)
	return vmuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vmuo *VehicleModelUpdateOne) SetNillableUpdatedBy(i *int) *VehicleModelUpdateOne {
	if i != nil {
		vmuo.SetUpdatedBy(*i)
	}
	return vmuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (vmuo *VehicleModelUpdateOne) AddUpdatedBy(i int) *VehicleModelUpdateOne {
	vmuo.mutation.AddUpdatedBy(i)
	return vmuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (vmuo *VehicleModelUpdateOne) ClearUpdatedBy() *VehicleModelUpdateOne {
	vmuo.mutation.ClearUpdatedBy()
	return vmuo
}

// SetUpdatedWith sets the "updated_with" field.
func (vmuo *VehicleModelUpdateOne) SetUpdatedWith(s string) *VehicleModelUpdateOne {
	vmuo.mutation.SetUpdatedWith(s)
	return vmuo
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (vmuo *VehicleModelUpdateOne) SetNillableUpdatedWith(s *string) *VehicleModelUpdateOne {
	if s != nil {
		vmuo.SetUpdatedWith(*s)
	}
	return vmuo
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (vmuo *VehicleModelUpdateOne) ClearUpdatedWith() *VehicleModelUpdateOne {
	vmuo.mutation.ClearUpdatedWith()
	return vmuo
}

// SetShort sets the "short" field.
func (vmuo *VehicleModelUpdateOne) SetShort(s string) *VehicleModelUpdateOne {
	vmuo.mutation.SetShort(s)
	return vmuo
}

// SetTitle sets the "title" field.
func (vmuo *VehicleModelUpdateOne) SetTitle(s string) *VehicleModelUpdateOne {
	vmuo.mutation.SetTitle(s)
	return vmuo
}

// SetDescription sets the "description" field.
func (vmuo *VehicleModelUpdateOne) SetDescription(s string) *VehicleModelUpdateOne {
	vmuo.mutation.SetDescription(s)
	return vmuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vmuo *VehicleModelUpdateOne) SetNillableDescription(s *string) *VehicleModelUpdateOne {
	if s != nil {
		vmuo.SetDescription(*s)
	}
	return vmuo
}

// ClearDescription clears the value of the "description" field.
func (vmuo *VehicleModelUpdateOne) ClearDescription() *VehicleModelUpdateOne {
	vmuo.mutation.ClearDescription()
	return vmuo
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (vmuo *VehicleModelUpdateOne) SetMetadataID(id int) *VehicleModelUpdateOne {
	vmuo.mutation.SetMetadataID(id)
	return vmuo
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (vmuo *VehicleModelUpdateOne) SetNillableMetadataID(id *int) *VehicleModelUpdateOne {
	if id != nil {
		vmuo = vmuo.SetMetadataID(*id)
	}
	return vmuo
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (vmuo *VehicleModelUpdateOne) SetMetadata(m *Metadata) *VehicleModelUpdateOne {
	return vmuo.SetMetadataID(m.ID)
}

// AddVehicleIDs adds the "vehicles" edge to the Vehicle entity by IDs.
func (vmuo *VehicleModelUpdateOne) AddVehicleIDs(ids ...int) *VehicleModelUpdateOne {
	vmuo.mutation.AddVehicleIDs(ids...)
	return vmuo
}

// AddVehicles adds the "vehicles" edges to the Vehicle entity.
func (vmuo *VehicleModelUpdateOne) AddVehicles(v ...*Vehicle) *VehicleModelUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vmuo.AddVehicleIDs(ids...)
}

// Mutation returns the VehicleModelMutation object of the builder.
func (vmuo *VehicleModelUpdateOne) Mutation() *VehicleModelMutation {
	return vmuo.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (vmuo *VehicleModelUpdateOne) ClearMetadata() *VehicleModelUpdateOne {
	vmuo.mutation.ClearMetadata()
	return vmuo
}

// ClearVehicles clears all "vehicles" edges to the Vehicle entity.
func (vmuo *VehicleModelUpdateOne) ClearVehicles() *VehicleModelUpdateOne {
	vmuo.mutation.ClearVehicles()
	return vmuo
}

// RemoveVehicleIDs removes the "vehicles" edge to Vehicle entities by IDs.
func (vmuo *VehicleModelUpdateOne) RemoveVehicleIDs(ids ...int) *VehicleModelUpdateOne {
	vmuo.mutation.RemoveVehicleIDs(ids...)
	return vmuo
}

// RemoveVehicles removes "vehicles" edges to Vehicle entities.
func (vmuo *VehicleModelUpdateOne) RemoveVehicles(v ...*Vehicle) *VehicleModelUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vmuo.RemoveVehicleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vmuo *VehicleModelUpdateOne) Select(field string, fields ...string) *VehicleModelUpdateOne {
	vmuo.fields = append([]string{field}, fields...)
	return vmuo
}

// Save executes the query and returns the updated VehicleModel entity.
func (vmuo *VehicleModelUpdateOne) Save(ctx context.Context) (*VehicleModel, error) {
	var (
		err  error
		node *VehicleModel
	)
	if err := vmuo.defaults(); err != nil {
		return nil, err
	}
	if len(vmuo.hooks) == 0 {
		if err = vmuo.check(); err != nil {
			return nil, err
		}
		node, err = vmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vmuo.check(); err != nil {
				return nil, err
			}
			vmuo.mutation = mutation
			node, err = vmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vmuo.hooks) - 1; i >= 0; i-- {
			if vmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vmuo *VehicleModelUpdateOne) SaveX(ctx context.Context) *VehicleModel {
	node, err := vmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vmuo *VehicleModelUpdateOne) Exec(ctx context.Context) error {
	_, err := vmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmuo *VehicleModelUpdateOne) ExecX(ctx context.Context) {
	if err := vmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vmuo *VehicleModelUpdateOne) defaults() error {
	if _, ok := vmuo.mutation.UpdatedAt(); !ok {
		if vehiclemodel.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized vehiclemodel.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := vehiclemodel.UpdateDefaultUpdatedAt()
		vmuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vmuo *VehicleModelUpdateOne) check() error {
	if v, ok := vmuo.mutation.Short(); ok {
		if err := vehiclemodel.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf("ent: validator failed for field \"short\": %w", err)}
		}
	}
	if v, ok := vmuo.mutation.Title(); ok {
		if err := vehiclemodel.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := vmuo.mutation.Description(); ok {
		if err := vehiclemodel.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (vmuo *VehicleModelUpdateOne) sqlSave(ctx context.Context) (_node *VehicleModel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vehiclemodel.Table,
			Columns: vehiclemodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehiclemodel.FieldID,
			},
		},
	}
	id, ok := vmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing VehicleModel.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := vmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehiclemodel.FieldID)
		for _, f := range fields {
			if !vehiclemodel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vehiclemodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vmuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldCreatedBy,
		})
	}
	if value, ok := vmuo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldCreatedBy,
		})
	}
	if vmuo.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: vehiclemodel.FieldCreatedBy,
		})
	}
	if value, ok := vmuo.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldCreatedWith,
		})
	}
	if vmuo.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehiclemodel.FieldCreatedWith,
		})
	}
	if value, ok := vmuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedAt,
		})
	}
	if value, ok := vmuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedBy,
		})
	}
	if value, ok := vmuo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedBy,
		})
	}
	if vmuo.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: vehiclemodel.FieldUpdatedBy,
		})
	}
	if value, ok := vmuo.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldUpdatedWith,
		})
	}
	if vmuo.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehiclemodel.FieldUpdatedWith,
		})
	}
	if value, ok := vmuo.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldShort,
		})
	}
	if value, ok := vmuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldTitle,
		})
	}
	if value, ok := vmuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehiclemodel.FieldDescription,
		})
	}
	if vmuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehiclemodel.FieldDescription,
		})
	}
	if vmuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vehiclemodel.MetadataTable,
			Columns: []string{vehiclemodel.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vmuo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vehiclemodel.MetadataTable,
			Columns: []string{vehiclemodel.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metadata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vmuo.mutation.VehiclesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vehiclemodel.VehiclesTable,
			Columns: []string{vehiclemodel.VehiclesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vehicle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vmuo.mutation.RemovedVehiclesIDs(); len(nodes) > 0 && !vmuo.mutation.VehiclesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vehiclemodel.VehiclesTable,
			Columns: []string{vehiclemodel.VehiclesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vehicle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vmuo.mutation.VehiclesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vehiclemodel.VehiclesTable,
			Columns: []string{vehiclemodel.VehiclesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vehicle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &VehicleModel{config: vmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehiclemodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
