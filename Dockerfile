# Use the official golang image as base
FROM --platform=linux/amd64 golang:latest AS builder

MAINTAINER "01himanshugautam@gmail.com"

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files for downloading dependencies
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download
RUN go mod verify

# Copy the entire project
COPY . .

# Build the Go app with a statically linked binary
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o go_test .

# Use a lightweight base image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/go_test .

# Copy template files (if any)
COPY public/ /app/public

# Command to run the executable
CMD ["./go_test"]
