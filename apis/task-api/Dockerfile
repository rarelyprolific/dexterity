# Stage 1: Build
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o task-api

# Stage 2: Run
FROM alpine:latest

COPY --from=builder /app/task-api /app/task-api

# Install curl for debugging!
RUN apk add --no-cache curl

# Run API
CMD ["/app/task-api"]
