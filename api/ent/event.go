// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/garguelles/archpass/ent/event"
	"github.com/garguelles/archpass/ent/user"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// EventSlug holds the value of the "event_slug" field.
	EventSlug string `json:"event_slug,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// ContractAddress holds the value of the "contract_address" field.
	ContractAddress string `json:"contract_address,omitempty"`
	// TransactionHash holds the value of the "transaction_hash" field.
	TransactionHash string `json:"transaction_hash,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber string `json:"block_number,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges        EventEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// Tickets holds the value of the tickets edge.
	Tickets []*Ticket `json:"tickets,omitempty"`
	// Attendees holds the value of the attendees edge.
	Attendees []*Attendee `json:"attendees,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventEdges) UsersOrErr() (*User, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// TicketsOrErr returns the Tickets value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) TicketsOrErr() ([]*Ticket, error) {
	if e.loadedTypes[1] {
		return e.Tickets, nil
	}
	return nil, &NotLoadedError{edge: "tickets"}
}

// AttendeesOrErr returns the Attendees value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) AttendeesOrErr() ([]*Attendee, error) {
	if e.loadedTypes[2] {
		return e.Attendees, nil
	}
	return nil, &NotLoadedError{edge: "attendees"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldID, event.FieldUserID:
			values[i] = new(sql.NullInt64)
		case event.FieldName, event.FieldDescription, event.FieldEventSlug, event.FieldLocation, event.FieldImageURL, event.FieldContractAddress, event.FieldTransactionHash, event.FieldBlockNumber:
			values[i] = new(sql.NullString)
		case event.FieldCreatedAt, event.FieldModifiedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case event.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case event.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case event.FieldEventSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_slug", values[i])
			} else if value.Valid {
				e.EventSlug = value.String
			}
		case event.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				e.Location = value.String
			}
		case event.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				e.ImageURL = value.String
			}
		case event.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				e.UserID = int(value.Int64)
			}
		case event.FieldContractAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract_address", values[i])
			} else if value.Valid {
				e.ContractAddress = value.String
			}
		case event.FieldTransactionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_hash", values[i])
			} else if value.Valid {
				e.TransactionHash = value.String
			}
		case event.FieldBlockNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field block_number", values[i])
			} else if value.Valid {
				e.BlockNumber = value.String
			}
		case event.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case event.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				e.ModifiedAt = value.Time
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Event.
// This includes values selected through modifiers, order, etc.
func (e *Event) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Event entity.
func (e *Event) QueryUsers() *UserQuery {
	return NewEventClient(e.config).QueryUsers(e)
}

// QueryTickets queries the "tickets" edge of the Event entity.
func (e *Event) QueryTickets() *TicketQuery {
	return NewEventClient(e.config).QueryTickets(e)
}

// QueryAttendees queries the "attendees" edge of the Event entity.
func (e *Event) QueryAttendees() *AttendeeQuery {
	return NewEventClient(e.config).QueryAttendees(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return NewEventClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(e.Description)
	builder.WriteString(", ")
	builder.WriteString("event_slug=")
	builder.WriteString(e.EventSlug)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(e.Location)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(e.ImageURL)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", e.UserID))
	builder.WriteString(", ")
	builder.WriteString("contract_address=")
	builder.WriteString(e.ContractAddress)
	builder.WriteString(", ")
	builder.WriteString("transaction_hash=")
	builder.WriteString(e.TransactionHash)
	builder.WriteString(", ")
	builder.WriteString("block_number=")
	builder.WriteString(e.BlockNumber)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(e.ModifiedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event
