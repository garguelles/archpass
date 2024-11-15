// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/garguelles/archpass/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// WalletAddress applies equality check predicate on the "wallet_address" field. It's identical to WalletAddressEQ.
func WalletAddress(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWalletAddress, v))
}

// Bio applies equality check predicate on the "bio" field. It's identical to BioEQ.
func Bio(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBio, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// WalletAddressEQ applies the EQ predicate on the "wallet_address" field.
func WalletAddressEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWalletAddress, v))
}

// WalletAddressNEQ applies the NEQ predicate on the "wallet_address" field.
func WalletAddressNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldWalletAddress, v))
}

// WalletAddressIn applies the In predicate on the "wallet_address" field.
func WalletAddressIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldWalletAddress, vs...))
}

// WalletAddressNotIn applies the NotIn predicate on the "wallet_address" field.
func WalletAddressNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldWalletAddress, vs...))
}

// WalletAddressGT applies the GT predicate on the "wallet_address" field.
func WalletAddressGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldWalletAddress, v))
}

// WalletAddressGTE applies the GTE predicate on the "wallet_address" field.
func WalletAddressGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldWalletAddress, v))
}

// WalletAddressLT applies the LT predicate on the "wallet_address" field.
func WalletAddressLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldWalletAddress, v))
}

// WalletAddressLTE applies the LTE predicate on the "wallet_address" field.
func WalletAddressLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldWalletAddress, v))
}

// WalletAddressContains applies the Contains predicate on the "wallet_address" field.
func WalletAddressContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldWalletAddress, v))
}

// WalletAddressHasPrefix applies the HasPrefix predicate on the "wallet_address" field.
func WalletAddressHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldWalletAddress, v))
}

// WalletAddressHasSuffix applies the HasSuffix predicate on the "wallet_address" field.
func WalletAddressHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldWalletAddress, v))
}

// WalletAddressEqualFold applies the EqualFold predicate on the "wallet_address" field.
func WalletAddressEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldWalletAddress, v))
}

// WalletAddressContainsFold applies the ContainsFold predicate on the "wallet_address" field.
func WalletAddressContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldWalletAddress, v))
}

// BioEQ applies the EQ predicate on the "bio" field.
func BioEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBio, v))
}

// BioNEQ applies the NEQ predicate on the "bio" field.
func BioNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBio, v))
}

// BioIn applies the In predicate on the "bio" field.
func BioIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldBio, vs...))
}

// BioNotIn applies the NotIn predicate on the "bio" field.
func BioNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBio, vs...))
}

// BioGT applies the GT predicate on the "bio" field.
func BioGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldBio, v))
}

// BioGTE applies the GTE predicate on the "bio" field.
func BioGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBio, v))
}

// BioLT applies the LT predicate on the "bio" field.
func BioLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldBio, v))
}

// BioLTE applies the LTE predicate on the "bio" field.
func BioLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBio, v))
}

// BioContains applies the Contains predicate on the "bio" field.
func BioContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldBio, v))
}

// BioHasPrefix applies the HasPrefix predicate on the "bio" field.
func BioHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldBio, v))
}

// BioHasSuffix applies the HasSuffix predicate on the "bio" field.
func BioHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldBio, v))
}

// BioEqualFold applies the EqualFold predicate on the "bio" field.
func BioEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldBio, v))
}

// BioContainsFold applies the ContainsFold predicate on the "bio" field.
func BioContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldBio, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newEventsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttendeeTickets applies the HasEdge predicate on the "attendee_tickets" edge.
func HasAttendeeTickets() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttendeeTicketsTable, AttendeeTicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttendeeTicketsWith applies the HasEdge predicate on the "attendee_tickets" edge with a given conditions (other predicates).
func HasAttendeeTicketsWith(preds ...predicate.Attendee) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAttendeeTicketsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
