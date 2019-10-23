FROM golang:latest as builder

ENV PROJECT_DIR=/go/src/github.com/srijanaravali/spinnaker-servicebroker
RUN mkdir -p $PROJECT_DIR
WORKDIR $PROJECT_DIR
ARG SOURCE_DIR="./"

COPY $SOURCE_DIR .

RUN go mod download && make linux

FROM alpine:latest

RUN apk update && \
    apk add curl

# Install spin-cli
RUN curl -LO https://storage.googleapis.com/spinnaker-artifacts/spin/$(curl -s https://storage.googleapis.com/spinnaker-artifacts/spin/latest)/linux/amd64/spin \
    && chmod +x spin \
    && mv spin /usr/local/bin/spin

RUN apk add --no-cache ca-certificates bash

COPY --from=builder /go/src/github.com/srijanaravali/spinnaker-servicebroker/servicebroker-linux /usr/local/bin/spinnaker-servicebroker
COPY --from=builder /go/src/github.com/srijanaravali/spinnaker-servicebroker/scripts/start_broker.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/start_broker.sh

CMD ["start_broker.sh"]