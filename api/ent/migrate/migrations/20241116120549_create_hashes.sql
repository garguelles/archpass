-- Modify "events" table
ALTER TABLE "events" ADD COLUMN "event_hash" character varying NULL;
-- Modify "tickets" table
ALTER TABLE "tickets" ADD COLUMN "ticket_hash" character varying NULL;
