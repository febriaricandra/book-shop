# Stage 1: Build the Go app
FROM golang:1.22 AS builder

WORKDIR /app

# Copy the application code into the container
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app/main.go

# Stage 2: Create a smaller image to run the app
FROM alpine:latest

WORKDIR /app

# Copy the binary and the .env file from the builder stage
COPY --from=builder /app/main .
COPY .env .env

# Install bash (optional, for debugging or shell-based scripts)
RUN apk add --no-cache bash

# Ensure the binary has execution permissions
RUN chmod +x main

# Expose the application port
EXPOSE 8080

# Set the entry point
CMD ["./main"]
