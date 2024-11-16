package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type Attendee struct {
	ent.Schema
}

func (Attendee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("event_id"),
		field.Int("ticket_id"),
		field.Int("token_id"),
		field.String("transaction_hash"),
		field.Int64("block_number"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Attendee) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token_id", "ticket_id", "user_id").
			Unique(),
	}
}

func (Attendee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("attendees").
			Field("event_id").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("attendee_tickets").
			Field("user_id").
			Unique().
			Required(),
		edge.From("ticket", Ticket.Type).
			Ref("attendees").
			Field("ticket_id").
			Unique().
			Required(),
	}
}
