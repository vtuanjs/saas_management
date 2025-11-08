-- This script runs after database creation
-- You can add any initial schema setup here

-- Example: Create TimescaleDB extension in each database
-- This will be executed against the default POSTGRES_DB

-- Enable TimescaleDB extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
