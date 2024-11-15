// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/garguelles/archpass/ent/attendee"
	"github.com/garguelles/archpass/ent/event"
	"github.com/garguelles/archpass/ent/ticket"
	"github.com/garguelles/archpass/ent/user"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *EventCreate) SetName(s string) *EventCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetDescription sets the "description" field.
func (ec *EventCreate) SetDescription(s string) *EventCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ec *EventCreate) SetNillableDescription(s *string) *EventCreate {
	if s != nil {
		ec.SetDescription(*s)
	}
	return ec
}

// SetEventSlug sets the "event_slug" field.
func (ec *EventCreate) SetEventSlug(s string) *EventCreate {
	ec.mutation.SetEventSlug(s)
	return ec
}

// SetStartDate sets the "start_date" field.
func (ec *EventCreate) SetStartDate(t time.Time) *EventCreate {
	ec.mutation.SetStartDate(t)
	return ec
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (ec *EventCreate) SetNillableStartDate(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetStartDate(*t)
	}
	return ec
}

// SetEndDate sets the "end_date" field.
func (ec *EventCreate) SetEndDate(t time.Time) *EventCreate {
	ec.mutation.SetEndDate(t)
	return ec
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (ec *EventCreate) SetNillableEndDate(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetEndDate(*t)
	}
	return ec
}

// SetDate sets the "date" field.
func (ec *EventCreate) SetDate(s string) *EventCreate {
	ec.mutation.SetDate(s)
	return ec
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ec *EventCreate) SetNillableDate(s *string) *EventCreate {
	if s != nil {
		ec.SetDate(*s)
	}
	return ec
}

// SetLocation sets the "location" field.
func (ec *EventCreate) SetLocation(s string) *EventCreate {
	ec.mutation.SetLocation(s)
	return ec
}

// SetImageURL sets the "image_url" field.
func (ec *EventCreate) SetImageURL(s string) *EventCreate {
	ec.mutation.SetImageURL(s)
	return ec
}

// SetUserID sets the "user_id" field.
func (ec *EventCreate) SetUserID(i int) *EventCreate {
	ec.mutation.SetUserID(i)
	return ec
}

// SetContractAddress sets the "contract_address" field.
func (ec *EventCreate) SetContractAddress(s string) *EventCreate {
	ec.mutation.SetContractAddress(s)
	return ec
}

// SetNillableContractAddress sets the "contract_address" field if the given value is not nil.
func (ec *EventCreate) SetNillableContractAddress(s *string) *EventCreate {
	if s != nil {
		ec.SetContractAddress(*s)
	}
	return ec
}

// SetTransactionHash sets the "transaction_hash" field.
func (ec *EventCreate) SetTransactionHash(s string) *EventCreate {
	ec.mutation.SetTransactionHash(s)
	return ec
}

// SetNillableTransactionHash sets the "transaction_hash" field if the given value is not nil.
func (ec *EventCreate) SetNillableTransactionHash(s *string) *EventCreate {
	if s != nil {
		ec.SetTransactionHash(*s)
	}
	return ec
}

// SetBlockNumber sets the "block_number" field.
func (ec *EventCreate) SetBlockNumber(s string) *EventCreate {
	ec.mutation.SetBlockNumber(s)
	return ec
}

// SetNillableBlockNumber sets the "block_number" field if the given value is not nil.
func (ec *EventCreate) SetNillableBlockNumber(s *string) *EventCreate {
	if s != nil {
		ec.SetBlockNumber(*s)
	}
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EventCreate) SetCreatedAt(t time.Time) *EventCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreatedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetModifiedAt sets the "modified_at" field.
func (ec *EventCreate) SetModifiedAt(t time.Time) *EventCreate {
	ec.mutation.SetModifiedAt(t)
	return ec
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableModifiedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetModifiedAt(*t)
	}
	return ec
}

// SetUser sets the "user" edge to the User entity.
func (ec *EventCreate) SetUser(u *User) *EventCreate {
	return ec.SetUserID(u.ID)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (ec *EventCreate) AddTicketIDs(ids ...int) *EventCreate {
	ec.mutation.AddTicketIDs(ids...)
	return ec
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (ec *EventCreate) AddTickets(t ...*Ticket) *EventCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ec.AddTicketIDs(ids...)
}

// AddAttendeeIDs adds the "attendees" edge to the Attendee entity by IDs.
func (ec *EventCreate) AddAttendeeIDs(ids ...int) *EventCreate {
	ec.mutation.AddAttendeeIDs(ids...)
	return ec
}

// AddAttendees adds the "attendees" edges to the Attendee entity.
func (ec *EventCreate) AddAttendees(a ...*Attendee) *EventCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ec.AddAttendeeIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := event.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.ModifiedAt(); !ok {
		v := event.DefaultModifiedAt()
		ec.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Event.name"`)}
	}
	if v, ok := ec.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if _, ok := ec.mutation.EventSlug(); !ok {
		return &ValidationError{Name: "event_slug", err: errors.New(`ent: missing required field "Event.event_slug"`)}
	}
	if v, ok := ec.mutation.EventSlug(); ok {
		if err := event.EventSlugValidator(v); err != nil {
			return &ValidationError{Name: "event_slug", err: fmt.Errorf(`ent: validator failed for field "Event.event_slug": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "Event.location"`)}
	}
	if _, ok := ec.mutation.ImageURL(); !ok {
		return &ValidationError{Name: "image_url", err: errors.New(`ent: missing required field "Event.image_url"`)}
	}
	if _, ok := ec.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Event.user_id"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Event.created_at"`)}
	}
	if _, ok := ec.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New(`ent: missing required field "Event.modified_at"`)}
	}
	if len(ec.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Event.user"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(event.Table, sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(event.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.SetField(event.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ec.mutation.EventSlug(); ok {
		_spec.SetField(event.FieldEventSlug, field.TypeString, value)
		_node.EventSlug = value
	}
	if value, ok := ec.mutation.StartDate(); ok {
		_spec.SetField(event.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := ec.mutation.EndDate(); ok {
		_spec.SetField(event.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := ec.mutation.Date(); ok {
		_spec.SetField(event.FieldDate, field.TypeString, value)
		_node.Date = value
	}
	if value, ok := ec.mutation.Location(); ok {
		_spec.SetField(event.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := ec.mutation.ImageURL(); ok {
		_spec.SetField(event.FieldImageURL, field.TypeString, value)
		_node.ImageURL = value
	}
	if value, ok := ec.mutation.ContractAddress(); ok {
		_spec.SetField(event.FieldContractAddress, field.TypeString, value)
		_node.ContractAddress = value
	}
	if value, ok := ec.mutation.TransactionHash(); ok {
		_spec.SetField(event.FieldTransactionHash, field.TypeString, value)
		_node.TransactionHash = value
	}
	if value, ok := ec.mutation.BlockNumber(); ok {
		_spec.SetField(event.FieldBlockNumber, field.TypeString, value)
		_node.BlockNumber = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.ModifiedAt(); ok {
		_spec.SetField(event.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if nodes := ec.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.UserTable,
			Columns: []string{event.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.TicketsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.AttendeesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	err      error
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
