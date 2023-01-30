// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmhafiz/go8/ent/gen/account"
	"github.com/gmhafiz/go8/ent/gen/ekyc"
	"github.com/gmhafiz/go8/ent/gen/predicate"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAccountName sets the "account_name" field.
func (au *AccountUpdate) SetAccountName(s string) *AccountUpdate {
	au.mutation.SetAccountName(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccountUpdate) SetCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// ClearCreatedAt clears the value of the "created_at" field.
func (au *AccountUpdate) ClearCreatedAt() *AccountUpdate {
	au.mutation.ClearCreatedAt()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUpdatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *AccountUpdate) ClearUpdatedAt() *AccountUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AccountUpdate) SetDeletedAt(t time.Time) *AccountUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeletedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *AccountUpdate) ClearDeletedAt() *AccountUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// AddEkycIDs adds the "ekycs" edge to the Ekyc entity by IDs.
func (au *AccountUpdate) AddEkycIDs(ids ...uint) *AccountUpdate {
	au.mutation.AddEkycIDs(ids...)
	return au
}

// AddEkycs adds the "ekycs" edges to the Ekyc entity.
func (au *AccountUpdate) AddEkycs(e ...*Ekyc) *AccountUpdate {
	ids := make([]uint, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.AddEkycIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearEkycs clears all "ekycs" edges to the Ekyc entity.
func (au *AccountUpdate) ClearEkycs() *AccountUpdate {
	au.mutation.ClearEkycs()
	return au
}

// RemoveEkycIDs removes the "ekycs" edge to Ekyc entities by IDs.
func (au *AccountUpdate) RemoveEkycIDs(ids ...uint) *AccountUpdate {
	au.mutation.RemoveEkycIDs(ids...)
	return au
}

// RemoveEkycs removes "ekycs" edges to Ekyc entities.
func (au *AccountUpdate) RemoveEkycs(e ...*Ekyc) *AccountUpdate {
	ids := make([]uint, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.RemoveEkycIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.AccountName(); ok {
		_spec.SetField(account.FieldAccountName, field.TypeString, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.CreatedAtCleared() {
		_spec.ClearField(account.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(account.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(account.FieldDeletedAt, field.TypeTime)
	}
	if au.mutation.EkycsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.EkycsTable,
			Columns: account.EkycsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: ekyc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedEkycsIDs(); len(nodes) > 0 && !au.mutation.EkycsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.EkycsTable,
			Columns: account.EkycsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: ekyc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EkycsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.EkycsTable,
			Columns: account.EkycsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: ekyc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetAccountName sets the "account_name" field.
func (auo *AccountUpdateOne) SetAccountName(s string) *AccountUpdateOne {
	auo.mutation.SetAccountName(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccountUpdateOne) SetCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (auo *AccountUpdateOne) ClearCreatedAt() *AccountUpdateOne {
	auo.mutation.ClearCreatedAt()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUpdatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *AccountUpdateOne) ClearUpdatedAt() *AccountUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AccountUpdateOne) SetDeletedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeletedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *AccountUpdateOne) ClearDeletedAt() *AccountUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// AddEkycIDs adds the "ekycs" edge to the Ekyc entity by IDs.
func (auo *AccountUpdateOne) AddEkycIDs(ids ...uint) *AccountUpdateOne {
	auo.mutation.AddEkycIDs(ids...)
	return auo
}

// AddEkycs adds the "ekycs" edges to the Ekyc entity.
func (auo *AccountUpdateOne) AddEkycs(e ...*Ekyc) *AccountUpdateOne {
	ids := make([]uint, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.AddEkycIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearEkycs clears all "ekycs" edges to the Ekyc entity.
func (auo *AccountUpdateOne) ClearEkycs() *AccountUpdateOne {
	auo.mutation.ClearEkycs()
	return auo
}

// RemoveEkycIDs removes the "ekycs" edge to Ekyc entities by IDs.
func (auo *AccountUpdateOne) RemoveEkycIDs(ids ...uint) *AccountUpdateOne {
	auo.mutation.RemoveEkycIDs(ids...)
	return auo
}

// RemoveEkycs removes "ekycs" edges to Ekyc entities.
func (auo *AccountUpdateOne) RemoveEkycs(e ...*Ekyc) *AccountUpdateOne {
	ids := make([]uint, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.RemoveEkycIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.AccountName(); ok {
		_spec.SetField(account.FieldAccountName, field.TypeString, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.CreatedAtCleared() {
		_spec.ClearField(account.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(account.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(account.FieldDeletedAt, field.TypeTime)
	}
	if auo.mutation.EkycsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.EkycsTable,
			Columns: account.EkycsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: ekyc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedEkycsIDs(); len(nodes) > 0 && !auo.mutation.EkycsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.EkycsTable,
			Columns: account.EkycsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: ekyc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EkycsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.EkycsTable,
			Columns: account.EkycsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: ekyc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
