// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/garguelles/archpass/ent/attendee"
	"github.com/garguelles/archpass/ent/event"
	"github.com/garguelles/archpass/ent/predicate"
	"github.com/garguelles/archpass/ent/ticket"
	"github.com/garguelles/archpass/ent/user"
)

// AttendeeUpdate is the builder for updating Attendee entities.
type AttendeeUpdate struct {
	config
	hooks    []Hook
	mutation *AttendeeMutation
}

// Where appends a list predicates to the AttendeeUpdate builder.
func (au *AttendeeUpdate) Where(ps ...predicate.Attendee) *AttendeeUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUserID sets the "user_id" field.
func (au *AttendeeUpdate) SetUserID(i int) *AttendeeUpdate {
	au.mutation.SetUserID(i)
	return au
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (au *AttendeeUpdate) SetNillableUserID(i *int) *AttendeeUpdate {
	if i != nil {
		au.SetUserID(*i)
	}
	return au
}

// SetEventID sets the "event_id" field.
func (au *AttendeeUpdate) SetEventID(i int) *AttendeeUpdate {
	au.mutation.SetEventID(i)
	return au
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (au *AttendeeUpdate) SetNillableEventID(i *int) *AttendeeUpdate {
	if i != nil {
		au.SetEventID(*i)
	}
	return au
}

// SetTicketID sets the "ticket_id" field.
func (au *AttendeeUpdate) SetTicketID(i int) *AttendeeUpdate {
	au.mutation.SetTicketID(i)
	return au
}

// SetNillableTicketID sets the "ticket_id" field if the given value is not nil.
func (au *AttendeeUpdate) SetNillableTicketID(i *int) *AttendeeUpdate {
	if i != nil {
		au.SetTicketID(*i)
	}
	return au
}

// SetEvent sets the "event" edge to the Event entity.
func (au *AttendeeUpdate) SetEvent(e *Event) *AttendeeUpdate {
	return au.SetEventID(e.ID)
}

// SetUser sets the "user" edge to the User entity.
func (au *AttendeeUpdate) SetUser(u *User) *AttendeeUpdate {
	return au.SetUserID(u.ID)
}

// SetTicket sets the "ticket" edge to the Ticket entity.
func (au *AttendeeUpdate) SetTicket(t *Ticket) *AttendeeUpdate {
	return au.SetTicketID(t.ID)
}

// Mutation returns the AttendeeMutation object of the builder.
func (au *AttendeeUpdate) Mutation() *AttendeeMutation {
	return au.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (au *AttendeeUpdate) ClearEvent() *AttendeeUpdate {
	au.mutation.ClearEvent()
	return au
}

// ClearUser clears the "user" edge to the User entity.
func (au *AttendeeUpdate) ClearUser() *AttendeeUpdate {
	au.mutation.ClearUser()
	return au
}

// ClearTicket clears the "ticket" edge to the Ticket entity.
func (au *AttendeeUpdate) ClearTicket() *AttendeeUpdate {
	au.mutation.ClearTicket()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttendeeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttendeeUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttendeeUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttendeeUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AttendeeUpdate) check() error {
	if au.mutation.EventCleared() && len(au.mutation.EventIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Attendee.event"`)
	}
	if au.mutation.UserCleared() && len(au.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Attendee.user"`)
	}
	if au.mutation.TicketCleared() && len(au.mutation.TicketIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Attendee.ticket"`)
	}
	return nil
}

func (au *AttendeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(attendee.Table, attendee.Columns, sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if au.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.EventTable,
			Columns: []string{attendee.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.EventTable,
			Columns: []string{attendee.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.UserTable,
			Columns: []string{attendee.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.UserTable,
			Columns: []string{attendee.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attendee.TicketTable,
			Columns: []string{attendee.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attendee.TicketTable,
			Columns: []string{attendee.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttendeeUpdateOne is the builder for updating a single Attendee entity.
type AttendeeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttendeeMutation
}

// SetUserID sets the "user_id" field.
func (auo *AttendeeUpdateOne) SetUserID(i int) *AttendeeUpdateOne {
	auo.mutation.SetUserID(i)
	return auo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (auo *AttendeeUpdateOne) SetNillableUserID(i *int) *AttendeeUpdateOne {
	if i != nil {
		auo.SetUserID(*i)
	}
	return auo
}

// SetEventID sets the "event_id" field.
func (auo *AttendeeUpdateOne) SetEventID(i int) *AttendeeUpdateOne {
	auo.mutation.SetEventID(i)
	return auo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (auo *AttendeeUpdateOne) SetNillableEventID(i *int) *AttendeeUpdateOne {
	if i != nil {
		auo.SetEventID(*i)
	}
	return auo
}

// SetTicketID sets the "ticket_id" field.
func (auo *AttendeeUpdateOne) SetTicketID(i int) *AttendeeUpdateOne {
	auo.mutation.SetTicketID(i)
	return auo
}

// SetNillableTicketID sets the "ticket_id" field if the given value is not nil.
func (auo *AttendeeUpdateOne) SetNillableTicketID(i *int) *AttendeeUpdateOne {
	if i != nil {
		auo.SetTicketID(*i)
	}
	return auo
}

// SetEvent sets the "event" edge to the Event entity.
func (auo *AttendeeUpdateOne) SetEvent(e *Event) *AttendeeUpdateOne {
	return auo.SetEventID(e.ID)
}

// SetUser sets the "user" edge to the User entity.
func (auo *AttendeeUpdateOne) SetUser(u *User) *AttendeeUpdateOne {
	return auo.SetUserID(u.ID)
}

// SetTicket sets the "ticket" edge to the Ticket entity.
func (auo *AttendeeUpdateOne) SetTicket(t *Ticket) *AttendeeUpdateOne {
	return auo.SetTicketID(t.ID)
}

// Mutation returns the AttendeeMutation object of the builder.
func (auo *AttendeeUpdateOne) Mutation() *AttendeeMutation {
	return auo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (auo *AttendeeUpdateOne) ClearEvent() *AttendeeUpdateOne {
	auo.mutation.ClearEvent()
	return auo
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AttendeeUpdateOne) ClearUser() *AttendeeUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// ClearTicket clears the "ticket" edge to the Ticket entity.
func (auo *AttendeeUpdateOne) ClearTicket() *AttendeeUpdateOne {
	auo.mutation.ClearTicket()
	return auo
}

// Where appends a list predicates to the AttendeeUpdate builder.
func (auo *AttendeeUpdateOne) Where(ps ...predicate.Attendee) *AttendeeUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttendeeUpdateOne) Select(field string, fields ...string) *AttendeeUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attendee entity.
func (auo *AttendeeUpdateOne) Save(ctx context.Context) (*Attendee, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttendeeUpdateOne) SaveX(ctx context.Context) *Attendee {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttendeeUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttendeeUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AttendeeUpdateOne) check() error {
	if auo.mutation.EventCleared() && len(auo.mutation.EventIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Attendee.event"`)
	}
	if auo.mutation.UserCleared() && len(auo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Attendee.user"`)
	}
	if auo.mutation.TicketCleared() && len(auo.mutation.TicketIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Attendee.ticket"`)
	}
	return nil
}

func (auo *AttendeeUpdateOne) sqlSave(ctx context.Context) (_node *Attendee, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(attendee.Table, attendee.Columns, sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attendee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendee.FieldID)
		for _, f := range fields {
			if !attendee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attendee.FieldID {
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
	if auo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.EventTable,
			Columns: []string{attendee.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.EventTable,
			Columns: []string{attendee.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.UserTable,
			Columns: []string{attendee.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.UserTable,
			Columns: []string{attendee.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attendee.TicketTable,
			Columns: []string{attendee.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attendee.TicketTable,
			Columns: []string{attendee.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attendee{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
