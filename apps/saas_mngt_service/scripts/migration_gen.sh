#!/bin/bash

# Only run in local development to generate migration files

# Set default values if environment variables are not set
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME_TEST=${DB_NAME_TEST:-test}
DB_SCHEMA=${DB_SCHEMA:-public}
MIGRATION_NAME=${MIGRATION_NAME:-}

# Check if MIGRATION_NAME is provided as a command-line argument
if [ -z "$MIGRATION_NAME" ] && [ -n "$1" ]; then
  MIGRATION_NAME="$1"
fi

if [ -z "$MIGRATION_NAME" ]; then
    echo "Please provide a name for the new migration (without .sql extension):"
    read -r MIGRATION_NAME_INPUT
    MIGRATION_NAME=${MIGRATION_NAME:-$MIGRATION_NAME_INPUT}
fi

# Build the database URL
# It will cover passwordless connection as well
if [ -n "$DB_PASSWORD" ]; then
    DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME_TEST}?search_path=${DB_SCHEMA}&sslmode=disable"
else
    DB_URL="postgres://${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME_TEST}?search_path=${DB_SCHEMA}&sslmode=disable"
fi

atlas migrate diff "$MIGRATION_NAME" \
  --dir "file://sql/migrations" \
  --to "ent://ent/schema" \
  --dev-url "$DB_URL"