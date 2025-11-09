#!/bin/bash

MIGRATION_NAME=${MIGRATION_NAME:-}

export GOOSE_DRIVER=postgres
export GOOSE_MIGRATION_DIR=./sql/migrations

goose fix