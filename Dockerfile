# Use the official Golang image as a base image
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the dependencies into the container
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Expose the port the app runs on
EXPOSE 9090

# Command to run the executable
CMD ["go", "run", "main.go"]
