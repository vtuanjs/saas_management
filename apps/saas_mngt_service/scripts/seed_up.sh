#!/bin/bash

DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_SCHEMA=${DB_SCHEMA:-public}
DB_NAME=${DB_NAME:-saas_mngt}
MIGRATION_NAME=${MIGRATION_NAME:-}

export GOOSE_DRIVER=postgres
export GOOSE_MIGRATION_DIR=./sql/seeds
export GOOSE_TABLE=goose_seed_revisions

# Build the database URL
# It will cover passwordless connection as well
if [ -n "$DB_PASSWORD" ]; then
    export GOOSE_DBSTRING="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}&sslmode=disable"
else
    export GOOSE_DBSTRING="postgres://${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}&sslmode=disable"
fi

goose up -allow-missing