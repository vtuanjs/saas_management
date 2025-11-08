#!/bin/bash

# Set default values if environment variables are not set
MIGRATION_NAME=${MIGRATION_NAME:-}

# Check if MIGRATION_NAME is provided as a command-line argument
if [ -z "$MIGRATION_NAME" ] && [ -n "$1" ]; then
  MIGRATION_NAME="$1"
fi

if [ -z "$MIGRATION_NAME" ]; then
    echo "Please input migration name to rebase (e.g., 20230915010101_initial):"
    read -r MIGRATION_NAME_INPUT
    MIGRATION_NAME=${MIGRATION_NAME_INPUT}
fi

# Ensure MIGRATION_NAME ends with .sql
if [[ "$MIGRATION_NAME" != *.sql ]]; then
  MIGRATION_NAME="${MIGRATION_NAME}.sql"
fi

atlas migrate rebase \
  --dir "file://sql/migrations" \
  "$MIGRATION_NAME"
