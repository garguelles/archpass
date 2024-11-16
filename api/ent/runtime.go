// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/garguelles/archpass/ent/attendee"
	"github.com/garguelles/archpass/ent/event"
	"github.com/garguelles/archpass/ent/schema"
	"github.com/garguelles/archpass/ent/ticket"
	"github.com/garguelles/archpass/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	attendeeFields := schema.Attendee{}.Fields()
	_ = attendeeFields
	// attendeeDescCreatedAt is the schema descriptor for created_at field.
	attendeeDescCreatedAt := attendeeFields[6].Descriptor()
	// attendee.DefaultCreatedAt holds the default value on creation for the created_at field.
	attendee.DefaultCreatedAt = attendeeDescCreatedAt.Default.(func() time.Time)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescName is the schema descriptor for name field.
	eventDescName := eventFields[0].Descriptor()
	// event.NameValidator is a validator for the "name" field. It is called by the builders before save.
	event.NameValidator = eventDescName.Validators[0].(func(string) error)
	// eventDescEventSlug is the schema descriptor for event_slug field.
	eventDescEventSlug := eventFields[2].Descriptor()
	// event.EventSlugValidator is a validator for the "event_slug" field. It is called by the builders before save.
	event.EventSlugValidator = eventDescEventSlug.Validators[0].(func(string) error)
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[13].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescModifiedAt is the schema descriptor for modified_at field.
	eventDescModifiedAt := eventFields[14].Descriptor()
	// event.DefaultModifiedAt holds the default value on creation for the modified_at field.
	event.DefaultModifiedAt = eventDescModifiedAt.Default.(func() time.Time)
	// event.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	event.UpdateDefaultModifiedAt = eventDescModifiedAt.UpdateDefault.(func() time.Time)
	ticketFields := schema.Ticket{}.Fields()
	_ = ticketFields
	// ticketDescTicketSlug is the schema descriptor for ticket_slug field.
	ticketDescTicketSlug := ticketFields[2].Descriptor()
	// ticket.TicketSlugValidator is a validator for the "ticket_slug" field. It is called by the builders before save.
	ticket.TicketSlugValidator = ticketDescTicketSlug.Validators[0].(func(string) error)
	// ticketDescCreatedAt is the schema descriptor for created_at field.
	ticketDescCreatedAt := ticketFields[11].Descriptor()
	// ticket.DefaultCreatedAt holds the default value on creation for the created_at field.
	ticket.DefaultCreatedAt = ticketDescCreatedAt.Default.(func() time.Time)
	// ticketDescUpdatedAt is the schema descriptor for updated_at field.
	ticketDescUpdatedAt := ticketFields[12].Descriptor()
	// ticket.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ticket.DefaultUpdatedAt = ticketDescUpdatedAt.Default.(func() time.Time)
	// ticket.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ticket.UpdateDefaultUpdatedAt = ticketDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescWalletAddress is the schema descriptor for wallet_address field.
	userDescWalletAddress := userFields[0].Descriptor()
	// user.WalletAddressValidator is a validator for the "wallet_address" field. It is called by the builders before save.
	user.WalletAddressValidator = userDescWalletAddress.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
