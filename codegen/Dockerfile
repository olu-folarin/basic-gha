# Use the official Golang image as the base image for the build stage
FROM golang:1.22.7-alpine3.19 as builder 

# Set the working directory inside the container
WORKDIR /app

# Copy only dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final image 
FROM alpine:latest

# Set the working directory inside the container to /root 
WORKDIR /root/

# Copy the binary from the builder stage into the final image 
COPY --from=builder /app/main .

# Command to run the application when the container starts
CMD ["./main"] 