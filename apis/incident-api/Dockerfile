# Stage 1: Build
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o incident-api

# Stage 2: Run
FROM alpine:latest

COPY --from=builder /app/incident-api /app/incident-api

# Install curl for debugging!
RUN apk add --no-cache curl

# Run API
CMD ["/app/incident-api"]
