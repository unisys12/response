// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/gameserver"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/player"
	"github.com/responserms/response/internal/ent/playeridentifier"
	"github.com/responserms/response/internal/ent/predicate"
)

// PlayerUpdate is the builder for updating Player entities.
type PlayerUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerMutation
}

// Where appends a list predicates to the PlayerUpdate builder.
func (pu *PlayerUpdate) Where(ps ...predicate.Player) *PlayerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *PlayerUpdate) SetCreatedBy(i int) *PlayerUpdate {
	pu.mutation.ResetCreatedBy()
	pu.mutation.SetCreatedBy(i)
	return pu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableCreatedBy(i *int) *PlayerUpdate {
	if i != nil {
		pu.SetCreatedBy(*i)
	}
	return pu
}

// AddCreatedBy adds i to the "created_by" field.
func (pu *PlayerUpdate) AddCreatedBy(i int) *PlayerUpdate {
	pu.mutation.AddCreatedBy(i)
	return pu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (pu *PlayerUpdate) ClearCreatedBy() *PlayerUpdate {
	pu.mutation.ClearCreatedBy()
	return pu
}

// SetCreatedWith sets the "created_with" field.
func (pu *PlayerUpdate) SetCreatedWith(s string) *PlayerUpdate {
	pu.mutation.SetCreatedWith(s)
	return pu
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableCreatedWith(s *string) *PlayerUpdate {
	if s != nil {
		pu.SetCreatedWith(*s)
	}
	return pu
}

// ClearCreatedWith clears the value of the "created_with" field.
func (pu *PlayerUpdate) ClearCreatedWith() *PlayerUpdate {
	pu.mutation.ClearCreatedWith()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PlayerUpdate) SetUpdatedAt(t time.Time) *PlayerUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PlayerUpdate) SetUpdatedBy(i int) *PlayerUpdate {
	pu.mutation.ResetUpdatedBy()
	pu.mutation.SetUpdatedBy(i)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableUpdatedBy(i *int) *PlayerUpdate {
	if i != nil {
		pu.SetUpdatedBy(*i)
	}
	return pu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (pu *PlayerUpdate) AddUpdatedBy(i int) *PlayerUpdate {
	pu.mutation.AddUpdatedBy(i)
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *PlayerUpdate) ClearUpdatedBy() *PlayerUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetUpdatedWith sets the "updated_with" field.
func (pu *PlayerUpdate) SetUpdatedWith(s string) *PlayerUpdate {
	pu.mutation.SetUpdatedWith(s)
	return pu
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableUpdatedWith(s *string) *PlayerUpdate {
	if s != nil {
		pu.SetUpdatedWith(*s)
	}
	return pu
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (pu *PlayerUpdate) ClearUpdatedWith() *PlayerUpdate {
	pu.mutation.ClearUpdatedWith()
	return pu
}

// SetName sets the "name" field.
func (pu *PlayerUpdate) SetName(s string) *PlayerUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (pu *PlayerUpdate) SetMetadataID(id int) *PlayerUpdate {
	pu.mutation.SetMetadataID(id)
	return pu
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (pu *PlayerUpdate) SetNillableMetadataID(id *int) *PlayerUpdate {
	if id != nil {
		pu = pu.SetMetadataID(*id)
	}
	return pu
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (pu *PlayerUpdate) SetMetadata(m *Metadata) *PlayerUpdate {
	return pu.SetMetadataID(m.ID)
}

// AddServerIDs adds the "servers" edge to the GameServer entity by IDs.
func (pu *PlayerUpdate) AddServerIDs(ids ...int) *PlayerUpdate {
	pu.mutation.AddServerIDs(ids...)
	return pu
}

// AddServers adds the "servers" edges to the GameServer entity.
func (pu *PlayerUpdate) AddServers(g ...*GameServer) *PlayerUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return pu.AddServerIDs(ids...)
}

// AddIdentifierIDs adds the "identifiers" edge to the PlayerIdentifier entity by IDs.
func (pu *PlayerUpdate) AddIdentifierIDs(ids ...int) *PlayerUpdate {
	pu.mutation.AddIdentifierIDs(ids...)
	return pu
}

// AddIdentifiers adds the "identifiers" edges to the PlayerIdentifier entity.
func (pu *PlayerUpdate) AddIdentifiers(p ...*PlayerIdentifier) *PlayerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddIdentifierIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pu *PlayerUpdate) Mutation() *PlayerMutation {
	return pu.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (pu *PlayerUpdate) ClearMetadata() *PlayerUpdate {
	pu.mutation.ClearMetadata()
	return pu
}

// ClearServers clears all "servers" edges to the GameServer entity.
func (pu *PlayerUpdate) ClearServers() *PlayerUpdate {
	pu.mutation.ClearServers()
	return pu
}

// RemoveServerIDs removes the "servers" edge to GameServer entities by IDs.
func (pu *PlayerUpdate) RemoveServerIDs(ids ...int) *PlayerUpdate {
	pu.mutation.RemoveServerIDs(ids...)
	return pu
}

// RemoveServers removes "servers" edges to GameServer entities.
func (pu *PlayerUpdate) RemoveServers(g ...*GameServer) *PlayerUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return pu.RemoveServerIDs(ids...)
}

// ClearIdentifiers clears all "identifiers" edges to the PlayerIdentifier entity.
func (pu *PlayerUpdate) ClearIdentifiers() *PlayerUpdate {
	pu.mutation.ClearIdentifiers()
	return pu
}

// RemoveIdentifierIDs removes the "identifiers" edge to PlayerIdentifier entities by IDs.
func (pu *PlayerUpdate) RemoveIdentifierIDs(ids ...int) *PlayerUpdate {
	pu.mutation.RemoveIdentifierIDs(ids...)
	return pu
}

// RemoveIdentifiers removes "identifiers" edges to PlayerIdentifier entities.
func (pu *PlayerUpdate) RemoveIdentifiers(p ...*PlayerIdentifier) *PlayerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveIdentifierIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlayerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlayerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlayerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PlayerUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if player.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized player.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := player.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pu *PlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   player.Table,
			Columns: player.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: player.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldCreatedBy,
		})
	}
	if value, ok := pu.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldCreatedBy,
		})
	}
	if pu.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: player.FieldCreatedBy,
		})
	}
	if value, ok := pu.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldCreatedWith,
		})
	}
	if pu.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: player.FieldCreatedWith,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: player.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldUpdatedBy,
		})
	}
	if value, ok := pu.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldUpdatedBy,
		})
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: player.FieldUpdatedBy,
		})
	}
	if value, ok := pu.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldUpdatedWith,
		})
	}
	if pu.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: player.FieldUpdatedWith,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldName,
		})
	}
	if pu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   player.MetadataTable,
			Columns: []string{player.MetadataColumn},
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
	if nodes := pu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   player.MetadataTable,
			Columns: []string{player.MetadataColumn},
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
	if pu.mutation.ServersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.ServersTable,
			Columns: player.ServersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameserver.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedServersIDs(); len(nodes) > 0 && !pu.mutation.ServersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.ServersTable,
			Columns: player.ServersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameserver.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ServersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.ServersTable,
			Columns: player.ServersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameserver.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.IdentifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.IdentifiersTable,
			Columns: []string{player.IdentifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playeridentifier.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedIdentifiersIDs(); len(nodes) > 0 && !pu.mutation.IdentifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.IdentifiersTable,
			Columns: []string{player.IdentifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playeridentifier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.IdentifiersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.IdentifiersTable,
			Columns: []string{player.IdentifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playeridentifier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PlayerUpdateOne is the builder for updating a single Player entity.
type PlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerMutation
}

// SetCreatedBy sets the "created_by" field.
func (puo *PlayerUpdateOne) SetCreatedBy(i int) *PlayerUpdateOne {
	puo.mutation.ResetCreatedBy()
	puo.mutation.SetCreatedBy(i)
	return puo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableCreatedBy(i *int) *PlayerUpdateOne {
	if i != nil {
		puo.SetCreatedBy(*i)
	}
	return puo
}

// AddCreatedBy adds i to the "created_by" field.
func (puo *PlayerUpdateOne) AddCreatedBy(i int) *PlayerUpdateOne {
	puo.mutation.AddCreatedBy(i)
	return puo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (puo *PlayerUpdateOne) ClearCreatedBy() *PlayerUpdateOne {
	puo.mutation.ClearCreatedBy()
	return puo
}

// SetCreatedWith sets the "created_with" field.
func (puo *PlayerUpdateOne) SetCreatedWith(s string) *PlayerUpdateOne {
	puo.mutation.SetCreatedWith(s)
	return puo
}

// SetNillableCreatedWith sets the "created_with" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableCreatedWith(s *string) *PlayerUpdateOne {
	if s != nil {
		puo.SetCreatedWith(*s)
	}
	return puo
}

// ClearCreatedWith clears the value of the "created_with" field.
func (puo *PlayerUpdateOne) ClearCreatedWith() *PlayerUpdateOne {
	puo.mutation.ClearCreatedWith()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PlayerUpdateOne) SetUpdatedAt(t time.Time) *PlayerUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PlayerUpdateOne) SetUpdatedBy(i int) *PlayerUpdateOne {
	puo.mutation.ResetUpdatedBy()
	puo.mutation.SetUpdatedBy(i)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableUpdatedBy(i *int) *PlayerUpdateOne {
	if i != nil {
		puo.SetUpdatedBy(*i)
	}
	return puo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (puo *PlayerUpdateOne) AddUpdatedBy(i int) *PlayerUpdateOne {
	puo.mutation.AddUpdatedBy(i)
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *PlayerUpdateOne) ClearUpdatedBy() *PlayerUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetUpdatedWith sets the "updated_with" field.
func (puo *PlayerUpdateOne) SetUpdatedWith(s string) *PlayerUpdateOne {
	puo.mutation.SetUpdatedWith(s)
	return puo
}

// SetNillableUpdatedWith sets the "updated_with" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableUpdatedWith(s *string) *PlayerUpdateOne {
	if s != nil {
		puo.SetUpdatedWith(*s)
	}
	return puo
}

// ClearUpdatedWith clears the value of the "updated_with" field.
func (puo *PlayerUpdateOne) ClearUpdatedWith() *PlayerUpdateOne {
	puo.mutation.ClearUpdatedWith()
	return puo
}

// SetName sets the "name" field.
func (puo *PlayerUpdateOne) SetName(s string) *PlayerUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (puo *PlayerUpdateOne) SetMetadataID(id int) *PlayerUpdateOne {
	puo.mutation.SetMetadataID(id)
	return puo
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableMetadataID(id *int) *PlayerUpdateOne {
	if id != nil {
		puo = puo.SetMetadataID(*id)
	}
	return puo
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (puo *PlayerUpdateOne) SetMetadata(m *Metadata) *PlayerUpdateOne {
	return puo.SetMetadataID(m.ID)
}

// AddServerIDs adds the "servers" edge to the GameServer entity by IDs.
func (puo *PlayerUpdateOne) AddServerIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.AddServerIDs(ids...)
	return puo
}

// AddServers adds the "servers" edges to the GameServer entity.
func (puo *PlayerUpdateOne) AddServers(g ...*GameServer) *PlayerUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return puo.AddServerIDs(ids...)
}

// AddIdentifierIDs adds the "identifiers" edge to the PlayerIdentifier entity by IDs.
func (puo *PlayerUpdateOne) AddIdentifierIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.AddIdentifierIDs(ids...)
	return puo
}

// AddIdentifiers adds the "identifiers" edges to the PlayerIdentifier entity.
func (puo *PlayerUpdateOne) AddIdentifiers(p ...*PlayerIdentifier) *PlayerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddIdentifierIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (puo *PlayerUpdateOne) Mutation() *PlayerMutation {
	return puo.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (puo *PlayerUpdateOne) ClearMetadata() *PlayerUpdateOne {
	puo.mutation.ClearMetadata()
	return puo
}

// ClearServers clears all "servers" edges to the GameServer entity.
func (puo *PlayerUpdateOne) ClearServers() *PlayerUpdateOne {
	puo.mutation.ClearServers()
	return puo
}

// RemoveServerIDs removes the "servers" edge to GameServer entities by IDs.
func (puo *PlayerUpdateOne) RemoveServerIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.RemoveServerIDs(ids...)
	return puo
}

// RemoveServers removes "servers" edges to GameServer entities.
func (puo *PlayerUpdateOne) RemoveServers(g ...*GameServer) *PlayerUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return puo.RemoveServerIDs(ids...)
}

// ClearIdentifiers clears all "identifiers" edges to the PlayerIdentifier entity.
func (puo *PlayerUpdateOne) ClearIdentifiers() *PlayerUpdateOne {
	puo.mutation.ClearIdentifiers()
	return puo
}

// RemoveIdentifierIDs removes the "identifiers" edge to PlayerIdentifier entities by IDs.
func (puo *PlayerUpdateOne) RemoveIdentifierIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.RemoveIdentifierIDs(ids...)
	return puo
}

// RemoveIdentifiers removes "identifiers" edges to PlayerIdentifier entities.
func (puo *PlayerUpdateOne) RemoveIdentifiers(p ...*PlayerIdentifier) *PlayerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveIdentifierIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlayerUpdateOne) Select(field string, fields ...string) *PlayerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Player entity.
func (puo *PlayerUpdateOne) Save(ctx context.Context) (*Player, error) {
	var (
		err  error
		node *Player
	)
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlayerUpdateOne) SaveX(ctx context.Context) *Player {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlayerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PlayerUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if player.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized player.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := player.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (puo *PlayerUpdateOne) sqlSave(ctx context.Context) (_node *Player, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   player.Table,
			Columns: player.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: player.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Player.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, player.FieldID)
		for _, f := range fields {
			if !player.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != player.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldCreatedBy,
		})
	}
	if value, ok := puo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldCreatedBy,
		})
	}
	if puo.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: player.FieldCreatedBy,
		})
	}
	if value, ok := puo.mutation.CreatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldCreatedWith,
		})
	}
	if puo.mutation.CreatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: player.FieldCreatedWith,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: player.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldUpdatedBy,
		})
	}
	if value, ok := puo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: player.FieldUpdatedBy,
		})
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: player.FieldUpdatedBy,
		})
	}
	if value, ok := puo.mutation.UpdatedWith(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldUpdatedWith,
		})
	}
	if puo.mutation.UpdatedWithCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: player.FieldUpdatedWith,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldName,
		})
	}
	if puo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   player.MetadataTable,
			Columns: []string{player.MetadataColumn},
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
	if nodes := puo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   player.MetadataTable,
			Columns: []string{player.MetadataColumn},
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
	if puo.mutation.ServersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.ServersTable,
			Columns: player.ServersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameserver.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedServersIDs(); len(nodes) > 0 && !puo.mutation.ServersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.ServersTable,
			Columns: player.ServersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameserver.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ServersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.ServersTable,
			Columns: player.ServersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gameserver.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.IdentifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.IdentifiersTable,
			Columns: []string{player.IdentifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playeridentifier.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedIdentifiersIDs(); len(nodes) > 0 && !puo.mutation.IdentifiersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.IdentifiersTable,
			Columns: []string{player.IdentifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playeridentifier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.IdentifiersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.IdentifiersTable,
			Columns: []string{player.IdentifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playeridentifier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Player{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
