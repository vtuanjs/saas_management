#!/bin/bash
set -e

# Function to setup TimescaleDB extension in specific databases
setup_timescaledb_in_db() {
    local database=$1
    echo "Setting up TimescaleDB extension in database: $database"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$database" <<-EOSQL
        CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
EOSQL
}

# Setup TimescaleDB in multiple databases
if [ -n "$POSTGRES_MULTIPLE_DATABASES" ]; then
    echo "Setting up TimescaleDB in multiple databases: $POSTGRES_MULTIPLE_DATABASES"
    for db in $(echo $POSTGRES_MULTIPLE_DATABASES | tr ',' ' '); do
        setup_timescaledb_in_db $db
    done
    echo "TimescaleDB extension setup completed for all databases"
fi
