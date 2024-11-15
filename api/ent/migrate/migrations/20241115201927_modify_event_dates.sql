-- Modify "events" table
ALTER TABLE "events" ALTER COLUMN "start_date" DROP NOT NULL, ALTER COLUMN "end_date" DROP NOT NULL, ADD COLUMN "date" character varying NULL;
