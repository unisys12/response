// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/player"
	"github.com/responserms/response/internal/ent/playeridentifier"
	"github.com/responserms/response/internal/ent/predicate"
)

// PlayerIdentifierUpdate is the builder for updating PlayerIdentifier entities.
type PlayerIdentifierUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerIdentifierMutation
}

// Where appends a list predicates to the PlayerIdentifierUpdate builder.
func (piu *PlayerIdentifierUpdate) Where(ps ...predicate.PlayerIdentifier) *PlayerIdentifierUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetCreatedBy sets the "created_by" field.
func (piu *PlayerIdentifierUpdate) SetCreatedBy(i int) *PlayerIdentifierUpdate {
	piu.mutation.ResetCreatedBy()
	piu.mutation.SetCreatedBy(i)
	return piu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (piu *PlayerIdentifierUpdate) SetNillableCreatedBy(i *int) *PlayerIdentifierUpdate {
	if i != nil {
		piu.SetCreatedBy(*i)
	}
	return piu
}

// AddCreatedBy adds i to the "created_by" field.
func (piu *PlayerIdentifierUpdate) AddCreatedBy(i int) *PlayerIdentifierUpdate {
	piu.mutation.AddCreatedBy(i)
	return piu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (piu *PlayerIdentifierUpdate) ClearCreatedBy() *PlayerIdentifierUpdate {
	piu.mutation.ClearCreatedBy()
	return piu
}

// SetCreatedWith sets the "created_with" field.
func (piu *PlayerIdentifierUpdate) SetCreatedWith(s string) *PlayerIdentifierUpdate {
	piu.mutation.SetCreatedWith(s)
	return piu
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (piu *PlayerIdentifierUpdate) SetNillableCreatedWith(s *string) *PlayerIdentifierUpdate {
	if s != nil {
		piu.SetCreatedWith(*s)
	}
	return piu
}

// ClearCreatedWith clears the value of the "created_with" field.
func (piu *PlayerIdentifierUpdate) ClearCreatedWith() *PlayerIdentifierUpdate {
	piu.mutation.ClearCreatedWith()
	return piu
}

// SetUpdatedAt sets the "updated_at" field.
func (piu *PlayerIdentifierUpdate) SetUpdatedAt(t time.Time) *PlayerIdentifierUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetUpdatedBy sets the "updated_by" field.
func (piu *PlayerIdentifierUpdate) SetUpdatedBy(i int) *PlayerIdentifierUpdate {
	piu.mutation.ResetUpdatedBy()
	piu.mutation.SetUpdatedBy(i)
	return piu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (piu *PlayerIdentifierUpdate) SetNillableUpdatedBy(i *int) *PlayerIdentifierUpdate {
	if i != nil {
		piu.SetUpdatedBy(*i)
	}
	return piu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (piu *PlayerIdentifierUpdate) AddUpdatedBy(i int) *PlayerIdentifierUpdate {
	piu.mutation.AddUpdatedBy(i)
	return piu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (piu *PlayerIdentifierUpdate) ClearUpdatedBy() *PlayerIdentifierUpdate {
	piu.mutation.ClearUpdatedBy()
	return piu
}

// SetUpdatedWith sets the "updated_with" field.
func (piu *PlayerIdentifierUpdate) SetUpdatedWith(s string) *PlayerIdentifierUpdate {
	piu.mutation.SetUpdatedWith(s)
	return piu
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (piu *PlayerIdentifierUpdate) SetNillableUpdatedWith(s *string) *PlayerIdentifierUpdate {
	if s != nil {
		piu.SetUpdatedWith(*s)
	}
	return piu
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (piu *PlayerIdentifierUpdate) ClearUpdatedWith() *PlayerIdentifierUpdate {
	piu.mutation.ClearUpdatedWith()
	return piu
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (piu *PlayerIdentifierUpdate) SetPlayerID(id int) *PlayerIdentifierUpdate {
	piu.mutation.SetPlayerID(id)
	return piu
}

// SetNillablePlayerID sets the "player" edge to the Player entity by ID if the given value is not nil.
func (piu *PlayerIdentifierUpdate) SetNillablePlayerID(id *int) *PlayerIdentifierUpdate {
	if id != nil {
		piu = piu.SetPlayerID(*id)
	}
	return piu
}

// SetPlayer sets the "player" edge to the Player entity.
func (piu *PlayerIdentifierUpdate) SetPlayer(p *Player) *PlayerIdentifierUpdate {
	return piu.SetPlayerID(p.ID)
}

// Mutation returns the PlayerIdentifierMutation object of the builder.
func (piu *PlayerIdentifierUpdate) Mutation() *PlayerIdentifierMutation {
	return piu.mutation
}

// ClearPlayer clears the "player" edge to the Player entity.
func (piu *PlayerIdentifierUpdate) ClearPlayer() *PlayerIdentifierUpdate {
	piu.mutation.ClearPlayer()
	return piu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *PlayerIdentifierUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := piu.defaults(); err != nil {
		return 0, err
	}
	if len(piu.hooks) == 0 {
		affected, err = piu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerIdentifierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piu.mutation = mutation
			affected, err = piu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(piu.hooks) - 1; i >= 0; i-- {
			if piu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = piu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PlayerIdentifierUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PlayerIdentifierUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PlayerIdentifierUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piu *PlayerIdentifierUpdate) defaults() error {
	if _, ok := piu.mutation.UpdatedAt(); !ok {
		if playeridentifier.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized playeridentifier.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := playeridentifier.UpdateDefaultUpdatedAt()
		piu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (piu *PlayerIdentifierUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playeridentifier.Table,
			Columns: playeridentifier.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playeridentifier.FieldID,
			},
		},
	}
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldCreatedBy,
		})
	}
	if value, ok := piu.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldCreatedBy,
		})
	}
	if piu.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: playeridentifier.FieldCreatedBy,
		})
	}
	if value, ok := piu.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldCreatedWith,
		})
	}
	if piu.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playeridentifier.FieldCreatedWith,
		})
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playeridentifier.FieldUpdatedAt,
		})
	}
	if value, ok := piu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldUpdatedBy,
		})
	}
	if value, ok := piu.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldUpdatedBy,
		})
	}
	if piu.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: playeridentifier.FieldUpdatedBy,
		})
	}
	if value, ok := piu.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldUpdatedWith,
		})
	}
	if piu.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playeridentifier.FieldUpdatedWith,
		})
	}
	if piu.mutation.PlayerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.PlayerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playeridentifier.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PlayerIdentifierUpdateOne is the builder for updating a single PlayerIdentifier entity.
type PlayerIdentifierUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerIdentifierMutation
}

// SetCreatedBy sets the "created_by" field.
func (piuo *PlayerIdentifierUpdateOne) SetCreatedBy(i int) *PlayerIdentifierUpdateOne {
	piuo.mutation.ResetCreatedBy()
	piuo.mutation.SetCreatedBy(i)
	return piuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (piuo *PlayerIdentifierUpdateOne) SetNillableCreatedBy(i *int) *PlayerIdentifierUpdateOne {
	if i != nil {
		piuo.SetCreatedBy(*i)
	}
	return piuo
}

// AddCreatedBy adds i to the "created_by" field.
func (piuo *PlayerIdentifierUpdateOne) AddCreatedBy(i int) *PlayerIdentifierUpdateOne {
	piuo.mutation.AddCreatedBy(i)
	return piuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (piuo *PlayerIdentifierUpdateOne) ClearCreatedBy() *PlayerIdentifierUpdateOne {
	piuo.mutation.ClearCreatedBy()
	return piuo
}

// SetCreatedWith sets the "created_with" field.
func (piuo *PlayerIdentifierUpdateOne) SetCreatedWith(s string) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetCreatedWith(s)
	return piuo
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (piuo *PlayerIdentifierUpdateOne) SetNillableCreatedWith(s *string) *PlayerIdentifierUpdateOne {
	if s != nil {
		piuo.SetCreatedWith(*s)
	}
	return piuo
}

// ClearCreatedWith clears the value of the "created_with" field.
func (piuo *PlayerIdentifierUpdateOne) ClearCreatedWith() *PlayerIdentifierUpdateOne {
	piuo.mutation.ClearCreatedWith()
	return piuo
}

// SetUpdatedAt sets the "updated_at" field.
func (piuo *PlayerIdentifierUpdateOne) SetUpdatedAt(t time.Time) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetUpdatedBy sets the "updated_by" field.
func (piuo *PlayerIdentifierUpdateOne) SetUpdatedBy(i int) *PlayerIdentifierUpdateOne {
	piuo.mutation.ResetUpdatedBy()
	piuo.mutation.SetUpdatedBy(i)
	return piuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (piuo *PlayerIdentifierUpdateOne) SetNillableUpdatedBy(i *int) *PlayerIdentifierUpdateOne {
	if i != nil {
		piuo.SetUpdatedBy(*i)
	}
	return piuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (piuo *PlayerIdentifierUpdateOne) AddUpdatedBy(i int) *PlayerIdentifierUpdateOne {
	piuo.mutation.AddUpdatedBy(i)
	return piuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (piuo *PlayerIdentifierUpdateOne) ClearUpdatedBy() *PlayerIdentifierUpdateOne {
	piuo.mutation.ClearUpdatedBy()
	return piuo
}

// SetUpdatedWith sets the "updated_with" field.
func (piuo *PlayerIdentifierUpdateOne) SetUpdatedWith(s string) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetUpdatedWith(s)
	return piuo
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (piuo *PlayerIdentifierUpdateOne) SetNillableUpdatedWith(s *string) *PlayerIdentifierUpdateOne {
	if s != nil {
		piuo.SetUpdatedWith(*s)
	}
	return piuo
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (piuo *PlayerIdentifierUpdateOne) ClearUpdatedWith() *PlayerIdentifierUpdateOne {
	piuo.mutation.ClearUpdatedWith()
	return piuo
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (piuo *PlayerIdentifierUpdateOne) SetPlayerID(id int) *PlayerIdentifierUpdateOne {
	piuo.mutation.SetPlayerID(id)
	return piuo
}

// SetNillablePlayerID sets the "player" edge to the Player entity by ID if the given value is not nil.
func (piuo *PlayerIdentifierUpdateOne) SetNillablePlayerID(id *int) *PlayerIdentifierUpdateOne {
	if id != nil {
		piuo = piuo.SetPlayerID(*id)
	}
	return piuo
}

// SetPlayer sets the "player" edge to the Player entity.
func (piuo *PlayerIdentifierUpdateOne) SetPlayer(p *Player) *PlayerIdentifierUpdateOne {
	return piuo.SetPlayerID(p.ID)
}

// Mutation returns the PlayerIdentifierMutation object of the builder.
func (piuo *PlayerIdentifierUpdateOne) Mutation() *PlayerIdentifierMutation {
	return piuo.mutation
}

// ClearPlayer clears the "player" edge to the Player entity.
func (piuo *PlayerIdentifierUpdateOne) ClearPlayer() *PlayerIdentifierUpdateOne {
	piuo.mutation.ClearPlayer()
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *PlayerIdentifierUpdateOne) Select(field string, fields ...string) *PlayerIdentifierUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated PlayerIdentifier entity.
func (piuo *PlayerIdentifierUpdateOne) Save(ctx context.Context) (*PlayerIdentifier, error) {
	var (
		err  error
		node *PlayerIdentifier
	)
	if err := piuo.defaults(); err != nil {
		return nil, err
	}
	if len(piuo.hooks) == 0 {
		node, err = piuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerIdentifierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piuo.mutation = mutation
			node, err = piuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(piuo.hooks) - 1; i >= 0; i-- {
			if piuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = piuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PlayerIdentifierUpdateOne) SaveX(ctx context.Context) *PlayerIdentifier {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *PlayerIdentifierUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PlayerIdentifierUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piuo *PlayerIdentifierUpdateOne) defaults() error {
	if _, ok := piuo.mutation.UpdatedAt(); !ok {
		if playeridentifier.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized playeridentifier.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := playeridentifier.UpdateDefaultUpdatedAt()
		piuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (piuo *PlayerIdentifierUpdateOne) sqlSave(ctx context.Context) (_node *PlayerIdentifier, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playeridentifier.Table,
			Columns: playeridentifier.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playeridentifier.FieldID,
			},
		},
	}
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PlayerIdentifier.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playeridentifier.FieldID)
		for _, f := range fields {
			if !playeridentifier.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != playeridentifier.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldCreatedBy,
		})
	}
	if value, ok := piuo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldCreatedBy,
		})
	}
	if piuo.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: playeridentifier.FieldCreatedBy,
		})
	}
	if value, ok := piuo.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldCreatedWith,
		})
	}
	if piuo.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playeridentifier.FieldCreatedWith,
		})
	}
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playeridentifier.FieldUpdatedAt,
		})
	}
	if value, ok := piuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldUpdatedBy,
		})
	}
	if value, ok := piuo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playeridentifier.FieldUpdatedBy,
		})
	}
	if piuo.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: playeridentifier.FieldUpdatedBy,
		})
	}
	if value, ok := piuo.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playeridentifier.FieldUpdatedWith,
		})
	}
	if piuo.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: playeridentifier.FieldUpdatedWith,
		})
	}
	if piuo.mutation.PlayerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.PlayerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PlayerIdentifier{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playeridentifier.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
