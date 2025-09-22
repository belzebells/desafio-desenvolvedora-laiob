# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# Copy go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with verbose output
RUN go build -v -o main .

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add sqlite ca-certificates

WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/main .

# Make sure the binary is executable
RUN chmod +x ./main

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
