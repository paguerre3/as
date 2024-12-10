# Use an official Go image as a base
FROM golang:1.23.1-alpine

# Install build-essential for C/C++ development tools
RUN apk update && \
    apk add --no-cache \
    bash \
    build-base && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY cmd/ cmd/
COPY internal/  internal/

# The ./... pattern ensures that all relevant packages within the project structure are installed.
COPY go.mod go.mod
RUN go mod tidy && \
    go build ./... && \
    go install ./...

# Expose the port
EXPOSE 8080

# Run the command when the container starts (then the cgo build constraint is likely disabled; try setting CGO_ENABLED=1 environment variable in your go build step.)
CMD CGO_ENABLED=1 go run ./...