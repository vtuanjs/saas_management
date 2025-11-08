.EXPORT_ALL_VARIABLES:

DOCKER_BUILDKIT=1
COMPOSE_DOCKER_CLI_BUILD=1
COMPOSE_IGNORE_PULL_FAILURES=1
ENV=local
GO_VERSION ?= "1.24.5"
BUF_TOKEN ?= $(shell test -f ../../secrets/buf_token && cat ../../secrets/buf_token || echo "")
MAKEFLAGS += --no-print-directory
SOPS_PGP_FP = $(shell cat ../../secrets/sops_public_key)
TOOL ?=

setup:
	@chmod +x ./*.sh
	@chmod +x ./scripts/*.sh

# buf
buf-update:
	buf dep update

buf-gen:
	buf generate

buf-lint:
	buf lint

buf-breaking:
	buf breaking --against "../../.git#branch=develop,subdir=proto"

# env
env-decrypt:
	export SOPS_PGP_FP=$(SOPS_PGP_FP)
	sops decrypt --input-type dotenv --output-type dotenv .env.enc > .env

env-encrypt:
	export SOPS_PGP_FP=$(SOPS_PGP_FP)
	sops encrypt --input-type dotenv --output-type dotenv .env > .env.enc

# bazel
build:
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

# git
commit:
	@echo "Committing changes..."
	@bash ./commit.sh

# golangci-lint
check:
	@echo "Running check..."
	@make buf-lint -C ../../ 
	@TOOL=CHECK docker compose up saas_mngt_service_tool

migrate-up:
	@echo "Running migrate up..."
	bash ./scripts/migration_up.sh

migrate-down:
	@echo "Running migrate down..."
	bash ./scripts/migration_down.sh

migrate-add:
	@echo "Adding new migration..."
	@TOOL=MIGRATE_ADD MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) docker compose up saas_mngt_service_tool

migrate-remove:
	@echo "Removing last migration..."
	@TOOL=MIGRATE_REMOVE MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) docker compose up saas_mngt_service_tool

migrate-status:
	@echo "Checking migration status..."
	@TOOL=MIGRATE_STATUS docker compose up saas_mngt_service_tool

migrate-hash:
	@echo "Hashing migration files..."
	@TOOL=MIGRATE_HASH docker compose up saas_mngt_service_tool

migrate-gen:
	@echo "Generating migration files..."
	@TOOL=MIGRATE_GEN MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) docker compose up saas_mngt_service_tool

migrate-rebase:
	@echo "Rebasing migration files..."
	@TOOL=MIGRATE_REBASE MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) docker compose up saas_mngt_service_tool

ad-hoc-up:
	@echo "Running ad-hoc migrate apply..."
	@TOOL=AD_HOC_UP docker compose up saas_mngt_service_tool

ad-hoc-down:
	@echo "Running ad-hoc migrate down..."
	@TOOL=AD_HOC_DOWN docker compose up saas_mngt_service_tool

ad-hoc-add:
	@echo "Adding new ad-hoc migration..."
	@TOOL=AD_HOC_ADD MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) docker compose up saas_mngt_service_tool

seed-up:
	@echo "Running seed migrate apply..."
	@TOOL=SEED_UP docker compose up saas_mngt_service_tool

seed-add:
	@echo "Adding new seed migration..."
	@TOOL=SEED_ADD MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) docker compose up saas_mngt_service_tool

seed-down:
	@echo "Running seed migrate down..."
	@TOOL=SEED_DOWN docker compose up saas_mngt_service_tool

#### FAST GEN - You can use "make gen" instead if you lazy ####
mod-tidy:
	@echo "Running go mod tidy..."
	TOOL=MOD_TIDY docker compose up saas_mngt_service_tool

wire:
	@echo "Generating wire..."
	wire ./internal/wire

%:
	@: