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
	"github.com/responserms/response/internal/ent/player"
	"github.com/responserms/response/internal/ent/playeridentifier"
)

// PlayerIdentifierCreate is the builder for creating a PlayerIdentifier entity.
type PlayerIdentifierCreate struct {
	config
	mutation *PlayerIdentifierMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pic *PlayerIdentifierCreate) SetCreatedAt(t time.Time) *PlayerIdentifierCreate {
	pic.mutation.SetCreatedAt(t)
	return pic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pic *PlayerIdentifierCreate) SetNillableCreatedAt(t *time.Time) *PlayerIdentifierCreate {
	if t != nil {
		pic.SetCreatedAt(*t)
	}
	return pic
}

// SetCreatedBy sets the "created_by" field.
func (pic *PlayerIdentifierCreate) SetCreatedBy(i int) *PlayerIdentifierCreate {
	pic.mutation.SetCreatedBy(i)
	return pic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pic *PlayerIdentifierCreate) SetNillableCreatedBy(i *int) *PlayerIdentifierCreate {
	if i != nil {
		pic.SetCreatedBy(*i)
	}
	return pic
}

// SetCreatedWith sets the "created_with" field.
func (pic *PlayerIdentifierCreate) SetCreatedWith(s string) *PlayerIdentifierCreate {
	pic.mutation.SetCreatedWith(s)
	return pic
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (pic *PlayerIdentifierCreate) SetNillableCreatedWith(s *string) *PlayerIdentifierCreate {
	if s != nil {
		pic.SetCreatedWith(*s)
	}
	return pic
}

// SetUpdatedAt sets the "updated_at" field.
func (pic *PlayerIdentifierCreate) SetUpdatedAt(t time.Time) *PlayerIdentifierCreate {
	pic.mutation.SetUpdatedAt(t)
	return pic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pic *PlayerIdentifierCreate) SetNillableUpdatedAt(t *time.Time) *PlayerIdentifierCreate {
	if t != nil {
		pic.SetUpdatedAt(*t)
	}
	return pic
}

// SetUpdatedBy sets the "updated_by" field.
func (pic *PlayerIdentifierCreate) SetUpdatedBy(i int) *PlayerIdentifierCreate {
	pic.mutation.SetUpdatedBy(i)
	return pic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pic *PlayerIdentifierCreate) SetNillableUpdatedBy(i *int) *PlayerIdentifierCreate {
	if i != nil {
		pic.SetUpdatedBy(*i)
	}
	return pic
}

// SetUpdatedWith sets the "updated_with" field.
func (pic *PlayerIdentifierCreate) SetUpdatedWith(s string) *PlayerIdentifierCreate {
	pic.mutation.SetUpdatedWith(s)
	return pic
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (pic *PlayerIdentifierCreate) SetNillableUpdatedWith(s *string) *PlayerIdentifierCreate {
	if s != nil {
		pic.SetUpdatedWith(*s)
	}
	return pic
}

// SetValue sets the "value" field.
func (pic *PlayerIdentifierCreate) SetValue(s string) *PlayerIdentifierCreate {
	pic.mutation.SetValue(s)
	return pic
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (pic *PlayerIdentifierCreate) SetPlayerID(id int) *PlayerIdentifierCreate {
	pic.mutation.SetPlayerID(id)
	return pic
}

// SetNillablePlayerID sets the "player" edge to the Player entity by ID if the given value is not nil.
func (pic *PlayerIdentifierCreate) SetNillablePlayerID(id *int) *PlayerIdentifierCreate {
	if id != nil {
		pic = pic.SetPlayerID(*id)
	}
	return pic
}

// SetPlayer sets the "player" edge to the Player entity.
func (pic *PlayerIdentifierCreate) SetPlayer(p *Player) *PlayerIdentifierCreate {
	return pic.SetPlayerID(p.ID)
}

// Mutation returns the PlayerIdentifierMutation object of the builder.
func (pic *PlayerIdentifierCreate) Mutation() *PlayerIdentifierMutation {
	return pic.mutation
}

// Save creates the PlayerIdentifier in the database.
func (pic *PlayerIdentifierCreate) Save(ctx context.Context) (*PlayerIdentifier, error) {
	var (
		err  error
		node *PlayerIdentifier
	)
	if err := pic.defaults(); err != nil {
		return nil, err
	}
	if len(pic.hooks) == 0 {
		if err = pic.check(); err != nil {
			return nil, err
		}
		node, err = pic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerIdentifierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pic.check(); err != nil {
				return nil, err
			}
			pic.mutation = mutation
			if node, err = pic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pic.hooks) - 1; i >= 0; i-- {
			if pic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PlayerIdentifierCreate) SaveX(ctx context.Context) *PlayerIdentifier {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PlayerIdentifierCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PlayerIdentifierCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *PlayerIdentifierCreate) defaults() error {
	if _, ok := pic.mutation.CreatedAt(); !ok {
		if playeridentifier.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized playeridentifier.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := playeridentifier.DefaultCreatedAt()
		pic.mutation.SetCreatedAt(v)
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		if playeridentifier.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized playeridentifier.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := playeridentifier.DefaultUpdatedAt()
		pic.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pic *PlayerIdentifierCreate) check() error {
	if _, ok := pic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := pic.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "value"`)}
	}
	if v, ok := pic.mutation.Value(); ok {
		if err := playeridentifier.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "value": %w`, err)}
		}
	}
	return nil
}

func (pic *PlayerIdentifierCreate) sqlSave(ctx context.Context) (*PlayerIdentifier, error) {
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pic *PlayerIdentifierCreate) createSpec() (*PlayerIdentifier, *sqlgraph.CreateSpec) {
	var (
		_node = &PlayerIdentifier{config: pic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: playeridentifier.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playeridentifier.FieldID,
			},
		}
	)
	_spec.OnConflict = pic.conflict
	if value, ok := pic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playeridentifier.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pic.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := pic.mutation.CreatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldCreatedWith,
		})
		_node.CreatedWith = value
	}
	if value, ok := pic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playeridentifier.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pic.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := pic.mutation.UpdatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldUpdatedWith,
		})
		_node.UpdatedWith = value
	}
	if value, ok := pic.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldValue,
		})
		_node.Value = value
	}
	if nodes := pic.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playeridentifier.PlayerTable,
			Columns: []string{playeridentifier.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_identifiers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlayerIdentifier.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerIdentifierUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pic *PlayerIdentifierCreate) OnConflict(opts ...sql.ConflictOption) *PlayerIdentifierUpsertOne {
	pic.conflict = opts
	return &PlayerIdentifierUpsertOne{
		create: pic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlayerIdentifier.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pic *PlayerIdentifierCreate) OnConflictColumns(columns ...string) *PlayerIdentifierUpsertOne {
	pic.conflict = append(pic.conflict, sql.ConflictColumns(columns...))
	return &PlayerIdentifierUpsertOne{
		create: pic,
	}
}

type (
	// PlayerIdentifierUpsertOne is the builder for "upsert"-ing
	//  one PlayerIdentifier node.
	PlayerIdentifierUpsertOne struct {
		create *PlayerIdentifierCreate
	}

	// PlayerIdentifierUpsert is the "OnConflict" setter.
	PlayerIdentifierUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *PlayerIdentifierUpsert) SetCreatedAt(v time.Time) *PlayerIdentifierUpsert {
	u.Set(playeridentifier.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PlayerIdentifierUpsert) UpdateCreatedAt() *PlayerIdentifierUpsert {
	u.SetExcluded(playeridentifier.FieldCreatedAt)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *PlayerIdentifierUpsert) SetCreatedBy(v int) *PlayerIdentifierUpsert {
	u.Set(playeridentifier.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *PlayerIdentifierUpsert) UpdateCreatedBy() *PlayerIdentifierUpsert {
	u.SetExcluded(playeridentifier.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *PlayerIdentifierUpsert) ClearCreatedBy() *PlayerIdentifierUpsert {
	u.SetNull(playeridentifier.FieldCreatedBy)
	return u
}

// SetCreatedWith sets the "created_with" field.
func (u *PlayerIdentifierUpsert) SetCreatedWith(v string) *PlayerIdentifierUpsert {
	u.Set(playeridentifier.FieldCreatedWith, v)
	return u
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *PlayerIdentifierUpsert) UpdateCreatedWith() *PlayerIdentifierUpsert {
	u.SetExcluded(playeridentifier.FieldCreatedWith)
	return u
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *PlayerIdentifierUpsert) ClearCreatedWith() *PlayerIdentifierUpsert {
	u.SetNull(playeridentifier.FieldCreatedWith)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlayerIdentifierUpsert) SetUpdatedAt(v time.Time) *PlayerIdentifierUpsert {
	u.Set(playeridentifier.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlayerIdentifierUpsert) UpdateUpdatedAt() *PlayerIdentifierUpsert {
	u.SetExcluded(playeridentifier.FieldUpdatedAt)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *PlayerIdentifierUpsert) SetUpdatedBy(v int) *PlayerIdentifierUpsert {
	u.Set(playeridentifier.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *PlayerIdentifierUpsert) UpdateUpdatedBy() *PlayerIdentifierUpsert {
	u.SetExcluded(playeridentifier.FieldUpdatedBy)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *PlayerIdentifierUpsert) ClearUpdatedBy() *PlayerIdentifierUpsert {
	u.SetNull(playeridentifier.FieldUpdatedBy)
	return u
}

// SetUpdatedWith sets the "updated_with" field.
func (u *PlayerIdentifierUpsert) SetUpdatedWith(v string) *PlayerIdentifierUpsert {
	u.Set(playeridentifier.FieldUpdatedWith, v)
	return u
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *PlayerIdentifierUpsert) UpdateUpdatedWith() *PlayerIdentifierUpsert {
	u.SetExcluded(playeridentifier.FieldUpdatedWith)
	return u
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *PlayerIdentifierUpsert) ClearUpdatedWith() *PlayerIdentifierUpsert {
	u.SetNull(playeridentifier.FieldUpdatedWith)
	return u
}

// SetValue sets the "value" field.
func (u *PlayerIdentifierUpsert) SetValue(v string) *PlayerIdentifierUpsert {
	u.Set(playeridentifier.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *PlayerIdentifierUpsert) UpdateValue() *PlayerIdentifierUpsert {
	u.SetExcluded(playeridentifier.FieldValue)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.PlayerIdentifier.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *PlayerIdentifierUpsertOne) UpdateNewValues() *PlayerIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.PlayerIdentifier.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PlayerIdentifierUpsertOne) Ignore() *PlayerIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerIdentifierUpsertOne) DoNothing() *PlayerIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerIdentifierCreate.OnConflict
// documentation for more info.
func (u *PlayerIdentifierUpsertOne) Update(set func(*PlayerIdentifierUpsert)) *PlayerIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerIdentifierUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PlayerIdentifierUpsertOne) SetCreatedAt(v time.Time) *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertOne) UpdateCreatedAt() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *PlayerIdentifierUpsertOne) SetCreatedBy(v int) *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertOne) UpdateCreatedBy() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *PlayerIdentifierUpsertOne) ClearCreatedBy() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *PlayerIdentifierUpsertOne) SetCreatedWith(v string) *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertOne) UpdateCreatedWith() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *PlayerIdentifierUpsertOne) ClearCreatedWith() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlayerIdentifierUpsertOne) SetUpdatedAt(v time.Time) *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertOne) UpdateUpdatedAt() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *PlayerIdentifierUpsertOne) SetUpdatedBy(v int) *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertOne) UpdateUpdatedBy() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *PlayerIdentifierUpsertOne) ClearUpdatedBy() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *PlayerIdentifierUpsertOne) SetUpdatedWith(v string) *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertOne) UpdateUpdatedWith() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *PlayerIdentifierUpsertOne) ClearUpdatedWith() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetValue sets the "value" field.
func (u *PlayerIdentifierUpsertOne) SetValue(v string) *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertOne) UpdateValue() *PlayerIdentifierUpsertOne {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *PlayerIdentifierUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerIdentifierCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerIdentifierUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlayerIdentifierUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlayerIdentifierUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlayerIdentifierCreateBulk is the builder for creating many PlayerIdentifier entities in bulk.
type PlayerIdentifierCreateBulk struct {
	config
	builders []*PlayerIdentifierCreate
	conflict []sql.ConflictOption
}

// Save creates the PlayerIdentifier entities in the database.
func (picb *PlayerIdentifierCreateBulk) Save(ctx context.Context) ([]*PlayerIdentifier, error) {
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PlayerIdentifier, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerIdentifierMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = picb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PlayerIdentifierCreateBulk) SaveX(ctx context.Context) []*PlayerIdentifier {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PlayerIdentifierCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PlayerIdentifierCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlayerIdentifier.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerIdentifierUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (picb *PlayerIdentifierCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlayerIdentifierUpsertBulk {
	picb.conflict = opts
	return &PlayerIdentifierUpsertBulk{
		create: picb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlayerIdentifier.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (picb *PlayerIdentifierCreateBulk) OnConflictColumns(columns ...string) *PlayerIdentifierUpsertBulk {
	picb.conflict = append(picb.conflict, sql.ConflictColumns(columns...))
	return &PlayerIdentifierUpsertBulk{
		create: picb,
	}
}

// PlayerIdentifierUpsertBulk is the builder for "upsert"-ing
// a bulk of PlayerIdentifier nodes.
type PlayerIdentifierUpsertBulk struct {
	create *PlayerIdentifierCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PlayerIdentifier.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *PlayerIdentifierUpsertBulk) UpdateNewValues() *PlayerIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PlayerIdentifier.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PlayerIdentifierUpsertBulk) Ignore() *PlayerIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerIdentifierUpsertBulk) DoNothing() *PlayerIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerIdentifierCreateBulk.OnConflict
// documentation for more info.
func (u *PlayerIdentifierUpsertBulk) Update(set func(*PlayerIdentifierUpsert)) *PlayerIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerIdentifierUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PlayerIdentifierUpsertBulk) SetCreatedAt(v time.Time) *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertBulk) UpdateCreatedAt() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *PlayerIdentifierUpsertBulk) SetCreatedBy(v int) *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertBulk) UpdateCreatedBy() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *PlayerIdentifierUpsertBulk) ClearCreatedBy() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *PlayerIdentifierUpsertBulk) SetCreatedWith(v string) *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertBulk) UpdateCreatedWith() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *PlayerIdentifierUpsertBulk) ClearCreatedWith() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlayerIdentifierUpsertBulk) SetUpdatedAt(v time.Time) *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertBulk) UpdateUpdatedAt() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *PlayerIdentifierUpsertBulk) SetUpdatedBy(v int) *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertBulk) UpdateUpdatedBy() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *PlayerIdentifierUpsertBulk) ClearUpdatedBy() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *PlayerIdentifierUpsertBulk) SetUpdatedWith(v string) *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertBulk) UpdateUpdatedWith() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *PlayerIdentifierUpsertBulk) ClearUpdatedWith() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetValue sets the "value" field.
func (u *PlayerIdentifierUpsertBulk) SetValue(v string) *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *PlayerIdentifierUpsertBulk) UpdateValue() *PlayerIdentifierUpsertBulk {
	return u.Update(func(s *PlayerIdentifierUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *PlayerIdentifierUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlayerIdentifierCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerIdentifierCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerIdentifierUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
