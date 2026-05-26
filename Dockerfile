# Stage 1: Builder
FROM golang:1.26.1-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary (CGO disabled for alpine compatibility)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o arena-api ./cmd/api

# Stage 2: Runtime (lightweight)
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/arena-api .

# Copy migrations (IMPORTANTE!)
COPY --from=builder /app/migrations ./migrations/

# Copy .env (optional, será sobrescrito por docker-compose)
COPY .env .

# Expose port
EXPOSE 8080

# Run binary
CMD ["./arena-api"]
