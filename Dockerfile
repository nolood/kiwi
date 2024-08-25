FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o start-bot ./cmd/bot/main.go

FROM ubuntu:latest

WORKDIR /usr/local/bin

COPY --from=builder /app/start-bot /usr/local/bin/start-bot
COPY --from=builder /app/config /usr/local/bin/config
COPY --from=builder /app/.env /usr/local/bin/.env

RUN chmod +x /usr/local/bin/start-bot

CMD ["start-bot", "-config=./config/dev.yml"]
