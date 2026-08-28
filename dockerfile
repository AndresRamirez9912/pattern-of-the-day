# Build the binary
FROM golang:1.26-alpine AS builder

# Define the working directory for the application
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code into the container
COPY . .

# Static build: no CGO, so the binary has no libc dependency and can run on
# scratch. If llm.base_url ever points to an https:// endpoint, this stage's
# /etc/ssl/certs/ca-certificates.crt would need to be copied below too.
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o pattern ./cmd/pattern-of-the-day/main.go

# Runtime stage: only the binary and the config it reads at startup, nothing else
FROM scratch

WORKDIR /app

COPY --from=builder /app/pattern ./pattern
COPY config.yaml ./config.yaml

# Set the entry point for the container
ENTRYPOINT ["./pattern", "start-server"]

# Expose the port the application will run on
EXPOSE 8080


