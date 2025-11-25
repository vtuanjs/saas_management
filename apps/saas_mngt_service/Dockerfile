# Stage 1: Build the Go application
FROM golang:1.24.5-alpine AS builder

WORKDIR /saas_management

# Copy shared packages
COPY packages/go packages/go

# Copy service source
COPY apps/saas_mngt_service apps/saas_mngt_service

WORKDIR /saas_management/apps/saas_mngt_service

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main ./cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /saas_management/apps/saas_mngt_service
	
COPY --from=builder /saas_management/apps/saas_mngt_service/.build/main .

CMD ["./main"]