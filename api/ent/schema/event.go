package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("description").
			Optional(),
		field.String("event_slug").
			NotEmpty().
			Unique(),
		field.Time("start_date").Optional(),
		field.Time("end_date").Optional(),
		field.String("date").Optional(),
		field.Text("location"),
		field.String("image_url"),
		field.Int("user_id"),
		field.String("contract_address").
			Optional(),
		field.String("transaction_hash").
			Optional(),
		field.String("block_number").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("events").
			Field("user_id").
			Unique().
			Required(),
		edge.To("tickets", Ticket.Type),
		edge.To("attendees", Attendee.Type),
	}
}
