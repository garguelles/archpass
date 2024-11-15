#!/bin/bash

migration_name=$1

source ./.env

atlas migrate new "$migration_name" --dir "file://ent/migrate/migrations"
