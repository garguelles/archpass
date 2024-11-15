#!/bin/bash

migration_name=$1

source ./.env
env=$ENV

echo "env is set to: $env"

dev_url="docker://postgres/16/oxstarter_ephemeral?search_path=public"

atlas migrate diff "$migration_name" --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "$dev_url"
