# Stage 1: Build the Go binary
FROM golang:1.23.6-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o main ./cmd/main.go

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose necessary ports (if your app listens on a port)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
