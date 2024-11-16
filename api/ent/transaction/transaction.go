// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventType holds the string denoting the event_type field in the database.
	FieldEventType = "event_type"
	// FieldWalletAddress holds the string denoting the wallet_address field in the database.
	FieldWalletAddress = "wallet_address"
	// FieldTransactionHash holds the string denoting the transaction_hash field in the database.
	FieldTransactionHash = "transaction_hash"
	// FieldBlockNumber holds the string denoting the block_number field in the database.
	FieldBlockNumber = "block_number"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldEventType,
	FieldWalletAddress,
	FieldTransactionHash,
	FieldBlockNumber,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Transaction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEventType orders the results by the event_type field.
func ByEventType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventType, opts...).ToFunc()
}

// ByWalletAddress orders the results by the wallet_address field.
func ByWalletAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWalletAddress, opts...).ToFunc()
}

// ByTransactionHash orders the results by the transaction_hash field.
func ByTransactionHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionHash, opts...).ToFunc()
}

// ByBlockNumber orders the results by the block_number field.
func ByBlockNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlockNumber, opts...).ToFunc()
}
