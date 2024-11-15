// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/garguelles/archpass/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// EventSlug applies equality check predicate on the "event_slug" field. It's identical to EventSlugEQ.
func EventSlug(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventSlug, v))
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEndDate, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldLocation, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldImageURL, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUserID, v))
}

// ContractAddress applies equality check predicate on the "contract_address" field. It's identical to ContractAddressEQ.
func ContractAddress(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldContractAddress, v))
}

// TransactionHash applies equality check predicate on the "transaction_hash" field. It's identical to TransactionHashEQ.
func TransactionHash(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTransactionHash, v))
}

// BlockNumber applies equality check predicate on the "block_number" field. It's identical to BlockNumberEQ.
func BlockNumber(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldBlockNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// ModifiedAt applies equality check predicate on the "modified_at" field. It's identical to ModifiedAtEQ.
func ModifiedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldModifiedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDescription, v))
}

// EventSlugEQ applies the EQ predicate on the "event_slug" field.
func EventSlugEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventSlug, v))
}

// EventSlugNEQ applies the NEQ predicate on the "event_slug" field.
func EventSlugNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEventSlug, v))
}

// EventSlugIn applies the In predicate on the "event_slug" field.
func EventSlugIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEventSlug, vs...))
}

// EventSlugNotIn applies the NotIn predicate on the "event_slug" field.
func EventSlugNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEventSlug, vs...))
}

// EventSlugGT applies the GT predicate on the "event_slug" field.
func EventSlugGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEventSlug, v))
}

// EventSlugGTE applies the GTE predicate on the "event_slug" field.
func EventSlugGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEventSlug, v))
}

// EventSlugLT applies the LT predicate on the "event_slug" field.
func EventSlugLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEventSlug, v))
}

// EventSlugLTE applies the LTE predicate on the "event_slug" field.
func EventSlugLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEventSlug, v))
}

// EventSlugContains applies the Contains predicate on the "event_slug" field.
func EventSlugContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldEventSlug, v))
}

// EventSlugHasPrefix applies the HasPrefix predicate on the "event_slug" field.
func EventSlugHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldEventSlug, v))
}

// EventSlugHasSuffix applies the HasSuffix predicate on the "event_slug" field.
func EventSlugHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldEventSlug, v))
}

// EventSlugEqualFold applies the EqualFold predicate on the "event_slug" field.
func EventSlugEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldEventSlug, v))
}

// EventSlugContainsFold applies the ContainsFold predicate on the "event_slug" field.
func EventSlugContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldEventSlug, v))
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldStartDate, v))
}

// StartDateIsNil applies the IsNil predicate on the "start_date" field.
func StartDateIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldStartDate))
}

// StartDateNotNil applies the NotNil predicate on the "start_date" field.
func StartDateNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldStartDate))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEndDate, v))
}

// EndDateIsNil applies the IsNil predicate on the "end_date" field.
func EndDateIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldEndDate))
}

// EndDateNotNil applies the NotNil predicate on the "end_date" field.
func EndDateNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldEndDate))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDate, v))
}

// DateContains applies the Contains predicate on the "date" field.
func DateContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDate, v))
}

// DateHasPrefix applies the HasPrefix predicate on the "date" field.
func DateHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDate, v))
}

// DateHasSuffix applies the HasSuffix predicate on the "date" field.
func DateHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDate, v))
}

// DateIsNil applies the IsNil predicate on the "date" field.
func DateIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDate))
}

// DateNotNil applies the NotNil predicate on the "date" field.
func DateNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDate))
}

// DateEqualFold applies the EqualFold predicate on the "date" field.
func DateEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDate, v))
}

// DateContainsFold applies the ContainsFold predicate on the "date" field.
func DateContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDate, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldLocation, v))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldImageURL, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUserID, vs...))
}

// ContractAddressEQ applies the EQ predicate on the "contract_address" field.
func ContractAddressEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldContractAddress, v))
}

// ContractAddressNEQ applies the NEQ predicate on the "contract_address" field.
func ContractAddressNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldContractAddress, v))
}

// ContractAddressIn applies the In predicate on the "contract_address" field.
func ContractAddressIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldContractAddress, vs...))
}

// ContractAddressNotIn applies the NotIn predicate on the "contract_address" field.
func ContractAddressNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldContractAddress, vs...))
}

// ContractAddressGT applies the GT predicate on the "contract_address" field.
func ContractAddressGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldContractAddress, v))
}

// ContractAddressGTE applies the GTE predicate on the "contract_address" field.
func ContractAddressGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldContractAddress, v))
}

// ContractAddressLT applies the LT predicate on the "contract_address" field.
func ContractAddressLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldContractAddress, v))
}

// ContractAddressLTE applies the LTE predicate on the "contract_address" field.
func ContractAddressLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldContractAddress, v))
}

// ContractAddressContains applies the Contains predicate on the "contract_address" field.
func ContractAddressContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldContractAddress, v))
}

// ContractAddressHasPrefix applies the HasPrefix predicate on the "contract_address" field.
func ContractAddressHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldContractAddress, v))
}

// ContractAddressHasSuffix applies the HasSuffix predicate on the "contract_address" field.
func ContractAddressHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldContractAddress, v))
}

// ContractAddressIsNil applies the IsNil predicate on the "contract_address" field.
func ContractAddressIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldContractAddress))
}

// ContractAddressNotNil applies the NotNil predicate on the "contract_address" field.
func ContractAddressNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldContractAddress))
}

// ContractAddressEqualFold applies the EqualFold predicate on the "contract_address" field.
func ContractAddressEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldContractAddress, v))
}

// ContractAddressContainsFold applies the ContainsFold predicate on the "contract_address" field.
func ContractAddressContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldContractAddress, v))
}

// TransactionHashEQ applies the EQ predicate on the "transaction_hash" field.
func TransactionHashEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTransactionHash, v))
}

// TransactionHashNEQ applies the NEQ predicate on the "transaction_hash" field.
func TransactionHashNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldTransactionHash, v))
}

// TransactionHashIn applies the In predicate on the "transaction_hash" field.
func TransactionHashIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldTransactionHash, vs...))
}

// TransactionHashNotIn applies the NotIn predicate on the "transaction_hash" field.
func TransactionHashNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldTransactionHash, vs...))
}

// TransactionHashGT applies the GT predicate on the "transaction_hash" field.
func TransactionHashGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldTransactionHash, v))
}

// TransactionHashGTE applies the GTE predicate on the "transaction_hash" field.
func TransactionHashGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldTransactionHash, v))
}

// TransactionHashLT applies the LT predicate on the "transaction_hash" field.
func TransactionHashLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldTransactionHash, v))
}

// TransactionHashLTE applies the LTE predicate on the "transaction_hash" field.
func TransactionHashLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldTransactionHash, v))
}

// TransactionHashContains applies the Contains predicate on the "transaction_hash" field.
func TransactionHashContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldTransactionHash, v))
}

// TransactionHashHasPrefix applies the HasPrefix predicate on the "transaction_hash" field.
func TransactionHashHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldTransactionHash, v))
}

// TransactionHashHasSuffix applies the HasSuffix predicate on the "transaction_hash" field.
func TransactionHashHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldTransactionHash, v))
}

// TransactionHashIsNil applies the IsNil predicate on the "transaction_hash" field.
func TransactionHashIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldTransactionHash))
}

// TransactionHashNotNil applies the NotNil predicate on the "transaction_hash" field.
func TransactionHashNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldTransactionHash))
}

// TransactionHashEqualFold applies the EqualFold predicate on the "transaction_hash" field.
func TransactionHashEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldTransactionHash, v))
}

// TransactionHashContainsFold applies the ContainsFold predicate on the "transaction_hash" field.
func TransactionHashContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldTransactionHash, v))
}

// BlockNumberEQ applies the EQ predicate on the "block_number" field.
func BlockNumberEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldBlockNumber, v))
}

// BlockNumberNEQ applies the NEQ predicate on the "block_number" field.
func BlockNumberNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldBlockNumber, v))
}

// BlockNumberIn applies the In predicate on the "block_number" field.
func BlockNumberIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldBlockNumber, vs...))
}

// BlockNumberNotIn applies the NotIn predicate on the "block_number" field.
func BlockNumberNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldBlockNumber, vs...))
}

// BlockNumberGT applies the GT predicate on the "block_number" field.
func BlockNumberGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldBlockNumber, v))
}

// BlockNumberGTE applies the GTE predicate on the "block_number" field.
func BlockNumberGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldBlockNumber, v))
}

// BlockNumberLT applies the LT predicate on the "block_number" field.
func BlockNumberLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldBlockNumber, v))
}

// BlockNumberLTE applies the LTE predicate on the "block_number" field.
func BlockNumberLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldBlockNumber, v))
}

// BlockNumberContains applies the Contains predicate on the "block_number" field.
func BlockNumberContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldBlockNumber, v))
}

// BlockNumberHasPrefix applies the HasPrefix predicate on the "block_number" field.
func BlockNumberHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldBlockNumber, v))
}

// BlockNumberHasSuffix applies the HasSuffix predicate on the "block_number" field.
func BlockNumberHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldBlockNumber, v))
}

// BlockNumberIsNil applies the IsNil predicate on the "block_number" field.
func BlockNumberIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldBlockNumber))
}

// BlockNumberNotNil applies the NotNil predicate on the "block_number" field.
func BlockNumberNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldBlockNumber))
}

// BlockNumberEqualFold applies the EqualFold predicate on the "block_number" field.
func BlockNumberEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldBlockNumber, v))
}

// BlockNumberContainsFold applies the ContainsFold predicate on the "block_number" field.
func BlockNumberContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldBlockNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// ModifiedAtEQ applies the EQ predicate on the "modified_at" field.
func ModifiedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldModifiedAt, v))
}

// ModifiedAtNEQ applies the NEQ predicate on the "modified_at" field.
func ModifiedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldModifiedAt, v))
}

// ModifiedAtIn applies the In predicate on the "modified_at" field.
func ModifiedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldModifiedAt, vs...))
}

// ModifiedAtNotIn applies the NotIn predicate on the "modified_at" field.
func ModifiedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldModifiedAt, vs...))
}

// ModifiedAtGT applies the GT predicate on the "modified_at" field.
func ModifiedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldModifiedAt, v))
}

// ModifiedAtGTE applies the GTE predicate on the "modified_at" field.
func ModifiedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldModifiedAt, v))
}

// ModifiedAtLT applies the LT predicate on the "modified_at" field.
func ModifiedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldModifiedAt, v))
}

// ModifiedAtLTE applies the LTE predicate on the "modified_at" field.
func ModifiedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldModifiedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTickets applies the HasEdge predicate on the "tickets" edge.
func HasTickets() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TicketsTable, TicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTicketsWith applies the HasEdge predicate on the "tickets" edge with a given conditions (other predicates).
func HasTicketsWith(preds ...predicate.Ticket) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newTicketsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttendees applies the HasEdge predicate on the "attendees" edge.
func HasAttendees() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttendeesTable, AttendeesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttendeesWith applies the HasEdge predicate on the "attendees" edge with a given conditions (other predicates).
func HasAttendeesWith(preds ...predicate.Attendee) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newAttendeesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}
