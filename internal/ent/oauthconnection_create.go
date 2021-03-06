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
	"github.com/responserms/response/internal/ent/oauthconnection"
	"github.com/responserms/response/internal/ent/user"
)

// OAuthConnectionCreate is the builder for creating a OAuthConnection entity.
type OAuthConnectionCreate struct {
	config
	mutation *OAuthConnectionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (occ *OAuthConnectionCreate) SetCreatedAt(t time.Time) *OAuthConnectionCreate {
	occ.mutation.SetCreatedAt(t)
	return occ
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (occ *OAuthConnectionCreate) SetNillableCreatedAt(t *time.Time) *OAuthConnectionCreate {
	if t != nil {
		occ.SetCreatedAt(*t)
	}
	return occ
}

// SetCreatedBy sets the "created_by" field.
func (occ *OAuthConnectionCreate) SetCreatedBy(i int) *OAuthConnectionCreate {
	occ.mutation.SetCreatedBy(i)
	return occ
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (occ *OAuthConnectionCreate) SetNillableCreatedBy(i *int) *OAuthConnectionCreate {
	if i != nil {
		occ.SetCreatedBy(*i)
	}
	return occ
}

// SetCreatedWith sets the "created_with" field.
func (occ *OAuthConnectionCreate) SetCreatedWith(s string) *OAuthConnectionCreate {
	occ.mutation.SetCreatedWith(s)
	return occ
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (occ *OAuthConnectionCreate) SetNillableCreatedWith(s *string) *OAuthConnectionCreate {
	if s != nil {
		occ.SetCreatedWith(*s)
	}
	return occ
}

// SetUpdatedAt sets the "updated_at" field.
func (occ *OAuthConnectionCreate) SetUpdatedAt(t time.Time) *OAuthConnectionCreate {
	occ.mutation.SetUpdatedAt(t)
	return occ
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (occ *OAuthConnectionCreate) SetNillableUpdatedAt(t *time.Time) *OAuthConnectionCreate {
	if t != nil {
		occ.SetUpdatedAt(*t)
	}
	return occ
}

// SetUpdatedBy sets the "updated_by" field.
func (occ *OAuthConnectionCreate) SetUpdatedBy(i int) *OAuthConnectionCreate {
	occ.mutation.SetUpdatedBy(i)
	return occ
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (occ *OAuthConnectionCreate) SetNillableUpdatedBy(i *int) *OAuthConnectionCreate {
	if i != nil {
		occ.SetUpdatedBy(*i)
	}
	return occ
}

// SetUpdatedWith sets the "updated_with" field.
func (occ *OAuthConnectionCreate) SetUpdatedWith(s string) *OAuthConnectionCreate {
	occ.mutation.SetUpdatedWith(s)
	return occ
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (occ *OAuthConnectionCreate) SetNillableUpdatedWith(s *string) *OAuthConnectionCreate {
	if s != nil {
		occ.SetUpdatedWith(*s)
	}
	return occ
}

// SetProvider sets the "provider" field.
func (occ *OAuthConnectionCreate) SetProvider(s string) *OAuthConnectionCreate {
	occ.mutation.SetProvider(s)
	return occ
}

// SetProviderUserID sets the "provider_user_id" field.
func (occ *OAuthConnectionCreate) SetProviderUserID(s string) *OAuthConnectionCreate {
	occ.mutation.SetProviderUserID(s)
	return occ
}

// SetName sets the "name" field.
func (occ *OAuthConnectionCreate) SetName(s string) *OAuthConnectionCreate {
	occ.mutation.SetName(s)
	return occ
}

// SetUserID sets the "user" edge to the User entity by ID.
func (occ *OAuthConnectionCreate) SetUserID(id int) *OAuthConnectionCreate {
	occ.mutation.SetUserID(id)
	return occ
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (occ *OAuthConnectionCreate) SetNillableUserID(id *int) *OAuthConnectionCreate {
	if id != nil {
		occ = occ.SetUserID(*id)
	}
	return occ
}

// SetUser sets the "user" edge to the User entity.
func (occ *OAuthConnectionCreate) SetUser(u *User) *OAuthConnectionCreate {
	return occ.SetUserID(u.ID)
}

// Mutation returns the OAuthConnectionMutation object of the builder.
func (occ *OAuthConnectionCreate) Mutation() *OAuthConnectionMutation {
	return occ.mutation
}

// Save creates the OAuthConnection in the database.
func (occ *OAuthConnectionCreate) Save(ctx context.Context) (*OAuthConnection, error) {
	var (
		err  error
		node *OAuthConnection
	)
	if err := occ.defaults(); err != nil {
		return nil, err
	}
	if len(occ.hooks) == 0 {
		if err = occ.check(); err != nil {
			return nil, err
		}
		node, err = occ.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuthConnectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = occ.check(); err != nil {
				return nil, err
			}
			occ.mutation = mutation
			if node, err = occ.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(occ.hooks) - 1; i >= 0; i-- {
			if occ.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = occ.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, occ.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (occ *OAuthConnectionCreate) SaveX(ctx context.Context) *OAuthConnection {
	v, err := occ.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occ *OAuthConnectionCreate) Exec(ctx context.Context) error {
	_, err := occ.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occ *OAuthConnectionCreate) ExecX(ctx context.Context) {
	if err := occ.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (occ *OAuthConnectionCreate) defaults() error {
	if _, ok := occ.mutation.CreatedAt(); !ok {
		if oauthconnection.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized oauthconnection.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := oauthconnection.DefaultCreatedAt()
		occ.mutation.SetCreatedAt(v)
	}
	if _, ok := occ.mutation.UpdatedAt(); !ok {
		if oauthconnection.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized oauthconnection.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := oauthconnection.DefaultUpdatedAt()
		occ.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (occ *OAuthConnectionCreate) check() error {
	if _, ok := occ.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := occ.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := occ.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "provider"`)}
	}
	if _, ok := occ.mutation.ProviderUserID(); !ok {
		return &ValidationError{Name: "provider_user_id", err: errors.New(`ent: missing required field "provider_user_id"`)}
	}
	if _, ok := occ.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (occ *OAuthConnectionCreate) sqlSave(ctx context.Context) (*OAuthConnection, error) {
	_node, _spec := occ.createSpec()
	if err := sqlgraph.CreateNode(ctx, occ.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (occ *OAuthConnectionCreate) createSpec() (*OAuthConnection, *sqlgraph.CreateSpec) {
	var (
		_node = &OAuthConnection{config: occ.config}
		_spec = &sqlgraph.CreateSpec{
			Table: oauthconnection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauthconnection.FieldID,
			},
		}
	)
	_spec.OnConflict = occ.conflict
	if value, ok := occ.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauthconnection.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := occ.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oauthconnection.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := occ.mutation.CreatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauthconnection.FieldCreatedWith,
		})
		_node.CreatedWith = value
	}
	if value, ok := occ.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauthconnection.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := occ.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: oauthconnection.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := occ.mutation.UpdatedWith(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauthconnection.FieldUpdatedWith,
		})
		_node.UpdatedWith = value
	}
	if value, ok := occ.mutation.Provider(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauthconnection.FieldProvider,
		})
		_node.Provider = value
	}
	if value, ok := occ.mutation.ProviderUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauthconnection.FieldProviderUserID,
		})
		_node.ProviderUserID = value
	}
	if value, ok := occ.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauthconnection.FieldName,
		})
		_node.Name = value
	}
	if nodes := occ.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthconnection.UserTable,
			Columns: []string{oauthconnection.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_oauth_connections = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OAuthConnection.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OAuthConnectionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (occ *OAuthConnectionCreate) OnConflict(opts ...sql.ConflictOption) *OAuthConnectionUpsertOne {
	occ.conflict = opts
	return &OAuthConnectionUpsertOne{
		create: occ,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OAuthConnection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (occ *OAuthConnectionCreate) OnConflictColumns(columns ...string) *OAuthConnectionUpsertOne {
	occ.conflict = append(occ.conflict, sql.ConflictColumns(columns...))
	return &OAuthConnectionUpsertOne{
		create: occ,
	}
}

type (
	// OAuthConnectionUpsertOne is the builder for "upsert"-ing
	//  one OAuthConnection node.
	OAuthConnectionUpsertOne struct {
		create *OAuthConnectionCreate
	}

	// OAuthConnectionUpsert is the "OnConflict" setter.
	OAuthConnectionUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *OAuthConnectionUpsert) SetCreatedAt(v time.Time) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateCreatedAt() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldCreatedAt)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *OAuthConnectionUpsert) SetCreatedBy(v int) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateCreatedBy() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *OAuthConnectionUpsert) ClearCreatedBy() *OAuthConnectionUpsert {
	u.SetNull(oauthconnection.FieldCreatedBy)
	return u
}

// SetCreatedWith sets the "created_with" field.
func (u *OAuthConnectionUpsert) SetCreatedWith(v string) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldCreatedWith, v)
	return u
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateCreatedWith() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldCreatedWith)
	return u
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *OAuthConnectionUpsert) ClearCreatedWith() *OAuthConnectionUpsert {
	u.SetNull(oauthconnection.FieldCreatedWith)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OAuthConnectionUpsert) SetUpdatedAt(v time.Time) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateUpdatedAt() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldUpdatedAt)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OAuthConnectionUpsert) SetUpdatedBy(v int) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateUpdatedBy() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldUpdatedBy)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OAuthConnectionUpsert) ClearUpdatedBy() *OAuthConnectionUpsert {
	u.SetNull(oauthconnection.FieldUpdatedBy)
	return u
}

// SetUpdatedWith sets the "updated_with" field.
func (u *OAuthConnectionUpsert) SetUpdatedWith(v string) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldUpdatedWith, v)
	return u
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateUpdatedWith() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldUpdatedWith)
	return u
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *OAuthConnectionUpsert) ClearUpdatedWith() *OAuthConnectionUpsert {
	u.SetNull(oauthconnection.FieldUpdatedWith)
	return u
}

// SetProvider sets the "provider" field.
func (u *OAuthConnectionUpsert) SetProvider(v string) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldProvider, v)
	return u
}

// UpdateProvider sets the "provider" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateProvider() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldProvider)
	return u
}

// SetProviderUserID sets the "provider_user_id" field.
func (u *OAuthConnectionUpsert) SetProviderUserID(v string) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldProviderUserID, v)
	return u
}

// UpdateProviderUserID sets the "provider_user_id" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateProviderUserID() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldProviderUserID)
	return u
}

// SetName sets the "name" field.
func (u *OAuthConnectionUpsert) SetName(v string) *OAuthConnectionUpsert {
	u.Set(oauthconnection.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OAuthConnectionUpsert) UpdateName() *OAuthConnectionUpsert {
	u.SetExcluded(oauthconnection.FieldName)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.OAuthConnection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OAuthConnectionUpsertOne) UpdateNewValues() *OAuthConnectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.OAuthConnection.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *OAuthConnectionUpsertOne) Ignore() *OAuthConnectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OAuthConnectionUpsertOne) DoNothing() *OAuthConnectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OAuthConnectionCreate.OnConflict
// documentation for more info.
func (u *OAuthConnectionUpsertOne) Update(set func(*OAuthConnectionUpsert)) *OAuthConnectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OAuthConnectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OAuthConnectionUpsertOne) SetCreatedAt(v time.Time) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateCreatedAt() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *OAuthConnectionUpsertOne) SetCreatedBy(v int) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateCreatedBy() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *OAuthConnectionUpsertOne) ClearCreatedBy() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *OAuthConnectionUpsertOne) SetCreatedWith(v string) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateCreatedWith() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *OAuthConnectionUpsertOne) ClearCreatedWith() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OAuthConnectionUpsertOne) SetUpdatedAt(v time.Time) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateUpdatedAt() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OAuthConnectionUpsertOne) SetUpdatedBy(v int) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateUpdatedBy() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OAuthConnectionUpsertOne) ClearUpdatedBy() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *OAuthConnectionUpsertOne) SetUpdatedWith(v string) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateUpdatedWith() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *OAuthConnectionUpsertOne) ClearUpdatedWith() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetProvider sets the "provider" field.
func (u *OAuthConnectionUpsertOne) SetProvider(v string) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetProvider(v)
	})
}

// UpdateProvider sets the "provider" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateProvider() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateProvider()
	})
}

// SetProviderUserID sets the "provider_user_id" field.
func (u *OAuthConnectionUpsertOne) SetProviderUserID(v string) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetProviderUserID(v)
	})
}

// UpdateProviderUserID sets the "provider_user_id" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateProviderUserID() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateProviderUserID()
	})
}

// SetName sets the "name" field.
func (u *OAuthConnectionUpsertOne) SetName(v string) *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OAuthConnectionUpsertOne) UpdateName() *OAuthConnectionUpsertOne {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *OAuthConnectionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OAuthConnectionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OAuthConnectionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OAuthConnectionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OAuthConnectionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OAuthConnectionCreateBulk is the builder for creating many OAuthConnection entities in bulk.
type OAuthConnectionCreateBulk struct {
	config
	builders []*OAuthConnectionCreate
	conflict []sql.ConflictOption
}

// Save creates the OAuthConnection entities in the database.
func (occb *OAuthConnectionCreateBulk) Save(ctx context.Context) ([]*OAuthConnection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(occb.builders))
	nodes := make([]*OAuthConnection, len(occb.builders))
	mutators := make([]Mutator, len(occb.builders))
	for i := range occb.builders {
		func(i int, root context.Context) {
			builder := occb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OAuthConnectionMutation)
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
					_, err = mutators[i+1].Mutate(root, occb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = occb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, occb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, occb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (occb *OAuthConnectionCreateBulk) SaveX(ctx context.Context) []*OAuthConnection {
	v, err := occb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occb *OAuthConnectionCreateBulk) Exec(ctx context.Context) error {
	_, err := occb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occb *OAuthConnectionCreateBulk) ExecX(ctx context.Context) {
	if err := occb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OAuthConnection.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OAuthConnectionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (occb *OAuthConnectionCreateBulk) OnConflict(opts ...sql.ConflictOption) *OAuthConnectionUpsertBulk {
	occb.conflict = opts
	return &OAuthConnectionUpsertBulk{
		create: occb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OAuthConnection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (occb *OAuthConnectionCreateBulk) OnConflictColumns(columns ...string) *OAuthConnectionUpsertBulk {
	occb.conflict = append(occb.conflict, sql.ConflictColumns(columns...))
	return &OAuthConnectionUpsertBulk{
		create: occb,
	}
}

// OAuthConnectionUpsertBulk is the builder for "upsert"-ing
// a bulk of OAuthConnection nodes.
type OAuthConnectionUpsertBulk struct {
	create *OAuthConnectionCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OAuthConnection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OAuthConnectionUpsertBulk) UpdateNewValues() *OAuthConnectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OAuthConnection.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *OAuthConnectionUpsertBulk) Ignore() *OAuthConnectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OAuthConnectionUpsertBulk) DoNothing() *OAuthConnectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OAuthConnectionCreateBulk.OnConflict
// documentation for more info.
func (u *OAuthConnectionUpsertBulk) Update(set func(*OAuthConnectionUpsert)) *OAuthConnectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OAuthConnectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OAuthConnectionUpsertBulk) SetCreatedAt(v time.Time) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateCreatedAt() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *OAuthConnectionUpsertBulk) SetCreatedBy(v int) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateCreatedBy() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *OAuthConnectionUpsertBulk) ClearCreatedBy() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearCreatedBy()
	})
}

// SetCreatedWith sets the "created_with" field.
func (u *OAuthConnectionUpsertBulk) SetCreatedWith(v string) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetCreatedWith(v)
	})
}

// UpdateCreatedWith sets the "created_with" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateCreatedWith() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateCreatedWith()
	})
}

// ClearCreatedWith clears the value of the "created_with" field.
func (u *OAuthConnectionUpsertBulk) ClearCreatedWith() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearCreatedWith()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OAuthConnectionUpsertBulk) SetUpdatedAt(v time.Time) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateUpdatedAt() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OAuthConnectionUpsertBulk) SetUpdatedBy(v int) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateUpdatedBy() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OAuthConnectionUpsertBulk) ClearUpdatedBy() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedWith sets the "updated_with" field.
func (u *OAuthConnectionUpsertBulk) SetUpdatedWith(v string) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetUpdatedWith(v)
	})
}

// UpdateUpdatedWith sets the "updated_with" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateUpdatedWith() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateUpdatedWith()
	})
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (u *OAuthConnectionUpsertBulk) ClearUpdatedWith() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.ClearUpdatedWith()
	})
}

// SetProvider sets the "provider" field.
func (u *OAuthConnectionUpsertBulk) SetProvider(v string) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetProvider(v)
	})
}

// UpdateProvider sets the "provider" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateProvider() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateProvider()
	})
}

// SetProviderUserID sets the "provider_user_id" field.
func (u *OAuthConnectionUpsertBulk) SetProviderUserID(v string) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetProviderUserID(v)
	})
}

// UpdateProviderUserID sets the "provider_user_id" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateProviderUserID() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateProviderUserID()
	})
}

// SetName sets the "name" field.
func (u *OAuthConnectionUpsertBulk) SetName(v string) *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OAuthConnectionUpsertBulk) UpdateName() *OAuthConnectionUpsertBulk {
	return u.Update(func(s *OAuthConnectionUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *OAuthConnectionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OAuthConnectionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OAuthConnectionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OAuthConnectionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
