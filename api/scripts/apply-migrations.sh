#!/bin/bash

source ./.env

database_url=$DATABASE_URL
env=$ENV

echo "env is set to: $env"

if [ "$env" = "local" ]; then
  docker compose run api atlas migrate apply --dir "file://ent/migrate/migrations" --url "$database_url"
else
  atlas migrate apply --dir "file://ent/migrate/migrations" --url "$database_url"
fi
