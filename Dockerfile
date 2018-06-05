FROM scratch
MAINTAINER Pavol Noha <pavol.noha@gmail.com>
EXPOSE 8080
WORKDIR /
COPY agent /
ENTRYPOINT ["./agent"]
