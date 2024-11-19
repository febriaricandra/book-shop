# Dockerfile

# Stage 1: Build the Go app
FROM golang:1.22.5 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy code from current directory into container
COPY . .

# Build the Go app into an executable named "main"
RUN go build -o main ./cmd/app/main.go

# Stage 2: Create a smaller image to run the app
FROM alpine:latest

# Set working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 (for HTTP traffic)
EXPOSE 8080

# Run the built executable
CMD ["./main"]