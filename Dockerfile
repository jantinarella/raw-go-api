# Build stage
FROM golang:1.21-alpine AS builder

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files first (for better layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application - matches your local build command
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o raw-api ./cmd

# Final stage - minimal image
FROM alpine:latest

# Create non-root user for security
RUN adduser -D -s /bin/sh appuser

WORKDIR /root/

# Copy the binary from builder stage (matches your output name 'raw-api')
COPY --from=builder /app/raw-api .

# Change ownership to non-root user
RUN chown appuser:appuser raw-api

# Switch to non-root user
USER appuser

# Expose port 8080 (matches your app's port)
EXPOSE 8080

# Run the application
CMD ["./raw-api"]