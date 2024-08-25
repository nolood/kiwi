FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o start-bot ./cmd/bot/main.go

FROM postgres:latest

WORKDIR /usr/local/bin

COPY --from=builder /app/start-bot .
COPY --from=builder /app/entrypoint.dev.sh .
COPY --from=builder /app/.env ./.env
COPY --from=builder /app/config /usr/local/kiwi-config
COPY --from=builder /app/.env /usr/local/kiwi-config/.env

RUN chmod +x ./entrypoint.dev.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.dev.sh"]

EXPOSE 8080
