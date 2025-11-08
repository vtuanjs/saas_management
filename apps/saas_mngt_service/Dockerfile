FROM golang:1.24.5

RUN go install github.com/cespare/reflex@v0.3.1

RUN go install github.com/go-delve/delve/cmd/dlv@v1.25.1

RUN go install github.com/bazelbuild/bazelisk@v1.26.0

# Install GPG and dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    gnupg \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN ln -sf $(which bazelisk) /go/bin/bazel