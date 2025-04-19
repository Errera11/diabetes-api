FROM golang:1.24-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

#ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod tidy

COPY . .

RUN go build -o server ./cmd

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    libpq-dev \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /usr/local/bin/server

CMD ["server"]
