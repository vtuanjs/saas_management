.EXPORT_ALL_VARIABLES:

DOCKER_BUILDKIT=1
COMPOSE_DOCKER_CLI_BUILD=1
COMPOSE_IGNORE_PULL_FAILURES=1
ENV=local
GO_VERSION ?= "1.24.5"
BUF_TOKEN ?= $(shell test -f ./tools/buf/buf_token && cat ./tools/buf/buf_token || echo "")
SCHEMA ?= 
TOOL ?=
DEBUG ?= false
MAKEFLAGS += --no-print-directory
SOPS_PGP_FP = 49031300F4AF150110DCECAD347C46C2BFE6D611

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
	make -C ./apps/saas_mngt_service setup


# If "Error response from daemon: pull access denied for go-local-1.24.5, repository does not exist or may require 'docker login': denied: requested access to the resource is denied"
# Just use it to fix the issue
fix:
	@echo "Initializing project..."
	@echo "Building docker image: go-local-1.24.5 ..."
	@docker build -t go-local-1.24.5:latest -f ./tools/go.Dockerfile . --no-cache
	@echo "Initializing shared networks"
	@docker network create --driver=bridge local-monorepo || echo "Network 'local-monorepo' already exists - skipping creation"

# infrastructure
infra-up:
	@docker compose -f ./environment/docker-compose.yml up -d --remove-orphans

infra-down:
	@docker compose -f ./environment/docker-compose.yml down
