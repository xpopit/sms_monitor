# Start by specifying the base image
FROM golang:1.22.4-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go install github.com/air-verse/air@latest

# Copy the source code into the container
COPY . .

# # Build the Go app
# RUN go build -o main .

# # Start a new stage from scratch
# FROM alpine:latest

# # Copy the Pre-built binary file from the previous stage
# COPY --from=builder /app/main /main
# COPY .env /app/.env

# Expose port 8080 to the outside world
EXPOSE 13000

# # Command to run the executable
# ENTRYPOINT ["/main"]