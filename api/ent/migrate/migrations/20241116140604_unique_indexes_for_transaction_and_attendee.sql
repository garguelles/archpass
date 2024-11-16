-- Create index "attendee_token_id_ticket_id_user_id" to table: "attendees"
CREATE UNIQUE INDEX "attendee_token_id_ticket_id_user_id" ON "attendees" ("token_id", "ticket_id", "user_id");
-- Create index "transaction_event_type_wallet_address_transaction_hash" to table: "transactions"
CREATE UNIQUE INDEX "transaction_event_type_wallet_address_transaction_hash" ON "transactions" ("event_type", "wallet_address", "transaction_hash");
