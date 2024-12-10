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

# CGO build constraint is likely disabled; setting CGO_ENABLED=1 environment for C/C++ libs
ENV CGO_ENABLED=1

# Tidy up dependencies and build the binary for the main entry point
RUN go mod tidy && \
    go build -o /app/main ./cmd/7_9_damaged_spaceship/main.go

# Expose the port
EXPOSE 8080

# Run the command when the container starts
CMD ["/app/main"]