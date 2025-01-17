FROM golang:1.17 as builder

RUN apt-get update -y
RUN apt-get install -y golang npm

WORKDIR /workspace

RUN echo 2

COPY / /workspace/
RUN make build

FROM debian:11
WORKDIR /
COPY --from=builder /workspace/shopbasket .
EXPOSE 8080
ENTRYPOINT ["/shopbasket"]
