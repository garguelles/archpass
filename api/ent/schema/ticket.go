package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("ticket_slug").
			NotEmpty().
			Unique(),
		field.String("mint_price"),
		field.Int("quantity"),
		field.Int("event_id"),
		field.String("contract_address").
			Optional(),
		field.String("transaction_hash").
			Optional(),
		field.String("block_number").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Ticket
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).
			Ref("tickets").
			Field("event_id").
			Unique().
			Required(),
	}
}
