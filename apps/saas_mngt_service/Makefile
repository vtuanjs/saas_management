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
DEBUG ?=

.PHONY: setup buf-update buf-gen buf-lint buf-breaking env-decrypt env-encrypt build up down debug check migrate-up migrate-down migrate-add migrate-remove migrate-status migrate-hash migrate-gen migrate-rebase ad-hoc-up ad-hoc-down ad-hoc-add seed-up seed-add seed-down mod-tidy wire %

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

# gen wire + mock
gen:
	@echo "Generating wire..."
	wire ./wire
	@echo "Generating mocks..."
	go generate ./internal/...
	@echo "Build after code generation..."
	make build

# env
env-decrypt:
	export SOPS_PGP_FP=$(SOPS_PGP_FP)
	sops decrypt --input-type dotenv --output-type dotenv .env.enc > .env

env-encrypt:
	export SOPS_PGP_FP=$(SOPS_PGP_FP)
	sops encrypt --input-type dotenv --output-type dotenv .env > .env.enc

# bazel
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

migrate-down:
	@echo "Running migrate down..."
	bash ./scripts/migration_down.sh

migrate-add:
	@echo "Adding new migration..."
	MIGRATION_NAME=$(word 2,$(MAKECMDGOALS)) bash ./scripts/migration_add.sh 

migrate-fix:
	@echo "Fixing last migration..."
	bash ./scripts/migration_fix.sh

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

%:
	@: