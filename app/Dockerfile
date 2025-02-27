FROM golang:1.24.0-bullseye AS builder

RUN useradd -u 1001 nonroot

WORKDIR /app

COPY go.mod ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o server

# Use a minimal scratch image as the base
FROM scratch 

# Copy /etc/passwd for non-root user information
COPY --from=builder /etc/passwd /etc/passwd

# Copy the compiled application binary
COPY --from=builder /app/server /server

# Use the non-root user
USER nonroot

# Expose the application port
EXPOSE 1343

# Define the command to run the application
CMD ["/server"]
