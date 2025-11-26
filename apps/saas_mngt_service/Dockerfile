# Need run from project root. See <root>/cicd/build.sh
# Stage 1: Build the Go application
FROM golang:1.24.5-alpine AS builder

WORKDIR /saas_management/apps/saas_mngt_service

# Cached layer
COPY apps/saas_mngt_service/go.mod apps/saas_mngt_service/go.sum ./

COPY packages/go/go.mod packages/go/go.sum /saas_management/packages/go/

RUN go mod download

# Dynamic layer
COPY packages/go /saas_management/packages/go

COPY apps/saas_mngt_service /saas_management/apps/saas_mngt_service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main ./cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /saas_management/apps/saas_mngt_service
	
COPY --from=builder /saas_management/apps/saas_mngt_service/.build/main .

CMD ["./main"]