# syntax=docker/dockerfile:1

# Godbo app Dockerfile.
# Use the latest official Golang image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

COPY ./src .

RUN ls -la
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main

EXPOSE 8080

# Command to run the executable
CMD ["./main"]
