-- Modify "attendees" table
ALTER TABLE "attendees" ADD COLUMN "token_id" bigint NOT NULL, ADD COLUMN "transaction_hash" character varying NOT NULL, ADD COLUMN "block_number" bigint NOT NULL, ADD COLUMN "created_at" timestamptz NOT NULL, ADD CONSTRAINT "attendees_tickets_attendees" FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Create index "attendees_ticket_id_key" to table: "attendees"
CREATE UNIQUE INDEX "attendees_ticket_id_key" ON "attendees" ("ticket_id");
