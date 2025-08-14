# Stage 1: Build
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o dexterity

# Stage 2: Run
FROM alpine:latest

COPY --from=builder /app/dexterity /usr/local/bin/dexterity
RUN ln -s /usr/local/bin/dexterity /usr/local/bin/dex

# Install curl for debugging!
RUN apk add --no-cache curl

# Keep alive!
CMD ["tail", "-f", "/dev/null"]
