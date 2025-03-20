FROM golang:1.19 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Use a minimal image for the final container
FROM alpine:latest
WORKDIR /root/

RUN apk --no-cache add ca-certificates

# Copy built binary from builder
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./main"]
