// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmhafiz/go8/ent/gen/author"
	"github.com/gmhafiz/go8/ent/gen/book"
	"github.com/gmhafiz/go8/ent/gen/predicate"
)

// AuthorQuery is the builder for querying Author entities.
type AuthorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Author
	withBooks  *BookQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthorQuery builder.
func (aq *AuthorQuery) Where(ps ...predicate.Author) *AuthorQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AuthorQuery) Limit(limit int) *AuthorQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AuthorQuery) Offset(offset int) *AuthorQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AuthorQuery) Unique(unique bool) *AuthorQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AuthorQuery) Order(o ...OrderFunc) *AuthorQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryBooks chains the current query on the "books" edge.
func (aq *AuthorQuery) QueryBooks() *BookQuery {
	query := &BookQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(author.Table, author.FieldID, selector),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, author.BooksTable, author.BooksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Author entity from the query.
// Returns a *NotFoundError when no Author was found.
func (aq *AuthorQuery) First(ctx context.Context) (*Author, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{author.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AuthorQuery) FirstX(ctx context.Context) *Author {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Author ID from the query.
// Returns a *NotFoundError when no Author ID was found.
func (aq *AuthorQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{author.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AuthorQuery) FirstIDX(ctx context.Context) uint {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Author entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Author entity is found.
// Returns a *NotFoundError when no Author entities are found.
func (aq *AuthorQuery) Only(ctx context.Context) (*Author, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{author.Label}
	default:
		return nil, &NotSingularError{author.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AuthorQuery) OnlyX(ctx context.Context) *Author {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Author ID in the query.
// Returns a *NotSingularError when more than one Author ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AuthorQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{author.Label}
	default:
		err = &NotSingularError{author.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AuthorQuery) OnlyIDX(ctx context.Context) uint {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Authors.
func (aq *AuthorQuery) All(ctx context.Context) ([]*Author, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AuthorQuery) AllX(ctx context.Context) []*Author {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Author IDs.
func (aq *AuthorQuery) IDs(ctx context.Context) ([]uint, error) {
	var ids []uint
	if err := aq.Select(author.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AuthorQuery) IDsX(ctx context.Context) []uint {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AuthorQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AuthorQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AuthorQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AuthorQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AuthorQuery) Clone() *AuthorQuery {
	if aq == nil {
		return nil
	}
	return &AuthorQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		predicates: append([]predicate.Author{}, aq.predicates...),
		withBooks:  aq.withBooks.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithBooks tells the query-builder to eager-load the nodes that are connected to
// the "books" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AuthorQuery) WithBooks(opts ...func(*BookQuery)) *AuthorQuery {
	query := &BookQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withBooks = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"first_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Author.Query().
//		GroupBy(author.FieldFirstName).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (aq *AuthorQuery) GroupBy(field string, fields ...string) *AuthorGroupBy {
	grbuild := &AuthorGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = author.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"first_name,omitempty"`
//	}
//
//	client.Author.Query().
//		Select(author.FieldFirstName).
//		Scan(ctx, &v)
func (aq *AuthorQuery) Select(fields ...string) *AuthorSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AuthorSelect{AuthorQuery: aq}
	selbuild.label = author.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a AuthorSelect configured with the given aggregations.
func (aq *AuthorQuery) Aggregate(fns ...AggregateFunc) *AuthorSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AuthorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !author.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AuthorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Author, error) {
	var (
		nodes       = []*Author{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withBooks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Author).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Author{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withBooks; query != nil {
		if err := aq.loadBooks(ctx, query, nodes,
			func(n *Author) { n.Edges.Books = []*Book{} },
			func(n *Author, e *Book) { n.Edges.Books = append(n.Edges.Books, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AuthorQuery) loadBooks(ctx context.Context, query *BookQuery, nodes []*Author, init func(*Author), assign func(*Author, *Book)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint]*Author)
	nids := make(map[uint]map[*Author]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(author.BooksTable)
		s.Join(joinT).On(s.C(book.FieldID), joinT.C(author.BooksPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(author.BooksPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(author.BooksPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := uint(values[0].(*sql.NullInt64).Int64)
			inValue := uint(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Author]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "books" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aq *AuthorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AuthorQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AuthorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   author.Table,
			Columns: author.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: author.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, author.FieldID)
		for i := range fields {
			if fields[i] != author.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AuthorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(author.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = author.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthorGroupBy is the group-by builder for Author entities.
type AuthorGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AuthorGroupBy) Aggregate(fns ...AggregateFunc) *AuthorGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AuthorGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AuthorGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !author.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AuthorGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AuthorSelect is the builder for selecting fields of Author entities.
type AuthorSelect struct {
	*AuthorQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AuthorSelect) Aggregate(fns ...AggregateFunc) *AuthorSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AuthorSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AuthorQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AuthorSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(as.sql))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		as.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		as.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}