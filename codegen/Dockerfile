# Use the official Golang image as the base image for the build stage
FROM golang:1.16 as builder 

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests into the container
COPY go.mod go.sum ./

# Download the Go modules for the application
RUN go mod download 

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main . 

# Use a minimal base image for the final image 
FROM alpine:latest

# Set the working directory inside the container to /root 
WORKDIR /root/

# Copy the binary from the builder stage into the final image 
COPY --from=builder /app/main . 

# Expose the application port on the container
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./main"] 