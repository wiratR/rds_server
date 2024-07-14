# Start from the official Go 1.22 base image
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy