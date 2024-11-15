package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Attendee struct {
	ent.Schema
}

func (Attendee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("event_id"),
		field.Int("ticket_id"),
	}
}

func (Attendee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).
			Ref("attendees").
			Field("event_id").
			Unique().
			Required(),
		edge.From("users", User.Type).
			Ref("attendee_tickets").
			Field("user_id").
			Unique().
			Required(),
	}
}
