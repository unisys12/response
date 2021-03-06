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
	"github.com/responserms/response/internal/ent/deathcertificate"
	"github.com/responserms/response/internal/ent/deathplace"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/predicate"
)

// DeathPlaceQuery is the builder for querying DeathPlace entities.
type DeathPlaceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeathPlace
	// eager-loading edges.
	withMetadata          *MetadataQuery
	withDeathCertificates *DeathCertificateQuery
	withFKs               bool
	modifiers             []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeathPlaceQuery builder.
func (dpq *DeathPlaceQuery) Where(ps ...predicate.DeathPlace) *DeathPlaceQuery {
	dpq.predicates = append(dpq.predicates, ps...)
	return dpq
}

// Limit adds a limit step to the query.
func (dpq *DeathPlaceQuery) Limit(limit int) *DeathPlaceQuery {
	dpq.limit = &limit
	return dpq
}

// Offset adds an offset step to the query.
func (dpq *DeathPlaceQuery) Offset(offset int) *DeathPlaceQuery {
	dpq.offset = &offset
	return dpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dpq *DeathPlaceQuery) Unique(unique bool) *DeathPlaceQuery {
	dpq.unique = &unique
	return dpq
}

// Order adds an order step to the query.
func (dpq *DeathPlaceQuery) Order(o ...OrderFunc) *DeathPlaceQuery {
	dpq.order = append(dpq.order, o...)
	return dpq
}

// QueryMetadata chains the current query on the "metadata" edge.
func (dpq *DeathPlaceQuery) QueryMetadata() *MetadataQuery {
	query := &MetadataQuery{config: dpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deathplace.Table, deathplace.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, deathplace.MetadataTable, deathplace.MetadataColumn),
		)
		fromU = sqlgraph.SetNeighbors(dpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDeathCertificates chains the current query on the "death_certificates" edge.
func (dpq *DeathPlaceQuery) QueryDeathCertificates() *DeathCertificateQuery {
	query := &DeathCertificateQuery{config: dpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deathplace.Table, deathplace.FieldID, selector),
			sqlgraph.To(deathcertificate.Table, deathcertificate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deathplace.DeathCertificatesTable, deathplace.DeathCertificatesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeathPlace entity from the query.
// Returns a *NotFoundError when no DeathPlace was found.
func (dpq *DeathPlaceQuery) First(ctx context.Context) (*DeathPlace, error) {
	nodes, err := dpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deathplace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dpq *DeathPlaceQuery) FirstX(ctx context.Context) *DeathPlace {
	node, err := dpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeathPlace ID from the query.
// Returns a *NotFoundError when no DeathPlace ID was found.
func (dpq *DeathPlaceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deathplace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dpq *DeathPlaceQuery) FirstIDX(ctx context.Context) int {
	id, err := dpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeathPlace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DeathPlace entity is not found.
// Returns a *NotFoundError when no DeathPlace entities are found.
func (dpq *DeathPlaceQuery) Only(ctx context.Context) (*DeathPlace, error) {
	nodes, err := dpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deathplace.Label}
	default:
		return nil, &NotSingularError{deathplace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dpq *DeathPlaceQuery) OnlyX(ctx context.Context) *DeathPlace {
	node, err := dpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeathPlace ID in the query.
// Returns a *NotSingularError when exactly one DeathPlace ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dpq *DeathPlaceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = &NotSingularError{deathplace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dpq *DeathPlaceQuery) OnlyIDX(ctx context.Context) int {
	id, err := dpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeathPlaces.
func (dpq *DeathPlaceQuery) All(ctx context.Context) ([]*DeathPlace, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dpq *DeathPlaceQuery) AllX(ctx context.Context) []*DeathPlace {
	nodes, err := dpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeathPlace IDs.
func (dpq *DeathPlaceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dpq.Select(deathplace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dpq *DeathPlaceQuery) IDsX(ctx context.Context) []int {
	ids, err := dpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dpq *DeathPlaceQuery) Count(ctx context.Context) (int, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dpq *DeathPlaceQuery) CountX(ctx context.Context) int {
	count, err := dpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dpq *DeathPlaceQuery) Exist(ctx context.Context) (bool, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dpq *DeathPlaceQuery) ExistX(ctx context.Context) bool {
	exist, err := dpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeathPlaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dpq *DeathPlaceQuery) Clone() *DeathPlaceQuery {
	if dpq == nil {
		return nil
	}
	return &DeathPlaceQuery{
		config:                dpq.config,
		limit:                 dpq.limit,
		offset:                dpq.offset,
		order:                 append([]OrderFunc{}, dpq.order...),
		predicates:            append([]predicate.DeathPlace{}, dpq.predicates...),
		withMetadata:          dpq.withMetadata.Clone(),
		withDeathCertificates: dpq.withDeathCertificates.Clone(),
		// clone intermediate query.
		sql:  dpq.sql.Clone(),
		path: dpq.path,
	}
}

// WithMetadata tells the query-builder to eager-load the nodes that are connected to
// the "metadata" edge. The optional arguments are used to configure the query builder of the edge.
func (dpq *DeathPlaceQuery) WithMetadata(opts ...func(*MetadataQuery)) *DeathPlaceQuery {
	query := &MetadataQuery{config: dpq.config}
	for _, opt := range opts {
		opt(query)
	}
	dpq.withMetadata = query
	return dpq
}

// WithDeathCertificates tells the query-builder to eager-load the nodes that are connected to
// the "death_certificates" edge. The optional arguments are used to configure the query builder of the edge.
func (dpq *DeathPlaceQuery) WithDeathCertificates(opts ...func(*DeathCertificateQuery)) *DeathPlaceQuery {
	query := &DeathCertificateQuery{config: dpq.config}
	for _, opt := range opts {
		opt(query)
	}
	dpq.withDeathCertificates = query
	return dpq
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
//	client.DeathPlace.Query().
//		GroupBy(deathplace.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dpq *DeathPlaceQuery) GroupBy(field string, fields ...string) *DeathPlaceGroupBy {
	group := &DeathPlaceGroupBy{config: dpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dpq.sqlQuery(ctx), nil
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
//	client.DeathPlace.Query().
//		Select(deathplace.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (dpq *DeathPlaceQuery) Select(fields ...string) *DeathPlaceSelect {
	dpq.fields = append(dpq.fields, fields...)
	return &DeathPlaceSelect{DeathPlaceQuery: dpq}
}

func (dpq *DeathPlaceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dpq.fields {
		if !deathplace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dpq.path != nil {
		prev, err := dpq.path(ctx)
		if err != nil {
			return err
		}
		dpq.sql = prev
	}
	return nil
}

func (dpq *DeathPlaceQuery) sqlAll(ctx context.Context) ([]*DeathPlace, error) {
	var (
		nodes       = []*DeathPlace{}
		withFKs     = dpq.withFKs
		_spec       = dpq.querySpec()
		loadedTypes = [2]bool{
			dpq.withMetadata != nil,
			dpq.withDeathCertificates != nil,
		}
	)
	if dpq.withMetadata != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, deathplace.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DeathPlace{config: dpq.config}
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
	if len(dpq.modifiers) > 0 {
		_spec.Modifiers = dpq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, dpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dpq.withMetadata; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DeathPlace)
		for i := range nodes {
			if nodes[i].death_place_metadata == nil {
				continue
			}
			fk := *nodes[i].death_place_metadata
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
				return nil, fmt.Errorf(`unexpected foreign-key "death_place_metadata" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Metadata = n
			}
		}
	}

	if query := dpq.withDeathCertificates; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*DeathPlace)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DeathCertificates = []*DeathCertificate{}
		}
		query.withFKs = true
		query.Where(predicate.DeathCertificate(func(s *sql.Selector) {
			s.Where(sql.InValues(deathplace.DeathCertificatesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.death_place_death_certificates
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "death_place_death_certificates" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "death_place_death_certificates" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DeathCertificates = append(node.Edges.DeathCertificates, n)
		}
	}

	return nodes, nil
}

func (dpq *DeathPlaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dpq.querySpec()
	if len(dpq.modifiers) > 0 {
		_spec.Modifiers = dpq.modifiers
	}
	return sqlgraph.CountNodes(ctx, dpq.driver, _spec)
}

func (dpq *DeathPlaceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dpq *DeathPlaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deathplace.Table,
			Columns: deathplace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deathplace.FieldID,
			},
		},
		From:   dpq.sql,
		Unique: true,
	}
	if unique := dpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deathplace.FieldID)
		for i := range fields {
			if fields[i] != deathplace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dpq *DeathPlaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dpq.driver.Dialect())
	t1 := builder.Table(deathplace.Table)
	columns := dpq.fields
	if len(columns) == 0 {
		columns = deathplace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dpq.sql != nil {
		selector = dpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range dpq.modifiers {
		m(selector)
	}
	for _, p := range dpq.predicates {
		p(selector)
	}
	for _, p := range dpq.order {
		p(selector)
	}
	if offset := dpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dpq *DeathPlaceQuery) ForUpdate(opts ...sql.LockOption) *DeathPlaceQuery {
	if dpq.driver.Dialect() == dialect.Postgres {
		dpq.Unique(false)
	}
	dpq.modifiers = append(dpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dpq *DeathPlaceQuery) ForShare(opts ...sql.LockOption) *DeathPlaceQuery {
	if dpq.driver.Dialect() == dialect.Postgres {
		dpq.Unique(false)
	}
	dpq.modifiers = append(dpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dpq *DeathPlaceQuery) Modify(modifiers ...func(s *sql.Selector)) *DeathPlaceSelect {
	dpq.modifiers = append(dpq.modifiers, modifiers...)
	return dpq.Select()
}

// DeathPlaceGroupBy is the group-by builder for DeathPlace entities.
type DeathPlaceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dpgb *DeathPlaceGroupBy) Aggregate(fns ...AggregateFunc) *DeathPlaceGroupBy {
	dpgb.fns = append(dpgb.fns, fns...)
	return dpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dpgb *DeathPlaceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dpgb.path(ctx)
	if err != nil {
		return err
	}
	dpgb.sql = query
	return dpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) StringsX(ctx context.Context) []string {
	v, err := dpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) StringX(ctx context.Context) string {
	v, err := dpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) IntsX(ctx context.Context) []int {
	v, err := dpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) IntX(ctx context.Context) int {
	v, err := dpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DeathPlaceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dpgb *DeathPlaceGroupBy) BoolX(ctx context.Context) bool {
	v, err := dpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dpgb *DeathPlaceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dpgb.fields {
		if !deathplace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dpgb *DeathPlaceGroupBy) sqlQuery() *sql.Selector {
	selector := dpgb.sql.Select()
	aggregation := make([]string, 0, len(dpgb.fns))
	for _, fn := range dpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dpgb.fields)+len(dpgb.fns))
		for _, f := range dpgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dpgb.fields...)...)
}

// DeathPlaceSelect is the builder for selecting fields of DeathPlace entities.
type DeathPlaceSelect struct {
	*DeathPlaceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dps *DeathPlaceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dps.prepareQuery(ctx); err != nil {
		return err
	}
	dps.sql = dps.DeathPlaceQuery.sqlQuery(ctx)
	return dps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dps *DeathPlaceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dps *DeathPlaceSelect) StringsX(ctx context.Context) []string {
	v, err := dps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dps *DeathPlaceSelect) StringX(ctx context.Context) string {
	v, err := dps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dps *DeathPlaceSelect) IntsX(ctx context.Context) []int {
	v, err := dps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dps *DeathPlaceSelect) IntX(ctx context.Context) int {
	v, err := dps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dps *DeathPlaceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dps *DeathPlaceSelect) Float64X(ctx context.Context) float64 {
	v, err := dps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DeathPlaceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dps *DeathPlaceSelect) BoolsX(ctx context.Context) []bool {
	v, err := dps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dps *DeathPlaceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deathplace.Label}
	default:
		err = fmt.Errorf("ent: DeathPlaceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dps *DeathPlaceSelect) BoolX(ctx context.Context) bool {
	v, err := dps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dps *DeathPlaceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dps.sql.Query()
	if err := dps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dps *DeathPlaceSelect) Modify(modifiers ...func(s *sql.Selector)) *DeathPlaceSelect {
	dps.modifiers = append(dps.modifiers, modifiers...)
	return dps
}
