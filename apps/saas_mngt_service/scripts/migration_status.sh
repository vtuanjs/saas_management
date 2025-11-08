#!/bin/bash

# Set default values if environment variables are not set
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME=${DB_NAME:-saas}
DB_SCHEMA=${DB_SCHEMA:-public}

# Build the database URL
# It will cover passwordless connection as well
if [ -n "$DB_PASSWORD" ]; then
    DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}&sslmode=disable"
else
    DB_URL="postgres://${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}&sslmode=disable"
fi

atlas migrate status \
  --dir "file://sql/migrations" \
  --url "$DB_URL"