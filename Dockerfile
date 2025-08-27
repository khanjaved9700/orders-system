# 1. Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Cache go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o orders ./cmd/main.go

# 2. Run stage (smaller image)
FROM alpine:3.18

# Install CA certificates for HTTPS calls (optional)
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/orders .

# Copy config.yaml if needed
COPY --from=builder /app/config/config.yaml ./config.yaml

# Expose HTTP port
EXPOSE 8080

# Command to run
CMD ["./orders"]
