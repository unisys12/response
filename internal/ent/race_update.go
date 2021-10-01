// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/person"
	"github.com/responserms/response/internal/ent/predicate"
	"github.com/responserms/response/internal/ent/race"
)

// RaceUpdate is the builder for updating Race entities.
type RaceUpdate struct {
	config
	hooks    []Hook
	mutation *RaceMutation
}

// Where appends a list predicates to the RaceUpdate builder.
func (ru *RaceUpdate) Where(ps ...predicate.Race) *RaceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedBy sets the "created_by" field.
func (ru *RaceUpdate) SetCreatedBy(i int) *RaceUpdate {
	ru.mutation.ResetCreatedBy()
	ru.mutation.SetCreatedBy(i)
	return ru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ru *RaceUpdate) SetNillableCreatedBy(i *int) *RaceUpdate {
	if i != nil {
		ru.SetCreatedBy(*i)
	}
	return ru
}

// AddCreatedBy adds i to the "created_by" field.
func (ru *RaceUpdate) AddCreatedBy(i int) *RaceUpdate {
	ru.mutation.AddCreatedBy(i)
	return ru
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ru *RaceUpdate) ClearCreatedBy() *RaceUpdate {
	ru.mutation.ClearCreatedBy()
	return ru
}

// SetCreatedWith sets the "created_with" field.
func (ru *RaceUpdate) SetCreatedWith(s string) *RaceUpdate {
	ru.mutation.SetCreatedWith(s)
	return ru
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (ru *RaceUpdate) SetNillableCreatedWith(s *string) *RaceUpdate {
	if s != nil {
		ru.SetCreatedWith(*s)
	}
	return ru
}

// ClearCreatedWith clears the value of the "created_with" field.
func (ru *RaceUpdate) ClearCreatedWith() *RaceUpdate {
	ru.mutation.ClearCreatedWith()
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RaceUpdate) SetUpdatedAt(t time.Time) *RaceUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetUpdatedBy sets the "updated_by" field.
func (ru *RaceUpdate) SetUpdatedBy(i int) *RaceUpdate {
	ru.mutation.ResetUpdatedBy()
	ru.mutation.SetUpdatedBy(i)
	return ru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ru *RaceUpdate) SetNillableUpdatedBy(i *int) *RaceUpdate {
	if i != nil {
		ru.SetUpdatedBy(*i)
	}
	return ru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ru *RaceUpdate) AddUpdatedBy(i int) *RaceUpdate {
	ru.mutation.AddUpdatedBy(i)
	return ru
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ru *RaceUpdate) ClearUpdatedBy() *RaceUpdate {
	ru.mutation.ClearUpdatedBy()
	return ru
}

// SetUpdatedWith sets the "updated_with" field.
func (ru *RaceUpdate) SetUpdatedWith(s string) *RaceUpdate {
	ru.mutation.SetUpdatedWith(s)
	return ru
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (ru *RaceUpdate) SetNillableUpdatedWith(s *string) *RaceUpdate {
	if s != nil {
		ru.SetUpdatedWith(*s)
	}
	return ru
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (ru *RaceUpdate) ClearUpdatedWith() *RaceUpdate {
	ru.mutation.ClearUpdatedWith()
	return ru
}

// SetShort sets the "short" field.
func (ru *RaceUpdate) SetShort(s string) *RaceUpdate {
	ru.mutation.SetShort(s)
	return ru
}

// SetTitle sets the "title" field.
func (ru *RaceUpdate) SetTitle(s string) *RaceUpdate {
	ru.mutation.SetTitle(s)
	return ru
}

// SetDescription sets the "description" field.
func (ru *RaceUpdate) SetDescription(s string) *RaceUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RaceUpdate) SetNillableDescription(s *string) *RaceUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// ClearDescription clears the value of the "description" field.
func (ru *RaceUpdate) ClearDescription() *RaceUpdate {
	ru.mutation.ClearDescription()
	return ru
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (ru *RaceUpdate) AddPersonIDs(ids ...int) *RaceUpdate {
	ru.mutation.AddPersonIDs(ids...)
	return ru
}

// AddPeople adds the "people" edges to the Person entity.
func (ru *RaceUpdate) AddPeople(p ...*Person) *RaceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPersonIDs(ids...)
}

// Mutation returns the RaceMutation object of the builder.
func (ru *RaceUpdate) Mutation() *RaceMutation {
	return ru.mutation
}

// ClearPeople clears all "people" edges to the Person entity.
func (ru *RaceUpdate) ClearPeople() *RaceUpdate {
	ru.mutation.ClearPeople()
	return ru
}

// RemovePersonIDs removes the "people" edge to Person entities by IDs.
func (ru *RaceUpdate) RemovePersonIDs(ids ...int) *RaceUpdate {
	ru.mutation.RemovePersonIDs(ids...)
	return ru
}

// RemovePeople removes "people" edges to Person entities.
func (ru *RaceUpdate) RemovePeople(p ...*Person) *RaceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePersonIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RaceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ru.defaults(); err != nil {
		return 0, err
	}
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RaceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RaceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RaceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RaceUpdate) defaults() error {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		if race.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized race.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := race.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ru *RaceUpdate) check() error {
	if v, ok := ru.mutation.Short(); ok {
		if err := race.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf("ent: validator failed for field \"short\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Title(); ok {
		if err := race.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Description(); ok {
		if err := race.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (ru *RaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   race.Table,
			Columns: race.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: race.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldCreatedBy,
		})
	}
	if value, ok := ru.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldCreatedBy,
		})
	}
	if ru.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: race.FieldCreatedBy,
		})
	}
	if value, ok := ru.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldCreatedWith,
		})
	}
	if ru.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: race.FieldCreatedWith,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: race.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldUpdatedBy,
		})
	}
	if value, ok := ru.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldUpdatedBy,
		})
	}
	if ru.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: race.FieldUpdatedBy,
		})
	}
	if value, ok := ru.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldUpdatedWith,
		})
	}
	if ru.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: race.FieldUpdatedWith,
		})
	}
	if value, ok := ru.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldShort,
		})
	}
	if value, ok := ru.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldTitle,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldDescription,
		})
	}
	if ru.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: race.FieldDescription,
		})
	}
	if ru.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.PeopleTable,
			Columns: []string{race.PeopleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPeopleIDs(); len(nodes) > 0 && !ru.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.PeopleTable,
			Columns: []string{race.PeopleColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.PeopleTable,
			Columns: []string{race.PeopleColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{race.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RaceUpdateOne is the builder for updating a single Race entity.
type RaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RaceMutation
}

// SetCreatedBy sets the "created_by" field.
func (ruo *RaceUpdateOne) SetCreatedBy(i int) *RaceUpdateOne {
	ruo.mutation.ResetCreatedBy()
	ruo.mutation.SetCreatedBy(i)
	return ruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ruo *RaceUpdateOne) SetNillableCreatedBy(i *int) *RaceUpdateOne {
	if i != nil {
		ruo.SetCreatedBy(*i)
	}
	return ruo
}

// AddCreatedBy adds i to the "created_by" field.
func (ruo *RaceUpdateOne) AddCreatedBy(i int) *RaceUpdateOne {
	ruo.mutation.AddCreatedBy(i)
	return ruo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ruo *RaceUpdateOne) ClearCreatedBy() *RaceUpdateOne {
	ruo.mutation.ClearCreatedBy()
	return ruo
}

// SetCreatedWith sets the "created_with" field.
func (ruo *RaceUpdateOne) SetCreatedWith(s string) *RaceUpdateOne {
	ruo.mutation.SetCreatedWith(s)
	return ruo
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (ruo *RaceUpdateOne) SetNillableCreatedWith(s *string) *RaceUpdateOne {
	if s != nil {
		ruo.SetCreatedWith(*s)
	}
	return ruo
}

// ClearCreatedWith clears the value of the "created_with" field.
func (ruo *RaceUpdateOne) ClearCreatedWith() *RaceUpdateOne {
	ruo.mutation.ClearCreatedWith()
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RaceUpdateOne) SetUpdatedAt(t time.Time) *RaceUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruo *RaceUpdateOne) SetUpdatedBy(i int) *RaceUpdateOne {
	ruo.mutation.ResetUpdatedBy()
	ruo.mutation.SetUpdatedBy(i)
	return ruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruo *RaceUpdateOne) SetNillableUpdatedBy(i *int) *RaceUpdateOne {
	if i != nil {
		ruo.SetUpdatedBy(*i)
	}
	return ruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ruo *RaceUpdateOne) AddUpdatedBy(i int) *RaceUpdateOne {
	ruo.mutation.AddUpdatedBy(i)
	return ruo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ruo *RaceUpdateOne) ClearUpdatedBy() *RaceUpdateOne {
	ruo.mutation.ClearUpdatedBy()
	return ruo
}

// SetUpdatedWith sets the "updated_with" field.
func (ruo *RaceUpdateOne) SetUpdatedWith(s string) *RaceUpdateOne {
	ruo.mutation.SetUpdatedWith(s)
	return ruo
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (ruo *RaceUpdateOne) SetNillableUpdatedWith(s *string) *RaceUpdateOne {
	if s != nil {
		ruo.SetUpdatedWith(*s)
	}
	return ruo
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (ruo *RaceUpdateOne) ClearUpdatedWith() *RaceUpdateOne {
	ruo.mutation.ClearUpdatedWith()
	return ruo
}

// SetShort sets the "short" field.
func (ruo *RaceUpdateOne) SetShort(s string) *RaceUpdateOne {
	ruo.mutation.SetShort(s)
	return ruo
}

// SetTitle sets the "title" field.
func (ruo *RaceUpdateOne) SetTitle(s string) *RaceUpdateOne {
	ruo.mutation.SetTitle(s)
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RaceUpdateOne) SetDescription(s string) *RaceUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RaceUpdateOne) SetNillableDescription(s *string) *RaceUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// ClearDescription clears the value of the "description" field.
func (ruo *RaceUpdateOne) ClearDescription() *RaceUpdateOne {
	ruo.mutation.ClearDescription()
	return ruo
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (ruo *RaceUpdateOne) AddPersonIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.AddPersonIDs(ids...)
	return ruo
}

// AddPeople adds the "people" edges to the Person entity.
func (ruo *RaceUpdateOne) AddPeople(p ...*Person) *RaceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPersonIDs(ids...)
}

// Mutation returns the RaceMutation object of the builder.
func (ruo *RaceUpdateOne) Mutation() *RaceMutation {
	return ruo.mutation
}

// ClearPeople clears all "people" edges to the Person entity.
func (ruo *RaceUpdateOne) ClearPeople() *RaceUpdateOne {
	ruo.mutation.ClearPeople()
	return ruo
}

// RemovePersonIDs removes the "people" edge to Person entities by IDs.
func (ruo *RaceUpdateOne) RemovePersonIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.RemovePersonIDs(ids...)
	return ruo
}

// RemovePeople removes "people" edges to Person entities.
func (ruo *RaceUpdateOne) RemovePeople(p ...*Person) *RaceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePersonIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RaceUpdateOne) Select(field string, fields ...string) *RaceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Race entity.
func (ruo *RaceUpdateOne) Save(ctx context.Context) (*Race, error) {
	var (
		err  error
		node *Race
	)
	if err := ruo.defaults(); err != nil {
		return nil, err
	}
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RaceUpdateOne) SaveX(ctx context.Context) *Race {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RaceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RaceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RaceUpdateOne) defaults() error {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		if race.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized race.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := race.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RaceUpdateOne) check() error {
	if v, ok := ruo.mutation.Short(); ok {
		if err := race.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf("ent: validator failed for field \"short\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Title(); ok {
		if err := race.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Description(); ok {
		if err := race.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (ruo *RaceUpdateOne) sqlSave(ctx context.Context) (_node *Race, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   race.Table,
			Columns: race.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: race.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Race.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, race.FieldID)
		for _, f := range fields {
			if !race.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != race.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldCreatedBy,
		})
	}
	if value, ok := ruo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldCreatedBy,
		})
	}
	if ruo.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: race.FieldCreatedBy,
		})
	}
	if value, ok := ruo.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldCreatedWith,
		})
	}
	if ruo.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: race.FieldCreatedWith,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: race.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldUpdatedBy,
		})
	}
	if value, ok := ruo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: race.FieldUpdatedBy,
		})
	}
	if ruo.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: race.FieldUpdatedBy,
		})
	}
	if value, ok := ruo.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldUpdatedWith,
		})
	}
	if ruo.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: race.FieldUpdatedWith,
		})
	}
	if value, ok := ruo.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldShort,
		})
	}
	if value, ok := ruo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldTitle,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: race.FieldDescription,
		})
	}
	if ruo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: race.FieldDescription,
		})
	}
	if ruo.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.PeopleTable,
			Columns: []string{race.PeopleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPeopleIDs(); len(nodes) > 0 && !ruo.mutation.PeopleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.PeopleTable,
			Columns: []string{race.PeopleColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.PeopleTable,
			Columns: []string{race.PeopleColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Race{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{race.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}