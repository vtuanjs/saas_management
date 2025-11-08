#!/bin/bash

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

atlas migrate new \
  --dir "file://sql/migrations" \
  "$MIGRATION_NAME"