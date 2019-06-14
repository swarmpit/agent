FROM golang:1.12 as build
RUN mkdir temp && cd temp && go mod init . && go get -d -v github.com/Kenits/agent@v2.1.1 && go build -v github.com/Kenits/agent


FROM scratch
LABEL maintainer="deimos.kenit@gmail.com"
EXPOSE 8080
WORKDIR /
COPY --from=build go/temp/agent /
ENTRYPOINT ["./agent"]