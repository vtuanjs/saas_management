.EXPORT_ALL_VARIABLES:

DOCKER_BUILDKIT=1
COMPOSE_DOCKER_CLI_BUILD=1
COMPOSE_IGNORE_PULL_FAILURES=1
GO_VERSION ?= "1.24.5"
BUF_TOKEN ?= $(shell test -f ./secrets/buf_token && cat ./secrets/buf_token || echo "")
MAKEFLAGS += --no-print-directory
SOPS_PGP_FP = $(shell cat ./secrets/sops_public_key)

setup:
	@echo "Installing required Go tools..."
	go install github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0
	go install go.uber.org/mock/mockgen@v0.6.0
	go install github.com/pressly/goose/v3/cmd/goose@v3.25.0
	go install github.com/getsops/sops/v3/cmd/sops@v3.9.0
	@echo "Importing SOPS PGP key..."
	gpg --import ./secrets/sops_private_key.asc
	@echo "Initializing shared networks"
	@docker network create --driver=bridge local-monorepo || echo "Network 'local-monorepo' already exists - skipping creation"
	@echo "Changing file permissions..."
	chmod +x ./commit.sh
	make -C ./apps/saas_mngt_service setup

# infrastructure
infra-up:
	@docker compose -f ./environment/docker-compose.yml up -d --remove-orphans

infra-down:
	@docker compose -f ./environment/docker-compose.yml down
