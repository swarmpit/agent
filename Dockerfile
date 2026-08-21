# cross-compile on the native builder instead of emulating the target
FROM --platform=$BUILDPLATFORM golang:1.12 AS builder

WORKDIR /build

ENV CGO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
# explicit GOOS/GOARCH/GOARM still win, so deploy-image.sh keeps working
ARG GOOS
ARG GOARCH
ARG GOARM

COPY . /build/
RUN export GOOS="${GOOS:-${TARGETOS:-linux}}" GOARCH="${GOARCH:-${TARGETARCH:-amd64}}" && \
    if [ -z "$GOARM" ]; then \
      case "$TARGETVARIANT" in v6) GOARM=6 ;; v7) GOARM=7 ;; esac; \
    fi && \
    export GOARM && \
    echo "building for GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM" && \
    go build -a -installsuffix cgo -o agent .

FROM scratch
MAINTAINER Pavol Noha <pavol.noha@gmail.com>
EXPOSE 8080
WORKDIR /
COPY --from=builder /build/agent /
ENTRYPOINT ["./agent"]