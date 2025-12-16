# Stage 1: Builder
FROM golang:alpine AS builder

WORKDIR /app

# 1. Install system dependencies (SSL Certs + TimeZone Data)
RUN apk add --no-cache git tzdata

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o finxplore-api ./cmd/server

# Stage 2: Final (The actual running container)
FROM scratch

WORKDIR /app

COPY --from=builder /app/finxplore-api .

# 2. COPY SSL Certs (Critical for calling external APIs)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 3. COPY TimeZone Data (Fixes 'unknown time zone' error)
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 8080 50051

CMD ["./finxplore-api"]