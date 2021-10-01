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
	"github.com/responserms/response/internal/ent/sex"
)

// SexCreate is the builder for creating a Sex entity.
type SexCreate struct {
	config
	mutation *SexMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SexCreate) SetCreatedAt(t time.Time) *SexCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SexCreate) SetNillableCreatedAt(t *time.Time) *SexCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *SexCreate) SetCreatedBy(i int) *SexCreate {
	sc.mutation.SetCreatedBy(i)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *SexCreate) SetNillableCreatedBy(i *int) *SexCreate {
	if i != nil {
		sc.SetCreatedBy(*i)
	}
	return sc
}

// SetCreatedWith sets the "created_with" field.
func (sc *SexCreate) SetCreatedWith(s string) *SexCreate {
	sc.mutation.SetCreatedWith(s)
	return sc
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (sc *SexCreate) SetNillableCreatedWith(s *string) *SexCreate {
	if s != nil {
		sc.SetCreatedWith(*s)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SexCreate) SetUpdatedAt(t time.Time) *SexCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SexCreate) SetNillableUpdatedAt(t *time.Time) *SexCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SexCreate) SetUpdatedBy(i int) *SexCreate {
	sc.mutation.SetUpdatedBy(i)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *SexCreate) SetNillableUpdatedBy(i *int) *SexCreate {
	if i != nil {
		sc.SetUpdatedBy(*i)
	}
	return sc
}

// SetUpdatedWith sets the "updated_with" field.
func (sc *SexCreate) SetUpdatedWith(s string) *SexCreate {
	sc.mutation.SetUpdatedWith(s)
	return sc
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (sc *SexCreate) SetNillableUpdatedWith(s *string) *SexCreate {
	if s != nil {
		sc.SetUpdatedWith(*s)
	}
	return sc
}

// SetShort sets the "short" field.
func (sc *SexCreate) SetShort(s string) *SexCreate {
	sc.mutation.SetShort(s)
	return sc
}

// SetTitle sets the "title" field.
func (sc *SexCreate) SetTitle(s string) *SexCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SexCreate) SetDescription(s string) *SexCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SexCreate) SetNillableDescription(s *string) *SexCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (sc *SexCreate) SetMetadataID(id int) *SexCreate {
	sc.mutation.SetMetadataID(id)
	return sc
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (sc *SexCreate) SetNillableMetadataID(id *int) *SexCreate {
	if id != nil {
		sc = sc.SetMetadataID(*id)
	}
	return sc
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (sc *SexCreate) SetMetadata(m *Metadata) *SexCreate {
	return sc.SetMetadataID(m.ID)
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (sc *SexCreate) AddPersonIDs(ids ...int) *SexCreate {
	sc.mutation.AddPersonIDs(ids...)
	return sc
}

// AddPeople adds the "people" edges to the Person entity.
func (sc *SexCreate) AddPeople(p ...*Person) *SexCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPersonIDs(ids...)
}

// Mutation returns the SexMutation object of the builder.
func (sc *SexCreate) Mutation() *SexMutation {
	return sc.mutation
}

// Save creates the Sex in the database.
func (sc *SexCreate) Save(ctx context.Context) (*Sex, error) {
	var (
		err  error
		node *Sex
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
			mutation, ok := m.(*SexMutation)
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
func (sc *SexCreate) SaveX(ctx context.Context) *Sex {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SexCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SexCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SexCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if sex.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized sex.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := sex.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if sex.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized sex.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := sex.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SexCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := sc.mutation.Short(); !ok {
		return &ValidationError{Name: "short", err: errors.New(`ent: missing required field "short"`)}
	}
	if v, ok := sc.mutation.Short(); ok {
		if err := sex.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf(`ent: validator failed for field "short": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := sc.mutation.Title(); ok {
		if err := sex.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Description(); ok {
		if err := sex.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "description": %w`, err)}
		}
	}
	return nil
}

func (sc *SexCreate) sqlSave(ctx context.Context) (*Sex, error) {
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

func (sc *SexCreate) createSpec() (*Sex, *sqlgraph.CreateSpec) {
	var (
		_node = &Sex{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sex.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sex.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sex.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sex.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.CreatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sex.FieldCreatedWith,
		})
		_node.CreatedWith = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sex.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sex.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.UpdatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sex.FieldUpdatedWith,
		})
		_node.UpdatedWith = value
	}
	if value, ok := sc.mutation.Short(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sex.FieldShort,
		})
		_node.Short = value
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sex.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sex.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := sc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sex.MetadataTable,
			Columns: []string{sex.MetadataColumn},
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
		_node.sex_metadata = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sex.PeopleTable,
			Columns: []string{sex.PeopleColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Sex.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SexUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (sc *SexCreate) OnConflict(opts ...sql.ConflictOption) *SexUpsertOne {
	sc.conflict = opts
	return &SexUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Sex.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *SexCreate) OnConflictColumns(columns ...string) *SexUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SexUpsertOne{
		create: sc,
	}
}

type (
	// SexUpsertOne is the builder for "upsert"-ing
	//  one Sex node.
	SexUpsertOne struct {
		create *SexCreate
	}

	// SexUpsert is the "OnConflict" setter.
	SexUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *SexUpsert) SetCreatedAt(v time.Time) *SexUpsert {
	u.Set(sex.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SexUpsert) UpdateCreatedAt() *SexUpsert {
	u.SetExcluded(sex.FieldCreatedAt)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *SexUpsert) SetCreatedBy(v int) *SexUpsert {
	u.Set(sex.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *SexUpsert) UpdateCreatedBy() *SexUpsert {
	u.SetExcluded(sex.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *SexUpsert) ClearCreatedBy() *SexUpsert {
	u.SetNull(sex.FieldCreatedBy)
	return u
}

// SetCreatedWith sets the "created_with" field.
func (u *SexUpsert) SetCreatedWith(v string) *SexUpsert {
	u.Set(sex.FieldCreatedWith, v)
	return u
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *SexUpsert) UpdateCreatedWith() *SexUpsert {
	u.SetExcluded(sex.FieldCreatedWith)
	return u
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *SexUpsert) ClearCreatedWith() *SexUpsert {
	u.SetNull(sex.FieldCreatedWith)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SexUpsert) SetUpdatedAt(v time.Time) *SexUpsert {
	u.Set(sex.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SexUpsert) UpdateUpdatedAt() *SexUpsert {
	u.SetExcluded(sex.FieldUpdatedAt)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SexUpsert) SetUpdatedBy(v int) *SexUpsert {
	u.Set(sex.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SexUpsert) UpdateUpdatedBy() *SexUpsert {
	u.SetExcluded(sex.FieldUpdatedBy)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SexUpsert) ClearUpdatedBy() *SexUpsert {
	u.SetNull(sex.FieldUpdatedBy)
	return u
}

// SetUpdatedWith sets the "updated_with" field.
func (u *SexUpsert) SetUpdatedWith(v string) *SexUpsert {
	u.Set(sex.FieldUpdatedWith, v)
	return u
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *SexUpsert) UpdateUpdatedWith() *SexUpsert {
	u.SetExcluded(sex.FieldUpdatedWith)
	return u
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *SexUpsert) ClearUpdatedWith() *SexUpsert {
	u.SetNull(sex.FieldUpdatedWith)
	return u
}

// SetShort sets the "short" field.
func (u *SexUpsert) SetShort(v string) *SexUpsert {
	u.Set(sex.FieldShort, v)
	return u
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *SexUpsert) UpdateShort() *SexUpsert {
	u.SetExcluded(sex.FieldShort)
	return u
}

// SetTitle sets the "title" field.
func (u *SexUpsert) SetTitle(v string) *SexUpsert {
	u.Set(sex.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *SexUpsert) UpdateTitle() *SexUpsert {
	u.SetExcluded(sex.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *SexUpsert) SetDescription(v string) *SexUpsert {
	u.Set(sex.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SexUpsert) UpdateDescription() *SexUpsert {
	u.SetExcluded(sex.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *SexUpsert) ClearDescription() *SexUpsert {
	u.SetNull(sex.FieldDescription)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Sex.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *SexUpsertOne) UpdateNewValues() *SexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Sex.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *SexUpsertOne) Ignore() *SexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SexUpsertOne) DoNothing() *SexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SexCreate.OnConflict
// documentation for more info.
func (u *SexUpsertOne) Update(set func(*SexUpsert)) *SexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SexUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SexUpsertOne) SetCreatedAt(v time.Time) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateCreatedAt() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *SexUpsertOne) SetCreatedBy(v int) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateCreatedBy() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *SexUpsertOne) ClearCreatedBy() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *SexUpsertOne) SetCreatedWith(v string) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateCreatedWith() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *SexUpsertOne) ClearCreatedWith() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SexUpsertOne) SetUpdatedAt(v time.Time) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateUpdatedAt() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SexUpsertOne) SetUpdatedBy(v int) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateUpdatedBy() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SexUpsertOne) ClearUpdatedBy() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *SexUpsertOne) SetUpdatedWith(v string) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateUpdatedWith() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *SexUpsertOne) ClearUpdatedWith() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetShort sets the "short" field.
func (u *SexUpsertOne) SetShort(v string) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateShort() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateShort()
	})
}

// SetTitle sets the "title" field.
func (u *SexUpsertOne) SetTitle(v string) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateTitle() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *SexUpsertOne) SetDescription(v string) *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SexUpsertOne) UpdateDescription() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *SexUpsertOne) ClearDescription() *SexUpsertOne {
	return u.Update(func(s *SexUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *SexUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SexCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SexUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SexUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SexUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SexCreateBulk is the builder for creating many Sex entities in bulk.
type SexCreateBulk struct {
	config
	builders []*SexCreate
	conflict []sql.ConflictOption
}

// Save creates the Sex entities in the database.
func (scb *SexCreateBulk) Save(ctx context.Context) ([]*Sex, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sex, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SexMutation)
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
func (scb *SexCreateBulk) SaveX(ctx context.Context) []*Sex {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SexCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SexCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Sex.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SexUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (scb *SexCreateBulk) OnConflict(opts ...sql.ConflictOption) *SexUpsertBulk {
	scb.conflict = opts
	return &SexUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Sex.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *SexCreateBulk) OnConflictColumns(columns ...string) *SexUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SexUpsertBulk{
		create: scb,
	}
}

// SexUpsertBulk is the builder for "upsert"-ing
// a bulk of Sex nodes.
type SexUpsertBulk struct {
	create *SexCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Sex.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *SexUpsertBulk) UpdateNewValues() *SexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Sex.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *SexUpsertBulk) Ignore() *SexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SexUpsertBulk) DoNothing() *SexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SexCreateBulk.OnConflict
// documentation for more info.
func (u *SexUpsertBulk) Update(set func(*SexUpsert)) *SexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SexUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SexUpsertBulk) SetCreatedAt(v time.Time) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateCreatedAt() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *SexUpsertBulk) SetCreatedBy(v int) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateCreatedBy() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *SexUpsertBulk) ClearCreatedBy() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *SexUpsertBulk) SetCreatedWith(v string) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateCreatedWith() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *SexUpsertBulk) ClearCreatedWith() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SexUpsertBulk) SetUpdatedAt(v time.Time) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateUpdatedAt() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SexUpsertBulk) SetUpdatedBy(v int) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateUpdatedBy() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SexUpsertBulk) ClearUpdatedBy() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *SexUpsertBulk) SetUpdatedWith(v string) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateUpdatedWith() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *SexUpsertBulk) ClearUpdatedWith() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetShort sets the "short" field.
func (u *SexUpsertBulk) SetShort(v string) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateShort() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateShort()
	})
}

// SetTitle sets the "title" field.
func (u *SexUpsertBulk) SetTitle(v string) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateTitle() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *SexUpsertBulk) SetDescription(v string) *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SexUpsertBulk) UpdateDescription() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *SexUpsertBulk) ClearDescription() *SexUpsertBulk {
	return u.Update(func(s *SexUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *SexUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SexCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SexCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SexUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
