# Stage 1: Build the Go application
FROM golang:1.23.6-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker caching
COPY go.mod  go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary from the main.go in cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-execute ./cmd/main.go

# Stage 2: Run the application using a smaller base image
FROM alpine:latest

# Install any dependencies your binary might need, e.g., if it relies on SSL
RUN apk --no-cache add ca-certificates && \
    update-ca-certificates

WORKDIR /app

# Create non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy the built binary from the builder stage
COPY --from=builder /app/go-execute /app/go-execute

# Set ownership and permissions
RUN chown appuser:appgroup /app/go-execute && \
    chmod 550 /app/go-execute

USER appuser

# Expose port 8000 (or any port your app uses)
EXPOSE 8000

# Run the Go binary
CMD ["/app/go-execute"]
