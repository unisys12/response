// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/person"
	"github.com/responserms/response/internal/ent/state"
	"github.com/responserms/response/internal/ent/vehicle"
	"github.com/responserms/response/internal/ent/vehicleregistration"
)

// VehicleRegistrationCreate is the builder for creating a VehicleRegistration entity.
type VehicleRegistrationCreate struct {
	config
	mutation *VehicleRegistrationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (vrc *VehicleRegistrationCreate) SetCreatedAt(t time.Time) *VehicleRegistrationCreate {
	vrc.mutation.SetCreatedAt(t)
	return vrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vrc *VehicleRegistrationCreate) SetNillableCreatedAt(t *time.Time) *VehicleRegistrationCreate {
	if t != nil {
		vrc.SetCreatedAt(*t)
	}
	return vrc
}

// SetCreatedBy sets the "created_by" field.
func (vrc *VehicleRegistrationCreate) SetCreatedBy(i int) *VehicleRegistrationCreate {
	vrc.mutation.SetCreatedBy(i)
	return vrc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vrc *VehicleRegistrationCreate) SetNillableCreatedBy(i *int) *VehicleRegistrationCreate {
	if i != nil {
		vrc.SetCreatedBy(*i)
	}
	return vrc
}

// SetCreatedWith sets the "created_with" field.
func (vrc *VehicleRegistrationCreate) SetCreatedWith(s string) *VehicleRegistrationCreate {
	vrc.mutation.SetCreatedWith(s)
	return vrc
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (vrc *VehicleRegistrationCreate) SetNillableCreatedWith(s *string) *VehicleRegistrationCreate {
	if s != nil {
		vrc.SetCreatedWith(*s)
	}
	return vrc
}

// SetUpdatedAt sets the "updated_at" field.
func (vrc *VehicleRegistrationCreate) SetUpdatedAt(t time.Time) *VehicleRegistrationCreate {
	vrc.mutation.SetUpdatedAt(t)
	return vrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vrc *VehicleRegistrationCreate) SetNillableUpdatedAt(t *time.Time) *VehicleRegistrationCreate {
	if t != nil {
		vrc.SetUpdatedAt(*t)
	}
	return vrc
}

// SetUpdatedBy sets the "updated_by" field.
func (vrc *VehicleRegistrationCreate) SetUpdatedBy(i int) *VehicleRegistrationCreate {
	vrc.mutation.SetUpdatedBy(i)
	return vrc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vrc *VehicleRegistrationCreate) SetNillableUpdatedBy(i *int) *VehicleRegistrationCreate {
	if i != nil {
		vrc.SetUpdatedBy(*i)
	}
	return vrc
}

// SetUpdatedWith sets the "updated_with" field.
func (vrc *VehicleRegistrationCreate) SetUpdatedWith(s string) *VehicleRegistrationCreate {
	vrc.mutation.SetUpdatedWith(s)
	return vrc
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (vrc *VehicleRegistrationCreate) SetNillableUpdatedWith(s *string) *VehicleRegistrationCreate {
	if s != nil {
		vrc.SetUpdatedWith(*s)
	}
	return vrc
}

// SetPlate sets the "plate" field.
func (vrc *VehicleRegistrationCreate) SetPlate(s string) *VehicleRegistrationCreate {
	vrc.mutation.SetPlate(s)
	return vrc
}

// SetExpiredAt sets the "expired_at" field.
func (vrc *VehicleRegistrationCreate) SetExpiredAt(t time.Time) *VehicleRegistrationCreate {
	vrc.mutation.SetExpiredAt(t)
	return vrc
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (vrc *VehicleRegistrationCreate) SetMetadataID(id int) *VehicleRegistrationCreate {
	vrc.mutation.SetMetadataID(id)
	return vrc
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (vrc *VehicleRegistrationCreate) SetNillableMetadataID(id *int) *VehicleRegistrationCreate {
	if id != nil {
		vrc = vrc.SetMetadataID(*id)
	}
	return vrc
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (vrc *VehicleRegistrationCreate) SetMetadata(m *Metadata) *VehicleRegistrationCreate {
	return vrc.SetMetadataID(m.ID)
}

// SetStateID sets the "state" edge to the State entity by ID.
func (vrc *VehicleRegistrationCreate) SetStateID(id int) *VehicleRegistrationCreate {
	vrc.mutation.SetStateID(id)
	return vrc
}

// SetState sets the "state" edge to the State entity.
func (vrc *VehicleRegistrationCreate) SetState(s *State) *VehicleRegistrationCreate {
	return vrc.SetStateID(s.ID)
}

// SetPersonID sets the "person" edge to the Person entity by ID.
func (vrc *VehicleRegistrationCreate) SetPersonID(id int) *VehicleRegistrationCreate {
	vrc.mutation.SetPersonID(id)
	return vrc
}

// SetPerson sets the "person" edge to the Person entity.
func (vrc *VehicleRegistrationCreate) SetPerson(p *Person) *VehicleRegistrationCreate {
	return vrc.SetPersonID(p.ID)
}

// SetVehicleID sets the "vehicle" edge to the Vehicle entity by ID.
func (vrc *VehicleRegistrationCreate) SetVehicleID(id int) *VehicleRegistrationCreate {
	vrc.mutation.SetVehicleID(id)
	return vrc
}

// SetVehicle sets the "vehicle" edge to the Vehicle entity.
func (vrc *VehicleRegistrationCreate) SetVehicle(v *Vehicle) *VehicleRegistrationCreate {
	return vrc.SetVehicleID(v.ID)
}

// Mutation returns the VehicleRegistrationMutation object of the builder.
func (vrc *VehicleRegistrationCreate) Mutation() *VehicleRegistrationMutation {
	return vrc.mutation
}

// Save creates the VehicleRegistration in the database.
func (vrc *VehicleRegistrationCreate) Save(ctx context.Context) (*VehicleRegistration, error) {
	var (
		err  error
		node *VehicleRegistration
	)
	if err := vrc.defaults(); err != nil {
		return nil, err
	}
	if len(vrc.hooks) == 0 {
		if err = vrc.check(); err != nil {
			return nil, err
		}
		node, err = vrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleRegistrationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vrc.check(); err != nil {
				return nil, err
			}
			vrc.mutation = mutation
			if node, err = vrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vrc.hooks) - 1; i >= 0; i-- {
			if vrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vrc *VehicleRegistrationCreate) SaveX(ctx context.Context) *VehicleRegistration {
	v, err := vrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vrc *VehicleRegistrationCreate) Exec(ctx context.Context) error {
	_, err := vrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vrc *VehicleRegistrationCreate) ExecX(ctx context.Context) {
	if err := vrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vrc *VehicleRegistrationCreate) defaults() error {
	if _, ok := vrc.mutation.CreatedAt(); !ok {
		if vehicleregistration.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized vehicleregistration.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := vehicleregistration.DefaultCreatedAt()
		vrc.mutation.SetCreatedAt(v)
	}
	if _, ok := vrc.mutation.UpdatedAt(); !ok {
		if vehicleregistration.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized vehicleregistration.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := vehicleregistration.DefaultUpdatedAt()
		vrc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vrc *VehicleRegistrationCreate) check() error {
	if _, ok := vrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := vrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := vrc.mutation.Plate(); !ok {
		return &ValidationError{Name: "plate", err: errors.New(`ent: missing required field "plate"`)}
	}
	if _, ok := vrc.mutation.ExpiredAt(); !ok {
		return &ValidationError{Name: "expired_at", err: errors.New(`ent: missing required field "expired_at"`)}
	}
	if _, ok := vrc.mutation.StateID(); !ok {
		return &ValidationError{Name: "state", err: errors.New("ent: missing required edge \"state\"")}
	}
	if _, ok := vrc.mutation.PersonID(); !ok {
		return &ValidationError{Name: "person", err: errors.New("ent: missing required edge \"person\"")}
	}
	if _, ok := vrc.mutation.VehicleID(); !ok {
		return &ValidationError{Name: "vehicle", err: errors.New("ent: missing required edge \"vehicle\"")}
	}
	return nil
}

func (vrc *VehicleRegistrationCreate) sqlSave(ctx context.Context) (*VehicleRegistration, error) {
	_node, _spec := vrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vrc *VehicleRegistrationCreate) createSpec() (*VehicleRegistration, *sqlgraph.CreateSpec) {
	var (
		_node = &VehicleRegistration{config: vrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vehicleregistration.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehicleregistration.FieldID,
			},
		}
	)
	_spec.OnConflict = vrc.conflict
	if value, ok := vrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicleregistration.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vrc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehicleregistration.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := vrc.mutation.CreatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicleregistration.FieldCreatedWith,
		})
		_node.CreatedWith = value
	}
	if value, ok := vrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicleregistration.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := vrc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vehicleregistration.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := vrc.mutation.UpdatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicleregistration.FieldUpdatedWith,
		})
		_node.UpdatedWith = value
	}
	if value, ok := vrc.mutation.Plate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicleregistration.FieldPlate,
		})
		_node.Plate = value
	}
	if value, ok := vrc.mutation.ExpiredAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicleregistration.FieldExpiredAt,
		})
		_node.ExpiredAt = value
	}
	if nodes := vrc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vehicleregistration.MetadataTable,
			Columns: []string{vehicleregistration.MetadataColumn},
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
		_node.vehicle_registration_metadata = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vrc.mutation.StateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehicleregistration.StateTable,
			Columns: []string{vehicleregistration.StateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: state.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.state_vehicle_registrations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vrc.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehicleregistration.PersonTable,
			Columns: []string{vehicleregistration.PersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.person_vehicle_registrations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vrc.mutation.VehicleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehicleregistration.VehicleTable,
			Columns: []string{vehicleregistration.VehicleColumn},
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
		_node.vehicle_registrations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.VehicleRegistration.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VehicleRegistrationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (vrc *VehicleRegistrationCreate) OnConflict(opts ...sql.ConflictOption) *VehicleRegistrationUpsertOne {
	vrc.conflict = opts
	return &VehicleRegistrationUpsertOne{
		create: vrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VehicleRegistration.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (vrc *VehicleRegistrationCreate) OnConflictColumns(columns ...string) *VehicleRegistrationUpsertOne {
	vrc.conflict = append(vrc.conflict, sql.ConflictColumns(columns...))
	return &VehicleRegistrationUpsertOne{
		create: vrc,
	}
}

type (
	// VehicleRegistrationUpsertOne is the builder for "upsert"-ing
	//  one VehicleRegistration node.
	VehicleRegistrationUpsertOne struct {
		create *VehicleRegistrationCreate
	}

	// VehicleRegistrationUpsert is the "OnConflict" setter.
	VehicleRegistrationUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *VehicleRegistrationUpsert) SetCreatedAt(v time.Time) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdateCreatedAt() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldCreatedAt)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *VehicleRegistrationUpsert) SetCreatedBy(v int) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdateCreatedBy() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *VehicleRegistrationUpsert) ClearCreatedBy() *VehicleRegistrationUpsert {
	u.SetNull(vehicleregistration.FieldCreatedBy)
	return u
}

// SetCreatedWith sets the "created_with" field.
func (u *VehicleRegistrationUpsert) SetCreatedWith(v string) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldCreatedWith, v)
	return u
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdateCreatedWith() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldCreatedWith)
	return u
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *VehicleRegistrationUpsert) ClearCreatedWith() *VehicleRegistrationUpsert {
	u.SetNull(vehicleregistration.FieldCreatedWith)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VehicleRegistrationUpsert) SetUpdatedAt(v time.Time) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdateUpdatedAt() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldUpdatedAt)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *VehicleRegistrationUpsert) SetUpdatedBy(v int) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdateUpdatedBy() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldUpdatedBy)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *VehicleRegistrationUpsert) ClearUpdatedBy() *VehicleRegistrationUpsert {
	u.SetNull(vehicleregistration.FieldUpdatedBy)
	return u
}

// SetUpdatedWith sets the "updated_with" field.
func (u *VehicleRegistrationUpsert) SetUpdatedWith(v string) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldUpdatedWith, v)
	return u
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdateUpdatedWith() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldUpdatedWith)
	return u
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *VehicleRegistrationUpsert) ClearUpdatedWith() *VehicleRegistrationUpsert {
	u.SetNull(vehicleregistration.FieldUpdatedWith)
	return u
}

// SetPlate sets the "plate" field.
func (u *VehicleRegistrationUpsert) SetPlate(v string) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldPlate, v)
	return u
}

// UpdatePlate sets the "plate" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdatePlate() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldPlate)
	return u
}

// SetExpiredAt sets the "expired_at" field.
func (u *VehicleRegistrationUpsert) SetExpiredAt(v time.Time) *VehicleRegistrationUpsert {
	u.Set(vehicleregistration.FieldExpiredAt, v)
	return u
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsert) UpdateExpiredAt() *VehicleRegistrationUpsert {
	u.SetExcluded(vehicleregistration.FieldExpiredAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.VehicleRegistration.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *VehicleRegistrationUpsertOne) UpdateNewValues() *VehicleRegistrationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.VehicleRegistration.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *VehicleRegistrationUpsertOne) Ignore() *VehicleRegistrationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VehicleRegistrationUpsertOne) DoNothing() *VehicleRegistrationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VehicleRegistrationCreate.OnConflict
// documentation for more info.
func (u *VehicleRegistrationUpsertOne) Update(set func(*VehicleRegistrationUpsert)) *VehicleRegistrationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VehicleRegistrationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *VehicleRegistrationUpsertOne) SetCreatedAt(v time.Time) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdateCreatedAt() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *VehicleRegistrationUpsertOne) SetCreatedBy(v int) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdateCreatedBy() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *VehicleRegistrationUpsertOne) ClearCreatedBy() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *VehicleRegistrationUpsertOne) SetCreatedWith(v string) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdateCreatedWith() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *VehicleRegistrationUpsertOne) ClearCreatedWith() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VehicleRegistrationUpsertOne) SetUpdatedAt(v time.Time) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdateUpdatedAt() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *VehicleRegistrationUpsertOne) SetUpdatedBy(v int) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdateUpdatedBy() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *VehicleRegistrationUpsertOne) ClearUpdatedBy() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *VehicleRegistrationUpsertOne) SetUpdatedWith(v string) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdateUpdatedWith() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *VehicleRegistrationUpsertOne) ClearUpdatedWith() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetPlate sets the "plate" field.
func (u *VehicleRegistrationUpsertOne) SetPlate(v string) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetPlate(v)
	})
}

// UpdatePlate sets the "plate" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdatePlate() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdatePlate()
	})
}

// SetExpiredAt sets the "expired_at" field.
func (u *VehicleRegistrationUpsertOne) SetExpiredAt(v time.Time) *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetExpiredAt(v)
	})
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertOne) UpdateExpiredAt() *VehicleRegistrationUpsertOne {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateExpiredAt()
	})
}

// Exec executes the query.
func (u *VehicleRegistrationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VehicleRegistrationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VehicleRegistrationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *VehicleRegistrationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *VehicleRegistrationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// VehicleRegistrationCreateBulk is the builder for creating many VehicleRegistration entities in bulk.
type VehicleRegistrationCreateBulk struct {
	config
	builders []*VehicleRegistrationCreate
	conflict []sql.ConflictOption
}

// Save creates the VehicleRegistration entities in the database.
func (vrcb *VehicleRegistrationCreateBulk) Save(ctx context.Context) ([]*VehicleRegistration, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vrcb.builders))
	nodes := make([]*VehicleRegistration, len(vrcb.builders))
	mutators := make([]Mutator, len(vrcb.builders))
	for i := range vrcb.builders {
		func(i int, root context.Context) {
			builder := vrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VehicleRegistrationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vrcb *VehicleRegistrationCreateBulk) SaveX(ctx context.Context) []*VehicleRegistration {
	v, err := vrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vrcb *VehicleRegistrationCreateBulk) Exec(ctx context.Context) error {
	_, err := vrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vrcb *VehicleRegistrationCreateBulk) ExecX(ctx context.Context) {
	if err := vrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.VehicleRegistration.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VehicleRegistrationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (vrcb *VehicleRegistrationCreateBulk) OnConflict(opts ...sql.ConflictOption) *VehicleRegistrationUpsertBulk {
	vrcb.conflict = opts
	return &VehicleRegistrationUpsertBulk{
		create: vrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VehicleRegistration.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (vrcb *VehicleRegistrationCreateBulk) OnConflictColumns(columns ...string) *VehicleRegistrationUpsertBulk {
	vrcb.conflict = append(vrcb.conflict, sql.ConflictColumns(columns...))
	return &VehicleRegistrationUpsertBulk{
		create: vrcb,
	}
}

// VehicleRegistrationUpsertBulk is the builder for "upsert"-ing
// a bulk of VehicleRegistration nodes.
type VehicleRegistrationUpsertBulk struct {
	create *VehicleRegistrationCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.VehicleRegistration.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *VehicleRegistrationUpsertBulk) UpdateNewValues() *VehicleRegistrationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.VehicleRegistration.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *VehicleRegistrationUpsertBulk) Ignore() *VehicleRegistrationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VehicleRegistrationUpsertBulk) DoNothing() *VehicleRegistrationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VehicleRegistrationCreateBulk.OnConflict
// documentation for more info.
func (u *VehicleRegistrationUpsertBulk) Update(set func(*VehicleRegistrationUpsert)) *VehicleRegistrationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VehicleRegistrationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *VehicleRegistrationUpsertBulk) SetCreatedAt(v time.Time) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdateCreatedAt() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *VehicleRegistrationUpsertBulk) SetCreatedBy(v int) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdateCreatedBy() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *VehicleRegistrationUpsertBulk) ClearCreatedBy() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *VehicleRegistrationUpsertBulk) SetCreatedWith(v string) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdateCreatedWith() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *VehicleRegistrationUpsertBulk) ClearCreatedWith() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VehicleRegistrationUpsertBulk) SetUpdatedAt(v time.Time) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdateUpdatedAt() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *VehicleRegistrationUpsertBulk) SetUpdatedBy(v int) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdateUpdatedBy() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *VehicleRegistrationUpsertBulk) ClearUpdatedBy() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *VehicleRegistrationUpsertBulk) SetUpdatedWith(v string) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdateUpdatedWith() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *VehicleRegistrationUpsertBulk) ClearUpdatedWith() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetPlate sets the "plate" field.
func (u *VehicleRegistrationUpsertBulk) SetPlate(v string) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetPlate(v)
	})
}

// UpdatePlate sets the "plate" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdatePlate() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdatePlate()
	})
}

// SetExpiredAt sets the "expired_at" field.
func (u *VehicleRegistrationUpsertBulk) SetExpiredAt(v time.Time) *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.SetExpiredAt(v)
	})
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *VehicleRegistrationUpsertBulk) UpdateExpiredAt() *VehicleRegistrationUpsertBulk {
	return u.Update(func(s *VehicleRegistrationUpsert) {
		s.UpdateExpiredAt()
	})
}

// Exec executes the query.
func (u *VehicleRegistrationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the VehicleRegistrationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VehicleRegistrationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VehicleRegistrationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
