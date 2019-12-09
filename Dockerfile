FROM golang:1.12 AS builder

WORKDIR /build

ENV CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
ARG GOARM

COPY . /build/
RUN go build -a -installsuffix cgo -o agent .

FROM scratch
MAINTAINER Pavol Noha <pavol.noha@gmail.com>
EXPOSE 8080
WORKDIR /
COPY --from=builder /build/agent /
ENTRYPOINT ["./agent"]