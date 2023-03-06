FROM golang:1.20-alpine3.16 AS builder

RUN go version

COPY . /TgBot/
WORKDIR /TgBot/

RUN go mod download
RUN go build -o ./.bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /TgBot/.bin/bot .
COPY --from=0 /TgBot/configs configs/

EXPOSE 80

CMD ["./bot"]