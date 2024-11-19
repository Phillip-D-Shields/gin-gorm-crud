# First stage: Build the application
FROM golang:1.23-alpine AS builder

WORKDIR /build

# Install build dependencies
ENV CGO_ENABLED=1
RUN apk update && \
    apk add --no-cache \
    build-base \
    gcc \
    wget \
    git

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application
RUN go build -v -o app .

# Second stage: Create the final image
FROM alpine:3.18

WORKDIR /app

# Install runtime dependencies
RUN apk update && \
    apk add --no-cache \
    sqlite \
    sqlite-libs

# Copy the binary from builder
COPY --from=builder /build/app .

# Create data directory structure only
RUN mkdir -p /data && \
    chmod 777 /data

ENTRYPOINT ["./app"]