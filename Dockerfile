FROM scratch
MAINTAINER Pavol Noha <pavol.noha@gmail.com>
WORKDIR /
COPY swarmpit-ec /
ENTRYPOINT ["./swarmpit-ec"]
