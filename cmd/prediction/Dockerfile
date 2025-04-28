FROM golang:1.24-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod tidy

COPY . .

RUN go build -o server ./cmd

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    libpq-dev \
    make \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /usr/local/bin/server
COPY --from=builder /app/Makefile /Makefile
COPY --from=builder /internal/prediction/infrastructure/db/migrations /migrations/

CMD ["sh", "-c", "make migrate-db && /usr/local/bin/server"]
