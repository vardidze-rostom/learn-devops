FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod /app/

RUN  apk update && apk add build-base

RUN go mod download

COPY . .

RUN go build -o go_bot main.go

FROM alpine

WORKDIR /app

RUN apk update && apk upgrade

RUN adduser --disabled-password --no-create-home john-doe && chown john-doe:john-doe -R /app/

COPY --from=builder /app/go_bot /app/

USER john-doe

ENTRYPOINT ["./go_bot"]