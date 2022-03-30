FROM golang:1.17 as builder

RUN apt-get update -y
RUN apt-get install -y golang npm

WORKDIR /workspace
COPY / /workspace/

RUN make build

FROM debian:11
WORKDIR /
COPY --from=builder /workspace/shopbasket .

ENTRYPOINT ["/shopbasket"]
