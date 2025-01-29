# Use a specific version of Golang Alpine
FROM golang:1.21-alpine3.19 AS builder

# Set the working directory
WORKDIR /app

# Install git and ca-certificates
RUN apk add --no-cache git ca-certificates

# Copy only dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build with security flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main .

# Use a specific Alpine version for the final image
FROM alpine:3.19

# Add non-root user
RUN adduser -D appuser

# Install CA certificates for HTTPS
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy only the binary
COPY --from=builder /app/main .

# Set ownership
RUN chown -R appuser:appuser /app

# Use non-root user
USER appuser

# Set environment variables
ENV JWT_SECRET_KEY=""
ENV USER_ROLE=""

# Command to run the application
CMD ["./main"]