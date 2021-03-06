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
	"github.com/responserms/response/internal/ent/setting"
	"github.com/responserms/response/internal/structs"
)

// SettingCreate is the builder for creating a Setting entity.
type SettingCreate struct {
	config
	mutation *SettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SettingCreate) SetCreatedAt(t time.Time) *SettingCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreatedAt(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *SettingCreate) SetCreatedBy(i int) *SettingCreate {
	sc.mutation.SetCreatedBy(i)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreatedBy(i *int) *SettingCreate {
	if i != nil {
		sc.SetCreatedBy(*i)
	}
	return sc
}

// SetCreatedWith sets the "created_with" field.
func (sc *SettingCreate) SetCreatedWith(s string) *SettingCreate {
	sc.mutation.SetCreatedWith(s)
	return sc
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreatedWith(s *string) *SettingCreate {
	if s != nil {
		sc.SetCreatedWith(*s)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SettingCreate) SetUpdatedAt(t time.Time) *SettingCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SettingCreate) SetNillableUpdatedAt(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SettingCreate) SetUpdatedBy(i int) *SettingCreate {
	sc.mutation.SetUpdatedBy(i)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *SettingCreate) SetNillableUpdatedBy(i *int) *SettingCreate {
	if i != nil {
		sc.SetUpdatedBy(*i)
	}
	return sc
}

// SetUpdatedWith sets the "updated_with" field.
func (sc *SettingCreate) SetUpdatedWith(s string) *SettingCreate {
	sc.mutation.SetUpdatedWith(s)
	return sc
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (sc *SettingCreate) SetNillableUpdatedWith(s *string) *SettingCreate {
	if s != nil {
		sc.SetUpdatedWith(*s)
	}
	return sc
}

// SetKey sets the "key" field.
func (sc *SettingCreate) SetKey(s string) *SettingCreate {
	sc.mutation.SetKey(s)
	return sc
}

// SetData sets the "data" field.
func (sc *SettingCreate) SetData(sd structs.SettingData) *SettingCreate {
	sc.mutation.SetData(sd)
	return sc
}

// SetNillableData sets the "data" field if the given value is not nil.
func (sc *SettingCreate) SetNillableData(sd *structs.SettingData) *SettingCreate {
	if sd != nil {
		sc.SetData(*sd)
	}
	return sc
}

// Mutation returns the SettingMutation object of the builder.
func (sc *SettingCreate) Mutation() *SettingMutation {
	return sc.mutation
}

// Save creates the Setting in the database.
func (sc *SettingCreate) Save(ctx context.Context) (*Setting, error) {
	var (
		err  error
		node *Setting
	)
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingCreate) SaveX(ctx context.Context) *Setting {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettingCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if setting.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized setting.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := setting.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if setting.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized setting.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := setting.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := sc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "key"`)}
	}
	return nil
}

func (sc *SettingCreate) sqlSave(ctx context.Context) (*Setting, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SettingCreate) createSpec() (*Setting, *sqlgraph.CreateSpec) {
	var (
		_node = &Setting{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: setting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: setting.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: setting.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: setting.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.CreatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: setting.FieldCreatedWith,
		})
		_node.CreatedWith = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: setting.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: setting.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.UpdatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: setting.FieldUpdatedWith,
		})
		_node.UpdatedWith = value
	}
	if value, ok := sc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: setting.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := sc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: setting.FieldData,
		})
		_node.Data = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (sc *SettingCreate) OnConflict(opts ...sql.ConflictOption) *SettingUpsertOne {
	sc.conflict = opts
	return &SettingUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *SettingCreate) OnConflictColumns(columns ...string) *SettingUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertOne{
		create: sc,
	}
}

type (
	// SettingUpsertOne is the builder for "upsert"-ing
	//  one Setting node.
	SettingUpsertOne struct {
		create *SettingCreate
	}

	// SettingUpsert is the "OnConflict" setter.
	SettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *SettingUpsert) SetCreatedAt(v time.Time) *SettingUpsert {
	u.Set(setting.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SettingUpsert) UpdateCreatedAt() *SettingUpsert {
	u.SetExcluded(setting.FieldCreatedAt)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *SettingUpsert) SetCreatedBy(v int) *SettingUpsert {
	u.Set(setting.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *SettingUpsert) UpdateCreatedBy() *SettingUpsert {
	u.SetExcluded(setting.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *SettingUpsert) ClearCreatedBy() *SettingUpsert {
	u.SetNull(setting.FieldCreatedBy)
	return u
}

// SetCreatedWith sets the "created_with" field.
func (u *SettingUpsert) SetCreatedWith(v string) *SettingUpsert {
	u.Set(setting.FieldCreatedWith, v)
	return u
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *SettingUpsert) UpdateCreatedWith() *SettingUpsert {
	u.SetExcluded(setting.FieldCreatedWith)
	return u
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *SettingUpsert) ClearCreatedWith() *SettingUpsert {
	u.SetNull(setting.FieldCreatedWith)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SettingUpsert) SetUpdatedAt(v time.Time) *SettingUpsert {
	u.Set(setting.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SettingUpsert) UpdateUpdatedAt() *SettingUpsert {
	u.SetExcluded(setting.FieldUpdatedAt)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SettingUpsert) SetUpdatedBy(v int) *SettingUpsert {
	u.Set(setting.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SettingUpsert) UpdateUpdatedBy() *SettingUpsert {
	u.SetExcluded(setting.FieldUpdatedBy)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SettingUpsert) ClearUpdatedBy() *SettingUpsert {
	u.SetNull(setting.FieldUpdatedBy)
	return u
}

// SetUpdatedWith sets the "updated_with" field.
func (u *SettingUpsert) SetUpdatedWith(v string) *SettingUpsert {
	u.Set(setting.FieldUpdatedWith, v)
	return u
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *SettingUpsert) UpdateUpdatedWith() *SettingUpsert {
	u.SetExcluded(setting.FieldUpdatedWith)
	return u
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *SettingUpsert) ClearUpdatedWith() *SettingUpsert {
	u.SetNull(setting.FieldUpdatedWith)
	return u
}

// SetKey sets the "key" field.
func (u *SettingUpsert) SetKey(v string) *SettingUpsert {
	u.Set(setting.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SettingUpsert) UpdateKey() *SettingUpsert {
	u.SetExcluded(setting.FieldKey)
	return u
}

// SetData sets the "data" field.
func (u *SettingUpsert) SetData(v structs.SettingData) *SettingUpsert {
	u.Set(setting.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *SettingUpsert) UpdateData() *SettingUpsert {
	u.SetExcluded(setting.FieldData)
	return u
}

// ClearData clears the value of the "data" field.
func (u *SettingUpsert) ClearData() *SettingUpsert {
	u.SetNull(setting.FieldData)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *SettingUpsertOne) UpdateNewValues() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Setting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *SettingUpsertOne) Ignore() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertOne) DoNothing() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreate.OnConflict
// documentation for more info.
func (u *SettingUpsertOne) Update(set func(*SettingUpsert)) *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SettingUpsertOne) SetCreatedAt(v time.Time) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateCreatedAt() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *SettingUpsertOne) SetCreatedBy(v int) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateCreatedBy() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *SettingUpsertOne) ClearCreatedBy() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *SettingUpsertOne) SetCreatedWith(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateCreatedWith() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *SettingUpsertOne) ClearCreatedWith() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SettingUpsertOne) SetUpdatedAt(v time.Time) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateUpdatedAt() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SettingUpsertOne) SetUpdatedBy(v int) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateUpdatedBy() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SettingUpsertOne) ClearUpdatedBy() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *SettingUpsertOne) SetUpdatedWith(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateUpdatedWith() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *SettingUpsertOne) ClearUpdatedWith() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetKey sets the "key" field.
func (u *SettingUpsertOne) SetKey(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateKey() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateKey()
	})
}

// SetData sets the "data" field.
func (u *SettingUpsertOne) SetData(v structs.SettingData) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateData() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *SettingUpsertOne) ClearData() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearData()
	})
}

// Exec executes the query.
func (u *SettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SettingUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SettingUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SettingCreateBulk is the builder for creating many Setting entities in bulk.
type SettingCreateBulk struct {
	config
	builders []*SettingCreate
	conflict []sql.ConflictOption
}

// Save creates the Setting entities in the database.
func (scb *SettingCreateBulk) Save(ctx context.Context) ([]*Setting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Setting, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingCreateBulk) SaveX(ctx context.Context) []*Setting {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (scb *SettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *SettingUpsertBulk {
	scb.conflict = opts
	return &SettingUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *SettingCreateBulk) OnConflictColumns(columns ...string) *SettingUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertBulk{
		create: scb,
	}
}

// SettingUpsertBulk is the builder for "upsert"-ing
// a bulk of Setting nodes.
type SettingUpsertBulk struct {
	create *SettingCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *SettingUpsertBulk) UpdateNewValues() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *SettingUpsertBulk) Ignore() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertBulk) DoNothing() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreateBulk.OnConflict
// documentation for more info.
func (u *SettingUpsertBulk) Update(set func(*SettingUpsert)) *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SettingUpsertBulk) SetCreatedAt(v time.Time) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateCreatedAt() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *SettingUpsertBulk) SetCreatedBy(v int) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateCreatedBy() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *SettingUpsertBulk) ClearCreatedBy() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *SettingUpsertBulk) SetCreatedWith(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateCreatedWith() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *SettingUpsertBulk) ClearCreatedWith() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SettingUpsertBulk) SetUpdatedAt(v time.Time) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateUpdatedAt() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SettingUpsertBulk) SetUpdatedBy(v int) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateUpdatedBy() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SettingUpsertBulk) ClearUpdatedBy() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *SettingUpsertBulk) SetUpdatedWith(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateUpdatedWith() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *SettingUpsertBulk) ClearUpdatedWith() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetKey sets the "key" field.
func (u *SettingUpsertBulk) SetKey(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateKey() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateKey()
	})
}

// SetData sets the "data" field.
func (u *SettingUpsertBulk) SetData(v structs.SettingData) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateData() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *SettingUpsertBulk) ClearData() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearData()
	})
}

// Exec executes the query.
func (u *SettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
