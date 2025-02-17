FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/ipfs-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/ipfs-svc /go/src/github.com/rarimo/ipfs-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/ipfs-svc /usr/local/bin/ipfs-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["ipfs-svc"]
