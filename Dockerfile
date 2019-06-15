FROM scratch
LABEL maintainer="deimos.kenit@gmail.com"
EXPOSE 8080
WORKDIR /
COPY agent /
ENTRYPOINT ["./agent"]