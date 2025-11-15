FROM golang:1.24.5

RUN go install github.com/getsops/sops/v3/cmd/sops@v3.9.0

RUN go install github.com/cespare/reflex@v0.3.1

RUN go install github.com/go-delve/delve/cmd/dlv@v1.25.1

RUN go install github.com/bazelbuild/bazelisk@v1.27.0

# Install dependencies and PostgreSQL client
RUN apt-get update && \
	apt-get install -y --no-install-recommends \
	gnupg \
	ca-certificates \
	curl \
	lsb-release && \
	# Add PostgreSQL official repository
	curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor -o /usr/share/keyrings/postgresql-keyring.gpg && \
	echo "deb [signed-by=/usr/share/keyrings/postgresql-keyring.gpg] http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list && \
	apt-get update && \
	apt-get install -y --no-install-recommends \
	postgresql-client-18 && \
	rm -rf /var/lib/apt/lists/*

RUN ln -sf $(which bazelisk) /go/bin/bazel