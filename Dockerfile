# Use the official Golang image as a build stage
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o books-service main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/books-service .

# Expose the application port
EXPOSE 8080

# Set environment variables (optional if using .env file)
ENV KAFKA_BROKERS=kafka:9092
ENV POSTGRES_HOST=postgres
ENV POSTGRES_PORT=_
