package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_type"),
		field.String("wallet_address"),
		field.String("transaction_hash"),
		field.Int64("block_number"),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type", "wallet_address", "transaction_hash").
			Unique(),
	}
}

func (Transaction) Edges() []ent.Edge {
	return nil
}
