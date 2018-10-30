FROM golang:1.11.1-alpine3.7 AS builder

WORKDIR /go/src/github.com/calinah/efk-demo
COPY ./ .

RUN apk add --no-cache git curl

RUN go get github.com/icrowley/fake

RUN go build -o efk-demo .

FROM alpine:3.7

COPY --from=builder /go/src/github.com/calinah/efk-demo/efk-demo /

ENTRYPOINT ["/efk-demo"]
