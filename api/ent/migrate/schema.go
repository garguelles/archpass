// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttendeesColumns holds the columns for the "attendees" table.
	AttendeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ticket_id", Type: field.TypeInt},
		{Name: "event_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// AttendeesTable holds the schema information for the "attendees" table.
	AttendeesTable = &schema.Table{
		Name:       "attendees",
		Columns:    AttendeesColumns,
		PrimaryKey: []*schema.Column{AttendeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attendees_events_attendees",
				Columns:    []*schema.Column{AttendeesColumns[2]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "attendees_users_attendee_tickets",
				Columns:    []*schema.Column{AttendeesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "event_slug", Type: field.TypeString, Unique: true},
		{Name: "location", Type: field.TypeString, Size: 2147483647},
		{Name: "image_url", Type: field.TypeString},
		{Name: "contract_address", Type: field.TypeString, Nullable: true},
		{Name: "transaction_hash", Type: field.TypeString, Nullable: true},
		{Name: "block_number", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_users_events",
				Columns:    []*schema.Column{EventsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TicketsColumns holds the columns for the "tickets" table.
	TicketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "ticket_slug", Type: field.TypeString, Unique: true},
		{Name: "mint_price", Type: field.TypeString},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "contract_address", Type: field.TypeString, Nullable: true},
		{Name: "transaction_hash", Type: field.TypeString, Nullable: true},
		{Name: "block_number", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "event_id", Type: field.TypeInt},
	}
	// TicketsTable holds the schema information for the "tickets" table.
	TicketsTable = &schema.Table{
		Name:       "tickets",
		Columns:    TicketsColumns,
		PrimaryKey: []*schema.Column{TicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickets_events_tickets",
				Columns:    []*schema.Column{TicketsColumns[11]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "wallet_address", Type: field.TypeString, Unique: true},
		{Name: "bio", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttendeesTable,
		EventsTable,
		TicketsTable,
		UsersTable,
	}
)

func init() {
	AttendeesTable.ForeignKeys[0].RefTable = EventsTable
	AttendeesTable.ForeignKeys[1].RefTable = UsersTable
	EventsTable.ForeignKeys[0].RefTable = UsersTable
	TicketsTable.ForeignKeys[0].RefTable = EventsTable
}
