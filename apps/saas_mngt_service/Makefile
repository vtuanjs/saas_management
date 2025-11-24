.EXPORT_ALL_VARIABLES:
DOCKER_BUILDKIT=1
COMPOSE_DOCKER_CLI_BUILD=1
COMPOSE_IGNORE_PULL_FAILURES=1
ENV=local
GO_VERSION ?= "1.24.5"
BUF_TOKEN ?= $(shell test -f ../../secrets/buf_token && cat ../../secrets/buf_token || echo "")
MAKEFLAGS += --no-print-directory
TOOL ?=
DEBUG ?=
DB_DOCKER_SV_NAME ?= postgres1
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_SCHEMA ?= public
DB_NAME ?= saas_mngt

.PHONY: setup buf-update buf-gen buf-lint buf-breaking env-decrypt env-encrypt build up down debug check migrate-up migrate-down migrate-add migrate-remove migrate-status migrate-hash migrate-gen migrate-rebase ad-hoc-up ad-hoc-down ad-hoc-add seed-up seed-add seed-down mod-tidy %

# gen
gen-proto:
	buf generate

gen-mock:
	@echo "\033[0;32m Generating mocks...\033[0m"
	go generate ./internal/...
	@echo "\033[0;32m Build after code generation...\033[0m"
	make build

gen-query:
	@echo "\033[0;32m Generating sqlc query...\033[0m"
	@sqlc generate
# 	@echo "\033[0;32m Build after code generation...\033[0m"
#	You can manually build after generation if needed
# 	@make build

gen-migration:
	@echo "Adding new migration..."
	MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) bash ./scripts/migration_add.sh

gen-seed:
	@echo "Adding new seed migration..."
	MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) bash ./scripts/seed_add.sh

gen-ad-hoc:
	@echo "Adding new ad-hoc migration..."
	MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) bash ./scripts/ad_hoc_add.sh

# buf commands
buf-update:
	buf dep update

## Another way: make gen-buf
buf-gen:
	buf generate

buf-lint:
	buf lint

buf-breaking:
	buf breaking --against "../../.git#branch=develop,subdir=proto"

# env
env-decrypt:
	TOOL=ENV_DECRYPT docker compose up saas_mngt_service

env-encrypt:
	TOOL=ENV_ENCRYPT docker compose up saas_mngt_service

# app build and run
build-no-cache:
	@echo "Building without cache..."
	docker compose build --no-cache saas_mngt_service
	TOOL=BUILD docker compose up saas_mngt_service --remove-orphans

build:
	@echo "Building..."
	TOOL=BUILD docker compose up saas_mngt_service --remove-orphans

up:
	@echo "Running in normal mode..."
	DEBUG=false docker compose up saas_mngt_service -d --remove-orphans

down:
	@echo "Stopping..."
	docker compose down

debug:
	@echo "Running in debug mode..."
	DEBUG=true docker compose up saas_mngt_service

# golangci-lint
check:
	@echo "Running check..."
	@make buf-lint -C ../../ 
	@TOOL=CHECK docker compose up saas_mngt_service_tool

# database migrations
migrate-up:
	@echo "Running migrate up..."
	bash ./scripts/migration_up.sh
	TOOL=DUMP_SQL docker compose up saas_mngt_service

migrate-down:
	@echo "Running migrate down..."
	bash ./scripts/migration_down.sh
	TOOL=DUMP_SQL docker compose up saas_mngt_service

## Another way: make migrate-add
migrate-add:
	@echo "Adding new migration..."
	MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) bash ./scripts/migration_add.sh 

migrate-fix:
	@echo "Fixing last migration..."
	bash ./scripts/migration_fix.sh

ad-hoc-up:
	@echo "Running ad-hoc migrate apply..."
	bash ./scripts/ad_hoc_up.sh

ad-hoc-down:
	@echo "Running ad-hoc migrate down..."
	bash ./scripts/ad_hoc_down.sh

## Another way: make gen-ad-hoc
ad-hoc-add:
	@echo "Adding new ad-hoc migration..."
	MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) bash ./scripts/ad_hoc_add.sh

seed-up:
	@echo "Running seed migrate apply..."
	bash ./scripts/seed_up.sh

## Another way: make gen-seed
seed-add:
	@echo "Adding new seed migration..."
	MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) bash ./scripts/seed_add.sh

seed-down:
	@echo "Running seed migrate down..."
	bash ./scripts/seed_down.sh

%:
	@: