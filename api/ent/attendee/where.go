// Code generated by ent, DO NOT EDIT.

package attendee

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/garguelles/archpass/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Attendee {
	return predicate.Attendee(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Attendee {
	return predicate.Attendee(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Attendee {
	return predicate.Attendee(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Attendee {
	return predicate.Attendee(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldUserID, v))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldEventID, v))
}

// TicketID applies equality check predicate on the "ticket_id" field. It's identical to TicketIDEQ.
func TicketID(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldTicketID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNotIn(FieldUserID, vs...))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNotIn(FieldEventID, vs...))
}

// TicketIDEQ applies the EQ predicate on the "ticket_id" field.
func TicketIDEQ(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldEQ(FieldTicketID, v))
}

// TicketIDNEQ applies the NEQ predicate on the "ticket_id" field.
func TicketIDNEQ(v int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNEQ(FieldTicketID, v))
}

// TicketIDIn applies the In predicate on the "ticket_id" field.
func TicketIDIn(vs ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldIn(FieldTicketID, vs...))
}

// TicketIDNotIn applies the NotIn predicate on the "ticket_id" field.
func TicketIDNotIn(vs ...int) predicate.Attendee {
	return predicate.Attendee(sql.FieldNotIn(FieldTicketID, vs...))
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Attendee {
	return predicate.Attendee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Attendee {
	return predicate.Attendee(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Attendee {
	return predicate.Attendee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Attendee {
	return predicate.Attendee(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTicket applies the HasEdge predicate on the "ticket" edge.
func HasTicket() predicate.Attendee {
	return predicate.Attendee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TicketTable, TicketColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTicketWith applies the HasEdge predicate on the "ticket" edge with a given conditions (other predicates).
func HasTicketWith(preds ...predicate.Ticket) predicate.Attendee {
	return predicate.Attendee(func(s *sql.Selector) {
		step := newTicketStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attendee) predicate.Attendee {
	return predicate.Attendee(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attendee) predicate.Attendee {
	return predicate.Attendee(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attendee) predicate.Attendee {
	return predicate.Attendee(sql.NotPredicates(p))
}
