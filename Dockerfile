FROM golang:1.12 as build
RUN go get -d -v github.com/Kenits/agent@v.2.1.0
RUN go build -v github.com/Kenits/agent


FROM scratch
LABEL maintainer="deimos.kenit@gmail.com"
EXPOSE 8080
WORKDIR /
COPY --from=build go/agent /
ENTRYPOINT ["./agent"]