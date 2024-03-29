// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Willsem/health-metrics/internal/generated/ent/access"
	"github.com/Willsem/health-metrics/internal/generated/ent/metric"
	"github.com/Willsem/health-metrics/internal/generated/ent/predicate"
	"github.com/Willsem/health-metrics/internal/generated/ent/user"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetLogin sets the "login" field.
func (uu *UserUpdate) SetLogin(s string) *UserUpdate {
	uu.mutation.SetLogin(s)
	return uu
}

// SetNillableLogin sets the "login" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLogin(s *string) *UserUpdate {
	if s != nil {
		uu.SetLogin(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// AddMetricIDs adds the "metrics" edge to the Metric entity by IDs.
func (uu *UserUpdate) AddMetricIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddMetricIDs(ids...)
	return uu
}

// AddMetrics adds the "metrics" edges to the Metric entity.
func (uu *UserUpdate) AddMetrics(m ...*Metric) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddMetricIDs(ids...)
}

// AddGivenAccessIDs adds the "given_accesses" edge to the Access entity by IDs.
func (uu *UserUpdate) AddGivenAccessIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddGivenAccessIDs(ids...)
	return uu
}

// AddGivenAccesses adds the "given_accesses" edges to the Access entity.
func (uu *UserUpdate) AddGivenAccesses(a ...*Access) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddGivenAccessIDs(ids...)
}

// AddRecievedAccessIDs adds the "recieved_accesses" edge to the Access entity by IDs.
func (uu *UserUpdate) AddRecievedAccessIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddRecievedAccessIDs(ids...)
	return uu
}

// AddRecievedAccesses adds the "recieved_accesses" edges to the Access entity.
func (uu *UserUpdate) AddRecievedAccesses(a ...*Access) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddRecievedAccessIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearMetrics clears all "metrics" edges to the Metric entity.
func (uu *UserUpdate) ClearMetrics() *UserUpdate {
	uu.mutation.ClearMetrics()
	return uu
}

// RemoveMetricIDs removes the "metrics" edge to Metric entities by IDs.
func (uu *UserUpdate) RemoveMetricIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveMetricIDs(ids...)
	return uu
}

// RemoveMetrics removes "metrics" edges to Metric entities.
func (uu *UserUpdate) RemoveMetrics(m ...*Metric) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveMetricIDs(ids...)
}

// ClearGivenAccesses clears all "given_accesses" edges to the Access entity.
func (uu *UserUpdate) ClearGivenAccesses() *UserUpdate {
	uu.mutation.ClearGivenAccesses()
	return uu
}

// RemoveGivenAccessIDs removes the "given_accesses" edge to Access entities by IDs.
func (uu *UserUpdate) RemoveGivenAccessIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveGivenAccessIDs(ids...)
	return uu
}

// RemoveGivenAccesses removes "given_accesses" edges to Access entities.
func (uu *UserUpdate) RemoveGivenAccesses(a ...*Access) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveGivenAccessIDs(ids...)
}

// ClearRecievedAccesses clears all "recieved_accesses" edges to the Access entity.
func (uu *UserUpdate) ClearRecievedAccesses() *UserUpdate {
	uu.mutation.ClearRecievedAccesses()
	return uu
}

// RemoveRecievedAccessIDs removes the "recieved_accesses" edge to Access entities by IDs.
func (uu *UserUpdate) RemoveRecievedAccessIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveRecievedAccessIDs(ids...)
	return uu
}

// RemoveRecievedAccesses removes "recieved_accesses" edges to Access entities.
func (uu *UserUpdate) RemoveRecievedAccesses(a ...*Access) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveRecievedAccessIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Login(); ok {
		_spec.SetField(user.FieldLogin, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetricsTable,
			Columns: []string{user.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metric.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedMetricsIDs(); len(nodes) > 0 && !uu.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetricsTable,
			Columns: []string{user.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metric.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetricsTable,
			Columns: []string{user.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metric.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GivenAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GivenAccessesTable,
			Columns: []string{user.GivenAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGivenAccessesIDs(); len(nodes) > 0 && !uu.mutation.GivenAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GivenAccessesTable,
			Columns: []string{user.GivenAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GivenAccessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GivenAccessesTable,
			Columns: []string{user.GivenAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RecievedAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RecievedAccessesTable,
			Columns: []string{user.RecievedAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRecievedAccessesIDs(); len(nodes) > 0 && !uu.mutation.RecievedAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RecievedAccessesTable,
			Columns: []string{user.RecievedAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RecievedAccessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RecievedAccessesTable,
			Columns: []string{user.RecievedAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetLogin sets the "login" field.
func (uuo *UserUpdateOne) SetLogin(s string) *UserUpdateOne {
	uuo.mutation.SetLogin(s)
	return uuo
}

// SetNillableLogin sets the "login" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLogin(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLogin(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// AddMetricIDs adds the "metrics" edge to the Metric entity by IDs.
func (uuo *UserUpdateOne) AddMetricIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddMetricIDs(ids...)
	return uuo
}

// AddMetrics adds the "metrics" edges to the Metric entity.
func (uuo *UserUpdateOne) AddMetrics(m ...*Metric) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddMetricIDs(ids...)
}

// AddGivenAccessIDs adds the "given_accesses" edge to the Access entity by IDs.
func (uuo *UserUpdateOne) AddGivenAccessIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddGivenAccessIDs(ids...)
	return uuo
}

// AddGivenAccesses adds the "given_accesses" edges to the Access entity.
func (uuo *UserUpdateOne) AddGivenAccesses(a ...*Access) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddGivenAccessIDs(ids...)
}

// AddRecievedAccessIDs adds the "recieved_accesses" edge to the Access entity by IDs.
func (uuo *UserUpdateOne) AddRecievedAccessIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddRecievedAccessIDs(ids...)
	return uuo
}

// AddRecievedAccesses adds the "recieved_accesses" edges to the Access entity.
func (uuo *UserUpdateOne) AddRecievedAccesses(a ...*Access) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddRecievedAccessIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearMetrics clears all "metrics" edges to the Metric entity.
func (uuo *UserUpdateOne) ClearMetrics() *UserUpdateOne {
	uuo.mutation.ClearMetrics()
	return uuo
}

// RemoveMetricIDs removes the "metrics" edge to Metric entities by IDs.
func (uuo *UserUpdateOne) RemoveMetricIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveMetricIDs(ids...)
	return uuo
}

// RemoveMetrics removes "metrics" edges to Metric entities.
func (uuo *UserUpdateOne) RemoveMetrics(m ...*Metric) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveMetricIDs(ids...)
}

// ClearGivenAccesses clears all "given_accesses" edges to the Access entity.
func (uuo *UserUpdateOne) ClearGivenAccesses() *UserUpdateOne {
	uuo.mutation.ClearGivenAccesses()
	return uuo
}

// RemoveGivenAccessIDs removes the "given_accesses" edge to Access entities by IDs.
func (uuo *UserUpdateOne) RemoveGivenAccessIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveGivenAccessIDs(ids...)
	return uuo
}

// RemoveGivenAccesses removes "given_accesses" edges to Access entities.
func (uuo *UserUpdateOne) RemoveGivenAccesses(a ...*Access) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveGivenAccessIDs(ids...)
}

// ClearRecievedAccesses clears all "recieved_accesses" edges to the Access entity.
func (uuo *UserUpdateOne) ClearRecievedAccesses() *UserUpdateOne {
	uuo.mutation.ClearRecievedAccesses()
	return uuo
}

// RemoveRecievedAccessIDs removes the "recieved_accesses" edge to Access entities by IDs.
func (uuo *UserUpdateOne) RemoveRecievedAccessIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveRecievedAccessIDs(ids...)
	return uuo
}

// RemoveRecievedAccesses removes "recieved_accesses" edges to Access entities.
func (uuo *UserUpdateOne) RemoveRecievedAccesses(a ...*Access) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveRecievedAccessIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Login(); ok {
		_spec.SetField(user.FieldLogin, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetricsTable,
			Columns: []string{user.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metric.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedMetricsIDs(); len(nodes) > 0 && !uuo.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetricsTable,
			Columns: []string{user.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metric.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MetricsTable,
			Columns: []string{user.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metric.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GivenAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GivenAccessesTable,
			Columns: []string{user.GivenAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGivenAccessesIDs(); len(nodes) > 0 && !uuo.mutation.GivenAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GivenAccessesTable,
			Columns: []string{user.GivenAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GivenAccessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GivenAccessesTable,
			Columns: []string{user.GivenAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RecievedAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RecievedAccessesTable,
			Columns: []string{user.RecievedAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRecievedAccessesIDs(); len(nodes) > 0 && !uuo.mutation.RecievedAccessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RecievedAccessesTable,
			Columns: []string{user.RecievedAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RecievedAccessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RecievedAccessesTable,
			Columns: []string{user.RecievedAccessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
