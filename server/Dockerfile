# Use an official Golang runtime as a base image
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /server

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main

# Final stage: Create a lightweight container with the Go binary
FROM alpine:latest

# Set the working directory for the final image
WORKDIR /server

# Copy the Go binary from the builder stage
COPY --from=builder /server/main .
COPY --from=builder /server/templates /server/templates
COPY --from=builder /server/assets /server/assets
COPY --from=builder /server/.env /server/.env

# Expose the port (adjust as needed)
EXPOSE 80

# Command to run the executable
CMD ["./main"]
