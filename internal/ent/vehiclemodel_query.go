// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/predicate"
	"github.com/responserms/response/internal/ent/vehicle"
	"github.com/responserms/response/internal/ent/vehiclemodel"
)

// VehicleModelQuery is the builder for querying VehicleModel entities.
type VehicleModelQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.VehicleModel
	// eager-loading edges.
	withMetadata *MetadataQuery
	withVehicles *VehicleQuery
	withFKs      bool
	modifiers    []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VehicleModelQuery builder.
func (vmq *VehicleModelQuery) Where(ps ...predicate.VehicleModel) *VehicleModelQuery {
	vmq.predicates = append(vmq.predicates, ps...)
	return vmq
}

// Limit adds a limit step to the query.
func (vmq *VehicleModelQuery) Limit(limit int) *VehicleModelQuery {
	vmq.limit = &limit
	return vmq
}

// Offset adds an offset step to the query.
func (vmq *VehicleModelQuery) Offset(offset int) *VehicleModelQuery {
	vmq.offset = &offset
	return vmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vmq *VehicleModelQuery) Unique(unique bool) *VehicleModelQuery {
	vmq.unique = &unique
	return vmq
}

// Order adds an order step to the query.
func (vmq *VehicleModelQuery) Order(o ...OrderFunc) *VehicleModelQuery {
	vmq.order = append(vmq.order, o...)
	return vmq
}

// QueryMetadata chains the current query on the "metadata" edge.
func (vmq *VehicleModelQuery) QueryMetadata() *MetadataQuery {
	query := &MetadataQuery{config: vmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vehiclemodel.Table, vehiclemodel.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, vehiclemodel.MetadataTable, vehiclemodel.MetadataColumn),
		)
		fromU = sqlgraph.SetNeighbors(vmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVehicles chains the current query on the "vehicles" edge.
func (vmq *VehicleModelQuery) QueryVehicles() *VehicleQuery {
	query := &VehicleQuery{config: vmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vehiclemodel.Table, vehiclemodel.FieldID, selector),
			sqlgraph.To(vehicle.Table, vehicle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vehiclemodel.VehiclesTable, vehiclemodel.VehiclesColumn),
		)
		fromU = sqlgraph.SetNeighbors(vmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VehicleModel entity from the query.
// Returns a *NotFoundError when no VehicleModel was found.
func (vmq *VehicleModelQuery) First(ctx context.Context) (*VehicleModel, error) {
	nodes, err := vmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vehiclemodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vmq *VehicleModelQuery) FirstX(ctx context.Context) *VehicleModel {
	node, err := vmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VehicleModel ID from the query.
// Returns a *NotFoundError when no VehicleModel ID was found.
func (vmq *VehicleModelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vehiclemodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vmq *VehicleModelQuery) FirstIDX(ctx context.Context) int {
	id, err := vmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VehicleModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one VehicleModel entity is not found.
// Returns a *NotFoundError when no VehicleModel entities are found.
func (vmq *VehicleModelQuery) Only(ctx context.Context) (*VehicleModel, error) {
	nodes, err := vmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vehiclemodel.Label}
	default:
		return nil, &NotSingularError{vehiclemodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vmq *VehicleModelQuery) OnlyX(ctx context.Context) *VehicleModel {
	node, err := vmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VehicleModel ID in the query.
// Returns a *NotSingularError when exactly one VehicleModel ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vmq *VehicleModelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = &NotSingularError{vehiclemodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vmq *VehicleModelQuery) OnlyIDX(ctx context.Context) int {
	id, err := vmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VehicleModels.
func (vmq *VehicleModelQuery) All(ctx context.Context) ([]*VehicleModel, error) {
	if err := vmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vmq *VehicleModelQuery) AllX(ctx context.Context) []*VehicleModel {
	nodes, err := vmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VehicleModel IDs.
func (vmq *VehicleModelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := vmq.Select(vehiclemodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vmq *VehicleModelQuery) IDsX(ctx context.Context) []int {
	ids, err := vmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vmq *VehicleModelQuery) Count(ctx context.Context) (int, error) {
	if err := vmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vmq *VehicleModelQuery) CountX(ctx context.Context) int {
	count, err := vmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vmq *VehicleModelQuery) Exist(ctx context.Context) (bool, error) {
	if err := vmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vmq *VehicleModelQuery) ExistX(ctx context.Context) bool {
	exist, err := vmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VehicleModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vmq *VehicleModelQuery) Clone() *VehicleModelQuery {
	if vmq == nil {
		return nil
	}
	return &VehicleModelQuery{
		config:       vmq.config,
		limit:        vmq.limit,
		offset:       vmq.offset,
		order:        append([]OrderFunc{}, vmq.order...),
		predicates:   append([]predicate.VehicleModel{}, vmq.predicates...),
		withMetadata: vmq.withMetadata.Clone(),
		withVehicles: vmq.withVehicles.Clone(),
		// clone intermediate query.
		sql:  vmq.sql.Clone(),
		path: vmq.path,
	}
}

// WithMetadata tells the query-builder to eager-load the nodes that are connected to
// the "metadata" edge. The optional arguments are used to configure the query builder of the edge.
func (vmq *VehicleModelQuery) WithMetadata(opts ...func(*MetadataQuery)) *VehicleModelQuery {
	query := &MetadataQuery{config: vmq.config}
	for _, opt := range opts {
		opt(query)
	}
	vmq.withMetadata = query
	return vmq
}

// WithVehicles tells the query-builder to eager-load the nodes that are connected to
// the "vehicles" edge. The optional arguments are used to configure the query builder of the edge.
func (vmq *VehicleModelQuery) WithVehicles(opts ...func(*VehicleQuery)) *VehicleModelQuery {
	query := &VehicleQuery{config: vmq.config}
	for _, opt := range opts {
		opt(query)
	}
	vmq.withVehicles = query
	return vmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VehicleModel.Query().
//		GroupBy(vehiclemodel.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vmq *VehicleModelQuery) GroupBy(field string, fields ...string) *VehicleModelGroupBy {
	group := &VehicleModelGroupBy{config: vmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vmq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.VehicleModel.Query().
//		Select(vehiclemodel.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (vmq *VehicleModelQuery) Select(fields ...string) *VehicleModelSelect {
	vmq.fields = append(vmq.fields, fields...)
	return &VehicleModelSelect{VehicleModelQuery: vmq}
}

func (vmq *VehicleModelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vmq.fields {
		if !vehiclemodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vmq.path != nil {
		prev, err := vmq.path(ctx)
		if err != nil {
			return err
		}
		vmq.sql = prev
	}
	return nil
}

func (vmq *VehicleModelQuery) sqlAll(ctx context.Context) ([]*VehicleModel, error) {
	var (
		nodes       = []*VehicleModel{}
		withFKs     = vmq.withFKs
		_spec       = vmq.querySpec()
		loadedTypes = [2]bool{
			vmq.withMetadata != nil,
			vmq.withVehicles != nil,
		}
	)
	if vmq.withMetadata != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, vehiclemodel.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &VehicleModel{config: vmq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vmq.modifiers) > 0 {
		_spec.Modifiers = vmq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, vmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := vmq.withMetadata; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*VehicleModel)
		for i := range nodes {
			if nodes[i].vehicle_model_metadata == nil {
				continue
			}
			fk := *nodes[i].vehicle_model_metadata
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(metadata.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "vehicle_model_metadata" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Metadata = n
			}
		}
	}

	if query := vmq.withVehicles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*VehicleModel)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Vehicles = []*Vehicle{}
		}
		query.withFKs = true
		query.Where(predicate.Vehicle(func(s *sql.Selector) {
			s.Where(sql.InValues(vehiclemodel.VehiclesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.vehicle_model_vehicles
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "vehicle_model_vehicles" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "vehicle_model_vehicles" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Vehicles = append(node.Edges.Vehicles, n)
		}
	}

	return nodes, nil
}

func (vmq *VehicleModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vmq.querySpec()
	if len(vmq.modifiers) > 0 {
		_spec.Modifiers = vmq.modifiers
	}
	return sqlgraph.CountNodes(ctx, vmq.driver, _spec)
}

func (vmq *VehicleModelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vmq *VehicleModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vehiclemodel.Table,
			Columns: vehiclemodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehiclemodel.FieldID,
			},
		},
		From:   vmq.sql,
		Unique: true,
	}
	if unique := vmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehiclemodel.FieldID)
		for i := range fields {
			if fields[i] != vehiclemodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vmq *VehicleModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vmq.driver.Dialect())
	t1 := builder.Table(vehiclemodel.Table)
	columns := vmq.fields
	if len(columns) == 0 {
		columns = vehiclemodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vmq.sql != nil {
		selector = vmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range vmq.modifiers {
		m(selector)
	}
	for _, p := range vmq.predicates {
		p(selector)
	}
	for _, p := range vmq.order {
		p(selector)
	}
	if offset := vmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (vmq *VehicleModelQuery) ForUpdate(opts ...sql.LockOption) *VehicleModelQuery {
	if vmq.driver.Dialect() == dialect.Postgres {
		vmq.Unique(false)
	}
	vmq.modifiers = append(vmq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return vmq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (vmq *VehicleModelQuery) ForShare(opts ...sql.LockOption) *VehicleModelQuery {
	if vmq.driver.Dialect() == dialect.Postgres {
		vmq.Unique(false)
	}
	vmq.modifiers = append(vmq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return vmq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vmq *VehicleModelQuery) Modify(modifiers ...func(s *sql.Selector)) *VehicleModelSelect {
	vmq.modifiers = append(vmq.modifiers, modifiers...)
	return vmq.Select()
}

// VehicleModelGroupBy is the group-by builder for VehicleModel entities.
type VehicleModelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vmgb *VehicleModelGroupBy) Aggregate(fns ...AggregateFunc) *VehicleModelGroupBy {
	vmgb.fns = append(vmgb.fns, fns...)
	return vmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vmgb *VehicleModelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vmgb.path(ctx)
	if err != nil {
		return err
	}
	vmgb.sql = query
	return vmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vmgb.fields) > 1 {
		return nil, errors.New("ent: VehicleModelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) StringsX(ctx context.Context) []string {
	v, err := vmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) StringX(ctx context.Context) string {
	v, err := vmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vmgb.fields) > 1 {
		return nil, errors.New("ent: VehicleModelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) IntsX(ctx context.Context) []int {
	v, err := vmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) IntX(ctx context.Context) int {
	v, err := vmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vmgb.fields) > 1 {
		return nil, errors.New("ent: VehicleModelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vmgb.fields) > 1 {
		return nil, errors.New("ent: VehicleModelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vmgb *VehicleModelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vmgb *VehicleModelGroupBy) BoolX(ctx context.Context) bool {
	v, err := vmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vmgb *VehicleModelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vmgb.fields {
		if !vehiclemodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vmgb *VehicleModelGroupBy) sqlQuery() *sql.Selector {
	selector := vmgb.sql.Select()
	aggregation := make([]string, 0, len(vmgb.fns))
	for _, fn := range vmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vmgb.fields)+len(vmgb.fns))
		for _, f := range vmgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vmgb.fields...)...)
}

// VehicleModelSelect is the builder for selecting fields of VehicleModel entities.
type VehicleModelSelect struct {
	*VehicleModelQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vms *VehicleModelSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vms.prepareQuery(ctx); err != nil {
		return err
	}
	vms.sql = vms.VehicleModelQuery.sqlQuery(ctx)
	return vms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vms *VehicleModelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vms.fields) > 1 {
		return nil, errors.New("ent: VehicleModelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vms *VehicleModelSelect) StringsX(ctx context.Context) []string {
	v, err := vms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vms *VehicleModelSelect) StringX(ctx context.Context) string {
	v, err := vms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vms.fields) > 1 {
		return nil, errors.New("ent: VehicleModelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vms *VehicleModelSelect) IntsX(ctx context.Context) []int {
	v, err := vms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vms *VehicleModelSelect) IntX(ctx context.Context) int {
	v, err := vms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vms.fields) > 1 {
		return nil, errors.New("ent: VehicleModelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vms *VehicleModelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vms *VehicleModelSelect) Float64X(ctx context.Context) float64 {
	v, err := vms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vms.fields) > 1 {
		return nil, errors.New("ent: VehicleModelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vms *VehicleModelSelect) BoolsX(ctx context.Context) []bool {
	v, err := vms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vms *VehicleModelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vehiclemodel.Label}
	default:
		err = fmt.Errorf("ent: VehicleModelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vms *VehicleModelSelect) BoolX(ctx context.Context) bool {
	v, err := vms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vms *VehicleModelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vms.sql.Query()
	if err := vms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vms *VehicleModelSelect) Modify(modifiers ...func(s *sql.Selector)) *VehicleModelSelect {
	vms.modifiers = append(vms.modifiers, modifiers...)
	return vms
}
