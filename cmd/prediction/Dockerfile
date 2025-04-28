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
    libpq-dev \
    make \
    curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list \
    apt-get update \
    apt-get install -y migrate \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /usr/local/bin/server
COPY --from=builder /app/Makefile /Makefile
COPY --from=builder /internal/prediction/infrastructure/db/migrations /migrations/

CMD ["sh", "-c", "make migrate-db && /usr/local/bin/server"]
