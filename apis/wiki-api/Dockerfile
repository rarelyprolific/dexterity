# Stage 1: Build
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o wiki-api

# Stage 2: Run
FROM alpine:latest

COPY --from=builder /app/wiki-api /app/wiki-api

# Install curl for debugging!
RUN apk add --no-cache curl

# Run API
CMD ["/app/wiki-api"]
