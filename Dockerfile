FROM golang:1.16.3-alpine3.13 as builder

RUN apk update && apk add gcc g++ musl-dev

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -ldflags="-w -s" -o go-puso .

FROM alpine:3.13.6

RUN apk add --no-cache libstdc++ libc6-compat

WORKDIR /app

COPY --from=builder /app/go-puso /usr/local/bin/

EXPOSE 3000

ENTRYPOINT [ "go-puso" ]
