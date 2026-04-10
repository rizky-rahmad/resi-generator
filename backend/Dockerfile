# ========================
# Stage 1: Builder
# ========================
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Hapus .env jika ikut ter-copy — paksa gunakan env var dari docker-compose
RUN rm -f .env

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app ./cmd/main.go

# ========================
# Stage 2: Runner
# ========================
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Jakarta

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]