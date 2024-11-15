-- Modify "events" table
ALTER TABLE "events" ADD COLUMN "start_date" timestamptz NOT NULL, ADD COLUMN "end_date" timestamptz NOT NULL;
