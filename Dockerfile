# Use an official Go image as a base
FROM golang:1.23.1-alpine

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

# Run the command when the container starts
CMD go run ./...