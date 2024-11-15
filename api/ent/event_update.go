// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/garguelles/archpass/ent/attendee"
	"github.com/garguelles/archpass/ent/event"
	"github.com/garguelles/archpass/ent/predicate"
	"github.com/garguelles/archpass/ent/ticket"
	"github.com/garguelles/archpass/ent/user"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventUpdate) SetName(s string) *EventUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EventUpdate) SetNillableName(s *string) *EventUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetDescription sets the "description" field.
func (eu *EventUpdate) SetDescription(s string) *EventUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *EventUpdate) SetNillableDescription(s *string) *EventUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// ClearDescription clears the value of the "description" field.
func (eu *EventUpdate) ClearDescription() *EventUpdate {
	eu.mutation.ClearDescription()
	return eu
}

// SetEventSlug sets the "event_slug" field.
func (eu *EventUpdate) SetEventSlug(s string) *EventUpdate {
	eu.mutation.SetEventSlug(s)
	return eu
}

// SetNillableEventSlug sets the "event_slug" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEventSlug(s *string) *EventUpdate {
	if s != nil {
		eu.SetEventSlug(*s)
	}
	return eu
}

// SetLocation sets the "location" field.
func (eu *EventUpdate) SetLocation(s string) *EventUpdate {
	eu.mutation.SetLocation(s)
	return eu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (eu *EventUpdate) SetNillableLocation(s *string) *EventUpdate {
	if s != nil {
		eu.SetLocation(*s)
	}
	return eu
}

// SetImageURL sets the "image_url" field.
func (eu *EventUpdate) SetImageURL(s string) *EventUpdate {
	eu.mutation.SetImageURL(s)
	return eu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (eu *EventUpdate) SetNillableImageURL(s *string) *EventUpdate {
	if s != nil {
		eu.SetImageURL(*s)
	}
	return eu
}

// SetUserID sets the "user_id" field.
func (eu *EventUpdate) SetUserID(i int) *EventUpdate {
	eu.mutation.SetUserID(i)
	return eu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (eu *EventUpdate) SetNillableUserID(i *int) *EventUpdate {
	if i != nil {
		eu.SetUserID(*i)
	}
	return eu
}

// SetContractAddress sets the "contract_address" field.
func (eu *EventUpdate) SetContractAddress(s string) *EventUpdate {
	eu.mutation.SetContractAddress(s)
	return eu
}

// SetNillableContractAddress sets the "contract_address" field if the given value is not nil.
func (eu *EventUpdate) SetNillableContractAddress(s *string) *EventUpdate {
	if s != nil {
		eu.SetContractAddress(*s)
	}
	return eu
}

// ClearContractAddress clears the value of the "contract_address" field.
func (eu *EventUpdate) ClearContractAddress() *EventUpdate {
	eu.mutation.ClearContractAddress()
	return eu
}

// SetTransactionHash sets the "transaction_hash" field.
func (eu *EventUpdate) SetTransactionHash(s string) *EventUpdate {
	eu.mutation.SetTransactionHash(s)
	return eu
}

// SetNillableTransactionHash sets the "transaction_hash" field if the given value is not nil.
func (eu *EventUpdate) SetNillableTransactionHash(s *string) *EventUpdate {
	if s != nil {
		eu.SetTransactionHash(*s)
	}
	return eu
}

// ClearTransactionHash clears the value of the "transaction_hash" field.
func (eu *EventUpdate) ClearTransactionHash() *EventUpdate {
	eu.mutation.ClearTransactionHash()
	return eu
}

// SetBlockNumber sets the "block_number" field.
func (eu *EventUpdate) SetBlockNumber(s string) *EventUpdate {
	eu.mutation.SetBlockNumber(s)
	return eu
}

// SetNillableBlockNumber sets the "block_number" field if the given value is not nil.
func (eu *EventUpdate) SetNillableBlockNumber(s *string) *EventUpdate {
	if s != nil {
		eu.SetBlockNumber(*s)
	}
	return eu
}

// ClearBlockNumber clears the value of the "block_number" field.
func (eu *EventUpdate) ClearBlockNumber() *EventUpdate {
	eu.mutation.ClearBlockNumber()
	return eu
}

// SetCreatedAt sets the "created_at" field.
func (eu *EventUpdate) SetCreatedAt(t time.Time) *EventUpdate {
	eu.mutation.SetCreatedAt(t)
	return eu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eu *EventUpdate) SetNillableCreatedAt(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetCreatedAt(*t)
	}
	return eu
}

// SetModifiedAt sets the "modified_at" field.
func (eu *EventUpdate) SetModifiedAt(t time.Time) *EventUpdate {
	eu.mutation.SetModifiedAt(t)
	return eu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (eu *EventUpdate) SetUsersID(id int) *EventUpdate {
	eu.mutation.SetUsersID(id)
	return eu
}

// SetUsers sets the "users" edge to the User entity.
func (eu *EventUpdate) SetUsers(u *User) *EventUpdate {
	return eu.SetUsersID(u.ID)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (eu *EventUpdate) AddTicketIDs(ids ...int) *EventUpdate {
	eu.mutation.AddTicketIDs(ids...)
	return eu
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (eu *EventUpdate) AddTickets(t ...*Ticket) *EventUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eu.AddTicketIDs(ids...)
}

// AddAttendeeIDs adds the "attendees" edge to the Attendee entity by IDs.
func (eu *EventUpdate) AddAttendeeIDs(ids ...int) *EventUpdate {
	eu.mutation.AddAttendeeIDs(ids...)
	return eu
}

// AddAttendees adds the "attendees" edges to the Attendee entity.
func (eu *EventUpdate) AddAttendees(a ...*Attendee) *EventUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.AddAttendeeIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (eu *EventUpdate) ClearUsers() *EventUpdate {
	eu.mutation.ClearUsers()
	return eu
}

// ClearTickets clears all "tickets" edges to the Ticket entity.
func (eu *EventUpdate) ClearTickets() *EventUpdate {
	eu.mutation.ClearTickets()
	return eu
}

// RemoveTicketIDs removes the "tickets" edge to Ticket entities by IDs.
func (eu *EventUpdate) RemoveTicketIDs(ids ...int) *EventUpdate {
	eu.mutation.RemoveTicketIDs(ids...)
	return eu
}

// RemoveTickets removes "tickets" edges to Ticket entities.
func (eu *EventUpdate) RemoveTickets(t ...*Ticket) *EventUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eu.RemoveTicketIDs(ids...)
}

// ClearAttendees clears all "attendees" edges to the Attendee entity.
func (eu *EventUpdate) ClearAttendees() *EventUpdate {
	eu.mutation.ClearAttendees()
	return eu
}

// RemoveAttendeeIDs removes the "attendees" edge to Attendee entities by IDs.
func (eu *EventUpdate) RemoveAttendeeIDs(ids ...int) *EventUpdate {
	eu.mutation.RemoveAttendeeIDs(ids...)
	return eu
}

// RemoveAttendees removes "attendees" edges to Attendee entities.
func (eu *EventUpdate) RemoveAttendees(a ...*Attendee) *EventUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.RemoveAttendeeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	eu.defaults()
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EventUpdate) defaults() {
	if _, ok := eu.mutation.ModifiedAt(); !ok {
		v := event.UpdateDefaultModifiedAt()
		eu.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.EventSlug(); ok {
		if err := event.EventSlugValidator(v); err != nil {
			return &ValidationError{Name: "event_slug", err: fmt.Errorf(`ent: validator failed for field "Event.event_slug": %w`, err)}
		}
	}
	if eu.mutation.UsersCleared() && len(eu.mutation.UsersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Event.users"`)
	}
	return nil
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(event.Table, event.Columns, sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(event.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.SetField(event.FieldDescription, field.TypeString, value)
	}
	if eu.mutation.DescriptionCleared() {
		_spec.ClearField(event.FieldDescription, field.TypeString)
	}
	if value, ok := eu.mutation.EventSlug(); ok {
		_spec.SetField(event.FieldEventSlug, field.TypeString, value)
	}
	if value, ok := eu.mutation.Location(); ok {
		_spec.SetField(event.FieldLocation, field.TypeString, value)
	}
	if value, ok := eu.mutation.ImageURL(); ok {
		_spec.SetField(event.FieldImageURL, field.TypeString, value)
	}
	if value, ok := eu.mutation.ContractAddress(); ok {
		_spec.SetField(event.FieldContractAddress, field.TypeString, value)
	}
	if eu.mutation.ContractAddressCleared() {
		_spec.ClearField(event.FieldContractAddress, field.TypeString)
	}
	if value, ok := eu.mutation.TransactionHash(); ok {
		_spec.SetField(event.FieldTransactionHash, field.TypeString, value)
	}
	if eu.mutation.TransactionHashCleared() {
		_spec.ClearField(event.FieldTransactionHash, field.TypeString)
	}
	if value, ok := eu.mutation.BlockNumber(); ok {
		_spec.SetField(event.FieldBlockNumber, field.TypeString, value)
	}
	if eu.mutation.BlockNumberCleared() {
		_spec.ClearField(event.FieldBlockNumber, field.TypeString)
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.ModifiedAt(); ok {
		_spec.SetField(event.FieldModifiedAt, field.TypeTime, value)
	}
	if eu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.UsersTable,
			Columns: []string{event.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.UsersTable,
			Columns: []string{event.UsersColumn},
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
	if eu.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedTicketsIDs(); len(nodes) > 0 && !eu.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
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
	if eu.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: []string{event.AttendeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedAttendeesIDs(); len(nodes) > 0 && !eu.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: []string{event.AttendeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.AttendeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: []string{event.AttendeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetName sets the "name" field.
func (euo *EventUpdateOne) SetName(s string) *EventUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableName(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetDescription sets the "description" field.
func (euo *EventUpdateOne) SetDescription(s string) *EventUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableDescription(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// ClearDescription clears the value of the "description" field.
func (euo *EventUpdateOne) ClearDescription() *EventUpdateOne {
	euo.mutation.ClearDescription()
	return euo
}

// SetEventSlug sets the "event_slug" field.
func (euo *EventUpdateOne) SetEventSlug(s string) *EventUpdateOne {
	euo.mutation.SetEventSlug(s)
	return euo
}

// SetNillableEventSlug sets the "event_slug" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEventSlug(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetEventSlug(*s)
	}
	return euo
}

// SetLocation sets the "location" field.
func (euo *EventUpdateOne) SetLocation(s string) *EventUpdateOne {
	euo.mutation.SetLocation(s)
	return euo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableLocation(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetLocation(*s)
	}
	return euo
}

// SetImageURL sets the "image_url" field.
func (euo *EventUpdateOne) SetImageURL(s string) *EventUpdateOne {
	euo.mutation.SetImageURL(s)
	return euo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableImageURL(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetImageURL(*s)
	}
	return euo
}

// SetUserID sets the "user_id" field.
func (euo *EventUpdateOne) SetUserID(i int) *EventUpdateOne {
	euo.mutation.SetUserID(i)
	return euo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableUserID(i *int) *EventUpdateOne {
	if i != nil {
		euo.SetUserID(*i)
	}
	return euo
}

// SetContractAddress sets the "contract_address" field.
func (euo *EventUpdateOne) SetContractAddress(s string) *EventUpdateOne {
	euo.mutation.SetContractAddress(s)
	return euo
}

// SetNillableContractAddress sets the "contract_address" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableContractAddress(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetContractAddress(*s)
	}
	return euo
}

// ClearContractAddress clears the value of the "contract_address" field.
func (euo *EventUpdateOne) ClearContractAddress() *EventUpdateOne {
	euo.mutation.ClearContractAddress()
	return euo
}

// SetTransactionHash sets the "transaction_hash" field.
func (euo *EventUpdateOne) SetTransactionHash(s string) *EventUpdateOne {
	euo.mutation.SetTransactionHash(s)
	return euo
}

// SetNillableTransactionHash sets the "transaction_hash" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableTransactionHash(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetTransactionHash(*s)
	}
	return euo
}

// ClearTransactionHash clears the value of the "transaction_hash" field.
func (euo *EventUpdateOne) ClearTransactionHash() *EventUpdateOne {
	euo.mutation.ClearTransactionHash()
	return euo
}

// SetBlockNumber sets the "block_number" field.
func (euo *EventUpdateOne) SetBlockNumber(s string) *EventUpdateOne {
	euo.mutation.SetBlockNumber(s)
	return euo
}

// SetNillableBlockNumber sets the "block_number" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableBlockNumber(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetBlockNumber(*s)
	}
	return euo
}

// ClearBlockNumber clears the value of the "block_number" field.
func (euo *EventUpdateOne) ClearBlockNumber() *EventUpdateOne {
	euo.mutation.ClearBlockNumber()
	return euo
}

// SetCreatedAt sets the "created_at" field.
func (euo *EventUpdateOne) SetCreatedAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetCreatedAt(t)
	return euo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCreatedAt(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetCreatedAt(*t)
	}
	return euo
}

// SetModifiedAt sets the "modified_at" field.
func (euo *EventUpdateOne) SetModifiedAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetModifiedAt(t)
	return euo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (euo *EventUpdateOne) SetUsersID(id int) *EventUpdateOne {
	euo.mutation.SetUsersID(id)
	return euo
}

// SetUsers sets the "users" edge to the User entity.
func (euo *EventUpdateOne) SetUsers(u *User) *EventUpdateOne {
	return euo.SetUsersID(u.ID)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (euo *EventUpdateOne) AddTicketIDs(ids ...int) *EventUpdateOne {
	euo.mutation.AddTicketIDs(ids...)
	return euo
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (euo *EventUpdateOne) AddTickets(t ...*Ticket) *EventUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return euo.AddTicketIDs(ids...)
}

// AddAttendeeIDs adds the "attendees" edge to the Attendee entity by IDs.
func (euo *EventUpdateOne) AddAttendeeIDs(ids ...int) *EventUpdateOne {
	euo.mutation.AddAttendeeIDs(ids...)
	return euo
}

// AddAttendees adds the "attendees" edges to the Attendee entity.
func (euo *EventUpdateOne) AddAttendees(a ...*Attendee) *EventUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.AddAttendeeIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (euo *EventUpdateOne) ClearUsers() *EventUpdateOne {
	euo.mutation.ClearUsers()
	return euo
}

// ClearTickets clears all "tickets" edges to the Ticket entity.
func (euo *EventUpdateOne) ClearTickets() *EventUpdateOne {
	euo.mutation.ClearTickets()
	return euo
}

// RemoveTicketIDs removes the "tickets" edge to Ticket entities by IDs.
func (euo *EventUpdateOne) RemoveTicketIDs(ids ...int) *EventUpdateOne {
	euo.mutation.RemoveTicketIDs(ids...)
	return euo
}

// RemoveTickets removes "tickets" edges to Ticket entities.
func (euo *EventUpdateOne) RemoveTickets(t ...*Ticket) *EventUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return euo.RemoveTicketIDs(ids...)
}

// ClearAttendees clears all "attendees" edges to the Attendee entity.
func (euo *EventUpdateOne) ClearAttendees() *EventUpdateOne {
	euo.mutation.ClearAttendees()
	return euo
}

// RemoveAttendeeIDs removes the "attendees" edge to Attendee entities by IDs.
func (euo *EventUpdateOne) RemoveAttendeeIDs(ids ...int) *EventUpdateOne {
	euo.mutation.RemoveAttendeeIDs(ids...)
	return euo
}

// RemoveAttendees removes "attendees" edges to Attendee entities.
func (euo *EventUpdateOne) RemoveAttendees(a ...*Attendee) *EventUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.RemoveAttendeeIDs(ids...)
}

// Where appends a list predicates to the EventUpdate builder.
func (euo *EventUpdateOne) Where(ps ...predicate.Event) *EventUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	euo.defaults()
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EventUpdateOne) defaults() {
	if _, ok := euo.mutation.ModifiedAt(); !ok {
		v := event.UpdateDefaultModifiedAt()
		euo.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.EventSlug(); ok {
		if err := event.EventSlugValidator(v); err != nil {
			return &ValidationError{Name: "event_slug", err: fmt.Errorf(`ent: validator failed for field "Event.event_slug": %w`, err)}
		}
	}
	if euo.mutation.UsersCleared() && len(euo.mutation.UsersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Event.users"`)
	}
	return nil
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(event.Table, event.Columns, sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(event.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.SetField(event.FieldDescription, field.TypeString, value)
	}
	if euo.mutation.DescriptionCleared() {
		_spec.ClearField(event.FieldDescription, field.TypeString)
	}
	if value, ok := euo.mutation.EventSlug(); ok {
		_spec.SetField(event.FieldEventSlug, field.TypeString, value)
	}
	if value, ok := euo.mutation.Location(); ok {
		_spec.SetField(event.FieldLocation, field.TypeString, value)
	}
	if value, ok := euo.mutation.ImageURL(); ok {
		_spec.SetField(event.FieldImageURL, field.TypeString, value)
	}
	if value, ok := euo.mutation.ContractAddress(); ok {
		_spec.SetField(event.FieldContractAddress, field.TypeString, value)
	}
	if euo.mutation.ContractAddressCleared() {
		_spec.ClearField(event.FieldContractAddress, field.TypeString)
	}
	if value, ok := euo.mutation.TransactionHash(); ok {
		_spec.SetField(event.FieldTransactionHash, field.TypeString, value)
	}
	if euo.mutation.TransactionHashCleared() {
		_spec.ClearField(event.FieldTransactionHash, field.TypeString)
	}
	if value, ok := euo.mutation.BlockNumber(); ok {
		_spec.SetField(event.FieldBlockNumber, field.TypeString, value)
	}
	if euo.mutation.BlockNumberCleared() {
		_spec.ClearField(event.FieldBlockNumber, field.TypeString)
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.ModifiedAt(); ok {
		_spec.SetField(event.FieldModifiedAt, field.TypeTime, value)
	}
	if euo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.UsersTable,
			Columns: []string{event.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.UsersTable,
			Columns: []string{event.UsersColumn},
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
	if euo.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedTicketsIDs(); len(nodes) > 0 && !euo.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
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
	if euo.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: []string{event.AttendeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedAttendeesIDs(); len(nodes) > 0 && !euo.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: []string{event.AttendeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.AttendeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: []string{event.AttendeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
