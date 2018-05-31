FROM scratch
MAINTAINER Pavol Noha <pavol.noha@gmail.com>
ENV PORT 8080
EXPOSE 8080
WORKDIR /
COPY agent /
ENTRYPOINT ["./agent"]
