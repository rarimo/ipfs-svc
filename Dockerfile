FROM golang:1.23-alpine as buildbase

ARG CI_JOB_TOKEN

RUN apk add git build-base ca-certificates
WORKDIR /go/src/github.com/rarimo/ipfs-svc
COPY . .

RUN git config --global url."https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com".insteadOf https://gitlab.com
RUN git config --global url."https://${CI_JOB_TOKEN}@github.com/".insteadOf https://github.com/
RUN go env -w GOPRIVATE=github.com/*,gitlab.com/*

RUN go mod tidy && go mod vendor
RUN CGO_ENABLED=1 GO111MODULE=on GOOS=linux go build  -o /usr/local/bin/ipfs-svc /go/src/github.com/rarimo/ipfs-svc

FROM scratch
COPY --from=alpine:3.9 /bin/sh /bin/sh
COPY --from=alpine:3.9 /usr /usr
COPY --from=alpine:3.9 /lib /lib

COPY --from=buildbase /usr/local/bin/ipfs-svc /usr/local/bin/ipfs-svc
COPY --from=buildbase /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["ipfs-svc"]